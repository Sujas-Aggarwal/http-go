package helper

type RequestHeader struct {
	Method        string
	Route         string
	Version       string
	Host          string
	UserAgent     string
	Accept        string
	ContentLength int
	ContentType   string
}

type Request struct {
	headers RequestHeader
	body    string
}
