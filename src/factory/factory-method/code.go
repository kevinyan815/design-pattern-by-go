package main

import "fmt"

// MathOperator 实际产品实现的接口--表示数学运算器应该有哪些行为
type MathOperator interface {
	SetOperandA(int)
	SetOperandB(int)
	ComputeResult() int
}

// OperatorFactory 工厂接口，由具体工厂类来实现
type OperatorFactory interface {
	Create() MathOperator
}

// BaseOperator 是所有 Operator 的基类
// 封装公用方法，因为Go不支持继承，具体Operator类
// 只能组合它来实现类似继承的行为表现。
type BaseOperator struct {
	operandA, operandB int
}

func (o *BaseOperator) SetOperandA(operand int) {
	o.operandA = operand
}

func (o *BaseOperator) SetOperandB(operand int) {
	o.operandB = operand
}

//PlusOperatorFactory 是 PlusOperator 加法运算器的工厂类
type PlusOperatorFactory struct{}

func (pf *PlusOperatorFactory) Create() MathOperator {
	return &PlusOperator{
		BaseOperator: &BaseOperator{},
	}
}

//PlusOperator 实际的产品类--加法运算器
type PlusOperator struct {
	*BaseOperator
}

//ComputeResult 计算并获取结果
func (p *PlusOperator) ComputeResult() int {
	return p.operandA + p.operandB
}

// MultiOperatorFactory 是乘法运算器产品的工厂
type MultiOperatorFactory struct {}

func (mf *MultiOperatorFactory) Create() MathOperator{
	return &MultiOperator{
		BaseOperator: &BaseOperator{},
	}
}

// MultiOperator 实际的产品类--乘法运算器
type MultiOperator struct {
	*BaseOperator
}
func (m *MultiOperator) ComputeResult() int {
	return m.operandA * m.operandB
}


// 测试运行
func main() {
	var factory OperatorFactory
	var mathOp MathOperator
	factory = &PlusOperatorFactory{}
	mathOp = factory.Create()
	mathOp.SetOperandB(3)
	mathOp.SetOperandA(2)
	fmt.Printf("Plus operation reuslt: %d\n", mathOp.ComputeResult())

	factory= &MultiOperatorFactory{}
	mathOp = factory.Create()
	mathOp.SetOperandB(3)
	mathOp.SetOperandA(2)
	fmt.Printf("Multiple operation reuslt: %d\n", mathOp.ComputeResult())
}
