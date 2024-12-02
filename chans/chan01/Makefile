CC = clang

ROOT_DIR = ../..

BUILD_DIR_BASE = $(ROOT_DIR)/build
BUILD_DIR = $(BUILD_DIR_BASE)/chans

OUT = $(BUILD_DIR)/chan01
SRC = $(ROOT_DIR)/chans/chan01

all: go

go:
	go build -o $(OUT) $(SRC)

.PHONY: go