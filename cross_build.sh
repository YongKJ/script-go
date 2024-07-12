export GOARCH=amd64
export GOOS=windows
export CGO_ENABLED=1

rm -rf ../dist

go build -o ../dist/script-go.exe
