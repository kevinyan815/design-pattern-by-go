package main

import (
    "fmt"
    "strings"
)

// 表示组织机构的接口
type Organization interface {
    display()
    duty()
}

// 组合对象--上级部门
type CompositeOrganization struct {
    orgName string
    depth   int
    list    []Organization
}

func NewCompositeOrganization(name string, depth int) *CompositeOrganization {
    return &CompositeOrganization{name, depth, []Organization{}}
}

func (c *CompositeOrganization) add(org Organization) {
    if c == nil {
        return
    }
    c.list = append(c.list, org)
}

func (c *CompositeOrganization) remove(org Organization) {
    if c == nil {
        return
    }
    for i, val := range c.list {
        if val == org {
            c.list = append(c.list[:i], c.list[i+1:]...)
            return
        }
    }
    return
}

func (c *CompositeOrganization) display() {
    if c == nil {
        return
    }
    fmt.Println(strings.Repeat("-", c.depth * 2), " ", c.orgName)
    for _, val := range c.list {
        val.display()
    }
}

func (c *CompositeOrganization) duty() {
    if c == nil {
        return
    }

    for _, val := range c.list {
        val.duty()
    }
}

// Leaf对象--人力资源部门
type HRDOrg struct {
    orgName string
    depth   int
}

func (o *HRDOrg) display() {
    if o == nil {
        return
    }
    fmt.Println(strings.Repeat("-", o.depth * 2), " ", o.orgName)
}

func (o *HRDOrg) duty() {
    if o == nil {
        return
    }
    fmt.Println(o.orgName, "员工招聘培训管理")
}

// Leaf对象--财务部门
type FinanceOrg struct {
    orgName string
    depth   int
}

func (f *FinanceOrg) display() {
    if f == nil {
        return
    }
    fmt.Println(strings.Repeat("-", f.depth * 2), " ", f.orgName)
}

func (f *FinanceOrg) duty() {
    if f == nil {
        return
    }
    fmt.Println(f.orgName, "员工招聘培训管理")
}

func main() {
    root := NewCompositeOrganization("北京总公司", 1)
    root.add(&HRDOrg{orgName: "总公司人力资源部", depth: 2})
    root.add(&FinanceOrg{orgName: "总公司财务部", depth: 2})

    compSh := NewCompositeOrganization("上海分公司", 2)
    compSh.add(&HRDOrg{orgName: "上海分公司人力资源部", depth: 3})
    compSh.add(&FinanceOrg{orgName: "上海分公司财务部", depth: 3})
    root.add(compSh)

    compGd := NewCompositeOrganization("广东分公司", 2)
    compGd.add(&HRDOrg{orgName: "广东分公司人力资源部", depth: 3})
    compGd.add(&FinanceOrg{orgName: "南京办事处财务部", depth: 3})
    root.add(compGd)

    fmt.Println("公司组织架构：")
    root.display()

    fmt.Println("各组织的职责：")
    root.duty()
}
