MOD=$(shell go list .)
NAME=$(notdir $(MOD))
DIST=./dist
OUTPUT:=$(DIST)/$(NAME)
ifeq ($(OS),Windows_NT)
OUTPUT:=$(OUTPUT).exe
endif

.PHONY: run build

run: build
	$(OUTPUT) COM3 COM4

build:
	mkdir -p $(DIST)
	go build -o $(OUTPUT) .

