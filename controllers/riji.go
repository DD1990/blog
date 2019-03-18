package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

type RijiController struct {
	BaseController
}

// 如果实现了NestPrepare() 接口，则请求会判断用户是否登录
/*func (ctx *RijiController) NestPrepare() {
	ctx.MustLogin()
	if ctx.User.Role != 0 {
		ctx.Abort500(syserrors.NewError("您没有权限修改文章", nil))
	}
}*/

// @router /new [get]
func (ctx *RijiController) NewPage() {
	ctx.Data["key"] = ctx.UUID()
	ctx.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (ctx *RijiController) EditPage() {
	key := ctx.Ctx.Input.Param(":key")
	note, err := ctx.Dao.QueryRijiByKeyAndUserId(key, int(ctx.User.ID))
	if err != nil {
		ctx.Abort500(syserrors.NewError("文章不存在", err))
	}
	ctx.Data["note"] = note
	ctx.Data["key"] = key
	ctx.TplName = "note_new.html"
}

// @router /del/:key [post]
func (ctx *RijiController) Del() {
	key := ctx.Ctx.Input.Param(":key")
	if err := ctx.Dao.DelRijiByKey(key, int(ctx.User.ID)); err != nil {
		ctx.Abort500(syserrors.NewError("删除失败", err))
	}
	ctx.JSONOk("删除成功！", "/")
}

// @router /save/:key [post]
func (ctx *RijiController) Save() {
	key := ctx.Ctx.Input.Param(":key")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	files := ctx.GetString("files", "")
	summary, _ := getRijiSummary(content)
	note, err := ctx.Dao.QueryRijiByKeyAndUserId(key, int(ctx.User.ID))
	var n models.Riji
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.Riji{
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
	if err := ctx.Dao.SaveRiji(&n); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOk("成功", "/details/"+key)
}

func getRijiSummary(content string) (string, error) {
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
func (c *RijiController) RijiEdit() {
	c.Data["key"] = c.UUID()
	/*c.Data["IsNoteManager2"] = true*/
	c.TplName = "logined_riji.html"
}

// @router /logined [get]
func (c *RijiController) RijiList() {
	c.Data["List"] = true
	c.Data["RijiList"] = true
	c.TplName = "logined.html"
}

// @router /mood [get]
func (c *RijiController) Mood() {
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




	now := time.Now()

	local2, err2 := time.LoadLocation("Local") //服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}

	note, err := c.Dao.QueryRijiByDate(now.In(local2).String()[0:10] + "%")
	if err != nil {
		//c.Abort500(syserrors.NewError("懒懒的大大飞今天没写日记哦", err))
		//TODO(lxf) 在日志区显示没有写日志
		n := models.Riji{
			Key:     "",
			Summary: "懒懒的大大飞今天没写日记哦",
			Title:   "",
			Files:   "",
			Content: "懒懒的大大飞今天没写日记哦",
		}
		c.Data["note"] = n
		c.TplName = "mood.html"
		return
	}

	datacha,_ := dateFormat(now.In(local2).String()[0:10])
	c.Data["datacha"] = datacha
	c.Data["note"] = note
	c.TplName = "mood.html"
}

// @router /moodajax [get]
func (c *RijiController) MoodAjax() {
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
		n := models.Riji{
			Key:     "",
			Summary: "懒懒的大大飞今天没写日记哦",
			Title:   "",
			Files:   "",
			Content: "懒懒的大大飞今天没写日记哦",
		}
		c.Data["note"] = n
		c.JSONOkData(1,n)
		return
	}

	datacha,_ := dateFormat(timeparam)
	c.Data["datacha"] = datacha
	note.Files = datacha  // 暂时使用files字段当中文日期
	c.Data["note"] = note


	//	c.JSONOk("成功", "/mood")
	c.JSONOkData(1,note)

}

// @router /riji/control/list [get]
func (c *RijiController) RijiControlList(){
	limit := 7;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if c.Dao==nil{
		c.Abort500(errors.New("数据库初始化失败！"))
	}
	ns, err := c.Dao.QueryRijisByPage(page, limit, title)
	if err != nil {
		c.Abort500(err)
	}
	if ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := c.Dao.QueryRijisCount(title)
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
func (self *RijiController) AjaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

func dateFormat(timestring string) (string, error) {
	year := timestring[0:4]
	year = strings.Replace(year, "1", "一", -1)
	year = strings.Replace(year, "2", "二", -1)
	year = strings.Replace(year, "3", "三", -1)
	year = strings.Replace(year, "4", "四", -1)
	year = strings.Replace(year, "5", "五", -1)
	year = strings.Replace(year, "6", "六", -1)
	year = strings.Replace(year, "7", "七", -1)
	year = strings.Replace(year, "8", "八", -1)
	year = strings.Replace(year, "9", "九", -1)
	year = strings.Replace(year, "0", "〇", -1)

	month := timestring[5:7]
	monthnum := make(map[string]string)
	monthnum["01"] = "一"
	monthnum["02"] = "二"
	monthnum["03"] = "三"
	monthnum["04"] = "四"
	monthnum["05"] = "五"
	monthnum["06"] = "六"
	monthnum["07"] = "七"
	monthnum["08"] = "八"
	monthnum["09"] = "九"
	monthnum["10"] = "十"
	monthnum["11"] = "十一"
	monthnum["12"] = "十二"
	month = monthnum[month]

	day := timestring[8:10]
	dayarray := [31]string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "二十一", "二十二", "二十三", "二十四", "二十五", "二十六", "二十七", "二十八", "二十九", "三十", "三十一"}
	dayint, _ := strconv.Atoi(day)
	day = dayarray[dayint-1]

	fmt.Println(year + "年" + month + "月" + day + "日")

	return year + "年" + month + "月" + day + "日", nil
}
