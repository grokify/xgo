package anyhttp

import (
	"io"
	"mime/multipart"
	"net"
	"net/url"

	"github.com/grokify/mogo/net/http/httputilmore"
)

type Request interface {
	Header(s string) []byte
	HeaderString(s string) string
	RemoteAddr() net.Addr
	RemoteAddress() string
	UserAgent() []byte
	Method() []byte
	ParseForm() error
	AllArgs() Args
	QueryArgs() Args
	PostArgs() Args
	MultipartForm() (*multipart.Form, error)
	RequestURI() []byte
	RequestURIVar(s string) string
	PostBody() ([]byte, error)
}

type Args interface {
	GetBytes(key string) []byte
	GetBytesSlice(key string) [][]byte
	GetString(key string) string
	GetStringSlice(key string) []string
	GetURLValues() url.Values
}

type Response interface {
	SetStatusCode(int)
	SetContentType(string)
	SetCookie(cookie *Cookie)
	GetHeader(key string) []byte
	SetHeader(key, val string)
	SetBodyBytes([]byte) (int, error)
	SetBodyStream(bodyStream io.Reader, bodySize int) error
}

func WriteSimpleJSON(w Response, status int, message string) (int, error) {
	w.SetStatusCode(status)
	w.SetContentType(httputilmore.ContentTypeAppJSONUtf8)
	resInfo := httputilmore.ResponseInfo{
		StatusCode: status,
		Body:       message}
	return w.SetBodyBytes(resInfo.ToJSON())
}

type MapStringString map[string]string

func (m MapStringString) Get(key string) string {
	if val, ok := m[key]; ok {
		return val
	}
	return ""
}

func (m MapStringString) GetSlice(key string) []string {
	return []string{m.Get(key)}
}

type ArgsMapStringString struct{ Raw MapStringString }

func NewArgsMapStringString(args MapStringString) ArgsMapStringString {
	return ArgsMapStringString{Raw: args}
}

func (args ArgsMapStringString) GetBytes(key string) []byte { return []byte(args.Raw.Get(key)) }
func (args ArgsMapStringString) GetBytesSlice(key string) [][]byte {
	output := make([][]byte, 1)
	output[0] = args.GetBytes(key)
	return output
}
func (args ArgsMapStringString) GetString(key string) string        { return args.Raw.Get(key) }
func (args ArgsMapStringString) GetStringSlice(key string) []string { return args.Raw.GetSlice(key) }

type ArgsURLValues struct{ Raw url.Values }

func NewArgsURLValues(args url.Values) ArgsURLValues {
	return ArgsURLValues{Raw: args}
}

func (args ArgsURLValues) GetBytes(key string) []byte { return []byte(args.Raw.Get(key)) }
func (args ArgsURLValues) GetBytesSlice(key string) [][]byte {
	newSlice := [][]byte{}
	if slice, ok := args.Raw[key]; ok {
		for _, item := range slice {
			newSlice = append(newSlice, []byte(item))
		}
	}
	return newSlice
}

func (args ArgsURLValues) GetString(key string) string { return args.Raw.Get(key) }
func (args ArgsURLValues) GetStringSlice(key string) []string {
	if slice, ok := args.Raw[key]; ok {
		return slice
	}
	return []string{}
}

func (args ArgsURLValues) GetURLValues() url.Values {
	return args.Raw
}

type Addr struct {
	Protocol string
	Address  string
}

func (addr Addr) Network() string { return addr.Protocol }
func (addr Addr) String() string  { return addr.Address }
