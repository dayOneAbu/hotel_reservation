build:
	@go build -o bin/hotelReservation.exe

run: build
	@./bin/hotelReservation.exe

seed:
	@ go run scripts/seed.go

test:
	@go test -v ./...
