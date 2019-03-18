package controllers

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
	"github.com/jinzhu/gorm"
	"time"
)

type ShiGeController struct {
	BaseController
}

// @router /shige [get]
func (c *ShiGeController) ShiGe() {
	limit := 10;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryShiGesByPage(page, limit, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryShiGesCount(title)
	if totcnt%limit == 0 {
		totpage = totcnt / limit
	} else {
		totpage = totcnt/limit + 1
	}
	c.Data["totpage"] = totpage
	c.Data["page"] = page
	c.Data["title"] = title
	c.TplName = "shige.html"
}
// @router /shige/new [get]
func (ctx *ShiGeController) NewShiGe() {
	ctx.Data["key"] = ctx.UUID()
	ctx.TplName = "shige_new.html"
}

// @router /save/:key [post]
func (ctx *ShiGeController) SaveShiGe() {
	key := ctx.Ctx.Input.Param(":key")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	files := ctx.GetString("files", "")
	summary, _ := getShiGeSummary(content)
	shige, err := ctx.Dao.QueryShiGeByKeyAndUserId(key, int(ctx.User.ID))
	var n models.ShiGe
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.ShiGe{
			Key:     key,
			Summary: summary,
			Title:   title,
			Files:   files,
			Content: content,
			UserID:  int(ctx.User.ID),
		}
	} else {
		n = shige
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.Files = files
		n.UpdatedAt = time.Now()
	}
	if err := ctx.Dao.SaveShiGe(&n); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOk("成功", "/details/"+key)
}


// @router /details/:key [get]
func (c *ShiGeController) GetShiGeDetail() {
	key := c.Ctx.Input.Param(":key")
	shige, err := c.Dao.QueryShiGeByKey(key)
	if err != nil {
		c.Abort500(syserrors.NewError("文章不存在", err))
	}
	go c.Dao.AllVisitOnShiGeCount(key)
	c.Data["praise"] = false
	//praise, err := c.Dao.QueryPraiseLog(key, int(c.User.ID), "note")
	//if err == nil {
	//	c.Data["praise"] = praise.Flag
	//}
	messages, _ := c.Dao.QueryMessageForNote(shige.Key)
	c.Data["messages"] = messages
	c.Data["note"] = shige
	c.TplName = "details.html"
}

func getShiGeSummary(content string) (string, error) {
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