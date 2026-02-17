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

```bash
# Install Task (install to a directory in your PATH).
$ sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin

# Clone E2E Tests.
$ git clone https://github.com/trulede/task-e2e.git

# Run E2E Tests.
$ cd task-e2e
$ make test
```


## References

* [Task repo][task_repo]
* [Task Doc][task_doc]
* [Testscript][testscript]


[task_repo]: https://github.com/go-task/task
[task_doc]: https://taskfile.dev/
[testscript]: https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript
