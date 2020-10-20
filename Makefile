# now=$()
# go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=${now}"

build:
	go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=`date +'%Y-%m-%d_%T'`" -o ./bin/ping-identity-exporter ./cmd/ping-identity-exporter

build_static:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=`date +'%Y-%m-%d_%T'`" -o ./bin/ping-identity-exporter -a -tags netgo -ldflags '-w' ./cmd/ping-identity-exporter

test:
	go test ./...
