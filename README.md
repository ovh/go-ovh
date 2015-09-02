govh
======

Simple Go wrapper for the OVH API.

## Register your app

Visit [https://eu.api.ovh.com/createApp](https://eu.api.ovh.com/createApp) and create your app. You'll get an application key and an application secret. To use the API you'll need a consumer key.

The consumer key has two types of restriction:

* path: eg. only the ```GET``` method on ```/me```
* time: eg. expire in 1 day


## Get a consumer key

Here's an example on how to get your consumer key.

```go
package main

import (
    "fmt"

    "github.com/gregdel/govh"
)

func main() {
    appKey := "you_app_key"
    ckReq := govh.NewCkRequest(govh.OvhEU, appKey)

    // Allow GET method on /me
    ckReq.AddRule("GET", "/me")

    // Allow GET method on /xdsl and all its sub routes
    ckReq.AddRule("GET", "/xdsl/*")

    // Run the request
    response, err := ckReq.Do()
    if err != nil {
        fmt.Printf("Error: %q\n", err)
        return
    }

    // Print the validation URL
    fmt.Printf("%s",response)
}
```

## Use the lib

### GET

```go
package main

import (
	"fmt"

	"github.com/gregdel/govh"
)

func main() {
	ak := "your_app_key"
	as := "your_app_secret"
	ck := "your_consumer_key"

	client := govh.NewClient(govh.OvhEU, ak, as, ck)

	// Get all the xdsl services
	xdslServices := []string{}
	if err := client.Get("/xdsl/", &xdslServices); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	// xdslAccess represents a xdsl access returned by the API
	type xdslAccess struct {
		Name   string `json:"accessName"`
		Status string `json:"status"`
		Pairs  int    `json:"pairsNumber"`
		// Insert the other properties here
	}

        // Get the details of each service
	for i, serviceName := range xdslServices {
		access := xdslAccess{}
		url := "/xdsl/" + serviceName

		if err := client.Get(url, &access); err != nil {
			fmt.Printf("Error: %q\n", err)
			return
		}
		fmt.Printf("#%d : %+v\n", i+1, access)
	}
}
```

### PUT

```go
package main

import (
	"fmt"

	"github.com/gregdel/govh"
)

func main() {
	ak := "your_app_key"
	as := "your_app_secret"
	ck := "your_consumer_key"

	client := govh.NewClient(govh.OvhEU, ak, as, ck)

	// Params
	type AccessPutParams struct {
		Description string `json:"description"`
	}

	// Update the description of the service
	params := &AccessPutParams{Description: "My awesome access"}
	if err := client.Put("/xdsl/xdsl-yourservice", params, nil); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	fmt.Println("Description updated")
}
```
