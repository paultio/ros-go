// Code generated by rclgo-gen. DO NOT EDIT.

package nav_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <nav_msgs/srv/load_map.h>
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
	typemap.RegisterService("nav_msgs/LoadMap", LoadMapTypeSupport)
	typemap.RegisterService("nav_msgs/srv/LoadMap", LoadMapTypeSupport)
}

type _LoadMapTypeSupport struct {}

func (s _LoadMapTypeSupport) Request() types.MessageTypeSupport {
	return LoadMap_RequestTypeSupport
}

func (s _LoadMapTypeSupport) Response() types.MessageTypeSupport {
	return LoadMap_ResponseTypeSupport
}

func (s _LoadMapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__nav_msgs__srv__LoadMap())
}

// Modifying this variable is undefined behavior.
var LoadMapTypeSupport types.ServiceTypeSupport = _LoadMapTypeSupport{}

// LoadMapClient wraps rclgo.Client to provide type safe helper
// functions
type LoadMapClient struct {
	*rclgo.Client
}

// NewLoadMapClient creates and returns a new client for the
// LoadMap
func NewLoadMapClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*LoadMapClient, error) {
	client, err := node.NewClient(serviceName, LoadMapTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &LoadMapClient{client}, nil
}

func (s *LoadMapClient) Send(ctx context.Context, req *LoadMap_Request) (*LoadMap_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*LoadMap_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type LoadMapServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s LoadMapServiceResponseSender) SendResponse(resp *LoadMap_Response) error {
	return s.sender.SendResponse(resp)
}

type LoadMapServiceRequestHandler func(*rclgo.ServiceInfo, *LoadMap_Request, LoadMapServiceResponseSender)

// LoadMapService wraps rclgo.Service to provide type safe helper
// functions
type LoadMapService struct {
	*rclgo.Service
}

// NewLoadMapService creates and returns a new service for the
// LoadMap
func NewLoadMapService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler LoadMapServiceRequestHandler) (*LoadMapService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*LoadMap_Request)
		responseSender := LoadMapServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, LoadMapTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &LoadMapService{service}, nil
}