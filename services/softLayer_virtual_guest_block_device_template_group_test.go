package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	common "github.com/maximilien/softlayer-go/common"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Virtual_Guest_Service", func() {
	var (
		username, apiKey string
		err              error

		fakeClient *slclientfakes.FakeSoftLayerClient

		vgbdtgService softlayer.SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service

		vgbdtGroup datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		vgbdtgService, err = fakeClient.GetSoftLayer_Virtual_Guest_Block_Device_Template_Group_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(vgbdtgService).ToNot(BeNil())

		vgbdtGroup = datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group{}
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := vgbdtgService.GetName()
			Expect(name).To(Equal("SoftLayer_Virtual_Guest_Block_Device_Template_Group"))
		})
	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 200150
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_getObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Virtual_Guest_Block_Device_Template_Group instance", func() {
			vgbdtg, err := vgbdtgService.GetObject(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(vgbdtg.AccountId).To(Equal(278444))
			Expect(vgbdtg.CreateDate).ToNot(BeNil())
			Expect(vgbdtg.Id).To(Equal(vgbdtGroup.Id))
			Expect(vgbdtg.Name).To(Equal("BOSH-eCPI-packer-centos-2014-08-12T15:54:16Z"))
			Expect(vgbdtg.Note).To(Equal("centos image created by packer at 2014-08-12T15:54:16Z"))
			Expect(vgbdtg.ParentId).To(BeNil())
			Expect(vgbdtg.PublicFlag).To(Equal(0))
			Expect(vgbdtg.StatusId).To(Equal(1))
			Expect(vgbdtg.Summary).To(Equal("centos image created by packer at 2014-08-12T15:54:16Z"))
			Expect(vgbdtg.TransactionId).To(BeNil())
			Expect(vgbdtg.UserRecordId).To(Equal(239954))
			Expect(vgbdtg.GlobalIdentifier).To(Equal("8071601b-5ee1-483e-a9e8-6e5582dcb9f7"))
		})
	})

	Context("#DeleteObject", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_deleteObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully deletes the SoftLayer_Virtual_Guest_Block_Device_Template_Group instance", func() {
			transaction, err := vgbdtgService.DeleteObject(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(transaction.CreateDate).ToNot(BeNil())
			Expect(transaction.ElapsedSeconds).To(Equal(1))
			Expect(transaction.GuestId).To(Equal(1234567))
			Expect(transaction.HardwareId).To(Equal(0))
			Expect(transaction.Id).To(Equal(11878004))
			Expect(transaction.ModifyDate).ToNot(BeNil())
			Expect(transaction.StatusChangeDate).ToNot(BeNil())

			Expect(transaction.TransactionGroup).To(Equal(datatypes.TransactionGroup{}))

			Expect(transaction.TransactionStatus.AverageDuration).To(Equal(".42"))
			Expect(transaction.TransactionStatus.FriendlyName).To(Equal("Cloud Reclaim Prep"))
			Expect(transaction.TransactionStatus.Name).To(Equal("CLOUD_RECLAIM_PREP"))
		})
	})

	Context("#GetDatacenters", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_getDatacenters.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves an array of SoftLayer_Location array for virtual guest device template group", func() {
			locations, err := vgbdtgService.GetDatacenters(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(locations)).To(BeNumerically("==", 2))

			Expect(locations[0].Id).To(Equal(265592))
			Expect(locations[0].LongName).To(Equal("Amsterdam 1"))
			Expect(locations[0].Name).To(Equal("ams01"))

			Expect(locations[1].Id).To(Equal(154820))
			Expect(locations[1].LongName).To(Equal("Dallas 6"))
			Expect(locations[1].Name).To(Equal("dal06"))
		})
	})

	Context("#GetSshKeys", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_getSshKeys.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves an array of SoftLayer_Security_Ssh_Key array for virtual guest device template group", func() {
			sshKeys, err := vgbdtgService.GetSshKeys(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(sshKeys)).To(BeNumerically(">", 0))

			for _, sshKey := range sshKeys {
				Expect(sshKey.CreateDate).ToNot(BeNil())
				Expect(sshKey.Fingerprint).To(Equal("f6:c2:9d:57:2f:74:be:a1:db:71:f2:e5:8e:0f:84:7e"))
				Expect(sshKey.Id).To(Equal(84386))
				Expect(sshKey.Key).ToNot(Equal(""))
				Expect(sshKey.Label).To(Equal("TEST:softlayer-go"))
				Expect(sshKey.ModifyDate).To(BeNil())
				Expect(sshKey.Label).To(Equal("TEST:softlayer-go"))
			}
		})
	})

	Context("#GetStatus", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_getStatus.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Virtual_Guest_Block_Device_Template_Group instance status", func() {
			status, err := vgbdtgService.GetStatus(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(status.Description).To(Equal("The Guest Block Device Template Group is available to all accounts"))
			Expect(status.KeyName).To(Equal("ACTIVE"))
			Expect(status.Name).To(Equal("Active"))
		})
	})

	Context("#GetStorageLocations", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_getStorageLocations.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Locations for the Virtual_Guest_Block_Device_Template_Group instance", func() {
			locations, err := vgbdtgService.GetStorageLocations(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(locations)).To(Equal(18))

			found := false
			for _, location := range locations {
				if location.LongName == "Paris 1" {
					Expect(location.Id).To(Equal(449500))
					Expect(location.Name).To(Equal("par01"))
					found = true
					break
				}
			}
			Expect(found).To(BeTrue())
		})
	})

	Context("#GetImageType", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_getImageType.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves the image type for the instance", func() {
			imageType, err := vgbdtgService.GetImageType(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(imageType.Description).To(Equal("a disk that may be replaced on upgrade"))
			Expect(imageType.KeyName).To(Equal("SYSTEM"))
			Expect(imageType.Name).To(Equal("System"))
		})
	})

	Context("#GetImageTypeKeyName", func() {
		BeforeEach(func() {
			vgbdtGroup.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_getImageTypeKeyName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves the image type key name for the instance", func() {
			imageTypeKeyName, err := vgbdtgService.GetImageTypeKeyName(vgbdtGroup.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(imageTypeKeyName).To(Equal("SYSTEM"))
		})
	})

	Context("#CreateFromExternalSource", func() {
		var configuration datatypes.SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration

		BeforeEach(func() {
			configuration = datatypes.SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration{
				Name: "fake-configuration-name",
				Note: "fake-configuration-note",
				OperatingSystemReferenceCode: "fake-operating-system-reference-code",
				Uri: "swift://FakeObjectStorageAccountName>@fake-clusterName/fake-containerName/fake-fileName.vhd",
			}
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_createFromExternalSource.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully creates a virtual guest device templte group from a configuration from an external VHD image", func() {
			vgbdtGroup, err := vgbdtgService.CreateFromExternalSource(configuration)
			Expect(err).ToNot(HaveOccurred())
			Expect(vgbdtGroup.AccountId).To(Equal(278444))
			Expect(vgbdtGroup.CreateDate).ToNot(BeNil())
			Expect(vgbdtGroup.Id).To(Equal(211582))
			Expect(vgbdtGroup.Name).To(Equal(" ubuntu-10.04-bosh-2168-IEM-itcs104-dea-stemcell"))
			Expect(vgbdtGroup.Note).To(Equal("fake-note"))
			Expect(*vgbdtGroup.ParentId).To(Equal(211578))
			Expect(vgbdtGroup.PublicFlag).To(Equal(0))
			Expect(vgbdtGroup.StatusId).To(Equal(1))
			Expect(vgbdtGroup.Summary).To(Equal("fake-summary"))
			Expect(vgbdtGroup.TransactionId).To(BeNil())
			Expect(vgbdtGroup.UserRecordId).To(Equal(180816))
			Expect(vgbdtGroup.GlobalIdentifier).To(Equal("fake-global-identifier"))
		})
	})

	Context("#CopyToExternalSource", func() {
		var configuration datatypes.SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration

		BeforeEach(func() {
			configuration = datatypes.SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration{
				Name: "fake-configuration-name",
				Note: "fake-configuration-note",
				OperatingSystemReferenceCode: "fake-operating-system-reference-code",
				Uri: "swift://FakeObjectStorageAccountName>@fake-clusterName/fake-containerName/fake-fileName.vhd",
			}
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_copyToExternalSource.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully copies the virtual guest device template group to an external source", func() {
			copied, err := vgbdtgService.CopyToExternalSource(configuration)
			Expect(err).ToNot(HaveOccurred())
			Expect(copied).To(BeTrue())
		})
	})
})
