PKG_NAME = lintsample
BIN_NAME = lintsample

RM = rm -rf

ifeq ($(OS),Windows_NT)
    RM = cmd.exe /c del /Q
    BIN_NAME = lintsample.exe
endif

.PHONY: build
build:
	@go build -o ${BIN_NAME} ./cmd/${PKG_NAME}/main.go

.PHONY: clean
clean:
	@$(RM) $(BIN_NAME)