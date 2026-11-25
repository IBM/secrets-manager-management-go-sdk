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

	Describe(`CreateVaultAdmintoken - Create admin token`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVaultAdmintoken(createVaultAdmintokenOptions *CreateVaultAdmintokenOptions)`, func() {
			createVaultAdmintokenOptions := &secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				ResponseWrapping: core.BoolPtr(true),
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
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
			}

			instance, response, err := secretsManagerInstanceManagementService.GetInstance(getInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceDestinations - List destinations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceDestinations(listInstanceDestinationsOptions *ListInstanceDestinationsOptions)`, func() {
			listInstanceDestinationsOptions := &secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				State: core.StringPtr("not_started"),
			}

			destinationCollection, response, err := secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateInstanceDestination - Create destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceDestination(createInstanceDestinationOptions *CreateInstanceDestinationOptions)`, func() {
			createDestinationRequestModel := &secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype{
				Name: core.StringPtr("my-postgres"),
				Type: core.StringPtr("ibm_cloud_database"),
				Description: core.StringPtr("Production PostgreSQL database"),
				Crn: core.StringPtr("crn:v1:staging:public:databases-for-postgresql:us-south:a/791f5fb10986423e97aa8512f18b7e65:11111111-0000-4000-8000-111111111111::"),
			}

			createInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationPrototype: createDestinationRequestModel,
			}

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceDestination - Get destination details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceDestination(getInstanceDestinationOptions *GetInstanceDestinationOptions)`, func() {
			getInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceDestination - Update destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceDestination(updateInstanceDestinationOptions *UpdateInstanceDestinationOptions)`, func() {
			destinationPatchModel := &secretsmanagerinstancemanagementv2.DestinationPatch{
				Name: core.StringPtr("my-postgres-updated"),
				Description: core.StringPtr("Production PostgreSQL database"),
			}
			destinationPatchModelAsPatch, asPatchErr := destinationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				DestinationPatch: destinationPatchModelAsPatch,
			}

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceHsmDestinations - List destinations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceHsmDestinations(listInstanceHsmDestinationsOptions *ListInstanceHsmDestinationsOptions)`, func() {
			listInstanceHsmDestinationsOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
			}

			hsmDestinationCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmDestinationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateInstanceHsmDestination - Create destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceHsmDestination(createInstanceHsmDestinationOptions *CreateInstanceHsmDestinationOptions)`, func() {
			createInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				Name: core.StringPtr("prod-hsm-destination"),
				Vendor: core.StringPtr("thales_luna"),
				Description: core.StringPtr("Primary Thales Luna destination"),
			}

			hsmDestination, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(hsmDestination).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceHsmDestination - Get destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceHsmDestination(getInstanceHsmDestinationOptions *GetInstanceHsmDestinationOptions)`, func() {
			getInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			hsmDestination, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmDestination).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceHsmDestination - Update destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptions *UpdateInstanceHsmDestinationOptions)`, func() {
			hsmDestinationPatchModel := &secretsmanagerinstancemanagementv2.HSMDestinationPatch{
				Name: core.StringPtr("my-hsm-destination"),
				Description: core.StringPtr("testString"),
			}
			hsmDestinationPatchModelAsPatch, asPatchErr := hsmDestinationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				HsmDestinationPatch: hsmDestinationPatchModelAsPatch,
			}

			hsmDestination, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmDestination).ToNot(BeNil())
		})
	})

	Describe(`GetHsmDestinationHealthSnapshot - Get destination health snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptions *GetHsmDestinationHealthSnapshotOptions)`, func() {
			getHsmDestinationHealthSnapshotOptions := &secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			destinationHealthSnapshot, response, err := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationHealthSnapshot).ToNot(BeNil())
		})
	})

	Describe(`RunHsmDestinationHealthCheck - Trigger an async destination health check`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptions *RunHsmDestinationHealthCheckOptions)`, func() {
			runHsmDestinationHealthCheckOptions := &secretsmanagerinstancemanagementv2.RunHsmDestinationHealthCheckOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			asyncHealthCheckAckResponse, response, err := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(asyncHealthCheckAckResponse).ToNot(BeNil())
		})
	})

	Describe(`ApplyHsmDestination - Trigger an async HSM configuration apply`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ApplyHsmDestination(applyHsmDestinationOptions *ApplyHsmDestinationOptions)`, func() {
			applyHsmDestinationOptions := &secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			asyncHealthAckResponse, response, err := secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(asyncHealthAckResponse).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceHsmDestinationHaGroups - List HSM HA Groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptions *ListInstanceHsmDestinationHaGroupsOptions)`, func() {
			listInstanceHsmDestinationHaGroupsOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			hsmhaGroupCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmhaGroupCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateInstanceHsmDestinationHaGroup - Create HA Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptions *CreateInstanceHsmDestinationHaGroupOptions)`, func() {
			createInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				Name: core.StringPtr("production-ha-group"),
				GroupLabel: core.StringPtr("prod_ha_group"),
				GroupSerial: core.StringPtr("11323465009866"),
				MemberSerials: []string{"123456", "123457"},
				Description: core.StringPtr("Primary HA Group for prod"),
			}

			hsmhaGroup, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(hsmhaGroup).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceHsmDestinationHaGroup - Get HA Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptions *GetInstanceHsmDestinationHaGroupOptions)`, func() {
			getInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				HaGroupID: core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0"),
			}

			hsmhaGroup, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmhaGroup).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceHsmDestinationHaGroup - Update HA Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptions *UpdateInstanceHsmDestinationHaGroupOptions)`, func() {
			hsmhaGroupPatchModel := &secretsmanagerinstancemanagementv2.HSMHAGroupPatch{
				Name: core.StringPtr("dr-ha-group"),
				Description: core.StringPtr("DR HA Group for failover"),
				MemberSerials: []string{"123456", "123457", "123458"},
			}
			hsmhaGroupPatchModelAsPatch, asPatchErr := hsmhaGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				HaGroupID: core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0"),
				HSMHAGroupPatch: hsmhaGroupPatchModelAsPatch,
			}

			hsmhaGroup, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmhaGroup).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceHsmDestinationPaths - List paths`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptions *ListInstanceHsmDestinationPathsOptions)`, func() {
			listInstanceHsmDestinationPathsOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			pathCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pathCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateInstanceHsmDestinationPath - Create path`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptions *CreateInstanceHsmDestinationPathOptions)`, func() {
			createInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				Name: core.StringPtr("hsm-node-1"),
				Crn: core.StringPtr("crn:v1:staging:public:databases-for-postgresql:us-south:a/791f5fb10986423e97aa8512f18b7e65:11111111-0000-4000-8000-111111111111::"),
				Description: core.StringPtr("testString"),
			}

			path, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(path).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceHsmDestinationPath - Get path`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions *GetInstanceHsmDestinationPathOptions)`, func() {
			getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
			}

			path, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(path).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceHsmDestinationPath - Update path`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptions *UpdateInstanceHsmDestinationPathOptions)`, func() {
			pathPatchModel := &secretsmanagerinstancemanagementv2.PathPatch{
				Name: core.StringPtr("hsm-node-1-renamed"),
				Description: core.StringPtr("testString"),
			}
			pathPatchModelAsPatch, asPatchErr := pathPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
				PathPatch: pathPatchModelAsPatch,
			}

			path, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(path).ToNot(BeNil())
		})
	})

	Describe(`CreatePathBuild - Build path (trigger provisioning)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePathBuild(createPathBuildOptions *CreatePathBuildOptions)`, func() {
			createPathBuildOptions := &secretsmanagerinstancemanagementv2.CreatePathBuildOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
			}

			vpeProvisioningResponse, response, err := secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vpeProvisioningResponse).ToNot(BeNil())
		})
	})

	Describe(`GetPathVpeStatus - Health check path (check provisioning status)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPathVpeStatus(getPathVpeStatusOptions *GetPathVpeStatusOptions)`, func() {
			getPathVpeStatusOptions := &secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
			}

			vpeStatusResult, response, err := secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpeStatusResult).ToNot(BeNil())
		})
	})

	Describe(`ResetHsmDestinationPath - Reset a path to its initial pending state`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResetHsmDestinationPath(resetHsmDestinationPathOptions *ResetHsmDestinationPathOptions)`, func() {
			resetHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.ResetHsmDestinationPathOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
			}

			hsmPathReset, response, err := secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmPathReset).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceHsmDestinationPathServerCas - List Server CA certificates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptions *ListInstanceHsmDestinationPathServerCasOptions)`, func() {
			listInstanceHsmDestinationPathServerCasOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
			}

			serverCaCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serverCaCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateInstanceHsmDestinationPathServerCa - Add Server CA certificate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptions *CreateInstanceHsmDestinationPathServerCaOptions)`, func() {
			createInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
				Name: core.StringPtr("HSM1_Server_CA"),
				Certificate: core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----"),
				Description: core.StringPtr("testString"),
			}

			serverCa, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serverCa).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceHsmDestinationPathServerCa - Get Server CA certificate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptions *GetInstanceHsmDestinationPathServerCaOptions)`, func() {
			getInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
				ServerCaID: core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"),
			}

			serverCa, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serverCa).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceHsmDestinationPathServerCa - Update or activate Server CA`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptions *UpdateInstanceHsmDestinationPathServerCaOptions)`, func() {
			serverCaPatchModel := &secretsmanagerinstancemanagementv2.ServerCaPatch{
				Name: core.StringPtr("HSM1_Server_CA_renamed"),
				Description: core.StringPtr("testString"),
				Status: core.StringPtr("active"),
			}
			serverCaPatchModelAsPatch, asPatchErr := serverCaPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
				ServerCaID: core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"),
				ServerCaPatch: serverCaPatchModelAsPatch,
			}

			serverCa, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serverCa).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceHsmDestinationClientCertificates - List client certificates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptions *ListInstanceHsmDestinationClientCertificatesOptions)`, func() {
			listInstanceHsmDestinationClientCertificatesOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			clientCertificateCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clientCertificateCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateInstanceHsmDestinationClientCertificate - Generate client certificate (CSR)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptions *CreateInstanceHsmDestinationClientCertificateOptions)`, func() {
			createInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				Name: core.StringPtr("hsm-client-cert-prod"),
				Description: core.StringPtr("testString"),
			}

			clientCertificate, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(clientCertificate).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceHsmDestinationClientCertificate - Get client certificate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptions *GetInstanceHsmDestinationClientCertificateOptions)`, func() {
			getInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				CertID: core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7"),
			}

			clientCertificate, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clientCertificate).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceHsmDestinationClientCertificate - Update, upload signed cert, and/or activate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptions *UpdateInstanceHsmDestinationClientCertificateOptions)`, func() {
			clientCertificatePatchModel := &secretsmanagerinstancemanagementv2.ClientCertificatePatch{
				Name: core.StringPtr("hsm-client-cert-renamed"),
				Description: core.StringPtr("testString"),
				Status: core.StringPtr("testString"),
				SignedCert: core.StringPtr("testString"),
			}
			clientCertificatePatchModelAsPatch, asPatchErr := clientCertificatePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				CertID: core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7"),
				ClientCertificatePatch: clientCertificatePatchModelAsPatch,
			}

			clientCertificate, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clientCertificate).ToNot(BeNil())
		})
	})

	Describe(`DeleteVaultAdmintokens - Delete admin tokens`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVaultAdmintokens(deleteVaultAdmintokensOptions *DeleteVaultAdmintokensOptions)`, func() {
			deleteVaultAdmintokensOptions := &secretsmanagerinstancemanagementv2.DeleteVaultAdmintokensOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
			}

			response, err := secretsManagerInstanceManagementService.DeleteVaultAdmintokens(deleteVaultAdmintokensOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteInstanceDestination - Delete destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceDestination(deleteInstanceDestinationOptions *DeleteInstanceDestinationOptions)`, func() {
			deleteInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
	})

	Describe(`DeleteInstanceHsmDestination - Delete a destination asynchronously`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptions *DeleteInstanceHsmDestinationOptions)`, func() {
			deleteInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
			}

			asyncDeleteResponse, response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(asyncDeleteResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteInstanceHsmDestinationHaGroup - Delete HA Group (synchronous hard delete)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptions *DeleteInstanceHsmDestinationHaGroupOptions)`, func() {
			deleteInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationHaGroupOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				HaGroupID: core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0"),
			}

			response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteInstanceHsmDestinationPath - Delete path (async hard delete)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptions *DeleteInstanceHsmDestinationPathOptions)`, func() {
			deleteInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
			}

			path, response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(path).ToNot(BeNil())
		})
	})

	Describe(`DeleteInstanceHsmDestinationPathServerCa - Delete Server CA certificate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptions *DeleteInstanceHsmDestinationPathServerCaOptions)`, func() {
			deleteInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathServerCaOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				PathID: core.StringPtr("fa3d1234-0000-0000-0000-000000000001"),
				ServerCaID: core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"),
			}

			response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteInstanceHsmDestinationClientCertificate - Delete client certificate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptions *DeleteInstanceHsmDestinationClientCertificateOptions)`, func() {
			deleteInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationClientCertificateOptions{
				InstanceID: core.StringPtr("fe5c5c27-3e44-474a-9484-55892e12d6d1"),
				DestinationID: core.StringPtr("3b96bc43-0d41-4908-9b6d-2e2682afc00e"),
				CertID: core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7"),
			}

			response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
