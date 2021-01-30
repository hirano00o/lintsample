BIN_NAME = lintsample

DIST_DIR = dist
RM = rm -rf

ifeq ($(OS),Windows_NT)
    RM = rd /Q
endif

.PHONY: dist-dir
dist-dir:
    @mkdir dist > NUL 2>&1

.PHONY: build
build: dist-dir
    @go build -o ${DIST_DIR}/${BIN_NAME} cmd/${BIN_NAME}/main.go

.PHONY: clean
clean:
    -$(RM) $(DIST_DIR)