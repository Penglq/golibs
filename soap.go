// Package soap provides a SOAP HTTP client.

package util

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

// XSINamespace is a link to the XML Schema instance namespace.
const XSINamespace = "http://www.w3.org/2001/XMLSchema-instance"

var xmlTyperType reflect.Type = reflect.TypeOf((*XMLTyper)(nil)).Elem()

// A RoundTripper executes a request passing the given req as the SOAP
// envelope body. The HTTP response is then de-serialized onto the resp
// object. Returns error in case an error occurs serializing req, making
// the HTTP request, or de-serializing the response.
type RoundTripper interface {
	RoundTrip(req, resp Message) error
	RoundTripSoap12(action string, req, resp Message) error
}

// Message is an opaque type used by the RoundTripper to carry XML
// documents for SOAP.
type Message interface{}

// Header is an opaque type used as the SOAP Header element in requests.
type Header interface{}

// AuthHeader is a Header to be encoded as the SOAP Header element in
// requests, to convey credentials for authentication.
type AuthHeader struct {
	Namespace string `xml:"xmlns:ns,attr"`
	Username  string `xml:"ns:username"`
	Password  string `xml:"ns:password"`
}

// Client is a SOAP client.
type Client struct {
	CTX                    context.Context
	URL                    string               // URL of the server
	Namespace              string               // SOAP Namespace
	ThisNamespace          string               // SOAP This-Namespace (tns)
	ExcludeActionNamespace bool                 // Include Namespace to SOAP Action header
	Envelope               string               // Optional SOAP Envelope
	NSA                    string               //
	Header                 Header               // Optional SOAP Header
	ContentType            string               // Optional Content-Type (default text/xml)
	Config                 *http.Client         // Optional HTTP client
	Pre                    func(*http.Request)  // Optional hook to modify outbound requests
	Post                   func(*http.Response) // Optional hook to snoop inbound responses
	TimeOut                time.Duration        // http请求超时时间
}

// XMLTyper is an abstract interface for types that can set an XML type.
type XMLTyper interface {
	SetXMLType()
}

func setXMLType(v reflect.Value) {
	if !v.IsValid() {
		return
	}
	switch v.Type().Kind() {
	case reflect.Interface:
		setXMLType(v.Elem())
	case reflect.Ptr:
		if v.IsNil() {
			break
		}
		ok := v.Type().Implements(xmlTyperType)
		if ok {
			v.MethodByName("SetXMLType").Call(nil)
		}
		setXMLType(v.Elem())
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			setXMLType(v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).CanAddr() {
				setXMLType(v.Field(i).Addr())
			} else {
				setXMLType(v.Field(i))
			}
		}
	}
}

func doRoundTrip(c *Client, setHeaders func(*http.Request), in Message, out interface{}) error {
	setXMLType(reflect.ValueOf(in))
	req := &Envelope{
		EnvelopeAttr: c.Envelope,
		//NSAttr:       c.Namespace,
		//TNSAttr:      c.ThisNamespace,
		//XSIAttr:      XSINamespace,
		Header: c.Header,
		//Body:   Ns1{Request: Att{Req: in, Ns1ATTR: c.NSA}},
		Body: in,
	}

	if req.EnvelopeAttr == "" {
		req.EnvelopeAttr = "http://schemas.xmlsoap.org/soap/envelope/"
	}
	//if req.Body.Request.Ns1ATTR == "" {
	//	req.Body.Request.Ns1ATTR = c.NSA
	//}
	//if req.TNSAttr == "" {
	//	req.TNSAttr = req.NSAttr
	//}
	var b bytes.Buffer
	err := xml.NewEncoder(&b).Encode(req)
	if err != nil {
		return err
	}
	QLog.GetLogger().Info("traceId", GetTraceIdFromCTX(c.CTX), "soap请求串", b.String())
	cli := c.Config
	if cli == nil {
		cli = http.DefaultClient
	}
	r, err := http.NewRequest("POST", c.URL, &b)
	if err != nil {
		return err
	}
	setHeaders(r)
	if c.Pre != nil {
		c.Pre(r)
	}
	if c.TimeOut.Seconds() != 0 {
		cli.Timeout = c.TimeOut
	}
	resp, err := cli.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if c.Post != nil {
		c.Post(resp)
	}
	if resp.StatusCode != http.StatusOK {
		// read only the first MiB of the body in error case
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("%s: %s", resp.Status, string(body)))
		//return &HTTPError{
		//	StatusResCode: resp.StatusCode,
		//	Status:     resp.Status,
		//	Msg:        string(body),
		//}
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		QLog.GetLogger().Info("traceId", GetTraceIdFromCTX(c.CTX), "soap响应串", string(body))
		return xml.Unmarshal(body, out)
	}
}

