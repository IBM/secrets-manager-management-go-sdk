// Package common contains shared utilities for the SDK
package common

import (
	"fmt"
	"runtime"

	"github.com/IBM/go-sdk-core/v5/core"
)

const (
	sdkName    = "secrets-manager-vaas-go-sdk"
	sdkVersion = "0.0.1"
)

// GetComponentInfo returns the name and version of this module
func GetComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(sdkName, sdkVersion)
}

// GetSdkHeaders returns custom headers to be sent with requests
func GetSdkHeaders(serviceName, serviceVersion, operationID string) map[string]string {
	headers := make(map[string]string)
	headers["User-Agent"] = GetUserAgentInfo()
	headers["X-SDK-Analytics"] = fmt.Sprintf("service_name=%s;service_version=%s;operation_id=%s", serviceName, serviceVersion, operationID)
	return headers
}

// GetUserAgentInfo returns the User-Agent header value
func GetUserAgentInfo() string {
	return fmt.Sprintf("%s/%s (%s; %s; %s)", sdkName, sdkVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
