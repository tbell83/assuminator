all:
	go build -ldflags="-s -w" && \
	upx --brute assuminator