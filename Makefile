GOCMD=go
GOBUILD=$(GOCMD) build
MKDIR=mkdir -p

SRC_ROOT=$(shell pwd)
SRC=src/tasks/main/main.go
SRC_OBJ=main

build:
	$(GOBUILD) -o $(SRC_OBJ) ${GOPATH}/$(SRC)

clean:
	@[ -f ./$(SRC) ] && rm ./$(SRC_OBJ) || true
