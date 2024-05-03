// Copyright 2017 the gousb Authors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

func main() {
	// Initialize a new Context.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open any device with a given VID/PID using a convenience function.
	dev, err := ctx.OpenDeviceWithVIDPID(0x04b8, 0x0e27)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	defer dev.Close()

	dev.SetAutoDetach(true)

	// Claim the default interface using a convenience function.
	// The default interface is always #0 alt #0 in the currently active
	// config.

	man, _ := dev.Manufacturer()
	pro, _ := dev.Product()

	fmt.Printf("%s %s\n", man, pro)

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

	data := []byte("Sua mensagem aqui\n")

	bytesWritten, err := epOut.Write(data)
	if err != nil {
		log.Fatalf("Erro ao escrever dados: %v", err)
	}
	fmt.Printf("%d bytes enviados para o dispositivo\n", bytesWritten)
}
