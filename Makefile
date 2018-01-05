
GIT_HASH=`git rev-parse HEAD`

all: version
	@go get ./cli/...
	@go build -i -o ./turtl ./cli/*.go

install: version
	@go install 
	@go get ./cli/...
	@go build -o $(GOPATH)/bin/turtl ./cli

libturtl: version
	@go install 

proto:
	@protoc -I . turtl.proto --gofast_out=plugins=grpc:cli

version:
	@echo "package turtl\n\nvar gitHash = \"$(GIT_HASH)\"\n" > githash.go

test:
	@go test -run '' -v

clean:
	@go clean -i github.com/andreiamatuni/
	@rm -f ./turtl
	@rm -f $(GOPATH)/bin/turtl
	