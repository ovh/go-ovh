/*
Package metrics provides easy to use methods over /metrics OVH API.

This package expose two ways to use OVH Metrics API.

The first one with raw functions calls, the second with a more "struct and methods style".

Raw usage

List Metrics services:
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	services, err := List(client)
	fmt.Printf("Metrics services: %+v", services)

	// Output:
	// ["xmvoqfviqqmf", "baliuvlqbvih", "uvqsdbviuqgh"]
	//

Get a token:
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	token, err := GetToken(client, "xmvoqfviqqmf", "cfa8f6e7-77a4-48fb-b446-cdf82749ac04")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	fmt.Printf("Token: %+v", token)

	// Output:
	// {ID: Access: Permission: Description: Labels: IsRevoked: CreatedAt: ExpireAt}
	//

Metric client usage

List Metrics services:
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	metricsClient : metrics.NewClient(client)

	services := metricsClient.Services()
	fmt.Printf("Services: %+v", services)

	// Output:
	//
	//

Get a token:
	service, err := metricsClient.Service("xmvoqfviqqmf")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	token, err := service.Token("cfa8f6e7-77a4-48fb-b446-cdf82749ac04")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	fmt.Printf("Token: %+v", token)
*/
package metrics
