
BINARY=thinkport_darwin_arm64

build:
	@echo "🔧 Building binary..."
	goreleaser build --snapshot --clean

build-test: build
	@echo "🔧 Test the binary"
	chmod +x ./dist/$(BINARY)/thinkport
	./dist/$(BINARY)/thinkport --version

release:
	@echo "🔧 Release binary …"
	goreleaser release --clean
