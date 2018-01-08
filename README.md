# gmake

## Demo

![Demo](https://github.com/joshi4/gmake/raw/master/demo.gif)

## Usage

`gmake` is a wrapper around make written in Go.

`gmake` invokes make with any arguments that it was passed.

`gmake` looks for a Makefile in the current working directory, if none is found it moves up to the parent directory.

This process is repeated till a Makefile is found or till the user's home directory is reached at which point `gmake` will make a best effort call to make from the original directory in which `gmake` was called

## Install

Install `gmake` by running the following command in your shell.

~~~sh
go get github.com/joshi4/gmake
~~~

To install `go` and setup the work environment see [this link](https://golang.org/doc/install#install).

**NOTE:** Make sure you've set the `GOPATH` variable correctly :)

## Motivation

The most common mistake I make while using make ( no pun intended ) is to invoke the command from a child directory where no Makefile  is present.

I estimate I run `make` about 50 - 100 times a day and hit the above error about 90% of the time. This tool is my attempt to solve this pain point.

### Alternatives

make provides a -C flag which accepts the path to the make file but that defeats the purpose as it is more effort/time consuming to type it out.

Using your bash history reduces the pain of repeating the -C flag, however, that fails when you have to switch amongst different repos/makefiles quite frequently.

Infact, using `gmake` will make it possible to use your bash history more consistently as `gmake test` or `gmake build` are common enough and the -C flag is abstracted away

## Testing

We have a simple end to end test for gmake in CI. Please see the `.travis.yml` file
