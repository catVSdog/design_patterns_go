package template_method

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	writePageA := Writer{PageA{}} // 通过一个结构体内嵌不同的接口实现来达到不同的效果
	writePageA.Write()

	writePageB := Writer{PageB{}}
	writePageB.Write()
}

type PageA struct {
	Writer
}

func (p PageA) Title() {
	fmt.Println("write Title A")
}

func (p PageA) Content() {
	fmt.Println("write content A")
}

func (p PageA) Quote() {
	fmt.Println("write quote A")
}

type PageB struct{}

func (p PageB) Title() {
	fmt.Println("write Title B")
}

func (p PageB) Content() {
	fmt.Println("write content B")
}

func (p PageB) Quote() {
	fmt.Println("write quote B")
}
