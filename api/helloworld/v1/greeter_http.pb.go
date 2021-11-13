// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 3.3.27

package v1

import (
	context "context"
	json "encoding/json"
	go_restful "github.com/emicklei/go-restful"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.go_restful.json.

type GreeterHTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

type GreeterHTTPHandler struct {
	srv GreeterHTTPServer
}

func newGreeterHTTPHandler(s GreeterHTTPServer) *GreeterHTTPHandler {
	return &GreeterHTTPHandler{srv: s}
}

func (h *GreeterHTTPHandler) SayHello(req *go_restful.Request, resp *go_restful.Response) {
	in := &HelloRequest{}
	if err := transportHTTP.GetQuery(req, in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.SayHello(req.Request.Context(), in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func RegisterGreeterHTTPServer(container *go_restful.Container, srv GreeterHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newGreeterHTTPHandler(srv)
	ws.Route(ws.GET("/helloworld/{name}").
		To(handler.SayHello))
}
