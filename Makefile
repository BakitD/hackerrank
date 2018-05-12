GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
MKDIR=mkdir -p

SRC_ROOT=$(shell pwd)
SRC_TEST=tests
SRC=src/tasks/main/main.go
SRC_OBJ=main

test:
	$(GOTEST) $(SRC_TEST)

clean:
	@[ -f ./$(SRC) ] && rm ./$(SRC_OBJ) || true
