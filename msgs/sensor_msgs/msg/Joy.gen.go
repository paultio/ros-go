// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	std_msgs_msg "paultio/ros-go/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/joy.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Joy", JoyTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/Joy", JoyTypeSupport)
}

type Joy struct {
	Header std_msgs_msg.Header `yaml:"header"`// The timestamp is the time at which data is received from the joystick.
	Axes []float32 `yaml:"axes"`// The axes measurements from a joystick.
	Buttons []int32 `yaml:"buttons"`// The buttons measurements from a joystick.
}

// NewJoy creates a new Joy with default values.
func NewJoy() *Joy {
	self := Joy{}
	self.SetDefaults()
	return &self
}

func (t *Joy) Clone() *Joy {
	c := &Joy{}
	c.Header = *t.Header.Clone()
	if t.Axes != nil {
		c.Axes = make([]float32, len(t.Axes))
		copy(c.Axes, t.Axes)
	}
	if t.Buttons != nil {
		c.Buttons = make([]int32, len(t.Buttons))
		copy(c.Buttons, t.Buttons)
	}
	return c
}

func (t *Joy) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Joy) SetDefaults() {
	t.Header.SetDefaults()
	t.Axes = nil
	t.Buttons = nil
}

func (t *Joy) GetTypeSupport() types.MessageTypeSupport {
	return JoyTypeSupport
}

// JoyPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type JoyPublisher struct {
	*rclgo.Publisher
}

// NewJoyPublisher creates and returns a new publisher for the
// Joy
func NewJoyPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*JoyPublisher, error) {
	pub, err := node.NewPublisher(topic_name, JoyTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &JoyPublisher{pub}, nil
}

func (p *JoyPublisher) Publish(msg *Joy) error {
	return p.Publisher.Publish(msg)
}

// JoySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type JoySubscription struct {
	*rclgo.Subscription
}

// JoySubscriptionCallback type is used to provide a subscription
// handler function for a JoySubscription.
type JoySubscriptionCallback func(msg *Joy, info *rclgo.MessageInfo, err error)

// NewJoySubscription creates and returns a new subscription for the
// Joy
func NewJoySubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback JoySubscriptionCallback) (*JoySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Joy
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, JoyTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &JoySubscription{sub}, nil
}

func (s *JoySubscription) TakeMessage(out *Joy) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneJoySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJoySlice(dst, src []Joy) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JoyTypeSupport types.MessageTypeSupport = _JoyTypeSupport{}

type _JoyTypeSupport struct{}

func (t _JoyTypeSupport) New() types.Message {
	return NewJoy()
}

func (t _JoyTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Joy
	return (unsafe.Pointer)(C.sensor_msgs__msg__Joy__create())
}

func (t _JoyTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Joy__destroy((*C.sensor_msgs__msg__Joy)(pointer_to_free))
}

func (t _JoyTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Joy)
	mem := (*C.sensor_msgs__msg__Joy)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.axes)), m.Axes)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.buttons)), m.Buttons)
}

func (t _JoyTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Joy)
	mem := (*C.sensor_msgs__msg__Joy)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.Float32__Sequence_to_Go(&m.Axes, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.axes)))
	primitives.Int32__Sequence_to_Go(&m.Buttons, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.buttons)))
}

func (t _JoyTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Joy())
}

type CJoy = C.sensor_msgs__msg__Joy
type CJoy__Sequence = C.sensor_msgs__msg__Joy__Sequence

func Joy__Sequence_to_Go(goSlice *[]Joy, cSlice CJoy__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Joy, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		JoyTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Joy__Sequence_to_C(cSlice *CJoy__Sequence, goSlice []Joy) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Joy)(C.malloc(C.sizeof_struct_sensor_msgs__msg__Joy * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		JoyTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Joy__Array_to_Go(goSlice []Joy, cSlice []CJoy) {
	for i := 0; i < len(cSlice); i++ {
		JoyTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Joy__Array_to_C(cSlice []CJoy, goSlice []Joy) {
	for i := 0; i < len(goSlice); i++ {
		JoyTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
