//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the secrets-manager-instance-management service.
//
// The following configuration properties are assumed to be defined:
// SECRETS_MANAGER_INSTANCE_MANAGEMENT_URL=<service base url>
// SECRETS_MANAGER_INSTANCE_MANAGEMENT_AUTH_TYPE=iam
// SECRETS_MANAGER_INSTANCE_MANAGEMENT_APIKEY=<IAM apikey>
// SECRETS_MANAGER_INSTANCE_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`SecretsManagerInstanceManagementV2 Examples Tests`, func() {

	const externalConfigFile = "../secrets_manager_instance_management_v2.env"

	var (
		secretsManagerInstanceManagementService *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(secretsmanagerinstancemanagementv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			secretsManagerInstanceManagementServiceOptions := &secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{}

			secretsManagerInstanceManagementService, err = secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2UsingExternalConfig(secretsManagerInstanceManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
		})
	})

	Describe(`SecretsManagerInstanceManagementV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVaultAdmintoken request example`, func() {
			fmt.Println("\nCreateVaultAdmintoken() result:")
			// begin-create_vault_admintoken

			createVaultAdmintokenOptions := secretsManagerInstanceManagementService.NewCreateVaultAdmintokenOptions(
				"bfc50c2e-d66d-4f37-9ccf-9713f8325b39",
			)

			token, response, err := secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(token, "", "  ")
			fmt.Println(string(b))

			// end-create_vault_admintoken

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(token).ToNot(BeNil())
		})
		It(`GetInstance request example`, func() {
			fmt.Println("\nGetInstance() result:")
			// begin-get_instance

			getInstanceOptions := secretsManagerInstanceManagementService.NewGetInstanceOptions(
				"bfc50c2e-d66d-4f37-9ccf-9713f8325b39",
			)

			instance, response, err := secretsManagerInstanceManagementService.GetInstance(getInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instance, "", "  ")
			fmt.Println(string(b))

			// end-get_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())
		})
		It(`DeleteInstanceAdmintokens request example`, func() {
			// begin-delete_instance_admintokens

			deleteInstanceAdmintokensOptions := secretsManagerInstanceManagementService.NewDeleteInstanceAdmintokensOptions(
				"bfc50c2e-d66d-4f37-9ccf-9713f8325b39",
			)

			response, err := secretsManagerInstanceManagementService.DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteInstanceAdmintokens(): %d\n", response.StatusCode)
			}

			// end-delete_instance_admintokens

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
