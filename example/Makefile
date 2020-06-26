build:
	env GOOS=darwin GOARCH=amd64 go build -o scaffolder ../main.go

gen-default:
	./scaffolder gen -o="./project" -d="github.com" -n="holmanskih/project"

gen-with-gomod-simple:
	./scaffolder gen -o="./project" -d="github.com" -n="holmanskih/project" -p="github.com/holmanskih/project"

gen-with-gomod-api:
	./scaffolder gen -o="./project" -d="github.com" -n="holmanskih/project" -p="github.com/holmanskih/project" --api

gen-with-gomod-db:
	./scaffolder gen -o="./project" -d="github.com" -n="holmanskih/project" -p="github.com/holmanskih/project" --db

gen:
	./scaffolder gen -o="./project" -n="holmanskih/project" --gomods

clean:
	rm -r ./project && rm ./scaffolder

test:
	make build gen