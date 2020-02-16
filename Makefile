run:
	go run main.go

bench:
	go test -bench=. main

test:
	go test main