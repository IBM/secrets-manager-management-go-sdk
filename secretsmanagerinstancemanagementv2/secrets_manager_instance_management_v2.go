/**
 * (C) Copyright IBM Corp. 2026.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

// Package secretsmanagerinstancemanagementv2 : Operations and models for the SecretsManagerInstanceManagementV2 service
package secretsmanagerinstancemanagementv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/secrets-manager-management-go-sdk/v2/common"
	"github.com/go-openapi/strfmt"
)

// SecretsManagerInstanceManagementV2 : Use the IBM  Cloud® Secrets Manager Instance Management API to manage service
// instances of the Vault Dedicated plan.
// - Get service instance details including cluster state, endpoints, and key management service.
// - Generate a Vault admin token for authenticating to your Vault Dedicated cluster.
// - Delete all active Vault admin tokens.
// - Manage destinations for your Vault Dedicated cluster.
// - Manage HSM destinations and their related resources, including HA groups, paths, server CAs, and client
// certificates.
// - Request payloads must not exceed 1 MB; requests larger than this limit will be rejected with a `413 Payload Too
// Large` response.
//
// API Version: 2.0.0
// See: https://cloud.ibm.com/docs/secrets-manager
type SecretsManagerInstanceManagementV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.secrets-manager.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "secrets_manager_instance_management"

const ParameterizedServiceURL = "https://{region}.secrets-manager.cloud.ibm.com"

var defaultUrlVariables = map[string]string{
	"region": "us-south",
}

// SecretsManagerInstanceManagementV2Options : Service options
type SecretsManagerInstanceManagementV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewSecretsManagerInstanceManagementV2UsingExternalConfig : constructs an instance of SecretsManagerInstanceManagementV2 with passed in options and external configuration.
func NewSecretsManagerInstanceManagementV2UsingExternalConfig(options *SecretsManagerInstanceManagementV2Options) (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	secretsManagerInstanceManagement, err = NewSecretsManagerInstanceManagementV2(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = secretsManagerInstanceManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = secretsManagerInstanceManagement.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewSecretsManagerInstanceManagementV2 : constructs an instance of SecretsManagerInstanceManagementV2 with passed in options.
func NewSecretsManagerInstanceManagementV2(options *SecretsManagerInstanceManagementV2Options) (service *SecretsManagerInstanceManagementV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &SecretsManagerInstanceManagementV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "secretsManagerInstanceManagement" suitable for processing requests.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) Clone() *SecretsManagerInstanceManagementV2 {
	if core.IsNil(secretsManagerInstanceManagement) {
		return nil
	}
	clone := *secretsManagerInstanceManagement
	clone.Service = secretsManagerInstanceManagement.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) SetServiceURL(url string) error {
	err := secretsManagerInstanceManagement.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetServiceURL() string {
	return secretsManagerInstanceManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) SetDefaultHeaders(headers http.Header) {
	secretsManagerInstanceManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) SetEnableGzipCompression(enableGzip bool) {
	secretsManagerInstanceManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetEnableGzipCompression() bool {
	return secretsManagerInstanceManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	secretsManagerInstanceManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DisableRetries() {
	secretsManagerInstanceManagement.Service.DisableRetries()
}

// CreateVaultAdmintoken : Create admin token
// Generate a Vault admin token for authenticating to your Vault Dedicated cluster. The token is valid for 1 hour and
// grants administrative privileges. Use only for initial setup and cluster management, then delete immediately.
// Optionally, set response_wrapping to true in the request body to receive a Vault response-wrapped token instead of a
// plain admin token.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateVaultAdmintoken(createVaultAdmintokenOptions *CreateVaultAdmintokenOptions) (result TokenIntf, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateVaultAdmintokenWithContext(context.Background(), createVaultAdmintokenOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateVaultAdmintokenWithContext is an alternate form of the CreateVaultAdmintoken method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateVaultAdmintokenWithContext(ctx context.Context, createVaultAdmintokenOptions *CreateVaultAdmintokenOptions) (result TokenIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVaultAdmintokenOptions, "createVaultAdmintokenOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createVaultAdmintokenOptions, "createVaultAdmintokenOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createVaultAdmintokenOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/admintokens`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreateVaultAdmintoken")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createVaultAdmintokenOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createVaultAdmintokenOptions.ResponseWrapping != nil {
		body["response_wrapping"] = createVaultAdmintokenOptions.ResponseWrapping
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_vault_admintoken", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalToken)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteVaultAdmintokens : Delete admin tokens
// Delete all active Vault admin tokens. This immediately invalidates all existing admin tokens including wrapped
// tokens.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteVaultAdmintokens(deleteVaultAdmintokensOptions *DeleteVaultAdmintokensOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.DeleteVaultAdmintokensWithContext(context.Background(), deleteVaultAdmintokensOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteVaultAdmintokensWithContext is an alternate form of the DeleteVaultAdmintokens method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteVaultAdmintokensWithContext(ctx context.Context, deleteVaultAdmintokensOptions *DeleteVaultAdmintokensOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVaultAdmintokensOptions, "deleteVaultAdmintokensOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteVaultAdmintokensOptions, "deleteVaultAdmintokensOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteVaultAdmintokensOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/admintokens`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteVaultAdmintokens")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteVaultAdmintokensOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_vault_admintokens", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetInstance : Get instance details
// Get service instance details including cluster state, endpoints, and key management service.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstance(getInstanceOptions *GetInstanceOptions) (result *Instance, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetInstanceWithContext(context.Background(), getInstanceOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceWithContext is an alternate form of the GetInstance method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceWithContext(ctx context.Context, getInstanceOptions *GetInstanceOptions) (result *Instance, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceOptions, "getInstanceOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getInstanceOptions, "getInstanceOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstance)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListInstanceDestinations : List destinations
// List all destinations for your Vault Dedicated cluster.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceDestinations(listInstanceDestinationsOptions *ListInstanceDestinationsOptions) (result *DestinationCollection, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ListInstanceDestinationsWithContext(context.Background(), listInstanceDestinationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListInstanceDestinationsWithContext is an alternate form of the ListInstanceDestinations method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceDestinationsWithContext(ctx context.Context, listInstanceDestinationsOptions *ListInstanceDestinationsOptions) (result *DestinationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listInstanceDestinationsOptions, "listInstanceDestinationsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listInstanceDestinationsOptions, "listInstanceDestinationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listInstanceDestinationsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ListInstanceDestinations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listInstanceDestinationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listInstanceDestinationsOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*listInstanceDestinationsOptions.State))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_instance_destinations", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDestinationCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateInstanceDestination : Create destination
// Create a new destination between your Vault Dedicated cluster and an IBM Cloud service instance.
//
// Returns `202 Accepted` with `state: provisioning`. Provisioning completes asynchronously — poll `GET
// /destinations/{id}` until `state` transitions to `succeeded` or `failed`.
//
// **Beta**: Only Gen 1 (Classic) IBM Cloud Database service instances are supported. Gen 2 instances are rejected with
// `422`. IBM Cloud Database service instances with no private endpoints are also rejected with `422`.
//
// **Rate Limit**: 10 requests per instance per minute
// **Quota**: Maximum 20 destinations per instance.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceDestination(createInstanceDestinationOptions *CreateInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateInstanceDestinationWithContext(context.Background(), createInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateInstanceDestinationWithContext is an alternate form of the CreateInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceDestinationWithContext(ctx context.Context, createInstanceDestinationOptions *CreateInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createInstanceDestinationOptions, "createInstanceDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createInstanceDestinationOptions, "createInstanceDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createInstanceDestinationOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreateInstanceDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createInstanceDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(createInstanceDestinationOptions.DestinationPrototype)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIbmCloudDatabaseDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceDestination : Get destination details
// Retrieve details and current state for a specific destination for your Vault Dedicated cluster.
//
// Returns `404` if the destination does not exist. A destination undergoing deletion may appear with a `deleting` state
// before it is fully removed.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceDestination(getInstanceDestinationOptions *GetInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetInstanceDestinationWithContext(context.Background(), getInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceDestinationWithContext is an alternate form of the GetInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceDestinationWithContext(ctx context.Context, getInstanceDestinationOptions *GetInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceDestinationOptions, "getInstanceDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getInstanceDestinationOptions, "getInstanceDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceDestinationOptions.InstanceID,
		"destination_id": *getInstanceDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetInstanceDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInstanceDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIbmCloudDatabaseDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateInstanceDestination : Update destination
// Update mutable metadata fields (`name`, `description`) on a destination for your Vault Dedicated cluster. All other
// fields are immutable after creation.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceDestination(updateInstanceDestinationOptions *UpdateInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.UpdateInstanceDestinationWithContext(context.Background(), updateInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateInstanceDestinationWithContext is an alternate form of the UpdateInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceDestinationWithContext(ctx context.Context, updateInstanceDestinationOptions *UpdateInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceDestinationOptions, "updateInstanceDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateInstanceDestinationOptions, "updateInstanceDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceDestinationOptions.InstanceID,
		"destination_id": *updateInstanceDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "UpdateInstanceDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateInstanceDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceDestinationOptions.DestinationPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIbmCloudDatabaseDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteInstanceDestination : Delete destination
// Delete a destination for your Vault Dedicated cluster. Returns `202 Accepted` and begins deletion asynchronously —
// the destination is removed from `GET /destinations` once deletion completes.
//
// A `failed` destination still counts against the per-instance quota until deleted.
//
// **Rate Limit**: 10 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceDestination(deleteInstanceDestinationOptions *DeleteInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.DeleteInstanceDestinationWithContext(context.Background(), deleteInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceDestinationWithContext is an alternate form of the DeleteInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceDestinationWithContext(ctx context.Context, deleteInstanceDestinationOptions *DeleteInstanceDestinationOptions) (result *IbmCloudDatabaseDestination, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteInstanceDestinationOptions, "deleteInstanceDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteInstanceDestinationOptions, "deleteInstanceDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteInstanceDestinationOptions.InstanceID,
		"destination_id": *deleteInstanceDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteInstanceDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteInstanceDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIbmCloudDatabaseDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListInstanceHsmDestinations : List destinations
// List all HSM Destinations for this SM instance. Typically returns a single entry since only one destination per
// vendor is allowed.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinations(listInstanceHsmDestinationsOptions *ListInstanceHsmDestinationsOptions) (result *HSMDestinationCollection, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ListInstanceHsmDestinationsWithContext(context.Background(), listInstanceHsmDestinationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListInstanceHsmDestinationsWithContext is an alternate form of the ListInstanceHsmDestinations method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationsWithContext(ctx context.Context, listInstanceHsmDestinationsOptions *ListInstanceHsmDestinationsOptions) (result *HSMDestinationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listInstanceHsmDestinationsOptions, "listInstanceHsmDestinationsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listInstanceHsmDestinationsOptions, "listInstanceHsmDestinationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listInstanceHsmDestinationsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ListInstanceHsmDestinations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listInstanceHsmDestinationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_instance_hsm_destinations", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMDestinationCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateInstanceHsmDestination : Create destination
// Create the HSM Destination for this SM instance. Each SM instance may have exactly **one** HSM Destination per
// vendor. Calling POST again returns `409 Conflict`.
//
// The response `account_id` is required for account policy setup on your gateway. Without this step, path provisioning
// cannot proceed.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestination(createInstanceHsmDestinationOptions *CreateInstanceHsmDestinationOptions) (result *HSMDestination, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateInstanceHsmDestinationWithContext(context.Background(), createInstanceHsmDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateInstanceHsmDestinationWithContext is an alternate form of the CreateInstanceHsmDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationWithContext(ctx context.Context, createInstanceHsmDestinationOptions *CreateInstanceHsmDestinationOptions) (result *HSMDestination, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createInstanceHsmDestinationOptions, "createInstanceHsmDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createInstanceHsmDestinationOptions, "createInstanceHsmDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createInstanceHsmDestinationOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreateInstanceHsmDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createInstanceHsmDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createInstanceHsmDestinationOptions.Name != nil {
		body["name"] = createInstanceHsmDestinationOptions.Name
	}
	if createInstanceHsmDestinationOptions.Vendor != nil {
		body["vendor"] = createInstanceHsmDestinationOptions.Vendor
	}
	if createInstanceHsmDestinationOptions.Description != nil {
		body["description"] = createInstanceHsmDestinationOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_instance_hsm_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceHsmDestination : Get destination
// Retrieve a single HSM Destination by its ID.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestination(getInstanceHsmDestinationOptions *GetInstanceHsmDestinationOptions) (result *HSMDestination, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetInstanceHsmDestinationWithContext(context.Background(), getInstanceHsmDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceHsmDestinationWithContext is an alternate form of the GetInstanceHsmDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationWithContext(ctx context.Context, getInstanceHsmDestinationOptions *GetInstanceHsmDestinationOptions) (result *HSMDestination, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceHsmDestinationOptions, "getInstanceHsmDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getInstanceHsmDestinationOptions, "getInstanceHsmDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceHsmDestinationOptions.InstanceID,
		"destination_id": *getInstanceHsmDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetInstanceHsmDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInstanceHsmDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance_hsm_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateInstanceHsmDestination : Update destination
// Update `name` or `description`. Does not stage any HSM config file changes. If `name` is omitted, the existing name
// is preserved (merge-patch semantics).
// `health_state` and `health_as_of` are read-only — if supplied they are silently ignored.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptions *UpdateInstanceHsmDestinationOptions) (result *HSMDestination, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.UpdateInstanceHsmDestinationWithContext(context.Background(), updateInstanceHsmDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateInstanceHsmDestinationWithContext is an alternate form of the UpdateInstanceHsmDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationWithContext(ctx context.Context, updateInstanceHsmDestinationOptions *UpdateInstanceHsmDestinationOptions) (result *HSMDestination, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceHsmDestinationOptions, "updateInstanceHsmDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateInstanceHsmDestinationOptions, "updateInstanceHsmDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceHsmDestinationOptions.InstanceID,
		"destination_id": *updateInstanceHsmDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "UpdateInstanceHsmDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateInstanceHsmDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceHsmDestinationOptions.HsmDestinationPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_instance_hsm_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteInstanceHsmDestination : Delete a destination asynchronously
// Initiate an asynchronous hard delete of the destination and all its sub-resources (HSM HA Groups, paths, Server CAs,
// client certificates, private keys). Sets `provisioning_state` to `deleting`. Poll `GET …/{destination_id}` until
// `404` to confirm completion. **This operation cannot be undone.**
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptions *DeleteInstanceHsmDestinationOptions) (result *AsyncDeleteResponse, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.DeleteInstanceHsmDestinationWithContext(context.Background(), deleteInstanceHsmDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceHsmDestinationWithContext is an alternate form of the DeleteInstanceHsmDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationWithContext(ctx context.Context, deleteInstanceHsmDestinationOptions *DeleteInstanceHsmDestinationOptions) (result *AsyncDeleteResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteInstanceHsmDestinationOptions, "deleteInstanceHsmDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteInstanceHsmDestinationOptions, "deleteInstanceHsmDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteInstanceHsmDestinationOptions.InstanceID,
		"destination_id": *deleteInstanceHsmDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteInstanceHsmDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteInstanceHsmDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_hsm_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAsyncDeleteResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetHsmDestinationHealthSnapshot : Get destination health snapshot
// Return the persisted AS-OF health snapshot for this destination. Does not trigger a new evaluation. Poll until
// `health_state` is no longer `evaluating` or `in_progress`.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptions *GetHsmDestinationHealthSnapshotOptions) (result *DestinationHealthSnapshot, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetHsmDestinationHealthSnapshotWithContext(context.Background(), getHsmDestinationHealthSnapshotOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetHsmDestinationHealthSnapshotWithContext is an alternate form of the GetHsmDestinationHealthSnapshot method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetHsmDestinationHealthSnapshotWithContext(ctx context.Context, getHsmDestinationHealthSnapshotOptions *GetHsmDestinationHealthSnapshotOptions) (result *DestinationHealthSnapshot, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getHsmDestinationHealthSnapshotOptions, "getHsmDestinationHealthSnapshotOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getHsmDestinationHealthSnapshotOptions, "getHsmDestinationHealthSnapshotOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getHsmDestinationHealthSnapshotOptions.InstanceID,
		"destination_id": *getHsmDestinationHealthSnapshotOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/health`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetHsmDestinationHealthSnapshot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getHsmDestinationHealthSnapshotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_hsm_destination_health_snapshot", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDestinationHealthSnapshot)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RunHsmDestinationHealthCheck : Trigger an async destination health check
// Start an asynchronous health check across all paths, HSM HA Groups, and client certificates. Sets `health_state` to
// `evaluating` on the destination and all sub-resources immediately. Returns `202 Accepted`. Poll `GET …/health` until
// `health_state` is no longer `evaluating`.
//
// Returns `409 health_check_in_progress` if `health_state` is already `evaluating` or `in_progress`.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptions *RunHsmDestinationHealthCheckOptions) (result *AsyncHealthCheckAckResponse, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.RunHsmDestinationHealthCheckWithContext(context.Background(), runHsmDestinationHealthCheckOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RunHsmDestinationHealthCheckWithContext is an alternate form of the RunHsmDestinationHealthCheck method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) RunHsmDestinationHealthCheckWithContext(ctx context.Context, runHsmDestinationHealthCheckOptions *RunHsmDestinationHealthCheckOptions) (result *AsyncHealthCheckAckResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runHsmDestinationHealthCheckOptions, "runHsmDestinationHealthCheckOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(runHsmDestinationHealthCheckOptions, "runHsmDestinationHealthCheckOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *runHsmDestinationHealthCheckOptions.InstanceID,
		"destination_id": *runHsmDestinationHealthCheckOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/health`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "RunHsmDestinationHealthCheck")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range runHsmDestinationHealthCheckOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "run_hsm_destination_health_check", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAsyncHealthCheckAckResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ApplyHsmDestination : Trigger an async HSM configuration apply
// Applies the current configuration (paths, HSM HA Groups, Server CAs, and client certificates) of the destination.
// Sets `health_state` to `in_progress` on the destination and all sub-resources immediately, then continues
// asynchronously. Returns `202 Accepted`. On success, automatically triggers a health check. On failure, sets all
// `health_state` to `unknown`.
//
// Returns `409 apply_not_allowed` if `health_state` is already `in_progress` or `evaluating`. Poll `GET …/health` until
// `health_state` is no longer `in_progress`.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ApplyHsmDestination(applyHsmDestinationOptions *ApplyHsmDestinationOptions) (result *AsyncHealthAckResponse, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ApplyHsmDestinationWithContext(context.Background(), applyHsmDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ApplyHsmDestinationWithContext is an alternate form of the ApplyHsmDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ApplyHsmDestinationWithContext(ctx context.Context, applyHsmDestinationOptions *ApplyHsmDestinationOptions) (result *AsyncHealthAckResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(applyHsmDestinationOptions, "applyHsmDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(applyHsmDestinationOptions, "applyHsmDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *applyHsmDestinationOptions.InstanceID,
		"destination_id": *applyHsmDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/apply`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ApplyHsmDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range applyHsmDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "apply_hsm_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAsyncHealthAckResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListInstanceHsmDestinationHaGroups : List HSM HA Groups
// List all HSM HA Groups for this destination. Returns an empty array if none have been created. Results are ordered by
// creation time ascending.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptions *ListInstanceHsmDestinationHaGroupsOptions) (result *HSMHAGroupCollection, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ListInstanceHsmDestinationHaGroupsWithContext(context.Background(), listInstanceHsmDestinationHaGroupsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListInstanceHsmDestinationHaGroupsWithContext is an alternate form of the ListInstanceHsmDestinationHaGroups method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationHaGroupsWithContext(ctx context.Context, listInstanceHsmDestinationHaGroupsOptions *ListInstanceHsmDestinationHaGroupsOptions) (result *HSMHAGroupCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listInstanceHsmDestinationHaGroupsOptions, "listInstanceHsmDestinationHaGroupsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listInstanceHsmDestinationHaGroupsOptions, "listInstanceHsmDestinationHaGroupsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listInstanceHsmDestinationHaGroupsOptions.InstanceID,
		"destination_id": *listInstanceHsmDestinationHaGroupsOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/ha_groups`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ListInstanceHsmDestinationHaGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listInstanceHsmDestinationHaGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_instance_hsm_destination_ha_groups", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMHAGroupCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateInstanceHsmDestinationHaGroup : Create HA Group
// Create a new HA Group. `group_label` must be unique within the destination — returns `409 group_label_already_exists`
// if duplicate. Max 20 HSM HA Groups per destination — returns `409 ha_group_limit_reached` if exceeded.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptions *CreateInstanceHsmDestinationHaGroupOptions) (result *HSMHAGroup, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateInstanceHsmDestinationHaGroupWithContext(context.Background(), createInstanceHsmDestinationHaGroupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateInstanceHsmDestinationHaGroupWithContext is an alternate form of the CreateInstanceHsmDestinationHaGroup method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationHaGroupWithContext(ctx context.Context, createInstanceHsmDestinationHaGroupOptions *CreateInstanceHsmDestinationHaGroupOptions) (result *HSMHAGroup, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createInstanceHsmDestinationHaGroupOptions, "createInstanceHsmDestinationHaGroupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createInstanceHsmDestinationHaGroupOptions, "createInstanceHsmDestinationHaGroupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createInstanceHsmDestinationHaGroupOptions.InstanceID,
		"destination_id": *createInstanceHsmDestinationHaGroupOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/ha_groups`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreateInstanceHsmDestinationHaGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createInstanceHsmDestinationHaGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createInstanceHsmDestinationHaGroupOptions.Name != nil {
		body["name"] = createInstanceHsmDestinationHaGroupOptions.Name
	}
	if createInstanceHsmDestinationHaGroupOptions.GroupLabel != nil {
		body["group_label"] = createInstanceHsmDestinationHaGroupOptions.GroupLabel
	}
	if createInstanceHsmDestinationHaGroupOptions.GroupSerial != nil {
		body["group_serial"] = createInstanceHsmDestinationHaGroupOptions.GroupSerial
	}
	if createInstanceHsmDestinationHaGroupOptions.MemberSerials != nil {
		body["member_serials"] = createInstanceHsmDestinationHaGroupOptions.MemberSerials
	}
	if createInstanceHsmDestinationHaGroupOptions.Description != nil {
		body["description"] = createInstanceHsmDestinationHaGroupOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_instance_hsm_destination_ha_group", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMHAGroup)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceHsmDestinationHaGroup : Get HA Group
// Retrieve a single HA Group by its ID.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptions *GetInstanceHsmDestinationHaGroupOptions) (result *HSMHAGroup, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetInstanceHsmDestinationHaGroupWithContext(context.Background(), getInstanceHsmDestinationHaGroupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceHsmDestinationHaGroupWithContext is an alternate form of the GetInstanceHsmDestinationHaGroup method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationHaGroupWithContext(ctx context.Context, getInstanceHsmDestinationHaGroupOptions *GetInstanceHsmDestinationHaGroupOptions) (result *HSMHAGroup, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceHsmDestinationHaGroupOptions, "getInstanceHsmDestinationHaGroupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getInstanceHsmDestinationHaGroupOptions, "getInstanceHsmDestinationHaGroupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceHsmDestinationHaGroupOptions.InstanceID,
		"destination_id": *getInstanceHsmDestinationHaGroupOptions.DestinationID,
		"ha_group_id": *getInstanceHsmDestinationHaGroupOptions.HaGroupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/ha_groups/{ha_group_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetInstanceHsmDestinationHaGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInstanceHsmDestinationHaGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance_hsm_destination_ha_group", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMHAGroup)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateInstanceHsmDestinationHaGroup : Update HA Group
// Update an HA Group. `group_serial` and `group_label` are immutable after creation — return `400` if supplied. If
// `name` is omitted, the existing name is preserved (merge-patch semantics). `name` cannot be set to `null`.
// `member_serials` can be updated to add or remove members (at least one entry required).
// `health_state` and `health_as_of` are read-only — silently ignored if supplied.
//
// **HSM Config File**: staged when `member_serials` is changed.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptions *UpdateInstanceHsmDestinationHaGroupOptions) (result *HSMHAGroup, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.UpdateInstanceHsmDestinationHaGroupWithContext(context.Background(), updateInstanceHsmDestinationHaGroupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateInstanceHsmDestinationHaGroupWithContext is an alternate form of the UpdateInstanceHsmDestinationHaGroup method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationHaGroupWithContext(ctx context.Context, updateInstanceHsmDestinationHaGroupOptions *UpdateInstanceHsmDestinationHaGroupOptions) (result *HSMHAGroup, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceHsmDestinationHaGroupOptions, "updateInstanceHsmDestinationHaGroupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateInstanceHsmDestinationHaGroupOptions, "updateInstanceHsmDestinationHaGroupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceHsmDestinationHaGroupOptions.InstanceID,
		"destination_id": *updateInstanceHsmDestinationHaGroupOptions.DestinationID,
		"ha_group_id": *updateInstanceHsmDestinationHaGroupOptions.HaGroupID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/ha_groups/{ha_group_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "UpdateInstanceHsmDestinationHaGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateInstanceHsmDestinationHaGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceHsmDestinationHaGroupOptions.HSMHAGroupPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_instance_hsm_destination_ha_group", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMHAGroup)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteInstanceHsmDestinationHaGroup : Delete HA Group (synchronous hard delete)
// Synchronous hard delete — returns 204 immediately.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptions *DeleteInstanceHsmDestinationHaGroupOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.DeleteInstanceHsmDestinationHaGroupWithContext(context.Background(), deleteInstanceHsmDestinationHaGroupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceHsmDestinationHaGroupWithContext is an alternate form of the DeleteInstanceHsmDestinationHaGroup method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationHaGroupWithContext(ctx context.Context, deleteInstanceHsmDestinationHaGroupOptions *DeleteInstanceHsmDestinationHaGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteInstanceHsmDestinationHaGroupOptions, "deleteInstanceHsmDestinationHaGroupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteInstanceHsmDestinationHaGroupOptions, "deleteInstanceHsmDestinationHaGroupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteInstanceHsmDestinationHaGroupOptions.InstanceID,
		"destination_id": *deleteInstanceHsmDestinationHaGroupOptions.DestinationID,
		"ha_group_id": *deleteInstanceHsmDestinationHaGroupOptions.HaGroupID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/ha_groups/{ha_group_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteInstanceHsmDestinationHaGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteInstanceHsmDestinationHaGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_hsm_destination_ha_group", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListInstanceHsmDestinationPaths : List paths
// List all paths (nodes) for this destination. Returns an empty array if none have been created. Results are ordered by
// creation time ascending.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptions *ListInstanceHsmDestinationPathsOptions) (result *PathCollection, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ListInstanceHsmDestinationPathsWithContext(context.Background(), listInstanceHsmDestinationPathsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListInstanceHsmDestinationPathsWithContext is an alternate form of the ListInstanceHsmDestinationPaths method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationPathsWithContext(ctx context.Context, listInstanceHsmDestinationPathsOptions *ListInstanceHsmDestinationPathsOptions) (result *PathCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listInstanceHsmDestinationPathsOptions, "listInstanceHsmDestinationPathsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listInstanceHsmDestinationPathsOptions, "listInstanceHsmDestinationPathsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listInstanceHsmDestinationPathsOptions.InstanceID,
		"destination_id": *listInstanceHsmDestinationPathsOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ListInstanceHsmDestinationPaths")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listInstanceHsmDestinationPathsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_instance_hsm_destination_paths", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPathCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateInstanceHsmDestinationPath : Create path
// Create a new path. `crn` is immutable. Must be unique within the destination — returns `409 crn_already_in_use` if
// duplicate. Max 20 paths per destination — returns `409 path_limit_reached` if exceeded.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptions *CreateInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateInstanceHsmDestinationPathWithContext(context.Background(), createInstanceHsmDestinationPathOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateInstanceHsmDestinationPathWithContext is an alternate form of the CreateInstanceHsmDestinationPath method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationPathWithContext(ctx context.Context, createInstanceHsmDestinationPathOptions *CreateInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createInstanceHsmDestinationPathOptions, "createInstanceHsmDestinationPathOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createInstanceHsmDestinationPathOptions, "createInstanceHsmDestinationPathOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createInstanceHsmDestinationPathOptions.InstanceID,
		"destination_id": *createInstanceHsmDestinationPathOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreateInstanceHsmDestinationPath")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createInstanceHsmDestinationPathOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createInstanceHsmDestinationPathOptions.Name != nil {
		body["name"] = createInstanceHsmDestinationPathOptions.Name
	}
	if createInstanceHsmDestinationPathOptions.Crn != nil {
		body["crn"] = createInstanceHsmDestinationPathOptions.Crn
	}
	if createInstanceHsmDestinationPathOptions.Description != nil {
		body["description"] = createInstanceHsmDestinationPathOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_instance_hsm_destination_path", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPath)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceHsmDestinationPath : Get path
// Retrieve a single path by its ID.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions *GetInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetInstanceHsmDestinationPathWithContext(context.Background(), getInstanceHsmDestinationPathOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceHsmDestinationPathWithContext is an alternate form of the GetInstanceHsmDestinationPath method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationPathWithContext(ctx context.Context, getInstanceHsmDestinationPathOptions *GetInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceHsmDestinationPathOptions, "getInstanceHsmDestinationPathOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getInstanceHsmDestinationPathOptions, "getInstanceHsmDestinationPathOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceHsmDestinationPathOptions.InstanceID,
		"destination_id": *getInstanceHsmDestinationPathOptions.DestinationID,
		"path_id": *getInstanceHsmDestinationPathOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetInstanceHsmDestinationPath")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInstanceHsmDestinationPathOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance_hsm_destination_path", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPath)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateInstanceHsmDestinationPath : Update path
// Update `name` or `description` (merge-patch semantics). `crn` is immutable.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptions *UpdateInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.UpdateInstanceHsmDestinationPathWithContext(context.Background(), updateInstanceHsmDestinationPathOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateInstanceHsmDestinationPathWithContext is an alternate form of the UpdateInstanceHsmDestinationPath method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationPathWithContext(ctx context.Context, updateInstanceHsmDestinationPathOptions *UpdateInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceHsmDestinationPathOptions, "updateInstanceHsmDestinationPathOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateInstanceHsmDestinationPathOptions, "updateInstanceHsmDestinationPathOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceHsmDestinationPathOptions.InstanceID,
		"destination_id": *updateInstanceHsmDestinationPathOptions.DestinationID,
		"path_id": *updateInstanceHsmDestinationPathOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "UpdateInstanceHsmDestinationPath")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateInstanceHsmDestinationPathOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceHsmDestinationPathOptions.PathPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_instance_hsm_destination_path", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPath)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteInstanceHsmDestinationPath : Delete path (async hard delete)
// Async hard delete of the path and its Server CAs. Sets `provisioning_state` to `deleting`. Poll `GET …/{path_id}`
// until `404`.
//
// **This operation cannot be undone.**
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptions *DeleteInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.DeleteInstanceHsmDestinationPathWithContext(context.Background(), deleteInstanceHsmDestinationPathOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceHsmDestinationPathWithContext is an alternate form of the DeleteInstanceHsmDestinationPath method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationPathWithContext(ctx context.Context, deleteInstanceHsmDestinationPathOptions *DeleteInstanceHsmDestinationPathOptions) (result *Path, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteInstanceHsmDestinationPathOptions, "deleteInstanceHsmDestinationPathOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteInstanceHsmDestinationPathOptions, "deleteInstanceHsmDestinationPathOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteInstanceHsmDestinationPathOptions.InstanceID,
		"destination_id": *deleteInstanceHsmDestinationPathOptions.DestinationID,
		"path_id": *deleteInstanceHsmDestinationPathOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteInstanceHsmDestinationPath")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteInstanceHsmDestinationPathOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_hsm_destination_path", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPath)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreatePathBuild : Build path (trigger provisioning)
// Trigger provisioning of the path. The customer must first approve the Secrets Manager account policy for the gateway.
// Moves path from `pending_account_policy` → `provisioning`. Returns `409` if already `provisioning`.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreatePathBuild(createPathBuildOptions *CreatePathBuildOptions) (result *VpeProvisioningResponse, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreatePathBuildWithContext(context.Background(), createPathBuildOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreatePathBuildWithContext is an alternate form of the CreatePathBuild method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreatePathBuildWithContext(ctx context.Context, createPathBuildOptions *CreatePathBuildOptions) (result *VpeProvisioningResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPathBuildOptions, "createPathBuildOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createPathBuildOptions, "createPathBuildOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createPathBuildOptions.InstanceID,
		"destination_id": *createPathBuildOptions.DestinationID,
		"path_id": *createPathBuildOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/build`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreatePathBuild")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createPathBuildOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_path_build", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVpeProvisioningResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetPathVpeStatus : Health check path (check provisioning status)
// Check the current provisioning status of this path. Can be called while the path is in `provisioning` or
// `pending_approval` state. Returns the latest health state and the resulting path `provisioning_state`.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetPathVpeStatus(getPathVpeStatusOptions *GetPathVpeStatusOptions) (result *VpeStatusResult, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetPathVpeStatusWithContext(context.Background(), getPathVpeStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPathVpeStatusWithContext is an alternate form of the GetPathVpeStatus method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetPathVpeStatusWithContext(ctx context.Context, getPathVpeStatusOptions *GetPathVpeStatusOptions) (result *VpeStatusResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPathVpeStatusOptions, "getPathVpeStatusOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getPathVpeStatusOptions, "getPathVpeStatusOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getPathVpeStatusOptions.InstanceID,
		"destination_id": *getPathVpeStatusOptions.DestinationID,
		"path_id": *getPathVpeStatusOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/health_check`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetPathVpeStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getPathVpeStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_path_vpe_status", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVpeStatusResult)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ResetHsmDestinationPath : Reset a path to its initial pending state
// Reset path to `pending_account_policy` so `POST /build` can be retried. Allowed from `crn_already_in_use`,
// `approval_declined`, `completed`, or `vpe_provisioning_failed`. Returns `409` otherwise.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ResetHsmDestinationPath(resetHsmDestinationPathOptions *ResetHsmDestinationPathOptions) (result *HSMPathReset, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ResetHsmDestinationPathWithContext(context.Background(), resetHsmDestinationPathOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ResetHsmDestinationPathWithContext is an alternate form of the ResetHsmDestinationPath method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ResetHsmDestinationPathWithContext(ctx context.Context, resetHsmDestinationPathOptions *ResetHsmDestinationPathOptions) (result *HSMPathReset, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resetHsmDestinationPathOptions, "resetHsmDestinationPathOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(resetHsmDestinationPathOptions, "resetHsmDestinationPathOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *resetHsmDestinationPathOptions.InstanceID,
		"destination_id": *resetHsmDestinationPathOptions.DestinationID,
		"path_id": *resetHsmDestinationPathOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/reset`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ResetHsmDestinationPath")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range resetHsmDestinationPathOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "reset_hsm_destination_path", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHSMPathReset)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListInstanceHsmDestinationPathServerCas : List Server CA certificates
// List all Server CA certificates for this path. Returns an empty array if none have been added. At most one
// certificate can be active at a time; up to 10 certificates total per path. Results are ordered by creation time
// ascending.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptions *ListInstanceHsmDestinationPathServerCasOptions) (result *ServerCaCollection, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ListInstanceHsmDestinationPathServerCasWithContext(context.Background(), listInstanceHsmDestinationPathServerCasOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListInstanceHsmDestinationPathServerCasWithContext is an alternate form of the ListInstanceHsmDestinationPathServerCas method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationPathServerCasWithContext(ctx context.Context, listInstanceHsmDestinationPathServerCasOptions *ListInstanceHsmDestinationPathServerCasOptions) (result *ServerCaCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listInstanceHsmDestinationPathServerCasOptions, "listInstanceHsmDestinationPathServerCasOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listInstanceHsmDestinationPathServerCasOptions, "listInstanceHsmDestinationPathServerCasOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listInstanceHsmDestinationPathServerCasOptions.InstanceID,
		"destination_id": *listInstanceHsmDestinationPathServerCasOptions.DestinationID,
		"path_id": *listInstanceHsmDestinationPathServerCasOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/server_cas`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ListInstanceHsmDestinationPathServerCas")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listInstanceHsmDestinationPathServerCasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_instance_hsm_destination_path_server_cas", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServerCaCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateInstanceHsmDestinationPathServerCa : Add Server CA certificate
// Add a Server CA (PEM). Starts `inactive`. Max 10 per path — returns `409 server_ca_limit_reached` if exceeded.
// Certificate must not be expired — returns `400 cert_expired`.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptions *CreateInstanceHsmDestinationPathServerCaOptions) (result *ServerCa, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateInstanceHsmDestinationPathServerCaWithContext(context.Background(), createInstanceHsmDestinationPathServerCaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateInstanceHsmDestinationPathServerCaWithContext is an alternate form of the CreateInstanceHsmDestinationPathServerCa method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationPathServerCaWithContext(ctx context.Context, createInstanceHsmDestinationPathServerCaOptions *CreateInstanceHsmDestinationPathServerCaOptions) (result *ServerCa, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createInstanceHsmDestinationPathServerCaOptions, "createInstanceHsmDestinationPathServerCaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createInstanceHsmDestinationPathServerCaOptions, "createInstanceHsmDestinationPathServerCaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createInstanceHsmDestinationPathServerCaOptions.InstanceID,
		"destination_id": *createInstanceHsmDestinationPathServerCaOptions.DestinationID,
		"path_id": *createInstanceHsmDestinationPathServerCaOptions.PathID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/server_cas`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreateInstanceHsmDestinationPathServerCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createInstanceHsmDestinationPathServerCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createInstanceHsmDestinationPathServerCaOptions.Name != nil {
		body["name"] = createInstanceHsmDestinationPathServerCaOptions.Name
	}
	if createInstanceHsmDestinationPathServerCaOptions.Certificate != nil {
		body["certificate"] = createInstanceHsmDestinationPathServerCaOptions.Certificate
	}
	if createInstanceHsmDestinationPathServerCaOptions.Description != nil {
		body["description"] = createInstanceHsmDestinationPathServerCaOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_instance_hsm_destination_path_server_ca", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServerCa)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceHsmDestinationPathServerCa : Get Server CA certificate
// Retrieve a single Server CA certificate by its ID.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptions *GetInstanceHsmDestinationPathServerCaOptions) (result *ServerCa, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetInstanceHsmDestinationPathServerCaWithContext(context.Background(), getInstanceHsmDestinationPathServerCaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceHsmDestinationPathServerCaWithContext is an alternate form of the GetInstanceHsmDestinationPathServerCa method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationPathServerCaWithContext(ctx context.Context, getInstanceHsmDestinationPathServerCaOptions *GetInstanceHsmDestinationPathServerCaOptions) (result *ServerCa, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceHsmDestinationPathServerCaOptions, "getInstanceHsmDestinationPathServerCaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getInstanceHsmDestinationPathServerCaOptions, "getInstanceHsmDestinationPathServerCaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceHsmDestinationPathServerCaOptions.InstanceID,
		"destination_id": *getInstanceHsmDestinationPathServerCaOptions.DestinationID,
		"path_id": *getInstanceHsmDestinationPathServerCaOptions.PathID,
		"server_ca_id": *getInstanceHsmDestinationPathServerCaOptions.ServerCaID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/server_cas/{server_ca_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetInstanceHsmDestinationPathServerCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInstanceHsmDestinationPathServerCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance_hsm_destination_path_server_ca", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServerCa)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateInstanceHsmDestinationPathServerCa : Update or activate Server CA
// Rename, update description, or activate a Server CA (merge-patch semantics). `certificate` is immutable. Activating
// demotes the currently active CA to `inactive`.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptions *UpdateInstanceHsmDestinationPathServerCaOptions) (result *ServerCa, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.UpdateInstanceHsmDestinationPathServerCaWithContext(context.Background(), updateInstanceHsmDestinationPathServerCaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateInstanceHsmDestinationPathServerCaWithContext is an alternate form of the UpdateInstanceHsmDestinationPathServerCa method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationPathServerCaWithContext(ctx context.Context, updateInstanceHsmDestinationPathServerCaOptions *UpdateInstanceHsmDestinationPathServerCaOptions) (result *ServerCa, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceHsmDestinationPathServerCaOptions, "updateInstanceHsmDestinationPathServerCaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateInstanceHsmDestinationPathServerCaOptions, "updateInstanceHsmDestinationPathServerCaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceHsmDestinationPathServerCaOptions.InstanceID,
		"destination_id": *updateInstanceHsmDestinationPathServerCaOptions.DestinationID,
		"path_id": *updateInstanceHsmDestinationPathServerCaOptions.PathID,
		"server_ca_id": *updateInstanceHsmDestinationPathServerCaOptions.ServerCaID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/server_cas/{server_ca_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "UpdateInstanceHsmDestinationPathServerCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateInstanceHsmDestinationPathServerCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceHsmDestinationPathServerCaOptions.ServerCaPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_instance_hsm_destination_path_server_ca", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServerCa)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteInstanceHsmDestinationPathServerCa : Delete Server CA certificate
// Hard delete a Server CA certificate. Cannot delete the active Server CA — returns `409`. Synchronous (204, no body).
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptions *DeleteInstanceHsmDestinationPathServerCaOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.DeleteInstanceHsmDestinationPathServerCaWithContext(context.Background(), deleteInstanceHsmDestinationPathServerCaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceHsmDestinationPathServerCaWithContext is an alternate form of the DeleteInstanceHsmDestinationPathServerCa method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationPathServerCaWithContext(ctx context.Context, deleteInstanceHsmDestinationPathServerCaOptions *DeleteInstanceHsmDestinationPathServerCaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteInstanceHsmDestinationPathServerCaOptions, "deleteInstanceHsmDestinationPathServerCaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteInstanceHsmDestinationPathServerCaOptions, "deleteInstanceHsmDestinationPathServerCaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteInstanceHsmDestinationPathServerCaOptions.InstanceID,
		"destination_id": *deleteInstanceHsmDestinationPathServerCaOptions.DestinationID,
		"path_id": *deleteInstanceHsmDestinationPathServerCaOptions.PathID,
		"server_ca_id": *deleteInstanceHsmDestinationPathServerCaOptions.ServerCaID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/paths/{path_id}/server_cas/{server_ca_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteInstanceHsmDestinationPathServerCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteInstanceHsmDestinationPathServerCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_hsm_destination_path_server_ca", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListInstanceHsmDestinationClientCertificates : List client certificates
// List all client certificates for this destination. Returns an empty array if none have been created. At most one
// certificate can be active; up to 10 inactive or pending certificates may exist in addition. Results are ordered by
// creation time ascending.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptions *ListInstanceHsmDestinationClientCertificatesOptions) (result *ClientCertificateCollection, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.ListInstanceHsmDestinationClientCertificatesWithContext(context.Background(), listInstanceHsmDestinationClientCertificatesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListInstanceHsmDestinationClientCertificatesWithContext is an alternate form of the ListInstanceHsmDestinationClientCertificates method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) ListInstanceHsmDestinationClientCertificatesWithContext(ctx context.Context, listInstanceHsmDestinationClientCertificatesOptions *ListInstanceHsmDestinationClientCertificatesOptions) (result *ClientCertificateCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listInstanceHsmDestinationClientCertificatesOptions, "listInstanceHsmDestinationClientCertificatesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listInstanceHsmDestinationClientCertificatesOptions, "listInstanceHsmDestinationClientCertificatesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listInstanceHsmDestinationClientCertificatesOptions.InstanceID,
		"destination_id": *listInstanceHsmDestinationClientCertificatesOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/client_certificates`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "ListInstanceHsmDestinationClientCertificates")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listInstanceHsmDestinationClientCertificatesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_instance_hsm_destination_client_certificates", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClientCertificateCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateInstanceHsmDestinationClientCertificate : Generate client certificate (CSR)
// Generate an RSA private key and CSR. CSR is immediately available in the response. Max 10 inactive/csr_generated
// certificates — returns `409 cert_limit_exceeded` if exceeded.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptions *CreateInstanceHsmDestinationClientCertificateOptions) (result *ClientCertificate, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateInstanceHsmDestinationClientCertificateWithContext(context.Background(), createInstanceHsmDestinationClientCertificateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateInstanceHsmDestinationClientCertificateWithContext is an alternate form of the CreateInstanceHsmDestinationClientCertificate method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceHsmDestinationClientCertificateWithContext(ctx context.Context, createInstanceHsmDestinationClientCertificateOptions *CreateInstanceHsmDestinationClientCertificateOptions) (result *ClientCertificate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createInstanceHsmDestinationClientCertificateOptions, "createInstanceHsmDestinationClientCertificateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createInstanceHsmDestinationClientCertificateOptions, "createInstanceHsmDestinationClientCertificateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createInstanceHsmDestinationClientCertificateOptions.InstanceID,
		"destination_id": *createInstanceHsmDestinationClientCertificateOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/client_certificates`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "CreateInstanceHsmDestinationClientCertificate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createInstanceHsmDestinationClientCertificateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createInstanceHsmDestinationClientCertificateOptions.Name != nil {
		body["name"] = createInstanceHsmDestinationClientCertificateOptions.Name
	}
	if createInstanceHsmDestinationClientCertificateOptions.Description != nil {
		body["description"] = createInstanceHsmDestinationClientCertificateOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_instance_hsm_destination_client_certificate", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClientCertificate)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceHsmDestinationClientCertificate : Get client certificate
// Retrieve a single client certificate by its ID.
//
// **Rate Limit**: 100 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptions *GetInstanceHsmDestinationClientCertificateOptions) (result *ClientCertificate, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.GetInstanceHsmDestinationClientCertificateWithContext(context.Background(), getInstanceHsmDestinationClientCertificateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceHsmDestinationClientCertificateWithContext is an alternate form of the GetInstanceHsmDestinationClientCertificate method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceHsmDestinationClientCertificateWithContext(ctx context.Context, getInstanceHsmDestinationClientCertificateOptions *GetInstanceHsmDestinationClientCertificateOptions) (result *ClientCertificate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceHsmDestinationClientCertificateOptions, "getInstanceHsmDestinationClientCertificateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getInstanceHsmDestinationClientCertificateOptions, "getInstanceHsmDestinationClientCertificateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceHsmDestinationClientCertificateOptions.InstanceID,
		"destination_id": *getInstanceHsmDestinationClientCertificateOptions.DestinationID,
		"cert_id": *getInstanceHsmDestinationClientCertificateOptions.CertID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/client_certificates/{cert_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "GetInstanceHsmDestinationClientCertificate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInstanceHsmDestinationClientCertificateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance_hsm_destination_client_certificate", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClientCertificate)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateInstanceHsmDestinationClientCertificate : Update, upload signed cert, and/or activate
// Update name/description, upload `signed_cert`, or activate (merge-patch semantics). `signed_cert` must match the
// generated CSR — returns `400 csr_mismatch` otherwise.
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptions *UpdateInstanceHsmDestinationClientCertificateOptions) (result *ClientCertificate, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.UpdateInstanceHsmDestinationClientCertificateWithContext(context.Background(), updateInstanceHsmDestinationClientCertificateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateInstanceHsmDestinationClientCertificateWithContext is an alternate form of the UpdateInstanceHsmDestinationClientCertificate method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceHsmDestinationClientCertificateWithContext(ctx context.Context, updateInstanceHsmDestinationClientCertificateOptions *UpdateInstanceHsmDestinationClientCertificateOptions) (result *ClientCertificate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceHsmDestinationClientCertificateOptions, "updateInstanceHsmDestinationClientCertificateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateInstanceHsmDestinationClientCertificateOptions, "updateInstanceHsmDestinationClientCertificateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceHsmDestinationClientCertificateOptions.InstanceID,
		"destination_id": *updateInstanceHsmDestinationClientCertificateOptions.DestinationID,
		"cert_id": *updateInstanceHsmDestinationClientCertificateOptions.CertID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/client_certificates/{cert_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "UpdateInstanceHsmDestinationClientCertificate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateInstanceHsmDestinationClientCertificateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceHsmDestinationClientCertificateOptions.ClientCertificatePatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = secretsManagerInstanceManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_instance_hsm_destination_client_certificate", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClientCertificate)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteInstanceHsmDestinationClientCertificate : Delete client certificate
// Delete a client certificate and purge its private key from Vault. Cannot delete the active certificate — returns `409
// Conflict`. Synchronous (204, no body).
//
// **Rate Limit**: 5 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptions *DeleteInstanceHsmDestinationClientCertificateOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.DeleteInstanceHsmDestinationClientCertificateWithContext(context.Background(), deleteInstanceHsmDestinationClientCertificateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceHsmDestinationClientCertificateWithContext is an alternate form of the DeleteInstanceHsmDestinationClientCertificate method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceHsmDestinationClientCertificateWithContext(ctx context.Context, deleteInstanceHsmDestinationClientCertificateOptions *DeleteInstanceHsmDestinationClientCertificateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteInstanceHsmDestinationClientCertificateOptions, "deleteInstanceHsmDestinationClientCertificateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteInstanceHsmDestinationClientCertificateOptions, "deleteInstanceHsmDestinationClientCertificateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteInstanceHsmDestinationClientCertificateOptions.InstanceID,
		"destination_id": *deleteInstanceHsmDestinationClientCertificateOptions.DestinationID,
		"cert_id": *deleteInstanceHsmDestinationClientCertificateOptions.CertID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{instance_id}/hsm_destinations/{destination_id}/client_certificates/{cert_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteInstanceHsmDestinationClientCertificate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteInstanceHsmDestinationClientCertificateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_hsm_destination_client_certificate", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "2.0.0")
}

// ApplyHsmDestinationOptions : The ApplyHsmDestination options.
type ApplyHsmDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewApplyHsmDestinationOptions : Instantiate ApplyHsmDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewApplyHsmDestinationOptions(instanceID string, destinationID string) *ApplyHsmDestinationOptions {
	return &ApplyHsmDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ApplyHsmDestinationOptions) SetInstanceID(instanceID string) *ApplyHsmDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *ApplyHsmDestinationOptions) SetDestinationID(destinationID string) *ApplyHsmDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ApplyHsmDestinationOptions) SetHeaders(param map[string]string) *ApplyHsmDestinationOptions {
	options.Headers = param
	return options
}

// AsyncDeleteResponse : Response body for async DELETE operations (202 Accepted). The `type` field is intentionally omitted — the caller
// knows which resource they deleted.
type AsyncDeleteResponse struct {
	// URL of the resource being deleted.
	Href *string `json:"href,omitempty"`

	// ID of the resource being deleted.
	ID *string `json:"id,omitempty"`

	// Lifecycle state of the deleted resource. Always `deleting` on a 202 response.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`
}

// Constants associated with the AsyncDeleteResponse.ProvisioningState property.
// Lifecycle state of the deleted resource. Always `deleting` on a 202 response.
const (
	AsyncDeleteResponse_ProvisioningState_Deleting = "deleting"
)

// UnmarshalAsyncDeleteResponse unmarshals an instance of AsyncDeleteResponse from the specified map of raw messages.
func UnmarshalAsyncDeleteResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AsyncDeleteResponse)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsyncHealthAckResponse : Response body for the async apply operation (202 Accepted). The body is currently an empty object;
// `href`, when present, is the `GET …/health` polling endpoint. Poll `GET …/health` until `health_state` is no longer
// `in_progress`.
type AsyncHealthAckResponse struct {
	// URL of the GET …/health snapshot endpoint to poll for results. Not currently returned.
	Href *string `json:"href,omitempty"`
}

// UnmarshalAsyncHealthAckResponse unmarshals an instance of AsyncHealthAckResponse from the specified map of raw messages.
func UnmarshalAsyncHealthAckResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AsyncHealthAckResponse)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsyncHealthCheckAckResponse : Response body for the async destination health-check operation (202 Accepted). Contains the `href` of the `GET
// …/health` polling endpoint and the health state known at acceptance time. Sub-resource arrays are empty when none
// exist. Poll `GET …/health` until `health_state` is no longer `evaluating`.
type AsyncHealthCheckAckResponse struct {
	// URL of the GET …/health snapshot endpoint to poll for results.
	Href *string `json:"href,omitempty"`

	// ID of the destination.
	DestinationID *string `json:"destination_id,omitempty"`

	// AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.
	// - `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.
	// - `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check
	// result once apply completes (or set to `unknown` on failure).
	// - `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1
	// active non-expired client certificate).
	// - `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.
	// - `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.
	// - `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.
	HealthState *string `json:"health_state" validate:"required"`

	// Timestamp of the health state, when available.
	HealthAsOf *strfmt.DateTime `json:"health_as_of,omitempty"`

	// Client certificates for this destination. Empty array if none exist.
	ClientCertificates []HealthSnapshotClientCert `json:"client_certificates" validate:"required"`

	// HSM HA Groups for this destination. Empty array if none exist.
	HaGroups []HealthSnapshotHSMHAGroup `json:"ha_groups" validate:"required"`

	// Paths for this destination. Empty array if none exist.
	Paths []HealthSnapshotConnection `json:"paths" validate:"required"`
}

// Constants associated with the AsyncHealthCheckAckResponse.HealthState property.
// AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.
// - `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.
// - `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result
// once apply completes (or set to `unknown` on failure).
// - `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1
// active non-expired client certificate).
// - `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.
// - `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.
// - `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.
const (
	AsyncHealthCheckAckResponse_HealthState_Degraded = "degraded"
	AsyncHealthCheckAckResponse_HealthState_Evaluating = "evaluating"
	AsyncHealthCheckAckResponse_HealthState_Healthy = "healthy"
	AsyncHealthCheckAckResponse_HealthState_InProgress = "in_progress"
	AsyncHealthCheckAckResponse_HealthState_NotConfigured = "not_configured"
	AsyncHealthCheckAckResponse_HealthState_Unhealthy = "unhealthy"
	AsyncHealthCheckAckResponse_HealthState_Unknown = "unknown"
)

// UnmarshalAsyncHealthCheckAckResponse unmarshals an instance of AsyncHealthCheckAckResponse from the specified map of raw messages.
func UnmarshalAsyncHealthCheckAckResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AsyncHealthCheckAckResponse)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_id", &obj.DestinationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_as_of", &obj.HealthAsOf)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_as_of-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "client_certificates", &obj.ClientCertificates, UnmarshalHealthSnapshotClientCert)
	if err != nil {
		err = core.SDKErrorf(err, "", "client_certificates-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "ha_groups", &obj.HaGroups, UnmarshalHealthSnapshotHSMHAGroup)
	if err != nil {
		err = core.SDKErrorf(err, "", "ha_groups-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "paths", &obj.Paths, UnmarshalHealthSnapshotConnection)
	if err != nil {
		err = core.SDKErrorf(err, "", "paths-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClientCertificate : A client certificate resource. csr, signed_cert, and expires_at are top-level fields. There is no vendor_config
// wrapper. CSR and private key are platform-generated. CN (`csr_cn`) is set to the SM instance UUID. The private key is
// stored securely and is NOT exposed in any API response. Max 10 inactive/csr_generated certificates at any time.
// Active certificate cannot be deleted.
// `csr` and `csr_cn` are populated on creation.
type ClientCertificate struct {
	// Certificate ID. Immutable.
	ID *string `json:"id,omitempty"`

	// URL of the client certificate resource.
	Href *string `json:"href,omitempty"`

	// Human-readable display name.
	Name *string `json:"name" validate:"required"`

	// Optional description. Omitted if not set.
	Description *string `json:"description,omitempty"`

	// Request-only. Set to `active` in a PATCH to activate this certificate. Never returned; read the current state from
	// `provisioning_state`.
	Status *string `json:"status,omitempty"`

	// Lifecycle state of the client certificate:
	// - `csr_generated`: CSR and private key generated successfully. `csr` and `csr_cn` are populated. Initial state on
	// POST. Waiting for signed_cert to be uploaded via PATCH.
	// - `inactive`: signed_cert has been uploaded but the certificate has not been activated yet.
	// - `active`: This certificate is the active client certificate for the destination. Only one can be active at a time.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`

	// AS OF health state of the client certificate. Written exclusively by POST …/health:
	// - `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.
	// - `evaluating`: A health check is in progress for this client certificate.
	// - `valid`: Certificate is within its validity period (expires_at is in the future).
	// - `expired`: Certificate has passed its expires_at date.
	HealthState *string `json:"health_state" validate:"required"`

	// Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
	HealthAsOf *strfmt.DateTime `json:"health_as_of" validate:"required"`

	// PEM-encoded CSR. Platform-generated. Immutable once set. Populated on creation (always present once the resource
	// exists).
	Csr *string `json:"csr" validate:"required"`

	// Common Name embedded in the CSR. Set to the SM instance UUID at generation time. Immutable once set. Populated on
	// creation (always present once the resource exists).
	CsrCn *string `json:"csr_cn" validate:"required"`

	// CA-signed certificate (PEM). Omitted until uploaded via PATCH.
	SignedCert *string `json:"signed_cert,omitempty"`

	// Parsed from signed_cert NotAfter field. Omitted until signed_cert is uploaded.
	ExpiresAt *strfmt.DateTime `json:"expires_at,omitempty"`

	// Computed field.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// Computed field.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Constants associated with the ClientCertificate.ProvisioningState property.
// Lifecycle state of the client certificate:
// - `csr_generated`: CSR and private key generated successfully. `csr` and `csr_cn` are populated. Initial state on
// POST. Waiting for signed_cert to be uploaded via PATCH.
// - `inactive`: signed_cert has been uploaded but the certificate has not been activated yet.
// - `active`: This certificate is the active client certificate for the destination. Only one can be active at a time.
const (
	ClientCertificate_ProvisioningState_Active = "active"
	ClientCertificate_ProvisioningState_CsrGenerated = "csr_generated"
	ClientCertificate_ProvisioningState_Inactive = "inactive"
)

// Constants associated with the ClientCertificate.HealthState property.
// AS OF health state of the client certificate. Written exclusively by POST …/health:
// - `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.
// - `evaluating`: A health check is in progress for this client certificate.
// - `valid`: Certificate is within its validity period (expires_at is in the future).
// - `expired`: Certificate has passed its expires_at date.
const (
	ClientCertificate_HealthState_Evaluating = "evaluating"
	ClientCertificate_HealthState_Expired = "expired"
	ClientCertificate_HealthState_Unknown = "unknown"
	ClientCertificate_HealthState_Valid = "valid"
)

// UnmarshalClientCertificate unmarshals an instance of ClientCertificate from the specified map of raw messages.
func UnmarshalClientCertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClientCertificate)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_as_of", &obj.HealthAsOf)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_as_of-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "csr", &obj.Csr)
	if err != nil {
		err = core.SDKErrorf(err, "", "csr-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "csr_cn", &obj.CsrCn)
	if err != nil {
		err = core.SDKErrorf(err, "", "csr_cn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "signed_cert", &obj.SignedCert)
	if err != nil {
		err = core.SDKErrorf(err, "", "signed_cert-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "expires_at", &obj.ExpiresAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "expires_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_by-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClientCertificateCollection : A collection of client certificates.
type ClientCertificateCollection struct {
	// The list of client certificates for this destination.
	ClientCertificates []ClientCertificate `json:"client_certificates" validate:"required"`
}

// UnmarshalClientCertificateCollection unmarshals an instance of ClientCertificateCollection from the specified map of raw messages.
func UnmarshalClientCertificateCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClientCertificateCollection)
	err = core.UnmarshalModel(m, "client_certificates", &obj.ClientCertificates, UnmarshalClientCertificate)
	if err != nil {
		err = core.SDKErrorf(err, "", "client_certificates-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClientCertificatePatch : Fields to update on the client certificate. Only supplied fields are changed (merge-patch semantics).
type ClientCertificatePatch struct {
	// Updated display name for the certificate.
	Name *string `json:"name,omitempty"`

	// Updated description. Set to empty string to clear.
	Description *string `json:"description,omitempty"`

	// Set to `active` to activate this certificate. Only `active` is allowed.
	Status *string `json:"status,omitempty"`

	// CA-signed certificate (PEM). Makes cert eligible for activation.
	SignedCert *string `json:"signed_cert,omitempty"`
}

// UnmarshalClientCertificatePatch unmarshals an instance of ClientCertificatePatch from the specified map of raw messages.
func UnmarshalClientCertificatePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClientCertificatePatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "signed_cert", &obj.SignedCert)
	if err != nil {
		err = core.SDKErrorf(err, "", "signed_cert-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the ClientCertificatePatch
func (clientCertificatePatch *ClientCertificatePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(clientCertificatePatch.Name) {
		_patch["name"] = clientCertificatePatch.Name
	}
	if !core.IsNil(clientCertificatePatch.Description) {
		_patch["description"] = clientCertificatePatch.Description
	}
	if !core.IsNil(clientCertificatePatch.Status) {
		_patch["status"] = clientCertificatePatch.Status
	}
	if !core.IsNil(clientCertificatePatch.SignedCert) {
		_patch["signed_cert"] = clientCertificatePatch.SignedCert
	}

	return
}

// CreateDestinationRequest : Request body for creating a destination.
// Models which "extend" this model:
// - CreateDestinationRequestIbmCloudDatabaseDestinationPrototype
type CreateDestinationRequest struct {
	// Destination name.
	Name *string `json:"name" validate:"required"`

	// Destination type.
	Type *string `json:"type" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// IBM Cloud Database service instance CRN.
	Crn *string `json:"crn" validate:"required"`
}

// Constants associated with the CreateDestinationRequest.Type property.
// Destination type.
const (
	CreateDestinationRequest_Type_IbmCloudDatabase = "ibm_cloud_database"
)
func (*CreateDestinationRequest) isaCreateDestinationRequest() bool {
	return true
}

type CreateDestinationRequestIntf interface {
	isaCreateDestinationRequest() bool
}

// UnmarshalCreateDestinationRequest unmarshals an instance of CreateDestinationRequest from the specified map of raw messages.
func UnmarshalCreateDestinationRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", common.GetComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'type' not found in JSON object", "missing-discriminator", common.GetComponentInfo())
		return
	}
	switch discValue {
	case "ibm_cloud_database":
		err = core.UnmarshalModel(m, "", result, UnmarshalCreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-CreateDestinationRequestIbmCloudDatabaseDestinationPrototype-error", common.GetComponentInfo())
		}
	default:
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", common.GetComponentInfo())
	}
	return
}

// CreateInstanceDestinationOptions : The CreateInstanceDestination options.
type CreateInstanceDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Request body for creating a destination.
	DestinationPrototype CreateDestinationRequestIntf `json:"destinationPrototype" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateInstanceDestinationOptions : Instantiate CreateInstanceDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewCreateInstanceDestinationOptions(instanceID string, destinationPrototype CreateDestinationRequestIntf) *CreateInstanceDestinationOptions {
	return &CreateInstanceDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationPrototype: destinationPrototype,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceDestinationOptions) SetInstanceID(instanceID string) *CreateInstanceDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationPrototype : Allow user to set DestinationPrototype
func (_options *CreateInstanceDestinationOptions) SetDestinationPrototype(destinationPrototype CreateDestinationRequestIntf) *CreateInstanceDestinationOptions {
	_options.DestinationPrototype = destinationPrototype
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceDestinationOptions) SetHeaders(param map[string]string) *CreateInstanceDestinationOptions {
	options.Headers = param
	return options
}

// CreateInstanceHsmDestinationClientCertificateOptions : The CreateInstanceHsmDestinationClientCertificate options.
type CreateInstanceHsmDestinationClientCertificateOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Human-readable display name. Required.
	Name *string `json:"name" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateInstanceHsmDestinationClientCertificateOptions : Instantiate CreateInstanceHsmDestinationClientCertificateOptions
func (*SecretsManagerInstanceManagementV2) NewCreateInstanceHsmDestinationClientCertificateOptions(instanceID string, destinationID string, name string) *CreateInstanceHsmDestinationClientCertificateOptions {
	return &CreateInstanceHsmDestinationClientCertificateOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		Name: core.StringPtr(name),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceHsmDestinationClientCertificateOptions) SetInstanceID(instanceID string) *CreateInstanceHsmDestinationClientCertificateOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *CreateInstanceHsmDestinationClientCertificateOptions) SetDestinationID(destinationID string) *CreateInstanceHsmDestinationClientCertificateOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateInstanceHsmDestinationClientCertificateOptions) SetName(name string) *CreateInstanceHsmDestinationClientCertificateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateInstanceHsmDestinationClientCertificateOptions) SetDescription(description string) *CreateInstanceHsmDestinationClientCertificateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceHsmDestinationClientCertificateOptions) SetHeaders(param map[string]string) *CreateInstanceHsmDestinationClientCertificateOptions {
	options.Headers = param
	return options
}

// CreateInstanceHsmDestinationHaGroupOptions : The CreateInstanceHsmDestinationHaGroup options.
type CreateInstanceHsmDestinationHaGroupOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Human-readable display name. Required.
	Name *string `json:"name" validate:"required"`

	// HA Group virtual token label. Immutable after creation.
	GroupLabel *string `json:"group_label" validate:"required"`

	// HA Group serial number. Immutable after creation.
	GroupSerial *string `json:"group_serial" validate:"required"`

	// Ordered list of HSM node member serial numbers for this HA Group. At least one entry required. Max 100 unique
	// members. Not linked to path resource IDs.
	MemberSerials []string `json:"member_serials" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateInstanceHsmDestinationHaGroupOptions : Instantiate CreateInstanceHsmDestinationHaGroupOptions
func (*SecretsManagerInstanceManagementV2) NewCreateInstanceHsmDestinationHaGroupOptions(instanceID string, destinationID string, name string, groupLabel string, groupSerial string, memberSerials []string) *CreateInstanceHsmDestinationHaGroupOptions {
	return &CreateInstanceHsmDestinationHaGroupOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		Name: core.StringPtr(name),
		GroupLabel: core.StringPtr(groupLabel),
		GroupSerial: core.StringPtr(groupSerial),
		MemberSerials: memberSerials,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceHsmDestinationHaGroupOptions) SetInstanceID(instanceID string) *CreateInstanceHsmDestinationHaGroupOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *CreateInstanceHsmDestinationHaGroupOptions) SetDestinationID(destinationID string) *CreateInstanceHsmDestinationHaGroupOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateInstanceHsmDestinationHaGroupOptions) SetName(name string) *CreateInstanceHsmDestinationHaGroupOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetGroupLabel : Allow user to set GroupLabel
func (_options *CreateInstanceHsmDestinationHaGroupOptions) SetGroupLabel(groupLabel string) *CreateInstanceHsmDestinationHaGroupOptions {
	_options.GroupLabel = core.StringPtr(groupLabel)
	return _options
}

// SetGroupSerial : Allow user to set GroupSerial
func (_options *CreateInstanceHsmDestinationHaGroupOptions) SetGroupSerial(groupSerial string) *CreateInstanceHsmDestinationHaGroupOptions {
	_options.GroupSerial = core.StringPtr(groupSerial)
	return _options
}

// SetMemberSerials : Allow user to set MemberSerials
func (_options *CreateInstanceHsmDestinationHaGroupOptions) SetMemberSerials(memberSerials []string) *CreateInstanceHsmDestinationHaGroupOptions {
	_options.MemberSerials = memberSerials
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateInstanceHsmDestinationHaGroupOptions) SetDescription(description string) *CreateInstanceHsmDestinationHaGroupOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceHsmDestinationHaGroupOptions) SetHeaders(param map[string]string) *CreateInstanceHsmDestinationHaGroupOptions {
	options.Headers = param
	return options
}

// CreateInstanceHsmDestinationOptions : The CreateInstanceHsmDestination options.
type CreateInstanceHsmDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Human-readable display name. Required.
	Name *string `json:"name" validate:"required"`

	// HSM vendor. Immutable after creation. Currently only `thales_luna` is accepted.
	Vendor *string `json:"vendor" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateInstanceHsmDestinationOptions.Vendor property.
// HSM vendor. Immutable after creation. Currently only `thales_luna` is accepted.
const (
	CreateInstanceHsmDestinationOptions_Vendor_ThalesLuna = "thales_luna"
)

// NewCreateInstanceHsmDestinationOptions : Instantiate CreateInstanceHsmDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewCreateInstanceHsmDestinationOptions(instanceID string, name string, vendor string) *CreateInstanceHsmDestinationOptions {
	return &CreateInstanceHsmDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		Name: core.StringPtr(name),
		Vendor: core.StringPtr(vendor),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceHsmDestinationOptions) SetInstanceID(instanceID string) *CreateInstanceHsmDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateInstanceHsmDestinationOptions) SetName(name string) *CreateInstanceHsmDestinationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetVendor : Allow user to set Vendor
func (_options *CreateInstanceHsmDestinationOptions) SetVendor(vendor string) *CreateInstanceHsmDestinationOptions {
	_options.Vendor = core.StringPtr(vendor)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateInstanceHsmDestinationOptions) SetDescription(description string) *CreateInstanceHsmDestinationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceHsmDestinationOptions) SetHeaders(param map[string]string) *CreateInstanceHsmDestinationOptions {
	options.Headers = param
	return options
}

// CreateInstanceHsmDestinationPathOptions : The CreateInstanceHsmDestinationPath options.
type CreateInstanceHsmDestinationPathOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Human-readable display name. Required.
	Name *string `json:"name" validate:"required"`

	// CRN of the gateway through which the HSM is reached. Immutable after creation.
	Crn *string `json:"crn" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateInstanceHsmDestinationPathOptions : Instantiate CreateInstanceHsmDestinationPathOptions
func (*SecretsManagerInstanceManagementV2) NewCreateInstanceHsmDestinationPathOptions(instanceID string, destinationID string, name string, crn string) *CreateInstanceHsmDestinationPathOptions {
	return &CreateInstanceHsmDestinationPathOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		Name: core.StringPtr(name),
		Crn: core.StringPtr(crn),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceHsmDestinationPathOptions) SetInstanceID(instanceID string) *CreateInstanceHsmDestinationPathOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *CreateInstanceHsmDestinationPathOptions) SetDestinationID(destinationID string) *CreateInstanceHsmDestinationPathOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateInstanceHsmDestinationPathOptions) SetName(name string) *CreateInstanceHsmDestinationPathOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCrn : Allow user to set Crn
func (_options *CreateInstanceHsmDestinationPathOptions) SetCrn(crn string) *CreateInstanceHsmDestinationPathOptions {
	_options.Crn = core.StringPtr(crn)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateInstanceHsmDestinationPathOptions) SetDescription(description string) *CreateInstanceHsmDestinationPathOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceHsmDestinationPathOptions) SetHeaders(param map[string]string) *CreateInstanceHsmDestinationPathOptions {
	options.Headers = param
	return options
}

// CreateInstanceHsmDestinationPathServerCaOptions : The CreateInstanceHsmDestinationPathServerCa options.
type CreateInstanceHsmDestinationPathServerCaOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// Human-readable display name. Required.
	Name *string `json:"name" validate:"required"`

	// PEM-encoded Server CA certificate. Immutable after creation.
	Certificate *string `json:"certificate" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateInstanceHsmDestinationPathServerCaOptions : Instantiate CreateInstanceHsmDestinationPathServerCaOptions
func (*SecretsManagerInstanceManagementV2) NewCreateInstanceHsmDestinationPathServerCaOptions(instanceID string, destinationID string, pathID string, name string, certificate string) *CreateInstanceHsmDestinationPathServerCaOptions {
	return &CreateInstanceHsmDestinationPathServerCaOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
		Name: core.StringPtr(name),
		Certificate: core.StringPtr(certificate),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceHsmDestinationPathServerCaOptions) SetInstanceID(instanceID string) *CreateInstanceHsmDestinationPathServerCaOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *CreateInstanceHsmDestinationPathServerCaOptions) SetDestinationID(destinationID string) *CreateInstanceHsmDestinationPathServerCaOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *CreateInstanceHsmDestinationPathServerCaOptions) SetPathID(pathID string) *CreateInstanceHsmDestinationPathServerCaOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateInstanceHsmDestinationPathServerCaOptions) SetName(name string) *CreateInstanceHsmDestinationPathServerCaOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCertificate : Allow user to set Certificate
func (_options *CreateInstanceHsmDestinationPathServerCaOptions) SetCertificate(certificate string) *CreateInstanceHsmDestinationPathServerCaOptions {
	_options.Certificate = core.StringPtr(certificate)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateInstanceHsmDestinationPathServerCaOptions) SetDescription(description string) *CreateInstanceHsmDestinationPathServerCaOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceHsmDestinationPathServerCaOptions) SetHeaders(param map[string]string) *CreateInstanceHsmDestinationPathServerCaOptions {
	options.Headers = param
	return options
}

// CreatePathBuildOptions : The CreatePathBuild options.
type CreatePathBuildOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreatePathBuildOptions : Instantiate CreatePathBuildOptions
func (*SecretsManagerInstanceManagementV2) NewCreatePathBuildOptions(instanceID string, destinationID string, pathID string) *CreatePathBuildOptions {
	return &CreatePathBuildOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreatePathBuildOptions) SetInstanceID(instanceID string) *CreatePathBuildOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *CreatePathBuildOptions) SetDestinationID(destinationID string) *CreatePathBuildOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *CreatePathBuildOptions) SetPathID(pathID string) *CreatePathBuildOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePathBuildOptions) SetHeaders(param map[string]string) *CreatePathBuildOptions {
	options.Headers = param
	return options
}

// CreateVaultAdmintokenOptions : The CreateVaultAdmintoken options.
type CreateVaultAdmintokenOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// If true, returns a Vault response-wrapped token (wrapped_token). If false or absent, returns a plain admin token
	// (token).
	ResponseWrapping *bool `json:"response_wrapping,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateVaultAdmintokenOptions : Instantiate CreateVaultAdmintokenOptions
func (*SecretsManagerInstanceManagementV2) NewCreateVaultAdmintokenOptions(instanceID string) *CreateVaultAdmintokenOptions {
	return &CreateVaultAdmintokenOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateVaultAdmintokenOptions) SetInstanceID(instanceID string) *CreateVaultAdmintokenOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResponseWrapping : Allow user to set ResponseWrapping
func (_options *CreateVaultAdmintokenOptions) SetResponseWrapping(responseWrapping bool) *CreateVaultAdmintokenOptions {
	_options.ResponseWrapping = core.BoolPtr(responseWrapping)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVaultAdmintokenOptions) SetHeaders(param map[string]string) *CreateVaultAdmintokenOptions {
	options.Headers = param
	return options
}

// DeleteInstanceDestinationOptions : The DeleteInstanceDestination options.
type DeleteInstanceDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceDestinationOptions : Instantiate DeleteInstanceDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceDestinationOptions(instanceID string, destinationID string) *DeleteInstanceDestinationOptions {
	return &DeleteInstanceDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteInstanceDestinationOptions) SetInstanceID(instanceID string) *DeleteInstanceDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteInstanceDestinationOptions) SetDestinationID(destinationID string) *DeleteInstanceDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceDestinationOptions) SetHeaders(param map[string]string) *DeleteInstanceDestinationOptions {
	options.Headers = param
	return options
}

// DeleteInstanceHsmDestinationClientCertificateOptions : The DeleteInstanceHsmDestinationClientCertificate options.
type DeleteInstanceHsmDestinationClientCertificateOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// The identifier of the Client Certificate to operate on. Matches the `id` field in the `ClientCertificate` response
	// schema.
	CertID *string `json:"cert_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceHsmDestinationClientCertificateOptions : Instantiate DeleteInstanceHsmDestinationClientCertificateOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceHsmDestinationClientCertificateOptions(instanceID string, destinationID string, certID string) *DeleteInstanceHsmDestinationClientCertificateOptions {
	return &DeleteInstanceHsmDestinationClientCertificateOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		CertID: core.StringPtr(certID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteInstanceHsmDestinationClientCertificateOptions) SetInstanceID(instanceID string) *DeleteInstanceHsmDestinationClientCertificateOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteInstanceHsmDestinationClientCertificateOptions) SetDestinationID(destinationID string) *DeleteInstanceHsmDestinationClientCertificateOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetCertID : Allow user to set CertID
func (_options *DeleteInstanceHsmDestinationClientCertificateOptions) SetCertID(certID string) *DeleteInstanceHsmDestinationClientCertificateOptions {
	_options.CertID = core.StringPtr(certID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceHsmDestinationClientCertificateOptions) SetHeaders(param map[string]string) *DeleteInstanceHsmDestinationClientCertificateOptions {
	options.Headers = param
	return options
}

// DeleteInstanceHsmDestinationHaGroupOptions : The DeleteInstanceHsmDestinationHaGroup options.
type DeleteInstanceHsmDestinationHaGroupOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// The identifier of the HA Group to operate on. Matches the `id` field in the `HSMHAGroup` response schema.
	HaGroupID *string `json:"ha_group_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceHsmDestinationHaGroupOptions : Instantiate DeleteInstanceHsmDestinationHaGroupOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceHsmDestinationHaGroupOptions(instanceID string, destinationID string, haGroupID string) *DeleteInstanceHsmDestinationHaGroupOptions {
	return &DeleteInstanceHsmDestinationHaGroupOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		HaGroupID: core.StringPtr(haGroupID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteInstanceHsmDestinationHaGroupOptions) SetInstanceID(instanceID string) *DeleteInstanceHsmDestinationHaGroupOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteInstanceHsmDestinationHaGroupOptions) SetDestinationID(destinationID string) *DeleteInstanceHsmDestinationHaGroupOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHaGroupID : Allow user to set HaGroupID
func (_options *DeleteInstanceHsmDestinationHaGroupOptions) SetHaGroupID(haGroupID string) *DeleteInstanceHsmDestinationHaGroupOptions {
	_options.HaGroupID = core.StringPtr(haGroupID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceHsmDestinationHaGroupOptions) SetHeaders(param map[string]string) *DeleteInstanceHsmDestinationHaGroupOptions {
	options.Headers = param
	return options
}

// DeleteInstanceHsmDestinationOptions : The DeleteInstanceHsmDestination options.
type DeleteInstanceHsmDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceHsmDestinationOptions : Instantiate DeleteInstanceHsmDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceHsmDestinationOptions(instanceID string, destinationID string) *DeleteInstanceHsmDestinationOptions {
	return &DeleteInstanceHsmDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteInstanceHsmDestinationOptions) SetInstanceID(instanceID string) *DeleteInstanceHsmDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteInstanceHsmDestinationOptions) SetDestinationID(destinationID string) *DeleteInstanceHsmDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceHsmDestinationOptions) SetHeaders(param map[string]string) *DeleteInstanceHsmDestinationOptions {
	options.Headers = param
	return options
}

// DeleteInstanceHsmDestinationPathOptions : The DeleteInstanceHsmDestinationPath options.
type DeleteInstanceHsmDestinationPathOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceHsmDestinationPathOptions : Instantiate DeleteInstanceHsmDestinationPathOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceHsmDestinationPathOptions(instanceID string, destinationID string, pathID string) *DeleteInstanceHsmDestinationPathOptions {
	return &DeleteInstanceHsmDestinationPathOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteInstanceHsmDestinationPathOptions) SetInstanceID(instanceID string) *DeleteInstanceHsmDestinationPathOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteInstanceHsmDestinationPathOptions) SetDestinationID(destinationID string) *DeleteInstanceHsmDestinationPathOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *DeleteInstanceHsmDestinationPathOptions) SetPathID(pathID string) *DeleteInstanceHsmDestinationPathOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceHsmDestinationPathOptions) SetHeaders(param map[string]string) *DeleteInstanceHsmDestinationPathOptions {
	options.Headers = param
	return options
}

// DeleteInstanceHsmDestinationPathServerCaOptions : The DeleteInstanceHsmDestinationPathServerCa options.
type DeleteInstanceHsmDestinationPathServerCaOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// The identifier of the Server CA to operate on. Matches the `id` field in the `ServerCa` response schema.
	ServerCaID *string `json:"server_ca_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceHsmDestinationPathServerCaOptions : Instantiate DeleteInstanceHsmDestinationPathServerCaOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceHsmDestinationPathServerCaOptions(instanceID string, destinationID string, pathID string, serverCaID string) *DeleteInstanceHsmDestinationPathServerCaOptions {
	return &DeleteInstanceHsmDestinationPathServerCaOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
		ServerCaID: core.StringPtr(serverCaID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteInstanceHsmDestinationPathServerCaOptions) SetInstanceID(instanceID string) *DeleteInstanceHsmDestinationPathServerCaOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteInstanceHsmDestinationPathServerCaOptions) SetDestinationID(destinationID string) *DeleteInstanceHsmDestinationPathServerCaOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *DeleteInstanceHsmDestinationPathServerCaOptions) SetPathID(pathID string) *DeleteInstanceHsmDestinationPathServerCaOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetServerCaID : Allow user to set ServerCaID
func (_options *DeleteInstanceHsmDestinationPathServerCaOptions) SetServerCaID(serverCaID string) *DeleteInstanceHsmDestinationPathServerCaOptions {
	_options.ServerCaID = core.StringPtr(serverCaID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceHsmDestinationPathServerCaOptions) SetHeaders(param map[string]string) *DeleteInstanceHsmDestinationPathServerCaOptions {
	options.Headers = param
	return options
}

// DeleteVaultAdmintokensOptions : The DeleteVaultAdmintokens options.
type DeleteVaultAdmintokensOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteVaultAdmintokensOptions : Instantiate DeleteVaultAdmintokensOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteVaultAdmintokensOptions(instanceID string) *DeleteVaultAdmintokensOptions {
	return &DeleteVaultAdmintokensOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteVaultAdmintokensOptions) SetInstanceID(instanceID string) *DeleteVaultAdmintokensOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVaultAdmintokensOptions) SetHeaders(param map[string]string) *DeleteVaultAdmintokensOptions {
	options.Headers = param
	return options
}

// Destination : A destination resource representing a private network link to a service instance on a Vault Dedicated cluster.
type Destination struct {
	// A standard UUID string (RFC 4122).
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The URL of the destination resource.
	Href *string `json:"href,omitempty"`

	// Destination name.
	Name *string `json:"name" validate:"required"`

	// Destination type.
	Type *string `json:"type" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// Destination state:
	// - `not_started`: Initial state before the first provisioning attempt begins
	// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
	// - `succeeded`: Destination ready and usable
	// - `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against
	// the per-instance quota until deleted.
	// - `deleting`: Deletion in progress.
	State *string `json:"state" validate:"required"`

	// Timestamp when the destination was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Timestamp when the destination was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// IAM identity that created the destination.
	CreatedBy *string `json:"created_by,omitempty"`

	// Human-readable message providing additional context about the current state. Present only when non-empty — set when
	// `state` is `failed`, describing why provisioning or deletion failed.
	Message *string `json:"message,omitempty"`
}

// Constants associated with the Destination.Type property.
// Destination type.
const (
	Destination_Type_IbmCloudDatabase = "ibm_cloud_database"
)

// Constants associated with the Destination.State property.
// Destination state:
// - `not_started`: Initial state before the first provisioning attempt begins
// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
// - `succeeded`: Destination ready and usable
// - `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against
// the per-instance quota until deleted.
// - `deleting`: Deletion in progress.
const (
	Destination_State_Deleting = "deleting"
	Destination_State_Failed = "failed"
	Destination_State_NotStarted = "not_started"
	Destination_State_Provisioning = "provisioning"
	Destination_State_Succeeded = "succeeded"
)

// UnmarshalDestination unmarshals an instance of Destination from the specified map of raw messages.
func UnmarshalDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Destination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DestinationCollection : List of destinations for a Vault Dedicated cluster.
type DestinationCollection struct {
	// List of destinations.
	Destinations []Destination `json:"destinations" validate:"required"`

	// Total number of destinations. Maximum 20 per instance.
	Total *int64 `json:"total" validate:"required"`
}

// UnmarshalDestinationCollection unmarshals an instance of DestinationCollection from the specified map of raw messages.
func UnmarshalDestinationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DestinationCollection)
	err = core.UnmarshalModel(m, "destinations", &obj.Destinations, UnmarshalDestination)
	if err != nil {
		err = core.SDKErrorf(err, "", "destinations-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		err = core.SDKErrorf(err, "", "total-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DestinationHealthSnapshot : Persisted AS-OF health snapshot for a destination. Returned by GET /hsm_destinations/{destination_id}/health and by
// the async POST /hsm_destinations/{destination_id}/health and POST /hsm_destinations/{destination_id}/apply flows
// (poll until health_state is no longer `evaluating` or `in_progress`). All health_state values reflect the last
// completed evaluation.
type DestinationHealthSnapshot struct {
	// URL of the health snapshot endpoint for this destination.
	Href *string `json:"href,omitempty"`

	// ID of the destination. Echoed from path param.
	DestinationID *string `json:"destination_id" validate:"required"`

	// AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.
	// - `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.
	// - `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check
	// result once apply completes (or set to `unknown` on failure).
	// - `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1
	// active non-expired client certificate).
	// - `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.
	// - `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.
	// - `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.
	HealthState *string `json:"health_state" validate:"required"`

	// Timestamp of this health check — always the current time of the live probe.
	HealthAsOf *strfmt.DateTime `json:"health_as_of" validate:"required"`

	// All client certificates for this destination (active, inactive, csr_generated). health_state freshly evaluated for
	// all certs with a signed_cert.
	ClientCertificates []HealthSnapshotClientCert `json:"client_certificates" validate:"required"`

	// All HSM HA Groups for this destination. Empty array if none exist.
	HaGroups []HealthSnapshotHSMHAGroup `json:"ha_groups" validate:"required"`

	// All paths for this destination. Empty array if none exist.
	Paths []HealthSnapshotConnection `json:"paths" validate:"required"`
}

// Constants associated with the DestinationHealthSnapshot.HealthState property.
// AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.
// - `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.
// - `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result
// once apply completes (or set to `unknown` on failure).
// - `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1
// active non-expired client certificate).
// - `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.
// - `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.
// - `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.
const (
	DestinationHealthSnapshot_HealthState_Degraded = "degraded"
	DestinationHealthSnapshot_HealthState_Evaluating = "evaluating"
	DestinationHealthSnapshot_HealthState_Healthy = "healthy"
	DestinationHealthSnapshot_HealthState_InProgress = "in_progress"
	DestinationHealthSnapshot_HealthState_NotConfigured = "not_configured"
	DestinationHealthSnapshot_HealthState_Unhealthy = "unhealthy"
	DestinationHealthSnapshot_HealthState_Unknown = "unknown"
)

// UnmarshalDestinationHealthSnapshot unmarshals an instance of DestinationHealthSnapshot from the specified map of raw messages.
func UnmarshalDestinationHealthSnapshot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DestinationHealthSnapshot)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_id", &obj.DestinationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_as_of", &obj.HealthAsOf)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_as_of-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "client_certificates", &obj.ClientCertificates, UnmarshalHealthSnapshotClientCert)
	if err != nil {
		err = core.SDKErrorf(err, "", "client_certificates-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "ha_groups", &obj.HaGroups, UnmarshalHealthSnapshotHSMHAGroup)
	if err != nil {
		err = core.SDKErrorf(err, "", "ha_groups-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "paths", &obj.Paths, UnmarshalHealthSnapshotConnection)
	if err != nil {
		err = core.SDKErrorf(err, "", "paths-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DestinationPatch : Fields to update on a destination. At least one field must be provided.
type DestinationPatch struct {
	// Updated name.
	Name *string `json:"name,omitempty"`

	// Updated description.
	Description *string `json:"description,omitempty"`
}

// UnmarshalDestinationPatch unmarshals an instance of DestinationPatch from the specified map of raw messages.
func UnmarshalDestinationPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DestinationPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the DestinationPatch
func (destinationPatch *DestinationPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(destinationPatch.Name) {
		_patch["name"] = destinationPatch.Name
	}
	if !core.IsNil(destinationPatch.Description) {
		_patch["description"] = destinationPatch.Description
	}

	return
}

// GetHsmDestinationHealthSnapshotOptions : The GetHsmDestinationHealthSnapshot options.
type GetHsmDestinationHealthSnapshotOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetHsmDestinationHealthSnapshotOptions : Instantiate GetHsmDestinationHealthSnapshotOptions
func (*SecretsManagerInstanceManagementV2) NewGetHsmDestinationHealthSnapshotOptions(instanceID string, destinationID string) *GetHsmDestinationHealthSnapshotOptions {
	return &GetHsmDestinationHealthSnapshotOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetHsmDestinationHealthSnapshotOptions) SetInstanceID(instanceID string) *GetHsmDestinationHealthSnapshotOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetHsmDestinationHealthSnapshotOptions) SetDestinationID(destinationID string) *GetHsmDestinationHealthSnapshotOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetHsmDestinationHealthSnapshotOptions) SetHeaders(param map[string]string) *GetHsmDestinationHealthSnapshotOptions {
	options.Headers = param
	return options
}

// GetInstanceDestinationOptions : The GetInstanceDestination options.
type GetInstanceDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceDestinationOptions : Instantiate GetInstanceDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceDestinationOptions(instanceID string, destinationID string) *GetInstanceDestinationOptions {
	return &GetInstanceDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceDestinationOptions) SetInstanceID(instanceID string) *GetInstanceDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetInstanceDestinationOptions) SetDestinationID(destinationID string) *GetInstanceDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceDestinationOptions) SetHeaders(param map[string]string) *GetInstanceDestinationOptions {
	options.Headers = param
	return options
}

// GetInstanceHsmDestinationClientCertificateOptions : The GetInstanceHsmDestinationClientCertificate options.
type GetInstanceHsmDestinationClientCertificateOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// The identifier of the Client Certificate to operate on. Matches the `id` field in the `ClientCertificate` response
	// schema.
	CertID *string `json:"cert_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceHsmDestinationClientCertificateOptions : Instantiate GetInstanceHsmDestinationClientCertificateOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceHsmDestinationClientCertificateOptions(instanceID string, destinationID string, certID string) *GetInstanceHsmDestinationClientCertificateOptions {
	return &GetInstanceHsmDestinationClientCertificateOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		CertID: core.StringPtr(certID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceHsmDestinationClientCertificateOptions) SetInstanceID(instanceID string) *GetInstanceHsmDestinationClientCertificateOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetInstanceHsmDestinationClientCertificateOptions) SetDestinationID(destinationID string) *GetInstanceHsmDestinationClientCertificateOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetCertID : Allow user to set CertID
func (_options *GetInstanceHsmDestinationClientCertificateOptions) SetCertID(certID string) *GetInstanceHsmDestinationClientCertificateOptions {
	_options.CertID = core.StringPtr(certID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceHsmDestinationClientCertificateOptions) SetHeaders(param map[string]string) *GetInstanceHsmDestinationClientCertificateOptions {
	options.Headers = param
	return options
}

// GetInstanceHsmDestinationHaGroupOptions : The GetInstanceHsmDestinationHaGroup options.
type GetInstanceHsmDestinationHaGroupOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// The identifier of the HA Group to operate on. Matches the `id` field in the `HSMHAGroup` response schema.
	HaGroupID *string `json:"ha_group_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceHsmDestinationHaGroupOptions : Instantiate GetInstanceHsmDestinationHaGroupOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceHsmDestinationHaGroupOptions(instanceID string, destinationID string, haGroupID string) *GetInstanceHsmDestinationHaGroupOptions {
	return &GetInstanceHsmDestinationHaGroupOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		HaGroupID: core.StringPtr(haGroupID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceHsmDestinationHaGroupOptions) SetInstanceID(instanceID string) *GetInstanceHsmDestinationHaGroupOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetInstanceHsmDestinationHaGroupOptions) SetDestinationID(destinationID string) *GetInstanceHsmDestinationHaGroupOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHaGroupID : Allow user to set HaGroupID
func (_options *GetInstanceHsmDestinationHaGroupOptions) SetHaGroupID(haGroupID string) *GetInstanceHsmDestinationHaGroupOptions {
	_options.HaGroupID = core.StringPtr(haGroupID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceHsmDestinationHaGroupOptions) SetHeaders(param map[string]string) *GetInstanceHsmDestinationHaGroupOptions {
	options.Headers = param
	return options
}

// GetInstanceHsmDestinationOptions : The GetInstanceHsmDestination options.
type GetInstanceHsmDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceHsmDestinationOptions : Instantiate GetInstanceHsmDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceHsmDestinationOptions(instanceID string, destinationID string) *GetInstanceHsmDestinationOptions {
	return &GetInstanceHsmDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceHsmDestinationOptions) SetInstanceID(instanceID string) *GetInstanceHsmDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetInstanceHsmDestinationOptions) SetDestinationID(destinationID string) *GetInstanceHsmDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceHsmDestinationOptions) SetHeaders(param map[string]string) *GetInstanceHsmDestinationOptions {
	options.Headers = param
	return options
}

// GetInstanceHsmDestinationPathOptions : The GetInstanceHsmDestinationPath options.
type GetInstanceHsmDestinationPathOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceHsmDestinationPathOptions : Instantiate GetInstanceHsmDestinationPathOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceHsmDestinationPathOptions(instanceID string, destinationID string, pathID string) *GetInstanceHsmDestinationPathOptions {
	return &GetInstanceHsmDestinationPathOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceHsmDestinationPathOptions) SetInstanceID(instanceID string) *GetInstanceHsmDestinationPathOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetInstanceHsmDestinationPathOptions) SetDestinationID(destinationID string) *GetInstanceHsmDestinationPathOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *GetInstanceHsmDestinationPathOptions) SetPathID(pathID string) *GetInstanceHsmDestinationPathOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceHsmDestinationPathOptions) SetHeaders(param map[string]string) *GetInstanceHsmDestinationPathOptions {
	options.Headers = param
	return options
}

// GetInstanceHsmDestinationPathServerCaOptions : The GetInstanceHsmDestinationPathServerCa options.
type GetInstanceHsmDestinationPathServerCaOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// The identifier of the Server CA to operate on. Matches the `id` field in the `ServerCa` response schema.
	ServerCaID *string `json:"server_ca_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceHsmDestinationPathServerCaOptions : Instantiate GetInstanceHsmDestinationPathServerCaOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceHsmDestinationPathServerCaOptions(instanceID string, destinationID string, pathID string, serverCaID string) *GetInstanceHsmDestinationPathServerCaOptions {
	return &GetInstanceHsmDestinationPathServerCaOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
		ServerCaID: core.StringPtr(serverCaID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceHsmDestinationPathServerCaOptions) SetInstanceID(instanceID string) *GetInstanceHsmDestinationPathServerCaOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetInstanceHsmDestinationPathServerCaOptions) SetDestinationID(destinationID string) *GetInstanceHsmDestinationPathServerCaOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *GetInstanceHsmDestinationPathServerCaOptions) SetPathID(pathID string) *GetInstanceHsmDestinationPathServerCaOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetServerCaID : Allow user to set ServerCaID
func (_options *GetInstanceHsmDestinationPathServerCaOptions) SetServerCaID(serverCaID string) *GetInstanceHsmDestinationPathServerCaOptions {
	_options.ServerCaID = core.StringPtr(serverCaID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceHsmDestinationPathServerCaOptions) SetHeaders(param map[string]string) *GetInstanceHsmDestinationPathServerCaOptions {
	options.Headers = param
	return options
}

// GetInstanceOptions : The GetInstance options.
type GetInstanceOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceOptions : Instantiate GetInstanceOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceOptions(instanceID string) *GetInstanceOptions {
	return &GetInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceOptions) SetInstanceID(instanceID string) *GetInstanceOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceOptions) SetHeaders(param map[string]string) *GetInstanceOptions {
	options.Headers = param
	return options
}

// GetPathVpeStatusOptions : The GetPathVpeStatus options.
type GetPathVpeStatusOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPathVpeStatusOptions : Instantiate GetPathVpeStatusOptions
func (*SecretsManagerInstanceManagementV2) NewGetPathVpeStatusOptions(instanceID string, destinationID string, pathID string) *GetPathVpeStatusOptions {
	return &GetPathVpeStatusOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetPathVpeStatusOptions) SetInstanceID(instanceID string) *GetPathVpeStatusOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetPathVpeStatusOptions) SetDestinationID(destinationID string) *GetPathVpeStatusOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *GetPathVpeStatusOptions) SetPathID(pathID string) *GetPathVpeStatusOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPathVpeStatusOptions) SetHeaders(param map[string]string) *GetPathVpeStatusOptions {
	options.Headers = param
	return options
}

// HSMDestination : An HSM Destination resource.
type HSMDestination struct {
	// Destination ID. Immutable. System-managed.
	ID *string `json:"id,omitempty"`

	// URL of the destination resource.
	Href *string `json:"href,omitempty"`

	// IBM Cloud account ID. System-managed. Required for account policy setup.
	AccountID *string `json:"account_id,omitempty"`

	// Human-readable display name.
	Name *string `json:"name" validate:"required"`

	// Optional description. Omitted if not set.
	Description *string `json:"description,omitempty"`

	// HSM vendor. Immutable after creation.
	Vendor *string `json:"vendor" validate:"required"`

	// Lifecycle state of the destination:
	// - `completed`: Destination created and available for configuration. Initial state on POST.
	// - `deleting`: Hard delete initiated; asynchronous cleanup in progress.
	// - `deletion_failed`: Asynchronous delete failed; the destination still exists. Retry DELETE to re-initiate.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`

	// AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.
	// - `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.
	// - `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check
	// result once apply completes (or set to `unknown` on failure).
	// - `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1
	// active non-expired client certificate).
	// - `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.
	// - `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.
	// - `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.
	HealthState *string `json:"health_state" validate:"required"`

	// Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
	HealthAsOf *strfmt.DateTime `json:"health_as_of" validate:"required"`

	// Computed field. When the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// IAM identity of creator. Computed field.
	CreatedBy *string `json:"created_by" validate:"required"`

	// IAM identity of last updater. Computed field. Omitted until first update.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Constants associated with the HSMDestination.Vendor property.
// HSM vendor. Immutable after creation.
const (
	HSMDestination_Vendor_ThalesLuna = "thales_luna"
)

// Constants associated with the HSMDestination.ProvisioningState property.
// Lifecycle state of the destination:
// - `completed`: Destination created and available for configuration. Initial state on POST.
// - `deleting`: Hard delete initiated; asynchronous cleanup in progress.
// - `deletion_failed`: Asynchronous delete failed; the destination still exists. Retry DELETE to re-initiate.
const (
	HSMDestination_ProvisioningState_Completed = "completed"
	HSMDestination_ProvisioningState_Deleting = "deleting"
	HSMDestination_ProvisioningState_DeletionFailed = "deletion_failed"
)

// Constants associated with the HSMDestination.HealthState property.
// AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.
// - `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.
// - `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result
// once apply completes (or set to `unknown` on failure).
// - `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1
// active non-expired client certificate).
// - `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.
// - `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.
// - `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.
const (
	HSMDestination_HealthState_Degraded = "degraded"
	HSMDestination_HealthState_Evaluating = "evaluating"
	HSMDestination_HealthState_Healthy = "healthy"
	HSMDestination_HealthState_InProgress = "in_progress"
	HSMDestination_HealthState_NotConfigured = "not_configured"
	HSMDestination_HealthState_Unhealthy = "unhealthy"
	HSMDestination_HealthState_Unknown = "unknown"
)

// UnmarshalHSMDestination unmarshals an instance of HSMDestination from the specified map of raw messages.
func UnmarshalHSMDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HSMDestination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		err = core.SDKErrorf(err, "", "account_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "vendor", &obj.Vendor)
	if err != nil {
		err = core.SDKErrorf(err, "", "vendor-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_as_of", &obj.HealthAsOf)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_as_of-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_by-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HSMDestinationCollection : A collection of HSM Destinations.
type HSMDestinationCollection struct {
	// Total number of HSM Destinations in the collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The list of HSM Destinations. At most 5 destinations may exist per SM instance.
	HsmDestinations []HSMDestination `json:"hsm_destinations" validate:"required"`
}

// UnmarshalHSMDestinationCollection unmarshals an instance of HSMDestinationCollection from the specified map of raw messages.
func UnmarshalHSMDestinationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HSMDestinationCollection)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "hsm_destinations", &obj.HsmDestinations, UnmarshalHSMDestination)
	if err != nil {
		err = core.SDKErrorf(err, "", "hsm_destinations-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HSMDestinationPatch : Fields to update on the HSM Destination. Only supplied fields are changed (merge-patch semantics).
type HSMDestinationPatch struct {
	// Updated display name for the destination.
	Name *string `json:"name,omitempty"`

	// Updated description for the destination. Set to empty string to clear.
	Description *string `json:"description,omitempty"`
}

// UnmarshalHSMDestinationPatch unmarshals an instance of HSMDestinationPatch from the specified map of raw messages.
func UnmarshalHSMDestinationPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HSMDestinationPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the HSMDestinationPatch
func (hSMDestinationPatch *HSMDestinationPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(hSMDestinationPatch.Name) {
		_patch["name"] = hSMDestinationPatch.Name
	}
	if !core.IsNil(hSMDestinationPatch.Description) {
		_patch["description"] = hSMDestinationPatch.Description
	}

	return
}

// HSMHAGroup : An HA Group resource. All fields are top-level — no vendor_config wrapper.
type HSMHAGroup struct {
	// HA Group ID. Immutable. System-managed.
	ID *string `json:"id,omitempty"`

	// URL of the HA Group resource.
	Href *string `json:"href,omitempty"`

	// Human-readable display name.
	Name *string `json:"name" validate:"required"`

	// Optional description. Omitted if not set.
	Description *string `json:"description,omitempty"`

	// Lifecycle state of the HA Group:
	// - `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only
	// state.
	//
	// **Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values
	// gracefully and MUST NOT fail on unrecognised states.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`

	// AS OF health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Initial state on creation. No health check has been performed yet.
	// - `evaluating`: A health check is in progress for this resource.
	// - `in_progress`: An apply operation is running across the destination.
	// - `healthy`: Configured AND all member HSM nodes are healthy.
	// - `degraded`: Configured AND any one member is unhealthy (but not all).
	// - `unhealthy`: Configured AND all members are unhealthy or unknown.
	// - `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active
	// non-expired client certificate).
	HealthState *string `json:"health_state" validate:"required"`

	// Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
	HealthAsOf *strfmt.DateTime `json:"health_as_of" validate:"required"`

	// HA Group virtual token label. Immutable after creation.
	GroupLabel *string `json:"group_label" validate:"required"`

	// HA Group serial number. Immutable after creation.
	GroupSerial *string `json:"group_serial" validate:"required"`

	// Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members.
	MemberSerials []string `json:"member_serials" validate:"required"`

	// Computed field.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// Computed field.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Constants associated with the HSMHAGroup.ProvisioningState property.
// Lifecycle state of the HA Group:
// - `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only
// state.
//
// **Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values
// gracefully and MUST NOT fail on unrecognised states.
const (
	HSMHAGroup_ProvisioningState_Completed = "completed"
)

// Constants associated with the HSMHAGroup.HealthState property.
// AS OF health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Initial state on creation. No health check has been performed yet.
// - `evaluating`: A health check is in progress for this resource.
// - `in_progress`: An apply operation is running across the destination.
// - `healthy`: Configured AND all member HSM nodes are healthy.
// - `degraded`: Configured AND any one member is unhealthy (but not all).
// - `unhealthy`: Configured AND all members are unhealthy or unknown.
// - `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active
// non-expired client certificate).
const (
	HSMHAGroup_HealthState_Degraded = "degraded"
	HSMHAGroup_HealthState_Evaluating = "evaluating"
	HSMHAGroup_HealthState_Healthy = "healthy"
	HSMHAGroup_HealthState_InProgress = "in_progress"
	HSMHAGroup_HealthState_NotConfigured = "not_configured"
	HSMHAGroup_HealthState_Unhealthy = "unhealthy"
	HSMHAGroup_HealthState_Unknown = "unknown"
)

// UnmarshalHSMHAGroup unmarshals an instance of HSMHAGroup from the specified map of raw messages.
func UnmarshalHSMHAGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HSMHAGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_as_of", &obj.HealthAsOf)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_as_of-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "group_label", &obj.GroupLabel)
	if err != nil {
		err = core.SDKErrorf(err, "", "group_label-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "group_serial", &obj.GroupSerial)
	if err != nil {
		err = core.SDKErrorf(err, "", "group_serial-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "member_serials", &obj.MemberSerials)
	if err != nil {
		err = core.SDKErrorf(err, "", "member_serials-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_by-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HSMHAGroupCollection : A collection of HSM HA Groups.
type HSMHAGroupCollection struct {
	// HSM HA Groups for this destination. At most 20 HSM HA Groups may exist per destination.
	HaGroups []HSMHAGroup `json:"ha_groups" validate:"required"`
}

// UnmarshalHSMHAGroupCollection unmarshals an instance of HSMHAGroupCollection from the specified map of raw messages.
func UnmarshalHSMHAGroupCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HSMHAGroupCollection)
	err = core.UnmarshalModel(m, "ha_groups", &obj.HaGroups, UnmarshalHSMHAGroup)
	if err != nil {
		err = core.SDKErrorf(err, "", "ha_groups-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HSMHAGroupPatch : Request body for updating an HA Group. group_serial and group_label are immutable — return 400 if supplied. If name
// is omitted, existing name is preserved (merge-patch semantics). name cannot be set to null. HSM Config File is staged
// when member_serials is changed.
type HSMHAGroupPatch struct {
	// Updated display name for the HA Group.
	Name *string `json:"name,omitempty"`

	// Updated description for the HA Group. Set to empty string to clear.
	Description *string `json:"description,omitempty"`

	// Update the member serial list. At least one entry required if supplied. Max 100 unique serials. Stages HSM config
	// file change.
	MemberSerials []string `json:"member_serials,omitempty"`
}

// UnmarshalHSMHAGroupPatch unmarshals an instance of HSMHAGroupPatch from the specified map of raw messages.
func UnmarshalHSMHAGroupPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HSMHAGroupPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "member_serials", &obj.MemberSerials)
	if err != nil {
		err = core.SDKErrorf(err, "", "member_serials-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the HSMHAGroupPatch
func (hSMHAGroupPatch *HSMHAGroupPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(hSMHAGroupPatch.Name) {
		_patch["name"] = hSMHAGroupPatch.Name
	}
	if !core.IsNil(hSMHAGroupPatch.Description) {
		_patch["description"] = hSMHAGroupPatch.Description
	}
	if !core.IsNil(hSMHAGroupPatch.MemberSerials) {
		_patch["member_serials"] = hSMHAGroupPatch.MemberSerials
	}

	return
}

// HSMPathReset : Response body for POST .../paths/{path_id}/reset.
type HSMPathReset struct {
	// URL of the path resource.
	Href *string `json:"href,omitempty"`

	// Lifecycle state of the path:
	// - `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.
	// - `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy
	// for the gateway before proceeding.
	// - `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.
	// - `pending_approval`: Connectivity created. Waiting for customer approval.
	// - `completed`: Connectivity established. Path fully operational.
	// - `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.
	// - `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager
	// instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.
	// - `approval_declined`: Customer declined the pending approval. Call POST /reset.
	// - `deleting`: Delete initiated; async cleanup in progress.
	// - `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`
}

// Constants associated with the HSMPathReset.ProvisioningState property.
// Lifecycle state of the path:
// - `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.
// - `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy for
// the gateway before proceeding.
// - `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.
// - `pending_approval`: Connectivity created. Waiting for customer approval.
// - `completed`: Connectivity established. Path fully operational.
// - `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.
// - `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager
// instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.
// - `approval_declined`: Customer declined the pending approval. Call POST /reset.
// - `deleting`: Delete initiated; async cleanup in progress.
// - `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.
const (
	HSMPathReset_ProvisioningState_ApprovalDeclined = "approval_declined"
	HSMPathReset_ProvisioningState_Completed = "completed"
	HSMPathReset_ProvisioningState_CrnAlreadyInUse = "crn_already_in_use"
	HSMPathReset_ProvisioningState_Deleting = "deleting"
	HSMPathReset_ProvisioningState_DeletionFailed = "deletion_failed"
	HSMPathReset_ProvisioningState_PendingAccountPolicy = "pending_account_policy"
	HSMPathReset_ProvisioningState_PendingApproval = "pending_approval"
	HSMPathReset_ProvisioningState_PendingServerCa = "pending_server_ca"
	HSMPathReset_ProvisioningState_Provisioning = "provisioning"
	HSMPathReset_ProvisioningState_VpeProvisioningFailed = "vpe_provisioning_failed"
)

// UnmarshalHSMPathReset unmarshals an instance of HSMPathReset from the specified map of raw messages.
func UnmarshalHSMPathReset(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HSMPathReset)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HealthSnapshotClientCert : Client certificate summary within a DestinationHealthSnapshot.
type HealthSnapshotClientCert struct {
	// Client certificate ID.
	ID *string `json:"id" validate:"required"`

	// Human-readable display name.
	Name *string `json:"name,omitempty"`

	// AS OF health state of the client certificate. Written exclusively by POST …/health:
	// - `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.
	// - `evaluating`: A health check is in progress for this client certificate.
	// - `valid`: Certificate is within its validity period (expires_at is in the future).
	// - `expired`: Certificate has passed its expires_at date.
	HealthState *string `json:"health_state" validate:"required"`
}

// Constants associated with the HealthSnapshotClientCert.HealthState property.
// AS OF health state of the client certificate. Written exclusively by POST …/health:
// - `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.
// - `evaluating`: A health check is in progress for this client certificate.
// - `valid`: Certificate is within its validity period (expires_at is in the future).
// - `expired`: Certificate has passed its expires_at date.
const (
	HealthSnapshotClientCert_HealthState_Evaluating = "evaluating"
	HealthSnapshotClientCert_HealthState_Expired = "expired"
	HealthSnapshotClientCert_HealthState_Unknown = "unknown"
	HealthSnapshotClientCert_HealthState_Valid = "valid"
)

// UnmarshalHealthSnapshotClientCert unmarshals an instance of HealthSnapshotClientCert from the specified map of raw messages.
func UnmarshalHealthSnapshotClientCert(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HealthSnapshotClientCert)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HealthSnapshotConnection : Path summary within a DestinationHealthSnapshot.
type HealthSnapshotConnection struct {
	// Path ID.
	ID *string `json:"id" validate:"required"`

	// Human-readable display name.
	Name *string `json:"name,omitempty"`

	// AS OF health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.
	// - `evaluating`: A health check is in progress for this path.
	// - `in_progress`: An apply operation is running across the destination.
	// - `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.
	// - `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval,
	// approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is
	// missing or expired, no_valid_server_ca takes priority.
	// - `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.
	// - `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.
	HealthState *string `json:"health_state" validate:"required"`

	// All Server CAs for this path (active and inactive). Empty array if none.
	ServerCas []HealthSnapshotServerCa `json:"server_cas" validate:"required"`
}

// Constants associated with the HealthSnapshotConnection.HealthState property.
// AS OF health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.
// - `evaluating`: A health check is in progress for this path.
// - `in_progress`: An apply operation is running across the destination.
// - `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.
// - `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval,
// approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is
// missing or expired, no_valid_server_ca takes priority.
// - `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.
// - `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.
const (
	HealthSnapshotConnection_HealthState_Evaluating = "evaluating"
	HealthSnapshotConnection_HealthState_Healthy = "healthy"
	HealthSnapshotConnection_HealthState_InProgress = "in_progress"
	HealthSnapshotConnection_HealthState_NoValidServerCa = "no_valid_server_ca"
	HealthSnapshotConnection_HealthState_NotConfigured = "not_configured"
	HealthSnapshotConnection_HealthState_Unhealthy = "unhealthy"
	HealthSnapshotConnection_HealthState_Unknown = "unknown"
)

// UnmarshalHealthSnapshotConnection unmarshals an instance of HealthSnapshotConnection from the specified map of raw messages.
func UnmarshalHealthSnapshotConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HealthSnapshotConnection)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "server_cas", &obj.ServerCas, UnmarshalHealthSnapshotServerCa)
	if err != nil {
		err = core.SDKErrorf(err, "", "server_cas-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HealthSnapshotHSMHAGroup : HA Group summary within a DestinationHealthSnapshot.
type HealthSnapshotHSMHAGroup struct {
	// HA Group ID.
	ID *string `json:"id" validate:"required"`

	// Human-readable display name.
	Name *string `json:"name,omitempty"`

	// AS OF health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Initial state on creation. No health check has been performed yet.
	// - `evaluating`: A health check is in progress for this resource.
	// - `in_progress`: An apply operation is running across the destination.
	// - `healthy`: Configured AND all member HSM nodes are healthy.
	// - `degraded`: Configured AND any one member is unhealthy (but not all).
	// - `unhealthy`: Configured AND all members are unhealthy or unknown.
	// - `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active
	// non-expired client certificate).
	HealthState *string `json:"health_state" validate:"required"`
}

// Constants associated with the HealthSnapshotHSMHAGroup.HealthState property.
// AS OF health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Initial state on creation. No health check has been performed yet.
// - `evaluating`: A health check is in progress for this resource.
// - `in_progress`: An apply operation is running across the destination.
// - `healthy`: Configured AND all member HSM nodes are healthy.
// - `degraded`: Configured AND any one member is unhealthy (but not all).
// - `unhealthy`: Configured AND all members are unhealthy or unknown.
// - `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active
// non-expired client certificate).
const (
	HealthSnapshotHSMHAGroup_HealthState_Degraded = "degraded"
	HealthSnapshotHSMHAGroup_HealthState_Evaluating = "evaluating"
	HealthSnapshotHSMHAGroup_HealthState_Healthy = "healthy"
	HealthSnapshotHSMHAGroup_HealthState_InProgress = "in_progress"
	HealthSnapshotHSMHAGroup_HealthState_NotConfigured = "not_configured"
	HealthSnapshotHSMHAGroup_HealthState_Unhealthy = "unhealthy"
	HealthSnapshotHSMHAGroup_HealthState_Unknown = "unknown"
)

// UnmarshalHealthSnapshotHSMHAGroup unmarshals an instance of HealthSnapshotHSMHAGroup from the specified map of raw messages.
func UnmarshalHealthSnapshotHSMHAGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HealthSnapshotHSMHAGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HealthSnapshotServerCa : Server CA summary within a HealthSnapshotConnection.
type HealthSnapshotServerCa struct {
	// Server CA ID.
	ID *string `json:"id" validate:"required"`

	// Human-readable display name.
	Name *string `json:"name,omitempty"`

	// AS OF health state of the Server CA certificate. Written exclusively by POST …/health:
	// - `unknown`: Initial state on creation. Health check not yet performed.
	// - `evaluating`: A health check is in progress for this Server CA.
	// - `valid`: Certificate is within its validity period (expires_at is in the future).
	// - `expired`: Certificate has passed its expires_at date.
	HealthState *string `json:"health_state" validate:"required"`
}

// Constants associated with the HealthSnapshotServerCa.HealthState property.
// AS OF health state of the Server CA certificate. Written exclusively by POST …/health:
// - `unknown`: Initial state on creation. Health check not yet performed.
// - `evaluating`: A health check is in progress for this Server CA.
// - `valid`: Certificate is within its validity period (expires_at is in the future).
// - `expired`: Certificate has passed its expires_at date.
const (
	HealthSnapshotServerCa_HealthState_Evaluating = "evaluating"
	HealthSnapshotServerCa_HealthState_Expired = "expired"
	HealthSnapshotServerCa_HealthState_Unknown = "unknown"
	HealthSnapshotServerCa_HealthState_Valid = "valid"
)

// UnmarshalHealthSnapshotServerCa unmarshals an instance of HealthSnapshotServerCa from the specified map of raw messages.
func UnmarshalHealthSnapshotServerCa(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HealthSnapshotServerCa)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IbmCloudDatabaseDestination : A destination resource representing a private network link to an IBM Cloud Database service instance on a Vault
// Dedicated cluster.
type IbmCloudDatabaseDestination struct {
	// A standard UUID string (RFC 4122).
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The URL of the destination resource.
	Href *string `json:"href,omitempty"`

	// Destination name.
	Name *string `json:"name" validate:"required"`

	// Destination type.
	Type *string `json:"type" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// Destination state:
	// - `not_started`: Initial state before the first provisioning attempt begins
	// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
	// - `succeeded`: Destination ready and usable
	// - `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against
	// the per-instance quota until deleted.
	// - `deleting`: Deletion in progress.
	State *string `json:"state" validate:"required"`

	// Timestamp when the destination was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Timestamp when the destination was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// IAM identity that created the destination.
	CreatedBy *string `json:"created_by,omitempty"`

	// Human-readable message providing additional context about the current state. Present only when non-empty — set when
	// `state` is `failed`, describing why provisioning or deletion failed.
	Message *string `json:"message,omitempty"`

	// IBM Cloud Database service instance CRN.
	Crn *string `json:"crn" validate:"required"`
}

// Constants associated with the IbmCloudDatabaseDestination.Type property.
// Destination type.
const (
	IbmCloudDatabaseDestination_Type_IbmCloudDatabase = "ibm_cloud_database"
)

// Constants associated with the IbmCloudDatabaseDestination.State property.
// Destination state:
// - `not_started`: Initial state before the first provisioning attempt begins
// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
// - `succeeded`: Destination ready and usable
// - `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against
// the per-instance quota until deleted.
// - `deleting`: Deletion in progress.
const (
	IbmCloudDatabaseDestination_State_Deleting = "deleting"
	IbmCloudDatabaseDestination_State_Failed = "failed"
	IbmCloudDatabaseDestination_State_NotStarted = "not_started"
	IbmCloudDatabaseDestination_State_Provisioning = "provisioning"
	IbmCloudDatabaseDestination_State_Succeeded = "succeeded"
)

// UnmarshalIbmCloudDatabaseDestination unmarshals an instance of IbmCloudDatabaseDestination from the specified map of raw messages.
func UnmarshalIbmCloudDatabaseDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IbmCloudDatabaseDestination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		err = core.SDKErrorf(err, "", "crn-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Instance : The service instance information.
type Instance struct {
	// A standard UUID string (RFC 4122).
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The instance name.
	Name *string `json:"name" validate:"required"`

	// The instance CRN identifier.
	InstanceCrn *string `json:"instance_crn" validate:"required"`

	// Instance plan name.
	Plan *string `json:"plan" validate:"required"`

	// Vault cluster information for Vault Dedicated instances.
	VaultCluster *VaultDedicatedCluster `json:"vault_cluster" validate:"required"`

	// Instance endpoints for Vault Dedicated instances.
	Endpoints *VaultDedicatedInstanceEndpoints `json:"endpoints" validate:"required"`

	// Vault encryption configuration for Vault Dedicated instances.
	Encryption *VaultDedicatedInstanceEncryption `json:"encryption" validate:"required"`

	// The URL of the instance resource.
	Href *string `json:"href,omitempty"`
}

// Constants associated with the Instance.Plan property.
// Instance plan name.
const (
	Instance_Plan_Dedicated = "dedicated"
)

// UnmarshalInstance unmarshals an instance of Instance from the specified map of raw messages.
func UnmarshalInstance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Instance)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_crn", &obj.InstanceCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "instance_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "plan", &obj.Plan)
	if err != nil {
		err = core.SDKErrorf(err, "", "plan-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "vault_cluster", &obj.VaultCluster, UnmarshalVaultDedicatedCluster)
	if err != nil {
		err = core.SDKErrorf(err, "", "vault_cluster-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "endpoints", &obj.Endpoints, UnmarshalVaultDedicatedInstanceEndpoints)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoints-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "encryption", &obj.Encryption, UnmarshalVaultDedicatedInstanceEncryption)
	if err != nil {
		err = core.SDKErrorf(err, "", "encryption-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListInstanceDestinationsOptions : The ListInstanceDestinations options.
type ListInstanceDestinationsOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Filter by destination state.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the ListInstanceDestinationsOptions.State property.
// Filter by destination state.
const (
	ListInstanceDestinationsOptions_State_Deleting = "deleting"
	ListInstanceDestinationsOptions_State_Failed = "failed"
	ListInstanceDestinationsOptions_State_NotStarted = "not_started"
	ListInstanceDestinationsOptions_State_Provisioning = "provisioning"
	ListInstanceDestinationsOptions_State_Succeeded = "succeeded"
)

// NewListInstanceDestinationsOptions : Instantiate ListInstanceDestinationsOptions
func (*SecretsManagerInstanceManagementV2) NewListInstanceDestinationsOptions(instanceID string) *ListInstanceDestinationsOptions {
	return &ListInstanceDestinationsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListInstanceDestinationsOptions) SetInstanceID(instanceID string) *ListInstanceDestinationsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetState : Allow user to set State
func (_options *ListInstanceDestinationsOptions) SetState(state string) *ListInstanceDestinationsOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListInstanceDestinationsOptions) SetHeaders(param map[string]string) *ListInstanceDestinationsOptions {
	options.Headers = param
	return options
}

// ListInstanceHsmDestinationClientCertificatesOptions : The ListInstanceHsmDestinationClientCertificates options.
type ListInstanceHsmDestinationClientCertificatesOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListInstanceHsmDestinationClientCertificatesOptions : Instantiate ListInstanceHsmDestinationClientCertificatesOptions
func (*SecretsManagerInstanceManagementV2) NewListInstanceHsmDestinationClientCertificatesOptions(instanceID string, destinationID string) *ListInstanceHsmDestinationClientCertificatesOptions {
	return &ListInstanceHsmDestinationClientCertificatesOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListInstanceHsmDestinationClientCertificatesOptions) SetInstanceID(instanceID string) *ListInstanceHsmDestinationClientCertificatesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *ListInstanceHsmDestinationClientCertificatesOptions) SetDestinationID(destinationID string) *ListInstanceHsmDestinationClientCertificatesOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListInstanceHsmDestinationClientCertificatesOptions) SetHeaders(param map[string]string) *ListInstanceHsmDestinationClientCertificatesOptions {
	options.Headers = param
	return options
}

// ListInstanceHsmDestinationHaGroupsOptions : The ListInstanceHsmDestinationHaGroups options.
type ListInstanceHsmDestinationHaGroupsOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListInstanceHsmDestinationHaGroupsOptions : Instantiate ListInstanceHsmDestinationHaGroupsOptions
func (*SecretsManagerInstanceManagementV2) NewListInstanceHsmDestinationHaGroupsOptions(instanceID string, destinationID string) *ListInstanceHsmDestinationHaGroupsOptions {
	return &ListInstanceHsmDestinationHaGroupsOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListInstanceHsmDestinationHaGroupsOptions) SetInstanceID(instanceID string) *ListInstanceHsmDestinationHaGroupsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *ListInstanceHsmDestinationHaGroupsOptions) SetDestinationID(destinationID string) *ListInstanceHsmDestinationHaGroupsOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListInstanceHsmDestinationHaGroupsOptions) SetHeaders(param map[string]string) *ListInstanceHsmDestinationHaGroupsOptions {
	options.Headers = param
	return options
}

// ListInstanceHsmDestinationPathServerCasOptions : The ListInstanceHsmDestinationPathServerCas options.
type ListInstanceHsmDestinationPathServerCasOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListInstanceHsmDestinationPathServerCasOptions : Instantiate ListInstanceHsmDestinationPathServerCasOptions
func (*SecretsManagerInstanceManagementV2) NewListInstanceHsmDestinationPathServerCasOptions(instanceID string, destinationID string, pathID string) *ListInstanceHsmDestinationPathServerCasOptions {
	return &ListInstanceHsmDestinationPathServerCasOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListInstanceHsmDestinationPathServerCasOptions) SetInstanceID(instanceID string) *ListInstanceHsmDestinationPathServerCasOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *ListInstanceHsmDestinationPathServerCasOptions) SetDestinationID(destinationID string) *ListInstanceHsmDestinationPathServerCasOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *ListInstanceHsmDestinationPathServerCasOptions) SetPathID(pathID string) *ListInstanceHsmDestinationPathServerCasOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListInstanceHsmDestinationPathServerCasOptions) SetHeaders(param map[string]string) *ListInstanceHsmDestinationPathServerCasOptions {
	options.Headers = param
	return options
}

// ListInstanceHsmDestinationPathsOptions : The ListInstanceHsmDestinationPaths options.
type ListInstanceHsmDestinationPathsOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListInstanceHsmDestinationPathsOptions : Instantiate ListInstanceHsmDestinationPathsOptions
func (*SecretsManagerInstanceManagementV2) NewListInstanceHsmDestinationPathsOptions(instanceID string, destinationID string) *ListInstanceHsmDestinationPathsOptions {
	return &ListInstanceHsmDestinationPathsOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListInstanceHsmDestinationPathsOptions) SetInstanceID(instanceID string) *ListInstanceHsmDestinationPathsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *ListInstanceHsmDestinationPathsOptions) SetDestinationID(destinationID string) *ListInstanceHsmDestinationPathsOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListInstanceHsmDestinationPathsOptions) SetHeaders(param map[string]string) *ListInstanceHsmDestinationPathsOptions {
	options.Headers = param
	return options
}

// ListInstanceHsmDestinationsOptions : The ListInstanceHsmDestinations options.
type ListInstanceHsmDestinationsOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListInstanceHsmDestinationsOptions : Instantiate ListInstanceHsmDestinationsOptions
func (*SecretsManagerInstanceManagementV2) NewListInstanceHsmDestinationsOptions(instanceID string) *ListInstanceHsmDestinationsOptions {
	return &ListInstanceHsmDestinationsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListInstanceHsmDestinationsOptions) SetInstanceID(instanceID string) *ListInstanceHsmDestinationsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListInstanceHsmDestinationsOptions) SetHeaders(param map[string]string) *ListInstanceHsmDestinationsOptions {
	options.Headers = param
	return options
}

// Path : A path resource. crn is a top-level field — no vendor_config wrapper. Paths are independent of HSM HA Groups. crn is
// immutable after creation.
type Path struct {
	// Path ID. Immutable. System-managed.
	ID *string `json:"id,omitempty"`

	// URL of the path resource.
	Href *string `json:"href,omitempty"`

	// Human-readable display name.
	Name *string `json:"name" validate:"required"`

	// Optional description. Omitted if not set.
	Description *string `json:"description,omitempty"`

	// CRN of the gateway through which the HSM is reached. Immutable.
	Crn *string `json:"crn" validate:"required"`

	// Lifecycle state of the path:
	// - `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.
	// - `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy
	// for the gateway before proceeding.
	// - `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.
	// - `pending_approval`: Connectivity created. Waiting for customer approval.
	// - `completed`: Connectivity established. Path fully operational.
	// - `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.
	// - `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager
	// instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.
	// - `approval_declined`: Customer declined the pending approval. Call POST /reset.
	// - `deleting`: Delete initiated; async cleanup in progress.
	// - `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`

	// AS OF health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.
	// - `evaluating`: A health check is in progress for this path.
	// - `in_progress`: An apply operation is running across the destination.
	// - `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.
	// - `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval,
	// approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is
	// missing or expired, no_valid_server_ca takes priority.
	// - `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.
	// - `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.
	HealthState *string `json:"health_state" validate:"required"`

	// Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
	HealthAsOf *strfmt.DateTime `json:"health_as_of" validate:"required"`

	// Computed field.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// Computed field.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Constants associated with the Path.ProvisioningState property.
// Lifecycle state of the path:
// - `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.
// - `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy for
// the gateway before proceeding.
// - `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.
// - `pending_approval`: Connectivity created. Waiting for customer approval.
// - `completed`: Connectivity established. Path fully operational.
// - `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.
// - `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager
// instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.
// - `approval_declined`: Customer declined the pending approval. Call POST /reset.
// - `deleting`: Delete initiated; async cleanup in progress.
// - `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.
const (
	Path_ProvisioningState_ApprovalDeclined = "approval_declined"
	Path_ProvisioningState_Completed = "completed"
	Path_ProvisioningState_CrnAlreadyInUse = "crn_already_in_use"
	Path_ProvisioningState_Deleting = "deleting"
	Path_ProvisioningState_DeletionFailed = "deletion_failed"
	Path_ProvisioningState_PendingAccountPolicy = "pending_account_policy"
	Path_ProvisioningState_PendingApproval = "pending_approval"
	Path_ProvisioningState_PendingServerCa = "pending_server_ca"
	Path_ProvisioningState_Provisioning = "provisioning"
	Path_ProvisioningState_VpeProvisioningFailed = "vpe_provisioning_failed"
)

// Constants associated with the Path.HealthState property.
// AS OF health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.
// - `evaluating`: A health check is in progress for this path.
// - `in_progress`: An apply operation is running across the destination.
// - `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.
// - `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval,
// approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is
// missing or expired, no_valid_server_ca takes priority.
// - `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.
// - `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.
const (
	Path_HealthState_Evaluating = "evaluating"
	Path_HealthState_Healthy = "healthy"
	Path_HealthState_InProgress = "in_progress"
	Path_HealthState_NoValidServerCa = "no_valid_server_ca"
	Path_HealthState_NotConfigured = "not_configured"
	Path_HealthState_Unhealthy = "unhealthy"
	Path_HealthState_Unknown = "unknown"
)

// UnmarshalPath unmarshals an instance of Path from the specified map of raw messages.
func UnmarshalPath(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Path)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		err = core.SDKErrorf(err, "", "crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_as_of", &obj.HealthAsOf)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_as_of-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_by-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PathCollection : A collection of paths.
type PathCollection struct {
	// Paths for this destination. At most 20 paths may exist per destination.
	Paths []Path `json:"paths" validate:"required"`
}

// UnmarshalPathCollection unmarshals an instance of PathCollection from the specified map of raw messages.
func UnmarshalPathCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PathCollection)
	err = core.UnmarshalModel(m, "paths", &obj.Paths, UnmarshalPath)
	if err != nil {
		err = core.SDKErrorf(err, "", "paths-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PathPatch : Fields to update on the path. Only supplied fields are changed (merge-patch semantics).
type PathPatch struct {
	// Updated display name for the path.
	Name *string `json:"name,omitempty"`

	// Updated description for the path. Set to empty string to clear.
	Description *string `json:"description,omitempty"`
}

// UnmarshalPathPatch unmarshals an instance of PathPatch from the specified map of raw messages.
func UnmarshalPathPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PathPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the PathPatch
func (pathPatch *PathPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(pathPatch.Name) {
		_patch["name"] = pathPatch.Name
	}
	if !core.IsNil(pathPatch.Description) {
		_patch["description"] = pathPatch.Description
	}

	return
}

// ResetHsmDestinationPathOptions : The ResetHsmDestinationPath options.
type ResetHsmDestinationPathOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewResetHsmDestinationPathOptions : Instantiate ResetHsmDestinationPathOptions
func (*SecretsManagerInstanceManagementV2) NewResetHsmDestinationPathOptions(instanceID string, destinationID string, pathID string) *ResetHsmDestinationPathOptions {
	return &ResetHsmDestinationPathOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ResetHsmDestinationPathOptions) SetInstanceID(instanceID string) *ResetHsmDestinationPathOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *ResetHsmDestinationPathOptions) SetDestinationID(destinationID string) *ResetHsmDestinationPathOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *ResetHsmDestinationPathOptions) SetPathID(pathID string) *ResetHsmDestinationPathOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResetHsmDestinationPathOptions) SetHeaders(param map[string]string) *ResetHsmDestinationPathOptions {
	options.Headers = param
	return options
}

// RunHsmDestinationHealthCheckOptions : The RunHsmDestinationHealthCheck options.
type RunHsmDestinationHealthCheckOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRunHsmDestinationHealthCheckOptions : Instantiate RunHsmDestinationHealthCheckOptions
func (*SecretsManagerInstanceManagementV2) NewRunHsmDestinationHealthCheckOptions(instanceID string, destinationID string) *RunHsmDestinationHealthCheckOptions {
	return &RunHsmDestinationHealthCheckOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *RunHsmDestinationHealthCheckOptions) SetInstanceID(instanceID string) *RunHsmDestinationHealthCheckOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *RunHsmDestinationHealthCheckOptions) SetDestinationID(destinationID string) *RunHsmDestinationHealthCheckOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunHsmDestinationHealthCheckOptions) SetHeaders(param map[string]string) *RunHsmDestinationHealthCheckOptions {
	options.Headers = param
	return options
}

// ServerCa : A Server CA certificate resource. certificate and expires_at are top-level fields. There is no vendor_config wrapper.
// Only one Server CA can be active per path. Constraint: Cannot leave a path with zero active Server CAs if a client
// certificate is already active.
type ServerCa struct {
	// Server CA ID. Immutable.
	ID *string `json:"id,omitempty"`

	// Human-readable display name.
	Name *string `json:"name" validate:"required"`

	// Optional description. Omitted if not set.
	Description *string `json:"description,omitempty"`

	// PEM-encoded Server CA certificate. Immutable.
	Certificate *string `json:"certificate,omitempty"`

	// Expiry parsed from certificate NotAfter field. System-managed. Populated on creation.
	ExpiresAt *strfmt.DateTime `json:"expires_at" validate:"required"`

	// Only one Server CA can be active per path. Starts inactive on creation.
	Status *string `json:"status" validate:"required"`

	// AS OF health state of the Server CA certificate. Written exclusively by POST …/health:
	// - `unknown`: Initial state on creation. Health check not yet performed.
	// - `evaluating`: A health check is in progress for this Server CA.
	// - `valid`: Certificate is within its validity period (expires_at is in the future).
	// - `expired`: Certificate has passed its expires_at date.
	HealthState *string `json:"health_state" validate:"required"`

	// Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.
	HealthAsOf *strfmt.DateTime `json:"health_as_of,omitempty"`

	// Computed field.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// Computed field.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Computed field. Omitted until first update.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Constants associated with the ServerCa.Status property.
// Only one Server CA can be active per path. Starts inactive on creation.
const (
	ServerCa_Status_Active = "active"
	ServerCa_Status_Inactive = "inactive"
)

// Constants associated with the ServerCa.HealthState property.
// AS OF health state of the Server CA certificate. Written exclusively by POST …/health:
// - `unknown`: Initial state on creation. Health check not yet performed.
// - `evaluating`: A health check is in progress for this Server CA.
// - `valid`: Certificate is within its validity period (expires_at is in the future).
// - `expired`: Certificate has passed its expires_at date.
const (
	ServerCa_HealthState_Evaluating = "evaluating"
	ServerCa_HealthState_Expired = "expired"
	ServerCa_HealthState_Unknown = "unknown"
	ServerCa_HealthState_Valid = "valid"
)

// UnmarshalServerCa unmarshals an instance of ServerCa from the specified map of raw messages.
func UnmarshalServerCa(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServerCa)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate", &obj.Certificate)
	if err != nil {
		err = core.SDKErrorf(err, "", "certificate-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "expires_at", &obj.ExpiresAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "expires_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_as_of", &obj.HealthAsOf)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_as_of-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_by-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServerCaCollection : A collection of Server CA certificates.
type ServerCaCollection struct {
	// The list of Server CA certificates for this path.
	ServerCas []ServerCa `json:"server_cas" validate:"required"`
}

// UnmarshalServerCaCollection unmarshals an instance of ServerCaCollection from the specified map of raw messages.
func UnmarshalServerCaCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServerCaCollection)
	err = core.UnmarshalModel(m, "server_cas", &obj.ServerCas, UnmarshalServerCa)
	if err != nil {
		err = core.SDKErrorf(err, "", "server_cas-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServerCaPatch : Fields to update on the Server CA. Only supplied fields are changed (merge-patch semantics).
type ServerCaPatch struct {
	// Updated display name for the Server CA.
	Name *string `json:"name,omitempty"`

	// Updated description. Set to empty string to clear.
	Description *string `json:"description,omitempty"`

	// Set to `active` to activate this certificate. Only `active` is allowed.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the ServerCaPatch.Status property.
// Set to `active` to activate this certificate. Only `active` is allowed.
const (
	ServerCaPatch_Status_Active = "active"
)

// UnmarshalServerCaPatch unmarshals an instance of ServerCaPatch from the specified map of raw messages.
func UnmarshalServerCaPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServerCaPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the ServerCaPatch
func (serverCaPatch *ServerCaPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(serverCaPatch.Name) {
		_patch["name"] = serverCaPatch.Name
	}
	if !core.IsNil(serverCaPatch.Description) {
		_patch["description"] = serverCaPatch.Description
	}
	if !core.IsNil(serverCaPatch.Status) {
		_patch["status"] = serverCaPatch.Status
	}

	return
}

// Token : Admin token response. Exactly one of token or wrapped_token is present, never both. wrapped_token is returned only
// when response_wrapping: true is requested.
// Models which "extend" this model:
// - TokenPlainAdminToken
// - TokenWrappedAdminToken
type Token struct {
	// The plain Vault admin token.
	Token *string `json:"token,omitempty"`

	// A Vault response-wrapped token. Present only when response_wrapping: true is requested.
	WrappedToken *string `json:"wrapped_token,omitempty"`
}
func (*Token) isaToken() bool {
	return true
}

type TokenIntf interface {
	isaToken() bool
}

// UnmarshalToken unmarshals an instance of Token from the specified map of raw messages.
func UnmarshalToken(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Token)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		err = core.SDKErrorf(err, "", "token-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "wrapped_token", &obj.WrappedToken)
	if err != nil {
		err = core.SDKErrorf(err, "", "wrapped_token-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateInstanceDestinationOptions : The UpdateInstanceDestination options.
type UpdateInstanceDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// JSON Merge-Patch content for update_instance_destination.
	DestinationPatch map[string]interface{} `json:"Destination_patch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateInstanceDestinationOptions : Instantiate UpdateInstanceDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewUpdateInstanceDestinationOptions(instanceID string, destinationID string, destinationPatch map[string]interface{}) *UpdateInstanceDestinationOptions {
	return &UpdateInstanceDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		DestinationPatch: destinationPatch,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceDestinationOptions) SetInstanceID(instanceID string) *UpdateInstanceDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *UpdateInstanceDestinationOptions) SetDestinationID(destinationID string) *UpdateInstanceDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetDestinationPatch : Allow user to set DestinationPatch
func (_options *UpdateInstanceDestinationOptions) SetDestinationPatch(destinationPatch map[string]interface{}) *UpdateInstanceDestinationOptions {
	_options.DestinationPatch = destinationPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceDestinationOptions) SetHeaders(param map[string]string) *UpdateInstanceDestinationOptions {
	options.Headers = param
	return options
}

// UpdateInstanceHsmDestinationClientCertificateOptions : The UpdateInstanceHsmDestinationClientCertificate options.
type UpdateInstanceHsmDestinationClientCertificateOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// The identifier of the Client Certificate to operate on. Matches the `id` field in the `ClientCertificate` response
	// schema.
	CertID *string `json:"cert_id" validate:"required,ne="`

	// JSON Merge-Patch content for update_instance_hsm_destination_client_certificate.
	ClientCertificatePatch map[string]interface{} `json:"clientCertificatePatch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateInstanceHsmDestinationClientCertificateOptions : Instantiate UpdateInstanceHsmDestinationClientCertificateOptions
func (*SecretsManagerInstanceManagementV2) NewUpdateInstanceHsmDestinationClientCertificateOptions(instanceID string, destinationID string, certID string, clientCertificatePatch map[string]interface{}) *UpdateInstanceHsmDestinationClientCertificateOptions {
	return &UpdateInstanceHsmDestinationClientCertificateOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		CertID: core.StringPtr(certID),
		ClientCertificatePatch: clientCertificatePatch,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceHsmDestinationClientCertificateOptions) SetInstanceID(instanceID string) *UpdateInstanceHsmDestinationClientCertificateOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *UpdateInstanceHsmDestinationClientCertificateOptions) SetDestinationID(destinationID string) *UpdateInstanceHsmDestinationClientCertificateOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetCertID : Allow user to set CertID
func (_options *UpdateInstanceHsmDestinationClientCertificateOptions) SetCertID(certID string) *UpdateInstanceHsmDestinationClientCertificateOptions {
	_options.CertID = core.StringPtr(certID)
	return _options
}

// SetClientCertificatePatch : Allow user to set ClientCertificatePatch
func (_options *UpdateInstanceHsmDestinationClientCertificateOptions) SetClientCertificatePatch(clientCertificatePatch map[string]interface{}) *UpdateInstanceHsmDestinationClientCertificateOptions {
	_options.ClientCertificatePatch = clientCertificatePatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceHsmDestinationClientCertificateOptions) SetHeaders(param map[string]string) *UpdateInstanceHsmDestinationClientCertificateOptions {
	options.Headers = param
	return options
}

// UpdateInstanceHsmDestinationHaGroupOptions : The UpdateInstanceHsmDestinationHaGroup options.
type UpdateInstanceHsmDestinationHaGroupOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// The identifier of the HA Group to operate on. Matches the `id` field in the `HSMHAGroup` response schema.
	HaGroupID *string `json:"ha_group_id" validate:"required,ne="`

	// JSON Merge-Patch content for update_instance_hsm_destination_ha_group.
	HSMHAGroupPatch map[string]interface{} `json:"HSMHAGroupPatch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateInstanceHsmDestinationHaGroupOptions : Instantiate UpdateInstanceHsmDestinationHaGroupOptions
func (*SecretsManagerInstanceManagementV2) NewUpdateInstanceHsmDestinationHaGroupOptions(instanceID string, destinationID string, haGroupID string, hSMHAGroupPatch map[string]interface{}) *UpdateInstanceHsmDestinationHaGroupOptions {
	return &UpdateInstanceHsmDestinationHaGroupOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		HaGroupID: core.StringPtr(haGroupID),
		HSMHAGroupPatch: hSMHAGroupPatch,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceHsmDestinationHaGroupOptions) SetInstanceID(instanceID string) *UpdateInstanceHsmDestinationHaGroupOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *UpdateInstanceHsmDestinationHaGroupOptions) SetDestinationID(destinationID string) *UpdateInstanceHsmDestinationHaGroupOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHaGroupID : Allow user to set HaGroupID
func (_options *UpdateInstanceHsmDestinationHaGroupOptions) SetHaGroupID(haGroupID string) *UpdateInstanceHsmDestinationHaGroupOptions {
	_options.HaGroupID = core.StringPtr(haGroupID)
	return _options
}

// SetHSMHAGroupPatch : Allow user to set HSMHAGroupPatch
func (_options *UpdateInstanceHsmDestinationHaGroupOptions) SetHSMHAGroupPatch(hSMHAGroupPatch map[string]interface{}) *UpdateInstanceHsmDestinationHaGroupOptions {
	_options.HSMHAGroupPatch = hSMHAGroupPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceHsmDestinationHaGroupOptions) SetHeaders(param map[string]string) *UpdateInstanceHsmDestinationHaGroupOptions {
	options.Headers = param
	return options
}

// UpdateInstanceHsmDestinationOptions : The UpdateInstanceHsmDestination options.
type UpdateInstanceHsmDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// JSON Merge-Patch content for update_instance_hsm_destination.
	HsmDestinationPatch map[string]interface{} `json:"hsmDestinationPatch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateInstanceHsmDestinationOptions : Instantiate UpdateInstanceHsmDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewUpdateInstanceHsmDestinationOptions(instanceID string, destinationID string, hsmDestinationPatch map[string]interface{}) *UpdateInstanceHsmDestinationOptions {
	return &UpdateInstanceHsmDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		HsmDestinationPatch: hsmDestinationPatch,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceHsmDestinationOptions) SetInstanceID(instanceID string) *UpdateInstanceHsmDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *UpdateInstanceHsmDestinationOptions) SetDestinationID(destinationID string) *UpdateInstanceHsmDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHsmDestinationPatch : Allow user to set HsmDestinationPatch
func (_options *UpdateInstanceHsmDestinationOptions) SetHsmDestinationPatch(hsmDestinationPatch map[string]interface{}) *UpdateInstanceHsmDestinationOptions {
	_options.HsmDestinationPatch = hsmDestinationPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceHsmDestinationOptions) SetHeaders(param map[string]string) *UpdateInstanceHsmDestinationOptions {
	options.Headers = param
	return options
}

// UpdateInstanceHsmDestinationPathOptions : The UpdateInstanceHsmDestinationPath options.
type UpdateInstanceHsmDestinationPathOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// JSON Merge-Patch content for update_instance_hsm_destination_path.
	PathPatch map[string]interface{} `json:"pathPatch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateInstanceHsmDestinationPathOptions : Instantiate UpdateInstanceHsmDestinationPathOptions
func (*SecretsManagerInstanceManagementV2) NewUpdateInstanceHsmDestinationPathOptions(instanceID string, destinationID string, pathID string, pathPatch map[string]interface{}) *UpdateInstanceHsmDestinationPathOptions {
	return &UpdateInstanceHsmDestinationPathOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
		PathPatch: pathPatch,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceHsmDestinationPathOptions) SetInstanceID(instanceID string) *UpdateInstanceHsmDestinationPathOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *UpdateInstanceHsmDestinationPathOptions) SetDestinationID(destinationID string) *UpdateInstanceHsmDestinationPathOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *UpdateInstanceHsmDestinationPathOptions) SetPathID(pathID string) *UpdateInstanceHsmDestinationPathOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetPathPatch : Allow user to set PathPatch
func (_options *UpdateInstanceHsmDestinationPathOptions) SetPathPatch(pathPatch map[string]interface{}) *UpdateInstanceHsmDestinationPathOptions {
	_options.PathPatch = pathPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceHsmDestinationPathOptions) SetHeaders(param map[string]string) *UpdateInstanceHsmDestinationPathOptions {
	options.Headers = param
	return options
}

// UpdateInstanceHsmDestinationPathServerCaOptions : The UpdateInstanceHsmDestinationPathServerCa options.
type UpdateInstanceHsmDestinationPathServerCaOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination ID.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Path ID.
	PathID *string `json:"path_id" validate:"required,ne="`

	// The identifier of the Server CA to operate on. Matches the `id` field in the `ServerCa` response schema.
	ServerCaID *string `json:"server_ca_id" validate:"required,ne="`

	// JSON Merge-Patch content for update_instance_hsm_destination_path_server_ca.
	ServerCaPatch map[string]interface{} `json:"serverCaPatch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateInstanceHsmDestinationPathServerCaOptions : Instantiate UpdateInstanceHsmDestinationPathServerCaOptions
func (*SecretsManagerInstanceManagementV2) NewUpdateInstanceHsmDestinationPathServerCaOptions(instanceID string, destinationID string, pathID string, serverCaID string, serverCaPatch map[string]interface{}) *UpdateInstanceHsmDestinationPathServerCaOptions {
	return &UpdateInstanceHsmDestinationPathServerCaOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		PathID: core.StringPtr(pathID),
		ServerCaID: core.StringPtr(serverCaID),
		ServerCaPatch: serverCaPatch,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceHsmDestinationPathServerCaOptions) SetInstanceID(instanceID string) *UpdateInstanceHsmDestinationPathServerCaOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *UpdateInstanceHsmDestinationPathServerCaOptions) SetDestinationID(destinationID string) *UpdateInstanceHsmDestinationPathServerCaOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetPathID : Allow user to set PathID
func (_options *UpdateInstanceHsmDestinationPathServerCaOptions) SetPathID(pathID string) *UpdateInstanceHsmDestinationPathServerCaOptions {
	_options.PathID = core.StringPtr(pathID)
	return _options
}

// SetServerCaID : Allow user to set ServerCaID
func (_options *UpdateInstanceHsmDestinationPathServerCaOptions) SetServerCaID(serverCaID string) *UpdateInstanceHsmDestinationPathServerCaOptions {
	_options.ServerCaID = core.StringPtr(serverCaID)
	return _options
}

// SetServerCaPatch : Allow user to set ServerCaPatch
func (_options *UpdateInstanceHsmDestinationPathServerCaOptions) SetServerCaPatch(serverCaPatch map[string]interface{}) *UpdateInstanceHsmDestinationPathServerCaOptions {
	_options.ServerCaPatch = serverCaPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceHsmDestinationPathServerCaOptions) SetHeaders(param map[string]string) *UpdateInstanceHsmDestinationPathServerCaOptions {
	options.Headers = param
	return options
}

// VaultDedicatedCluster : Vault cluster information for Vault Dedicated instances.
type VaultDedicatedCluster struct {
	// Vault cluster status. Possible values:
	// - sealed: The Vault cluster is sealed and requires unsealing to access secrets
	// - not_initialized: The Vault cluster has not been initialized yet
	// - healthy: The Vault cluster is operational and ready to serve requests.
	Status *string `json:"status" validate:"required"`

	// Vault cluster version.
	Version *string `json:"version" validate:"required"`
}

// Constants associated with the VaultDedicatedCluster.Status property.
// Vault cluster status. Possible values:
// - sealed: The Vault cluster is sealed and requires unsealing to access secrets
// - not_initialized: The Vault cluster has not been initialized yet
// - healthy: The Vault cluster is operational and ready to serve requests.
const (
	VaultDedicatedCluster_Status_Healthy = "healthy"
	VaultDedicatedCluster_Status_NotInitialized = "not_initialized"
	VaultDedicatedCluster_Status_Sealed = "sealed"
)

// UnmarshalVaultDedicatedCluster unmarshals an instance of VaultDedicatedCluster from the specified map of raw messages.
func UnmarshalVaultDedicatedCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultDedicatedCluster)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VaultDedicatedEndpointsData : Endpoint URLs for accessing the Vault Dedicated instance.
type VaultDedicatedEndpointsData struct {
	// Vault API endpoint URL.
	VaultApi *string `json:"vault_api" validate:"required"`

	// Vault UI endpoint URL.
	VaultUi *string `json:"vault_ui" validate:"required"`
}

// UnmarshalVaultDedicatedEndpointsData unmarshals an instance of VaultDedicatedEndpointsData from the specified map of raw messages.
func UnmarshalVaultDedicatedEndpointsData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultDedicatedEndpointsData)
	err = core.UnmarshalPrimitive(m, "vault_api", &obj.VaultApi)
	if err != nil {
		err = core.SDKErrorf(err, "", "vault_api-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "vault_ui", &obj.VaultUi)
	if err != nil {
		err = core.SDKErrorf(err, "", "vault_ui-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VaultDedicatedInstanceEncryption : Vault encryption configuration for Vault Dedicated instances.
type VaultDedicatedInstanceEncryption struct {
	// Vault encryption mode.
	Mode *string `json:"mode" validate:"required"`

	// Vault encryption provider (only present for customer_managed mode). Valid value - 'key_protect'.
	Provider *string `json:"provider,omitempty"`

	// Vault encryption key CRN (only present for customer_managed mode).
	KeyCrn *string `json:"key_crn,omitempty"`
}

// Constants associated with the VaultDedicatedInstanceEncryption.Mode property.
// Vault encryption mode.
const (
	VaultDedicatedInstanceEncryption_Mode_CustomerManaged = "customer_managed"
	VaultDedicatedInstanceEncryption_Mode_ServiceManaged = "service_managed"
)

// UnmarshalVaultDedicatedInstanceEncryption unmarshals an instance of VaultDedicatedInstanceEncryption from the specified map of raw messages.
func UnmarshalVaultDedicatedInstanceEncryption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultDedicatedInstanceEncryption)
	err = core.UnmarshalPrimitive(m, "mode", &obj.Mode)
	if err != nil {
		err = core.SDKErrorf(err, "", "mode-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		err = core.SDKErrorf(err, "", "provider-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "key_crn", &obj.KeyCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "key_crn-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VaultDedicatedInstanceEndpoints : Instance endpoints for Vault Dedicated instances.
type VaultDedicatedInstanceEndpoints struct {
	// Endpoint URLs for accessing the Vault Dedicated instance.
	Public *VaultDedicatedEndpointsData `json:"public,omitempty"`

	// Endpoint URLs for accessing the Vault Dedicated instance.
	Private *VaultDedicatedEndpointsData `json:"private" validate:"required"`
}

// UnmarshalVaultDedicatedInstanceEndpoints unmarshals an instance of VaultDedicatedInstanceEndpoints from the specified map of raw messages.
func UnmarshalVaultDedicatedInstanceEndpoints(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultDedicatedInstanceEndpoints)
	err = core.UnmarshalModel(m, "public", &obj.Public, UnmarshalVaultDedicatedEndpointsData)
	if err != nil {
		err = core.SDKErrorf(err, "", "public-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "private", &obj.Private, UnmarshalVaultDedicatedEndpointsData)
	if err != nil {
		err = core.SDKErrorf(err, "", "private-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VpeProvisioningResponse : Response body for POST .../paths/{path_id}/build (202 Accepted).
type VpeProvisioningResponse struct {
	// ID of the path being provisioned.
	ID *string `json:"id,omitempty"`

	// Always `provisioning` on a 202 response.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`
}

// Constants associated with the VpeProvisioningResponse.ProvisioningState property.
// Always `provisioning` on a 202 response.
const (
	VpeProvisioningResponse_ProvisioningState_Provisioning = "provisioning"
)

// UnmarshalVpeProvisioningResponse unmarshals an instance of VpeProvisioningResponse from the specified map of raw messages.
func UnmarshalVpeProvisioningResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VpeProvisioningResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VpeStatusResult : Response body for POST .../paths/{path_id}/health_check. Reflects the current health state and provisioning state of
// the HSM path.
type VpeStatusResult struct {
	// AS OF health state. Written exclusively by POST …/health and POST …/apply:
	// - `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.
	// - `evaluating`: A health check is in progress for this path.
	// - `in_progress`: An apply operation is running across the destination.
	// - `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.
	// - `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval,
	// approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is
	// missing or expired, no_valid_server_ca takes priority.
	// - `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.
	// - `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.
	HealthState *string `json:"health_state" validate:"required"`

	// RFC 3339 timestamp of the last health evaluation.
	HealthTimeStamp *strfmt.DateTime `json:"health_time_stamp" validate:"required"`

	// Lifecycle state of the path:
	// - `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.
	// - `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy
	// for the gateway before proceeding.
	// - `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.
	// - `pending_approval`: Connectivity created. Waiting for customer approval.
	// - `completed`: Connectivity established. Path fully operational.
	// - `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.
	// - `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager
	// instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.
	// - `approval_declined`: Customer declined the pending approval. Call POST /reset.
	// - `deleting`: Delete initiated; async cleanup in progress.
	// - `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.
	ProvisioningState *string `json:"provisioning_state" validate:"required"`
}

// Constants associated with the VpeStatusResult.HealthState property.
// AS OF health state. Written exclusively by POST …/health and POST …/apply:
// - `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.
// - `evaluating`: A health check is in progress for this path.
// - `in_progress`: An apply operation is running across the destination.
// - `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.
// - `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval,
// approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is
// missing or expired, no_valid_server_ca takes priority.
// - `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.
// - `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.
const (
	VpeStatusResult_HealthState_Evaluating = "evaluating"
	VpeStatusResult_HealthState_Healthy = "healthy"
	VpeStatusResult_HealthState_InProgress = "in_progress"
	VpeStatusResult_HealthState_NoValidServerCa = "no_valid_server_ca"
	VpeStatusResult_HealthState_NotConfigured = "not_configured"
	VpeStatusResult_HealthState_Unhealthy = "unhealthy"
	VpeStatusResult_HealthState_Unknown = "unknown"
)

// Constants associated with the VpeStatusResult.ProvisioningState property.
// Lifecycle state of the path:
// - `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.
// - `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy for
// the gateway before proceeding.
// - `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.
// - `pending_approval`: Connectivity created. Waiting for customer approval.
// - `completed`: Connectivity established. Path fully operational.
// - `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.
// - `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager
// instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.
// - `approval_declined`: Customer declined the pending approval. Call POST /reset.
// - `deleting`: Delete initiated; async cleanup in progress.
// - `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.
const (
	VpeStatusResult_ProvisioningState_ApprovalDeclined = "approval_declined"
	VpeStatusResult_ProvisioningState_Completed = "completed"
	VpeStatusResult_ProvisioningState_CrnAlreadyInUse = "crn_already_in_use"
	VpeStatusResult_ProvisioningState_Deleting = "deleting"
	VpeStatusResult_ProvisioningState_DeletionFailed = "deletion_failed"
	VpeStatusResult_ProvisioningState_PendingAccountPolicy = "pending_account_policy"
	VpeStatusResult_ProvisioningState_PendingApproval = "pending_approval"
	VpeStatusResult_ProvisioningState_PendingServerCa = "pending_server_ca"
	VpeStatusResult_ProvisioningState_Provisioning = "provisioning"
	VpeStatusResult_ProvisioningState_VpeProvisioningFailed = "vpe_provisioning_failed"
)

// UnmarshalVpeStatusResult unmarshals an instance of VpeStatusResult from the specified map of raw messages.
func UnmarshalVpeStatusResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VpeStatusResult)
	err = core.UnmarshalPrimitive(m, "health_state", &obj.HealthState)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "health_time_stamp", &obj.HealthTimeStamp)
	if err != nil {
		err = core.SDKErrorf(err, "", "health_time_stamp-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "provisioning_state", &obj.ProvisioningState)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateDestinationRequestIbmCloudDatabaseDestinationPrototype : Request body for creating an IBM Cloud Database destination.
// This model "extends" CreateDestinationRequest
type CreateDestinationRequestIbmCloudDatabaseDestinationPrototype struct {
	// Destination name.
	Name *string `json:"name" validate:"required"`

	// Destination type.
	Type *string `json:"type" validate:"required"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// IBM Cloud Database service instance CRN.
	Crn *string `json:"crn" validate:"required"`
}

// Constants associated with the CreateDestinationRequestIbmCloudDatabaseDestinationPrototype.Type property.
// Destination type.
const (
	CreateDestinationRequestIbmCloudDatabaseDestinationPrototype_Type_IbmCloudDatabase = "ibm_cloud_database"
)

// NewCreateDestinationRequestIbmCloudDatabaseDestinationPrototype : Instantiate CreateDestinationRequestIbmCloudDatabaseDestinationPrototype (Generic Model Constructor)
func (*SecretsManagerInstanceManagementV2) NewCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(name string, typeVar string, crn string) (_model *CreateDestinationRequestIbmCloudDatabaseDestinationPrototype, err error) {
	_model = &CreateDestinationRequestIbmCloudDatabaseDestinationPrototype{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
		Crn: core.StringPtr(crn),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

func (*CreateDestinationRequestIbmCloudDatabaseDestinationPrototype) isaCreateDestinationRequest() bool {
	return true
}

// UnmarshalCreateDestinationRequestIbmCloudDatabaseDestinationPrototype unmarshals an instance of CreateDestinationRequestIbmCloudDatabaseDestinationPrototype from the specified map of raw messages.
func UnmarshalCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		err = core.SDKErrorf(err, "", "crn-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TokenPlainAdminToken : Response containing a plain Vault admin token.
// This model "extends" Token
type TokenPlainAdminToken struct {
	// The plain Vault admin token.
	Token *string `json:"token" validate:"required"`
}

func (*TokenPlainAdminToken) isaToken() bool {
	return true
}

// UnmarshalTokenPlainAdminToken unmarshals an instance of TokenPlainAdminToken from the specified map of raw messages.
func UnmarshalTokenPlainAdminToken(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenPlainAdminToken)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		err = core.SDKErrorf(err, "", "token-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TokenWrappedAdminToken : Response containing a Vault response-wrapped token.
// This model "extends" Token
type TokenWrappedAdminToken struct {
	// A Vault response-wrapped token. Present only when response_wrapping: true is requested.
	WrappedToken *string `json:"wrapped_token" validate:"required"`
}

func (*TokenWrappedAdminToken) isaToken() bool {
	return true
}

// UnmarshalTokenWrappedAdminToken unmarshals an instance of TokenWrappedAdminToken from the specified map of raw messages.
func UnmarshalTokenWrappedAdminToken(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenWrappedAdminToken)
	err = core.UnmarshalPrimitive(m, "wrapped_token", &obj.WrappedToken)
	if err != nil {
		err = core.SDKErrorf(err, "", "wrapped_token-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
