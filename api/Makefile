SERVICE_NAME=api
MY_PACKAGES=$(shell find . -type d -name vendor -prune -o -name "*.go" -print | xargs -I{} dirname {} | sort | uniq)

.PHONY: fmt, gen

fmt:
	go fmt ${MY_PACKAGES}

gen:
	go generate ${MY_PACKAGES}
