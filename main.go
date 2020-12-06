package main

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
	"log"
)

func main() {
	// set the environment variables using $ . ./set_env_variables
	// $ go run main.go
	config := common.DefaultConfigProvider()
	tenancyId, err := config.TenancyOCID()
	FatalIfError(err)
	fmt.Println(tenancyId)

	client, err := identity.NewIdentityClientWithConfigurationProvider(config)
	FatalIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	response, err := client.ListAvailabilityDomains(context.Background(), request)
	FatalIfError(err)

	log.Printf("list of available domains: %v", response.Items)
	fmt.Println("list of available domains completed")
}

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
