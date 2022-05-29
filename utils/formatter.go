package utils

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

type Formatter struct {
	ctx *gin.Context
}

type IFormatter interface {
	FormatURL(string) (*url.URL, error)
}

func (formatter *Formatter) FormatURL(path string) (*url.URL, error) {
	req := formatter.ctx.Request

	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}

	url, err := url.Parse(scheme + "://" + req.URL.Host + req.URL.Path)
	if err != nil {
		defaultUrl, _ := url.Parse(req.URL.String() + path)
		return defaultUrl, nil
	}

	return url, nil
}

func NewFormatter(ctx *gin.Context) IFormatter {
	return &Formatter{
		ctx: ctx,
	}
}
