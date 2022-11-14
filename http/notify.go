package http

import (
	"errors"
	"github.com/valyala/fasthttp"
	"publish/consts"
)

func Notify() error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(consts.ApiNotifyUrl)
	req.Header.SetMethod("POST")
	req.Header.Set("Token", consts.PublisherToken)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK && resp.StatusCode() != fasthttp.StatusCreated {
		return err
	}
	res := string(resp.Body())
	if res != "ok" {
		return errors.New("error while sending message")
	}
	return nil
}
