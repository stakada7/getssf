VERSION=0.1.0
TARGETS_NOVENDOR=$(shell glide novendor)

all: bin/getssf

build-cross: cmd/getssf/getssf.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/getssf-${VERSION}/getssf cmd/getssf/getssf.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/getssf-${VERSION}/getssf cmd/getssf/getssf.go

dist: build-cross
	cd bin/linux/amd64 && tar zcvf getssf-linux-amd64-${VERSION}.tar.gz getssf-${VERSION}
	cd bin/darwin/amd64 && tar zcvf getssf-darwin-amd64-${VERSION}.tar.gz getssf-${VERSION}

bin/getssf: cmd/getssf/getssf.go
	go build -o bin/getssf cmd/getssf/getssf.go

bundle:
	glide install

fmt:
	@echo $(TARGETS_NOVENDOR) | xargs go fmt

check:
	go test -v $(TARGETS_NOVENDOR)

clean:
	rm -rf bin/*
