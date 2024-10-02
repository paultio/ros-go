// Code generated by rclgo-gen. DO NOT EDIT.

package diagnostic_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "paultio/ros-go/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <diagnostic_msgs/msg/diagnostic_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("diagnostic_msgs/DiagnosticArray", DiagnosticArrayTypeSupport)
	typemap.RegisterMessage("diagnostic_msgs/msg/DiagnosticArray", DiagnosticArrayTypeSupport)
}

type DiagnosticArray struct {
	Header std_msgs_msg.Header `yaml:"header"`// for timestamp. This message is used to send diagnostic information about the state of the robot.
	Status []DiagnosticStatus `yaml:"status"`// an array of components being reported on
}

// NewDiagnosticArray creates a new DiagnosticArray with default values.
func NewDiagnosticArray() *DiagnosticArray {
	self := DiagnosticArray{}
	self.SetDefaults()
	return &self
}

func (t *DiagnosticArray) Clone() *DiagnosticArray {
	c := &DiagnosticArray{}
	c.Header = *t.Header.Clone()
	if t.Status != nil {
		c.Status = make([]DiagnosticStatus, len(t.Status))
		CloneDiagnosticStatusSlice(c.Status, t.Status)
	}
	return c
}

func (t *DiagnosticArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *DiagnosticArray) SetDefaults() {
	t.Header.SetDefaults()
	t.Status = nil
}

func (t *DiagnosticArray) GetTypeSupport() types.MessageTypeSupport {
	return DiagnosticArrayTypeSupport
}

// DiagnosticArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type DiagnosticArrayPublisher struct {
	*rclgo.Publisher
}

// NewDiagnosticArrayPublisher creates and returns a new publisher for the
// DiagnosticArray
func NewDiagnosticArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*DiagnosticArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, DiagnosticArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &DiagnosticArrayPublisher{pub}, nil
}

func (p *DiagnosticArrayPublisher) Publish(msg *DiagnosticArray) error {
	return p.Publisher.Publish(msg)
}

// DiagnosticArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type DiagnosticArraySubscription struct {
	*rclgo.Subscription
}

// DiagnosticArraySubscriptionCallback type is used to provide a subscription
// handler function for a DiagnosticArraySubscription.
type DiagnosticArraySubscriptionCallback func(msg *DiagnosticArray, info *rclgo.MessageInfo, err error)

// NewDiagnosticArraySubscription creates and returns a new subscription for the
// DiagnosticArray
func NewDiagnosticArraySubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback DiagnosticArraySubscriptionCallback) (*DiagnosticArraySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg DiagnosticArray
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, DiagnosticArrayTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &DiagnosticArraySubscription{sub}, nil
}

func (s *DiagnosticArraySubscription) TakeMessage(out *DiagnosticArray) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneDiagnosticArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDiagnosticArraySlice(dst, src []DiagnosticArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DiagnosticArrayTypeSupport types.MessageTypeSupport = _DiagnosticArrayTypeSupport{}

type _DiagnosticArrayTypeSupport struct{}

func (t _DiagnosticArrayTypeSupport) New() types.Message {
	return NewDiagnosticArray()
}

func (t _DiagnosticArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.diagnostic_msgs__msg__DiagnosticArray
	return (unsafe.Pointer)(C.diagnostic_msgs__msg__DiagnosticArray__create())
}

func (t _DiagnosticArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.diagnostic_msgs__msg__DiagnosticArray__destroy((*C.diagnostic_msgs__msg__DiagnosticArray)(pointer_to_free))
}

func (t _DiagnosticArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*DiagnosticArray)
	mem := (*C.diagnostic_msgs__msg__DiagnosticArray)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	DiagnosticStatus__Sequence_to_C(&mem.status, m.Status)
}

func (t _DiagnosticArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*DiagnosticArray)
	mem := (*C.diagnostic_msgs__msg__DiagnosticArray)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	DiagnosticStatus__Sequence_to_Go(&m.Status, mem.status)
}

func (t _DiagnosticArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__diagnostic_msgs__msg__DiagnosticArray())
}

type CDiagnosticArray = C.diagnostic_msgs__msg__DiagnosticArray
type CDiagnosticArray__Sequence = C.diagnostic_msgs__msg__DiagnosticArray__Sequence

func DiagnosticArray__Sequence_to_Go(goSlice *[]DiagnosticArray, cSlice CDiagnosticArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DiagnosticArray, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		DiagnosticArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func DiagnosticArray__Sequence_to_C(cSlice *CDiagnosticArray__Sequence, goSlice []DiagnosticArray) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.diagnostic_msgs__msg__DiagnosticArray)(C.malloc(C.sizeof_struct_diagnostic_msgs__msg__DiagnosticArray * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		DiagnosticArrayTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func DiagnosticArray__Array_to_Go(goSlice []DiagnosticArray, cSlice []CDiagnosticArray) {
	for i := 0; i < len(cSlice); i++ {
		DiagnosticArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func DiagnosticArray__Array_to_C(cSlice []CDiagnosticArray, goSlice []DiagnosticArray) {
	for i := 0; i < len(goSlice); i++ {
		DiagnosticArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
