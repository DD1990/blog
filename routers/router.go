package routers

import (
	"github.com/jicg/liteblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
		&controllers.ShiGeController{},
		&controllers.PhoalbController{},
		&controllers.PhotosController{},
		&controllers.RijiController{},
	)
	beego.AddNamespace(
		beego.NewNamespace(
			"note",
			beego.NSInclude(&controllers.NoteController{}),
		),
		beego.NewNamespace(
			"upload",
			beego.NSInclude(&controllers.UploadController{}),
		),
		beego.NewNamespace(
			"praise",
			beego.NSInclude(&controllers.PraiseController{}),
		),
		beego.NewNamespace(
			"message",
			beego.NSInclude(&controllers.MessageController{}),
		),
		beego.NewNamespace(
			"simple",
			beego.NSInclude(&controllers.SimpleController{}),
		),
	)
}
