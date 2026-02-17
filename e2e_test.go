package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/rogpeppe/go-internal/testscript"
)

var tests = map[string]string{
	"CLI":   filepath.Join("tests", "cli"),
	"Watch": filepath.Join("tests", "watch"),
}

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){})
}

func TestE2E(t *testing.T) {
	for test, path := range tests {
		t.Run(test, func(t *testing.T) { runTest(path, t) })
	}
}

func runTest(dir string, t *testing.T) {
	files := []string{}
	ext := ".txtar"
	fileSystem := os.DirFS(dir)
	fs.WalkDir(fileSystem, ".", func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return nil
		}
		if !d.IsDir() && filepath.Ext(s) == ext {
			files = append(files, filepath.Join(dir, s))
		}
		return nil
	})
	testscript.Run(t, testscript.Params{
		Files: files,
		Cmds: map[string]func(ts *testscript.TestScript, neg bool, args []string){
			"task":         task,
			"sleep":        sleep,
			"touch":        touch,
			"filecontains": filecontains,
		},
	})
}

func task(ts *testscript.TestScript, neg bool, args []string) {
	ts.Exec("task", args...)
}

func filecontains(ts *testscript.TestScript, neg bool, args []string) {
	if len(args) != 2 {
		ts.Fatalf("filecontains <file> <text>")
	}
	got := ts.ReadFile(args[0])
	want := args[1]
	if strings.Contains(got, want) == neg {
		ts.Fatalf("filecontains %q; %q not found in file:\n%q", args[0], want, got)
	}
}

func sleep(ts *testscript.TestScript, neg bool, args []string) {
	duration := time.Second
	if len(args) == 1 {
		d, err := time.ParseDuration(args[0])
		ts.Check(err)
		duration = d
	}
	time.Sleep(duration)
}

func touch(ts *testscript.TestScript, neg bool, args []string) {
	if len(args) != 1 {
		ts.Fatalf("touch <file>")
	}
	// Get the relative path to the scripts current directory.
	path := ts.MkAbs(args[0])
	// Create the file (if necessary).
	err := os.MkdirAll(filepath.Dir(path), 0750)
	ts.Check(err)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	ts.Check(err)
	err = file.Close()
	ts.Check(err)
	// Now update the timestamp.
	t := time.Now()
	err = os.Chtimes(path, t, t)
	ts.Check(err)
}
