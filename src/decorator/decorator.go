package main

import "fmt"

// PS5 产品接口
type PS5 interface {
	StartGPUEngine()
	GetPrice() int64
}

// CD 版 PS5主机
type PS5WithCD struct{}

func (p PS5WithCD) StartGPUEngine() {
	fmt.Println("start engine")
}
func (p PS5WithCD) GetPrice() int64 {
	return 5000
}

// PS5 数字版主机
type PS5WithDigital struct{}

func (p PS5WithDigital) StartGPUEngine() {
	fmt.Println("start normal gpu engine")
}

func (p PS5WithDigital) GetPrice() int64 {
	return 3600
}

type PS5MachinePlus struct {
	ps5Machine PS5
}

func (p *PS5MachinePlus) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5MachinePlus) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("start plus plugin")
}

func (p PS5MachinePlus) GetPrice() int64 {
	return p.ps5Machine.GetPrice() + 500
}

// 主题色版 PS5 主机
type PS5WithTopicColor struct {
	ps5Machine PS5
}

func (p *PS5WithTopicColor) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5WithTopicColor) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("尊贵的主题色主机，GPU启动")
}
func (p PS5WithTopicColor) GetPrice() int64 {
	return p.ps5Machine.GetPrice() + 200
}

func main() {
	ps5MachinePlus := PS5MachinePlus{}
	ps5MachinePlus.SetPS5Machine(PS5WithCD{})
	// ps5MachinePlus.SetPS5Machine(PS5WithDigital{}) // 可以在更换主机
	ps5MachinePlus.StartGPUEngine()
	price := ps5MachinePlus.GetPrice()
	fmt.Printf("PS5 CD 豪华Plus版，价格: %d 元\n\n", price)

	ps5WithTopicColor := PS5WithTopicColor{}
	ps5WithTopicColor.SetPS5Machine(ps5MachinePlus)
	ps5WithTopicColor.StartGPUEngine()
	price = ps5WithTopicColor.GetPrice()
	fmt.Printf("PS5 CD 豪华Plus 经典主题配色版，价格: %d 元\n", price)
}
