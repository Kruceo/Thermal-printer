package epson

import (
	"fmt"
	"thermal-printer/lib"

	"github.com/google/gousb"
)

func feedLines(epOut gousb.OutEndpoint, n int) {
	argIndex := len(FeedLines) - 1

	var command = make(lib.CommandBytes, len(FeedLines))
	copy(command, FeedLines)

	command[argIndex] = byte(n)
	for _, v := range command {
		fmt.Printf("%X ", v)
	}
	epOut.Write(command)
}

func CreateEpsonPrinter(outEndpoint gousb.OutEndpoint) lib.Printer {

	printer := lib.Printer{
		OutEndpoint: outEndpoint,
		FeedLines:   func(n int) { feedLines(outEndpoint, n) },
		FullCut:     func() { outEndpoint.Write(FullCut) },
		Clear:       func() { outEndpoint.Write(Init) },
	}

	return printer

}
