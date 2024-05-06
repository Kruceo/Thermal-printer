package lib

import "github.com/google/gousb"

/* CharacterSetBytes are just []bytes, this name is used for better types */
type CharacterSetBytes []byte

/* CommandBytes are just []bytes, this name is used for better types */
type CommandBytes []byte

type Printer struct {
	OutEndpoint gousb.OutEndpoint
	FeedLines   func(n int)
	FullCut     func()
	Clear       func()
}

func (p Printer) Write(content []byte) error {
	_, err := p.OutEndpoint.Write(content)
	return err
}
