// Package for dingding message robots
package dingding

import (
	"fmt"
	"time"
    "crypto/hmac"
	"crypto/sha256"
    "encoding/base64"
    "net/url"
    "encoding/json"
    "github.com/marknown/ohttp"
)

// Dingding struct
type Dingding struct {
    Token string
    Secret string
}

// Response of dingding
type Response struct {
	Errcode int
	Errmsg  string
}

func (d *Dingding) getSign(timestamp int64) string {
    preSign := fmt.Sprintf("%d\n%s", timestamp, d.Secret)
    h := hmac.New(sha256.New, []byte(d.Secret))
    h.Write([]byte(preSign))
	return url.QueryEscape(base64.StdEncoding.EncodeToString([]byte(h.Sum(nil))))
}

func (d *Dingding) getURL() string {
    timestamp := time.Now().UnixNano()/1000000
    sign := d.getSign(timestamp);
    return fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", d.Token, timestamp, sign);
}

func (d *Dingding) request(raw string) *Response {
    result := &Response{}

	settings := ohttp.InitSetttings()
	settings.ContentType = "application/json;charset=utf-8"
	settings.Timeout = 10 * time.Second

	content, _, err := settings.Post(d.getURL(), raw)
	if nil != err {
        result.Errcode = 999999
        result.Errmsg  = err.Error()
		return result
    }

    err = json.Unmarshal([]byte(content), result)
    if nil != err {
        result.Errcode = 999999
        result.Errmsg  = err.Error()
		return result
    }

    return result
}

// NotifyText notify text to dingding robots
func (d *Dingding) NotifyText(text string) *Response {
    raw := fmt.Sprintf(`{"msgtype":"text","text":{"content":"%s"}}`, text)
    return d.request(raw);
}

// NotifyLink notify link to dingding robots
func (d *Dingding) NotifyLink(title string, text string, messageURL string, picURL string) *Response {
    raw := fmt.Sprintf(`{"msgtype":"link","link":{"text":"%s","title":"%s","picUrl":"%s","messageUrl":"%s"}}`, text, title, picURL, messageURL)
    return d.request(raw);
}