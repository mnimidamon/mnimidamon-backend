package factory

import (
	"fmt"
	"github.com/go-openapi/loads"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints"
	"mnimidamonbackend/adapter/restapi/endpoints/operations"
)

func NewServer() (*endpoints.Server, error) {
	swaggerSpec, err := loads.Analyzed(endpoints.SwaggerJSON, "")
	if err != nil {
		return nil, fmt.Errorf("swagger spec loading error %w", err)
	}

	api := operations.NewMnimidamonAPI(swaggerSpec)
	server := endpoints.NewServer(api)

	server.Port = restapi.GlobalConfig.Port
	server.ConfigureAPI()

	return server, nil
}
