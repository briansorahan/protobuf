OUTPUT = $(patsubst %.proto,%.pb.go,$(wildcard *.proto))

all: $(OUTPUT)

clean:
	@rm -f *.pb.go

%.pb.go: %.proto
	protoc -I=. -I=$(GOPATH)/src -I=../gogoproto --go_out=. $<

test: $(OUTPUT) $(wildcard *_test.go)
	@go test

.PHONY: all clean test
