package utils

import (
	"context"
	"errors"
	"github.com/penglq/QLog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpSend struct {
	ctx                                    *context.Context
	client                                 *http.Client
	url, postValue, httpProxy, contentType string
	postFormValue                          url.Values
	timeout                                time.Duration
	tryNum                                 int
	doNum                                  int
	intervalTime                           time.Duration
}

func NewHttpClient(ctx *context.Context, options ...HttpOptions) *HttpSend {
	client := new(HttpSend)
	client.ctx = ctx
	for i := 0; i < len(options); i++ {
		options[i](client)
	}
	client.NewClient(ctx)
	if client.contentType == "" {
		client.contentType = "application/x-www-form-urlencoded"
	}

	if client.timeout != 0 {
		client.client.Timeout = client.timeout
	}

	return client
}

func (h *HttpSend) NewClient(ctx *context.Context) {
	h.client = new(http.Client)
	if h.httpProxy != "" {
		proxyVal, err := url.Parse(h.httpProxy)
		if err != nil {
			QLog.GetLogger().Alert(TraceKey, GetTraceIdFromCTX(*ctx), "method", "NewHttpClient", "action", "初始化httpClient-httpProxy解析错误", "error", err)
			panic("httpProxy解析错误")
		}
		h.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyVal),
		}
	}
}

// 使用config中httpProxy即可
func (h *HttpSend) HttpGet() ([]byte, error) {
	defer func(t time.Time) {
		QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "took/s", time.Since(t).Seconds())
	}(time.Now())
	QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "request", h.url)
	var (
		err  error = nil
		body []byte
		resp = new(http.Response)
	)

	resp, err = h.client.Get(h.url)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, err = ioutil.ReadAll(resp.Body)
		} else {
			err = errors.New(resp.Status)
		}
	} else {
		for h.doNum < h.tryNum {
			h.doNum++
			QLog.GetLogger().Alert(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "重试次数", h.doNum, "error", err)
			time.Sleep(h.intervalTime)
			h.NewClient(h.ctx)
			body, err = h.HttpGet()
			QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", err)
			return body, err
		}
	}
	QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", err)
	return body, err
}

func (h *HttpSend) HttpPost() (string, error) {
	QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "request", h.postValue)
	var (
		err  error = nil
		body []byte
		resp = new(http.Response)
	)

	resp, err = h.client.Post(h.url, h.contentType, strings.NewReader(h.postValue))
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, err = ioutil.ReadAll(resp.Body)
		} else {
			err = errors.New(resp.Status)
		}
	} else {
		for h.doNum < h.tryNum {
			h.doNum++
			QLog.GetLogger().Alert(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "重试次数", h.doNum, "error", err)
			time.Sleep(h.intervalTime)
			h.NewClient(h.ctx)
			str := ""
			str, err = h.HttpPost()
			QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", str, "error", err)
			return str, err
		}
	}
	QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", err)
	return string(body), err
}

func (h *HttpSend) HttpPostForm() (string, error) {
	QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "request", h.postValue)
	var (
		err  error = nil
		body []byte
		resp = new(http.Response)
	)
	resp, err = h.client.PostForm(h.url, h.postFormValue)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, err = ioutil.ReadAll(resp.Body)
		} else {
			err = errors.New(resp.Status)
		}
	} else {
		for h.doNum < h.tryNum {
			h.doNum++
			QLog.GetLogger().Alert(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "重试次数", h.doNum, "error", err)
			time.Sleep(h.intervalTime)
			h.NewClient(h.ctx)
			str := ""
			str, err = h.HttpPostForm()
			QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", str, "error", err)
			return h.HttpPostForm()
		}
	}
	QLog.GetLogger().Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", err)
	return string(body), err
}

type HttpOptions func(*HttpSend)

func UrlOptions(url string) HttpOptions {
	return func(client *HttpSend) {
		client.url = url
	}
}
func PostValueOptions(postValue string) HttpOptions {
	return func(client *HttpSend) {
		client.postValue = postValue
	}
}

func PostFormValueOptions(postFormValue url.Values) HttpOptions {
	return func(client *HttpSend) {
		client.postFormValue = postFormValue
	}
}

func TimeOptions(timeout time.Duration) HttpOptions {
	return func(client *HttpSend) {
		client.timeout = timeout
	}
}

func ProxyOptions(proxy string) HttpOptions {
	return func(client *HttpSend) {
		client.httpProxy = proxy
	}
}

func ContentTypeOptions(contentType string) HttpOptions {
	return func(client *HttpSend) {
		client.contentType = contentType
	}
}

func TryOptions(tryNum int, intervalTime time.Duration) HttpOptions {
	return func(client *HttpSend) {
		client.tryNum = tryNum
		client.intervalTime = intervalTime
	}
}
