build:
	env GOOS=darwin GOARCH=amd64 go build -o scaffolder ../main.go

gen-default:
	./scaffolder gen -o="./project" -d="github.com" -n="holmanskih/project"

gen-with-gomod:
	./scaffolder gen -o="./project" -d="github.com" -n="holmanskih/project" -p="github.com/holmanskih/project"

clean:
	rm -r ./project && rm ./scaffolder

test:
	make build gen-with-gomod