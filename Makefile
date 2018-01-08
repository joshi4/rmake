build:
	go build -o test/nested/dir/gmake github.com/joshi4/gmake

test: build
	go test -v ./...
