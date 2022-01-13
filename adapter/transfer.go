package adapter

import "fmt"

type PPrint interface {
	PrintWeak()
	PrintStrong()
}

type BBanner struct {
	content string
}

func (b BBanner) printWithParen() {
	fmt.Println("(", b.content, ")")
}

func (b BBanner) printWithAster() {
	fmt.Println("*", b.content, "*")
}

type PPrinterBanner struct {
	b Banner
}

func (p PPrinterBanner) PrintWeak() {
	p.b.printWithParen() // 通过原代码实例来调用原方法
}

func (p PPrinterBanner) PrintStrong() {
	p.b.printWithAster()
}
