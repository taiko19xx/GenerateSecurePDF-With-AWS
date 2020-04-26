.PHONY: all build clean resource

setup: build resource

build:
	cd copyobject && env GOOS=linux go build -ldflags="-s -w" -o ../bin/copyobject/main main.go
	cd encrypt && env GOOS=linux go build -ldflags="-s -w" -o ../bin/encrypt/main main.go
	cd property && env GOOS=linux go build -ldflags="-s -w" -o ../bin/property/main main.go
	cd sign && env GOOS=linux go build -ldflags="-s -w" -o ../bin/sign/main main.go
	cd watermark && env GOOS=linux go build -ldflags="-s -w" -o ../bin/watermark/main main.go

clean:
	rm -rf ./bin

resource:
	cp sign/TanukiMagic.ttf bin/sign/