package gohlml

import (
	"testing"
	"log"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	cnt, err := DeviceCount()

	assert.NotNil(t, err, "Error should be raised when HLML isn't enabled")

	start := time.Now()
	err = Initialize()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "TestInitialize()", duration)
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

	start := time.Now()
	err := InitWithLogs()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "TestInitWithLogs()", duration)
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

	start := time.Now()
	dev, err := DeviceHandleByIndex(0)
	duration := time.Since(start)
	log.Printf("%v: %v\n", "TestDeviceHandleByIndex()", duration)
	assert.Nil(t, err, "Should be able to get device handle")

	start = time.Now()
	idx, err := dev.MinorNumber()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "MinorNumber()", duration)
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

	start := time.Now()
	uuid, err := dev.UUID()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "UUID()", duration)

	assert.Nil(t, err, "Should be able to get device UUID")

	start = time.Now()
	dev2, err := DeviceHandleByUUID(uuid)
	duration = time.Since(start)
	log.Printf("%v: %v\n", "DeviceHandleByUUID()", duration)

	assert.Nil(t, err, "Should be able to get device UUID")

	start = time.Now()
	devMinor, err := dev.MinorNumber()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "MinorNumber()", duration)

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

	start := time.Now()
	name, err := dev.Name()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "Name()", duration)
	
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

	start := time.Now()
	_, err = dev.PCIDomain()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "PCIDomain()", duration)
	assert.Nil(t, err, "Should be able to get PCI domain")

	start = time.Now()
	_, err = dev.PCIBus()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "PCIBus()", duration)
	assert.Nil(t, err, "Should be able to get PCI bus")

    start = time.Now()
	busID, err := dev.PCIBusID()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "PCIBusID()", duration)
	assert.Nil(t, err, "Should be able to get PCI busID")
	assert.Greater(t, len(busID), 0, "busID should have a length")

    start = time.Now()
	_, err = dev.PCIID()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "PCIID()", duration)
	assert.Nil(t, err, "Should be able to get PCIID")

    start = time.Now()
	_, err = dev.PCILinkSpeed()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "PCILinkSpeed()", duration)
	assert.Nil(t, err, "Should be able to get PCILinkSpeed")

    start = time.Now()
	_, err = dev.PCILinkWidth()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "PCILinkWidth()", duration)
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

	start := time.Now()
	total, used, free, err := dev.MemoryInfo()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "MemoryInfo()", duration)
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

	start := time.Now()
	_, err = dev.UtilizationInfo()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "UtilizationInfo()", duration)
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

	start := time.Now()
	freqSoc, socErr := dev.SOCClockInfo()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "SOCClockInfo()", duration)

	start = time.Now()
	freqIC, icErr := dev.ICClockInfo()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "ICClockInfo()", duration)

	start = time.Now()
	freqMME, mmeErr := dev.MMEClockInfo()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "MMEClockInfo", duration)

	start = time.Now()
	freqTPC, tpcErr := dev.TPCClockInfo()
	duration = time.Since(start)
	log.Printf("%v: %v\n", "TPCClockInfo()", duration)


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

	start := time.Now()
	power, err := dev.PowerUsage()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "PowerUsage()", duration)
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

	start := time.Now()
	onchip, onboard, err := dev.Temperature()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "Temperature()", duration)

	assert.Nil(t, err, "Should be able to get temperature")
	assert.GreaterOrEqual(t, onchip, uint(30), "onchip temperature value cannot be less than 0")
	assert.LessOrEqual(t, onchip, uint(80), "onchip temperature value cannot be more than 80")

	// onboard check
	assert.GreaterOrEqual(t, onboard, uint(25), "onboard temperature value cannot be less than 0")
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

	start := time.Now()
	eccErr, err := dev.ECCVolatileErrors()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "ECCVolatileErrors()", duration)
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

	start := time.Now()
	eccErr, err := dev.ECCAggregateErrors()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "ECCAggregateErrors()", duration)
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

	start := time.Now()
	rev, err := dev.HLRevision()
	duration := time.Since(start)
	log.Printf("%v: %v\n", "HLRevision()", duration)
	assert.Nil(t, err, "Should be able to get HL revision")
	assert.Equal(t, rev, int(3), "HLRevision should be 3")

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

	start := time.Now()
	pcbVer, err := dev.PCBVersion()
	duration := time.Since(start)
	log.Printf("%v: %v", "PCBVersion", duration)
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

	start := time.Now()
	pcbVer, err := dev.PCBAssemblyVersion()
	duration := time.Since(start)
	log.Printf("%v: %v", "PCBAssemblyVersion()", duration)
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

	start := time.Now()
	serial, err := dev.SerialNumber()
	duration := time.Since(start)
	log.Printf("%v: %v", "SerialNumber()", duration)
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

	start := time.Now()
	_, err = dev.BoardID()
	duration := time.Since(start)
	log.Printf("%v: %v", "BoardID()", duration)
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

	start := time.Now()
	_, err = dev.PCIeTX()
	duration := time.Since(start)
	log.Printf("%v: %v", "PCIeTX()", duration)
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

	start := time.Now()
	_, err = dev.PCIeRX()
	duration := time.Since(start)
	log.Printf("%v: %v", "PCIeRX()", duration)
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

	start := time.Now()
	replayCnt, err := dev.PCIReplayCounter()
	duration := time.Since(start)
	log.Printf("%v: %v", "PCIReplayCounter()", duration)
	assert.Nil(t, err, "Should be able to get pcie replay count")
	assert.Equal(t, uint(0), replayCnt, "PCIReplayCounter should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

/*
func TestPCIeLinkGeneration(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	start := time.Now()
	gen, err := dev.PCIeLinkGeneration()
	duration := time.Since(start)
	log.Printf("%v: %v", "PCIeLinkGeneration()", duration)
	assert.Nil(t, err, "Should be able to get pcie link generation")
	assert.Equal(t, int(4), gen, "PCIeLinkGeneration should be 4") // maybe 5

	err = Shutdown()
	assert.Nil(t, err, err)
}
*/
/*
func TestPCIeLinkWidth(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	start := time.Now()
	width, err := dev.PCIeLinkWidth()
	duration := time.Since(start)
	log.Printf("%v: %v", "PCIeLinkWidth()", duration)
	assert.Nil(t, err, "Should be able to get pcie link width")
	assert.Equal(t, int(16), width, "PCIeLinkWidth should be 16")

	err = Shutdown()
	assert.Nil(t, err, err)
}
*/
func TestEnergyConsumptionCounter(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")

	start := time.Now()
	energy, err := dev.EnergyConsumptionCounter()
	duration := time.Since(start)
	log.Printf("%v: %v", "EnergyConsumptionCounter()", duration)
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

	start := time.Now()
	kernel, uboot, err := FWVersion(0)
	duration := time.Since(start)
	log.Printf("%v: %v", "FWVersion()", duration)
	assert.Nil(t, err, "Should be able to get FWVersion")
	assert.Greater(t, len(kernel), 10, "kernel version too short")
	assert.Greater(t, len(uboot), 5, "uboot version too short")

	start = time.Now()
	ver, err := SystemDriverVersion()
	duration = time.Since(start)
	log.Printf("%v: %v", "SystemDriverVersion()", duration)
	assert.Nil(t, err, "Should be able to get SystemDriverVersion")
	assert.Greater(t, len(ver), 7, "driver version too short")

	err = Shutdown()
	assert.Nil(t, err, err)
}

