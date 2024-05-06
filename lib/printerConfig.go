package lib

type PrinterConfig struct {
	feedLines func(n int)
	cut       func()
	clear     func()
}
