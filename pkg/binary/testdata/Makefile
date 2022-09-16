
helloworld-%: helloworld.go
	GOOS=$$(echo $* | cut -f1 -d-) GOARCH=$$(echo $* | cut -f2 -d-) go build -o $@ $^

all: helloworld-linux-amd64 helloworld-linux-arm64 helloworld-darwin-arm64
