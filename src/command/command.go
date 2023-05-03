package main

import "fmt"

// 命令接收者，负责逻辑的执行
type CPU struct{}

func (CPU) ADoSomething(param int) {
	fmt.Printf("a do something with param %v\n", param)
}
func (CPU) BDoSomething(param1 string, param2 int) {
	fmt.Printf("b do something with params %v and %v \n", param1, param2)
}
func (CPU) CDoSomething() {
	fmt.Println("c do something with no params")
}

// 接口中仅声明一个执行命令的方法 Execute()
type Command interface {
	Execute()
}

// 命令对象持有一个指向接收者的引用，以及请求中的所有参数，
type ACommand struct {
	cpu *CPU
	param int
}
// 命令不会进行逻辑处理，调用Execute方法会将发送者的请求委派给接收者对象。 
func (a ACommand) Execute() {
	a.cpu.ADoSomething(a.param)
	a.cpu.CDoSomething()// 可以执行多个接收者的操作完成命令宏
}

func NewACommand(cpu *CPU, param int) Command {
	return ACommand{cpu, param}
}

type BCommand struct {
	state bool // Command 里可以添加些状态用作逻辑判断
	cpu *CPU
	param1 string
	param2 int
}

func (b BCommand) Execute() {
	if b.state {
		return
	}
	b.cpu.BDoSomething(b.param1, b.param2)
	b.state = true
	b.cpu.CDoSomething()
}

func NewBCommand(cpu *CPU, param1 string, param2 int) Command {
	return BCommand{false,cpu, param1, param2}
}

type PS5 struct {
	commands map[string]Command
}

// SetCommand方法来将 Command 指令设定给PS5。
func (p *PS5) SetCommand(name string, command Command) {
	p.commands[name] = command
}
// DoCommand方法选择要执行的命令
func (p *PS5) DoCommand(name string) {
	p.commands[name].Execute()
}

func main() {
	cpu := CPU{}
    // main方法充当客户端，创建并配置具体命令对象, 完成命令与执行操作的接收者的关联。
	ps5 := PS5{make(map[string]Command)}
	ps5.SetCommand("a", NewACommand(&cpu, 1))
	ps5.SetCommand("b", NewBCommand(&cpu, "hello", 2))
	ps5.DoCommand("a")
	ps5.DoCommand("b")
}
