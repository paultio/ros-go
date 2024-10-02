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

#include <geometry_msgs/msg/wrench_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/WrenchStamped", WrenchStampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/WrenchStamped", WrenchStampedTypeSupport)
}

type WrenchStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Wrench Wrench `yaml:"wrench"`
}

// NewWrenchStamped creates a new WrenchStamped with default values.
func NewWrenchStamped() *WrenchStamped {
	self := WrenchStamped{}
	self.SetDefaults()
	return &self
}

func (t *WrenchStamped) Clone() *WrenchStamped {
	c := &WrenchStamped{}
	c.Header = *t.Header.Clone()
	c.Wrench = *t.Wrench.Clone()
	return c
}

func (t *WrenchStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WrenchStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Wrench.SetDefaults()
}

func (t *WrenchStamped) GetTypeSupport() types.MessageTypeSupport {
	return WrenchStampedTypeSupport
}

// WrenchStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WrenchStampedPublisher struct {
	*rclgo.Publisher
}

// NewWrenchStampedPublisher creates and returns a new publisher for the
// WrenchStamped
func NewWrenchStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WrenchStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, WrenchStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WrenchStampedPublisher{pub}, nil
}

func (p *WrenchStampedPublisher) Publish(msg *WrenchStamped) error {
	return p.Publisher.Publish(msg)
}

// WrenchStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WrenchStampedSubscription struct {
	*rclgo.Subscription
}

// WrenchStampedSubscriptionCallback type is used to provide a subscription
// handler function for a WrenchStampedSubscription.
type WrenchStampedSubscriptionCallback func(msg *WrenchStamped, info *rclgo.MessageInfo, err error)

// NewWrenchStampedSubscription creates and returns a new subscription for the
// WrenchStamped
func NewWrenchStampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback WrenchStampedSubscriptionCallback) (*WrenchStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg WrenchStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WrenchStampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &WrenchStampedSubscription{sub}, nil
}

func (s *WrenchStampedSubscription) TakeMessage(out *WrenchStamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWrenchStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWrenchStampedSlice(dst, src []WrenchStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WrenchStampedTypeSupport types.MessageTypeSupport = _WrenchStampedTypeSupport{}

type _WrenchStampedTypeSupport struct{}

func (t _WrenchStampedTypeSupport) New() types.Message {
	return NewWrenchStamped()
}

func (t _WrenchStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__WrenchStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__WrenchStamped__create())
}

func (t _WrenchStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__WrenchStamped__destroy((*C.geometry_msgs__msg__WrenchStamped)(pointer_to_free))
}

func (t _WrenchStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WrenchStamped)
	mem := (*C.geometry_msgs__msg__WrenchStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	WrenchTypeSupport.AsCStruct(unsafe.Pointer(&mem.wrench), &m.Wrench)
}

func (t _WrenchStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WrenchStamped)
	mem := (*C.geometry_msgs__msg__WrenchStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	WrenchTypeSupport.AsGoStruct(&m.Wrench, unsafe.Pointer(&mem.wrench))
}

func (t _WrenchStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__WrenchStamped())
}

type CWrenchStamped = C.geometry_msgs__msg__WrenchStamped
type CWrenchStamped__Sequence = C.geometry_msgs__msg__WrenchStamped__Sequence

func WrenchStamped__Sequence_to_Go(goSlice *[]WrenchStamped, cSlice CWrenchStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WrenchStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WrenchStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func WrenchStamped__Sequence_to_C(cSlice *CWrenchStamped__Sequence, goSlice []WrenchStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__WrenchStamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__WrenchStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WrenchStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func WrenchStamped__Array_to_Go(goSlice []WrenchStamped, cSlice []CWrenchStamped) {
	for i := 0; i < len(cSlice); i++ {
		WrenchStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WrenchStamped__Array_to_C(cSlice []CWrenchStamped, goSlice []WrenchStamped) {
	for i := 0; i < len(goSlice); i++ {
		WrenchStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
