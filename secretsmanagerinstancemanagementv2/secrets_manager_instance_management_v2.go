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
 * IBM OpenAPI SDK Code Generator Version: 3.114.3-943fbc81-20260603-173645
 */

// Package secretsmanagerinstancemanagementv2 : Operations and models for the SecretsManagerInstanceManagementV2 service
package secretsmanagerinstancemanagementv2

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/secrets-manager-management-go-sdk/v2/common"
)

// SecretsManagerInstanceManagementV2 : With IBM Cloud® Secrets Manager Instance Management API, you can manage service
// instances of the Vault Dedicated plan. Use the API for the following operations:
// - Get service instance details including cluster state, endpoints, and key management service.
// - Generate a Vault admin token for authenticating to your Vault Dedicated cluster.
// - Revoke all active Vault admin tokens.
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

// CreateVaultAdmintoken : Generate admin token
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
		"instance_id": *createVaultAdmintokenOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/api/v2/instances/{instance_id}/admintokens`, pathParamsMap)
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

// DeleteInstanceAdmintokens : Revoke admin tokens
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
		"instance_id": *deleteInstanceAdmintokensOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/api/v2/instances/{instance_id}/admintokens`, pathParamsMap)
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
		"instance_id": *getInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = secretsManagerInstanceManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(secretsManagerInstanceManagement.Service.Options.URL, `/api/v2/instances/{instance_id}`, pathParamsMap)
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
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "2.0.0")
}

// CreateVaultAdmintokenOptions : The CreateVaultAdmintoken options.
type CreateVaultAdmintokenOptions struct {
	// The service instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

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

// SetHeaders : Allow user to set Headers
func (options *CreateVaultAdmintokenOptions) SetHeaders(param map[string]string) *CreateVaultAdmintokenOptions {
	options.Headers = param
	return options
}

// DeleteInstanceAdmintokensOptions : The DeleteInstanceAdmintokens options.
type DeleteInstanceAdmintokensOptions struct {
	// The service instance ID.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteInstanceAdmintokensOptions : Instantiate DeleteInstanceAdmintokensOptions
func (*SecretsManagerInstanceManagementV2) NewDeleteInstanceAdmintokensOptions(instanceID string) *DeleteInstanceAdmintokensOptions {
	return &DeleteInstanceAdmintokensOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteInstanceAdmintokensOptions) SetInstanceID(instanceID string) *DeleteInstanceAdmintokensOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteInstanceAdmintokensOptions) SetHeaders(param map[string]string) *DeleteInstanceAdmintokensOptions {
	options.Headers = param
	return options
}

// GetInstanceOptions : The GetInstance options.
type GetInstanceOptions struct {
	// The service instance ID.
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

// Instance : Instance.
type Instance struct {
	// Instance metadata for Vault Dedicated instances.
	Instance *VaultDedicatedInstanceMetadata `json:"instance" validate:"required"`

	// Vault cluster information for Vault Dedicated instances.
	VaultCluster *VaultDedicatedCluster `json:"vault_cluster" validate:"required"`

	// Instance endpoints for Vault Dedicated instances.
	Endpoints *VaultDedicatedInstanceEndpoints `json:"endpoints" validate:"required"`

	// Vault encryption configuration for Vault Dedicated instances.
	Encryption *VaultDedicatedInstanceEncryption `json:"encryption" validate:"required"`
}

// UnmarshalInstance unmarshals an instance of Instance from the specified map of raw messages.
func UnmarshalInstance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Instance)
	err = core.UnmarshalModel(m, "instance", &obj.Instance, UnmarshalVaultDedicatedInstanceMetadata)
	if err != nil {
		err = core.SDKErrorf(err, "", "instance-error", common.GetComponentInfo())
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

// VaultDedicatedInstanceMetadata : Instance metadata for Vault Dedicated instances.
type VaultDedicatedInstanceMetadata struct {
	// The instance CRN identifier.
	ID *string `json:"id" validate:"required"`

	// Instance plan information.
	Plan *VaultDedicatedInstancePlan `json:"plan" validate:"required"`
}

// UnmarshalVaultDedicatedInstanceMetadata unmarshals an instance of VaultDedicatedInstanceMetadata from the specified map of raw messages.
func UnmarshalVaultDedicatedInstanceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultDedicatedInstanceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "plan", &obj.Plan, UnmarshalVaultDedicatedInstancePlan)
	if err != nil {
		err = core.SDKErrorf(err, "", "plan-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VaultDedicatedInstancePlan : Instance plan information.
type VaultDedicatedInstancePlan struct {
	// The plan name of this instance.
	Name *string `json:"name" validate:"required"`
}

// Constants associated with the VaultDedicatedInstancePlan.Name property.
// The plan name of this instance.
const (
	VaultDedicatedInstancePlan_Name_Dedicated = "dedicated"
	VaultDedicatedInstancePlan_Name_Standard = "standard"
	VaultDedicatedInstancePlan_Name_Trial = "trial"
)

// UnmarshalVaultDedicatedInstancePlan unmarshals an instance of VaultDedicatedInstancePlan from the specified map of raw messages.
func UnmarshalVaultDedicatedInstancePlan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultDedicatedInstancePlan)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
