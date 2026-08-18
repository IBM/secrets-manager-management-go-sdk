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
 * IBM OpenAPI SDK Code Generator Version: 3.116.0-df613dbc-20260803-154903
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
// - Revoke all active Vault admin tokens.
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
// grants administrative privileges. Use only for initial setup and cluster management, then revoke immediately.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateVaultAdmintoken(createVaultAdmintokenOptions *CreateVaultAdmintokenOptions) (result *Token, response *core.DetailedResponse, err error) {
	result, response, err = secretsManagerInstanceManagement.CreateVaultAdmintokenWithContext(context.Background(), createVaultAdmintokenOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateVaultAdmintokenWithContext is an alternate form of the CreateVaultAdmintoken method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateVaultAdmintokenWithContext(ctx context.Context, createVaultAdmintokenOptions *CreateVaultAdmintokenOptions) (result *Token, response *core.DetailedResponse, err error) {
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
		"id": *createVaultAdmintokenOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{id}/admintokens`, pathParamsMap)
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

// DeleteInstanceAdmintokens : Delete admin tokens
// Revoke all active Vault admin tokens. This immediately invalidates all existing admin tokens.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptions *DeleteInstanceAdmintokensOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.DeleteInstanceAdmintokensWithContext(context.Background(), deleteInstanceAdmintokensOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceAdmintokensWithContext is an alternate form of the DeleteInstanceAdmintokens method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceAdmintokensWithContext(ctx context.Context, deleteInstanceAdmintokensOptions *DeleteInstanceAdmintokensOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteInstanceAdmintokensOptions, "deleteInstanceAdmintokensOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteInstanceAdmintokensOptions, "deleteInstanceAdmintokensOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteInstanceAdmintokensOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{id}/admintokens`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("secrets_manager_instance_management", "V2", "DeleteInstanceAdmintokens")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteInstanceAdmintokensOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_admintokens", getServiceComponentInfo())
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
		"id": *getInstanceOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/v2/instances/{id}`, pathParamsMap)
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
// Returns `202 Accepted` with `state: not_started`. Provisioning completes asynchronously — poll `GET
// /destinations/{id}` until `state` transitions to `succeeded` or `failed`.
//
// **Beta**: Only Gen 1 (Classic) IBM Cloud Database service instances are supported. Gen 2 instances are rejected with
// `422`. IBM Cloud Database service instances with no private endpoints are also rejected with `422`.
//
// **Rate Limit**: 10 requests per instance per minute
// **Quota**: Maximum 20 destinations per instance.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceDestination(createInstanceDestinationOptions *CreateInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.CreateInstanceDestinationWithContext(context.Background(), createInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateInstanceDestinationWithContext is an alternate form of the CreateInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) CreateInstanceDestinationWithContext(ctx context.Context, createInstanceDestinationOptions *CreateInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
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

	body := make(map[string]interface{})
	if createInstanceDestinationOptions.Name != nil {
		body["name"] = createInstanceDestinationOptions.Name
	}
	if createInstanceDestinationOptions.Type != nil {
		body["type"] = createInstanceDestinationOptions.Type
	}
	if createInstanceDestinationOptions.Description != nil {
		body["description"] = createInstanceDestinationOptions.Description
	}
	if createInstanceDestinationOptions.Crn != nil {
		body["crn"] = createInstanceDestinationOptions.Crn
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

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetInstanceDestination : Get destination details
// Retrieve details and current state for a specific destination for your Vault Dedicated cluster.
//
// Returns `404` if the destination does not exist. A deleted destination is immediately absent from GET — the
// `deleting` state is internal only and never returned to callers.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceDestination(getInstanceDestinationOptions *GetInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.GetInstanceDestinationWithContext(context.Background(), getInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInstanceDestinationWithContext is an alternate form of the GetInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) GetInstanceDestinationWithContext(ctx context.Context, getInstanceDestinationOptions *GetInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
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

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateInstanceDestination : Update destination
// Update mutable metadata fields (`name`, `description`) on a destination for your Vault Dedicated cluster. All other
// fields are immutable after creation.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceDestination(updateInstanceDestinationOptions *UpdateInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.UpdateInstanceDestinationWithContext(context.Background(), updateInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateInstanceDestinationWithContext is an alternate form of the UpdateInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) UpdateInstanceDestinationWithContext(ctx context.Context, updateInstanceDestinationOptions *UpdateInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
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

	_, err = builder.SetBodyContentJSON(updateInstanceDestinationOptions.RequestBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// DeleteInstanceDestination : Delete destination
// Delete a destination for your Vault Dedicated cluster. A deleted destination is immediately absent from GET after
// this call returns 204.
//
// A `failed` destination still counts against the per-instance quota until deleted.
//
// **Rate Limit**: 10 requests per instance per minute.
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceDestination(deleteInstanceDestinationOptions *DeleteInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
	response, err = secretsManagerInstanceManagement.DeleteInstanceDestinationWithContext(context.Background(), deleteInstanceDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteInstanceDestinationWithContext is an alternate form of the DeleteInstanceDestination method which supports a Context parameter
func (secretsManagerInstanceManagement *SecretsManagerInstanceManagementV2) DeleteInstanceDestinationWithContext(ctx context.Context, deleteInstanceDestinationOptions *DeleteInstanceDestinationOptions) (response *core.DetailedResponse, err error) {
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

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = secretsManagerInstanceManagement.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_instance_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "2.0.0")
}

// CreateInstanceDestinationOptions : The CreateInstanceDestination options.
type CreateInstanceDestinationOptions struct {
	// Secrets Manager instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Destination name.
	Name *string `json:"name,omitempty"`

	// Destination type.
	Type *string `json:"type,omitempty"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	// IBM Cloud Database service instance CRN.
	Crn *string `json:"crn,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateInstanceDestinationOptions.Type property.
// Destination type.
const (
	CreateInstanceDestinationOptions_Type_IbmCloudDatabase = "ibm_cloud_database"
)

// NewCreateInstanceDestinationOptions : Instantiate CreateInstanceDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewCreateInstanceDestinationOptions(instanceID string) *CreateInstanceDestinationOptions {
	return &CreateInstanceDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceDestinationOptions) SetInstanceID(instanceID string) *CreateInstanceDestinationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateInstanceDestinationOptions) SetName(name string) *CreateInstanceDestinationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateInstanceDestinationOptions) SetType(typeVar string) *CreateInstanceDestinationOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateInstanceDestinationOptions) SetDescription(description string) *CreateInstanceDestinationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetCrn : Allow user to set Crn
func (_options *CreateInstanceDestinationOptions) SetCrn(crn string) *CreateInstanceDestinationOptions {
	_options.Crn = core.StringPtr(crn)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceDestinationOptions) SetHeaders(param map[string]string) *CreateInstanceDestinationOptions {
	options.Headers = param
	return options
}

// CreateVaultAdmintokenOptions : The CreateVaultAdmintoken options.
type CreateVaultAdmintokenOptions struct {
	// Secrets Manager instance ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateVaultAdmintokenOptions : Instantiate CreateVaultAdmintokenOptions
func (*SecretsManagerInstanceManagementV2) NewCreateVaultAdmintokenOptions(id string) *CreateVaultAdmintokenOptions {
	return &CreateVaultAdmintokenOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *CreateVaultAdmintokenOptions) SetID(id string) *CreateVaultAdmintokenOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVaultAdmintokenOptions) SetHeaders(param map[string]string) *CreateVaultAdmintokenOptions {
	options.Headers = param
	return options
}

// DeleteInstanceAdmintokensOptions : The DeleteInstanceAdmintokens options.
type DeleteInstanceAdmintokensOptions struct {
	// Secrets Manager instance ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceAdmintokensOptions : Instantiate DeleteInstanceAdmintokensOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceAdmintokensOptions(id string) *DeleteInstanceAdmintokensOptions {
	return &DeleteInstanceAdmintokensOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteInstanceAdmintokensOptions) SetID(id string) *DeleteInstanceAdmintokensOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceAdmintokensOptions) SetHeaders(param map[string]string) *DeleteInstanceAdmintokensOptions {
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

// Destination : A destination resource representing a private network link to a service instance on a Vault Dedicated cluster.
type Destination struct {
	// Destination ID.
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
	// - `not_started`: Job accepted, waiting to start provisioning
	// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
	// - `succeeded`: Destination ready and usable
	// - `failed`: Provisioning failed — terminal state; delete and recreate.
	//   A `failed` destination still counts against the per-instance quota until deleted.
	State *string `json:"state" validate:"required"`

	// Timestamp when the destination was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Timestamp when the destination was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// IAM identity that created the destination.
	CreatedBy *string `json:"created_by,omitempty"`
}

// Constants associated with the Destination.Type property.
// Destination type.
const (
	Destination_Type_IbmCloudDatabase = "ibm_cloud_database"
)

// Constants associated with the Destination.State property.
// Destination state:
// - `not_started`: Job accepted, waiting to start provisioning
// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
// - `succeeded`: Destination ready and usable
// - `failed`: Provisioning failed — terminal state; delete and recreate.
//   A `failed` destination still counts against the per-instance quota until deleted.
const (
	Destination_State_Failed = "failed"
	Destination_State_NotStarted = "not_started"
	Destination_State_Provisioning = "provisioning"
	Destination_State_Succeeded = "succeeded"
)
func (*Destination) isaDestination() bool {
	return true
}

type DestinationIntf interface {
	isaDestination() bool
}

// UnmarshalDestination unmarshals an instance of Destination from the specified map of raw messages.
func UnmarshalDestination(m map[string]json.RawMessage, result interface{}) (err error) {
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
	if discValue == "ibm_cloud_database" {
		err = core.UnmarshalModel(m, "", result, UnmarshalIbmCloudDatabaseDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-IbmCloudDatabaseDestination-error", common.GetComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", common.GetComponentInfo())
	}
	return
}

// DestinationCollection : List of destinations for a Vault Dedicated cluster.
type DestinationCollection struct {
	// List of destinations.
	Destinations []DestinationIntf `json:"destinations" validate:"required"`

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

// GetInstanceOptions : The GetInstance options.
type GetInstanceOptions struct {
	// Secrets Manager instance ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInstanceOptions : Instantiate GetInstanceOptions
func (*SecretsManagerInstanceManagementV2) NewGetInstanceOptions(id string) *GetInstanceOptions {
	return &GetInstanceOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetInstanceOptions) SetID(id string) *GetInstanceOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceOptions) SetHeaders(param map[string]string) *GetInstanceOptions {
	options.Headers = param
	return options
}

// Instance : The service instance information.
type Instance struct {
	// The instance ID.
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

// Token : Admin Token response.
type Token struct {
	// The token value.
	Token *string `json:"token" validate:"required"`
}

// UnmarshalToken unmarshals an instance of Token from the specified map of raw messages.
func UnmarshalToken(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Token)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		err = core.SDKErrorf(err, "", "token-error", common.GetComponentInfo())
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
	RequestBody map[string]interface{} `json:"request_body" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateInstanceDestinationOptions : Instantiate UpdateInstanceDestinationOptions
func (*SecretsManagerInstanceManagementV2) NewUpdateInstanceDestinationOptions(instanceID string, destinationID string, requestBody map[string]interface{}) *UpdateInstanceDestinationOptions {
	return &UpdateInstanceDestinationOptions{
		InstanceID: core.StringPtr(instanceID),
		DestinationID: core.StringPtr(destinationID),
		RequestBody: requestBody,
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

// SetRequestBody : Allow user to set RequestBody
func (_options *UpdateInstanceDestinationOptions) SetRequestBody(requestBody map[string]interface{}) *UpdateInstanceDestinationOptions {
	_options.RequestBody = requestBody
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceDestinationOptions) SetHeaders(param map[string]string) *UpdateInstanceDestinationOptions {
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

// IbmCloudDatabaseDestination : A destination resource representing a private network link to an IBM Cloud Database service instance on a Vault
// Dedicated cluster.
// This model "extends" Destination
type IbmCloudDatabaseDestination struct {
	// Destination ID.
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
	// - `not_started`: Job accepted, waiting to start provisioning
	// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
	// - `succeeded`: Destination ready and usable
	// - `failed`: Provisioning failed — terminal state; delete and recreate.
	//   A `failed` destination still counts against the per-instance quota until deleted.
	State *string `json:"state" validate:"required"`

	// Timestamp when the destination was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Timestamp when the destination was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// IAM identity that created the destination.
	CreatedBy *string `json:"created_by,omitempty"`

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
// - `not_started`: Job accepted, waiting to start provisioning
// - `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`
// - `succeeded`: Destination ready and usable
// - `failed`: Provisioning failed — terminal state; delete and recreate.
//   A `failed` destination still counts against the per-instance quota until deleted.
const (
	IbmCloudDatabaseDestination_State_Failed = "failed"
	IbmCloudDatabaseDestination_State_NotStarted = "not_started"
	IbmCloudDatabaseDestination_State_Provisioning = "provisioning"
	IbmCloudDatabaseDestination_State_Succeeded = "succeeded"
)

func (*IbmCloudDatabaseDestination) isaDestination() bool {
	return true
}

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
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		err = core.SDKErrorf(err, "", "crn-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
