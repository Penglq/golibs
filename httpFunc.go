package golibs

import (
	"context"
	"errors"
	"github.com/Penglq/QLog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	contentType = "Content-Type"
)

type HttpSend struct {
	ctx                       *context.Context
	client                    *http.Client
	request                   *http.Request
	log                       QLog.LoggerInterface
	url, postValue, httpProxy string
	header                    map[string]string
	postFormValue             url.Values
	timeout                   time.Duration
	tryNum                    int
	doNum                     int
	intervalTime              time.Duration
	err                       error
}

func NewHttpClient(ctx *context.Context, options ...HttpOptions) *HttpSend {
	client := new(HttpSend)
	client.ctx = ctx
	for i := 0; i < len(options); i++ {
		options[i](client)
	}
	client.NewClient(ctx)
	if client.timeout != 0 {
		client.client.Timeout = client.timeout
	}
	return client
}

func (h *HttpSend) NewClient(ctx *context.Context) {
	h.client = new(http.Client)
	if h.httpProxy != "" {
		var proxyUrl *url.URL
		proxyUrl, h.err = url.Parse(h.httpProxy)
		if h.err != nil {
			h.log.AlertWithLevel(QLog.ALERTALERT, TraceKey, GetTraceIdFromCTX(*ctx), "method", "NewHttpClient", "action", "初始化httpClient-httpProxy解析错误", "error", h.err)
		}
		h.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}
}
func (h *HttpSend) HttpDo() []byte {
	var resp *http.Response
	var body []byte
	resp, h.err = h.client.Do(h.request)
	if h.err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, h.err = ioutil.ReadAll(resp.Body)
		} else {
			h.err = errors.New(resp.Status)
		}
	} else {
		for h.doNum < h.tryNum {
			h.doNum++
			h.log.AlertWithLevel(QLog.ALERTALERT, TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "重试次数", h.doNum, "error", h.err)
			time.Sleep(h.intervalTime)
			h.NewClient(h.ctx)
			body = h.HttpDo()
			// h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", h.err)
			return body
		}
	}
	return body
}

// 使用config中httpProxy即可
func (h *HttpSend) HttpGet() ([]byte, error) {
	if h.err != nil {
		return []byte{}, h.err
	}
	defer func(t time.Time) {
		h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "took/s", time.Since(t).Seconds())
	}(time.Now())
	h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "request", h.url)
	h.request, h.err = http.NewRequest(http.MethodGet, h.url, nil)
	if h.err != nil {
		return []byte{}, h.err
	}
	h.setHeader()
	var body []byte
	body = h.HttpDo()
	h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", h.err)
	return body, h.err
}

func (h *HttpSend) HttpPost() ([]byte, error) {
	if h.err != nil {
		return []byte{}, h.err
	}
	h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "request", h.postValue)
	h.request, h.err = http.NewRequest(http.MethodPost, h.url, strings.NewReader(h.postValue))
	if _, ok := h.header[contentType]; !ok {
		h.request.Header.Set(contentType, "application/json")
	}
	h.setHeader()
	var body []byte
	body = h.HttpDo()
	h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", h.err)
	return body, h.err
}

func (h *HttpSend) HttpPostForm() ([]byte, error) {
	if h.err != nil {
		return []byte{}, h.err
	}
	h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "request", h.postValue)
	h.request, h.err = http.NewRequest(http.MethodPost, h.url, strings.NewReader(h.postFormValue.Encode()))
	if h.err != nil {
		return []byte{}, h.err
	}
	h.request.Header.Set(contentType, "application/x-www-form-urlencoded")
	var body []byte
	body = h.HttpDo()
	h.log.Info(TraceKey, GetTraceIdFromCTX(*(h.ctx)), "url", h.url, "response", string(body), "error", h.err)
	return body, h.err
}
func (h *HttpSend) setHeader() {
	for v, k := range h.header {
		h.request.Header.Set(k, v)
	}
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

func HeaderOptions(header map[string]string) HttpOptions {
	return func(client *HttpSend) {
		client.header = header
	}
}

func TryOptions(tryNum int, intervalTime time.Duration) HttpOptions {
	return func(client *HttpSend) {
		client.tryNum = tryNum
		client.intervalTime = intervalTime
	}
}

func LoggerOptions(l QLog.LoggerInterface) HttpOptions {
	return func(client *HttpSend) {
		client.log = l
	}
}
