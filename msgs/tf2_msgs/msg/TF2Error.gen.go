// Code generated by rclgo-gen. DO NOT EDIT.

package tf2_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <tf2_msgs/msg/tf2_error.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("tf2_msgs/TF2Error", TF2ErrorTypeSupport)
	typemap.RegisterMessage("tf2_msgs/msg/TF2Error", TF2ErrorTypeSupport)
}
const (
	TF2Error_NO_ERROR uint8 = 0
	TF2Error_LOOKUP_ERROR uint8 = 1
	TF2Error_CONNECTIVITY_ERROR uint8 = 2
	TF2Error_EXTRAPOLATION_ERROR uint8 = 3
	TF2Error_INVALID_ARGUMENT_ERROR uint8 = 4
	TF2Error_TIMEOUT_ERROR uint8 = 5
	TF2Error_TRANSFORM_ERROR uint8 = 6
)

type TF2Error struct {
	Error uint8 `yaml:"error"`
	ErrorString string `yaml:"error_string"`
}

// NewTF2Error creates a new TF2Error with default values.
func NewTF2Error() *TF2Error {
	self := TF2Error{}
	self.SetDefaults()
	return &self
}

func (t *TF2Error) Clone() *TF2Error {
	c := &TF2Error{}
	c.Error = t.Error
	c.ErrorString = t.ErrorString
	return c
}

func (t *TF2Error) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TF2Error) SetDefaults() {
	t.Error = 0
	t.ErrorString = ""
}

func (t *TF2Error) GetTypeSupport() types.MessageTypeSupport {
	return TF2ErrorTypeSupport
}

// TF2ErrorPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TF2ErrorPublisher struct {
	*rclgo.Publisher
}

// NewTF2ErrorPublisher creates and returns a new publisher for the
// TF2Error
func NewTF2ErrorPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TF2ErrorPublisher, error) {
	pub, err := node.NewPublisher(topic_name, TF2ErrorTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TF2ErrorPublisher{pub}, nil
}

func (p *TF2ErrorPublisher) Publish(msg *TF2Error) error {
	return p.Publisher.Publish(msg)
}

// TF2ErrorSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TF2ErrorSubscription struct {
	*rclgo.Subscription
}

// TF2ErrorSubscriptionCallback type is used to provide a subscription
// handler function for a TF2ErrorSubscription.
type TF2ErrorSubscriptionCallback func(msg *TF2Error, info *rclgo.MessageInfo, err error)

// NewTF2ErrorSubscription creates and returns a new subscription for the
// TF2Error
func NewTF2ErrorSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback TF2ErrorSubscriptionCallback) (*TF2ErrorSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg TF2Error
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TF2ErrorTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &TF2ErrorSubscription{sub}, nil
}

func (s *TF2ErrorSubscription) TakeMessage(out *TF2Error) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTF2ErrorSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTF2ErrorSlice(dst, src []TF2Error) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TF2ErrorTypeSupport types.MessageTypeSupport = _TF2ErrorTypeSupport{}

type _TF2ErrorTypeSupport struct{}

func (t _TF2ErrorTypeSupport) New() types.Message {
	return NewTF2Error()
}

func (t _TF2ErrorTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.tf2_msgs__msg__TF2Error
	return (unsafe.Pointer)(C.tf2_msgs__msg__TF2Error__create())
}

func (t _TF2ErrorTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.tf2_msgs__msg__TF2Error__destroy((*C.tf2_msgs__msg__TF2Error)(pointer_to_free))
}

func (t _TF2ErrorTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TF2Error)
	mem := (*C.tf2_msgs__msg__TF2Error)(dst)
	mem.error = C.uint8_t(m.Error)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.error_string), m.ErrorString)
}

func (t _TF2ErrorTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TF2Error)
	mem := (*C.tf2_msgs__msg__TF2Error)(ros2_message_buffer)
	m.Error = uint8(mem.error)
	primitives.StringAsGoStruct(&m.ErrorString, unsafe.Pointer(&mem.error_string))
}

func (t _TF2ErrorTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__tf2_msgs__msg__TF2Error())
}

type CTF2Error = C.tf2_msgs__msg__TF2Error
type CTF2Error__Sequence = C.tf2_msgs__msg__TF2Error__Sequence

func TF2Error__Sequence_to_Go(goSlice *[]TF2Error, cSlice CTF2Error__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TF2Error, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		TF2ErrorTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func TF2Error__Sequence_to_C(cSlice *CTF2Error__Sequence, goSlice []TF2Error) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.tf2_msgs__msg__TF2Error)(C.malloc(C.sizeof_struct_tf2_msgs__msg__TF2Error * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		TF2ErrorTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func TF2Error__Array_to_Go(goSlice []TF2Error, cSlice []CTF2Error) {
	for i := 0; i < len(cSlice); i++ {
		TF2ErrorTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TF2Error__Array_to_C(cSlice []CTF2Error, goSlice []TF2Error) {
	for i := 0; i < len(goSlice); i++ {
		TF2ErrorTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
