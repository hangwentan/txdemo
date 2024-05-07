package modelfunc

// 定义接口
type Product interface {
	GetName() string
}

// 定义产品A
type ProductA struct {
}

func (p *ProductA) GetName() string {
	return "ProductA"
}

// 定义产品B
type ProductB struct {
}

func (p *ProductB) GetName() string {
	return "ProductB"
}

// 定义工厂接口
type Factory interface {
	CreateProduct() Product
}

// 定义产品A工厂
type ProductAFactory struct {
}

func (f *ProductAFactory) CreateProduct() Product {
	return &ProductA{}
}

// 定义产品B工厂
type ProductBFactory struct {
}

func (f *ProductBFactory) CreateProduct() Product {
	return &ProductB{}
}
