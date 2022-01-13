package adapter

import "testing"

func TestAdapterInherit(t *testing.T) {
	p := PrinterBanner{Banner{"hello"}}
	p.PrintStrong()
	p.PrintWeak()
}

func TestAdapterTransfer(t *testing.T) {
	p := PPrinterBanner{b: Banner{"hello"}}
	p.PrintStrong()
	p.PrintWeak()
}
