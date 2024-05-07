package main

import (
	"example.com/m/modelfunc"
	"fmt"
)

func main() {
	// 创建单例
	svc := modelfunc.Service()
	svc.Restart()
	svc.Startup()
	svc.Stop()

	// 创建产品A工厂
	factoryA := &modelfunc.ProductAFactory{}

	// 使用产品A工厂创建产品A
	productA := factoryA.CreateProduct()

	// 打印产品A的名称
	fmt.Println(productA.GetName())

	// 创建产品B工厂
	factoryB := &modelfunc.ProductBFactory{}

	// 使用产品B工厂创建产品B
	productB := factoryB.CreateProduct()

	// 打印产品B的名称
	fmt.Println(productB.GetName())
}
