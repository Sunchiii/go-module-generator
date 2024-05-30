:: Build the Go binary
go build -o go-gen ./cmd/generate/main.go

:: Move the binary to a directory in your PATH
move go-gen C:\Windows\System32\go-gen
