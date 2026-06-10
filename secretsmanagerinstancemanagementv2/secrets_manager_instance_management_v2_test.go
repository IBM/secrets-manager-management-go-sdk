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

package secretsmanagerinstancemanagementv2_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`SecretsManagerInstanceManagementV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(secretsManagerInstanceManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
				URL: "https://secretsmanagerinstancemanagementv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(secretsManagerInstanceManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECRETS_MANAGER_INSTANCE_MANAGEMENT_URL": "https://secretsmanagerinstancemanagementv2/api",
				"SECRETS_MANAGER_INSTANCE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2UsingExternalConfig(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
				})
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := secretsManagerInstanceManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != secretsManagerInstanceManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(secretsManagerInstanceManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(secretsManagerInstanceManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2UsingExternalConfig(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL: "https://testService/api",
				})
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := secretsManagerInstanceManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != secretsManagerInstanceManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(secretsManagerInstanceManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(secretsManagerInstanceManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2UsingExternalConfig(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
				})
				err := secretsManagerInstanceManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := secretsManagerInstanceManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != secretsManagerInstanceManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(secretsManagerInstanceManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(secretsManagerInstanceManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECRETS_MANAGER_INSTANCE_MANAGEMENT_URL": "https://secretsmanagerinstancemanagementv2/api",
				"SECRETS_MANAGER_INSTANCE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2UsingExternalConfig(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(secretsManagerInstanceManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECRETS_MANAGER_INSTANCE_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2UsingExternalConfig(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(secretsManagerInstanceManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = secretsmanagerinstancemanagementv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := secretsmanagerinstancemanagementv2.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.secrets-manager.cloud.ibm.com"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := secretsmanagerinstancemanagementv2.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`CreateVaultAdmintoken(createVaultAdmintokenOptions *CreateVaultAdmintokenOptions) - Operation response error`, func() {
		createVaultAdmintokenPath := "/api/v2/instances/60b40daa-1fd3-4f35-a994-2409cc0f270c/admintokens"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVaultAdmintokenPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateVaultAdmintoken with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateVaultAdmintokenOptions model
				createVaultAdmintokenOptionsModel := new(secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions)
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				createVaultAdmintokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVaultAdmintoken(createVaultAdmintokenOptions *CreateVaultAdmintokenOptions)`, func() {
		createVaultAdmintokenPath := "/api/v2/instances/60b40daa-1fd3-4f35-a994-2409cc0f270c/admintokens"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVaultAdmintokenPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"token": "hvs.CAESIIG_PILmULFYOsEyWHxkZ2mF2a8V...example...p3ZnpWbDF1RUNjUkNTZEg"}`)
				}))
			})
			It(`Invoke CreateVaultAdmintoken successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateVaultAdmintokenOptions model
				createVaultAdmintokenOptionsModel := new(secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions)
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				createVaultAdmintokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreateVaultAdmintokenWithContext(ctx, createVaultAdmintokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreateVaultAdmintokenWithContext(ctx, createVaultAdmintokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVaultAdmintokenPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"token": "hvs.CAESIIG_PILmULFYOsEyWHxkZ2mF2a8V...example...p3ZnpWbDF1RUNjUkNTZEg"}`)
				}))
			})
			It(`Invoke CreateVaultAdmintoken successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreateVaultAdmintoken(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateVaultAdmintokenOptions model
				createVaultAdmintokenOptionsModel := new(secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions)
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				createVaultAdmintokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVaultAdmintoken with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateVaultAdmintokenOptions model
				createVaultAdmintokenOptionsModel := new(secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions)
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				createVaultAdmintokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVaultAdmintokenOptions model with no property values
				createVaultAdmintokenOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateVaultAdmintoken successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateVaultAdmintokenOptions model
				createVaultAdmintokenOptionsModel := new(secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions)
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				createVaultAdmintokenOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptions *DeleteInstanceAdmintokensOptions)`, func() {
		deleteInstanceAdmintokensPath := "/api/v2/instances/60b40daa-1fd3-4f35-a994-2409cc0f270c/admintokens"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceAdmintokensPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteInstanceAdmintokens successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceAdmintokens(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteInstanceAdmintokensOptions model
				deleteInstanceAdmintokensOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceAdmintokensOptions)
				deleteInstanceAdmintokensOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				deleteInstanceAdmintokensOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteInstanceAdmintokens with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceAdmintokensOptions model
				deleteInstanceAdmintokensOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceAdmintokensOptions)
				deleteInstanceAdmintokensOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				deleteInstanceAdmintokensOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteInstanceAdmintokensOptions model with no property values
				deleteInstanceAdmintokensOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteInstanceAdmintokensOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstance(getInstanceOptions *GetInstanceOptions) - Operation response error`, func() {
		getInstancePath := "/api/v2/instances/60b40daa-1fd3-4f35-a994-2409cc0f270c"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstance with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstance(getInstanceOptions *GetInstanceOptions)`, func() {
		getInstancePath := "/api/v2/instances/60b40daa-1fd3-4f35-a994-2409cc0f270c"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"instance": {"id": "crn:v1:bluemix:public:secrets-manager:us-south:a/791f3fb10486421e97aa8512f18b7e65:b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5::", "plan": {"name": "standard"}}, "vault_cluster": {"status": "healthy", "version": "1.21.2+ent.hsm"}, "endpoints": {"public": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}, "private": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}}, "encryption": {"mode": "service_managed", "provider": "key_protect", "key_crn": "crn:v1:bluemix:public:kms:us-south:a/791f5fb10986423e97aa8512f18b7e65:31639268-42e8-4420-9872-590a6ee20506:key:b4af8f76-e6ea-4dc5-89cc-5f1b9bb207cc"}}`)
				}))
			})
			It(`Invoke GetInstance successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetInstanceWithContext(ctx, getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetInstanceWithContext(ctx, getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"instance": {"id": "crn:v1:bluemix:public:secrets-manager:us-south:a/791f3fb10486421e97aa8512f18b7e65:b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5::", "plan": {"name": "standard"}}, "vault_cluster": {"status": "healthy", "version": "1.21.2+ent.hsm"}, "endpoints": {"public": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}, "private": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}}, "encryption": {"mode": "service_managed", "provider": "key_protect", "key_crn": "crn:v1:bluemix:public:kms:us-south:a/791f5fb10986423e97aa8512f18b7e65:31639268-42e8-4420-9872-590a6ee20506:key:b4af8f76-e6ea-4dc5-89cc-5f1b9bb207cc"}}`)
				}))
			})
			It(`Invoke GetInstance successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstance with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceOptions model with no property values
				getInstanceOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstance(getInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetInstance successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			secretsManagerInstanceManagementService, _ := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
				URL:           "http://secretsmanagerinstancemanagementv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateVaultAdmintokenOptions successfully`, func() {
				// Construct an instance of the CreateVaultAdmintokenOptions model
				instanceID := "60b40daa-1fd3-4f35-a994-2409cc0f270c"
				createVaultAdmintokenOptionsModel := secretsManagerInstanceManagementService.NewCreateVaultAdmintokenOptions(instanceID)
				createVaultAdmintokenOptionsModel.SetInstanceID("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				createVaultAdmintokenOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVaultAdmintokenOptionsModel).ToNot(BeNil())
				Expect(createVaultAdmintokenOptionsModel.InstanceID).To(Equal(core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")))
				Expect(createVaultAdmintokenOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteInstanceAdmintokensOptions successfully`, func() {
				// Construct an instance of the DeleteInstanceAdmintokensOptions model
				instanceID := "60b40daa-1fd3-4f35-a994-2409cc0f270c"
				deleteInstanceAdmintokensOptionsModel := secretsManagerInstanceManagementService.NewDeleteInstanceAdmintokensOptions(instanceID)
				deleteInstanceAdmintokensOptionsModel.SetInstanceID("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				deleteInstanceAdmintokensOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteInstanceAdmintokensOptionsModel).ToNot(BeNil())
				Expect(deleteInstanceAdmintokensOptionsModel.InstanceID).To(Equal(core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")))
				Expect(deleteInstanceAdmintokensOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceOptions successfully`, func() {
				// Construct an instance of the GetInstanceOptions model
				instanceID := "60b40daa-1fd3-4f35-a994-2409cc0f270c"
				getInstanceOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceOptions(instanceID)
				getInstanceOptionsModel.SetInstanceID("60b40daa-1fd3-4f35-a994-2409cc0f270c")
				getInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceOptionsModel).ToNot(BeNil())
				Expect(getInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("60b40daa-1fd3-4f35-a994-2409cc0f270c")))
				Expect(getInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
