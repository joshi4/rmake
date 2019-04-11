package main

import "testing"

func TestFindMakefileWithNoPathSeparator(t *testing.T) {
	if dir, _ := findMakefile("test", ""); dir != "" {
		t.Errorf("unexpected dir=%s, dir with no PathSeparator should return empty string", dir)
	}
}
