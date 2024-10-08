// Code generated by rclgo-gen. DO NOT EDIT.

package nav_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <nav_msgs/srv/set_map.h>
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
	typemap.RegisterService("nav_msgs/SetMap", SetMapTypeSupport)
	typemap.RegisterService("nav_msgs/srv/SetMap", SetMapTypeSupport)
}

type _SetMapTypeSupport struct {}

func (s _SetMapTypeSupport) Request() types.MessageTypeSupport {
	return SetMap_RequestTypeSupport
}

func (s _SetMapTypeSupport) Response() types.MessageTypeSupport {
	return SetMap_ResponseTypeSupport
}

func (s _SetMapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__nav_msgs__srv__SetMap())
}

// Modifying this variable is undefined behavior.
var SetMapTypeSupport types.ServiceTypeSupport = _SetMapTypeSupport{}

// SetMapClient wraps rclgo.Client to provide type safe helper
// functions
type SetMapClient struct {
	*rclgo.Client
}

// NewSetMapClient creates and returns a new client for the
// SetMap
func NewSetMapClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*SetMapClient, error) {
	client, err := node.NewClient(serviceName, SetMapTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetMapClient{client}, nil
}

func (s *SetMapClient) Send(ctx context.Context, req *SetMap_Request) (*SetMap_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*SetMap_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type SetMapServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s SetMapServiceResponseSender) SendResponse(resp *SetMap_Response) error {
	return s.sender.SendResponse(resp)
}

type SetMapServiceRequestHandler func(*rclgo.ServiceInfo, *SetMap_Request, SetMapServiceResponseSender)

// SetMapService wraps rclgo.Service to provide type safe helper
// functions
type SetMapService struct {
	*rclgo.Service
}

// NewSetMapService creates and returns a new service for the
// SetMap
func NewSetMapService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler SetMapServiceRequestHandler) (*SetMapService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*SetMap_Request)
		responseSender := SetMapServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, SetMapTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &SetMapService{service}, nil
}