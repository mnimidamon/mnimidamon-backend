# ![md-s-icon](./public/mnimidamon-server-icon.svg) mnimidamon server
Backend for _mnimidamon_ written in Go.

## Running

The _mnimidamon_ server can be started with this command:

``go run ./cmd/mnimidamon-server/main.go``

or build the executable and run it later:

``go build ./cmd/mnimidamon-server/main.go``

## OpenAPI code generation

The server uses HTTP communications to exchange JSON entities. The endpoints follow REST specifications
which are defined in `./public/spec/swagger.yml`. Client and Server communication code is generated by the
`go-swagger` toolkit. If the specifications change, run the following commands respectfully.

Generate endpoints code:

``swagger generate server -A mnimidamon -f ./public/spec/swagger.yaml -C swagger-layout.yml -t ./adapter/restapi -s endpoints -m modelapi
``

Generate clients code:

``swagger generate client -A mnimidamon -f ./public/spec/swagger.yaml ``

## Installation

Generate GUI mnimidamon server .exe file for windows:

`` go build -ldflags -H=windowsgui cmd/mnimidamon-server-gui/main.go ``

For other platforms:

`` go build cmd/mnimidamon-server-gui/main.go ``

Generate command line mnimidamon server .exe file:

`` go build cmd/mnimidamon-server/main.go ``

## Application sneak peek

![configure](https://i.imgur.com/6M7qEHa.png)
![start](https://i.imgur.com/bdHWFtT.png)
![started](https://i.imgur.com/CTOM97I.png)
![reconfigure](https://i.imgur.com/b25uPsL.png)
