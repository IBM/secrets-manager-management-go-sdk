# IBM Cloud Secrets Manager Management Go SDK
A Go client library to interact with the IBM Cloud® Secrets Manager Instance Management APIs.

> **Important:** This SDK is for use with instances of the IBM Cloud Secrets Manager **Vault Enterprise plan only**. It is not compatible with other Secrets Manager plans.

<details>
<summary>Table of Contents</summary>

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Authentication](#authentication)
* [Using the SDK](#using-the-sdk)
* [Issues](#issues)
* [Contributing](#contributing)
* [License](#license)
</details>

## Overview
The IBM Cloud Secrets Manager Management Go SDK allows developers to programmatically interact with the following IBM Cloud services:

| Service name                              | Package name                              |
|-------------------------------------------|-------------------------------------------|
| Secrets Manager Management | SecretsManagerInstanceManagementV2 |

## Prerequisites

- An [IBM Cloud account](https://cloud.ibm.com/registration).
- A [Secrets Manager service instance](https://cloud.ibm.com/catalog/services/secrets-manager) with the **Vault Enterprise plan**.
- An [IBM Cloud API key](https://cloud.ibm.com/iam/apikeys) that allows the SDK to access your account. 
- Go version 1.19 or above. 

  This SDK is tested with Go versions 1.19 and up. The SDK may work on previous versions, but this is not officially supported.

## Installation
There are a few different ways to download and install the Secrets Manager Management Go SDK project for use by your Go application.

#### `go get` command  
Use this command to download and install the SDK:

```
go get -u github.com/IBM/secrets-manager-management-go-sdk/v2
```

#### Go modules  
If your application uses Go modules, you can add the following import to your Go application:

```go
import (
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)
```

Then run `go mod tidy` to download and install the new dependency and update the `go.mod` file for your application.

## Authentication

Secrets Manager uses token-based Identity and Access Management (IAM) authentication.  

With IAM authentication, you supply an API key that is used to generate an access token. Then, the access token is included in each API request to Secrets Manager. Access tokens are valid for a limited amount of time and must be regenerated.  
Authentication for this SDK is accomplished by using [IAM authenticators](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md#authentication). Import authenticators from `github.com/IBM/go-sdk-core/v5/core`.

### Examples
#### Programmatic credentials

```go
import "github.com/IBM/go-sdk-core/v5/core"

authenticator := &core.IamAuthenticator{
  ApiKey: "{apikey}",
}
```
To learn more about IAM authenticators and how to use them in your Go application, see the [IBM Go SDK Core documentation](https://github.com/IBM/go-sdk-core/blob/master/Authentication.md).

## Using the SDK

### Basic usage

- All methods return a response and an error. The response contains the body, the headers, the status code, and the status text.
- Use the `URL` parameter to set the endpoint URL that is specific to your Secrets Manager service instance. To find your endpoint URL, you can copy it from the **Endpoints** section on the **Overview** page in the Secrets Manager UI.

#### Examples

Construct a service client and use it to generate an admin token and get instance details.

Here's an example `main.go` file:

```go
package main

import (
    "fmt"
    "github.com/IBM/go-sdk-core/v5/core"
    sm "github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func main() {

    secretsManagerInstanceManagement, err := sm.NewSecretsManagerInstanceManagementV2(&sm.SecretsManagerInstanceManagementV2Options{
        URL: "<SERVICE_URL>",
        Authenticator: &core.IamAuthenticator{
            ApiKey: "<IBM_CLOUD_API_KEY>",
        },
    })

    if err != nil {
        panic(err)
    }

    // Generate admin token
    generateTokenOptions := secretsManagerInstanceManagement.NewAdminTokenGenerateOptions(
        "<INSTANCE_CRN>",
    )

    token, response, err := secretsManagerInstanceManagement.AdminTokenGenerate(generateTokenOptions)
    if err != nil {
        panic(err)
    }
    fmt.Println("Admin token generated!")

    // Get instance details
    getInstanceOptions := secretsManagerInstanceManagement.NewInstanceDetailsOptions(
        "<INSTANCE_CRN>",
    )

    instance, response, err := secretsManagerInstanceManagement.InstanceDetails(getInstanceOptions)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Instance details: %+v\n", instance)
}
```

Replace the `URL` and `ApiKey` values. Then run the `go run main.go` command to compile and run your Go program.

For more information and IBM Cloud SDK usage examples for Go, see the [IBM Cloud SDK Common documentation](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md).  

## Issues

If you encounter an issue with the project, you're welcome to submit a [bug report](https://github.com/IBM/secrets-manager-management-go-sdk/issues) to help us improve.

## Contributing
For general contribution guidelines, see [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](LICENSE).


