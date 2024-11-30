CC = clang

ROOT_DIR = ../..

BUILD_DIR_BASE = $(ROOT_DIR)/build
BUILD_DIR = $(BUILD_DIR_BASE)/go_and_c
BUILD_DIR_LIBS = $(BUILD_DIR_BASE)/libs

OUT_LIB = $(BUILD_DIR_LIBS)/libarray02.so
SRC_LIB = $(ROOT_DIR)/go_and_c/array02/array02.c
OUT = $(BUILD_DIR)/array02
SRC = $(ROOT_DIR)/go_and_c/array02

all: lib go

lib:
	$(CC) -o $(OUT_LIB) -O2 -g -shared $(SRC_LIB)

go:
	CGO_ENABLED=1 CGO_LDFLAGS="-L $(BUILD_DIR_LIBS)" go build -o $(OUT) $(SRC)

.PHONY: lib go