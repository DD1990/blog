package controllers

import (
	"errors"
	"fmt"
	"github.com/jicg/liteblog/syserrors"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Get() {
	limit := 7;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryNotesByPage(page, limit, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryNotesCount(title)
	if totcnt%limit == 0 {
		totpage = totcnt / limit
	} else {
		totpage = totcnt/limit + 1
	}
	c.Data["totpage"] = totpage
	c.Data["page"] = page
	c.Data["title"] = title
	c.TplName = "index.html"
}

// @router /blog [get]
func (c *IndexController) Blog() {
	limit := 7;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryNotesByPage(page, limit, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryNotesCount(title)
	if totcnt%limit == 0 {
		totpage = totcnt / limit
	} else {
		totpage = totcnt/limit + 1
	}
	c.Data["totpage"] = totpage
	c.Data["page"] = page
	c.Data["title"] = title
	c.TplName = "index_liteblog.html"
}

// @router /notelist [get]
func (c *IndexController) NoteList(){
	c.Data["TabName"] = "文章列表"
	c.Data["NoteList"] = true
	c.Data["List"] = true
	c.TplName = "logined.html"
}

// @router /notelist2 [get]
func (c *IndexController) NoteList2(){
	limit := 7;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryNotesByPage(page, limit, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryNotesCount(title)
	if totcnt%limit == 0 {
		totpage = totcnt / limit
	} else {
		totpage = totcnt/limit + 1
	}
	c.Data["totpage"] = totpage
	c.Data["page"] = page
	c.Data["title"] = title

	list := make([]map[string]interface{}, len(ns))
	for k, v := range ns {
		row := make(map[string]interface{})
		row["id"] = v.UserID
		row["title"] = v.Title
		row["picture"] = "/static/image/timg.jpg"
		list[k] = row
	}

	c.AjaxList("success",0,int64(len(ns)),list)
}
func (self *IndexController) AjaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}
// @router /details/:key [get]
func (c *IndexController) GetDetail() {
	key := c.Ctx.Input.Param(":key")
	note, err := c.Dao.QueryNoteByKey(key)
	if err != nil {
		c.Abort500(syserrors.NewError("文章不存在", err))
	}
	go c.Dao.AllVisitCount(key)
	c.Data["praise"] = false
	//praise, err := c.Dao.QueryPraiseLog(key, int(c.User.ID), "note")
	//if err == nil {
	//	c.Data["praise"] = praise.Flag
	//}
	messages, _ := c.Dao.QueryMessageForNote(note.Key)
	c.Data["messages"] = messages
	c.Data["note"] = note
	c.TplName = "details.html"
}

// @router /comment/:key [get]
func (c *IndexController) GetComment() {

	key := c.Ctx.Input.Param(":key")
	note, err := c.Dao.QueryNoteByKey(key)
	if err != nil {
		c.Abort500(syserrors.NewError("文章不存在", err))
	}

	c.Data["note"] = note
	c.TplName = "comment.html"
}


// @router /myinfo [get]
func (c *IndexController) MyInfo() {
	//c.Data["IsNoteManager1"] = true
	c.TplName = "myinfo.html"
}

// @router /photo [get]
func (c *IndexController) Photo() {
	//c.Data["IsNoteManager1"] = true
	limit := 3;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	//查询生活类相册
	ns0, err := c.Dao.QueryPhoalbsByPageAndType(page, limit, 0, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns0 != nil {
		c.Data["notes"] = ns0
	}
	//查询旅游类相册
	ns1, err := c.Dao.QueryPhoalbsByPageAndType(page, limit, 1, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns1 != nil {
		c.Data["notes1"] = ns1
	}
	//查询美食类相册
	ns2, err := c.Dao.QueryPhoalbsByPageAndType(page, limit,2, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns2 != nil {
		c.Data["notes2"] = ns2
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryPhoalbsCountByType(0)
	totcnt1, _ := c.Dao.QueryPhoalbsCountByType(1)
	totcnt2, _ := c.Dao.QueryPhoalbsCountByType(2)
	if totcnt%limit == 0 {
		totpage = totcnt / limit
	} else {
		totpage = totcnt/limit + 1
	}
	c.Data["totpage"] = totpage
	c.Data["page"] = page
	c.Data["title"] = title


	c.Data["total1"] = totcnt
	c.Data["total2"] = totcnt1
	c.Data["total3"] = totcnt2

	c.TplName = "photo.html"
}

// @router /logined [get]
func (c *IndexController) Logined() {
	c.Data["IsNoteManager1"] = true
	c.TplName = "logined.html"
}
// @router /logined/mu1 [get]
func (c *IndexController) Logined1() {
	c.Data["IsNoteManager1"] = true
	c.TplName = "logined.html"
}
// @router /logined/mu2 [get]
func (c *IndexController) Logined2() {
	c.Data["key"] = c.UUID()
	/*c.Data["IsNoteManager2"] = true*/
	c.TplName = "logined_notenew.html"
}

// @router /user [get]
func (c *IndexController) GetUser() {
	if c.IsLogin {
		c.Data["IsNoteManager1"] = true
		c.TplName = "logined.html"
	}else {
		c.TplName = "user.html"
	}
}

// @router /reg [get]
func (c *IndexController) GetReg() {
	c.TplName = "reg.html"
}

// @router /message [get]
func (c *IndexController) GetMessage() {
	messages, err := c.Dao.QueryMessageForNote("")
	if err != nil {
		c.Abort500(err)
	}
	fmt.Printf("%v \n", messages)
	c.Data["messages"] = messages
	c.TplName = "message.html"
}

// @router /about [get]
func (c *IndexController) GetAbout() {
	c.TplName = "about.html"
}


// @router /love [get]
func (c *IndexController) Love(){
	c.TplName = "love.html"
}