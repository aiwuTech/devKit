// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package controller

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/globalways/response"
	"net/url"
	"sort"
	"time"
)

var (
	err_miss_appid        = "缺少参数appid."
	err_invalid_appid     = "参数appid不正确."
	err_miss_sign         = "缺少参数signature."
	err_invalid_sign      = "请求签名不正确."
	err_miss_timestamp    = "缺少参数timestamp."
	err_format_timestamp  = "参数timestamp格式错误,应为2006-01-02 15:04:05."
	err_invalid_timestamp = "请求已过期."
)

func serveJson(ctx *context.Context, data interface{}) {
	ctx.Output.Json(data, false, false)
}

type AppIdToAppSecret func(string) string

func APIBaiscAuth(appid, appkey string) beego.FilterFunc {
	ft := func(aid string) string {
		if aid == appid {
			return appkey
		}
		return ""
	}
	return APIAuthWithFunc(ft, 300)
}

func APIAuthWithFunc(f AppIdToAppSecret, timeout int) beego.FilterFunc {
	return func(ctx *context.Context) {
		if ctx.Input.IsOptions() {
			return
		}

		ctx.Input.Request.ParseForm()
		beego.Debug(ctx.Input.Request.Form)
		if ctx.Input.Query("appid") == "" {
			beego.Debug("miss query param: appid")
			rspMsg := response.NewResponseMsgInvalidAuth(err_miss_appid)
			serveJson(ctx, rspMsg)
			return
		}
		appsecret := f(ctx.Input.Query("appid"))
		if appsecret == "" {
			beego.Debug("not exist this appid")
			rspMsg := response.NewResponseMsgInvalidAuth(err_invalid_appid)
			serveJson(ctx, rspMsg)
			return
		}
		signature := ctx.Input.Query("signature")
		if signature == "" {
			beego.Debug("miss query param: signature")
			rspMsg := response.NewResponseMsgInvalidAuth(err_miss_sign)
			serveJson(ctx, rspMsg)
			return
		}
		if ctx.Input.Query("timestamp") == "" {
			beego.Debug("miss query param: timestamp")
			rspMsg := response.NewResponseMsgInvalidAuth(err_miss_timestamp)
			serveJson(ctx, rspMsg)
			return
		}
		u, err := time.Parse("2006-01-02 15:04:05", ctx.Input.Query("timestamp"))
		if err != nil {
			beego.Debug("timestamp format is error, should 2006-01-02 15:04:05")
			rspMsg := response.NewResponseMsgInvalidAuth(err_format_timestamp)
			serveJson(ctx, rspMsg)
			return
		}
		t := time.Now()
		if t.Sub(u).Seconds() > float64(timeout) {
			beego.Debug("timeout! the request time is long ago, please try again")
			rspMsg := response.NewResponseMsgInvalidAuth(err_invalid_timestamp)
			serveJson(ctx, rspMsg)
			return
		}
		result := Signature(appsecret, ctx.Input.Method(), ctx.Request.Form, ctx.Input.Url())
		beego.Debug("ctx.Input.Query('signature'):", signature)
		beego.Debug("ctx.Input.Uri():", ctx.Input.Uri())
		beego.Debug("Signature:", result)
		if signature != result {
			beego.Debug("auth failed")
			rspMsg := response.NewResponseMsgInvalidAuth(err_invalid_sign)
			serveJson(ctx, rspMsg)
			return
		}
	}
}

func Signature(appsecret, method string, params url.Values, requestURL string) (result string) {
	var query string
	pa := make(map[string]string)
	for k, v := range params {
		pa[k] = v[0]
	}
	vs := mapSorter(pa)
	vs.Sort()
	for i := 0; i < vs.Len(); i++ {
		if vs.Keys[i] == "signature" {
			continue
		}
		if vs.Keys[i] != "" && vs.Vals[i] != "" {
			query = fmt.Sprintf("%v%v%v", query, vs.Keys[i], vs.Vals[i])
		}
	}
	beego.Debug("query:", query)
	string_to_sign := fmt.Sprintf("%v\n%v\n%v\n", method, query, requestURL)
	beego.Debug("string to sign:", string_to_sign)

	sha256 := sha256.New
	hash := hmac.New(sha256, []byte(appsecret))
	hash.Write([]byte(string_to_sign))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

type valSorter struct {
	Keys []string
	Vals []string
}

func mapSorter(m map[string]string) *valSorter {
	vs := &valSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]string, 0, len(m)),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *valSorter) Sort() {
	sort.Sort(vs)
}

func (vs *valSorter) Len() int           { return len(vs.Keys) }
func (vs *valSorter) Less(i, j int) bool { return vs.Keys[i] < vs.Keys[j] }
func (vs *valSorter) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}
