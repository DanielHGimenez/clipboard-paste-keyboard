.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -o bin/cbpk-amd64.exe # 64-bit
	GOOS=windows GOARCH=386 go build -o bin/cbpk-386.exe # 32-bit

	#GOOS=darwin GOARCH=amd64 go build -o bin/cbpk-amd64-darwin # 64-bit
	#GOOS=darwin GOARCH=386 go build -o bin/cbpk-386-darwin # 32-bit
	#GOOS=darwin GOARCH=arm64 go build -o bin/cbpk-arm64-darwin # Apple Silicon

	GOOS=linux GOARCH=amd64 go build -o bin/cbpk-amd64-linux # 64-bit
	#GOOS=linux GOARCH=386 go build -o bin/cbpk-386-linux # 32-bit
