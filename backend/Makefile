APP_NAME ?= backend
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME):$(DOCKER_TAG)
# VERIFY_IGNORE is a grep pattern to exclude files and directories from verification
VERIFY_IGNORE := /vendor\|/automock

# DIRS_TO_CHECK is a command used to determine which directories should be verified
DIRS_TO_CHECK = go list ./... | grep -v "$(VERIFY_IGNORE)"

.PHONY: fmt-local
fmt-local:
	go fmt $$($(DIRS_TO_CHECK))

# Use common makefile
include $(realpath $(shell pwd)/../)/common_makefile.mk
