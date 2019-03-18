package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
	"strings"
	"time"
)

type PhotosController struct {
	BaseController
}

func (ctx *PhotosController) NestPrepare() {
	ctx.MustLogin()
	if ctx.User.Role != 0 {
		ctx.Abort500(syserrors.NewError("您没有权限修改文章", nil))
	}
}

// @router /new [get]
func (ctx *PhotosController) NewPage() {
	ctx.Data["key"] = ctx.UUID()
	ctx.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (ctx *PhotosController) EditPage() {
	key := ctx.Ctx.Input.Param(":key")
	note, err := ctx.Dao.QueryPhotosByKeyAndUserId(key, int(ctx.User.ID))
	if err != nil {
		ctx.Abort500(syserrors.NewError("文章不存在", err))
	}
	ctx.Data["note"] = note
	ctx.Data["key"] = key
	ctx.TplName = "note_new.html"
}

// @router /del/:key [post]
func (ctx *PhotosController) Del() {
	key := ctx.Ctx.Input.Param(":key")
	if err := ctx.Dao.DelPhotosByKey(key, int(ctx.User.ID)); err != nil {
		ctx.Abort500(syserrors.NewError("删除失败", err))
	}
	ctx.JSONOk("删除成功！", "/")
}

// @router /save/:key [post]
func (ctx *PhotosController) Save() {
	/*key := ctx.Ctx.Input.Param(":key")*/
	key := ctx.GetMustString("key", "相册不能为空！")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	files := ctx.GetString("files", "")
	summary, _ := getPhotosSummary(content)
	//note, err := ctx.Dao.QueryPhotosByKeyAndUserId(key, int(ctx.User.ID))

	var n models.Photos
	fileArray := strings.Split(files,";")
	for _, file := range fileArray {
		n = models.Photos{
			Key:     key,
			Summary: summary,
			Title:   title,
			Files:   file,
			Content: content,
			UserID:  int(ctx.User.ID),
		}
		if err := ctx.Dao.SavePhotos(&n); err != nil {
			continue
			logs.Info("保存失败！", err)
			//ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
	}

	ctx.JSONOk("成功", "/details/"+key)
/*	var n models.Photos
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.Photos{
			Key:     key,
			Summary: summary,
			Title:   title,
			Files:   files,
			Content: content,
			UserID:  int(ctx.User.ID),
		}
	} else {
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.Files = files
		n.UpdatedAt = time.Now()
	}
	if err := ctx.Dao.SavePhotos(&n); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOk("成功", "/details/"+key)*/
}

func getPhotosSummary(content string) (string, error) {
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
func (c *PhotosController) LoginedPhotos() {
	/*c.Data["key"] = c.UUID()*/
	/*c.Data["IsNoteManager2"] = true*/
	ns, err := c.Dao.QueryPhoalbs()
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	c.TplName = "logined_photos.html"
}

// @router /column/query [get]
func (c *PhotosController) ColumnQuery() {
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



	// /mood?2018-12-14&_=1544843218564
	now := time.Now()
	local2, err2 := time.LoadLocation("Local") //服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	timeparam := now.In(local2).String()[0:10]

	uri := c.Ctx.Input.URI()
	if len(uri) > 16 {
		timeparam = uri[10:20]
	}

	note, err := c.Dao.QueryRijiByDate( timeparam + "%")
	if err != nil {
		//c.Abort500(syserrors.NewError("懒懒的大大飞今天没写日记哦", err))
		//TODO(lxf) 在日志区显示没有写日志
		n := []models.Riji{
			{
				Key:     "",
				Summary: "懒懒的大大飞今天没写日记哦",
				Title:   "",
				Files:   "/assert/files/2018November/1543303486624994700.png",
				Content: "懒懒的大大飞今天没写日记哦",
			},{
				Key:     "",
				Summary: "懒懒的大大飞今天没写日记哦",
				Title:   "",
				Files:   "/assert/files/2018November/1543303486624994700.png",
				Content: "懒懒的大大飞今天没写日记哦",
			},{
				Key:     "",
				Summary: "懒懒的大大飞今天没写日记哦",
				Title:   "",
				Files:   "/assert/files/2018November/1543303486624994700.png",
				Content: "懒懒的大大飞今天没写日记哦",
			},{
				Key:     "",
				Summary: "懒懒的大大飞今天没写日记哦",
				Title:   "",
				Files:   "/assert/files/2018November/1543303486624994700.png",
				Content: "懒懒的大大飞今天没写日记哦",
			},{
				Key:     "",
				Summary: "懒懒的大大飞今天没写日记哦",
				Title:   "",
				Files:   "/assert/files/2018November/1543303486624994700.png",
				Content: "懒懒的大大飞今天没写日记哦",
			},{
				Key:     "",
				Summary: "懒懒的大大飞今天没写日记哦",
				Title:   "",
				Files:   "/assert/files/2018November/1543303486624994700.png",
				Content: "懒懒的大大飞今天没写日记哦",
			},{
				Key:     "",
				Summary: "懒懒的大大飞今天没写日记哦",
				Title:   "",
				Files:   "/assert/files/2018November/1543303486624994700.png",
				Content: "懒懒的大大飞今天没写日记哦",
			},
		}
		c.Data["note"] = n
		//c.JSONOkData(1,n)

		list := make([]map[string]interface{}, len(ns))
		for k, v := range n {
			row := make(map[string]interface{})
			row["id"] = v.UserID
			row["title"] = v.Title
			row["picture"] = v.Files
			list[k] = row
		}

		c.AjaxList("success",0,int64(len(n)),list)

		return
	}

	note2 := []models.Riji{
		{
			Key:     "",
			Summary: "懒懒的大大飞今天没写日记哦",
			Title:   "",
			Files:   "/assert/files/2018November/1543303486624994700.png",
			Content: "懒懒的大大飞今天没写日记哦",
		},{
			Key:     "",
			Summary: "懒懒的大大飞今天没写日记哦",
			Title:   "",
			Files:   "/assert/files/2018November/1543303486624994700.png",
			Content: "懒懒的大大飞今天没写日记哦",
		},
	}

	datacha,_ := dateFormat(timeparam)
	c.Data["datacha"] = datacha
	c.Data["note2"] = note2
	c.Data["note"] = note
	c.Data["success"] = true


	list := make([]map[string]interface{}, len(ns))
	for k, v := range note2 {
		row := make(map[string]interface{})
		row["id"] = v.UserID
		row["title"] = v.Title
		row["picture"] = v.Files
		list[k] = row
	}

	c.AjaxList("success",0,int64(len(note2)),list)


	//	c.JSONOk("成功", "/mood")
	//c.JSONOkData(1,note2)

}
func (self *PhotosController) AjaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}