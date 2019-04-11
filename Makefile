build:
	go build -o test/nested/dir/rmake github.com/joshi4/rmake

test: build
	go test -v ./...
