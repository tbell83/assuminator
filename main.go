package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

var print = fmt.Println

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	targetRole := flag.String("target-arn", "", "ARN of the role to assume.")
	sessionToken := flag.String("session-token", "", "Session token ID")
	sessionDuration := flag.Int64("duration", 900, "Duration of assumed role session in seconds")
	flag.Parse()

	if *targetRole == "" {
		print("Target ARN not defined.")
		os.Exit(1)
	}

	if *sessionToken == "" {
		print("Session token not defined.")
		os.Exit(1)
	}

	stscreds.DefaultDuration = time.Hour
	sess := session.Must(session.NewSession())
	svc := sts.New(sess)
	params := &sts.AssumeRoleInput{
		RoleArn:         aws.String(*targetRole),
		RoleSessionName: aws.String(*sessionToken),
		DurationSeconds: aws.Int64(*sessionDuration),
	}
	resp, err := svc.AssumeRole(params)
	check(err)

	currentDir, err := os.Getwd()
	check(err)
	file, err := os.Create(currentDir + fmt.Sprintf("/%s_creds", *sessionToken))
	check(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("export AWS_ACCESS_KEY_ID=\"%s\"\n", *resp.Credentials.AccessKeyId))
	writer.WriteString(fmt.Sprintf("export AWS_SESSION_TOKEN=\"%s\"\n", *resp.Credentials.SessionToken))
	writer.WriteString(fmt.Sprintf("export AWS_SECRET_ACCESS_KEY=\"%s\"\n", *resp.Credentials.SecretAccessKey))
	writer.Flush()
}
