package lib

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

type DeviceShortcut struct {
	VID         gousb.ID
	PID         gousb.ID
	VendorName  string
	ProductName string
	Out         gousb.OutEndpoint
}

func GetDevice(vid, pid gousb.ID) DeviceShortcut {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open any device with a given VID/PID using a convenience function.0x04b8, 0x0e27
	dev, err := ctx.OpenDeviceWithVIDPID(vid, pid)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	defer dev.Close()
	dev.SetAutoDetach(true)

	man, _ := dev.Manufacturer()
	pro, _ := dev.Product()

	// NOTE
	// Melhorar a parte de config, pegar automaticamente
	// pelo primeiro index da lista de configs
	// ou atraves de uma configuração passada na funcao
	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("Error getting the device config.\n%s", err.Error())
	}
	defer cfg.Close()

	i, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("Error getting the device interface.\n%s", err.Error())
	}

	epOut, err := i.OutEndpoint(1)
	if err != nil {
		log.Fatalf("Error getting the device output endpoint.\n%s", err.Error())
	}

	return DeviceShortcut{VID: vid, PID: pid, VendorName: man, ProductName: pro, Out: *epOut}
}

func GetDeviceByName(vendorName string, productName string) DeviceShortcut {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool { return true })
	if err != nil {
		log.Fatalf("Error getting the devices:\n%s", err.Error())
	}

	var selectedDev *gousb.Device

	for _, dev := range devs {
		defer dev.Close()
		devName, err := dev.Product()
		if err != nil {
			log.Fatalf("Error getting product name:\n%s", err.Error())
		}
		venName, err := dev.Manufacturer()

		if err != nil {
			log.Fatalf("Error getting vendor name:\n%s", err.Error())
		}
		if devName == productName && vendorName == venName {
			selectedDev = dev
		}
	}

	if selectedDev == nil {
		log.Fatalf("No one device with given name.")
	}

	selectedDev.SetAutoDetach(true)

	cfg, err := selectedDev.Config(1)
	if err != nil {
		log.Fatalf("Error getting the device config.\n%s", err.Error())
	}
	defer cfg.Close()

	i, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("Error getting the device interface.\n%s", err.Error())
	}

	epOut, err := i.OutEndpoint(1)
	if err != nil {
		log.Fatalf("Error getting the device output endpoint.\n%s", err.Error())
	}

	return DeviceShortcut{VID: selectedDev.Desc.Product, PID: selectedDev.Desc.Vendor, VendorName: vendorName, ProductName: productName, Out: *epOut}
}

func ListDevices() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool { return true })
	if err != nil {
		log.Fatalf("Error getting the devices:\n%s", err.Error())
	}

	for _, dev := range devs {
		defer dev.Close()
		devName, err := dev.Product()
		if err != nil {
			log.Fatalf("Error getting product name:\n%s", err.Error())
		}
		venName, err := dev.Manufacturer()

		if err != nil {
			log.Fatalf("Error getting vendor name:\n%s", err.Error())
		}
		fmt.Printf("%-45s %s\n", venName, devName)
	}
}
