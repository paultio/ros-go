// Code generated by rclgo-gen. DO NOT EDIT.

package rcl_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/msg/parameter_type.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/ParameterType", ParameterTypeTypeSupport)
	typemap.RegisterMessage("rcl_interfaces/msg/ParameterType", ParameterTypeTypeSupport)
}
const (
	ParameterType_PARAMETER_NOT_SET uint8 = 0// Default value, which implies this is not a valid parameter.
	ParameterType_PARAMETER_BOOL uint8 = 1
	ParameterType_PARAMETER_INTEGER uint8 = 2
	ParameterType_PARAMETER_DOUBLE uint8 = 3
	ParameterType_PARAMETER_STRING uint8 = 4
	ParameterType_PARAMETER_BYTE_ARRAY uint8 = 5
	ParameterType_PARAMETER_BOOL_ARRAY uint8 = 6
	ParameterType_PARAMETER_INTEGER_ARRAY uint8 = 7
	ParameterType_PARAMETER_DOUBLE_ARRAY uint8 = 8
	ParameterType_PARAMETER_STRING_ARRAY uint8 = 9
)

type ParameterType struct {
}

// NewParameterType creates a new ParameterType with default values.
func NewParameterType() *ParameterType {
	self := ParameterType{}
	self.SetDefaults()
	return &self
}

func (t *ParameterType) Clone() *ParameterType {
	c := &ParameterType{}
	return c
}

func (t *ParameterType) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ParameterType) SetDefaults() {
}

func (t *ParameterType) GetTypeSupport() types.MessageTypeSupport {
	return ParameterTypeTypeSupport
}

// ParameterTypePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ParameterTypePublisher struct {
	*rclgo.Publisher
}

// NewParameterTypePublisher creates and returns a new publisher for the
// ParameterType
func NewParameterTypePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ParameterTypePublisher, error) {
	pub, err := node.NewPublisher(topic_name, ParameterTypeTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ParameterTypePublisher{pub}, nil
}

func (p *ParameterTypePublisher) Publish(msg *ParameterType) error {
	return p.Publisher.Publish(msg)
}

// ParameterTypeSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ParameterTypeSubscription struct {
	*rclgo.Subscription
}

// ParameterTypeSubscriptionCallback type is used to provide a subscription
// handler function for a ParameterTypeSubscription.
type ParameterTypeSubscriptionCallback func(msg *ParameterType, info *rclgo.MessageInfo, err error)

// NewParameterTypeSubscription creates and returns a new subscription for the
// ParameterType
func NewParameterTypeSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ParameterTypeSubscriptionCallback) (*ParameterTypeSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ParameterType
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ParameterTypeTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ParameterTypeSubscription{sub}, nil
}

func (s *ParameterTypeSubscription) TakeMessage(out *ParameterType) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneParameterTypeSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneParameterTypeSlice(dst, src []ParameterType) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ParameterTypeTypeSupport types.MessageTypeSupport = _ParameterTypeTypeSupport{}

type _ParameterTypeTypeSupport struct{}

func (t _ParameterTypeTypeSupport) New() types.Message {
	return NewParameterType()
}

func (t _ParameterTypeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__msg__ParameterType
	return (unsafe.Pointer)(C.rcl_interfaces__msg__ParameterType__create())
}

func (t _ParameterTypeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__msg__ParameterType__destroy((*C.rcl_interfaces__msg__ParameterType)(pointer_to_free))
}

func (t _ParameterTypeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _ParameterTypeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _ParameterTypeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__msg__ParameterType())
}

type CParameterType = C.rcl_interfaces__msg__ParameterType
type CParameterType__Sequence = C.rcl_interfaces__msg__ParameterType__Sequence

func ParameterType__Sequence_to_Go(goSlice *[]ParameterType, cSlice CParameterType__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ParameterType, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ParameterTypeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ParameterType__Sequence_to_C(cSlice *CParameterType__Sequence, goSlice []ParameterType) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rcl_interfaces__msg__ParameterType)(C.malloc(C.sizeof_struct_rcl_interfaces__msg__ParameterType * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ParameterTypeTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ParameterType__Array_to_Go(goSlice []ParameterType, cSlice []CParameterType) {
	for i := 0; i < len(cSlice); i++ {
		ParameterTypeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ParameterType__Array_to_C(cSlice []CParameterType, goSlice []ParameterType) {
	for i := 0; i < len(goSlice); i++ {
		ParameterTypeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
