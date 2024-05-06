package lib

import (
	"log"

	"github.com/google/gousb"
)

type DevideShortcut struct {
	VID         gousb.ID
	PID         gousb.ID
	VendorName  string
	ProductName string
	Out         gousb.OutEndpoint
}

func (p DevideShortcut) getting() {

}

func GetDevice(vid, pid gousb.ID) DevideShortcut {
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
		log.Fatalf("Error until getting device config.\n%s", err.Error())
	}
	defer cfg.Close()

	i, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("Error until getting device interface.\n%s", err.Error())
	}

	epOut, err := i.OutEndpoint(1)
	if err != nil {
		log.Fatalf("Error until getting device output endpoint.\n%s", err.Error())
	}

	return DevideShortcut{VID: vid, PID: pid, VendorName: man, ProductName: pro, Out: *epOut}
}
