CC = clang

ROOT_DIR = ../..

BUILD_DIR_BASE = $(ROOT_DIR)/build
BUILD_DIR = $(BUILD_DIR_BASE)/structs

OUT = $(BUILD_DIR)/interface_to_slice01
SRC = $(ROOT_DIR)/structs/interface_to_slice01

all: go

go:
	go build -o $(OUT) $(SRC)

.PHONY: go