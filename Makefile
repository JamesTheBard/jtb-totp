PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

release: $(PLATFORMS)

clean:
	rm -rf builds

$(PLATFORMS):
	@printf "Building platform $@..."
	@if [ "$(os)" = "windows" ]; then \
		GOOS=$(os) GOARCH=$(arch) go build -o 'builds/jtb-totp_$(os)_$(arch).exe' ; \
	else \
		GOOS=$(os) GOARCH=$(arch) go build -o 'builds/jtb-totp_$(os)_$(arch)' ; \
	fi
	@echo "done!"