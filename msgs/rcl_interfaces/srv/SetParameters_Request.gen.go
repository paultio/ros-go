// Code generated by rclgo-gen. DO NOT EDIT.

package rcl_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	rcl_interfaces_msg "paultio/ros-go/msgs/rcl_interfaces/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/srv/set_parameters.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/SetParameters_Request", SetParameters_RequestTypeSupport)
	typemap.RegisterMessage("rcl_interfaces/srv/SetParameters_Request", SetParameters_RequestTypeSupport)
}

type SetParameters_Request struct {
	Parameters []rcl_interfaces_msg.Parameter `yaml:"parameters"`// A list of parameters to set.
}

// NewSetParameters_Request creates a new SetParameters_Request with default values.
func NewSetParameters_Request() *SetParameters_Request {
	self := SetParameters_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetParameters_Request) Clone() *SetParameters_Request {
	c := &SetParameters_Request{}
	if t.Parameters != nil {
		c.Parameters = make([]rcl_interfaces_msg.Parameter, len(t.Parameters))
		rcl_interfaces_msg.CloneParameterSlice(c.Parameters, t.Parameters)
	}
	return c
}

func (t *SetParameters_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetParameters_Request) SetDefaults() {
	t.Parameters = nil
}

func (t *SetParameters_Request) GetTypeSupport() types.MessageTypeSupport {
	return SetParameters_RequestTypeSupport
}

// SetParameters_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetParameters_RequestPublisher struct {
	*rclgo.Publisher
}

// NewSetParameters_RequestPublisher creates and returns a new publisher for the
// SetParameters_Request
func NewSetParameters_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetParameters_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetParameters_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetParameters_RequestPublisher{pub}, nil
}

func (p *SetParameters_RequestPublisher) Publish(msg *SetParameters_Request) error {
	return p.Publisher.Publish(msg)
}

// SetParameters_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetParameters_RequestSubscription struct {
	*rclgo.Subscription
}

// SetParameters_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a SetParameters_RequestSubscription.
type SetParameters_RequestSubscriptionCallback func(msg *SetParameters_Request, info *rclgo.MessageInfo, err error)

// NewSetParameters_RequestSubscription creates and returns a new subscription for the
// SetParameters_Request
func NewSetParameters_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback SetParameters_RequestSubscriptionCallback) (*SetParameters_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetParameters_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetParameters_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &SetParameters_RequestSubscription{sub}, nil
}

func (s *SetParameters_RequestSubscription) TakeMessage(out *SetParameters_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetParameters_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetParameters_RequestSlice(dst, src []SetParameters_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetParameters_RequestTypeSupport types.MessageTypeSupport = _SetParameters_RequestTypeSupport{}

type _SetParameters_RequestTypeSupport struct{}

func (t _SetParameters_RequestTypeSupport) New() types.Message {
	return NewSetParameters_Request()
}

func (t _SetParameters_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__SetParameters_Request
	return (unsafe.Pointer)(C.rcl_interfaces__srv__SetParameters_Request__create())
}

func (t _SetParameters_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__SetParameters_Request__destroy((*C.rcl_interfaces__srv__SetParameters_Request)(pointer_to_free))
}

func (t _SetParameters_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetParameters_Request)
	mem := (*C.rcl_interfaces__srv__SetParameters_Request)(dst)
	rcl_interfaces_msg.Parameter__Sequence_to_C((*rcl_interfaces_msg.CParameter__Sequence)(unsafe.Pointer(&mem.parameters)), m.Parameters)
}

func (t _SetParameters_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetParameters_Request)
	mem := (*C.rcl_interfaces__srv__SetParameters_Request)(ros2_message_buffer)
	rcl_interfaces_msg.Parameter__Sequence_to_Go(&m.Parameters, *(*rcl_interfaces_msg.CParameter__Sequence)(unsafe.Pointer(&mem.parameters)))
}

func (t _SetParameters_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__SetParameters_Request())
}

type CSetParameters_Request = C.rcl_interfaces__srv__SetParameters_Request
type CSetParameters_Request__Sequence = C.rcl_interfaces__srv__SetParameters_Request__Sequence

func SetParameters_Request__Sequence_to_Go(goSlice *[]SetParameters_Request, cSlice CSetParameters_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetParameters_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetParameters_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetParameters_Request__Sequence_to_C(cSlice *CSetParameters_Request__Sequence, goSlice []SetParameters_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__SetParameters_Request)(C.malloc(C.sizeof_struct_rcl_interfaces__srv__SetParameters_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetParameters_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetParameters_Request__Array_to_Go(goSlice []SetParameters_Request, cSlice []CSetParameters_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetParameters_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetParameters_Request__Array_to_C(cSlice []CSetParameters_Request, goSlice []SetParameters_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetParameters_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
