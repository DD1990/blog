package controllers

import (
	"errors"
	"github.com/jicg/liteblog/models"
	"strconv"
	"strings"
	"time"
	"github.com/jicg/liteblog/syserrors"
	"github.com/jinzhu/gorm"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

type PhoalbController struct {
	BaseController
}

func (ctx *PhoalbController) NestPrepare() {
	ctx.MustLogin()
	if ctx.User.Role != 0 {
		ctx.Abort500(syserrors.NewError("您没有权限修改文章", nil))
	}
}

// @router /new [get]
func (ctx *PhoalbController) NewPage() {
	ctx.Data["key"] = ctx.UUID()
	ctx.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (ctx *PhoalbController) EditPage() {
	key := ctx.Ctx.Input.Param(":key")
	note, err := ctx.Dao.QueryPhoalbByKeyAndUserId(key, int(ctx.User.ID))
	if err != nil {
		ctx.Abort500(syserrors.NewError("文章不存在", err))
	}
	ctx.Data["note"] = note
	ctx.Data["key"] = key
	ctx.TplName = "note_new.html"
}

// @router /del/:key [post]
func (ctx *PhoalbController) Del() {
	key := ctx.Ctx.Input.Param(":key")
	if err := ctx.Dao.DelPhoalbByKey(key, int(ctx.User.ID)); err != nil {
		ctx.Abort500(syserrors.NewError("删除失败", err))
	}
	ctx.JSONOk("删除成功！", "/")
}

// @router /save/:key [post]
func (ctx *PhoalbController) Save() {
	key := ctx.Ctx.Input.Param(":key")
	visitstr := ctx.GetString("visit")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	files := ctx.GetString("files", "")
	summary, _ := getPhoalbSummary(content)
	note, err := ctx.Dao.QueryPhoalbByKeyAndUserId(key, int(ctx.User.ID))

	visit, _ := strconv.Atoi(visitstr)
	filefengmian := strings.Split(files,";")[0]
	var n models.Phoalb
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.Phoalb{
			Key:     key,
			Summary: summary,
			Title:   title,
			Files:   filefengmian,
			Content: content,
			Visit:   visit,
			UserID:  int(ctx.User.ID),
		}
	} else {
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.Files = filefengmian
		n.Visit = visit
		n.UpdatedAt = time.Now()
	}
	if err := ctx.Dao.SavePhoalb(&n); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOk("成功", "/photolist")
}

func getPhoalbSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))
	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	str := doc.Find("body").Text()
	strRune := []rune(str)
	if len(strRune) > 400 {
		strRune = strRune[:400]
	}
	return string(strRune) + "...", nil
}


// @router /logined/photos [get]
func (c *PhoalbController) LoginedPhoalb() {
	c.Data["key"] = c.UUID()
	/*c.Data["IsNoteManager2"] = true*/
	c.TplName = "logined_phoalb.html"
}

// @router /photolist [get]
func (c *PhoalbController) PhotoList(){
	c.Data["TabName"] = "相册列表"
	c.Data["PhotoList"] = true
	c.Data["List"] = true
	c.TplName = "logined.html"
}

// @router /notelist2 [get]
func (c *PhoalbController) PhoalbList(){
	limit := 7;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryPhoalbsByPage(page, limit, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryPhoalbsCount(title)
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
		row["picture"] = v.Files
		list[k] = row
	}

	c.AjaxList("success",0,int64(len(ns)),list)
}




// @router /nextthree/1/:key [get]
func (c *PhoalbController) NextThree1(){

	key := c.Ctx.Input.Param(":key")

	limit := 3;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}

	keynum,_ := strconv.ParseInt(key, 10,32)
	page = int(keynum) + 1


	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryPhoalbsByPageAndType(page, limit, 0, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryPhoalbsCount(title)
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
		row["picture"] = v.Files
		row["key"] = v.Key
		list[k] = row
	}

	c.AjaxList("success",0,int64(len(ns)),list)
}

// @router /nextthree/2/:key [get]
func (c *PhoalbController) NextThree2(){

	key := c.Ctx.Input.Param(":key")

	limit := 3;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}

	keynum,_ := strconv.ParseInt(key, 10,32)
	page = int(keynum) + 1


	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryPhoalbsByPageAndType(page, limit, 1, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryPhoalbsCount(title)
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
		row["picture"] = v.Files
		row["key"] = v.Key
		list[k] = row
	}

	c.AjaxList("success",0,int64(len(ns)),list)
}
// @router /nextthree/3/:key [get]
func (c *PhoalbController) NextThree3(){

	key := c.Ctx.Input.Param(":key")

	limit := 3;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}

	keynum,_ := strconv.ParseInt(key, 10,32)
	page = int(keynum) + 1


	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryPhoalbsByPageAndType(page, limit, 2, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryPhoalbsCount(title)
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
		row["picture"] = v.Files
		row["key"] = v.Key
		list[k] = row
	}

	c.AjaxList("success",0,int64(len(ns)),list)
}

func (self *PhoalbController) AjaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}
