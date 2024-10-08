// Code generated by rclgo-gen. DO NOT EDIT.

package shape_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	geometry_msgs_msg "paultio/ros-go/msgs/geometry_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <shape_msgs/msg/mesh.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("shape_msgs/Mesh", MeshTypeSupport)
	typemap.RegisterMessage("shape_msgs/msg/Mesh", MeshTypeSupport)
}

type Mesh struct {
	Triangles []MeshTriangle `yaml:"triangles"`// List of triangles; the index values refer to positions in vertices[].
	Vertices []geometry_msgs_msg.Point `yaml:"vertices"`// The actual vertices that make up the mesh.
}

// NewMesh creates a new Mesh with default values.
func NewMesh() *Mesh {
	self := Mesh{}
	self.SetDefaults()
	return &self
}

func (t *Mesh) Clone() *Mesh {
	c := &Mesh{}
	if t.Triangles != nil {
		c.Triangles = make([]MeshTriangle, len(t.Triangles))
		CloneMeshTriangleSlice(c.Triangles, t.Triangles)
	}
	if t.Vertices != nil {
		c.Vertices = make([]geometry_msgs_msg.Point, len(t.Vertices))
		geometry_msgs_msg.ClonePointSlice(c.Vertices, t.Vertices)
	}
	return c
}

func (t *Mesh) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Mesh) SetDefaults() {
	t.Triangles = nil
	t.Vertices = nil
}

func (t *Mesh) GetTypeSupport() types.MessageTypeSupport {
	return MeshTypeSupport
}

// MeshPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MeshPublisher struct {
	*rclgo.Publisher
}

// NewMeshPublisher creates and returns a new publisher for the
// Mesh
func NewMeshPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MeshPublisher, error) {
	pub, err := node.NewPublisher(topic_name, MeshTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MeshPublisher{pub}, nil
}

func (p *MeshPublisher) Publish(msg *Mesh) error {
	return p.Publisher.Publish(msg)
}

// MeshSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MeshSubscription struct {
	*rclgo.Subscription
}

// MeshSubscriptionCallback type is used to provide a subscription
// handler function for a MeshSubscription.
type MeshSubscriptionCallback func(msg *Mesh, info *rclgo.MessageInfo, err error)

// NewMeshSubscription creates and returns a new subscription for the
// Mesh
func NewMeshSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MeshSubscriptionCallback) (*MeshSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Mesh
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MeshTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MeshSubscription{sub}, nil
}

func (s *MeshSubscription) TakeMessage(out *Mesh) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMeshSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMeshSlice(dst, src []Mesh) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MeshTypeSupport types.MessageTypeSupport = _MeshTypeSupport{}

type _MeshTypeSupport struct{}

func (t _MeshTypeSupport) New() types.Message {
	return NewMesh()
}

func (t _MeshTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.shape_msgs__msg__Mesh
	return (unsafe.Pointer)(C.shape_msgs__msg__Mesh__create())
}

func (t _MeshTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.shape_msgs__msg__Mesh__destroy((*C.shape_msgs__msg__Mesh)(pointer_to_free))
}

func (t _MeshTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Mesh)
	mem := (*C.shape_msgs__msg__Mesh)(dst)
	MeshTriangle__Sequence_to_C(&mem.triangles, m.Triangles)
	geometry_msgs_msg.Point__Sequence_to_C((*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.vertices)), m.Vertices)
}

func (t _MeshTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Mesh)
	mem := (*C.shape_msgs__msg__Mesh)(ros2_message_buffer)
	MeshTriangle__Sequence_to_Go(&m.Triangles, mem.triangles)
	geometry_msgs_msg.Point__Sequence_to_Go(&m.Vertices, *(*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.vertices)))
}

func (t _MeshTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__shape_msgs__msg__Mesh())
}

type CMesh = C.shape_msgs__msg__Mesh
type CMesh__Sequence = C.shape_msgs__msg__Mesh__Sequence

func Mesh__Sequence_to_Go(goSlice *[]Mesh, cSlice CMesh__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Mesh, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MeshTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Mesh__Sequence_to_C(cSlice *CMesh__Sequence, goSlice []Mesh) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.shape_msgs__msg__Mesh)(C.malloc(C.sizeof_struct_shape_msgs__msg__Mesh * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MeshTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Mesh__Array_to_Go(goSlice []Mesh, cSlice []CMesh) {
	for i := 0; i < len(cSlice); i++ {
		MeshTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Mesh__Array_to_C(cSlice []CMesh, goSlice []Mesh) {
	for i := 0; i < len(goSlice); i++ {
		MeshTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
