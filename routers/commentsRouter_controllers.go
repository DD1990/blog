package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetAbout",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetComment",
			Router: `/comment/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetDetail",
			Router: `/details/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetMessage",
			Router: `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetReg",
			Router: `/reg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Logined",
			Router: `/logined`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "NoteList",
			Router: `/notelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "NoteList2",
			Router: `/notelist2`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Logined1",
			Router: `/logined/mu1`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Logined2",
			Router: `/logined/mu2`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "MyInfo",
			Router: `/myinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Photo",
			Router: `/photo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Blog",
			Router: `/blog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Count",
			Router: `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:MessageController"],
		beego.ControllerComments{
			Method: "NewMessage",
			Router: `/new/?:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "EditPage",
			Router: `/edit/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "NewPage",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PraiseController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PraiseController"],
		beego.ControllerComments{
			Method: "Parse",
			Router: `/:type/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UploadController"],
		beego.ControllerComments{
			Method: "UploadFile",
			Router: `/uploadfile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UploadController"],
		beego.ControllerComments{
			Method: "UploadImg",
			Router: `/uploadimg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UploadController"],
		beego.ControllerComments{
			Method: "UploadImgs",
			Router: `/upload/imgs`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Reg",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"],
		beego.ControllerComments{
			Method: "ShiGe",
			Router: `/shige`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"],
		beego.ControllerComments{
			Method: "NewShiGe",
			Router: `/shige/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"],
		beego.ControllerComments{
			Method: "SaveShiGe",
			Router: `/shige/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:ShiGeController"],
		beego.ControllerComments{
			Method: "GetShiGeDetail",
			Router: `/shige/details/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:SimpleController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:SimpleController"],
		beego.ControllerComments{
			Method: "Photo",
			Router: `/photo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:SimpleController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:SimpleController"],
		beego.ControllerComments{
			Method: "PhotoDetails",
			Router: `/photodetails/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

/************************************************************PhoalbController****************************************************************/

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "EditPage",
			Router: `/edit/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "NewPage",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/phoalb/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "LoginedPhoalb",
			Router: `/logined/phoalb`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "PhotoList",
			Router: `/photolist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "PhoalbList",
			Router: `/phoalb/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "NextThree1",
			Router: `/nextthree/1/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "NextThree2",
			Router: `/nextthree/2/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhoalbController"],
		beego.ControllerComments{
			Method: "NextThree3",
			Router: `/nextthree/3/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	/************************************************************PhotosController****************************************************************/


	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"],
		beego.ControllerComments{
			Method: "EditPage",
			Router: `/edit/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"],
		beego.ControllerComments{
			Method: "NewPage",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/photos/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"],
		beego.ControllerComments{
			Method: "LoginedPhotos",
			Router: `/logined/photos`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})


	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:PhotosController"],
		beego.ControllerComments{
			Method: "ColumnQuery",
			Router: `/column/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	/************************************************************RijiController****************************************************************/


	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "EditPage",
			Router: `/edit/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "NewPage",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/riji/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "RijiEdit",
			Router: `/riji/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "RijiList",
			Router: `/riji/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "Mood",
			Router: `/mood`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "MoodAjax",
			Router: `/moodajax`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:RijiController"],
		beego.ControllerComments{
			Method: "RijiControlList",
			Router: `/riji/control/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

/****************************************************************** Love ***************************************************************************/


	beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Love",
			Router: `/love`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})
}
