all:
	go build -ldflags="-s -w" && \
	upx --brute assuminator

release_build:
	gox -os="linux";
	for BIN in $(ls ./bin/); do upx --brute "./bin/${BIN}"; done

local:
	go build -ldflags="-s -w"