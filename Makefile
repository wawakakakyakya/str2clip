# /bin/sh

fileName := str2clip

# all: all-test export-all-coverage

mod_tidy:
	go mod tidy
build_linux: clean
	GOOS=linux GOARCH=amd64 go build -o ${fileName} main.go
build_windows: clean
	GOOS=windows GOARCH=amd64 go build -o ${fileName}.exe main.go
clean:
	rm -f ${fileName}
test:
	./${filename}
