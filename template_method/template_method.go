package template_method

type PaperWriter interface {
	Title()
	Content()
	Quote()
}

type Writer struct { // 借助该struct实现控制方法
	PaperWriter
}

func (w Writer) Write() {
	w.Title()
	w.Content()
	w.Quote()
}
