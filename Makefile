
default: help 

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make info              - Print info about the current setup."
	@echo "  make test              - Run E2E Test Suite using installed Task binary."
	@echo "  make install           - Install the latest Task binary."
	@echo ""
	@echo "Run a specific test: go test -v ./e2e_test.go --run TestE2E/docker/docker.taskfile.owner"

.PHONY: info
info:
	@echo "Task: $$(which task)"
	@echo "Version: $$(task --version)"
	@. /etc/os-release; echo "OS: $$(uname) $$(uname -m) $$PRETTY_NAME"

.PHONY: test 
test: info
	go test -v ./e2e_test.go

.PHONY: install
install:
