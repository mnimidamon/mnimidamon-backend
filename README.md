# mnimidamon-backend
Backend for mnimidamon written in Go.


## Running

The _mnimidamon_ server can be started with this command:

``go run ./cmd/mnimidamon-server/main.go``

or build the executable and run it later:

``go build ./cmd/mnimidamon-server/main.go``

## OpenAPI specific

Generate endpoints code:

``swagger generate server -A mnimidamon -f ./public/spec/swagger.yaml ``

Generate clients code:

``swagger generate client -A mnimidamon -f ./public/spec/swagger.yaml ``

