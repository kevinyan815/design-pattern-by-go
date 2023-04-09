package main

import "fmt"

// 订单实体类，实现IOrderService 接口
type Order struct {
	ID int
	Customer string
	City string
	Product string
	Quantity int
}

func NewOrder(id int, customer string, city string, product string, quantity int) *Order {
	return &Order{
		id, customer,city,product,quantity,
	}
}

// 订单服务接口
type IOrderService interface {
	Save(order *Order) error
	// 有的教程里把接收 visitor 实现的方法名定义成 Accept
	Accept(visitor IOrderVisitor)
}

type IOrderVisitor interface {
	// 这里参数不能定义成 IOrderService
	Visit(order *Order)
	Report()
}

// 销售订单服务，实现IOrderService接口
type OrderService struct {
	orders map[int]*Order
}

func (mo *OrderService) Save(o *Order) error {
	mo.orders[o.ID] = o
	return nil
}

func (mo *OrderService) Accept(visitor IOrderVisitor) {
	for _, v := range mo.orders {
		visitor.Visit(v)
	}
}

func NewOrderService() IOrderService {
	return &OrderService{
		orders: make(map[int]*Order, 0),
	}
}

// 区域销售报表, 按城市汇总销售情况, 实现IOrderVisitor接口
type CityVisitor struct {
	cities map[string]int
}

func (cv *CityVisitor) Visit(o *Order) {
	n, ok := cv.cities[o.City]
	if ok {
		cv.cities[o.City] = n + o.Quantity
	} else {
		cv.cities[o.City] = o.Quantity
	}
}

func (cv *CityVisitor) Report() {
	for k,v := range cv.cities {
		fmt.Printf("city=%s, sum=%v\n", k, v)
	}
}

func NewCityVisitor() IOrderVisitor {
	return &CityVisitor{
		cities: make(map[string]int, 0),
	}
}

// 品类销售报表, 按产品汇总销售情况, 实现ISaleOrderVisitor接口
type ProductVisitor struct {
	products map[string]int
}

func (pv *ProductVisitor) Visit(it *Order) {
	n,ok := pv.products[it.Product]
	if ok {
		pv.products[it.Product] = n + it.Quantity
	} else {
		pv.products[it.Product] = it.Quantity
	}
}

func (pv *ProductVisitor) Report() {
	for k,v := range pv.products {
		fmt.Printf("product=%s, sum=%v\n", k, v)
	}
}

func NewProductVisitor() IOrderVisitor {
	return &ProductVisitor{
		products: make(map[string]int,0),
	}
}


func main() {
	orderService := NewOrderService()
	orderService.Save(NewOrder(1, "张三", "广州", "电视", 10))
	orderService.Save(NewOrder(2, "李四", "深圳", "冰箱", 20))
	orderService.Save(NewOrder(3, "王五", "东莞", "空调", 30))
	orderService.Save(NewOrder(4, "张三三", "广州", "空调", 10))
	orderService.Save(NewOrder(5, "李四四", "深圳", "电视", 20))
	orderService.Save(NewOrder(6, "王五五", "东莞", "冰箱", 30))

	cv := NewCityVisitor()
	orderService.Accept(cv)
	cv.Report()

	pv := NewProductVisitor()
	orderService.Accept(pv)
	pv.Report()
}