// RoundTrip implements the RoundTripper interface.
func (c *Client) RoundTrip(in Message, out *Ret) error {
	headerFunc := func(r *http.Request) {
		var actionName, soapAction string
		if in != nil {
			soapAction = reflect.TypeOf(in).Elem().Name()
		}
		ct := c.ContentType
		if ct == "" {
			ct = "text/xml"
		}
		r.Header.Set("Content-Type", ct)
		if in != nil {
			if c.ExcludeActionNamespace {
				actionName = soapAction
			} else {
				actionName = fmt.Sprintf("%s/%s", c.Namespace, soapAction)
			}
			r.Header.Add("SOAPAction", actionName)
		}
	}
	return doRoundTrip(c, headerFunc, in, out)
}

// RoundTripWithAction implements the RoundTripper interface for SOAP clients
// that need to set the SOAPAction header.
func (c *Client) RoundTripWithAction(soapAction string, in Message, out *Ret) error {
	headerFunc := func(r *http.Request) {
		var actionName string
		ct := c.ContentType
		if ct == "" {
			ct = "text/xml"
		}
		r.Header.Set("Content-Type", ct)
		if in != nil {
			if c.ExcludeActionNamespace {
				actionName = soapAction
			} else {
				actionName = fmt.Sprintf("%s/%s", c.Namespace, soapAction)
			}
			r.Header.Add("SOAPAction", actionName)
		}
	}
	return doRoundTrip(c, headerFunc, in, out)
}

// RoundTripSoap12 implements the RoundTripper interface for SOAP 1.2.
func (c *Client) RoundTripSoap12(action string, in Message, out interface{}) error {
	headerFunc := func(r *http.Request) {
		r.Header.Add("Content-Type", fmt.Sprintf("application/soap+xml; charset=utf-8; action=\"%s\"", action))
	}
	return doRoundTrip(c, headerFunc, in, out)
}

// HTTPError is detailed soap http error
type HTTPError struct {
	StatusCode int
	Status     string
	Msg        string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%q: %q", e.Status, e.Msg)
}

// 请求结构
type Envelope struct {
	XMLName      xml.Name    `xml:"soap:Envelope"`
	EnvelopeAttr string      `xml:"xmlns:soap,attr"`
	Header       Message     `xml:"soap:Header"`
	Body         interface{} `xml:"soap:Body"`
}
type Ns1 struct {
	Request Att `xml:"ns1:singleCertAndAccountVerify"`
}
type Att struct {
	Req     Message `xml:"request"`
	Ns1ATTR string  `xml:"xmlns:ns1,attr"`
}

// 返回结构
type Ret struct {
	XMLName xml.Name `xml:"Envelope"`
	//EnvelopeAttr string   `xml:"xmlns:,attr"`
	Body Ns1RES `xml:"Body"`
}
type Ns1RES struct {
	Return Retu `xml:"singleCertAndAccountVerifyResponse"`
}
type Retu struct {
	Req ResStruct `xml:"return"`
}
type ResStruct struct {
	Name string `xml:"name"`
}

// Envelope is a SOAP envelope.
//type Envelope struct {
//	XMLName      xml.Name `xml:"SOAP-ENV:Envelope"`
//	EnvelopeAttr string   `xml:"xmlns:SOAP-ENV,attr"`
//	NSAttr       string   `xml:"xmlns:ns,attr"`
//	TNSAttr      string   `xml:"xmlns:tns,attr"`
//	XSIAttr      string   `xml:"xmlns:xsi,attr,omitempty"`
//	Header       Message  `xml:"SOAP-ENV:Header"`
//	Body         Message  `xml:"SOAP-ENV:Body"`
//}
