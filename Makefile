run:
	go run ./main

bench:
	go test -bench=. main

test:
	go test main