package gee

type HttpMethod string

const (
	GETMethod     HttpMethod = "GET"
	POSTMethod    HttpMethod = "POST"
	DELETEMethod  HttpMethod = "DELETE"
	PUTMethod     HttpMethod = "PUT"
	HEADMethod    HttpMethod = "HEAD"
	OPTIONSMethod HttpMethod = "OPTIONS"
	PATCHMethod   HttpMethod = "PATCH"
)
