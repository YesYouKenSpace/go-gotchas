SHELL := /bin/bash

all: fmt run_all clean_stack_traces


fmt:
	@echo "Running go fmt on all Go files"
	@find . -name "main.go" -execdir bash -c 'go fmt main.go' \;
	
run_all:
	@echo "Running all main.go files and saving output to sample.out.draft in their respective directories"
	@for dir in $(shell find . -name "main.go" -exec dirname {} \;); do \
		$(MAKE) run path=$$dir; \
	done


run: run_path clean_stack_traces_path

run_path:
	@echo "Running main.go file and saving output to sample.out.draft in specified directory: $(path)"
	@if [ -n "$(path)" ]; then \
		(cd $(path) && go run main.go > sample.out.draft 2>&1 || true); \
	else \
		echo "Please specify a path using 'make run_path path=<directory>'"; \
	fi

clean_stack_traces_path:
	@echo "Cleaning stack traces from sample.out file in specified directory: $(path)"
	@if [ -n "$(path)" ]; then \
		(cd $(path) && sed -E "/^(\s|\t)+\\//d" sample.out.draft > sample.out); \
	else \
		echo "Please specify a path using 'make clean_stack_traces_path path=<directory>'"; \
	fi

