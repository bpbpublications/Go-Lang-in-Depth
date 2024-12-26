package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"go_web_app_beego/models"
	"strconv"
)

type ManageController struct {
	beego.Controller
}

func (manage *ManageController) Home() {
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplName = "manage/home.tpl"
}

func (manage *ManageController) Delete() {
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplName = "manage/home.tpl"

	productId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))

	ormGorm := orm.NewOrm()
	ormGorm.Using("default")

	product := models.Product{}

	if exist := ormGorm.QueryTable(product.TableName()).Filter("Id", productId).Exist(); exist {
		if num, err := ormGorm.Delete(&models.Product{Id: productId}); err == nil {
			beego.Info("Record Deleted. ", num)
		} else {
			beego.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		beego.Info("Record Doesn't exist.")
	}
}

func (manage *ManageController) Update() {
	ormGorm := orm.NewOrm()
	ormGorm.Using("default")
	flash := beego.NewFlash()

	if productId, err := strconv.Atoi(manage.Ctx.Input.Param(":id")); err == nil {
		product := models.Product{Id: productId}
		if ormGorm.Read(&product) == nil {
			product.Category = "Sitepoint"
			product.Image = "http://www.google.com"
			if num, err := ormGorm.Update(&product); err == nil {
				flash.Notice("Record Was Updated.")
				flash.Store(&manage.Controller)
				beego.Info("Record Was Updated. ", num)
			}
		} else {
			flash.Notice("Record Was NOT Updated.")
			flash.Store(&manage.Controller)
			beego.Error("Couldn't find product matching id: ", productId)
		}
	} else {
		flash.Notice("Record Was NOT Updated.")
		flash.Store(&manage.Controller)
		beego.Error("Couldn't convert id from a string to a number. ", err)
	}

	manage.Redirect("/manage/view", 302)
}

func (manage *ManageController) View() {
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplName = "manage/view.tpl"

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		manage.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		manage.Data["notices"] = ok
	}

	ormGorm := orm.NewOrm()
	ormGorm.Using("default")

	var products []*models.Product
	num, err := ormGorm.QueryTable("products").All(&products)

	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = products
	}
}

func (manage *ManageController) Add() {
	manage.Data["Form"] = &models.Product{}
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplName = "manage/add.tpl"

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		manage.Data["flash"] = ok
	}

	ormGorm := orm.NewOrm()
	ormGorm.Using("default")

	product := models.Product{}

	if err := manage.ParseForm(&product); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		manage.Data["Products"] = product
		valid := validation.Validation{}
		isValid, _ := valid.Valid(product)

		if manage.Ctx.Input.Method() == "POST" {
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {
				searchProduct := models.Product{Name: product.Name}
				beego.Debug("Product name supplied:", product.Name)
				err = ormGorm.Read(&searchProduct)
				beego.Debug("Err:", err)
				flash := beego.NewFlash()

				if err == orm.ErrNoRows || err == orm.ErrMissPK {
					beego.Debug("No product found matching details supplied. Attempting to insert product: ", product)
					id, error := ormGorm.Insert(&product)
					if error == nil {
						msg := fmt.Sprintf("Product inserted with id:", id)
						beego.Debug(msg)
						flash.Notice(msg)
						flash.Store(&manage.Controller)
					} else {
						msg := fmt.Sprintf("Couldn't insert new product. Reason: ", err)
						beego.Debug(msg)
						flash.Error(msg)
						flash.Store(&manage.Controller)
					}
				} else {
					beego.Debug("Product found matching details supplied. Cannot insert")
				}
			}
		}
	}
}
