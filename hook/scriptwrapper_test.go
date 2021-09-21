package hook

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/buildkite/agent/v3/bootstrap/shell"
	"github.com/buildkite/agent/v3/env"
	"github.com/stretchr/testify/assert"
)

func TestRunningHookDetectsChangedEnvironment(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	var script []string

	if runtime.GOOS != "windows" {
		script = []string{
			"#!/bin/bash",
			"export LLAMAS=rock",
			"export Alpacas=\"are ok\"",
			"echo hello world",
		}
	} else {
		script = []string{
			"@echo off",
			"set LLAMAS=rock",
			"set Alpacas=are ok",
			"echo hello world",
		}
	}

	wrapper := newTestScriptWrapper(t, script)
	defer os.Remove(wrapper.Path())

	sh := shell.NewTestShell(t)

	if err := sh.RunScript(ctx, wrapper.Path(), nil); err != nil {
		t.Fatal(err)
	}

	changes, err := wrapper.Changes()
	if err != nil {
		t.Fatal(err)
	}

	// The strict equals check here also ensures we aren't bubbling up the
	// internal BUILDKITE_HOOK_EXIT_STATUS and BUILDKITE_HOOK_WORKING_DIR
	// environment variables
	assert.Equal(t, env.Diff {
		Added: map[string]string {
			"LLAMAS": "rock",
			"Alpacas": "are ok",
		},
		Changed: map[string]env.Pair{},
		Removed: map[string]struct{}{},
	}, changes.Diff)
}

func TestRunningHookDetectsChangedWorkingDirectory(t *testing.T) {
	t.Parallel()

	tempDir, err := ioutil.TempDir("", "hookwrapperdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	ctx := context.Background()
	var script []string

	if runtime.GOOS != "windows" {
		script = []string{
			"#!/bin/bash",
			"mkdir mysubdir",
			"cd mysubdir",
			"echo hello world",
		}
	} else {
		script = []string{
			"@echo off",
			"mkdir mysubdir",
			"cd mysubdir",
			"echo hello world",
		}
	}

	wrapper := newTestScriptWrapper(t, script)
	defer os.Remove(wrapper.Path())

	sh := shell.NewTestShell(t)
	if err := sh.Chdir(tempDir); err != nil {
		t.Fatal(err)
	}

	if err := sh.RunScript(ctx, wrapper.Path(), nil); err != nil {
		t.Fatal(err)
	}

	changes, err := wrapper.Changes()
	if err != nil {
		t.Fatal(err)
	}

	expected, err := filepath.EvalSymlinks(filepath.Join(tempDir, "mysubdir"))
	if err != nil {
		t.Fatal(err)
	}

	changesDir, err := filepath.EvalSymlinks(changes.Dir)
	if err != nil {
		t.Fatal(err)
	}

	if changesDir != expected {
		t.Fatalf("Expected working dir of %q, got %q", expected, changesDir)
	}
}

func newTestScriptWrapper(t *testing.T, script []string) *ScriptWrapper {
	hookName := "hookwrapper"
	if runtime.GOOS == "windows" {
		hookName += ".bat"
	}

	hookFile, err := shell.TempFileWithExtension(hookName)
	if err != nil {
		t.Fatal(err)
	}

	for _, line := range script {
		if _, err = fmt.Fprintln(hookFile, line); err != nil {
			t.Fatal(err)
		}
	}

	hookFile.Close()

	wrapper, err := CreateScriptWrapper(hookFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	return wrapper
}
