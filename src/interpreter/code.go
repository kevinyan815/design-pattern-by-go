package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Expression interface {
	Interpret() int
}

type NumberExpression struct {
	val int
}

func (n *NumberExpression) Interpret() int {
	return n.val
}

type AdditionExpression struct {
	left, right Expression
}

func (n *AdditionExpression) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

type SubtractionExpression struct {
	left, right Expression
}

func (n *SubtractionExpression) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Expression
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAdditionExpression()
		case "-":
			p.prev = p.newSubtractionExpression()
		default:
			p.prev = p.newNumberExpression()
		}
	}
}

func (p *Parser) newAdditionExpression() Expression {
	p.index++
	return &AdditionExpression{
		left:  p.prev,
		right: p.newNumberExpression(),
	}
}

func (p *Parser) newSubtractionExpression() Expression {
	p.index++
	return &SubtractionExpression{
		left:  p.prev,
		right: p.newNumberExpression(),
	}
}

func (p *Parser) newNumberExpression() Expression {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &NumberExpression{
		val: v,
	}
}

func (p *Parser) Result() Expression {
	return p.prev
}

func main() {
	p := &Parser{}
	p.Parse("1 + 3 + 3 + 3 - 3")
	res := p.Result().Interpret()
	expect := 7
	if res != expect {
		log.Fatalf("error: expect %d got %d", expect, res)
	}

	fmt.Printf("expect: %d, got: %d", expect, res)
}
