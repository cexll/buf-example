// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             (unknown)
// source: demo/v1/demo.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDemoServiceCreate = "/demo.api.v1.DemoService/Create"

type DemoServiceHTTPServer interface {
	Create(context.Context, *DemoServiceCreateRequest) (*DemoServiceCreateResponse, error)
}

func RegisterDemoServiceHTTPServer(s *http.Server, srv DemoServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/demo/create", _DemoService_Create0_HTTP_Handler(srv))
}

func _DemoService_Create0_HTTP_Handler(srv DemoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DemoServiceCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoServiceCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*DemoServiceCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DemoServiceCreateResponse)
		return ctx.Result(200, reply)
	}
}

type DemoServiceHTTPClient interface {
	Create(ctx context.Context, req *DemoServiceCreateRequest, opts ...http.CallOption) (rsp *DemoServiceCreateResponse, err error)
}

type DemoServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDemoServiceHTTPClient(client *http.Client) DemoServiceHTTPClient {
	return &DemoServiceHTTPClientImpl{client}
}

func (c *DemoServiceHTTPClientImpl) Create(ctx context.Context, in *DemoServiceCreateRequest, opts ...http.CallOption) (*DemoServiceCreateResponse, error) {
	var out DemoServiceCreateResponse
	pattern := "/v1/demo/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDemoServiceCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}