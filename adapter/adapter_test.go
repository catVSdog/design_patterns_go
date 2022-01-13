package adapter

import "testing"

func TestAdapterInherit(t *testing.T) {
	var p Print // p 为接口类型
	p = PrinterBanner{Banner{"hello"}}
	p.PrintStrong()
	p.PrintWeak()
}

func TestAdapterTransfer(t *testing.T) {
	var p PPrint // p 为接口类型
	p = PPrinterBanner{b: Banner{"hello"}}
	p.PrintStrong()
	p.PrintWeak()
}
