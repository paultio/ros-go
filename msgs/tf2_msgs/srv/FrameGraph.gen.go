// Code generated by rclgo-gen. DO NOT EDIT.

package tf2_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <tf2_msgs/srv/frame_graph.h>
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
	typemap.RegisterService("tf2_msgs/FrameGraph", FrameGraphTypeSupport)
	typemap.RegisterService("tf2_msgs/srv/FrameGraph", FrameGraphTypeSupport)
}

type _FrameGraphTypeSupport struct {}

func (s _FrameGraphTypeSupport) Request() types.MessageTypeSupport {
	return FrameGraph_RequestTypeSupport
}

func (s _FrameGraphTypeSupport) Response() types.MessageTypeSupport {
	return FrameGraph_ResponseTypeSupport
}

func (s _FrameGraphTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__tf2_msgs__srv__FrameGraph())
}

// Modifying this variable is undefined behavior.
var FrameGraphTypeSupport types.ServiceTypeSupport = _FrameGraphTypeSupport{}

// FrameGraphClient wraps rclgo.Client to provide type safe helper
// functions
type FrameGraphClient struct {
	*rclgo.Client
}

// NewFrameGraphClient creates and returns a new client for the
// FrameGraph
func NewFrameGraphClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*FrameGraphClient, error) {
	client, err := node.NewClient(serviceName, FrameGraphTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FrameGraphClient{client}, nil
}

func (s *FrameGraphClient) Send(ctx context.Context, req *FrameGraph_Request) (*FrameGraph_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*FrameGraph_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type FrameGraphServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s FrameGraphServiceResponseSender) SendResponse(resp *FrameGraph_Response) error {
	return s.sender.SendResponse(resp)
}

type FrameGraphServiceRequestHandler func(*rclgo.ServiceInfo, *FrameGraph_Request, FrameGraphServiceResponseSender)

// FrameGraphService wraps rclgo.Service to provide type safe helper
// functions
type FrameGraphService struct {
	*rclgo.Service
}

// NewFrameGraphService creates and returns a new service for the
// FrameGraph
func NewFrameGraphService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler FrameGraphServiceRequestHandler) (*FrameGraphService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*FrameGraph_Request)
		responseSender := FrameGraphServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, FrameGraphTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &FrameGraphService{service}, nil
}