build:
	go build -o messege cmd/*.go

run: build
	./messege ${INPUT_PATH}; rm messege

test:
	go test ./random/... -count=1

