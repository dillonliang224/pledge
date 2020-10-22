clean:
	bazel clean
	rm -rf dockerbuild

PKG_LIST := $(shell go list ./... | grep -v /vendor/)

lint:
	golint ${PKG_LIST}

vet:
	go vet ${PKG_LIST}

doc:
	swagger generate spec