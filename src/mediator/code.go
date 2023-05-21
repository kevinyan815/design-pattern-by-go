package main

import "fmt"

// 中介者--机场指挥塔的接口定义
type mediator interface {
	canLanding(airplane airplane) bool
	notifyAboutDeparture()
}

// 组件--飞行器的接口定义
type airplane interface {
	landing()
	takeOff()
	permitLanding()
}

// 组件1--波音飞机
type boeingPlane struct {
	mediator
}

func (b *boeingPlane) landing() {
	if !b.mediator.canLanding(b) {
		fmt.Println("Airplane Boeing: 飞机跑到正在被占用，无法降落！")
		return
	}
	fmt.Println("Airplane Boeing: 已成功降落！")
}

func (b *boeingPlane)takeOff() {
	fmt.Println("Airplane Boeing: 正在起飞离开跑道！")
	b.mediator.notifyAboutDeparture()
}

func (b *boeingPlane)permitLanding() {
	fmt.Println("Airplane Boeing: 收到指挥塔信号，允许降落，正在降落！")
	b.landing()

}

// 组件2--空客飞机
type airBusPlane struct {
	mediator mediator
}

func (airbus *airBusPlane) landing() {
	if !airbus.mediator.canLanding(airbus) {
		fmt.Println("Airplane AirBus: 飞机跑到正在被占用，无法降落！")
		return
	}
	fmt.Println("Airplane AirBus: 已成功降落！")
}

func (airbus *airBusPlane) takeOff() {
	fmt.Println("Airplane AirBus: 正在起飞离开跑道！")
	airbus.mediator.notifyAboutDeparture()
}

func (airbus *airBusPlane)permitLanding() {
	fmt.Println("Airplane AirBus: 收到指挥塔信号，允许降落，正在降落！")
	airbus.landing()
}

// 中介者实现--指挥塔
type manageTower struct {
	isRunwayFree bool
	airportQueue []airplane
}

func (tower *manageTower) canLanding(airplane airplane) bool {
	if tower.isRunwayFree {
		// 跑道空闲，允许降落，同时把状态变为繁忙
		tower.isRunwayFree = false
		return true
	}
	// 跑道繁忙，把飞机加入等待通知的队列
	tower.airportQueue = append(tower.airportQueue, airplane)
	return false
}

func (tower *manageTower) notifyAboutDeparture() {
	if !tower.isRunwayFree {
		tower.isRunwayFree = true
	}
	if len(tower.airportQueue) > 0 {
		firstPlaneInWaitingQueue := tower.airportQueue[0]
		tower.airportQueue = tower.airportQueue[1:]
		firstPlaneInWaitingQueue.permitLanding()
	}
}

func newManageTower() *manageTower {
	return &manageTower{
		isRunwayFree: true,
	}
}



func main() {
	tower := newManageTower()
	boeing := &boeingPlane{
		mediator: tower,
	}
	airbus := &airBusPlane{
		mediator: tower,
	}
	boeing.landing()
	airbus.landing()
	boeing.takeOff()
}
