// Code generated by rclgo-gen. DO NOT EDIT.

package geographic_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <geographic_msgs/srv/get_geo_path.h>
*/
import "C"

import (
	"context"
	"errors"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
)

func init() {
	typemap.RegisterService("geographic_msgs/GetGeoPath", GetGeoPathTypeSupport)
	typemap.RegisterService("geographic_msgs/srv/GetGeoPath", GetGeoPathTypeSupport)
}

type _GetGeoPathTypeSupport struct {}

func (s _GetGeoPathTypeSupport) Request() types.MessageTypeSupport {
	return GetGeoPath_RequestTypeSupport
}

func (s _GetGeoPathTypeSupport) Response() types.MessageTypeSupport {
	return GetGeoPath_ResponseTypeSupport
}

func (s _GetGeoPathTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__geographic_msgs__srv__GetGeoPath())
}

// Modifying this variable is undefined behavior.
var GetGeoPathTypeSupport types.ServiceTypeSupport = _GetGeoPathTypeSupport{}

// GetGeoPathClient wraps rclgo.Client to provide type safe helper
// functions
type GetGeoPathClient struct {
	*rclgo.Client
}

// NewGetGeoPathClient creates and returns a new client for the
// GetGeoPath
func NewGetGeoPathClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*GetGeoPathClient, error) {
	client, err := node.NewClient(serviceName, GetGeoPathTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GetGeoPathClient{client}, nil
}

func (s *GetGeoPathClient) Send(ctx context.Context, req *GetGeoPath_Request) (*GetGeoPath_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*GetGeoPath_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type GetGeoPathServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s GetGeoPathServiceResponseSender) SendResponse(resp *GetGeoPath_Response) error {
	return s.sender.SendResponse(resp)
}

type GetGeoPathServiceRequestHandler func(*rclgo.ServiceInfo, *GetGeoPath_Request, GetGeoPathServiceResponseSender)

// GetGeoPathService wraps rclgo.Service to provide type safe helper
// functions
type GetGeoPathService struct {
	*rclgo.Service
}

// NewGetGeoPathService creates and returns a new service for the
// GetGeoPath
func NewGetGeoPathService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler GetGeoPathServiceRequestHandler) (*GetGeoPathService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*GetGeoPath_Request)
		responseSender := GetGeoPathServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, GetGeoPathTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &GetGeoPathService{service}, nil
}