
build-darwin:
  @echo "Building for Darwin (MacOS): amd64 (64 bit)"
  @env GOOS=darwin GOARCH=amd64 go build -ldflags='-extldflags=-static' -o build/smoke_x64-darwin main.go

build-linux:
  @echo "Building for Linux: amd64 (64 bit) & i386 (32 bit)"
  @env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags='-extldflags=-static' -o build/smoke_x64-linux main.go
  @env GOOS=linux GOARCH=386 go build -ldflags='-extldflags=-static' -o build/smoke_i386-linux main.go

build-windows:
  @echo "Building for Windows: amd64 (64 bit) & i386 (32 bit)"
  @env GOOS=windows GOARCH=amd64 go build -ldflags='-extldflags=-static' -o build/smoke_amd64-windows.exe main.go
  @env GOOS=windows GOARCH=386 go build -ldflags='-extldflags=-static' -o build/smoke_i386-windows.exe main.go

clean:
  @rm -rf build

build: clean build-darwin build-linux build-windows
  @sha256sum build/smoke_x64-darwin >> build/BUILD_LOG
  @sha256sum build/smoke_x64-linux >> build/BUILD_LOG
  @sha256sum build/smoke_i386-linux >> build/BUILD_LOG
  @sha256sum build/smoke_amd64-windows.exe >> build/BUILD_LOG
  @sha256sum build/smoke_i386-windows.exe >> build/BUILD_LOG
