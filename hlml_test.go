package gohlml

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	cnt, err := DeviceCount()
	assert.NotNil(t, err, "Error should be raised when HLML isn't enabled")

	start := time.Now()
	err = Initialize()
	printDuration("TestInitialize()", time.Since(start))
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
	printDuration("TestInitWithLogs()", time.Since(start))
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
	printDuration("TestDeviceHandleByIndex()", time.Since(start))
	assert.Nil(t, err, "Should be able to get device handle")

	start = time.Now()
	idx, err := dev.MinorNumber()
	printDuration("MinorNumber()", time.Since(start))
	assert.Nil(t, err, "Should be able to get device index")
	assert.Equal(t, uint(0), idx, "Index of device 0 should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestGetDeviceByUUID(t *testing.T) {
	dev := prepareDevice(t)
	start := time.Now()
	uuid, err := dev.UUID()
	printDuration("UUID()", time.Since(start))

	assert.Nil(t, err, "Should be able to get device UUID")

	start = time.Now()
	dev2, err := DeviceHandleByUUID(uuid)
	printDuration("DeviceHandleByUUID()", time.Since(start))

	assert.Nil(t, err, "Should be able to get device UUID")

	start = time.Now()
	devMinor, err := dev.MinorNumber()
	printDuration("MinorNumber()", time.Since(start))

	assert.Nil(t, err, "Should be able to get minor number ")

	devMinor2, err := dev2.MinorNumber()
	assert.Nil(t, err, "Should be able to get minor number ")

	assert.Equal(t, devMinor, devMinor2, "Query by idx or UUID should give same device")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestGetDeviceBySerial(t *testing.T) {
	dev := prepareDevice(t)

	serial, err := dev.SerialNumber()
	assert.Nil(t, err, "Should be able to get device serial")

	dev2, err := DeviceHandleBySerial(serial)
	assert.Nil(t, err, "Should be able to get device by serial")

	dev2Serial, err := dev2.SerialNumber()
	assert.Nil(t, err, "Should be able to get device serial")
	assert.Equal(t, serial, dev2Serial, "The serial numbers should match")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestGetDeviceName(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	name, err := dev.Name()
	printDuration("Name()", time.Since(start))

	assert.Nil(t, err, "Should be able to get device UUID")
	assert.Greater(t, len(name), 0, "Name should have a length")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIFunctions(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	_, err := dev.PCIDomain()
	printDuration("PCIDomain()", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCI domain")

	start = time.Now()
	_, err = dev.PCIBus()
	printDuration("PCIBus()", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCI bus")

	start = time.Now()
	busID, err := dev.PCIBusID()
	printDuration("PCIBusID()", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCI busID")
	assert.Greater(t, len(busID), 0, "busID should have a length")

	start = time.Now()
	_, err = dev.PCIID()
	printDuration("PCIID()", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCIID")

	start = time.Now()
	_, err = dev.PCILinkSpeed()
	printDuration("PCILinkSpeed()", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCILinkSpeed")

	start = time.Now()
	_, err = dev.PCILinkWidth()
	printDuration("PCILinkWidth()", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCILinkSpeed")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestMemoryMetrics(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	total, used, free, err := dev.MemoryInfo()
	printDuration("MemoryInfo()", time.Since(start))
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
	dev := prepareDevice(t)

	start := time.Now()
	_, err := dev.UtilizationInfo()
	printDuration("UtilizationInfo()", time.Since(start))
	assert.Nil(t, err, "Should be able to get UtilizationInfo")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestClockMetrics(t *testing.T) {
	dev := prepareDevice(t)

	// Current freq for clocks
	start := time.Now()
	freqSoc, socErr := dev.SOCClockInfo()
	printDuration("SOCClockInfo()", time.Since(start))

	// Max freq for clocks
	start = time.Now()
	maxFreqSoc, maxSocErr := dev.SOCClockMax()
	printDuration("SOCClockMax()", time.Since(start))

	if socErr != nil {
		assert.Equal(t, false, true, "Error retrieving any clock")
	}

	if maxSocErr != nil {
		assert.Equal(t, false, true, "Error retrieving any clock for maximum frequency")
	}

	if socErr == nil {
		assert.GreaterOrEqual(t, freqSoc, uint(1000), "soc frequency too low")
		assert.LessOrEqual(t, freqSoc, uint(maxFreqSoc), "soc frequency too high")
	}

	err := Shutdown()
	assert.Nil(t, err, err)
}

func TestPowerUsage(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	power, err := dev.PowerUsage()
	printDuration("PowerUsage()", time.Since(start))
	assert.Nil(t, err, "Should be able to get power")
	assert.GreaterOrEqual(t, power, uint(5000), "power temperature value cannot be less than 0")
	assert.LessOrEqual(t, power, uint(400000), "power temperature value cannot be more than 80")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestTemperatureOnBoard(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	onboard, err := dev.TemperatureOnBoard()
	printDuration("Temperature()", time.Since(start))

	assert.Nil(t, err, "Should be able to get temperature")

	assert.GreaterOrEqual(t, onboard, uint(25), "onboard temperature value cannot be less than 0")
	assert.LessOrEqual(t, onboard, uint(80), "onboard temperature value cannot be more than 80")

	err = Shutdown()
	assert.Nil(t, err, err)
}
func TestTemperatureOnChip(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	onchip, err := dev.TemperatureOnChip()
	printDuration("Temperature()", time.Since(start))

	assert.Nil(t, err, "Should be able to get temperature")
	assert.GreaterOrEqual(t, onchip, uint(0), "onchip temperature value cannot be less than 0")
	assert.LessOrEqual(t, onchip, uint(80), "onchip temperature value cannot be more than 80")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestTemperatureThresholds(t *testing.T) {
	dev := prepareDevice(t)
	testCases := []struct {
		name string
		test func() (uint, error)
	}{
		{name: "Shutdown Temperature Threshold", test: dev.TemperatureThresholdShutdown},
		{name: "Slowdown Temperature Threshold", test: dev.TemperatureThresholdSlowdown},
		{name: "Memory Temperature Threshold", test: dev.TemperatureThresholdMemory},
		{name: "GPU Temperature Threshold", test: dev.TemperatureThresholdGPU},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			temp, err := tc.test()
			printDuration(tc.name, time.Since(start))
			assert.Nil(t, err, "Should not get an error when getting the temperature threshold")
			assert.Greater(t, temp, uint(0), "Threshold should be greater than 0")
		})
	}
	err := Shutdown()
	assert.Nil(t, err, err)
}

func TestPowerManagementDefaultLimit(t *testing.T) {
	dev := prepareDevice(t)

	limit, err := dev.PowerManagementDefaultLimit()
	assert.Nil(t, err, "Should be able to get the power management default limit")
	assert.Greater(t, limit, uint(0), "limit should be greater than 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestECCMode(t *testing.T) {
	dev := prepareDevice(t)

	currentECC, pendingECC, err := dev.ECCMode()
	assert.Nil(t, err, "Should be able to get the ecc mode")
	assert.GreaterOrEqual(t, currentECC, uint(0), "current ECC Mode should be Greater or equal than 0")
	assert.GreaterOrEqual(t, pendingECC, uint(0), "current ECC Mode should be Greater or equal than 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestHLRevision(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	rev, err := dev.HLRevision()
	printDuration("HLRevision()", time.Since(start))
	assert.Nil(t, err, "Should be able to get HL revision")
	assert.Equal(t, rev, int(3), "HLRevision should be 3")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCBVersion(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	pcbVer, err := dev.PCBVersion()
	printDuration("PCBVersion", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCB Version")
	assert.Greater(t, len(pcbVer), 2)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCBAssemblyVersion(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	pcbVer, err := dev.PCBAssemblyVersion()
	printDuration("PCBAssemblyVersion()", time.Since(start))
	assert.Nil(t, err, "Should be able to get PCB Assembly Version")
	assert.Greater(t, len(pcbVer), 2)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestSerialNumber(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	serial, err := dev.SerialNumber()
	printDuration("SerialNumber()", time.Since(start))
	assert.Nil(t, err, "Should be able to get serial number")
	assert.Greater(t, len(serial), 8)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestBoardID(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	_, err := dev.BoardID()
	printDuration("BoardID()", time.Since(start))
	assert.Nil(t, err, "Should be able to get board id")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIeTX(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	_, err := dev.PCIeTX()
	printDuration("PCIeTX()", time.Since(start))
	assert.Nil(t, err, "Should be able to get pcie transmit id")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIeRX(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	_, err := dev.PCIeRX()
	printDuration("PCIeRX()", time.Since(start))
	assert.Nil(t, err, "Should be able to get pcie receive id")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIReplayCounter(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	replayCnt, err := dev.PCIReplayCounter()
	printDuration("PCIReplayCounter()", time.Since(start))
	assert.Nil(t, err, "Should be able to get pcie replay count")
	assert.Equal(t, uint(0), replayCnt, "PCIReplayCounter should be 0")

	err = Shutdown()
	assert.Nil(t, err, err)
}

// TODO: PCIe Data collection is broken in v0.11

func TestPCIeLinkGeneration(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	gen, err := dev.PCIeLinkGeneration()
	printDuration("PCIeLinkGeneration()", time.Since(start))
	assert.Nil(t, err, "Should be able to get pcie link generation")
	assert.GreaterOrEqual(t, gen, uint(1), "PCIeLinkGeneration should be less than 5") // maybe 5
	assert.LessOrEqual(t, gen, uint(5), "PCIeLinkGeneration should be greater 1")      // maybe 5

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestPCIeLinkWidth(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	width, err := dev.PCIeLinkWidth()
	printDuration("PCIeLinkWidth()", time.Since(start))
	assert.Nil(t, err, "Should be able to get pcie link width")
	assert.Equal(t, uint(16), width, "PCIeLinkWidth should be 16")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestEnergyConsumptionCounter(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	energy, err := dev.EnergyConsumptionCounter()
	printDuration("EnergyConsumptionCounter()", time.Since(start))
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
	printDuration("FWVersion()", time.Since(start))
	assert.Nil(t, err, "Should be able to get FWVersion")
	assert.Greater(t, len(kernel), 10, "kernel version too short")
	assert.Greater(t, len(uboot), 5, "uboot version too short")

	start = time.Now()
	ver, err := SystemDriverVersion()
	printDuration("SystemDriverVersion()", time.Since(start))
	assert.Nil(t, err, "Should be able to get SystemDriverVersion")
	assert.Greater(t, len(ver), 7, "driver version too short")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestNewEventSet(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err, err)

	eventSet := NewEventSet()

	DeleteEventSet(eventSet)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func printDuration(msg string, duration time.Duration) {
	log.Printf("%v: %v", msg, duration)
}

func TestMacAddressInfo(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	ports, err := dev.MacAddressInfo()
	printDuration("MacAddressInfo()", time.Since(start))
	assert.Nil(t, err, "Should be able to get MacAddress info")
	assert.Equal(t, len(ports), int(10), "there should be 10 ports in a device")

	err = Shutdown()
	assert.Nil(t, err, err)
}

// func TestNicLinkStatus(t *testing.T) {
// 	dev := prepareDevice(t)

// 	start := time.Now()
// 	up, err := dev.NicLinkStatus(0)
// 	printDuration("NicLinkStatus()", time.Since(start))
// 	assert.Nil(t, err, "Should be able to get NicLinkStatus info")
// 	assert.GreaterOrEqual(t, up, uint(0), "Link status should be up (1) or down (0)")

// 	err = Shutdown()
// 	assert.Nil(t, err, err)
// }

// func TestThermalViolationStatus(t *testing.T) {
// 	dev := prepareDevice(t)

// 	start := time.Now()
// 	_, _, err := dev.ThermalViolationStatus()
// 	printDuration("TestThermalViolationStatus()", time.Since(start))
// 	assert.Nil(t, err, "Should be able to get thermal violation status")
// 	err = Shutdown()
// 	assert.Nil(t, err, err)
// }
// func TestPowerViolationStatus(t *testing.T) {
// 	dev := prepareDevice(t)

// 	start := time.Now()
// 	_, _, err := dev.PowerViolationStatus()
// 	printDuration("PowerViolationStatus()", time.Since(start))
// 	assert.Nil(t, err, "Should be able to get power violation status")

// 	err = Shutdown()
// 	assert.Nil(t, err, err)
// }

func TestReplacedRowDoubleBitECC(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	count, err := dev.ReplacedRowDoubleBitECC()
	printDuration("ReplacedRowDoubleBitECC()", time.Since((start)))
	assert.Nil(t, err, "Should be able to get the rows with double-bit errors")
	assert.Equal(t, uint(0), count, "Expected 0 rows with error, got %d", count)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestReplacedRowSingleBitECC(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	count, err := dev.ReplacedRowSingleBitECC()
	printDuration("ReplacedRowSingleBitECC()", time.Since((start)))
	assert.Nil(t, err, "Should be able to get the rows with single-bit errors")
	assert.Equal(t, uint(0), count, "Expected 0 rows with error, got %d", count)

	err = Shutdown()
	assert.Nil(t, err, err)
}

func TestIsReplacedRowsPendingStatus(t *testing.T) {
	dev := prepareDevice(t)

	start := time.Now()
	isPending, err := dev.IsReplacedRowsPendingStatus()
	printDuration("isReplacedRowsPendingStatus()", time.Since(start))
	assert.Nil(t, err, "Should be able to get pending rows status")
	assert.IsType(t, int(1), isPending, "Function should return a bool")

	err = Shutdown()
	assert.Nil(t, err, err)
}

func prepareDevice(t *testing.T) Device {
	err := Initialize()
	assert.Nil(t, err, err)

	cnt, err := DeviceCount()

	assert.Nil(t, err, "Error should not be raised when HLML is initialized")
	assert.Greater(t, cnt, uint(0), "Should detect at least 1 device")

	dev, err := DeviceHandleByIndex(0)
	assert.Nil(t, err, "Should be able to get device handle")
	return dev
}
