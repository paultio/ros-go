// Code generated by rclgo-gen. DO NOT EDIT.

package rosgraph_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "paultio/ros-go/msgs/builtin_interfaces/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rosgraph_msgs/msg/clock.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rosgraph_msgs/Clock", ClockTypeSupport)
	typemap.RegisterMessage("rosgraph_msgs/msg/Clock", ClockTypeSupport)
}

type Clock struct {
	Clock builtin_interfaces_msg.Time `yaml:"clock"`// This message communicates the current time.For more information, see https://design.ros2.org/articles/clock_and_time.html.
}

// NewClock creates a new Clock with default values.
func NewClock() *Clock {
	self := Clock{}
	self.SetDefaults()
	return &self
}

func (t *Clock) Clone() *Clock {
	c := &Clock{}
	c.Clock = *t.Clock.Clone()
	return c
}

func (t *Clock) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Clock) SetDefaults() {
	t.Clock.SetDefaults()
}

func (t *Clock) GetTypeSupport() types.MessageTypeSupport {
	return ClockTypeSupport
}

// ClockPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ClockPublisher struct {
	*rclgo.Publisher
}

// NewClockPublisher creates and returns a new publisher for the
// Clock
func NewClockPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ClockPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ClockTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ClockPublisher{pub}, nil
}

func (p *ClockPublisher) Publish(msg *Clock) error {
	return p.Publisher.Publish(msg)
}

// ClockSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ClockSubscription struct {
	*rclgo.Subscription
}

// ClockSubscriptionCallback type is used to provide a subscription
// handler function for a ClockSubscription.
type ClockSubscriptionCallback func(msg *Clock, info *rclgo.MessageInfo, err error)

// NewClockSubscription creates and returns a new subscription for the
// Clock
func NewClockSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ClockSubscriptionCallback) (*ClockSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Clock
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ClockTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ClockSubscription{sub}, nil
}

func (s *ClockSubscription) TakeMessage(out *Clock) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneClockSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneClockSlice(dst, src []Clock) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ClockTypeSupport types.MessageTypeSupport = _ClockTypeSupport{}

type _ClockTypeSupport struct{}

func (t _ClockTypeSupport) New() types.Message {
	return NewClock()
}

func (t _ClockTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rosgraph_msgs__msg__Clock
	return (unsafe.Pointer)(C.rosgraph_msgs__msg__Clock__create())
}

func (t _ClockTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rosgraph_msgs__msg__Clock__destroy((*C.rosgraph_msgs__msg__Clock)(pointer_to_free))
}

func (t _ClockTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Clock)
	mem := (*C.rosgraph_msgs__msg__Clock)(dst)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.clock), &m.Clock)
}

func (t _ClockTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Clock)
	mem := (*C.rosgraph_msgs__msg__Clock)(ros2_message_buffer)
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.Clock, unsafe.Pointer(&mem.clock))
}

func (t _ClockTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rosgraph_msgs__msg__Clock())
}

type CClock = C.rosgraph_msgs__msg__Clock
type CClock__Sequence = C.rosgraph_msgs__msg__Clock__Sequence

func Clock__Sequence_to_Go(goSlice *[]Clock, cSlice CClock__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Clock, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ClockTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Clock__Sequence_to_C(cSlice *CClock__Sequence, goSlice []Clock) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rosgraph_msgs__msg__Clock)(C.malloc(C.sizeof_struct_rosgraph_msgs__msg__Clock * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ClockTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Clock__Array_to_Go(goSlice []Clock, cSlice []CClock) {
	for i := 0; i < len(cSlice); i++ {
		ClockTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Clock__Array_to_C(cSlice []CClock, goSlice []Clock) {
	for i := 0; i < len(goSlice); i++ {
		ClockTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
