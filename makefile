all:
	go build -ldflags="-s -w" && \
	upx --brute assuminator

release_build:
	gox --output "./bin/{{.Dir}}.{{.OS}}_{{.Arch}}" -osarch="linux/386 linux/amd64 darwin/386 darwin/amd64"

local:
	go build -ldflags="-s -w"