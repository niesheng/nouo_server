package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("server", "Nouo Web Server")
	ctx.Response.Header.Set("version", "0.1.1")

	if string(ctx.RequestURI()) == "/index.html" {
		b, err := ioutil.ReadFile("index.html")
		if err != nil {
			Exit(err)
		}
		ctx.SetContentType("text/html")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(b)
		return
	}

	getMap := make(map[string]interface{})
	postMap := make(map[string]interface{})
	cookieMap := make(map[string]interface{})
	headerMap := make(map[string]interface{})

	ctx.QueryArgs().VisitAll(fmtParams(getMap))
	ctx.PostArgs().VisitAll(fmtParams(postMap))
	ctx.Request.Header.VisitAllCookie(fmtParams(cookieMap))
	ctx.Request.Header.VisitAll(fmtParams(headerMap))

	// ctx.Form

	request := webRequest{
		Url:    string(ctx.RequestURI()),
		Method: string(ctx.Method()),
		Get:    getMap,
		Post:   postMap,
		Cookie: cookieMap,
		Header: headerMap,
		Ip:     ctx.RemoteIP().String(),
	}
	fmt.Println(string(ctx.Request.Header.ContentType()))
	if string(ctx.Method()) == "POST" && strings.HasPrefix(string(ctx.Request.Header.ContentType()), "multipart/form-data") {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.Error(err.Error(), 500)
			return
		}
		request.Files = file_handle(form.File)
	}
	result, err := json.Marshal(request)
	if err != nil {
		ctx.Error("Request Error", 500)
		return
	}
	fmt.Fprintf(ctx, "%s", result)

}

func file_handle(mf map[string][]*multipart.FileHeader) map[string][]webFile {
	files := make(map[string][]webFile)
	for key, value := range mf {
		vfiles := make([]webFile, len(value))
		for i := 0; i < len(value); i++ {
			s, err := value[i].Open()
			if err != nil {
				//error log
				break
			}
			truePath := Config_.Upload + value[i].Filename
			b, _ := ioutil.ReadAll(s)
			ioutil.WriteFile(truePath, b, 644)
			os.Chown(truePath, Uid_, Gid_)
			vfiles[i] = webFile{
				Path: value[i].Filename,
				Size: value[i].Size,
				Name: value[i].Filename,
			}
		}
		files[key] = vfiles
	}
	return files
}
