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
		// throw hail mary and pass args to make directly
		execMake(args, cwd)
		return
	}

	cu, err := user.Current()
	if err != nil {
		// throw hail mary and pass args to make directly
		execMake(args, cwd)
		return
	}
	// TODO: wrap in a single mkPath type
	mkDir, _ := findMakefile(cwd, cu.HomeDir)
	// set mkPath as cwd for exec.Command
	if mkDir == "" {
		// hail mary on dir the cwd
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

func findMakefile(start, end string) (string, string) {
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
		// lookup order from: https://www.gnu.org/software/make/manual/html_node/Makefile-Names.html
		case GNUMakefile:
			return start, GNUMakefile
		case Makefile:
			return start, Makefile
		case makefile:
			return start, makefile
		}
	}

	idx := strings.LastIndex(start, string(os.PathSeparator))
	return findMakefile(start[:idx], end)
}
