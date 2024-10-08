// Code generated by rclgo-gen. DO NOT EDIT.

package diagnostic_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <diagnostic_msgs/srv/self_test.h>
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
	typemap.RegisterService("diagnostic_msgs/SelfTest", SelfTestTypeSupport)
	typemap.RegisterService("diagnostic_msgs/srv/SelfTest", SelfTestTypeSupport)
}

type _SelfTestTypeSupport struct {}

func (s _SelfTestTypeSupport) Request() types.MessageTypeSupport {
	return SelfTest_RequestTypeSupport
}

func (s _SelfTestTypeSupport) Response() types.MessageTypeSupport {
	return SelfTest_ResponseTypeSupport
}

func (s _SelfTestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__diagnostic_msgs__srv__SelfTest())
}

// Modifying this variable is undefined behavior.
var SelfTestTypeSupport types.ServiceTypeSupport = _SelfTestTypeSupport{}

// SelfTestClient wraps rclgo.Client to provide type safe helper
// functions
type SelfTestClient struct {
	*rclgo.Client
}

// NewSelfTestClient creates and returns a new client for the
// SelfTest
func NewSelfTestClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*SelfTestClient, error) {
	client, err := node.NewClient(serviceName, SelfTestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SelfTestClient{client}, nil
}

func (s *SelfTestClient) Send(ctx context.Context, req *SelfTest_Request) (*SelfTest_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*SelfTest_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type SelfTestServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s SelfTestServiceResponseSender) SendResponse(resp *SelfTest_Response) error {
	return s.sender.SendResponse(resp)
}

type SelfTestServiceRequestHandler func(*rclgo.ServiceInfo, *SelfTest_Request, SelfTestServiceResponseSender)

// SelfTestService wraps rclgo.Service to provide type safe helper
// functions
type SelfTestService struct {
	*rclgo.Service
}

// NewSelfTestService creates and returns a new service for the
// SelfTest
func NewSelfTestService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler SelfTestServiceRequestHandler) (*SelfTestService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*SelfTest_Request)
		responseSender := SelfTestServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, SelfTestTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &SelfTestService{service}, nil
}