//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the secretsmanagerinstancemanagementv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SecretsManagerInstanceManagementV2 Integration Tests`, func() {
	const externalConfigFile = "../secrets_manager_instance_management_v2.env"

	var (
		err          error
		secretsManagerInstanceManagementService *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(secretsmanagerinstancemanagementv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			secretsManagerInstanceManagementServiceOptions := &secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{}

			secretsManagerInstanceManagementService, err = secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2UsingExternalConfig(secretsManagerInstanceManagementServiceOptions)
			Expect(err).To(BeNil())
			Expect(secretsManagerInstanceManagementService).ToNot(BeNil())
			Expect(secretsManagerInstanceManagementService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			secretsManagerInstanceManagementService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateVaultAdmintoken - Generate admin token`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVaultAdmintoken(createVaultAdmintokenOptions *CreateVaultAdmintokenOptions)`, func() {
			createVaultAdmintokenOptions := &secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions{
				InstanceID: core.StringPtr("bfc50c2e-d66d-4f37-9ccf-9713f8325b39"),
			}

			token, response, err := secretsManagerInstanceManagementService.CreateVaultAdmintoken(createVaultAdmintokenOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(token).ToNot(BeNil())
		})
	})

	Describe(`GetInstance - Get instance details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstance(getInstanceOptions *GetInstanceOptions)`, func() {
			getInstanceOptions := &secretsmanagerinstancemanagementv2.GetInstanceOptions{
				InstanceID: core.StringPtr("bfc50c2e-d66d-4f37-9ccf-9713f8325b39"),
			}

			instance, response, err := secretsManagerInstanceManagementService.GetInstance(getInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())
		})
	})

	Describe(`DeleteInstanceAdmintokens - Revoke admin tokens`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptions *DeleteInstanceAdmintokensOptions)`, func() {
			deleteInstanceAdmintokensOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceAdmintokensOptions{
				InstanceID: core.StringPtr("bfc50c2e-d66d-4f37-9ccf-9713f8325b39"),
			}

			response, err := secretsManagerInstanceManagementService.DeleteInstanceAdmintokens(deleteInstanceAdmintokensOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
