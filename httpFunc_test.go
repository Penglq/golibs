package golibs

import (
	"context"
	"github.com/Penglq/QLog"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func init() {
	QLog.GetLogger().SetConfig(QLog.DEBUG, "", QLog.WithConsoleOPT())
}
func TestHttpmock(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		QLog.GetLogger().Info("request", *r)
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":9999", nil)
}

func TestHttpSend_HttpGet(t *testing.T) {
	ctx := context.Background()
	res, err := NewHttpClient(&ctx, LoggerOptions(QLog.GetLogger()), UrlOptions("http://localhost:9999/"), TryOptions(2, time.Second*3)).HttpGet()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestHttpSend_HttpPost(t *testing.T) {
	ctx := context.Background()
	res, err := NewHttpClient(&ctx, LoggerOptions(QLog.GetLogger()), UrlOptions("http://localhost:9999/"), PostValueOptions("hello"), TryOptions(2, time.Second*3)).HttpPost()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestHttpSend_HttpPostForm(t *testing.T) {
	ctx := context.Background()
	data := url.Values{}
	data.Set("key", "value")
	res, err := NewHttpClient(&ctx, LoggerOptions(QLog.GetLogger()), UrlOptions("http://localhost:9999/"), PostFormValueOptions(data), TryOptions(2, time.Second*3)).HttpPostForm()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
