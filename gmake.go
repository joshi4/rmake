package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

const (
	GNUMakefile = "GNUMakefile"
	Makefile    = "Makefile"
	makefile    = "makefile"
	makeSh      = "make.sh"
)

var fileToCommand = map[string]string{
	GNUMakefile: "make",
	Makefile:    "make",
	makefile:    "make",
	makeSh:      "./make.sh",
}

func main() {
	args := os.Args[1:]

	cwd, err := os.Getwd()
	if err != nil {
		// hail mary
		execMake(args, cwd, "")
		return
	}

	cu, err := user.Current()
	if err != nil {
		execMake(args, cwd, "")
		return
	}

	mkDir, fname := findMakefile(cwd, cu.HomeDir)
	if mkDir == "" {
		execMake(args, cwd, fname)
		return
	}
	execMake(args, mkDir, fname)
	return
}

func execMake(args []string, dir, fname string) {
	c := fileToCommand[fname]
	if c == "" {
		c = "make"
	}
	cmd := exec.Command(c, args...)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func findMakefile(start, end string) (string, string) {
	// exit if we've traversed beyond end
	if len(start) < len(end) {
		return "", ""
	}

	files, err := ioutil.ReadDir(start)
	if err != nil {
		return "", ""
	}

	for _, f := range files {
		// ignore all directories
		if f.IsDir() {
			continue
		}

		switch f.Name() {
		// from: https://www.gnu.org/software/make/manual/html_node/Makefile-Names.html
		case GNUMakefile:
			return start, GNUMakefile
		case Makefile:
			return start, Makefile
		case makefile:
			return start, makefile
		case makeSh:
			return start, makeSh
		}
	}

	idx := strings.LastIndex(start, string(os.PathSeparator))
	// be defensive
	if idx == -1 {
		return "", ""
	}
	return findMakefile(start[:idx], end)
}
