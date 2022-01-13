package adapter

import "fmt"

type Print interface {
	PrintWeak()
	PrintStrong()
}

type Banner struct {
	content string
}

func (b Banner) printWithParen() {
	fmt.Println("(", b.content, ")")
}

func (b Banner) printWithAster() {
	fmt.Println("*", b.content, "*")
}

type PrinterBanner struct {
	Banner
}

func (p PrinterBanner) PrintWeak() {
	p.printWithParen() // 直接调用原代码的方法
}

func (p PrinterBanner) PrintStrong() {
	p.printWithAster()
}
