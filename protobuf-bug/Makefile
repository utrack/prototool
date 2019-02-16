.DEFAULT_GOAL := bazel

BAZEL_VERSION ?= 0.22.0

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
TMP_BASE := .tmp
TMP := $(TMP_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
TMP_BIN = $(TMP)/bin
TMP_BAZEL := $(TMP)/bazel-$(BAZEL_VERSION)
BAZEL := $(abspath $(TMP_BAZEL)/bin/bazel)

ifeq ($(UNAME_OS),Darwin)
BAZEL_OS := darwin
else
BAZEL_OS = linux
endif
BAZEL_ARCH := $(UNAME_ARCH)
$(BAZEL):
	rm -f $(TMP_BIN)/bazel
	rm -rf $(TMP_BAZEL)
	mkdir -p $(TMP_BIN)
	mkdir -p $(TMP_BAZEL)
	curl -SSL https://github.com/bazelbuild/bazel/releases/download/$(BAZEL_VERSION)/bazel-$(BAZEL_VERSION)-installer-$(BAZEL_OS)-$(BAZEL_ARCH).sh -o $(TMP_BAZEL)/bazel-installer.sh
	chmod +x $(TMP_BAZEL)/bazel-installer.sh
	$(TMP_BAZEL)/bazel-installer.sh --base=$(abspath $(TMP_BAZEL)) --bin=$(abspath $(TMP_BIN))

.PHONY: bazel
bazel: $(BAZEL)
	$(BAZEL) build //...:all
