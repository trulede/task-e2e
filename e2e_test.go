package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){})
}

func TestE2E(t *testing.T) {
	testDir := "tests"
	curDir := testDir
	files := []string{}
	ext := ".txtar"
	fileSystem := os.DirFS(testDir)
	fs.WalkDir(fileSystem, ".", func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return nil
		}
		if d.IsDir() {
			// Run any files from the current dir.
			if len(files) > 0 {
				t.Run(curDir, func(t *testing.T) { runTest(files, t) })
			}
			// Reset for the next dir.
			curDir = s
			files = []string{}
		} else {
			if filepath.Ext(s) == ext {
				files = append(files, filepath.Join(testDir, s))
			}
		}
		return nil
	})
	// Run the remaining/last dir/files.
	if len(files) > 0 {
		t.Run(curDir, func(t *testing.T) { runTest(files, t) })
	}
}

func runTest(files []string, t *testing.T) {
	testscript.Run(t, testscript.Params{
		Cmds: map[string]func(ts *testscript.TestScript, neg bool, args []string){
			"task":         task,
			"sleep":        sleep,
			"touch":        touch,
			"filecontains": filecontains,
		},
		Files: files,
		Setup: func(e *testscript.Env) error {
			var vars = []string{
				fmt.Sprintf("VERSION=%s", os.Getenv("VERSION")),
			}
			e.Vars = append(e.Vars, vars...)
			return nil
		},
	})
}

func task(ts *testscript.TestScript, neg bool, args []string) {
	ts.Exec("task", args...)
}

func filecontains(ts *testscript.TestScript, neg bool, args []string) {
	if len(args) != 2 {
		ts.Fatalf("filecontains <file> <file>|<text>")
	}
	got := ts.ReadFile(args[0])
	want := args[1]
	if data, err := os.ReadFile(ts.MkAbs(want)); err == nil {
		want = string(data)
	}
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
