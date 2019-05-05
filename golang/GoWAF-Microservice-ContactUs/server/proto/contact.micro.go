// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: contact.proto

/*
Package go_micro_srv_contact is a generated protocol buffer package.

It is generated from these files:
	contact.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	ViewRequest
	ViewResponse
	Empty
	ListResponse
	Contact
*/
package go_micro_srv_contact

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ContactMicro service

type ContactMicroService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	View(ctx context.Context, in *ViewRequest, opts ...client.CallOption) (*ViewResponse, error)
	List(ctx context.Context, in *Empty, opts ...client.CallOption) (*ListResponse, error)
}

type contactMicroService struct {
	c    client.Client
	name string
}

func NewContactMicroService(name string, c client.Client) ContactMicroService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.contact"
	}
	return &contactMicroService{
		c:    c,
		name: name,
	}
}

func (c *contactMicroService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "ContactMicro.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactMicroService) View(ctx context.Context, in *ViewRequest, opts ...client.CallOption) (*ViewResponse, error) {
	req := c.c.NewRequest(c.name, "ContactMicro.View", in)
	out := new(ViewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactMicroService) List(ctx context.Context, in *Empty, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "ContactMicro.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContactMicro service

type ContactMicroHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	View(context.Context, *ViewRequest, *ViewResponse) error
	List(context.Context, *Empty, *ListResponse) error
}

func RegisterContactMicroHandler(s server.Server, hdlr ContactMicroHandler, opts ...server.HandlerOption) error {
	type contactMicro interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		View(ctx context.Context, in *ViewRequest, out *ViewResponse) error
		List(ctx context.Context, in *Empty, out *ListResponse) error
	}
	type ContactMicro struct {
		contactMicro
	}
	h := &contactMicroHandler{hdlr}
	return s.Handle(s.NewHandler(&ContactMicro{h}, opts...))
}

type contactMicroHandler struct {
	ContactMicroHandler
}

func (h *contactMicroHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ContactMicroHandler.Create(ctx, in, out)
}

func (h *contactMicroHandler) View(ctx context.Context, in *ViewRequest, out *ViewResponse) error {
	return h.ContactMicroHandler.View(ctx, in, out)
}

func (h *contactMicroHandler) List(ctx context.Context, in *Empty, out *ListResponse) error {
	return h.ContactMicroHandler.List(ctx, in, out)
}