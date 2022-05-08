#BUILD_ARCH=arm64
#BUILD_OS=darwin


run:
	docker run --name ggore --rm --env-file .env --volume $(PWD)/data:/app/data:rw plin2k/ggore-tg-channels

	#go run cmd/app/main.go

build:
	docker build -t plin2k/ggore-tg-channels .
	#go build go build -o bin/app_main_darwin_arm64 cmd/app/main.go

stop:
	docker stop ggore-tg-channels
	#go build go build -o bin/app_main_darwin_arm64 cmd/app/main.go

compile:
	GOOS=darwin GOARCH=arm64 go build -o bin/app_main_darwin_arm64 cmd/app/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/app_main_darwin_amd64 cmd/app/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/app_main_linux_amd64 cmd/app/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/app_main_linux_arm64 cmd/app/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/app_main_windows_amd64 cmd/app/main.go

