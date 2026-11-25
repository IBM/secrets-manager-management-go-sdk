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
	"encoding/json"
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
		createVaultAdmintokenPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/admintokens"
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
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createVaultAdmintokenOptionsModel.ResponseWrapping = core.BoolPtr(true)
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
		createVaultAdmintokenPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/admintokens"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVaultAdmintokenPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"token": "hvs.CAESIC202dQW-jogohoywlivU-AMllTdeNT0QBiob9OBmyfnGigKImh2cy5yUkxpcmVaZURzTmw5cgtevFoyelBpYkguYW82SlMQyr0Q"}`)
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
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createVaultAdmintokenOptionsModel.ResponseWrapping = core.BoolPtr(true)
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

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"token": "hvs.CAESIC202dQW-jogohoywlivU-AMllTdeNT0QBiob9OBmyfnGigKImh2cy5yUkxpcmVaZURzTmw5cgtevFoyelBpYkguYW82SlMQyr0Q"}`)
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
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createVaultAdmintokenOptionsModel.ResponseWrapping = core.BoolPtr(true)
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
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createVaultAdmintokenOptionsModel.ResponseWrapping = core.BoolPtr(true)
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
				createVaultAdmintokenOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createVaultAdmintokenOptionsModel.ResponseWrapping = core.BoolPtr(true)
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
	Describe(`DeleteVaultAdmintokens(deleteVaultAdmintokensOptions *DeleteVaultAdmintokensOptions)`, func() {
		deleteVaultAdmintokensPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/admintokens"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVaultAdmintokensPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteVaultAdmintokens successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := secretsManagerInstanceManagementService.DeleteVaultAdmintokens(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVaultAdmintokensOptions model
				deleteVaultAdmintokensOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteVaultAdmintokensOptions)
				deleteVaultAdmintokensOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteVaultAdmintokensOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteVaultAdmintokens(deleteVaultAdmintokensOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteVaultAdmintokens with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteVaultAdmintokensOptions model
				deleteVaultAdmintokensOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteVaultAdmintokensOptions)
				deleteVaultAdmintokensOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteVaultAdmintokensOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := secretsManagerInstanceManagementService.DeleteVaultAdmintokens(deleteVaultAdmintokensOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteVaultAdmintokensOptions model with no property values
				deleteVaultAdmintokensOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteVaultAdmintokensOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteVaultAdmintokens(deleteVaultAdmintokensOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstance(getInstanceOptions *GetInstanceOptions) - Operation response error`, func() {
		getInstancePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1"
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
				getInstanceOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
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
		getInstancePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1"
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
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "name": "Name", "instance_crn": "crn:v1:bluemix:public:secrets-manager:us-south:a/791f3fb10486421e97aa8512f18b7e65:b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5::", "plan": "dedicated", "vault_cluster": {"status": "healthy", "version": "1.21.2+ent.hsm"}, "endpoints": {"public": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}, "private": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}}, "encryption": {"mode": "service_managed", "provider": "key_protect", "key_crn": "crn:v1:bluemix:public:kms:us-south:a/791f5fb10986423e97aa8512f18b7e65:31639268-42e8-4420-9872-590a6ee20506:key:b4af8f76-e6ea-4dc5-89cc-5f1b9bb207cc"}, "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1"}`)
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
				getInstanceOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
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
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "name": "Name", "instance_crn": "crn:v1:bluemix:public:secrets-manager:us-south:a/791f3fb10486421e97aa8512f18b7e65:b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5::", "plan": "dedicated", "vault_cluster": {"status": "healthy", "version": "1.21.2+ent.hsm"}, "endpoints": {"public": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}, "private": {"vault_api": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud", "vault_ui": "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"}}, "encryption": {"mode": "service_managed", "provider": "key_protect", "key_crn": "crn:v1:bluemix:public:kms:us-south:a/791f5fb10986423e97aa8512f18b7e65:31639268-42e8-4420-9872-590a6ee20506:key:b4af8f76-e6ea-4dc5-89cc-5f1b9bb207cc"}, "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1"}`)
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
				getInstanceOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
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
				getInstanceOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
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
				getInstanceOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
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
	Describe(`ListInstanceDestinations(listInstanceDestinationsOptions *ListInstanceDestinationsOptions) - Operation response error`, func() {
		listInstanceDestinationsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceDestinationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"not_started"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceDestinations with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceDestinationsOptions model
				listInstanceDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions)
				listInstanceDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceDestinationsOptionsModel.State = core.StringPtr("not_started")
				listInstanceDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceDestinations(listInstanceDestinationsOptions *ListInstanceDestinationsOptions)`, func() {
		listInstanceDestinationsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["state"]).To(Equal([]string{"not_started"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"destinations": [{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision"}], "total": 0}`)
				}))
			})
			It(`Invoke ListInstanceDestinations successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListInstanceDestinationsOptions model
				listInstanceDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions)
				listInstanceDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceDestinationsOptionsModel.State = core.StringPtr("not_started")
				listInstanceDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ListInstanceDestinationsWithContext(ctx, listInstanceDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ListInstanceDestinationsWithContext(ctx, listInstanceDestinationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["state"]).To(Equal([]string{"not_started"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"destinations": [{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision"}], "total": 0}`)
				}))
			})
			It(`Invoke ListInstanceDestinations successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceDestinations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceDestinationsOptions model
				listInstanceDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions)
				listInstanceDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceDestinationsOptionsModel.State = core.StringPtr("not_started")
				listInstanceDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListInstanceDestinations with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceDestinationsOptions model
				listInstanceDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions)
				listInstanceDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceDestinationsOptionsModel.State = core.StringPtr("not_started")
				listInstanceDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListInstanceDestinationsOptions model with no property values
				listInstanceDestinationsOptionsModelNew := new(secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptionsModelNew)
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
			It(`Invoke ListInstanceDestinations successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceDestinationsOptions model
				listInstanceDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions)
				listInstanceDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceDestinationsOptionsModel.State = core.StringPtr("not_started")
				listInstanceDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptionsModel)
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
	Describe(`CreateInstanceDestination(createInstanceDestinationOptions *CreateInstanceDestinationOptions) - Operation response error`, func() {
		createInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceDestinationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateInstanceDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateDestinationRequestIbmCloudDatabaseDestinationPrototype model
				createDestinationRequestModel := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
				createDestinationRequestModel.Name = core.StringPtr("my-postgres")
				createDestinationRequestModel.Type = core.StringPtr("ibm_cloud_database")
				createDestinationRequestModel.Description = core.StringPtr("Production PostgreSQL database")
				createDestinationRequestModel.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

				// Construct an instance of the CreateInstanceDestinationOptions model
				createInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions)
				createInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceDestinationOptionsModel.DestinationPrototype = createDestinationRequestModel
				createInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateInstanceDestination(createInstanceDestinationOptions *CreateInstanceDestinationOptions)`, func() {
		createInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke CreateInstanceDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateDestinationRequestIbmCloudDatabaseDestinationPrototype model
				createDestinationRequestModel := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
				createDestinationRequestModel.Name = core.StringPtr("my-postgres")
				createDestinationRequestModel.Type = core.StringPtr("ibm_cloud_database")
				createDestinationRequestModel.Description = core.StringPtr("Production PostgreSQL database")
				createDestinationRequestModel.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

				// Construct an instance of the CreateInstanceDestinationOptions model
				createInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions)
				createInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceDestinationOptionsModel.DestinationPrototype = createDestinationRequestModel
				createInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreateInstanceDestinationWithContext(ctx, createInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreateInstanceDestinationWithContext(ctx, createInstanceDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke CreateInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDestinationRequestIbmCloudDatabaseDestinationPrototype model
				createDestinationRequestModel := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
				createDestinationRequestModel.Name = core.StringPtr("my-postgres")
				createDestinationRequestModel.Type = core.StringPtr("ibm_cloud_database")
				createDestinationRequestModel.Description = core.StringPtr("Production PostgreSQL database")
				createDestinationRequestModel.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

				// Construct an instance of the CreateInstanceDestinationOptions model
				createInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions)
				createInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceDestinationOptionsModel.DestinationPrototype = createDestinationRequestModel
				createInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateInstanceDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateDestinationRequestIbmCloudDatabaseDestinationPrototype model
				createDestinationRequestModel := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
				createDestinationRequestModel.Name = core.StringPtr("my-postgres")
				createDestinationRequestModel.Type = core.StringPtr("ibm_cloud_database")
				createDestinationRequestModel.Description = core.StringPtr("Production PostgreSQL database")
				createDestinationRequestModel.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

				// Construct an instance of the CreateInstanceDestinationOptions model
				createInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions)
				createInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceDestinationOptionsModel.DestinationPrototype = createDestinationRequestModel
				createInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateInstanceDestinationOptions model with no property values
				createInstanceDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateDestinationRequestIbmCloudDatabaseDestinationPrototype model
				createDestinationRequestModel := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
				createDestinationRequestModel.Name = core.StringPtr("my-postgres")
				createDestinationRequestModel.Type = core.StringPtr("ibm_cloud_database")
				createDestinationRequestModel.Description = core.StringPtr("Production PostgreSQL database")
				createDestinationRequestModel.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

				// Construct an instance of the CreateInstanceDestinationOptions model
				createInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions)
				createInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceDestinationOptionsModel.DestinationPrototype = createDestinationRequestModel
				createInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptionsModel)
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
	Describe(`GetInstanceDestination(getInstanceDestinationOptions *GetInstanceDestinationOptions) - Operation response error`, func() {
		getInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceDestinationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceDestinationOptions model
				getInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions)
				getInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceDestination(getInstanceDestinationOptions *GetInstanceDestinationOptions)`, func() {
		getInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke GetInstanceDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceDestinationOptions model
				getInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions)
				getInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetInstanceDestinationWithContext(ctx, getInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetInstanceDestinationWithContext(ctx, getInstanceDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke GetInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceDestinationOptions model
				getInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions)
				getInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceDestinationOptions model
				getInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions)
				getInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceDestinationOptions model with no property values
				getInstanceDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptionsModelNew)
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
			It(`Invoke GetInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceDestinationOptions model
				getInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions)
				getInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptionsModel)
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
	Describe(`UpdateInstanceDestination(updateInstanceDestinationOptions *UpdateInstanceDestinationOptions) - Operation response error`, func() {
		updateInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DestinationPatch model
				destinationPatchModel := new(secretsmanagerinstancemanagementv2.DestinationPatch)
				destinationPatchModel.Name = core.StringPtr("my-postgres-updated")
				destinationPatchModel.Description = core.StringPtr("Production PostgreSQL database")
				destinationPatchModelAsPatch, asPatchErr := destinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceDestinationOptions model
				updateInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions)
				updateInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceDestinationOptionsModel.DestinationPatch = destinationPatchModelAsPatch
				updateInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceDestination(updateInstanceDestinationOptions *UpdateInstanceDestinationOptions)`, func() {
		updateInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke UpdateInstanceDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the DestinationPatch model
				destinationPatchModel := new(secretsmanagerinstancemanagementv2.DestinationPatch)
				destinationPatchModel.Name = core.StringPtr("my-postgres-updated")
				destinationPatchModel.Description = core.StringPtr("Production PostgreSQL database")
				destinationPatchModelAsPatch, asPatchErr := destinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceDestinationOptions model
				updateInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions)
				updateInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceDestinationOptionsModel.DestinationPatch = destinationPatchModelAsPatch
				updateInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.UpdateInstanceDestinationWithContext(ctx, updateInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.UpdateInstanceDestinationWithContext(ctx, updateInstanceDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke UpdateInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DestinationPatch model
				destinationPatchModel := new(secretsmanagerinstancemanagementv2.DestinationPatch)
				destinationPatchModel.Name = core.StringPtr("my-postgres-updated")
				destinationPatchModel.Description = core.StringPtr("Production PostgreSQL database")
				destinationPatchModelAsPatch, asPatchErr := destinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceDestinationOptions model
				updateInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions)
				updateInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceDestinationOptionsModel.DestinationPatch = destinationPatchModelAsPatch
				updateInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DestinationPatch model
				destinationPatchModel := new(secretsmanagerinstancemanagementv2.DestinationPatch)
				destinationPatchModel.Name = core.StringPtr("my-postgres-updated")
				destinationPatchModel.Description = core.StringPtr("Production PostgreSQL database")
				destinationPatchModelAsPatch, asPatchErr := destinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceDestinationOptions model
				updateInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions)
				updateInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceDestinationOptionsModel.DestinationPatch = destinationPatchModelAsPatch
				updateInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceDestinationOptions model with no property values
				updateInstanceDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptionsModelNew)
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
			It(`Invoke UpdateInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DestinationPatch model
				destinationPatchModel := new(secretsmanagerinstancemanagementv2.DestinationPatch)
				destinationPatchModel.Name = core.StringPtr("my-postgres-updated")
				destinationPatchModel.Description = core.StringPtr("Production PostgreSQL database")
				destinationPatchModelAsPatch, asPatchErr := destinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceDestinationOptions model
				updateInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions)
				updateInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceDestinationOptionsModel.DestinationPatch = destinationPatchModelAsPatch
				updateInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptionsModel)
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
	Describe(`DeleteInstanceDestination(deleteInstanceDestinationOptions *DeleteInstanceDestinationOptions) - Operation response error`, func() {
		deleteInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteInstanceDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceDestinationOptions model
				deleteInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions)
				deleteInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteInstanceDestination(deleteInstanceDestinationOptions *DeleteInstanceDestinationOptions)`, func() {
		deleteInstanceDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke DeleteInstanceDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the DeleteInstanceDestinationOptions model
				deleteInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions)
				deleteInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.DeleteInstanceDestinationWithContext(ctx, deleteInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.DeleteInstanceDestinationWithContext(ctx, deleteInstanceDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb", "name": "Name", "type": "ibm_cloud_database", "description": "Description", "state": "not_started", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "message": "failed to provision", "crn": "Crn"}`)
				}))
			})
			It(`Invoke DeleteInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteInstanceDestinationOptions model
				deleteInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions)
				deleteInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteInstanceDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceDestinationOptions model
				deleteInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions)
				deleteInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteInstanceDestinationOptions model with no property values
				deleteInstanceDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteInstanceDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceDestinationOptions model
				deleteInstanceDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions)
				deleteInstanceDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptionsModel)
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
	Describe(`ListInstanceHsmDestinations(listInstanceHsmDestinationsOptions *ListInstanceHsmDestinationsOptions) - Operation response error`, func() {
		listInstanceHsmDestinationsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinations with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationsOptions model
				listInstanceHsmDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions)
				listInstanceHsmDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceHsmDestinations(listInstanceHsmDestinationsOptions *ListInstanceHsmDestinationsOptions)`, func() {
		listInstanceHsmDestinationsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "hsm_destinations": [{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinations successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListInstanceHsmDestinationsOptions model
				listInstanceHsmDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions)
				listInstanceHsmDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationsWithContext(ctx, listInstanceHsmDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationsWithContext(ctx, listInstanceHsmDestinationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "hsm_destinations": [{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinations successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationsOptions model
				listInstanceHsmDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions)
				listInstanceHsmDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListInstanceHsmDestinations with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationsOptions model
				listInstanceHsmDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions)
				listInstanceHsmDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListInstanceHsmDestinationsOptions model with no property values
				listInstanceHsmDestinationsOptionsModelNew := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptionsModelNew)
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
			It(`Invoke ListInstanceHsmDestinations successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationsOptions model
				listInstanceHsmDestinationsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions)
				listInstanceHsmDestinationsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptionsModel)
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
	Describe(`CreateInstanceHsmDestination(createInstanceHsmDestinationOptions *CreateInstanceHsmDestinationOptions) - Operation response error`, func() {
		createInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationOptions model
				createInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions)
				createInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationOptionsModel.Name = core.StringPtr("prod-hsm-destination")
				createInstanceHsmDestinationOptionsModel.Vendor = core.StringPtr("thales_luna")
				createInstanceHsmDestinationOptionsModel.Description = core.StringPtr("Primary Thales Luna destination")
				createInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateInstanceHsmDestination(createInstanceHsmDestinationOptions *CreateInstanceHsmDestinationOptions)`, func() {
		createInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateInstanceHsmDestinationOptions model
				createInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions)
				createInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationOptionsModel.Name = core.StringPtr("prod-hsm-destination")
				createInstanceHsmDestinationOptionsModel.Vendor = core.StringPtr("thales_luna")
				createInstanceHsmDestinationOptionsModel.Description = core.StringPtr("Primary Thales Luna destination")
				createInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationWithContext(ctx, createInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationWithContext(ctx, createInstanceHsmDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationOptions model
				createInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions)
				createInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationOptionsModel.Name = core.StringPtr("prod-hsm-destination")
				createInstanceHsmDestinationOptionsModel.Vendor = core.StringPtr("thales_luna")
				createInstanceHsmDestinationOptionsModel.Description = core.StringPtr("Primary Thales Luna destination")
				createInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateInstanceHsmDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationOptions model
				createInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions)
				createInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationOptionsModel.Name = core.StringPtr("prod-hsm-destination")
				createInstanceHsmDestinationOptionsModel.Vendor = core.StringPtr("thales_luna")
				createInstanceHsmDestinationOptionsModel.Description = core.StringPtr("Primary Thales Luna destination")
				createInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateInstanceHsmDestinationOptions model with no property values
				createInstanceHsmDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptionsModelNew)
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
			It(`Invoke CreateInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationOptions model
				createInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions)
				createInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationOptionsModel.Name = core.StringPtr("prod-hsm-destination")
				createInstanceHsmDestinationOptionsModel.Vendor = core.StringPtr("thales_luna")
				createInstanceHsmDestinationOptionsModel.Description = core.StringPtr("Primary Thales Luna destination")
				createInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptionsModel)
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
	Describe(`GetInstanceHsmDestination(getInstanceHsmDestinationOptions *GetInstanceHsmDestinationOptions) - Operation response error`, func() {
		getInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceHsmDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationOptions model
				getInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions)
				getInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceHsmDestination(getInstanceHsmDestinationOptions *GetInstanceHsmDestinationOptions)`, func() {
		getInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceHsmDestinationOptions model
				getInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions)
				getInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationWithContext(ctx, getInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationWithContext(ctx, getInstanceHsmDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationOptions model
				getInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions)
				getInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceHsmDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationOptions model
				getInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions)
				getInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceHsmDestinationOptions model with no property values
				getInstanceHsmDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptionsModelNew)
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
			It(`Invoke GetInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationOptions model
				getInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions)
				getInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptionsModel)
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
	Describe(`UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptions *UpdateInstanceHsmDestinationOptions) - Operation response error`, func() {
		updateInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the HSMDestinationPatch model
				hsmDestinationPatchModel := new(secretsmanagerinstancemanagementv2.HSMDestinationPatch)
				hsmDestinationPatchModel.Name = core.StringPtr("my-hsm-destination")
				hsmDestinationPatchModel.Description = core.StringPtr("testString")
				hsmDestinationPatchModelAsPatch, asPatchErr := hsmDestinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationOptions model
				updateInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions)
				updateInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationOptionsModel.HsmDestinationPatch = hsmDestinationPatchModelAsPatch
				updateInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptions *UpdateInstanceHsmDestinationOptions)`, func() {
		updateInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the HSMDestinationPatch model
				hsmDestinationPatchModel := new(secretsmanagerinstancemanagementv2.HSMDestinationPatch)
				hsmDestinationPatchModel.Name = core.StringPtr("my-hsm-destination")
				hsmDestinationPatchModel.Description = core.StringPtr("testString")
				hsmDestinationPatchModelAsPatch, asPatchErr := hsmDestinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationOptions model
				updateInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions)
				updateInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationOptionsModel.HsmDestinationPatch = hsmDestinationPatchModelAsPatch
				updateInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationWithContext(ctx, updateInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationWithContext(ctx, updateInstanceHsmDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "account_id": "a/abc123def456", "name": "prod-hsm-destination", "description": "Primary Thales Luna destination", "vendor": "thales_luna", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HSMDestinationPatch model
				hsmDestinationPatchModel := new(secretsmanagerinstancemanagementv2.HSMDestinationPatch)
				hsmDestinationPatchModel.Name = core.StringPtr("my-hsm-destination")
				hsmDestinationPatchModel.Description = core.StringPtr("testString")
				hsmDestinationPatchModelAsPatch, asPatchErr := hsmDestinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationOptions model
				updateInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions)
				updateInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationOptionsModel.HsmDestinationPatch = hsmDestinationPatchModelAsPatch
				updateInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceHsmDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the HSMDestinationPatch model
				hsmDestinationPatchModel := new(secretsmanagerinstancemanagementv2.HSMDestinationPatch)
				hsmDestinationPatchModel.Name = core.StringPtr("my-hsm-destination")
				hsmDestinationPatchModel.Description = core.StringPtr("testString")
				hsmDestinationPatchModelAsPatch, asPatchErr := hsmDestinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationOptions model
				updateInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions)
				updateInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationOptionsModel.HsmDestinationPatch = hsmDestinationPatchModelAsPatch
				updateInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceHsmDestinationOptions model with no property values
				updateInstanceHsmDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptionsModelNew)
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
			It(`Invoke UpdateInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the HSMDestinationPatch model
				hsmDestinationPatchModel := new(secretsmanagerinstancemanagementv2.HSMDestinationPatch)
				hsmDestinationPatchModel.Name = core.StringPtr("my-hsm-destination")
				hsmDestinationPatchModel.Description = core.StringPtr("testString")
				hsmDestinationPatchModelAsPatch, asPatchErr := hsmDestinationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationOptions model
				updateInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions)
				updateInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationOptionsModel.HsmDestinationPatch = hsmDestinationPatchModelAsPatch
				updateInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptionsModel)
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
	Describe(`DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptions *DeleteInstanceHsmDestinationOptions) - Operation response error`, func() {
		deleteInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationOptions model
				deleteInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions)
				deleteInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptions *DeleteInstanceHsmDestinationOptions)`, func() {
		deleteInstanceHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "provisioning_state": "deleting"}`)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the DeleteInstanceHsmDestinationOptions model
				deleteInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions)
				deleteInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationWithContext(ctx, deleteInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationWithContext(ctx, deleteInstanceHsmDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02", "id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "provisioning_state": "deleting"}`)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationOptions model
				deleteInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions)
				deleteInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteInstanceHsmDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationOptions model
				deleteInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions)
				deleteInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteInstanceHsmDestinationOptions model with no property values
				deleteInstanceHsmDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationOptions model
				deleteInstanceHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions)
				deleteInstanceHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptionsModel)
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
	Describe(`GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptions *GetHsmDestinationHealthSnapshotOptions) - Operation response error`, func() {
		getHsmDestinationHealthSnapshotPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/health"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHsmDestinationHealthSnapshotPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHsmDestinationHealthSnapshot with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetHsmDestinationHealthSnapshotOptions model
				getHsmDestinationHealthSnapshotOptionsModel := new(secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions)
				getHsmDestinationHealthSnapshotOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getHsmDestinationHealthSnapshotOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getHsmDestinationHealthSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptions *GetHsmDestinationHealthSnapshotOptions)`, func() {
		getHsmDestinationHealthSnapshotPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/health"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHsmDestinationHealthSnapshotPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/health", "destination_id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "client_certificates": [{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "health_state": "valid"}], "ha_groups": [{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "health_state": "healthy"}], "paths": [{"id": "fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "health_state": "healthy", "server_cas": [{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "health_state": "valid"}]}]}`)
				}))
			})
			It(`Invoke GetHsmDestinationHealthSnapshot successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetHsmDestinationHealthSnapshotOptions model
				getHsmDestinationHealthSnapshotOptionsModel := new(secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions)
				getHsmDestinationHealthSnapshotOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getHsmDestinationHealthSnapshotOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getHsmDestinationHealthSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshotWithContext(ctx, getHsmDestinationHealthSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshotWithContext(ctx, getHsmDestinationHealthSnapshotOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getHsmDestinationHealthSnapshotPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/health", "destination_id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "client_certificates": [{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "health_state": "valid"}], "ha_groups": [{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "health_state": "healthy"}], "paths": [{"id": "fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "health_state": "healthy", "server_cas": [{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "health_state": "valid"}]}]}`)
				}))
			})
			It(`Invoke GetHsmDestinationHealthSnapshot successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHsmDestinationHealthSnapshotOptions model
				getHsmDestinationHealthSnapshotOptionsModel := new(secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions)
				getHsmDestinationHealthSnapshotOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getHsmDestinationHealthSnapshotOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getHsmDestinationHealthSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetHsmDestinationHealthSnapshot with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetHsmDestinationHealthSnapshotOptions model
				getHsmDestinationHealthSnapshotOptionsModel := new(secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions)
				getHsmDestinationHealthSnapshotOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getHsmDestinationHealthSnapshotOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getHsmDestinationHealthSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetHsmDestinationHealthSnapshotOptions model with no property values
				getHsmDestinationHealthSnapshotOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptionsModelNew)
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
			It(`Invoke GetHsmDestinationHealthSnapshot successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetHsmDestinationHealthSnapshotOptions model
				getHsmDestinationHealthSnapshotOptionsModel := new(secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions)
				getHsmDestinationHealthSnapshotOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getHsmDestinationHealthSnapshotOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getHsmDestinationHealthSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptionsModel)
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
	Describe(`RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptions *RunHsmDestinationHealthCheckOptions) - Operation response error`, func() {
		runHsmDestinationHealthCheckPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/health"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runHsmDestinationHealthCheckPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunHsmDestinationHealthCheck with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the RunHsmDestinationHealthCheckOptions model
				runHsmDestinationHealthCheckOptionsModel := new(secretsmanagerinstancemanagementv2.RunHsmDestinationHealthCheckOptions)
				runHsmDestinationHealthCheckOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				runHsmDestinationHealthCheckOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				runHsmDestinationHealthCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptions *RunHsmDestinationHealthCheckOptions)`, func() {
		runHsmDestinationHealthCheckPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/health"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runHsmDestinationHealthCheckPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/health", "destination_id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "client_certificates": [{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "health_state": "valid"}], "ha_groups": [{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "health_state": "healthy"}], "paths": [{"id": "fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "health_state": "healthy", "server_cas": [{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "health_state": "valid"}]}]}`)
				}))
			})
			It(`Invoke RunHsmDestinationHealthCheck successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the RunHsmDestinationHealthCheckOptions model
				runHsmDestinationHealthCheckOptionsModel := new(secretsmanagerinstancemanagementv2.RunHsmDestinationHealthCheckOptions)
				runHsmDestinationHealthCheckOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				runHsmDestinationHealthCheckOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				runHsmDestinationHealthCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheckWithContext(ctx, runHsmDestinationHealthCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.RunHsmDestinationHealthCheckWithContext(ctx, runHsmDestinationHealthCheckOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(runHsmDestinationHealthCheckPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/health", "destination_id": "b49ad385-0cb7-3694-a71c-2579d6e94b02", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "client_certificates": [{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "health_state": "valid"}], "ha_groups": [{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "health_state": "healthy"}], "paths": [{"id": "fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "health_state": "healthy", "server_cas": [{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "health_state": "valid"}]}]}`)
				}))
			})
			It(`Invoke RunHsmDestinationHealthCheck successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RunHsmDestinationHealthCheckOptions model
				runHsmDestinationHealthCheckOptionsModel := new(secretsmanagerinstancemanagementv2.RunHsmDestinationHealthCheckOptions)
				runHsmDestinationHealthCheckOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				runHsmDestinationHealthCheckOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				runHsmDestinationHealthCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RunHsmDestinationHealthCheck with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the RunHsmDestinationHealthCheckOptions model
				runHsmDestinationHealthCheckOptionsModel := new(secretsmanagerinstancemanagementv2.RunHsmDestinationHealthCheckOptions)
				runHsmDestinationHealthCheckOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				runHsmDestinationHealthCheckOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				runHsmDestinationHealthCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RunHsmDestinationHealthCheckOptions model with no property values
				runHsmDestinationHealthCheckOptionsModelNew := new(secretsmanagerinstancemanagementv2.RunHsmDestinationHealthCheckOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke RunHsmDestinationHealthCheck successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the RunHsmDestinationHealthCheckOptions model
				runHsmDestinationHealthCheckOptionsModel := new(secretsmanagerinstancemanagementv2.RunHsmDestinationHealthCheckOptions)
				runHsmDestinationHealthCheckOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				runHsmDestinationHealthCheckOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				runHsmDestinationHealthCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptionsModel)
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
	Describe(`ApplyHsmDestination(applyHsmDestinationOptions *ApplyHsmDestinationOptions) - Operation response error`, func() {
		applyHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/apply"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(applyHsmDestinationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ApplyHsmDestination with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ApplyHsmDestinationOptions model
				applyHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions)
				applyHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				applyHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				applyHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ApplyHsmDestination(applyHsmDestinationOptions *ApplyHsmDestinationOptions)`, func() {
		applyHsmDestinationPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/apply"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(applyHsmDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/health"}`)
				}))
			})
			It(`Invoke ApplyHsmDestination successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ApplyHsmDestinationOptions model
				applyHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions)
				applyHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				applyHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				applyHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ApplyHsmDestinationWithContext(ctx, applyHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ApplyHsmDestinationWithContext(ctx, applyHsmDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(applyHsmDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/health"}`)
				}))
			})
			It(`Invoke ApplyHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ApplyHsmDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApplyHsmDestinationOptions model
				applyHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions)
				applyHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				applyHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				applyHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ApplyHsmDestination with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ApplyHsmDestinationOptions model
				applyHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions)
				applyHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				applyHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				applyHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ApplyHsmDestinationOptions model with no property values
				applyHsmDestinationOptionsModelNew := new(secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke ApplyHsmDestination successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ApplyHsmDestinationOptions model
				applyHsmDestinationOptionsModel := new(secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions)
				applyHsmDestinationOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				applyHsmDestinationOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				applyHsmDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptionsModel)
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
	Describe(`ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptions *ListInstanceHsmDestinationHaGroupsOptions) - Operation response error`, func() {
		listInstanceHsmDestinationHaGroupsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationHaGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationHaGroups with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationHaGroupsOptions model
				listInstanceHsmDestinationHaGroupsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions)
				listInstanceHsmDestinationHaGroupsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationHaGroupsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationHaGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptions *ListInstanceHsmDestinationHaGroupsOptions)`, func() {
		listInstanceHsmDestinationHaGroupsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationHaGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ha_groups": [{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationHaGroups successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListInstanceHsmDestinationHaGroupsOptions model
				listInstanceHsmDestinationHaGroupsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions)
				listInstanceHsmDestinationHaGroupsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationHaGroupsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationHaGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroupsWithContext(ctx, listInstanceHsmDestinationHaGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroupsWithContext(ctx, listInstanceHsmDestinationHaGroupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationHaGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ha_groups": [{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationHaGroups successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationHaGroupsOptions model
				listInstanceHsmDestinationHaGroupsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions)
				listInstanceHsmDestinationHaGroupsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationHaGroupsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationHaGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListInstanceHsmDestinationHaGroups with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationHaGroupsOptions model
				listInstanceHsmDestinationHaGroupsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions)
				listInstanceHsmDestinationHaGroupsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationHaGroupsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationHaGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListInstanceHsmDestinationHaGroupsOptions model with no property values
				listInstanceHsmDestinationHaGroupsOptionsModelNew := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptionsModelNew)
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
			It(`Invoke ListInstanceHsmDestinationHaGroups successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationHaGroupsOptions model
				listInstanceHsmDestinationHaGroupsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions)
				listInstanceHsmDestinationHaGroupsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationHaGroupsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationHaGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptionsModel)
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
	Describe(`CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptions *CreateInstanceHsmDestinationHaGroupOptions) - Operation response error`, func() {
		createInstanceHsmDestinationHaGroupPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationHaGroup with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationHaGroupOptions model
				createInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions)
				createInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationHaGroupOptionsModel.Name = core.StringPtr("production-ha-group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupLabel = core.StringPtr("prod_ha_group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupSerial = core.StringPtr("11323465009866")
				createInstanceHsmDestinationHaGroupOptionsModel.MemberSerials = []string{"123456", "123457"}
				createInstanceHsmDestinationHaGroupOptionsModel.Description = core.StringPtr("Primary HA Group for prod")
				createInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptions *CreateInstanceHsmDestinationHaGroupOptions)`, func() {
		createInstanceHsmDestinationHaGroupPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationHaGroup successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateInstanceHsmDestinationHaGroupOptions model
				createInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions)
				createInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationHaGroupOptionsModel.Name = core.StringPtr("production-ha-group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupLabel = core.StringPtr("prod_ha_group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupSerial = core.StringPtr("11323465009866")
				createInstanceHsmDestinationHaGroupOptionsModel.MemberSerials = []string{"123456", "123457"}
				createInstanceHsmDestinationHaGroupOptionsModel.Description = core.StringPtr("Primary HA Group for prod")
				createInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroupWithContext(ctx, createInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroupWithContext(ctx, createInstanceHsmDestinationHaGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationHaGroup successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationHaGroupOptions model
				createInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions)
				createInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationHaGroupOptionsModel.Name = core.StringPtr("production-ha-group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupLabel = core.StringPtr("prod_ha_group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupSerial = core.StringPtr("11323465009866")
				createInstanceHsmDestinationHaGroupOptionsModel.MemberSerials = []string{"123456", "123457"}
				createInstanceHsmDestinationHaGroupOptionsModel.Description = core.StringPtr("Primary HA Group for prod")
				createInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateInstanceHsmDestinationHaGroup with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationHaGroupOptions model
				createInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions)
				createInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationHaGroupOptionsModel.Name = core.StringPtr("production-ha-group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupLabel = core.StringPtr("prod_ha_group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupSerial = core.StringPtr("11323465009866")
				createInstanceHsmDestinationHaGroupOptionsModel.MemberSerials = []string{"123456", "123457"}
				createInstanceHsmDestinationHaGroupOptionsModel.Description = core.StringPtr("Primary HA Group for prod")
				createInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateInstanceHsmDestinationHaGroupOptions model with no property values
				createInstanceHsmDestinationHaGroupOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptionsModelNew)
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
			It(`Invoke CreateInstanceHsmDestinationHaGroup successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationHaGroupOptions model
				createInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions)
				createInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationHaGroupOptionsModel.Name = core.StringPtr("production-ha-group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupLabel = core.StringPtr("prod_ha_group")
				createInstanceHsmDestinationHaGroupOptionsModel.GroupSerial = core.StringPtr("11323465009866")
				createInstanceHsmDestinationHaGroupOptionsModel.MemberSerials = []string{"123456", "123457"}
				createInstanceHsmDestinationHaGroupOptionsModel.Description = core.StringPtr("Primary HA Group for prod")
				createInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptionsModel)
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
	Describe(`GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptions *GetInstanceHsmDestinationHaGroupOptions) - Operation response error`, func() {
		getInstanceHsmDestinationHaGroupPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationHaGroup with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationHaGroupOptions model
				getInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions)
				getInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				getInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptions *GetInstanceHsmDestinationHaGroupOptions)`, func() {
		getInstanceHsmDestinationHaGroupPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationHaGroup successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceHsmDestinationHaGroupOptions model
				getInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions)
				getInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				getInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroupWithContext(ctx, getInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroupWithContext(ctx, getInstanceHsmDestinationHaGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationHaGroup successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationHaGroupOptions model
				getInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions)
				getInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				getInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceHsmDestinationHaGroup with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationHaGroupOptions model
				getInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions)
				getInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				getInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceHsmDestinationHaGroupOptions model with no property values
				getInstanceHsmDestinationHaGroupOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptionsModelNew)
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
			It(`Invoke GetInstanceHsmDestinationHaGroup successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationHaGroupOptions model
				getInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions)
				getInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				getInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptionsModel)
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
	Describe(`UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptions *UpdateInstanceHsmDestinationHaGroupOptions) - Operation response error`, func() {
		updateInstanceHsmDestinationHaGroupPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationHaGroup with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the HSMHAGroupPatch model
				hsmhaGroupPatchModel := new(secretsmanagerinstancemanagementv2.HSMHAGroupPatch)
				hsmhaGroupPatchModel.Name = core.StringPtr("dr-ha-group")
				hsmhaGroupPatchModel.Description = core.StringPtr("DR HA Group for failover")
				hsmhaGroupPatchModel.MemberSerials = []string{"123456", "123457", "123458"}
				hsmhaGroupPatchModelAsPatch, asPatchErr := hsmhaGroupPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationHaGroupOptions model
				updateInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions)
				updateInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				updateInstanceHsmDestinationHaGroupOptionsModel.HSMHAGroupPatch = hsmhaGroupPatchModelAsPatch
				updateInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptions *UpdateInstanceHsmDestinationHaGroupOptions)`, func() {
		updateInstanceHsmDestinationHaGroupPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationHaGroup successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the HSMHAGroupPatch model
				hsmhaGroupPatchModel := new(secretsmanagerinstancemanagementv2.HSMHAGroupPatch)
				hsmhaGroupPatchModel.Name = core.StringPtr("dr-ha-group")
				hsmhaGroupPatchModel.Description = core.StringPtr("DR HA Group for failover")
				hsmhaGroupPatchModel.MemberSerials = []string{"123456", "123457", "123458"}
				hsmhaGroupPatchModelAsPatch, asPatchErr := hsmhaGroupPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationHaGroupOptions model
				updateInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions)
				updateInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				updateInstanceHsmDestinationHaGroupOptionsModel.HSMHAGroupPatch = hsmhaGroupPatchModelAsPatch
				updateInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroupWithContext(ctx, updateInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroupWithContext(ctx, updateInstanceHsmDestinationHaGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ee52ebb6-1728-4580-8ede-13f6504e3ae0", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0", "name": "production-ha-group", "description": "Primary HA Group for prod", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T16:00:00.000Z", "group_label": "prod_ha_group", "group_serial": "11323465009866", "member_serials": ["MemberSerials"], "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T16:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationHaGroup successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HSMHAGroupPatch model
				hsmhaGroupPatchModel := new(secretsmanagerinstancemanagementv2.HSMHAGroupPatch)
				hsmhaGroupPatchModel.Name = core.StringPtr("dr-ha-group")
				hsmhaGroupPatchModel.Description = core.StringPtr("DR HA Group for failover")
				hsmhaGroupPatchModel.MemberSerials = []string{"123456", "123457", "123458"}
				hsmhaGroupPatchModelAsPatch, asPatchErr := hsmhaGroupPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationHaGroupOptions model
				updateInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions)
				updateInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				updateInstanceHsmDestinationHaGroupOptionsModel.HSMHAGroupPatch = hsmhaGroupPatchModelAsPatch
				updateInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceHsmDestinationHaGroup with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the HSMHAGroupPatch model
				hsmhaGroupPatchModel := new(secretsmanagerinstancemanagementv2.HSMHAGroupPatch)
				hsmhaGroupPatchModel.Name = core.StringPtr("dr-ha-group")
				hsmhaGroupPatchModel.Description = core.StringPtr("DR HA Group for failover")
				hsmhaGroupPatchModel.MemberSerials = []string{"123456", "123457", "123458"}
				hsmhaGroupPatchModelAsPatch, asPatchErr := hsmhaGroupPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationHaGroupOptions model
				updateInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions)
				updateInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				updateInstanceHsmDestinationHaGroupOptionsModel.HSMHAGroupPatch = hsmhaGroupPatchModelAsPatch
				updateInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceHsmDestinationHaGroupOptions model with no property values
				updateInstanceHsmDestinationHaGroupOptionsModelNew := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptionsModelNew)
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
			It(`Invoke UpdateInstanceHsmDestinationHaGroup successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the HSMHAGroupPatch model
				hsmhaGroupPatchModel := new(secretsmanagerinstancemanagementv2.HSMHAGroupPatch)
				hsmhaGroupPatchModel.Name = core.StringPtr("dr-ha-group")
				hsmhaGroupPatchModel.Description = core.StringPtr("DR HA Group for failover")
				hsmhaGroupPatchModel.MemberSerials = []string{"123456", "123457", "123458"}
				hsmhaGroupPatchModelAsPatch, asPatchErr := hsmhaGroupPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationHaGroupOptions model
				updateInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions)
				updateInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				updateInstanceHsmDestinationHaGroupOptionsModel.HSMHAGroupPatch = hsmhaGroupPatchModelAsPatch
				updateInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptionsModel)
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
	Describe(`DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptions *DeleteInstanceHsmDestinationHaGroupOptions)`, func() {
		deleteInstanceHsmDestinationHaGroupPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationHaGroupPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestinationHaGroup successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationHaGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationHaGroupOptions model
				deleteInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationHaGroupOptions)
				deleteInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				deleteInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteInstanceHsmDestinationHaGroup with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationHaGroupOptions model
				deleteInstanceHsmDestinationHaGroupOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationHaGroupOptions)
				deleteInstanceHsmDestinationHaGroupOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationHaGroupOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationHaGroupOptionsModel.HaGroupID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				deleteInstanceHsmDestinationHaGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteInstanceHsmDestinationHaGroupOptions model with no property values
				deleteInstanceHsmDestinationHaGroupOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationHaGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptions *ListInstanceHsmDestinationPathsOptions) - Operation response error`, func() {
		listInstanceHsmDestinationPathsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationPathsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationPaths with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathsOptions model
				listInstanceHsmDestinationPathsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions)
				listInstanceHsmDestinationPathsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptions *ListInstanceHsmDestinationPathsOptions)`, func() {
		listInstanceHsmDestinationPathsPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationPathsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"paths": [{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationPaths successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListInstanceHsmDestinationPathsOptions model
				listInstanceHsmDestinationPathsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions)
				listInstanceHsmDestinationPathsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathsWithContext(ctx, listInstanceHsmDestinationPathsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathsWithContext(ctx, listInstanceHsmDestinationPathsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationPathsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"paths": [{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationPaths successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathsOptions model
				listInstanceHsmDestinationPathsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions)
				listInstanceHsmDestinationPathsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListInstanceHsmDestinationPaths with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathsOptions model
				listInstanceHsmDestinationPathsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions)
				listInstanceHsmDestinationPathsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListInstanceHsmDestinationPathsOptions model with no property values
				listInstanceHsmDestinationPathsOptionsModelNew := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptionsModelNew)
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
			It(`Invoke ListInstanceHsmDestinationPaths successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathsOptions model
				listInstanceHsmDestinationPathsOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions)
				listInstanceHsmDestinationPathsOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathsOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptionsModel)
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
	Describe(`CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptions *CreateInstanceHsmDestinationPathOptions) - Operation response error`, func() {
		createInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationPath with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathOptions model
				createInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions)
				createInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathOptionsModel.Name = core.StringPtr("hsm-node-1")
				createInstanceHsmDestinationPathOptionsModel.Crn = core.StringPtr("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")
				createInstanceHsmDestinationPathOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptions *CreateInstanceHsmDestinationPathOptions)`, func() {
		createInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationPath successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateInstanceHsmDestinationPathOptions model
				createInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions)
				createInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathOptionsModel.Name = core.StringPtr("hsm-node-1")
				createInstanceHsmDestinationPathOptionsModel.Crn = core.StringPtr("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")
				createInstanceHsmDestinationPathOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathWithContext(ctx, createInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathWithContext(ctx, createInstanceHsmDestinationPathOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathOptions model
				createInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions)
				createInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathOptionsModel.Name = core.StringPtr("hsm-node-1")
				createInstanceHsmDestinationPathOptionsModel.Crn = core.StringPtr("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")
				createInstanceHsmDestinationPathOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateInstanceHsmDestinationPath with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathOptions model
				createInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions)
				createInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathOptionsModel.Name = core.StringPtr("hsm-node-1")
				createInstanceHsmDestinationPathOptionsModel.Crn = core.StringPtr("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")
				createInstanceHsmDestinationPathOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateInstanceHsmDestinationPathOptions model with no property values
				createInstanceHsmDestinationPathOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptionsModelNew)
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
			It(`Invoke CreateInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathOptions model
				createInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions)
				createInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathOptionsModel.Name = core.StringPtr("hsm-node-1")
				createInstanceHsmDestinationPathOptionsModel.Crn = core.StringPtr("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")
				createInstanceHsmDestinationPathOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptionsModel)
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
	Describe(`GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions *GetInstanceHsmDestinationPathOptions) - Operation response error`, func() {
		getInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationPath with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathOptions model
				getInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions)
				getInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions *GetInstanceHsmDestinationPathOptions)`, func() {
		getInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationPath successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceHsmDestinationPathOptions model
				getInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions)
				getInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathWithContext(ctx, getInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathWithContext(ctx, getInstanceHsmDestinationPathOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathOptions model
				getInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions)
				getInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceHsmDestinationPath with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathOptions model
				getInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions)
				getInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceHsmDestinationPathOptions model with no property values
				getInstanceHsmDestinationPathOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptionsModelNew)
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
			It(`Invoke GetInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathOptions model
				getInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions)
				getInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptionsModel)
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
	Describe(`UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptions *UpdateInstanceHsmDestinationPathOptions) - Operation response error`, func() {
		updateInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationPath with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the PathPatch model
				pathPatchModel := new(secretsmanagerinstancemanagementv2.PathPatch)
				pathPatchModel.Name = core.StringPtr("hsm-node-1-renamed")
				pathPatchModel.Description = core.StringPtr("testString")
				pathPatchModelAsPatch, asPatchErr := pathPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathOptions model
				updateInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions)
				updateInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathOptionsModel.PathPatch = pathPatchModelAsPatch
				updateInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptions *UpdateInstanceHsmDestinationPathOptions)`, func() {
		updateInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationPath successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the PathPatch model
				pathPatchModel := new(secretsmanagerinstancemanagementv2.PathPatch)
				pathPatchModel.Name = core.StringPtr("hsm-node-1-renamed")
				pathPatchModel.Description = core.StringPtr("testString")
				pathPatchModelAsPatch, asPatchErr := pathPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathOptions model
				updateInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions)
				updateInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathOptionsModel.PathPatch = pathPatchModelAsPatch
				updateInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathWithContext(ctx, updateInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathWithContext(ctx, updateInstanceHsmDestinationPathOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PathPatch model
				pathPatchModel := new(secretsmanagerinstancemanagementv2.PathPatch)
				pathPatchModel.Name = core.StringPtr("hsm-node-1-renamed")
				pathPatchModel.Description = core.StringPtr("testString")
				pathPatchModelAsPatch, asPatchErr := pathPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathOptions model
				updateInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions)
				updateInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathOptionsModel.PathPatch = pathPatchModelAsPatch
				updateInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceHsmDestinationPath with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the PathPatch model
				pathPatchModel := new(secretsmanagerinstancemanagementv2.PathPatch)
				pathPatchModel.Name = core.StringPtr("hsm-node-1-renamed")
				pathPatchModel.Description = core.StringPtr("testString")
				pathPatchModelAsPatch, asPatchErr := pathPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathOptions model
				updateInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions)
				updateInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathOptionsModel.PathPatch = pathPatchModelAsPatch
				updateInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceHsmDestinationPathOptions model with no property values
				updateInstanceHsmDestinationPathOptionsModelNew := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptionsModelNew)
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
			It(`Invoke UpdateInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the PathPatch model
				pathPatchModel := new(secretsmanagerinstancemanagementv2.PathPatch)
				pathPatchModel.Name = core.StringPtr("hsm-node-1-renamed")
				pathPatchModel.Description = core.StringPtr("testString")
				pathPatchModelAsPatch, asPatchErr := pathPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathOptions model
				updateInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions)
				updateInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathOptionsModel.PathPatch = pathPatchModelAsPatch
				updateInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptionsModel)
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
	Describe(`DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptions *DeleteInstanceHsmDestinationPathOptions) - Operation response error`, func() {
		deleteInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestinationPath with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationPathOptions model
				deleteInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions)
				deleteInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptions *DeleteInstanceHsmDestinationPathOptions)`, func() {
		deleteInstanceHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestinationPath successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the DeleteInstanceHsmDestinationPathOptions model
				deleteInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions)
				deleteInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathWithContext(ctx, deleteInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathWithContext(ctx, deleteInstanceHsmDestinationPathOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationPathPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001", "name": "hsm-node-1", "description": "Primary HSM node for prod", "crn": "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1", "provisioning_state": "completed", "health_state": "healthy", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:00:00.000Z", "updated_at": "2026-07-09T15:30:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationPathOptions model
				deleteInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions)
				deleteInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteInstanceHsmDestinationPath with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationPathOptions model
				deleteInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions)
				deleteInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteInstanceHsmDestinationPathOptions model with no property values
				deleteInstanceHsmDestinationPathOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationPathOptions model
				deleteInstanceHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions)
				deleteInstanceHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptionsModel)
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
	Describe(`CreatePathBuild(createPathBuildOptions *CreatePathBuildOptions) - Operation response error`, func() {
		createPathBuildPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/build"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPathBuildPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePathBuild with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreatePathBuildOptions model
				createPathBuildOptionsModel := new(secretsmanagerinstancemanagementv2.CreatePathBuildOptions)
				createPathBuildOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createPathBuildOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createPathBuildOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createPathBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePathBuild(createPathBuildOptions *CreatePathBuildOptions)`, func() {
		createPathBuildPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/build"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPathBuildPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "provisioning_state": "provisioning"}`)
				}))
			})
			It(`Invoke CreatePathBuild successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreatePathBuildOptions model
				createPathBuildOptionsModel := new(secretsmanagerinstancemanagementv2.CreatePathBuildOptions)
				createPathBuildOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createPathBuildOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createPathBuildOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createPathBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreatePathBuildWithContext(ctx, createPathBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreatePathBuildWithContext(ctx, createPathBuildOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createPathBuildPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "fa3d1234-0000-0000-0000-000000000001", "provisioning_state": "provisioning"}`)
				}))
			})
			It(`Invoke CreatePathBuild successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreatePathBuild(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreatePathBuildOptions model
				createPathBuildOptionsModel := new(secretsmanagerinstancemanagementv2.CreatePathBuildOptions)
				createPathBuildOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createPathBuildOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createPathBuildOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createPathBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreatePathBuild with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreatePathBuildOptions model
				createPathBuildOptionsModel := new(secretsmanagerinstancemanagementv2.CreatePathBuildOptions)
				createPathBuildOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createPathBuildOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createPathBuildOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createPathBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePathBuildOptions model with no property values
				createPathBuildOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreatePathBuildOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreatePathBuild successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreatePathBuildOptions model
				createPathBuildOptionsModel := new(secretsmanagerinstancemanagementv2.CreatePathBuildOptions)
				createPathBuildOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createPathBuildOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createPathBuildOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createPathBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptionsModel)
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
	Describe(`GetPathVpeStatus(getPathVpeStatusOptions *GetPathVpeStatusOptions) - Operation response error`, func() {
		getPathVpeStatusPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/health_check"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPathVpeStatusPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPathVpeStatus with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetPathVpeStatusOptions model
				getPathVpeStatusOptionsModel := new(secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions)
				getPathVpeStatusOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getPathVpeStatusOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getPathVpeStatusOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getPathVpeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPathVpeStatus(getPathVpeStatusOptions *GetPathVpeStatusOptions)`, func() {
		getPathVpeStatusPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/health_check"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPathVpeStatusPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"health_state": "healthy", "health_time_stamp": "2026-11-13T07:14:16.000Z", "provisioning_state": "completed"}`)
				}))
			})
			It(`Invoke GetPathVpeStatus successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetPathVpeStatusOptions model
				getPathVpeStatusOptionsModel := new(secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions)
				getPathVpeStatusOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getPathVpeStatusOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getPathVpeStatusOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getPathVpeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetPathVpeStatusWithContext(ctx, getPathVpeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetPathVpeStatusWithContext(ctx, getPathVpeStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPathVpeStatusPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"health_state": "healthy", "health_time_stamp": "2026-11-13T07:14:16.000Z", "provisioning_state": "completed"}`)
				}))
			})
			It(`Invoke GetPathVpeStatus successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetPathVpeStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPathVpeStatusOptions model
				getPathVpeStatusOptionsModel := new(secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions)
				getPathVpeStatusOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getPathVpeStatusOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getPathVpeStatusOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getPathVpeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPathVpeStatus with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetPathVpeStatusOptions model
				getPathVpeStatusOptionsModel := new(secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions)
				getPathVpeStatusOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getPathVpeStatusOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getPathVpeStatusOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getPathVpeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPathVpeStatusOptions model with no property values
				getPathVpeStatusOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptionsModelNew)
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
			It(`Invoke GetPathVpeStatus successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetPathVpeStatusOptions model
				getPathVpeStatusOptionsModel := new(secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions)
				getPathVpeStatusOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getPathVpeStatusOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getPathVpeStatusOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getPathVpeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptionsModel)
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
	Describe(`ResetHsmDestinationPath(resetHsmDestinationPathOptions *ResetHsmDestinationPathOptions) - Operation response error`, func() {
		resetHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/reset"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resetHsmDestinationPathPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ResetHsmDestinationPath with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ResetHsmDestinationPathOptions model
				resetHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.ResetHsmDestinationPathOptions)
				resetHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				resetHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				resetHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				resetHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResetHsmDestinationPath(resetHsmDestinationPathOptions *ResetHsmDestinationPathOptions)`, func() {
		resetHsmDestinationPathPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/reset"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resetHsmDestinationPathPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "provisioning_state": "completed"}`)
				}))
			})
			It(`Invoke ResetHsmDestinationPath successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ResetHsmDestinationPathOptions model
				resetHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.ResetHsmDestinationPathOptions)
				resetHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				resetHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				resetHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				resetHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ResetHsmDestinationPathWithContext(ctx, resetHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ResetHsmDestinationPathWithContext(ctx, resetHsmDestinationPathOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(resetHsmDestinationPathPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"href": "Href", "provisioning_state": "completed"}`)
				}))
			})
			It(`Invoke ResetHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ResetHsmDestinationPath(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResetHsmDestinationPathOptions model
				resetHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.ResetHsmDestinationPathOptions)
				resetHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				resetHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				resetHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				resetHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ResetHsmDestinationPath with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ResetHsmDestinationPathOptions model
				resetHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.ResetHsmDestinationPathOptions)
				resetHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				resetHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				resetHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				resetHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ResetHsmDestinationPathOptions model with no property values
				resetHsmDestinationPathOptionsModelNew := new(secretsmanagerinstancemanagementv2.ResetHsmDestinationPathOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptionsModelNew)
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
			It(`Invoke ResetHsmDestinationPath successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ResetHsmDestinationPathOptions model
				resetHsmDestinationPathOptionsModel := new(secretsmanagerinstancemanagementv2.ResetHsmDestinationPathOptions)
				resetHsmDestinationPathOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				resetHsmDestinationPathOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				resetHsmDestinationPathOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				resetHsmDestinationPathOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptionsModel)
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
	Describe(`ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptions *ListInstanceHsmDestinationPathServerCasOptions) - Operation response error`, func() {
		listInstanceHsmDestinationPathServerCasPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationPathServerCasPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationPathServerCas with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathServerCasOptions model
				listInstanceHsmDestinationPathServerCasOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions)
				listInstanceHsmDestinationPathServerCasOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathServerCasOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathServerCasOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				listInstanceHsmDestinationPathServerCasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptions *ListInstanceHsmDestinationPathServerCasOptions)`, func() {
		listInstanceHsmDestinationPathServerCasPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationPathServerCasPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"server_cas": [{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationPathServerCas successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListInstanceHsmDestinationPathServerCasOptions model
				listInstanceHsmDestinationPathServerCasOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions)
				listInstanceHsmDestinationPathServerCasOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathServerCasOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathServerCasOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				listInstanceHsmDestinationPathServerCasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCasWithContext(ctx, listInstanceHsmDestinationPathServerCasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCasWithContext(ctx, listInstanceHsmDestinationPathServerCasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationPathServerCasPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"server_cas": [{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationPathServerCas successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathServerCasOptions model
				listInstanceHsmDestinationPathServerCasOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions)
				listInstanceHsmDestinationPathServerCasOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathServerCasOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathServerCasOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				listInstanceHsmDestinationPathServerCasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListInstanceHsmDestinationPathServerCas with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathServerCasOptions model
				listInstanceHsmDestinationPathServerCasOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions)
				listInstanceHsmDestinationPathServerCasOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathServerCasOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathServerCasOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				listInstanceHsmDestinationPathServerCasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListInstanceHsmDestinationPathServerCasOptions model with no property values
				listInstanceHsmDestinationPathServerCasOptionsModelNew := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptionsModelNew)
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
			It(`Invoke ListInstanceHsmDestinationPathServerCas successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationPathServerCasOptions model
				listInstanceHsmDestinationPathServerCasOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions)
				listInstanceHsmDestinationPathServerCasOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathServerCasOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathServerCasOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				listInstanceHsmDestinationPathServerCasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptionsModel)
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
	Describe(`CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptions *CreateInstanceHsmDestinationPathServerCaOptions) - Operation response error`, func() {
		createInstanceHsmDestinationPathServerCaPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationPathServerCa with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathServerCaOptions model
				createInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions)
				createInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createInstanceHsmDestinationPathServerCaOptionsModel.Name = core.StringPtr("HSM1_Server_CA")
				createInstanceHsmDestinationPathServerCaOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
				createInstanceHsmDestinationPathServerCaOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptions *CreateInstanceHsmDestinationPathServerCaOptions)`, func() {
		createInstanceHsmDestinationPathServerCaPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationPathServerCa successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateInstanceHsmDestinationPathServerCaOptions model
				createInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions)
				createInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createInstanceHsmDestinationPathServerCaOptionsModel.Name = core.StringPtr("HSM1_Server_CA")
				createInstanceHsmDestinationPathServerCaOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
				createInstanceHsmDestinationPathServerCaOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCaWithContext(ctx, createInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCaWithContext(ctx, createInstanceHsmDestinationPathServerCaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationPathServerCa successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathServerCaOptions model
				createInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions)
				createInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createInstanceHsmDestinationPathServerCaOptionsModel.Name = core.StringPtr("HSM1_Server_CA")
				createInstanceHsmDestinationPathServerCaOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
				createInstanceHsmDestinationPathServerCaOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateInstanceHsmDestinationPathServerCa with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathServerCaOptions model
				createInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions)
				createInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createInstanceHsmDestinationPathServerCaOptionsModel.Name = core.StringPtr("HSM1_Server_CA")
				createInstanceHsmDestinationPathServerCaOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
				createInstanceHsmDestinationPathServerCaOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateInstanceHsmDestinationPathServerCaOptions model with no property values
				createInstanceHsmDestinationPathServerCaOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptionsModelNew)
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
			It(`Invoke CreateInstanceHsmDestinationPathServerCa successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationPathServerCaOptions model
				createInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions)
				createInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				createInstanceHsmDestinationPathServerCaOptionsModel.Name = core.StringPtr("HSM1_Server_CA")
				createInstanceHsmDestinationPathServerCaOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
				createInstanceHsmDestinationPathServerCaOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptionsModel)
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
	Describe(`GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptions *GetInstanceHsmDestinationPathServerCaOptions) - Operation response error`, func() {
		getInstanceHsmDestinationPathServerCaPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas/5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationPathServerCa with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathServerCaOptions model
				getInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions)
				getInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				getInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptions *GetInstanceHsmDestinationPathServerCaOptions)`, func() {
		getInstanceHsmDestinationPathServerCaPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas/5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationPathServerCa successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceHsmDestinationPathServerCaOptions model
				getInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions)
				getInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				getInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCaWithContext(ctx, getInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCaWithContext(ctx, getInstanceHsmDestinationPathServerCaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationPathServerCa successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathServerCaOptions model
				getInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions)
				getInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				getInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceHsmDestinationPathServerCa with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathServerCaOptions model
				getInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions)
				getInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				getInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceHsmDestinationPathServerCaOptions model with no property values
				getInstanceHsmDestinationPathServerCaOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptionsModelNew)
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
			It(`Invoke GetInstanceHsmDestinationPathServerCa successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationPathServerCaOptions model
				getInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions)
				getInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				getInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptionsModel)
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
	Describe(`UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptions *UpdateInstanceHsmDestinationPathServerCaOptions) - Operation response error`, func() {
		updateInstanceHsmDestinationPathServerCaPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas/5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationPathServerCa with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ServerCaPatch model
				serverCaPatchModel := new(secretsmanagerinstancemanagementv2.ServerCaPatch)
				serverCaPatchModel.Name = core.StringPtr("HSM1_Server_CA_renamed")
				serverCaPatchModel.Description = core.StringPtr("testString")
				serverCaPatchModel.Status = core.StringPtr("active")
				serverCaPatchModelAsPatch, asPatchErr := serverCaPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathServerCaOptions model
				updateInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions)
				updateInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaPatch = serverCaPatchModelAsPatch
				updateInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptions *UpdateInstanceHsmDestinationPathServerCaOptions)`, func() {
		updateInstanceHsmDestinationPathServerCaPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas/5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationPathServerCa successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ServerCaPatch model
				serverCaPatchModel := new(secretsmanagerinstancemanagementv2.ServerCaPatch)
				serverCaPatchModel.Name = core.StringPtr("HSM1_Server_CA_renamed")
				serverCaPatchModel.Description = core.StringPtr("testString")
				serverCaPatchModel.Status = core.StringPtr("active")
				serverCaPatchModelAsPatch, asPatchErr := serverCaPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathServerCaOptions model
				updateInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions)
				updateInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaPatch = serverCaPatchModelAsPatch
				updateInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCaWithContext(ctx, updateInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCaWithContext(ctx, updateInstanceHsmDestinationPathServerCaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f", "name": "HSM1_Server_CA", "description": "Primary CA for HSM node 1", "certificate": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T15:10:00.000Z", "status": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "created_at": "2026-07-09T15:10:00.000Z", "updated_at": "2026-07-09T15:20:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationPathServerCa successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ServerCaPatch model
				serverCaPatchModel := new(secretsmanagerinstancemanagementv2.ServerCaPatch)
				serverCaPatchModel.Name = core.StringPtr("HSM1_Server_CA_renamed")
				serverCaPatchModel.Description = core.StringPtr("testString")
				serverCaPatchModel.Status = core.StringPtr("active")
				serverCaPatchModelAsPatch, asPatchErr := serverCaPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathServerCaOptions model
				updateInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions)
				updateInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaPatch = serverCaPatchModelAsPatch
				updateInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceHsmDestinationPathServerCa with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ServerCaPatch model
				serverCaPatchModel := new(secretsmanagerinstancemanagementv2.ServerCaPatch)
				serverCaPatchModel.Name = core.StringPtr("HSM1_Server_CA_renamed")
				serverCaPatchModel.Description = core.StringPtr("testString")
				serverCaPatchModel.Status = core.StringPtr("active")
				serverCaPatchModelAsPatch, asPatchErr := serverCaPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathServerCaOptions model
				updateInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions)
				updateInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaPatch = serverCaPatchModelAsPatch
				updateInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceHsmDestinationPathServerCaOptions model with no property values
				updateInstanceHsmDestinationPathServerCaOptionsModelNew := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptionsModelNew)
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
			It(`Invoke UpdateInstanceHsmDestinationPathServerCa successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ServerCaPatch model
				serverCaPatchModel := new(secretsmanagerinstancemanagementv2.ServerCaPatch)
				serverCaPatchModel.Name = core.StringPtr("HSM1_Server_CA_renamed")
				serverCaPatchModel.Description = core.StringPtr("testString")
				serverCaPatchModel.Status = core.StringPtr("active")
				serverCaPatchModelAsPatch, asPatchErr := serverCaPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationPathServerCaOptions model
				updateInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions)
				updateInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaPatch = serverCaPatchModelAsPatch
				updateInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptionsModel)
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
	Describe(`DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptions *DeleteInstanceHsmDestinationPathServerCaOptions)`, func() {
		deleteInstanceHsmDestinationPathServerCaPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/paths/fa3d1234-0000-0000-0000-000000000001/server_cas/5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationPathServerCaPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestinationPathServerCa successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathServerCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationPathServerCaOptions model
				deleteInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathServerCaOptions)
				deleteInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteInstanceHsmDestinationPathServerCa with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationPathServerCaOptions model
				deleteInstanceHsmDestinationPathServerCaOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathServerCaOptions)
				deleteInstanceHsmDestinationPathServerCaOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.PathID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteInstanceHsmDestinationPathServerCaOptions model with no property values
				deleteInstanceHsmDestinationPathServerCaOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathServerCaOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptions *ListInstanceHsmDestinationClientCertificatesOptions) - Operation response error`, func() {
		listInstanceHsmDestinationClientCertificatesPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationClientCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationClientCertificates with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationClientCertificatesOptions model
				listInstanceHsmDestinationClientCertificatesOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions)
				listInstanceHsmDestinationClientCertificatesOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationClientCertificatesOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationClientCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptions *ListInstanceHsmDestinationClientCertificatesOptions)`, func() {
		listInstanceHsmDestinationClientCertificatesPath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationClientCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"client_certificates": [{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationClientCertificates successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListInstanceHsmDestinationClientCertificatesOptions model
				listInstanceHsmDestinationClientCertificatesOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions)
				listInstanceHsmDestinationClientCertificatesOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationClientCertificatesOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationClientCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificatesWithContext(ctx, listInstanceHsmDestinationClientCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificatesWithContext(ctx, listInstanceHsmDestinationClientCertificatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceHsmDestinationClientCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"client_certificates": [{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}]}`)
				}))
			})
			It(`Invoke ListInstanceHsmDestinationClientCertificates successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationClientCertificatesOptions model
				listInstanceHsmDestinationClientCertificatesOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions)
				listInstanceHsmDestinationClientCertificatesOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationClientCertificatesOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationClientCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListInstanceHsmDestinationClientCertificates with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationClientCertificatesOptions model
				listInstanceHsmDestinationClientCertificatesOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions)
				listInstanceHsmDestinationClientCertificatesOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationClientCertificatesOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationClientCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListInstanceHsmDestinationClientCertificatesOptions model with no property values
				listInstanceHsmDestinationClientCertificatesOptionsModelNew := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptionsModelNew)
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
			It(`Invoke ListInstanceHsmDestinationClientCertificates successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ListInstanceHsmDestinationClientCertificatesOptions model
				listInstanceHsmDestinationClientCertificatesOptionsModel := new(secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions)
				listInstanceHsmDestinationClientCertificatesOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationClientCertificatesOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationClientCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptionsModel)
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
	Describe(`CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptions *CreateInstanceHsmDestinationClientCertificateOptions) - Operation response error`, func() {
		createInstanceHsmDestinationClientCertificatePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationClientCertificate with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationClientCertificateOptions model
				createInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions)
				createInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationClientCertificateOptionsModel.Name = core.StringPtr("hsm-client-cert-prod")
				createInstanceHsmDestinationClientCertificateOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptions *CreateInstanceHsmDestinationClientCertificateOptions)`, func() {
		createInstanceHsmDestinationClientCertificatePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationClientCertificate successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateInstanceHsmDestinationClientCertificateOptions model
				createInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions)
				createInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationClientCertificateOptionsModel.Name = core.StringPtr("hsm-client-cert-prod")
				createInstanceHsmDestinationClientCertificateOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificateWithContext(ctx, createInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificateWithContext(ctx, createInstanceHsmDestinationClientCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke CreateInstanceHsmDestinationClientCertificate successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationClientCertificateOptions model
				createInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions)
				createInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationClientCertificateOptionsModel.Name = core.StringPtr("hsm-client-cert-prod")
				createInstanceHsmDestinationClientCertificateOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateInstanceHsmDestinationClientCertificate with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationClientCertificateOptions model
				createInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions)
				createInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationClientCertificateOptionsModel.Name = core.StringPtr("hsm-client-cert-prod")
				createInstanceHsmDestinationClientCertificateOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateInstanceHsmDestinationClientCertificateOptions model with no property values
				createInstanceHsmDestinationClientCertificateOptionsModelNew := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptionsModelNew)
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
			It(`Invoke CreateInstanceHsmDestinationClientCertificate successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the CreateInstanceHsmDestinationClientCertificateOptions model
				createInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions)
				createInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationClientCertificateOptionsModel.Name = core.StringPtr("hsm-client-cert-prod")
				createInstanceHsmDestinationClientCertificateOptionsModel.Description = core.StringPtr("testString")
				createInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptionsModel)
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
	Describe(`GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptions *GetInstanceHsmDestinationClientCertificateOptions) - Operation response error`, func() {
		getInstanceHsmDestinationClientCertificatePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationClientCertificate with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationClientCertificateOptions model
				getInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions)
				getInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				getInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptions *GetInstanceHsmDestinationClientCertificateOptions)`, func() {
		getInstanceHsmDestinationClientCertificatePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationClientCertificate successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceHsmDestinationClientCertificateOptions model
				getInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions)
				getInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				getInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificateWithContext(ctx, getInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificateWithContext(ctx, getInstanceHsmDestinationClientCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke GetInstanceHsmDestinationClientCertificate successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationClientCertificateOptions model
				getInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions)
				getInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				getInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceHsmDestinationClientCertificate with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationClientCertificateOptions model
				getInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions)
				getInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				getInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceHsmDestinationClientCertificateOptions model with no property values
				getInstanceHsmDestinationClientCertificateOptionsModelNew := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptionsModelNew)
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
			It(`Invoke GetInstanceHsmDestinationClientCertificate successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the GetInstanceHsmDestinationClientCertificateOptions model
				getInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions)
				getInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				getInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptionsModel)
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
	Describe(`UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptions *UpdateInstanceHsmDestinationClientCertificateOptions) - Operation response error`, func() {
		updateInstanceHsmDestinationClientCertificatePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationClientCertificate with error: Operation response processing error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ClientCertificatePatch model
				clientCertificatePatchModel := new(secretsmanagerinstancemanagementv2.ClientCertificatePatch)
				clientCertificatePatchModel.Name = core.StringPtr("hsm-client-cert-renamed")
				clientCertificatePatchModel.Description = core.StringPtr("testString")
				clientCertificatePatchModel.Status = core.StringPtr("testString")
				clientCertificatePatchModel.SignedCert = core.StringPtr("testString")
				clientCertificatePatchModelAsPatch, asPatchErr := clientCertificatePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationClientCertificateOptions model
				updateInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions)
				updateInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				updateInstanceHsmDestinationClientCertificateOptionsModel.ClientCertificatePatch = clientCertificatePatchModelAsPatch
				updateInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				secretsManagerInstanceManagementService.EnableRetries(0, 0)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptions *UpdateInstanceHsmDestinationClientCertificateOptions)`, func() {
		updateInstanceHsmDestinationClientCertificatePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationClientCertificate successfully with retries`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
				secretsManagerInstanceManagementService.EnableRetries(0, 0)

				// Construct an instance of the ClientCertificatePatch model
				clientCertificatePatchModel := new(secretsmanagerinstancemanagementv2.ClientCertificatePatch)
				clientCertificatePatchModel.Name = core.StringPtr("hsm-client-cert-renamed")
				clientCertificatePatchModel.Description = core.StringPtr("testString")
				clientCertificatePatchModel.Status = core.StringPtr("testString")
				clientCertificatePatchModel.SignedCert = core.StringPtr("testString")
				clientCertificatePatchModelAsPatch, asPatchErr := clientCertificatePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationClientCertificateOptions model
				updateInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions)
				updateInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				updateInstanceHsmDestinationClientCertificateOptionsModel.ClientCertificatePatch = clientCertificatePatchModelAsPatch
				updateInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificateWithContext(ctx, updateInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				secretsManagerInstanceManagementService.DisableRetries()
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificateWithContext(ctx, updateInstanceHsmDestinationClientCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7c9e6679-7425-40de-944b-e07fc1f90ae7", "href": "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7", "name": "hsm-client-cert-prod", "description": "Primary client cert for prod", "status": "Status", "provisioning_state": "active", "health_state": "valid", "health_as_of": "2026-07-09T15:30:00.000Z", "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICvDCCAaQCAQAwdzEL...\n-----END CERTIFICATE REQUEST-----", "csr_cn": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "signed_cert": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----", "expires_at": "2028-07-09T10:00:00.000Z", "created_at": "2026-07-09T10:00:00.000Z", "updated_at": "2026-07-09T11:00:00.000Z", "created_by": "iam-ServiceId-abc123", "updated_by": "iam-ServiceId-abc123"}`)
				}))
			})
			It(`Invoke UpdateInstanceHsmDestinationClientCertificate successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClientCertificatePatch model
				clientCertificatePatchModel := new(secretsmanagerinstancemanagementv2.ClientCertificatePatch)
				clientCertificatePatchModel.Name = core.StringPtr("hsm-client-cert-renamed")
				clientCertificatePatchModel.Description = core.StringPtr("testString")
				clientCertificatePatchModel.Status = core.StringPtr("testString")
				clientCertificatePatchModel.SignedCert = core.StringPtr("testString")
				clientCertificatePatchModelAsPatch, asPatchErr := clientCertificatePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationClientCertificateOptions model
				updateInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions)
				updateInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				updateInstanceHsmDestinationClientCertificateOptionsModel.ClientCertificatePatch = clientCertificatePatchModelAsPatch
				updateInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceHsmDestinationClientCertificate with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ClientCertificatePatch model
				clientCertificatePatchModel := new(secretsmanagerinstancemanagementv2.ClientCertificatePatch)
				clientCertificatePatchModel.Name = core.StringPtr("hsm-client-cert-renamed")
				clientCertificatePatchModel.Description = core.StringPtr("testString")
				clientCertificatePatchModel.Status = core.StringPtr("testString")
				clientCertificatePatchModel.SignedCert = core.StringPtr("testString")
				clientCertificatePatchModelAsPatch, asPatchErr := clientCertificatePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationClientCertificateOptions model
				updateInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions)
				updateInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				updateInstanceHsmDestinationClientCertificateOptionsModel.ClientCertificatePatch = clientCertificatePatchModelAsPatch
				updateInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceHsmDestinationClientCertificateOptions model with no property values
				updateInstanceHsmDestinationClientCertificateOptionsModelNew := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptionsModelNew)
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
			It(`Invoke UpdateInstanceHsmDestinationClientCertificate successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the ClientCertificatePatch model
				clientCertificatePatchModel := new(secretsmanagerinstancemanagementv2.ClientCertificatePatch)
				clientCertificatePatchModel.Name = core.StringPtr("hsm-client-cert-renamed")
				clientCertificatePatchModel.Description = core.StringPtr("testString")
				clientCertificatePatchModel.Status = core.StringPtr("testString")
				clientCertificatePatchModel.SignedCert = core.StringPtr("testString")
				clientCertificatePatchModelAsPatch, asPatchErr := clientCertificatePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateInstanceHsmDestinationClientCertificateOptions model
				updateInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions)
				updateInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				updateInstanceHsmDestinationClientCertificateOptionsModel.ClientCertificatePatch = clientCertificatePatchModelAsPatch
				updateInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptionsModel)
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
	Describe(`DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptions *DeleteInstanceHsmDestinationClientCertificateOptions)`, func() {
		deleteInstanceHsmDestinationClientCertificatePath := "/v2/instances/fe5c5c27-3e44-474a-9484-55892e12d6d1/hsm_destinations/b2c3d4e5-f6a7-8901-bcde-f12345678901/client_certificates/7c9e6679-7425-40de-944b-e07fc1f90ae7"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceHsmDestinationClientCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteInstanceHsmDestinationClientCertificate successfully`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationClientCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationClientCertificateOptions model
				deleteInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationClientCertificateOptions)
				deleteInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteInstanceHsmDestinationClientCertificate with error: Operation validation and request error`, func() {
				secretsManagerInstanceManagementService, serviceErr := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(secretsManagerInstanceManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteInstanceHsmDestinationClientCertificateOptions model
				deleteInstanceHsmDestinationClientCertificateOptionsModel := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationClientCertificateOptions)
				deleteInstanceHsmDestinationClientCertificateOptionsModel.InstanceID = core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.DestinationID = core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.CertID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := secretsManagerInstanceManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteInstanceHsmDestinationClientCertificateOptions model with no property values
				deleteInstanceHsmDestinationClientCertificateOptionsModelNew := new(secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationClientCertificateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
			It(`Invoke NewApplyHsmDestinationOptions successfully`, func() {
				// Construct an instance of the ApplyHsmDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				applyHsmDestinationOptionsModel := secretsManagerInstanceManagementService.NewApplyHsmDestinationOptions(instanceID, destinationID)
				applyHsmDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				applyHsmDestinationOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				applyHsmDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(applyHsmDestinationOptionsModel).ToNot(BeNil())
				Expect(applyHsmDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(applyHsmDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(applyHsmDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateInstanceDestinationOptions successfully`, func() {
				// Construct an instance of the CreateDestinationRequestIbmCloudDatabaseDestinationPrototype model
				createDestinationRequestModel := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
				Expect(createDestinationRequestModel).ToNot(BeNil())
				createDestinationRequestModel.Name = core.StringPtr("testString")
				createDestinationRequestModel.Type = core.StringPtr("ibm_cloud_database")
				createDestinationRequestModel.Description = core.StringPtr("Production PostgreSQL database")
				createDestinationRequestModel.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")
				Expect(createDestinationRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationRequestModel.Type).To(Equal(core.StringPtr("ibm_cloud_database")))
				Expect(createDestinationRequestModel.Description).To(Equal(core.StringPtr("Production PostgreSQL database")))
				Expect(createDestinationRequestModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")))

				// Construct an instance of the CreateInstanceDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				var destinationPrototype secretsmanagerinstancemanagementv2.CreateDestinationRequestIntf = nil
				createInstanceDestinationOptionsModel := secretsManagerInstanceManagementService.NewCreateInstanceDestinationOptions(instanceID, destinationPrototype)
				createInstanceDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceDestinationOptionsModel.SetDestinationPrototype(createDestinationRequestModel)
				createInstanceDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createInstanceDestinationOptionsModel).ToNot(BeNil())
				Expect(createInstanceDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createInstanceDestinationOptionsModel.DestinationPrototype).To(Equal(createDestinationRequestModel))
				Expect(createInstanceDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateInstanceHsmDestinationClientCertificateOptions successfully`, func() {
				// Construct an instance of the CreateInstanceHsmDestinationClientCertificateOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				createInstanceHsmDestinationClientCertificateOptionsName := "hsm-client-cert-prod"
				createInstanceHsmDestinationClientCertificateOptionsModel := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationClientCertificateOptions(instanceID, destinationID, createInstanceHsmDestinationClientCertificateOptionsName)
				createInstanceHsmDestinationClientCertificateOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationClientCertificateOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationClientCertificateOptionsModel.SetName("hsm-client-cert-prod")
				createInstanceHsmDestinationClientCertificateOptionsModel.SetDescription("testString")
				createInstanceHsmDestinationClientCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createInstanceHsmDestinationClientCertificateOptionsModel).ToNot(BeNil())
				Expect(createInstanceHsmDestinationClientCertificateOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createInstanceHsmDestinationClientCertificateOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(createInstanceHsmDestinationClientCertificateOptionsModel.Name).To(Equal(core.StringPtr("hsm-client-cert-prod")))
				Expect(createInstanceHsmDestinationClientCertificateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createInstanceHsmDestinationClientCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateInstanceHsmDestinationHaGroupOptions successfully`, func() {
				// Construct an instance of the CreateInstanceHsmDestinationHaGroupOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				createInstanceHsmDestinationHaGroupOptionsName := "production-ha-group"
				createInstanceHsmDestinationHaGroupOptionsGroupLabel := "prod_ha_group"
				createInstanceHsmDestinationHaGroupOptionsGroupSerial := "11323465009866"
				createInstanceHsmDestinationHaGroupOptionsMemberSerials := []string{"123456", "123457"}
				createInstanceHsmDestinationHaGroupOptionsModel := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationHaGroupOptions(instanceID, destinationID, createInstanceHsmDestinationHaGroupOptionsName, createInstanceHsmDestinationHaGroupOptionsGroupLabel, createInstanceHsmDestinationHaGroupOptionsGroupSerial, createInstanceHsmDestinationHaGroupOptionsMemberSerials)
				createInstanceHsmDestinationHaGroupOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationHaGroupOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationHaGroupOptionsModel.SetName("production-ha-group")
				createInstanceHsmDestinationHaGroupOptionsModel.SetGroupLabel("prod_ha_group")
				createInstanceHsmDestinationHaGroupOptionsModel.SetGroupSerial("11323465009866")
				createInstanceHsmDestinationHaGroupOptionsModel.SetMemberSerials([]string{"123456", "123457"})
				createInstanceHsmDestinationHaGroupOptionsModel.SetDescription("Primary HA Group for prod")
				createInstanceHsmDestinationHaGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createInstanceHsmDestinationHaGroupOptionsModel).ToNot(BeNil())
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.Name).To(Equal(core.StringPtr("production-ha-group")))
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.GroupLabel).To(Equal(core.StringPtr("prod_ha_group")))
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.GroupSerial).To(Equal(core.StringPtr("11323465009866")))
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.MemberSerials).To(Equal([]string{"123456", "123457"}))
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.Description).To(Equal(core.StringPtr("Primary HA Group for prod")))
				Expect(createInstanceHsmDestinationHaGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateInstanceHsmDestinationOptions successfully`, func() {
				// Construct an instance of the CreateInstanceHsmDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				createInstanceHsmDestinationOptionsName := "prod-hsm-destination"
				createInstanceHsmDestinationOptionsVendor := "thales_luna"
				createInstanceHsmDestinationOptionsModel := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationOptions(instanceID, createInstanceHsmDestinationOptionsName, createInstanceHsmDestinationOptionsVendor)
				createInstanceHsmDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationOptionsModel.SetName("prod-hsm-destination")
				createInstanceHsmDestinationOptionsModel.SetVendor("thales_luna")
				createInstanceHsmDestinationOptionsModel.SetDescription("Primary Thales Luna destination")
				createInstanceHsmDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createInstanceHsmDestinationOptionsModel).ToNot(BeNil())
				Expect(createInstanceHsmDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createInstanceHsmDestinationOptionsModel.Name).To(Equal(core.StringPtr("prod-hsm-destination")))
				Expect(createInstanceHsmDestinationOptionsModel.Vendor).To(Equal(core.StringPtr("thales_luna")))
				Expect(createInstanceHsmDestinationOptionsModel.Description).To(Equal(core.StringPtr("Primary Thales Luna destination")))
				Expect(createInstanceHsmDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateInstanceHsmDestinationPathOptions successfully`, func() {
				// Construct an instance of the CreateInstanceHsmDestinationPathOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				createInstanceHsmDestinationPathOptionsName := "hsm-node-1"
				createInstanceHsmDestinationPathOptionsCrn := "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1"
				createInstanceHsmDestinationPathOptionsModel := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationPathOptions(instanceID, destinationID, createInstanceHsmDestinationPathOptionsName, createInstanceHsmDestinationPathOptionsCrn)
				createInstanceHsmDestinationPathOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathOptionsModel.SetName("hsm-node-1")
				createInstanceHsmDestinationPathOptionsModel.SetCrn("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")
				createInstanceHsmDestinationPathOptionsModel.SetDescription("testString")
				createInstanceHsmDestinationPathOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createInstanceHsmDestinationPathOptionsModel).ToNot(BeNil())
				Expect(createInstanceHsmDestinationPathOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createInstanceHsmDestinationPathOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(createInstanceHsmDestinationPathOptionsModel.Name).To(Equal(core.StringPtr("hsm-node-1")))
				Expect(createInstanceHsmDestinationPathOptionsModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")))
				Expect(createInstanceHsmDestinationPathOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createInstanceHsmDestinationPathOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateInstanceHsmDestinationPathServerCaOptions successfully`, func() {
				// Construct an instance of the CreateInstanceHsmDestinationPathServerCaOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				createInstanceHsmDestinationPathServerCaOptionsName := "HSM1_Server_CA"
				createInstanceHsmDestinationPathServerCaOptionsCertificate := "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----"
				createInstanceHsmDestinationPathServerCaOptionsModel := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationPathServerCaOptions(instanceID, destinationID, pathID, createInstanceHsmDestinationPathServerCaOptionsName, createInstanceHsmDestinationPathServerCaOptionsCertificate)
				createInstanceHsmDestinationPathServerCaOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createInstanceHsmDestinationPathServerCaOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createInstanceHsmDestinationPathServerCaOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				createInstanceHsmDestinationPathServerCaOptionsModel.SetName("HSM1_Server_CA")
				createInstanceHsmDestinationPathServerCaOptionsModel.SetCertificate("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
				createInstanceHsmDestinationPathServerCaOptionsModel.SetDescription("testString")
				createInstanceHsmDestinationPathServerCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel).ToNot(BeNil())
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel.Name).To(Equal(core.StringPtr("HSM1_Server_CA")))
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel.Certificate).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")))
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createInstanceHsmDestinationPathServerCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePathBuildOptions successfully`, func() {
				// Construct an instance of the CreatePathBuildOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				createPathBuildOptionsModel := secretsManagerInstanceManagementService.NewCreatePathBuildOptions(instanceID, destinationID, pathID)
				createPathBuildOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createPathBuildOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				createPathBuildOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				createPathBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPathBuildOptionsModel).ToNot(BeNil())
				Expect(createPathBuildOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createPathBuildOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(createPathBuildOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(createPathBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVaultAdmintokenOptions successfully`, func() {
				// Construct an instance of the CreateVaultAdmintokenOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				createVaultAdmintokenOptionsModel := secretsManagerInstanceManagementService.NewCreateVaultAdmintokenOptions(instanceID)
				createVaultAdmintokenOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				createVaultAdmintokenOptionsModel.SetResponseWrapping(true)
				createVaultAdmintokenOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVaultAdmintokenOptionsModel).ToNot(BeNil())
				Expect(createVaultAdmintokenOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(createVaultAdmintokenOptionsModel.ResponseWrapping).To(Equal(core.BoolPtr(true)))
				Expect(createVaultAdmintokenOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteInstanceDestinationOptions successfully`, func() {
				// Construct an instance of the DeleteInstanceDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				deleteInstanceDestinationOptionsModel := secretsManagerInstanceManagementService.NewDeleteInstanceDestinationOptions(instanceID, destinationID)
				deleteInstanceDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceDestinationOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteInstanceDestinationOptionsModel).ToNot(BeNil())
				Expect(deleteInstanceDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(deleteInstanceDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(deleteInstanceDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteInstanceHsmDestinationClientCertificateOptions successfully`, func() {
				// Construct an instance of the DeleteInstanceHsmDestinationClientCertificateOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				certID := "7c9e6679-7425-40de-944b-e07fc1f90ae7"
				deleteInstanceHsmDestinationClientCertificateOptionsModel := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationClientCertificateOptions(instanceID, destinationID, certID)
				deleteInstanceHsmDestinationClientCertificateOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.SetCertID("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				deleteInstanceHsmDestinationClientCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteInstanceHsmDestinationClientCertificateOptionsModel).ToNot(BeNil())
				Expect(deleteInstanceHsmDestinationClientCertificateOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(deleteInstanceHsmDestinationClientCertificateOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(deleteInstanceHsmDestinationClientCertificateOptionsModel.CertID).To(Equal(core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")))
				Expect(deleteInstanceHsmDestinationClientCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteInstanceHsmDestinationHaGroupOptions successfully`, func() {
				// Construct an instance of the DeleteInstanceHsmDestinationHaGroupOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				haGroupID := "ee52ebb6-1728-4580-8ede-13f6504e3ae0"
				deleteInstanceHsmDestinationHaGroupOptionsModel := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationHaGroupOptions(instanceID, destinationID, haGroupID)
				deleteInstanceHsmDestinationHaGroupOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationHaGroupOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationHaGroupOptionsModel.SetHaGroupID("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				deleteInstanceHsmDestinationHaGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteInstanceHsmDestinationHaGroupOptionsModel).ToNot(BeNil())
				Expect(deleteInstanceHsmDestinationHaGroupOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(deleteInstanceHsmDestinationHaGroupOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(deleteInstanceHsmDestinationHaGroupOptionsModel.HaGroupID).To(Equal(core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")))
				Expect(deleteInstanceHsmDestinationHaGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteInstanceHsmDestinationOptions successfully`, func() {
				// Construct an instance of the DeleteInstanceHsmDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				deleteInstanceHsmDestinationOptionsModel := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationOptions(instanceID, destinationID)
				deleteInstanceHsmDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteInstanceHsmDestinationOptionsModel).ToNot(BeNil())
				Expect(deleteInstanceHsmDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(deleteInstanceHsmDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(deleteInstanceHsmDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteInstanceHsmDestinationPathOptions successfully`, func() {
				// Construct an instance of the DeleteInstanceHsmDestinationPathOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				deleteInstanceHsmDestinationPathOptionsModel := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationPathOptions(instanceID, destinationID, pathID)
				deleteInstanceHsmDestinationPathOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteInstanceHsmDestinationPathOptionsModel).ToNot(BeNil())
				Expect(deleteInstanceHsmDestinationPathOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(deleteInstanceHsmDestinationPathOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(deleteInstanceHsmDestinationPathOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(deleteInstanceHsmDestinationPathOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteInstanceHsmDestinationPathServerCaOptions successfully`, func() {
				// Construct an instance of the DeleteInstanceHsmDestinationPathServerCaOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				serverCaID := "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
				deleteInstanceHsmDestinationPathServerCaOptionsModel := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationPathServerCaOptions(instanceID, destinationID, pathID, serverCaID)
				deleteInstanceHsmDestinationPathServerCaOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.SetServerCaID("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				deleteInstanceHsmDestinationPathServerCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteInstanceHsmDestinationPathServerCaOptionsModel).ToNot(BeNil())
				Expect(deleteInstanceHsmDestinationPathServerCaOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(deleteInstanceHsmDestinationPathServerCaOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(deleteInstanceHsmDestinationPathServerCaOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(deleteInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID).To(Equal(core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")))
				Expect(deleteInstanceHsmDestinationPathServerCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVaultAdmintokensOptions successfully`, func() {
				// Construct an instance of the DeleteVaultAdmintokensOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				deleteVaultAdmintokensOptionsModel := secretsManagerInstanceManagementService.NewDeleteVaultAdmintokensOptions(instanceID)
				deleteVaultAdmintokensOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				deleteVaultAdmintokensOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVaultAdmintokensOptionsModel).ToNot(BeNil())
				Expect(deleteVaultAdmintokensOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(deleteVaultAdmintokensOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHsmDestinationHealthSnapshotOptions successfully`, func() {
				// Construct an instance of the GetHsmDestinationHealthSnapshotOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				getHsmDestinationHealthSnapshotOptionsModel := secretsManagerInstanceManagementService.NewGetHsmDestinationHealthSnapshotOptions(instanceID, destinationID)
				getHsmDestinationHealthSnapshotOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getHsmDestinationHealthSnapshotOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getHsmDestinationHealthSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHsmDestinationHealthSnapshotOptionsModel).ToNot(BeNil())
				Expect(getHsmDestinationHealthSnapshotOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getHsmDestinationHealthSnapshotOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getHsmDestinationHealthSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceDestinationOptions successfully`, func() {
				// Construct an instance of the GetInstanceDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				getInstanceDestinationOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceDestinationOptions(instanceID, destinationID)
				getInstanceDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceDestinationOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceDestinationOptionsModel).ToNot(BeNil())
				Expect(getInstanceDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getInstanceDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getInstanceDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceHsmDestinationClientCertificateOptions successfully`, func() {
				// Construct an instance of the GetInstanceHsmDestinationClientCertificateOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				certID := "7c9e6679-7425-40de-944b-e07fc1f90ae7"
				getInstanceHsmDestinationClientCertificateOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationClientCertificateOptions(instanceID, destinationID, certID)
				getInstanceHsmDestinationClientCertificateOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationClientCertificateOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationClientCertificateOptionsModel.SetCertID("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				getInstanceHsmDestinationClientCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceHsmDestinationClientCertificateOptionsModel).ToNot(BeNil())
				Expect(getInstanceHsmDestinationClientCertificateOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getInstanceHsmDestinationClientCertificateOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getInstanceHsmDestinationClientCertificateOptionsModel.CertID).To(Equal(core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")))
				Expect(getInstanceHsmDestinationClientCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceHsmDestinationHaGroupOptions successfully`, func() {
				// Construct an instance of the GetInstanceHsmDestinationHaGroupOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				haGroupID := "ee52ebb6-1728-4580-8ede-13f6504e3ae0"
				getInstanceHsmDestinationHaGroupOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationHaGroupOptions(instanceID, destinationID, haGroupID)
				getInstanceHsmDestinationHaGroupOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationHaGroupOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationHaGroupOptionsModel.SetHaGroupID("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				getInstanceHsmDestinationHaGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceHsmDestinationHaGroupOptionsModel).ToNot(BeNil())
				Expect(getInstanceHsmDestinationHaGroupOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getInstanceHsmDestinationHaGroupOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getInstanceHsmDestinationHaGroupOptionsModel.HaGroupID).To(Equal(core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")))
				Expect(getInstanceHsmDestinationHaGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceHsmDestinationOptions successfully`, func() {
				// Construct an instance of the GetInstanceHsmDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				getInstanceHsmDestinationOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationOptions(instanceID, destinationID)
				getInstanceHsmDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceHsmDestinationOptionsModel).ToNot(BeNil())
				Expect(getInstanceHsmDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getInstanceHsmDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getInstanceHsmDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceHsmDestinationPathOptions successfully`, func() {
				// Construct an instance of the GetInstanceHsmDestinationPathOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				getInstanceHsmDestinationPathOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationPathOptions(instanceID, destinationID, pathID)
				getInstanceHsmDestinationPathOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceHsmDestinationPathOptionsModel).ToNot(BeNil())
				Expect(getInstanceHsmDestinationPathOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getInstanceHsmDestinationPathOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getInstanceHsmDestinationPathOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(getInstanceHsmDestinationPathOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceHsmDestinationPathServerCaOptions successfully`, func() {
				// Construct an instance of the GetInstanceHsmDestinationPathServerCaOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				serverCaID := "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
				getInstanceHsmDestinationPathServerCaOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationPathServerCaOptions(instanceID, destinationID, pathID, serverCaID)
				getInstanceHsmDestinationPathServerCaOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceHsmDestinationPathServerCaOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getInstanceHsmDestinationPathServerCaOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				getInstanceHsmDestinationPathServerCaOptionsModel.SetServerCaID("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				getInstanceHsmDestinationPathServerCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceHsmDestinationPathServerCaOptionsModel).ToNot(BeNil())
				Expect(getInstanceHsmDestinationPathServerCaOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getInstanceHsmDestinationPathServerCaOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getInstanceHsmDestinationPathServerCaOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(getInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID).To(Equal(core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")))
				Expect(getInstanceHsmDestinationPathServerCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceOptions successfully`, func() {
				// Construct an instance of the GetInstanceOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				getInstanceOptionsModel := secretsManagerInstanceManagementService.NewGetInstanceOptions(instanceID)
				getInstanceOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceOptionsModel).ToNot(BeNil())
				Expect(getInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPathVpeStatusOptions successfully`, func() {
				// Construct an instance of the GetPathVpeStatusOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				getPathVpeStatusOptionsModel := secretsManagerInstanceManagementService.NewGetPathVpeStatusOptions(instanceID, destinationID, pathID)
				getPathVpeStatusOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				getPathVpeStatusOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				getPathVpeStatusOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				getPathVpeStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPathVpeStatusOptionsModel).ToNot(BeNil())
				Expect(getPathVpeStatusOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(getPathVpeStatusOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(getPathVpeStatusOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(getPathVpeStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceDestinationsOptions successfully`, func() {
				// Construct an instance of the ListInstanceDestinationsOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				listInstanceDestinationsOptionsModel := secretsManagerInstanceManagementService.NewListInstanceDestinationsOptions(instanceID)
				listInstanceDestinationsOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceDestinationsOptionsModel.SetState("not_started")
				listInstanceDestinationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceDestinationsOptionsModel).ToNot(BeNil())
				Expect(listInstanceDestinationsOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(listInstanceDestinationsOptionsModel.State).To(Equal(core.StringPtr("not_started")))
				Expect(listInstanceDestinationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceHsmDestinationClientCertificatesOptions successfully`, func() {
				// Construct an instance of the ListInstanceHsmDestinationClientCertificatesOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				listInstanceHsmDestinationClientCertificatesOptionsModel := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationClientCertificatesOptions(instanceID, destinationID)
				listInstanceHsmDestinationClientCertificatesOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationClientCertificatesOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationClientCertificatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceHsmDestinationClientCertificatesOptionsModel).ToNot(BeNil())
				Expect(listInstanceHsmDestinationClientCertificatesOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(listInstanceHsmDestinationClientCertificatesOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(listInstanceHsmDestinationClientCertificatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceHsmDestinationHaGroupsOptions successfully`, func() {
				// Construct an instance of the ListInstanceHsmDestinationHaGroupsOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				listInstanceHsmDestinationHaGroupsOptionsModel := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationHaGroupsOptions(instanceID, destinationID)
				listInstanceHsmDestinationHaGroupsOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationHaGroupsOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationHaGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceHsmDestinationHaGroupsOptionsModel).ToNot(BeNil())
				Expect(listInstanceHsmDestinationHaGroupsOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(listInstanceHsmDestinationHaGroupsOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(listInstanceHsmDestinationHaGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceHsmDestinationPathServerCasOptions successfully`, func() {
				// Construct an instance of the ListInstanceHsmDestinationPathServerCasOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				listInstanceHsmDestinationPathServerCasOptionsModel := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationPathServerCasOptions(instanceID, destinationID, pathID)
				listInstanceHsmDestinationPathServerCasOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathServerCasOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathServerCasOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				listInstanceHsmDestinationPathServerCasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceHsmDestinationPathServerCasOptionsModel).ToNot(BeNil())
				Expect(listInstanceHsmDestinationPathServerCasOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(listInstanceHsmDestinationPathServerCasOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(listInstanceHsmDestinationPathServerCasOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(listInstanceHsmDestinationPathServerCasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceHsmDestinationPathsOptions successfully`, func() {
				// Construct an instance of the ListInstanceHsmDestinationPathsOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				listInstanceHsmDestinationPathsOptionsModel := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationPathsOptions(instanceID, destinationID)
				listInstanceHsmDestinationPathsOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationPathsOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				listInstanceHsmDestinationPathsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceHsmDestinationPathsOptionsModel).ToNot(BeNil())
				Expect(listInstanceHsmDestinationPathsOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(listInstanceHsmDestinationPathsOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(listInstanceHsmDestinationPathsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceHsmDestinationsOptions successfully`, func() {
				// Construct an instance of the ListInstanceHsmDestinationsOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				listInstanceHsmDestinationsOptionsModel := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationsOptions(instanceID)
				listInstanceHsmDestinationsOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				listInstanceHsmDestinationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceHsmDestinationsOptionsModel).ToNot(BeNil())
				Expect(listInstanceHsmDestinationsOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(listInstanceHsmDestinationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResetHsmDestinationPathOptions successfully`, func() {
				// Construct an instance of the ResetHsmDestinationPathOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				resetHsmDestinationPathOptionsModel := secretsManagerInstanceManagementService.NewResetHsmDestinationPathOptions(instanceID, destinationID, pathID)
				resetHsmDestinationPathOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				resetHsmDestinationPathOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				resetHsmDestinationPathOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				resetHsmDestinationPathOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resetHsmDestinationPathOptionsModel).ToNot(BeNil())
				Expect(resetHsmDestinationPathOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(resetHsmDestinationPathOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(resetHsmDestinationPathOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(resetHsmDestinationPathOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunHsmDestinationHealthCheckOptions successfully`, func() {
				// Construct an instance of the RunHsmDestinationHealthCheckOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				runHsmDestinationHealthCheckOptionsModel := secretsManagerInstanceManagementService.NewRunHsmDestinationHealthCheckOptions(instanceID, destinationID)
				runHsmDestinationHealthCheckOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				runHsmDestinationHealthCheckOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				runHsmDestinationHealthCheckOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runHsmDestinationHealthCheckOptionsModel).ToNot(BeNil())
				Expect(runHsmDestinationHealthCheckOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(runHsmDestinationHealthCheckOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(runHsmDestinationHealthCheckOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceDestinationOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				destinationPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateInstanceDestinationOptionsModel := secretsManagerInstanceManagementService.NewUpdateInstanceDestinationOptions(instanceID, destinationID, destinationPatch)
				updateInstanceDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceDestinationOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceDestinationOptionsModel.SetDestinationPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateInstanceDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceDestinationOptionsModel).ToNot(BeNil())
				Expect(updateInstanceDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(updateInstanceDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(updateInstanceDestinationOptionsModel.DestinationPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateInstanceDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceHsmDestinationClientCertificateOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceHsmDestinationClientCertificateOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				certID := "7c9e6679-7425-40de-944b-e07fc1f90ae7"
				clientCertificatePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateInstanceHsmDestinationClientCertificateOptionsModel := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationClientCertificateOptions(instanceID, destinationID, certID, clientCertificatePatch)
				updateInstanceHsmDestinationClientCertificateOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationClientCertificateOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationClientCertificateOptionsModel.SetCertID("7c9e6679-7425-40de-944b-e07fc1f90ae7")
				updateInstanceHsmDestinationClientCertificateOptionsModel.SetClientCertificatePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateInstanceHsmDestinationClientCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceHsmDestinationClientCertificateOptionsModel).ToNot(BeNil())
				Expect(updateInstanceHsmDestinationClientCertificateOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(updateInstanceHsmDestinationClientCertificateOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(updateInstanceHsmDestinationClientCertificateOptionsModel.CertID).To(Equal(core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")))
				Expect(updateInstanceHsmDestinationClientCertificateOptionsModel.ClientCertificatePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateInstanceHsmDestinationClientCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceHsmDestinationHaGroupOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceHsmDestinationHaGroupOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				haGroupID := "ee52ebb6-1728-4580-8ede-13f6504e3ae0"
				hSmhaGroupPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateInstanceHsmDestinationHaGroupOptionsModel := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationHaGroupOptions(instanceID, destinationID, haGroupID, hSmhaGroupPatch)
				updateInstanceHsmDestinationHaGroupOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationHaGroupOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationHaGroupOptionsModel.SetHaGroupID("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
				updateInstanceHsmDestinationHaGroupOptionsModel.SetHSMHAGroupPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateInstanceHsmDestinationHaGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceHsmDestinationHaGroupOptionsModel).ToNot(BeNil())
				Expect(updateInstanceHsmDestinationHaGroupOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(updateInstanceHsmDestinationHaGroupOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(updateInstanceHsmDestinationHaGroupOptionsModel.HaGroupID).To(Equal(core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")))
				Expect(updateInstanceHsmDestinationHaGroupOptionsModel.HSMHAGroupPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateInstanceHsmDestinationHaGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceHsmDestinationOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceHsmDestinationOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				hsmDestinationPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateInstanceHsmDestinationOptionsModel := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationOptions(instanceID, destinationID, hsmDestinationPatch)
				updateInstanceHsmDestinationOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationOptionsModel.SetHsmDestinationPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateInstanceHsmDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceHsmDestinationOptionsModel).ToNot(BeNil())
				Expect(updateInstanceHsmDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(updateInstanceHsmDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(updateInstanceHsmDestinationOptionsModel.HsmDestinationPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateInstanceHsmDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceHsmDestinationPathOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceHsmDestinationPathOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				pathPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateInstanceHsmDestinationPathOptionsModel := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationPathOptions(instanceID, destinationID, pathID, pathPatch)
				updateInstanceHsmDestinationPathOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathOptionsModel.SetPathPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateInstanceHsmDestinationPathOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceHsmDestinationPathOptionsModel).ToNot(BeNil())
				Expect(updateInstanceHsmDestinationPathOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(updateInstanceHsmDestinationPathOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(updateInstanceHsmDestinationPathOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(updateInstanceHsmDestinationPathOptionsModel.PathPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateInstanceHsmDestinationPathOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceHsmDestinationPathServerCaOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceHsmDestinationPathServerCaOptions model
				instanceID := "fe5c5c27-3e44-474a-9484-55892e12d6d1"
				destinationID := "b2c3d4e5-f6a7-8901-bcde-f12345678901"
				pathID := "fa3d1234-0000-0000-0000-000000000001"
				serverCaID := "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
				serverCaPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateInstanceHsmDestinationPathServerCaOptionsModel := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationPathServerCaOptions(instanceID, destinationID, pathID, serverCaID, serverCaPatch)
				updateInstanceHsmDestinationPathServerCaOptionsModel.SetInstanceID("fe5c5c27-3e44-474a-9484-55892e12d6d1")
				updateInstanceHsmDestinationPathServerCaOptionsModel.SetDestinationID("b2c3d4e5-f6a7-8901-bcde-f12345678901")
				updateInstanceHsmDestinationPathServerCaOptionsModel.SetPathID("fa3d1234-0000-0000-0000-000000000001")
				updateInstanceHsmDestinationPathServerCaOptionsModel.SetServerCaID("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
				updateInstanceHsmDestinationPathServerCaOptionsModel.SetServerCaPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateInstanceHsmDestinationPathServerCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceHsmDestinationPathServerCaOptionsModel).ToNot(BeNil())
				Expect(updateInstanceHsmDestinationPathServerCaOptionsModel.InstanceID).To(Equal(core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1")))
				Expect(updateInstanceHsmDestinationPathServerCaOptionsModel.DestinationID).To(Equal(core.StringPtr("b2c3d4e5-f6a7-8901-bcde-f12345678901")))
				Expect(updateInstanceHsmDestinationPathServerCaOptionsModel.PathID).To(Equal(core.StringPtr("fa3d1234-0000-0000-0000-000000000001")))
				Expect(updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaID).To(Equal(core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")))
				Expect(updateInstanceHsmDestinationPathServerCaOptionsModel.ServerCaPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateInstanceHsmDestinationPathServerCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDestinationRequestIbmCloudDatabaseDestinationPrototype successfully`, func() {
				name := "testString"
				typeVar := "ibm_cloud_database"
				crn := "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::"
				_model, err := secretsManagerInstanceManagementService.NewCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(name, typeVar, crn)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalClientCertificatePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.ClientCertificatePatch)
			model.Name = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Status = core.StringPtr("testString")
			model.SignedCert = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *secretsmanagerinstancemanagementv2.ClientCertificatePatch
			err = secretsmanagerinstancemanagementv2.UnmarshalClientCertificatePatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateDestinationRequest successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.CreateDestinationRequest)
			model.Name = core.StringPtr("testString")
			model.Type = core.StringPtr("ibm_cloud_database")
			model.Description = core.StringPtr("Production PostgreSQL database")
			model.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result interface{}
			err = secretsmanagerinstancemanagementv2.UnmarshalCreateDestinationRequest(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
		It(`Invoke UnmarshalDestinationPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.DestinationPatch)
			model.Name = core.StringPtr("production-postgres-db")
			model.Description = core.StringPtr("Updated description for production database")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *secretsmanagerinstancemanagementv2.DestinationPatch
			err = secretsmanagerinstancemanagementv2.UnmarshalDestinationPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalHSMDestinationPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.HSMDestinationPatch)
			model.Name = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *secretsmanagerinstancemanagementv2.HSMDestinationPatch
			err = secretsmanagerinstancemanagementv2.UnmarshalHSMDestinationPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalHSMHAGroupPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.HSMHAGroupPatch)
			model.Name = core.StringPtr("dr-ha-group")
			model.Description = core.StringPtr("DR HA Group for failover")
			model.MemberSerials = []string{"123456", "123457", "123458"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *secretsmanagerinstancemanagementv2.HSMHAGroupPatch
			err = secretsmanagerinstancemanagementv2.UnmarshalHSMHAGroupPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPathPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.PathPatch)
			model.Name = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *secretsmanagerinstancemanagementv2.PathPatch
			err = secretsmanagerinstancemanagementv2.UnmarshalPathPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalServerCaPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.ServerCaPatch)
			model.Name = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Status = core.StringPtr("active")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *secretsmanagerinstancemanagementv2.ServerCaPatch
			err = secretsmanagerinstancemanagementv2.UnmarshalServerCaPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateDestinationRequestIbmCloudDatabaseDestinationPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
			model.Name = core.StringPtr("testString")
			model.Type = core.StringPtr("ibm_cloud_database")
			model.Description = core.StringPtr("Production PostgreSQL database")
			model.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype
			err = secretsmanagerinstancemanagementv2.UnmarshalCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
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
