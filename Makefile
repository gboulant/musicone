all: test

build:
	@go build

test: build
	@./musicone

clean:
	@go clean
	@rm -f *~
