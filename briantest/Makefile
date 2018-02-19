OUTPUT = $(patsubst %.proto,%.pb.go,$(wildcard *.proto))

all: $(OUTPUT)

%.pb.go: %.proto
	protoc -I=. -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --go_out=. $<
