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
)

func main() {
	args := os.Args[1:]
	_ = args
	cwd, err := os.Getwd()
	if err != nil {
		// hail mary
		execMake(args, cwd)
		return
	}

	cu, err := user.Current()
	if err != nil {
		execMake(args, cwd)
		// hail mary
		return
	}
	mkDir := findMakefile(cwd, cu.HomeDir)
	if mkDir == "" {
		// hail mary
		execMake(args, cwd)
		return
	}
	execMake(args, mkDir)
	return
}

func execMake(args []string, dir string) {
	cmd := exec.Command("make", args...)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func findMakefile(start, end string) string {
	// exit if we've traversed beyond end
	if len(start) < len(end) {
		return ""
	}

	files, err := ioutil.ReadDir(start)
	if err != nil {
		return ""
	}

	for _, f := range files {
		// ignore all directories
		if f.IsDir() {
			continue
		}

		switch f.Name() {
		// from: https://www.gnu.org/software/make/manual/html_node/Makefile-Names.html
		case GNUMakefile, Makefile, makefile:
			return start
		}
	}

	idx := strings.LastIndex(start, string(os.PathSeparator))
	return findMakefile(start[:idx], end)
}
