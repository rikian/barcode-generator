tests:
	go test ./... -coverprofile coverage.out
	rm -rf coverage.out

build:
	go build -o barcode-generator .

