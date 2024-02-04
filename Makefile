run:
	templ generate
	go run . runserver
prebuild:
	templ generate
	yarn build
build:
	make prebuild
	go build -ldflags="-w -s" -o gostart
linux:
	make prebuild
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o linuxgostart
