package gohlml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	cnt, err := DeviceCount()

	assert.NotNil(t, err, "Error should be raised when HLML isn't enabled")

	err = Initialize()
	assert.Nil(t, err, err)

	cnt, err = DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestRedundantInitialize(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	err = Initialize()
	assert.NotNil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestInitWithLogs(t *testing.T) {
	err := InitWithLogs()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestDeviceHandleByIndex(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	idx, err := dev.MinorNumber()
	assert.Nil(t, err, "Should be able to get device index")
	assert.Equal(t, uint(0), idx, "Index of device 0 should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestGetDeviceByUUID(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	uuid, err := dev.UUID()
	assert.Nil(t, err, "Should be able to get device UUID")

	dev2, err := DeviceHandleByUUID(uuid)
	assert.Nil(t, err, "Should be able to get device UUID")

	devMinor, err := dev.MinorNumber()
	assert.Nil(t, err, "Should be able to get minor number ")

	devMinor2, err := dev2.MinorNumber()
	assert.Nil(t, err, "Should be able to get minor number ")

	assert.Equal(t, devMinor, devMinor2, "Query by idx or UUID should give same device")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestGetDeviceName(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	name, err := dev.Name()
	assert.Nil(t, err, "Should be able to get device UUID")
	assert.Greater(t, len(name), 0, "Name should have a length")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIFunctions(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	_, err = dev.PCIDomain()
	assert.Nil(t, err, "Should be able to get PCI domain")

	_, err = dev.PCIBus()
	assert.Nil(t, err, "Should be able to get PCI bus")

	busID, err := dev.PCIBusID()
	assert.Nil(t, err, "Should be able to get PCI busID")
	assert.Greater(t, len(busID), 0, "busID should have a length")

	_, err = dev.PCIID()
	assert.Nil(t, err, "Should be able to get PCIID")

	_, err = dev.PCILinkSpeed()
	assert.Nil(t, err, "Should be able to get PCILinkSpeed")

	_, err = dev.PCILinkWidth()
	assert.Nil(t, err, "Should be able to get PCILinkSpeed")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestMemoryMetrics(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	total, used, free, err := dev.MemoryInfo()
	assert.Nil(t, err, "Should be able to get MemoryInfo")

	assert.GreaterOrEqual(t, free, uint64(0), "free.bytes value cannot be less than 0")
	assert.LessOrEqual(t, free, uint64(7e10), "free.bytes value cannot be more than 7e10")

	assert.GreaterOrEqual(t, used, uint64(0), "used.bytes value cannot be less than 0")
	assert.LessOrEqual(t, used, uint64(7e10), "used.bytes value cannot be more than 7e10")

	assert.GreaterOrEqual(t, total, uint64(0), "total.bytes value cannot be less than 0")
	assert.LessOrEqual(t, total, uint64(7e10), "total.bytes value cannot be more than 7e10")

	assert.Equal(t, free+used, total, "total.bytes must equal free + used")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestUtilizationInfo(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	_, err = dev.UtilizationInfo()
	assert.Nil(t, err, "Should be able to get UtilizationInfo")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestClockMetrics(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	freqSoc, socErr := dev.SOCClockInfo()
	freqIC, icErr := dev.ICClockInfo()
	freqMME, mmeErr := dev.MMEClockInfo()
	freqTPC, tpcErr := dev.TPCClockInfo()
	if socErr != nil && icErr != nil && mmeErr != nil && tpcErr != nil {
		assert.Equal(t, false, true, "Error retrieving any clock")
	}

	if socErr == nil {
		assert.GreaterOrEqual(t, freqSoc, uint(1000), "soc frequency too low")
		assert.LessOrEqual(t, freqSoc, uint(3000), "soc frequency too high")
	}

	if icErr == nil {
		assert.GreaterOrEqual(t, freqIC, uint(50), "ic frequency too low")
		assert.LessOrEqual(t, freqIC, uint(3000), "ic frequency too high")
	}

	if mmeErr == nil {
		assert.GreaterOrEqual(t, freqMME, uint(50), "mme frequency too low")
		assert.LessOrEqual(t, freqMME, uint(3000), "mme frequency too high")
	}

	if tpcErr == nil {
		assert.GreaterOrEqual(t, freqTPC, uint(50), "tpc frequency too low")
		assert.LessOrEqual(t, freqTPC, uint(3000), "tpc frequency too high")
	}

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPowerUsage(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	power, err := dev.PowerUsage()
	assert.Nil(t, err, "Should be able to get power")
	assert.GreaterOrEqual(t, power, uint(5000), "power temperature value cannot be less than 0")
	assert.LessOrEqual(t, power, uint(400000), "power temperature value cannot be more than 80")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestTemperature(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	onchip, onboard, err := dev.Temperature()
	assert.Nil(t, err, "Should be able to get temperature")
	assert.GreaterOrEqual(t, onchip, uint(30), "onchip temperature value cannot be less than 0")
	assert.LessOrEqual(t, onchip, uint(80), "onchip temperature value cannot be more than 80")

	// onboard check
	assert.GreaterOrEqual(t, onboard, uint(30), "onboard temperature value cannot be less than 0")
	assert.LessOrEqual(t, onboard, uint(80), "onboard temperature value cannot be more than 80")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestECCVolatileErrors(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	eccErr, err := dev.ECCVolatileErrors()
	assert.Nil(t, err, "Should be able to get ecc volatile errors")
	assert.Equal(t, uint64(0), eccErr, "ECCVolatileErrors value should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestECCAggregateErrors(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	eccErr, err := dev.ECCAggregateErrors()
	assert.Nil(t, err, "Should be able to get ecc aggregate errors")
	assert.Equal(t, uint64(0), eccErr, "ECCVolatileErrors value should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestHLRevision(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	rev, err := dev.HLRevision()
	assert.Nil(t, err, "Should be able to get HL revision")
	assert.Equal(t, int(0), rev, "HLRevision should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCBVersion(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	pcbVer, err := dev.PCBVersion()
	assert.Nil(t, err, "Should be able to get PCB Version")
	assert.Greater(t, len(pcbVer), 2)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCBAssemblyVersion(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	pcbVer, err := dev.PCBAssemblyVersion()
	assert.Nil(t, err, "Should be able to get PCB Assembly Version")
	assert.Greater(t, len(pcbVer), 2)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestSerialNumber(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	serial, err := dev.SerialNumber()
	assert.Nil(t, err, "Should be able to get serial number")
	assert.Greater(t, len(serial), 8)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestBoardID(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	_, err := dev.BoardID()
	assert.Nil(t, err, "Should be able to get board id")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIeTX(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	_, err := dev.PCIeTX()
	assert.Nil(t, err, "Should be able to get pcie transmit id")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIeRX(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	_, err := dev.PCIeRX()
	assert.Nil(t, err, "Should be able to get pcie receive id")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIReplayCounter(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	replayCnt, err := dev.PCIReplayCounter()
	assert.Nil(t, err, "Should be able to get pcie replay count")
	assert.Equal(t, uint(0), replayCnt, "PCIReplayCounter should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIeLinkGeneration(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	gen, err := dev.PCIeLinkGeneration()
	assert.Nil(t, err, "Should be able to get pcie link generation")
	assert.Equal(t, int(4), gen, "PCIeLinkGeneration should be 4") // maybe 5

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIeLinkWidth(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	width, err := dev.PCIeLinkWidth()
	assert.Nil(t, err, "Should be able to get pcie link width")
	assert.Equal(t, int(4), width, "PCIeLinkWidth should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestEnergyConsumptionCounter(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	energy, err := dev.EnergyConsumptionCounter()
	assert.Nil(t, err, "Should be able to get pcie link width")
	assert.Greater(t, energy, uint64(0), "EnergyConsumptionCounter should be > 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestStaticInfo(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	kernel, uboot, err := FWVersion(0)
	assert.Nil(t, err, "Should be able to get FWVersion")
	assert.Greater(t, len(kernel), 10, "kernel version too short")
	assert.Greater(t, len(uboot), 5, "uboot version too short")

	ver, err := SystemDriverVersion()
	assert.Nil(t, err, "Should be able to get SystemDriverVersion")
	assert.Greater(t, len(ver), 7, "driver version too short")

	err = Shutdown()
	assert.Nil(t, err, err)
}
