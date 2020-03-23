package controllers

import (
	"bytes"
	"github.com/kataras/iris"
	"iris-gateway/common"
	"net/http"
	"time"
)

type BaseController struct{}

//校验Token
func (c *BaseController) VerifyToken(ctx *iris.Context) (err error) {
	//todo 此处可用于token校验
	return err
}

//返回正确时数据结构
func (c *BaseController) RenderSuccess(ctx *iris.Context, data interface{}) {
	rsp := common.Success(data)
	c.RenderJson(ctx, rsp)
}

//返回错误时数据结构
func (c *BaseController) RenderError(ctx *iris.Context, code uint16, message string) {
	rsp := common.Error(code, message)
	c.RenderJson(ctx, rsp)
}

//返回json数据
func (c *BaseController) RenderJson(ctx *iris.Context, response interface{}) {
	_, err := (*ctx).JSON(response)
	if err != nil {
		panic(err)
	}
	(*ctx).Next()
}

//直接返回文本
func (c *BaseController) RenderText(ctx *iris.Context, html string) {
	_, err := (*ctx).Text(html)
	if err != nil {
		panic(err)
	}
	(*ctx).Next()
}

//不保存本地，直接从内存下载文件
func (c *BaseController) DownloadXlsx(ctx *iris.Context, filename string, buffer bytes.Buffer) {
	req := (*ctx).Request()
	wt := (*ctx).ResponseWriter()
	wt.Header().Add("Content-Disposition", "attachment; filename="+filename+".xlsx")
	wt.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(wt, req, filename, time.Now(), r)
}

func (c *BaseController) DownloadDocx(ctx *iris.Context, filename string, buffer bytes.Buffer) {
	req := (*ctx).Request()
	wt := (*ctx).ResponseWriter()
	wt.Header().Add("Content-Disposition", "attachment; filename="+filename+".docx")
	wt.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")

	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(wt, req, filename, time.Now(), r)
}

func (c *BaseController) DownloadPdf(ctx *iris.Context, filename string, buffer bytes.Buffer) {
	req := (*ctx).Request()
	wt := (*ctx).ResponseWriter()
	wt.Header().Add("Content-Disposition", "attachment; filename="+filename+".pdf")
	wt.Header().Add("Content-Type", "application/pdf")

	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(wt, req, filename, time.Now(), r)
}
