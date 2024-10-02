// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "paultio/ros-go/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/inertia_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/InertiaStamped", InertiaStampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/InertiaStamped", InertiaStampedTypeSupport)
}

type InertiaStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Inertia Inertia `yaml:"inertia"`
}

// NewInertiaStamped creates a new InertiaStamped with default values.
func NewInertiaStamped() *InertiaStamped {
	self := InertiaStamped{}
	self.SetDefaults()
	return &self
}

func (t *InertiaStamped) Clone() *InertiaStamped {
	c := &InertiaStamped{}
	c.Header = *t.Header.Clone()
	c.Inertia = *t.Inertia.Clone()
	return c
}

func (t *InertiaStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InertiaStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Inertia.SetDefaults()
}

func (t *InertiaStamped) GetTypeSupport() types.MessageTypeSupport {
	return InertiaStampedTypeSupport
}

// InertiaStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type InertiaStampedPublisher struct {
	*rclgo.Publisher
}

// NewInertiaStampedPublisher creates and returns a new publisher for the
// InertiaStamped
func NewInertiaStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*InertiaStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, InertiaStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &InertiaStampedPublisher{pub}, nil
}

func (p *InertiaStampedPublisher) Publish(msg *InertiaStamped) error {
	return p.Publisher.Publish(msg)
}

// InertiaStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type InertiaStampedSubscription struct {
	*rclgo.Subscription
}

// InertiaStampedSubscriptionCallback type is used to provide a subscription
// handler function for a InertiaStampedSubscription.
type InertiaStampedSubscriptionCallback func(msg *InertiaStamped, info *rclgo.MessageInfo, err error)

// NewInertiaStampedSubscription creates and returns a new subscription for the
// InertiaStamped
func NewInertiaStampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback InertiaStampedSubscriptionCallback) (*InertiaStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg InertiaStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, InertiaStampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &InertiaStampedSubscription{sub}, nil
}

func (s *InertiaStampedSubscription) TakeMessage(out *InertiaStamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneInertiaStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInertiaStampedSlice(dst, src []InertiaStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InertiaStampedTypeSupport types.MessageTypeSupport = _InertiaStampedTypeSupport{}

type _InertiaStampedTypeSupport struct{}

func (t _InertiaStampedTypeSupport) New() types.Message {
	return NewInertiaStamped()
}

func (t _InertiaStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__InertiaStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__InertiaStamped__create())
}

func (t _InertiaStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__InertiaStamped__destroy((*C.geometry_msgs__msg__InertiaStamped)(pointer_to_free))
}

func (t _InertiaStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InertiaStamped)
	mem := (*C.geometry_msgs__msg__InertiaStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	InertiaTypeSupport.AsCStruct(unsafe.Pointer(&mem.inertia), &m.Inertia)
}

func (t _InertiaStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InertiaStamped)
	mem := (*C.geometry_msgs__msg__InertiaStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	InertiaTypeSupport.AsGoStruct(&m.Inertia, unsafe.Pointer(&mem.inertia))
}

func (t _InertiaStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__InertiaStamped())
}

type CInertiaStamped = C.geometry_msgs__msg__InertiaStamped
type CInertiaStamped__Sequence = C.geometry_msgs__msg__InertiaStamped__Sequence

func InertiaStamped__Sequence_to_Go(goSlice *[]InertiaStamped, cSlice CInertiaStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InertiaStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		InertiaStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func InertiaStamped__Sequence_to_C(cSlice *CInertiaStamped__Sequence, goSlice []InertiaStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__InertiaStamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__InertiaStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		InertiaStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func InertiaStamped__Array_to_Go(goSlice []InertiaStamped, cSlice []CInertiaStamped) {
	for i := 0; i < len(cSlice); i++ {
		InertiaStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InertiaStamped__Array_to_C(cSlice []CInertiaStamped, goSlice []InertiaStamped) {
	for i := 0; i < len(goSlice); i++ {
		InertiaStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
