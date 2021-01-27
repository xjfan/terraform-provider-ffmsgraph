HOSTNAME=hashicorp.com
NAMESPACE=xjfan
NAME=ffmsgraph
BINARY=terraform-provider-${NAME}
VERSION=1.0
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

build-linux:
	GOOS=linux GOARCH=amd64 go build -o plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/linux_amd64/${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
