// Code generated by rclgo-gen. DO NOT EDIT.

package lifecycle_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	lifecycle_msgs_msg "paultio/ros-go/msgs/lifecycle_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <lifecycle_msgs/srv/change_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("lifecycle_msgs/ChangeState_Request", ChangeState_RequestTypeSupport)
	typemap.RegisterMessage("lifecycle_msgs/srv/ChangeState_Request", ChangeState_RequestTypeSupport)
}

type ChangeState_Request struct {
	Transition lifecycle_msgs_msg.Transition `yaml:"transition"`// The requested transition.This change state service will fail if the transition is not possible.
}

// NewChangeState_Request creates a new ChangeState_Request with default values.
func NewChangeState_Request() *ChangeState_Request {
	self := ChangeState_Request{}
	self.SetDefaults()
	return &self
}

func (t *ChangeState_Request) Clone() *ChangeState_Request {
	c := &ChangeState_Request{}
	c.Transition = *t.Transition.Clone()
	return c
}

func (t *ChangeState_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ChangeState_Request) SetDefaults() {
	t.Transition.SetDefaults()
}

func (t *ChangeState_Request) GetTypeSupport() types.MessageTypeSupport {
	return ChangeState_RequestTypeSupport
}

// ChangeState_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ChangeState_RequestPublisher struct {
	*rclgo.Publisher
}

// NewChangeState_RequestPublisher creates and returns a new publisher for the
// ChangeState_Request
func NewChangeState_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ChangeState_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ChangeState_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ChangeState_RequestPublisher{pub}, nil
}

func (p *ChangeState_RequestPublisher) Publish(msg *ChangeState_Request) error {
	return p.Publisher.Publish(msg)
}

// ChangeState_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ChangeState_RequestSubscription struct {
	*rclgo.Subscription
}

// ChangeState_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a ChangeState_RequestSubscription.
type ChangeState_RequestSubscriptionCallback func(msg *ChangeState_Request, info *rclgo.MessageInfo, err error)

// NewChangeState_RequestSubscription creates and returns a new subscription for the
// ChangeState_Request
func NewChangeState_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ChangeState_RequestSubscriptionCallback) (*ChangeState_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ChangeState_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ChangeState_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ChangeState_RequestSubscription{sub}, nil
}

func (s *ChangeState_RequestSubscription) TakeMessage(out *ChangeState_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneChangeState_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneChangeState_RequestSlice(dst, src []ChangeState_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ChangeState_RequestTypeSupport types.MessageTypeSupport = _ChangeState_RequestTypeSupport{}

type _ChangeState_RequestTypeSupport struct{}

func (t _ChangeState_RequestTypeSupport) New() types.Message {
	return NewChangeState_Request()
}

func (t _ChangeState_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.lifecycle_msgs__srv__ChangeState_Request
	return (unsafe.Pointer)(C.lifecycle_msgs__srv__ChangeState_Request__create())
}

func (t _ChangeState_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.lifecycle_msgs__srv__ChangeState_Request__destroy((*C.lifecycle_msgs__srv__ChangeState_Request)(pointer_to_free))
}

func (t _ChangeState_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ChangeState_Request)
	mem := (*C.lifecycle_msgs__srv__ChangeState_Request)(dst)
	lifecycle_msgs_msg.TransitionTypeSupport.AsCStruct(unsafe.Pointer(&mem.transition), &m.Transition)
}

func (t _ChangeState_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ChangeState_Request)
	mem := (*C.lifecycle_msgs__srv__ChangeState_Request)(ros2_message_buffer)
	lifecycle_msgs_msg.TransitionTypeSupport.AsGoStruct(&m.Transition, unsafe.Pointer(&mem.transition))
}

func (t _ChangeState_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__lifecycle_msgs__srv__ChangeState_Request())
}

type CChangeState_Request = C.lifecycle_msgs__srv__ChangeState_Request
type CChangeState_Request__Sequence = C.lifecycle_msgs__srv__ChangeState_Request__Sequence

func ChangeState_Request__Sequence_to_Go(goSlice *[]ChangeState_Request, cSlice CChangeState_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ChangeState_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ChangeState_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ChangeState_Request__Sequence_to_C(cSlice *CChangeState_Request__Sequence, goSlice []ChangeState_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.lifecycle_msgs__srv__ChangeState_Request)(C.malloc(C.sizeof_struct_lifecycle_msgs__srv__ChangeState_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ChangeState_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ChangeState_Request__Array_to_Go(goSlice []ChangeState_Request, cSlice []CChangeState_Request) {
	for i := 0; i < len(cSlice); i++ {
		ChangeState_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ChangeState_Request__Array_to_C(cSlice []CChangeState_Request, goSlice []ChangeState_Request) {
	for i := 0; i < len(goSlice); i++ {
		ChangeState_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
