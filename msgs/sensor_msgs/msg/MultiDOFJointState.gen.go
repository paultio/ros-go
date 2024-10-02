// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	geometry_msgs_msg "paultio/ros-go/msgs/geometry_msgs/msg"
	std_msgs_msg "paultio/ros-go/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/multi_dof_joint_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/MultiDOFJointState", MultiDOFJointStateTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/MultiDOFJointState", MultiDOFJointStateTypeSupport)
}

type MultiDOFJointState struct {
	Header std_msgs_msg.Header `yaml:"header"`
	JointNames []string `yaml:"joint_names"`
	Transforms []geometry_msgs_msg.Transform `yaml:"transforms"`
	Twist []geometry_msgs_msg.Twist `yaml:"twist"`
	Wrench []geometry_msgs_msg.Wrench `yaml:"wrench"`
}

// NewMultiDOFJointState creates a new MultiDOFJointState with default values.
func NewMultiDOFJointState() *MultiDOFJointState {
	self := MultiDOFJointState{}
	self.SetDefaults()
	return &self
}

func (t *MultiDOFJointState) Clone() *MultiDOFJointState {
	c := &MultiDOFJointState{}
	c.Header = *t.Header.Clone()
	if t.JointNames != nil {
		c.JointNames = make([]string, len(t.JointNames))
		copy(c.JointNames, t.JointNames)
	}
	if t.Transforms != nil {
		c.Transforms = make([]geometry_msgs_msg.Transform, len(t.Transforms))
		geometry_msgs_msg.CloneTransformSlice(c.Transforms, t.Transforms)
	}
	if t.Twist != nil {
		c.Twist = make([]geometry_msgs_msg.Twist, len(t.Twist))
		geometry_msgs_msg.CloneTwistSlice(c.Twist, t.Twist)
	}
	if t.Wrench != nil {
		c.Wrench = make([]geometry_msgs_msg.Wrench, len(t.Wrench))
		geometry_msgs_msg.CloneWrenchSlice(c.Wrench, t.Wrench)
	}
	return c
}

func (t *MultiDOFJointState) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MultiDOFJointState) SetDefaults() {
	t.Header.SetDefaults()
	t.JointNames = nil
	t.Transforms = nil
	t.Twist = nil
	t.Wrench = nil
}

func (t *MultiDOFJointState) GetTypeSupport() types.MessageTypeSupport {
	return MultiDOFJointStateTypeSupport
}

// MultiDOFJointStatePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MultiDOFJointStatePublisher struct {
	*rclgo.Publisher
}

// NewMultiDOFJointStatePublisher creates and returns a new publisher for the
// MultiDOFJointState
func NewMultiDOFJointStatePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MultiDOFJointStatePublisher, error) {
	pub, err := node.NewPublisher(topic_name, MultiDOFJointStateTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MultiDOFJointStatePublisher{pub}, nil
}

func (p *MultiDOFJointStatePublisher) Publish(msg *MultiDOFJointState) error {
	return p.Publisher.Publish(msg)
}

// MultiDOFJointStateSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MultiDOFJointStateSubscription struct {
	*rclgo.Subscription
}

// MultiDOFJointStateSubscriptionCallback type is used to provide a subscription
// handler function for a MultiDOFJointStateSubscription.
type MultiDOFJointStateSubscriptionCallback func(msg *MultiDOFJointState, info *rclgo.MessageInfo, err error)

// NewMultiDOFJointStateSubscription creates and returns a new subscription for the
// MultiDOFJointState
func NewMultiDOFJointStateSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MultiDOFJointStateSubscriptionCallback) (*MultiDOFJointStateSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MultiDOFJointState
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MultiDOFJointStateTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MultiDOFJointStateSubscription{sub}, nil
}

func (s *MultiDOFJointStateSubscription) TakeMessage(out *MultiDOFJointState) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMultiDOFJointStateSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMultiDOFJointStateSlice(dst, src []MultiDOFJointState) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MultiDOFJointStateTypeSupport types.MessageTypeSupport = _MultiDOFJointStateTypeSupport{}

type _MultiDOFJointStateTypeSupport struct{}

func (t _MultiDOFJointStateTypeSupport) New() types.Message {
	return NewMultiDOFJointState()
}

func (t _MultiDOFJointStateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__MultiDOFJointState
	return (unsafe.Pointer)(C.sensor_msgs__msg__MultiDOFJointState__create())
}

func (t _MultiDOFJointStateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__MultiDOFJointState__destroy((*C.sensor_msgs__msg__MultiDOFJointState)(pointer_to_free))
}

func (t _MultiDOFJointStateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultiDOFJointState)
	mem := (*C.sensor_msgs__msg__MultiDOFJointState)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.String__Sequence_to_C((*primitives.CString__Sequence)(unsafe.Pointer(&mem.joint_names)), m.JointNames)
	geometry_msgs_msg.Transform__Sequence_to_C((*geometry_msgs_msg.CTransform__Sequence)(unsafe.Pointer(&mem.transforms)), m.Transforms)
	geometry_msgs_msg.Twist__Sequence_to_C((*geometry_msgs_msg.CTwist__Sequence)(unsafe.Pointer(&mem.twist)), m.Twist)
	geometry_msgs_msg.Wrench__Sequence_to_C((*geometry_msgs_msg.CWrench__Sequence)(unsafe.Pointer(&mem.wrench)), m.Wrench)
}

func (t _MultiDOFJointStateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultiDOFJointState)
	mem := (*C.sensor_msgs__msg__MultiDOFJointState)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.String__Sequence_to_Go(&m.JointNames, *(*primitives.CString__Sequence)(unsafe.Pointer(&mem.joint_names)))
	geometry_msgs_msg.Transform__Sequence_to_Go(&m.Transforms, *(*geometry_msgs_msg.CTransform__Sequence)(unsafe.Pointer(&mem.transforms)))
	geometry_msgs_msg.Twist__Sequence_to_Go(&m.Twist, *(*geometry_msgs_msg.CTwist__Sequence)(unsafe.Pointer(&mem.twist)))
	geometry_msgs_msg.Wrench__Sequence_to_Go(&m.Wrench, *(*geometry_msgs_msg.CWrench__Sequence)(unsafe.Pointer(&mem.wrench)))
}

func (t _MultiDOFJointStateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__MultiDOFJointState())
}

type CMultiDOFJointState = C.sensor_msgs__msg__MultiDOFJointState
type CMultiDOFJointState__Sequence = C.sensor_msgs__msg__MultiDOFJointState__Sequence

func MultiDOFJointState__Sequence_to_Go(goSlice *[]MultiDOFJointState, cSlice CMultiDOFJointState__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiDOFJointState, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MultiDOFJointStateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MultiDOFJointState__Sequence_to_C(cSlice *CMultiDOFJointState__Sequence, goSlice []MultiDOFJointState) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__MultiDOFJointState)(C.malloc(C.sizeof_struct_sensor_msgs__msg__MultiDOFJointState * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MultiDOFJointStateTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MultiDOFJointState__Array_to_Go(goSlice []MultiDOFJointState, cSlice []CMultiDOFJointState) {
	for i := 0; i < len(cSlice); i++ {
		MultiDOFJointStateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultiDOFJointState__Array_to_C(cSlice []CMultiDOFJointState, goSlice []MultiDOFJointState) {
	for i := 0; i < len(goSlice); i++ {
		MultiDOFJointStateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
