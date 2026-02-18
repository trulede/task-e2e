[![Sync Task Version](https://github.com/trulede/task-e2e/actions/workflows/sync-versions.yml/badge.svg)](https://github.com/trulede/task-e2e/actions/workflows/sync-versions.yml)
[![Task E2E Test](https://github.com/trulede/task-e2e/actions/workflows/test-e2e.yml/badge.svg)](https://github.com/trulede/task-e2e/actions/workflows/test-e2e.yml)

# Task E2E Test Suite

## Repo Layout

```text
task-e2e
└── .github/workflows
    └── sync-versions.yml   <-- Runs periodically to check for new Task releases.
    └── test-e2e.yml        <-- E2E Test workflow.
└── tests
    └── <dir>               <-- Collection of tests.
        └── <file>.txtar    <-- Testcase file (Testscript/txtar).
└── e2e_test.go             <-- E2E Test configuration and extensions.
└── Makefile                <-- Automation interface.
```


## Usage

### Running Tests

```bash
# Install Task (install to a directory in your PATH).
$ sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin

# Clone E2E Tests.
$ git clone https://github.com/trulede/task-e2e.git

# Run E2E Tests.
$ cd task-e2e
$ make test
```


### Writing Tests

Test are written in `Testscript` and saved in `txtar` files. Place new tests in a subdirectory of the `tests` directory, using the sub-directory structure to organize tests. The E2E Test will automatically discover and run the new tests.

```testscript
# Test: task version
task --version
cp stdout task.version
cmpenv task.version expect

-- expect --
${VERSION}
```


### Testscript Extensions

* <code>task args...</code> - Run Task!
* <code>sleep duration</code> -  Sleep for the specified duration (e.g. 1s).
* <code>touch file</code> - Touch the specified file, updating file access time. Create if the file does not exist.
* <code>[!] filecontains file1 file2|text</code> - Test that file1 contains file2 or the provided text.

## References

* [Task repo][task_repo]
* [Task Doc][task_doc]
* [Testscript][testscript]


[task_repo]: https://github.com/go-task/task
[task_doc]: https://taskfile.dev/
[testscript]: https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript
