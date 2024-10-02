// Code generated by rclgo-gen. DO NOT EDIT.

package rcl_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	builtin_interfaces_msg "paultio/ros-go/msgs/builtin_interfaces/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/msg/log.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/Log", LogTypeSupport)
	typemap.RegisterMessage("rcl_interfaces/msg/Log", LogTypeSupport)
}
const (
	Log_DEBUG byte = 10// Debug is for pedantic information, which is useful when debugging issues.
	Log_INFO byte = 20// Info is the standard informational level and is used to report expectedinformation.
	Log_WARN byte = 30// Warning is for information that may potentially cause issues or possibly unexpectedbehavior.
	Log_ERROR byte = 40// Error is for information that this node cannot resolve.
	Log_FATAL byte = 50// Information about a impending node shutdown.
)

type Log struct {
	Stamp builtin_interfaces_msg.Time `yaml:"stamp"`// Timestamp when this message was generated by the node.
	Level uint8 `yaml:"level"`// Corresponding log level, see above definitions.
	Name string `yaml:"name"`// The name representing the logger this message came from.
	Msg string `yaml:"msg"`// The full log message.
	File string `yaml:"file"`// The file the message came from.
	Function string `yaml:"function"`// The function the message came from.
	Line uint32 `yaml:"line"`// The line in the file the message came from.
}

// NewLog creates a new Log with default values.
func NewLog() *Log {
	self := Log{}
	self.SetDefaults()
	return &self
}

func (t *Log) Clone() *Log {
	c := &Log{}
	c.Stamp = *t.Stamp.Clone()
	c.Level = t.Level
	c.Name = t.Name
	c.Msg = t.Msg
	c.File = t.File
	c.Function = t.Function
	c.Line = t.Line
	return c
}

func (t *Log) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Log) SetDefaults() {
	t.Stamp.SetDefaults()
	t.Level = 0
	t.Name = ""
	t.Msg = ""
	t.File = ""
	t.Function = ""
	t.Line = 0
}

func (t *Log) GetTypeSupport() types.MessageTypeSupport {
	return LogTypeSupport
}

// LogPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type LogPublisher struct {
	*rclgo.Publisher
}

// NewLogPublisher creates and returns a new publisher for the
// Log
func NewLogPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*LogPublisher, error) {
	pub, err := node.NewPublisher(topic_name, LogTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &LogPublisher{pub}, nil
}

func (p *LogPublisher) Publish(msg *Log) error {
	return p.Publisher.Publish(msg)
}

// LogSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type LogSubscription struct {
	*rclgo.Subscription
}

// LogSubscriptionCallback type is used to provide a subscription
// handler function for a LogSubscription.
type LogSubscriptionCallback func(msg *Log, info *rclgo.MessageInfo, err error)

// NewLogSubscription creates and returns a new subscription for the
// Log
func NewLogSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback LogSubscriptionCallback) (*LogSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Log
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, LogTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &LogSubscription{sub}, nil
}

func (s *LogSubscription) TakeMessage(out *Log) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneLogSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneLogSlice(dst, src []Log) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var LogTypeSupport types.MessageTypeSupport = _LogTypeSupport{}

type _LogTypeSupport struct{}

func (t _LogTypeSupport) New() types.Message {
	return NewLog()
}

func (t _LogTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__msg__Log
	return (unsafe.Pointer)(C.rcl_interfaces__msg__Log__create())
}

func (t _LogTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__msg__Log__destroy((*C.rcl_interfaces__msg__Log)(pointer_to_free))
}

func (t _LogTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Log)
	mem := (*C.rcl_interfaces__msg__Log)(dst)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.stamp), &m.Stamp)
	mem.level = C.uint8_t(m.Level)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.msg), m.Msg)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.file), m.File)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.function), m.Function)
	mem.line = C.uint32_t(m.Line)
}

func (t _LogTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Log)
	mem := (*C.rcl_interfaces__msg__Log)(ros2_message_buffer)
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.Stamp, unsafe.Pointer(&mem.stamp))
	m.Level = uint8(mem.level)
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
	primitives.StringAsGoStruct(&m.Msg, unsafe.Pointer(&mem.msg))
	primitives.StringAsGoStruct(&m.File, unsafe.Pointer(&mem.file))
	primitives.StringAsGoStruct(&m.Function, unsafe.Pointer(&mem.function))
	m.Line = uint32(mem.line)
}

func (t _LogTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__msg__Log())
}

type CLog = C.rcl_interfaces__msg__Log
type CLog__Sequence = C.rcl_interfaces__msg__Log__Sequence

func Log__Sequence_to_Go(goSlice *[]Log, cSlice CLog__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Log, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		LogTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Log__Sequence_to_C(cSlice *CLog__Sequence, goSlice []Log) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rcl_interfaces__msg__Log)(C.malloc(C.sizeof_struct_rcl_interfaces__msg__Log * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		LogTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Log__Array_to_Go(goSlice []Log, cSlice []CLog) {
	for i := 0; i < len(cSlice); i++ {
		LogTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Log__Array_to_C(cSlice []CLog, goSlice []Log) {
	for i := 0; i < len(goSlice); i++ {
		LogTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
