//go:build cgo && onnxruntime

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"path"
	"github.com/google/go-cmp/cmp"
)

func getProjectRoot() string {
	// Get the absolute path to the project root dynamically
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Traverse up to the root of the project
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			panic("could not find project root")
		}
		dir = parent
	}
}

func TestCLI(t *testing.T) {
	const basicDir = "../tests_data/basic"
	var (
		files = []string{
			path.Join(basicDir, "python/code.py"),
			path.Join(basicDir, "zip/magika_test.zip"),
		}
		b strings.Builder
	)
	if err := cli(&b, files...); err != nil {
		t.Fatal(err)
	}
	if d := cmp.Diff(strings.TrimSpace(b.String()), strings.Join([]string{
		"../tests_data/basic/python/code.py: python",
		"../tests_data/basic/zip/magika_test.zip: zip",
	}, "\n")); d != "" {
		t.Errorf("mismatch (-want +got):\n%s", d)
	}
}

func TestCLI2(t *testing.T) {
	rootDir := getProjectRoot()
	basicDir := filepath.Join(rootDir, "cli/tests_data")

	var (
		files = []string{
			filepath.Join(basicDir, "magika_test_pptx.txt"),
			filepath.Join(basicDir, "magika_test.zip"),
		}
		b strings.Builder
	)

	if err := cli(&b, files...); err != nil {
		t.Fatal(err)
	}

	want := strings.Join([]string{
		filepath.Join(basicDir, "magika_test_pptx.txt") + ": txt",
		filepath.Join(basicDir, "magika_test.zip") + ": zip",
	}, "\n")

	if d := cmp.Diff(strings.TrimSpace(b.String()), want); d != "" {
		t.Errorf("mismatch (-want +got):\n%s", d)
	}
}


