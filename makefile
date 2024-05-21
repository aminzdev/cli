compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/zdevs main.go

.PHONY: compile