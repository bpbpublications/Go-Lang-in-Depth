package models

type Product struct {
	Id       int    `form:"-"`
	Name     string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
	Category string `form:"category,text,category:"`
	Image    string `form:"image,text,image:"`
}

func (a *Product) TableName() string {
	return "products"
}
