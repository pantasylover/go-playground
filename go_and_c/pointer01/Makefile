CC = clang

ROOT_DIR = ../..

BUILD_DIR = $(ROOT_DIR)/build
BUILD_DIR_LIBS = $(BUILD_DIR)/libs
BUILD_DIR_GO_AND_C = $(BUILD_DIR)/go_and_c

OUT_LIB = $(BUILD_DIR_LIBS)/libpointer01.so
SRC_LIB = $(ROOT_DIR)/go_and_c/pointer01/pointer01.c
OUT = $(BUILD_DIR_GO_AND_C)/pointer01
SRC = $(ROOT_DIR)/go_and_c/pointer01

all: lib go

lib:
	$(CC) -o $(OUT_LIB) -O2 -g -shared $(SRC_LIB)

go:
	CGO_ENABLED=1 CGO_LDFLAGS="-L $(BUILD_DIR_LIBS)" go build -o $(OUT) $(SRC)

.PHONY: lib go