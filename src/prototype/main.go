package main

import (
  "dom"
  "fmt"
)

func main() {
  // 职级节点--总监
 directorNode := dom.NewElement("Director of Engineering")
  // 职级节点--研发经理
 engManagerNode := dom.NewElement("Engineering Manager")
 engManagerNode.AddChild(dom.NewElement("Lead Software Engineer"))
  // 研发经理是总监的下级
 directorNode.AddChild(engManagerNode)
 directorNode.AddChild(engManagerNode)
  // 办公室经理也是总监的下级
 officeManagerNode := dom.NewElement("Office Manager")
 directorNode.AddChild(officeManagerNode)
 fmt.Println("")
 fmt.Println("# Company Hierarchy")
 fmt.Print(directorNode)
 fmt.Println("")
  // 从研发经理节点克隆出一颗新的树
 fmt.Println("# Team Hiearachy")
 fmt.Print(engManagerNode.Clone())
}
