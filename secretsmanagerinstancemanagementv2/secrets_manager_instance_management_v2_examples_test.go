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
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
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
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
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
		It(`ListInstanceDestinations request example`, func() {
			fmt.Println("\nListInstanceDestinations() result:")
			// begin-list_instance_destinations

			listInstanceDestinationsOptions := secretsManagerInstanceManagementService.NewListInstanceDestinationsOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
			)

			destinationCollection, response, err := secretsManagerInstanceManagementService.ListInstanceDestinations(listInstanceDestinationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destinationCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_destinations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationCollection).ToNot(BeNil())
		})
		It(`CreateInstanceDestination request example`, func() {
			fmt.Println("\nCreateInstanceDestination() result:")
			// begin-create_instance_destination

			createDestinationRequestModel := &secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype{
				Name: core.StringPtr("my-postgres"),
				Type: core.StringPtr("ibm_cloud_database"),
				Description: core.StringPtr("Production PostgreSQL database"),
				Crn: core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::"),
			}

			createInstanceDestinationOptions := secretsManagerInstanceManagementService.NewCreateInstanceDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				createDestinationRequestModel,
			)

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.CreateInstanceDestination(createInstanceDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ibmCloudDatabaseDestination, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
		It(`GetInstanceDestination request example`, func() {
			fmt.Println("\nGetInstanceDestination() result:")
			// begin-get_instance_destination

			getInstanceDestinationOptions := secretsManagerInstanceManagementService.NewGetInstanceDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.GetInstanceDestination(getInstanceDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ibmCloudDatabaseDestination, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
		It(`UpdateInstanceDestination request example`, func() {
			fmt.Println("\nUpdateInstanceDestination() result:")
			// begin-update_instance_destination

			destinationPatchModel := &secretsmanagerinstancemanagementv2.DestinationPatch{
				Name: core.StringPtr("my-postgres-updated"),
				Description: core.StringPtr("Production PostgreSQL database"),
			}
			destinationPatchModelAsPatch, asPatchErr := destinationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceDestinationOptions := secretsManagerInstanceManagementService.NewUpdateInstanceDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				destinationPatchModelAsPatch,
			)

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.UpdateInstanceDestination(updateInstanceDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ibmCloudDatabaseDestination, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
		It(`ListInstanceHsmDestinations request example`, func() {
			fmt.Println("\nListInstanceHsmDestinations() result:")
			// begin-list_instance_hsm_destinations

			listInstanceHsmDestinationsOptions := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationsOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
			)

			hsmDestinationCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinations(listInstanceHsmDestinationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmDestinationCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_hsm_destinations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmDestinationCollection).ToNot(BeNil())
		})
		It(`CreateInstanceHsmDestination request example`, func() {
			fmt.Println("\nCreateInstanceHsmDestination() result:")
			// begin-create_instance_hsm_destination

			createInstanceHsmDestinationOptions := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"prod-hsm-destination",
				"thales_luna",
			)
			createInstanceHsmDestinationOptions.SetDescription("Primary Thales Luna destination")

			hsmDestination, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestination(createInstanceHsmDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmDestination, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_hsm_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(hsmDestination).ToNot(BeNil())
		})
		It(`GetInstanceHsmDestination request example`, func() {
			fmt.Println("\nGetInstanceHsmDestination() result:")
			// begin-get_instance_hsm_destination

			getInstanceHsmDestinationOptions := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			hsmDestination, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestination(getInstanceHsmDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmDestination, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_hsm_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmDestination).ToNot(BeNil())
		})
		It(`UpdateInstanceHsmDestination request example`, func() {
			fmt.Println("\nUpdateInstanceHsmDestination() result:")
			// begin-update_instance_hsm_destination

			hsmDestinationPatchModel := &secretsmanagerinstancemanagementv2.HSMDestinationPatch{
				Name: core.StringPtr("my-hsm-destination"),
			}
			hsmDestinationPatchModelAsPatch, asPatchErr := hsmDestinationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationOptions := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				hsmDestinationPatchModelAsPatch,
			)

			hsmDestination, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestination(updateInstanceHsmDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmDestination, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_hsm_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmDestination).ToNot(BeNil())
		})
		It(`GetHsmDestinationHealthSnapshot request example`, func() {
			fmt.Println("\nGetHsmDestinationHealthSnapshot() result:")
			// begin-get_hsm_destination_health_snapshot

			getHsmDestinationHealthSnapshotOptions := secretsManagerInstanceManagementService.NewGetHsmDestinationHealthSnapshotOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			destinationHealthSnapshot, response, err := secretsManagerInstanceManagementService.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destinationHealthSnapshot, "", "  ")
			fmt.Println(string(b))

			// end-get_hsm_destination_health_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationHealthSnapshot).ToNot(BeNil())
		})
		It(`RunHsmDestinationHealthCheck request example`, func() {
			fmt.Println("\nRunHsmDestinationHealthCheck() result:")
			// begin-run_hsm_destination_health_check

			runHsmDestinationHealthCheckOptions := secretsManagerInstanceManagementService.NewRunHsmDestinationHealthCheckOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			asyncHealthCheckAckResponse, response, err := secretsManagerInstanceManagementService.RunHsmDestinationHealthCheck(runHsmDestinationHealthCheckOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(asyncHealthCheckAckResponse, "", "  ")
			fmt.Println(string(b))

			// end-run_hsm_destination_health_check

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(asyncHealthCheckAckResponse).ToNot(BeNil())
		})
		It(`ApplyHsmDestination request example`, func() {
			fmt.Println("\nApplyHsmDestination() result:")
			// begin-apply_hsm_destination

			applyHsmDestinationOptions := secretsManagerInstanceManagementService.NewApplyHsmDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			asyncHealthAckResponse, response, err := secretsManagerInstanceManagementService.ApplyHsmDestination(applyHsmDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(asyncHealthAckResponse, "", "  ")
			fmt.Println(string(b))

			// end-apply_hsm_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(asyncHealthAckResponse).ToNot(BeNil())
		})
		It(`ListInstanceHsmDestinationHaGroups request example`, func() {
			fmt.Println("\nListInstanceHsmDestinationHaGroups() result:")
			// begin-list_instance_hsm_destination_ha_groups

			listInstanceHsmDestinationHaGroupsOptions := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationHaGroupsOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			hsmhaGroupCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationHaGroups(listInstanceHsmDestinationHaGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmhaGroupCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_hsm_destination_ha_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmhaGroupCollection).ToNot(BeNil())
		})
		It(`CreateInstanceHsmDestinationHaGroup request example`, func() {
			fmt.Println("\nCreateInstanceHsmDestinationHaGroup() result:")
			// begin-create_instance_hsm_destination_ha_group

			createInstanceHsmDestinationHaGroupOptions := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationHaGroupOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"production-ha-group",
				"prod_ha_group",
				"11323465009866",
				[]string{"123456", "123457"},
			)
			createInstanceHsmDestinationHaGroupOptions.SetDescription("Primary HA Group for prod")

			hsmhaGroup, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationHaGroup(createInstanceHsmDestinationHaGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmhaGroup, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_hsm_destination_ha_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(hsmhaGroup).ToNot(BeNil())
		})
		It(`GetInstanceHsmDestinationHaGroup request example`, func() {
			fmt.Println("\nGetInstanceHsmDestinationHaGroup() result:")
			// begin-get_instance_hsm_destination_ha_group

			getInstanceHsmDestinationHaGroupOptions := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationHaGroupOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"ee52ebb6-1728-4580-8ede-13f6504e3ae0",
			)

			hsmhaGroup, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationHaGroup(getInstanceHsmDestinationHaGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmhaGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_hsm_destination_ha_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmhaGroup).ToNot(BeNil())
		})
		It(`UpdateInstanceHsmDestinationHaGroup request example`, func() {
			fmt.Println("\nUpdateInstanceHsmDestinationHaGroup() result:")
			// begin-update_instance_hsm_destination_ha_group

			hsmhaGroupPatchModel := &secretsmanagerinstancemanagementv2.HSMHAGroupPatch{
				Name: core.StringPtr("dr-ha-group"),
			}
			hsmhaGroupPatchModelAsPatch, asPatchErr := hsmhaGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationHaGroupOptions := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationHaGroupOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"ee52ebb6-1728-4580-8ede-13f6504e3ae0",
				hsmhaGroupPatchModelAsPatch,
			)

			hsmhaGroup, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationHaGroup(updateInstanceHsmDestinationHaGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmhaGroup, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_hsm_destination_ha_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmhaGroup).ToNot(BeNil())
		})
		It(`ListInstanceHsmDestinationPaths request example`, func() {
			fmt.Println("\nListInstanceHsmDestinationPaths() result:")
			// begin-list_instance_hsm_destination_paths

			listInstanceHsmDestinationPathsOptions := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationPathsOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			pathCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPaths(listInstanceHsmDestinationPathsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pathCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_hsm_destination_paths

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pathCollection).ToNot(BeNil())
		})
		It(`CreateInstanceHsmDestinationPath request example`, func() {
			fmt.Println("\nCreateInstanceHsmDestinationPath() result:")
			// begin-create_instance_hsm_destination_path

			createInstanceHsmDestinationPathOptions := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationPathOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"hsm-node-1",
				"crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1",
			)

			path, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPath(createInstanceHsmDestinationPathOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(path, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_hsm_destination_path

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(path).ToNot(BeNil())
		})
		It(`GetInstanceHsmDestinationPath request example`, func() {
			fmt.Println("\nGetInstanceHsmDestinationPath() result:")
			// begin-get_instance_hsm_destination_path

			getInstanceHsmDestinationPathOptions := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationPathOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
			)

			path, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(path, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_hsm_destination_path

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(path).ToNot(BeNil())
		})
		It(`UpdateInstanceHsmDestinationPath request example`, func() {
			fmt.Println("\nUpdateInstanceHsmDestinationPath() result:")
			// begin-update_instance_hsm_destination_path

			pathPatchModel := &secretsmanagerinstancemanagementv2.PathPatch{
				Name: core.StringPtr("hsm-node-1-renamed"),
			}
			pathPatchModelAsPatch, asPatchErr := pathPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationPathOptions := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationPathOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
				pathPatchModelAsPatch,
			)

			path, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPath(updateInstanceHsmDestinationPathOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(path, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_hsm_destination_path

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(path).ToNot(BeNil())
		})
		It(`CreatePathBuild request example`, func() {
			fmt.Println("\nCreatePathBuild() result:")
			// begin-create_path_build

			createPathBuildOptions := secretsManagerInstanceManagementService.NewCreatePathBuildOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
			)

			vpeProvisioningResponse, response, err := secretsManagerInstanceManagementService.CreatePathBuild(createPathBuildOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpeProvisioningResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_path_build

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vpeProvisioningResponse).ToNot(BeNil())
		})
		It(`GetPathVpeStatus request example`, func() {
			fmt.Println("\nGetPathVpeStatus() result:")
			// begin-get_path_vpe_status

			getPathVpeStatusOptions := secretsManagerInstanceManagementService.NewGetPathVpeStatusOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
			)

			vpeStatusResult, response, err := secretsManagerInstanceManagementService.GetPathVpeStatus(getPathVpeStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpeStatusResult, "", "  ")
			fmt.Println(string(b))

			// end-get_path_vpe_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpeStatusResult).ToNot(BeNil())
		})
		It(`ResetHsmDestinationPath request example`, func() {
			fmt.Println("\nResetHsmDestinationPath() result:")
			// begin-reset_hsm_destination_path

			resetHsmDestinationPathOptions := secretsManagerInstanceManagementService.NewResetHsmDestinationPathOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
			)

			hsmPathReset, response, err := secretsManagerInstanceManagementService.ResetHsmDestinationPath(resetHsmDestinationPathOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(hsmPathReset, "", "  ")
			fmt.Println(string(b))

			// end-reset_hsm_destination_path

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hsmPathReset).ToNot(BeNil())
		})
		It(`ListInstanceHsmDestinationPathServerCas request example`, func() {
			fmt.Println("\nListInstanceHsmDestinationPathServerCas() result:")
			// begin-list_instance_hsm_destination_path_server_cas

			listInstanceHsmDestinationPathServerCasOptions := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationPathServerCasOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
			)

			serverCaCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationPathServerCas(listInstanceHsmDestinationPathServerCasOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serverCaCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_hsm_destination_path_server_cas

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serverCaCollection).ToNot(BeNil())
		})
		It(`CreateInstanceHsmDestinationPathServerCa request example`, func() {
			fmt.Println("\nCreateInstanceHsmDestinationPathServerCa() result:")
			// begin-create_instance_hsm_destination_path_server_ca

			createInstanceHsmDestinationPathServerCaOptions := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationPathServerCaOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
				"HSM1_Server_CA",
				"-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----",
			)

			serverCa, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationPathServerCa(createInstanceHsmDestinationPathServerCaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serverCa, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_hsm_destination_path_server_ca

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serverCa).ToNot(BeNil())
		})
		It(`GetInstanceHsmDestinationPathServerCa request example`, func() {
			fmt.Println("\nGetInstanceHsmDestinationPathServerCa() result:")
			// begin-get_instance_hsm_destination_path_server_ca

			getInstanceHsmDestinationPathServerCaOptions := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationPathServerCaOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
				"5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f",
			)

			serverCa, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationPathServerCa(getInstanceHsmDestinationPathServerCaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serverCa, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_hsm_destination_path_server_ca

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serverCa).ToNot(BeNil())
		})
		It(`UpdateInstanceHsmDestinationPathServerCa request example`, func() {
			fmt.Println("\nUpdateInstanceHsmDestinationPathServerCa() result:")
			// begin-update_instance_hsm_destination_path_server_ca

			serverCaPatchModel := &secretsmanagerinstancemanagementv2.ServerCaPatch{
				Name: core.StringPtr("HSM1_Server_CA_renamed"),
			}
			serverCaPatchModelAsPatch, asPatchErr := serverCaPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationPathServerCaOptions := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationPathServerCaOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
				"5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f",
				serverCaPatchModelAsPatch,
			)

			serverCa, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationPathServerCa(updateInstanceHsmDestinationPathServerCaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serverCa, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_hsm_destination_path_server_ca

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serverCa).ToNot(BeNil())
		})
		It(`ListInstanceHsmDestinationClientCertificates request example`, func() {
			fmt.Println("\nListInstanceHsmDestinationClientCertificates() result:")
			// begin-list_instance_hsm_destination_client_certificates

			listInstanceHsmDestinationClientCertificatesOptions := secretsManagerInstanceManagementService.NewListInstanceHsmDestinationClientCertificatesOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			clientCertificateCollection, response, err := secretsManagerInstanceManagementService.ListInstanceHsmDestinationClientCertificates(listInstanceHsmDestinationClientCertificatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clientCertificateCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_hsm_destination_client_certificates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clientCertificateCollection).ToNot(BeNil())
		})
		It(`CreateInstanceHsmDestinationClientCertificate request example`, func() {
			fmt.Println("\nCreateInstanceHsmDestinationClientCertificate() result:")
			// begin-create_instance_hsm_destination_client_certificate

			createInstanceHsmDestinationClientCertificateOptions := secretsManagerInstanceManagementService.NewCreateInstanceHsmDestinationClientCertificateOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"hsm-client-cert-prod",
			)

			clientCertificate, response, err := secretsManagerInstanceManagementService.CreateInstanceHsmDestinationClientCertificate(createInstanceHsmDestinationClientCertificateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clientCertificate, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_hsm_destination_client_certificate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(clientCertificate).ToNot(BeNil())
		})
		It(`GetInstanceHsmDestinationClientCertificate request example`, func() {
			fmt.Println("\nGetInstanceHsmDestinationClientCertificate() result:")
			// begin-get_instance_hsm_destination_client_certificate

			getInstanceHsmDestinationClientCertificateOptions := secretsManagerInstanceManagementService.NewGetInstanceHsmDestinationClientCertificateOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"7c9e6679-7425-40de-944b-e07fc1f90ae7",
			)

			clientCertificate, response, err := secretsManagerInstanceManagementService.GetInstanceHsmDestinationClientCertificate(getInstanceHsmDestinationClientCertificateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clientCertificate, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_hsm_destination_client_certificate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clientCertificate).ToNot(BeNil())
		})
		It(`UpdateInstanceHsmDestinationClientCertificate request example`, func() {
			fmt.Println("\nUpdateInstanceHsmDestinationClientCertificate() result:")
			// begin-update_instance_hsm_destination_client_certificate

			clientCertificatePatchModel := &secretsmanagerinstancemanagementv2.ClientCertificatePatch{
				Name: core.StringPtr("hsm-client-cert-renamed"),
			}
			clientCertificatePatchModelAsPatch, asPatchErr := clientCertificatePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceHsmDestinationClientCertificateOptions := secretsManagerInstanceManagementService.NewUpdateInstanceHsmDestinationClientCertificateOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"7c9e6679-7425-40de-944b-e07fc1f90ae7",
				clientCertificatePatchModelAsPatch,
			)

			clientCertificate, response, err := secretsManagerInstanceManagementService.UpdateInstanceHsmDestinationClientCertificate(updateInstanceHsmDestinationClientCertificateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clientCertificate, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_hsm_destination_client_certificate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clientCertificate).ToNot(BeNil())
		})
		It(`DeleteVaultAdmintokens request example`, func() {
			// begin-delete_vault_admintokens

			deleteVaultAdmintokensOptions := secretsManagerInstanceManagementService.NewDeleteVaultAdmintokensOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
			)

			response, err := secretsManagerInstanceManagementService.DeleteVaultAdmintokens(deleteVaultAdmintokensOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteVaultAdmintokens(): %d\n", response.StatusCode)
			}

			// end-delete_vault_admintokens

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteInstanceDestination request example`, func() {
			fmt.Println("\nDeleteInstanceDestination() result:")
			// begin-delete_instance_destination

			deleteInstanceDestinationOptions := secretsManagerInstanceManagementService.NewDeleteInstanceDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementService.DeleteInstanceDestination(deleteInstanceDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ibmCloudDatabaseDestination, "", "  ")
			fmt.Println(string(b))

			// end-delete_instance_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(ibmCloudDatabaseDestination).ToNot(BeNil())
		})
		It(`DeleteInstanceHsmDestination request example`, func() {
			fmt.Println("\nDeleteInstanceHsmDestination() result:")
			// begin-delete_instance_hsm_destination

			deleteInstanceHsmDestinationOptions := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
			)

			asyncDeleteResponse, response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestination(deleteInstanceHsmDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(asyncDeleteResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_instance_hsm_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(asyncDeleteResponse).ToNot(BeNil())
		})
		It(`DeleteInstanceHsmDestinationHaGroup request example`, func() {
			// begin-delete_instance_hsm_destination_ha_group

			deleteInstanceHsmDestinationHaGroupOptions := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationHaGroupOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"ee52ebb6-1728-4580-8ede-13f6504e3ae0",
			)

			response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationHaGroup(deleteInstanceHsmDestinationHaGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteInstanceHsmDestinationHaGroup(): %d\n", response.StatusCode)
			}

			// end-delete_instance_hsm_destination_ha_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteInstanceHsmDestinationPath request example`, func() {
			fmt.Println("\nDeleteInstanceHsmDestinationPath() result:")
			// begin-delete_instance_hsm_destination_path

			deleteInstanceHsmDestinationPathOptions := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationPathOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
			)

			path, response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPath(deleteInstanceHsmDestinationPathOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(path, "", "  ")
			fmt.Println(string(b))

			// end-delete_instance_hsm_destination_path

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(path).ToNot(BeNil())
		})
		It(`DeleteInstanceHsmDestinationPathServerCa request example`, func() {
			// begin-delete_instance_hsm_destination_path_server_ca

			deleteInstanceHsmDestinationPathServerCaOptions := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationPathServerCaOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"fa3d1234-0000-0000-0000-000000000001",
				"5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f",
			)

			response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationPathServerCa(deleteInstanceHsmDestinationPathServerCaOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteInstanceHsmDestinationPathServerCa(): %d\n", response.StatusCode)
			}

			// end-delete_instance_hsm_destination_path_server_ca

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteInstanceHsmDestinationClientCertificate request example`, func() {
			// begin-delete_instance_hsm_destination_client_certificate

			deleteInstanceHsmDestinationClientCertificateOptions := secretsManagerInstanceManagementService.NewDeleteInstanceHsmDestinationClientCertificateOptions(
				"fe5c5c27-3e44-474a-9484-55892e12d6d1",
				"b2c3d4e5-f6a7-8901-bcde-f12345678901",
				"7c9e6679-7425-40de-944b-e07fc1f90ae7",
			)

			response, err := secretsManagerInstanceManagementService.DeleteInstanceHsmDestinationClientCertificate(deleteInstanceHsmDestinationClientCertificateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteInstanceHsmDestinationClientCertificate(): %d\n", response.StatusCode)
			}

			// end-delete_instance_hsm_destination_client_certificate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
