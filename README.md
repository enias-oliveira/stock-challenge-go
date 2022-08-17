# stock-challenge-go

Using go version 1.18

Recommended way to run is using [Air](https://github.com/cosmtrek/air) for Live reload. 

Will add more to this README ASAP. 

Swagger Endpoint: http://localhost:8080/swagger/index.html

Dont forget to convert the .env-example to .env 

To run it using Nix:
- First [install Nix](https://nixos.org/download.html) (It's awesome).
- On the project root, enter a nix-shell with .... `nix-shell`
- Once in the nix-shell, run `setup_db`, in order to create the database
- Get all dependecies with `go mod tidy`
- Finally, run `go run cmd/api/main.go` or `air` as recommended
