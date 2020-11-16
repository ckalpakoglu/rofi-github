VERSION         := v0.1
BUILD_NAME      := rofi-github
BUILD_DIR       := _release
BUILD_TIME      := $(shell date +%FT%T%z)
GO_LDFLAGS      := -ldflags="-s -w -X main.Version=${VERSION}"

# checksum
SHASUMCMD       := $(shell command -v sha256sum || command -v shasum; 2> /dev/null)

.PHONY: all clean

export GO111MODULE=on

all: linux
    @echo Version: $(VERSION)

linux:
	GOOS=linux GOARCH=amd64 go build -v ${GO_LDFLAGS} -o ${BUILD_DIR}/${BUILD_NAME}
	${SHASUMCMD} ${BUILD_DIR}/${BUILD_NAME} | cut -d ' ' -f1 > ${BUILD_DIR}/${BUILD_NAME}.sha256

clean:
	rm -rf ${BUILD_DIR}
