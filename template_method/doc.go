package template_method

// 模版方法

// 在父类中定义处理流程的框架，在子类中实现具体处理方式
// 类似 Builder 模式，但是模版方法更简单，更直接——父类直接控制子类

// 但是 go 中 interface 只支持定义抽象方法，所以，需要另外一个struct来实现控制流程
