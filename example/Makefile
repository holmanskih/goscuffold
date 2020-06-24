build:
	env GOOS=darwin GOARCH=amd64 go build -o scaffolder ../main.go

gen-default:
	./scaffolder gen --tmpl="./testdata" -d="github.com" -n="holmanskih/project" -o="./project"

test:
	make build gen-default