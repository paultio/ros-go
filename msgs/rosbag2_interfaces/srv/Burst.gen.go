// Code generated by rclgo-gen. DO NOT EDIT.

package rosbag2_interfaces_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rosbag2_interfaces/srv/burst.h>
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
	typemap.RegisterService("rosbag2_interfaces/Burst", BurstTypeSupport)
	typemap.RegisterService("rosbag2_interfaces/srv/Burst", BurstTypeSupport)
}

type _BurstTypeSupport struct {}

func (s _BurstTypeSupport) Request() types.MessageTypeSupport {
	return Burst_RequestTypeSupport
}

func (s _BurstTypeSupport) Response() types.MessageTypeSupport {
	return Burst_ResponseTypeSupport
}

func (s _BurstTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__rosbag2_interfaces__srv__Burst())
}

// Modifying this variable is undefined behavior.
var BurstTypeSupport types.ServiceTypeSupport = _BurstTypeSupport{}

// BurstClient wraps rclgo.Client to provide type safe helper
// functions
type BurstClient struct {
	*rclgo.Client
}

// NewBurstClient creates and returns a new client for the
// Burst
func NewBurstClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*BurstClient, error) {
	client, err := node.NewClient(serviceName, BurstTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BurstClient{client}, nil
}

func (s *BurstClient) Send(ctx context.Context, req *Burst_Request) (*Burst_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Burst_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type BurstServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s BurstServiceResponseSender) SendResponse(resp *Burst_Response) error {
	return s.sender.SendResponse(resp)
}

type BurstServiceRequestHandler func(*rclgo.ServiceInfo, *Burst_Request, BurstServiceResponseSender)

// BurstService wraps rclgo.Service to provide type safe helper
// functions
type BurstService struct {
	*rclgo.Service
}

// NewBurstService creates and returns a new service for the
// Burst
func NewBurstService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler BurstServiceRequestHandler) (*BurstService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Burst_Request)
		responseSender := BurstServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, BurstTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &BurstService{service}, nil
}