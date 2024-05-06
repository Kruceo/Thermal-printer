package epson

import (
	"thermal-printer/lib"

	"github.com/google/gousb"
)

func feedLines(epOut gousb.OutEndpoint, n int) {
	argIndex := len(FeedLines) - 1

	var command = make(lib.CommandBytes, len(FeedLines))
	copy(command, FeedLines)

	command[argIndex] = byte(n)

	epOut.Write(command)
}

func CreateEpsonPrinter(outEndpoint gousb.OutEndpoint, charset lib.CharacterSetBytes) lib.Printer {

	printer := lib.Printer{
		OutEndpoint: outEndpoint,
		FeedLines:   func(n int) { feedLines(outEndpoint, n) },
		FullCut:     func() { outEndpoint.Write(FullCut) },
		Clear: func() {
			outEndpoint.Write(Init)
			outEndpoint.Write(charset)
		},
	}

	return printer

}
