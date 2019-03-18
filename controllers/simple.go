package controllers

import (
	"errors"
	"github.com/jicg/liteblog/models"
	"time"
	"github.com/jicg/liteblog/syserrors"
	"github.com/jinzhu/gorm"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

type SimpleController struct {
	BaseController
}

func (ctx *SimpleController) NestPrepare() {
	ctx.MustLogin()
	if ctx.User.Role != 0 {
		ctx.Abort500(syserrors.NewError("您没有权限修改文章", nil))
	}
}

// @router /new [get]
func (ctx *SimpleController) Photo() {
	//ctx.Data["key"] = ctx.UUID() ./../layuiSimpleBlog/html/album.html
	ctx.TplName = "html/album.html"
}

// @router /photodetails [get]
func (ctx *SimpleController) PhotoDetails() {
	//ctx.Data["key"] = ctx.UUID() ./../layuiSimpleBlog/html/album.html

	key := ctx.Ctx.Input.Param(":key")
	ns, err := ctx.Dao.QueryPhotosAllByKey(key)

	title := ctx.GetString("title", "")
	imgsrc := ctx.GetString("files", "")
	if ctx.Dao==nil{
		ctx.Abort500(errors.New("数据库初始化失败！"))
	}
	if err != nil {
		ctx.Abort500(err)
	}
	if ns != nil {
		ctx.Data["notes"] = ns
	}
	ctx.Data["title"] = title
	ctx.Data["imgsrc"] = imgsrc

	/*ctx.TplName = "html/photo_details.html"*/
	ctx.TplName = "photo_view.html"
}

// @router /edit/:key [get]
func (ctx *SimpleController) EditPage() {
	key := ctx.Ctx.Input.Param(":key")
	note, err := ctx.Dao.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
	if err != nil {
		ctx.Abort500(syserrors.NewError("文章不存在", err))
	}
	ctx.Data["note"] = note
	ctx.Data["key"] = key
	ctx.TplName = "note_new.html"
}

// @router /del/:key [post]
func (ctx *SimpleController) Del() {
	key := ctx.Ctx.Input.Param(":key")
	if err := ctx.Dao.DelNoteByKey(key, int(ctx.User.ID)); err != nil {
		ctx.Abort500(syserrors.NewError("删除失败", err))
	}
	ctx.JSONOk("删除成功！", "/")
}

// @router /save/:key [post]
func (ctx *SimpleController) Save() {
	key := ctx.Ctx.Input.Param(":key")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	files := ctx.GetString("files", "")
	summary, _ := getSummary(content)
	note, err := ctx.Dao.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
	var n models.Note
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.Note{
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
	if err := ctx.Dao.SaveNote(&n); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOk("成功", "/details/"+key)
}

func getSimpleSummary(content string) (string, error) {
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
