package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"reflect"
	"strings"

	"github.com/valyala/fasthttp"
)

func router_handle(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("server", "Nouo Web Router")
	ctx.Response.Header.Set("version", "0.1.1")

	getMap := make(map[string]interface{})
	cookieMap := make(map[string]interface{})
	headerMap := make(map[string]interface{})

	ctx.QueryArgs().VisitAll(fmtParams(getMap))
	ctx.Request.Header.VisitAllCookie(fmtParams(cookieMap))
	ctx.Request.Header.VisitAll(fmtParams(headerMap))

	request := webRequest{
		Path:   string(ctx.Path()[1:]),
		Host:   string(ctx.Host()),
		Tls:    ctx.IsTLS(),
		Method: string(ctx.Method()),
		Get:    getMap,
		Cookie: cookieMap,
		Header: headerMap,
		Ip:     ctx.RemoteIP().String(),
	}
	if string(ctx.Method()) == "POST" {
		postMap := make(map[string]interface{})
		ctx.PostArgs().VisitAll(fmtParams(postMap))
		request.Post = postMap
		if strings.HasPrefix(string(ctx.Request.Header.ContentType()), "multipart/form-data") {
			form, err := ctx.MultipartForm()
			if err != nil {
				ctx.Error(err.Error(), 500)
				return
			}
			request.Files = file_handle(form.File)
		}
	}

	s, code := sql_router(request)
	ctx.SetStatusCode(code)

	for k, v := range s.Header {
		if reflect.TypeOf(v).String() != "string" {
			val, err := json.Marshal(v)
			if err != nil {
				ctx.Error(err.Error(), 500)
				return
			}
			ctx.Response.Header.Set(k, string(val))
		} else {
			ctx.Response.Header.Set(k, v.(string))
		}
	}

	if s.File.Name == "" {
		ctx.SetBodyString(s.Body)
	} else {
		f, err := os.Open(s.File.Path)
		if err != nil {
			Exit(err)
		}
		ctx.SetBodyStream(f, int(s.File.Size))
	}
	return
}

func file_handle(mf map[string][]*multipart.FileHeader) map[string]interface{} {
	files := make(map[string]interface{})
	for key, value := range mf {
		vfiles := make([]webFile, len(value))
		for i := 0; i < len(value); i++ {
			s, err := value[i].Open()
			if err != nil {
				//error log
				break
			}
			truePath := Config_.Server.Upload.Path + value[i].Filename
			b, _ := ioutil.ReadAll(s)
			acc := false
			fmt.Println(value[i].Filename)

			for _, v := range Config_.Server.Upload.Allow {
				fmt.Println(v + " - " + GetFileType(b[:10]))

				if GetFileType(b[:10]) == v {
					acc = true
				}
			}
			if acc {
				ioutil.WriteFile(truePath, b, 644)
				os.Chown(truePath, Uid_, Gid_)
				vfiles[i] = webFile{
					Path: truePath,
					Size: value[i].Size,
					Name: value[i].Filename,
				}
			}
		}
		if len(vfiles) == 1 {
			files[key] = vfiles[0]
		} else {
			files[key] = vfiles
		}
	}
	return files
}
