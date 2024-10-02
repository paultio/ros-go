// Code generated by rclgo-gen. DO NOT EDIT.

package diagnostic_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <diagnostic_msgs/srv/add_diagnostics.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("diagnostic_msgs/AddDiagnostics_Response", AddDiagnostics_ResponseTypeSupport)
	typemap.RegisterMessage("diagnostic_msgs/srv/AddDiagnostics_Response", AddDiagnostics_ResponseTypeSupport)
}

type AddDiagnostics_Response struct {
	Success bool `yaml:"success"`// True if diagnostic aggregator was updated with new diagnostics, Falseotherwise. A false return value means that either there is a bond in theaggregator which already used the requested namespace, or the initializationof analyzers failed.
	Message string `yaml:"message"`// Message with additional information about the success or failure
}

// NewAddDiagnostics_Response creates a new AddDiagnostics_Response with default values.
func NewAddDiagnostics_Response() *AddDiagnostics_Response {
	self := AddDiagnostics_Response{}
	self.SetDefaults()
	return &self
}

func (t *AddDiagnostics_Response) Clone() *AddDiagnostics_Response {
	c := &AddDiagnostics_Response{}
	c.Success = t.Success
	c.Message = t.Message
	return c
}

func (t *AddDiagnostics_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *AddDiagnostics_Response) SetDefaults() {
	t.Success = false
	t.Message = ""
}

func (t *AddDiagnostics_Response) GetTypeSupport() types.MessageTypeSupport {
	return AddDiagnostics_ResponseTypeSupport
}

// AddDiagnostics_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type AddDiagnostics_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewAddDiagnostics_ResponsePublisher creates and returns a new publisher for the
// AddDiagnostics_Response
func NewAddDiagnostics_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*AddDiagnostics_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, AddDiagnostics_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &AddDiagnostics_ResponsePublisher{pub}, nil
}

func (p *AddDiagnostics_ResponsePublisher) Publish(msg *AddDiagnostics_Response) error {
	return p.Publisher.Publish(msg)
}

// AddDiagnostics_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type AddDiagnostics_ResponseSubscription struct {
	*rclgo.Subscription
}

// AddDiagnostics_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a AddDiagnostics_ResponseSubscription.
type AddDiagnostics_ResponseSubscriptionCallback func(msg *AddDiagnostics_Response, info *rclgo.MessageInfo, err error)

// NewAddDiagnostics_ResponseSubscription creates and returns a new subscription for the
// AddDiagnostics_Response
func NewAddDiagnostics_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback AddDiagnostics_ResponseSubscriptionCallback) (*AddDiagnostics_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg AddDiagnostics_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, AddDiagnostics_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &AddDiagnostics_ResponseSubscription{sub}, nil
}

func (s *AddDiagnostics_ResponseSubscription) TakeMessage(out *AddDiagnostics_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneAddDiagnostics_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAddDiagnostics_ResponseSlice(dst, src []AddDiagnostics_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AddDiagnostics_ResponseTypeSupport types.MessageTypeSupport = _AddDiagnostics_ResponseTypeSupport{}

type _AddDiagnostics_ResponseTypeSupport struct{}

func (t _AddDiagnostics_ResponseTypeSupport) New() types.Message {
	return NewAddDiagnostics_Response()
}

func (t _AddDiagnostics_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.diagnostic_msgs__srv__AddDiagnostics_Response
	return (unsafe.Pointer)(C.diagnostic_msgs__srv__AddDiagnostics_Response__create())
}

func (t _AddDiagnostics_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.diagnostic_msgs__srv__AddDiagnostics_Response__destroy((*C.diagnostic_msgs__srv__AddDiagnostics_Response)(pointer_to_free))
}

func (t _AddDiagnostics_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AddDiagnostics_Response)
	mem := (*C.diagnostic_msgs__srv__AddDiagnostics_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.message), m.Message)
}

func (t _AddDiagnostics_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AddDiagnostics_Response)
	mem := (*C.diagnostic_msgs__srv__AddDiagnostics_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.Message, unsafe.Pointer(&mem.message))
}

func (t _AddDiagnostics_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__diagnostic_msgs__srv__AddDiagnostics_Response())
}

type CAddDiagnostics_Response = C.diagnostic_msgs__srv__AddDiagnostics_Response
type CAddDiagnostics_Response__Sequence = C.diagnostic_msgs__srv__AddDiagnostics_Response__Sequence

func AddDiagnostics_Response__Sequence_to_Go(goSlice *[]AddDiagnostics_Response, cSlice CAddDiagnostics_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AddDiagnostics_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		AddDiagnostics_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func AddDiagnostics_Response__Sequence_to_C(cSlice *CAddDiagnostics_Response__Sequence, goSlice []AddDiagnostics_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.diagnostic_msgs__srv__AddDiagnostics_Response)(C.malloc(C.sizeof_struct_diagnostic_msgs__srv__AddDiagnostics_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		AddDiagnostics_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func AddDiagnostics_Response__Array_to_Go(goSlice []AddDiagnostics_Response, cSlice []CAddDiagnostics_Response) {
	for i := 0; i < len(cSlice); i++ {
		AddDiagnostics_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AddDiagnostics_Response__Array_to_C(cSlice []CAddDiagnostics_Response, goSlice []AddDiagnostics_Response) {
	for i := 0; i < len(goSlice); i++ {
		AddDiagnostics_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
