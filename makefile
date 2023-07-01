build:
	@go build -o bin/hotelReservation.exe

run: build
	@./bin/hotelReservation.exe

test:
	@go test -v ./...
