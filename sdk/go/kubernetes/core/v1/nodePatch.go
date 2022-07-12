// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd).
type NodePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec defines the behavior of a node. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec NodeSpecPatchPtrOutput `pulumi:"spec"`
	// Most recently observed status of the node. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status NodeStatusPatchPtrOutput `pulumi:"status"`
}

// NewNodePatch registers a new resource with the given unique name, arguments, and options.
func NewNodePatch(ctx *pulumi.Context,
	name string, args *NodePatchArgs, opts ...pulumi.ResourceOption) (*NodePatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Node")
	var resource NodePatch
	err := ctx.RegisterResource("kubernetes:core/v1:NodePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodePatch gets an existing NodePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodePatchState, opts ...pulumi.ResourceOption) (*NodePatch, error) {
	var resource NodePatch
	err := ctx.ReadResource("kubernetes:core/v1:NodePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodePatch resources.
type nodePatchState struct {
}

type NodePatchState struct {
}

func (NodePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePatchState)(nil)).Elem()
}

type nodePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec defines the behavior of a node. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *NodeSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a NodePatch resource.
type NodePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchInput
	// Spec defines the behavior of a node. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec NodeSpecPatchPtrInput
}

func (NodePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePatchArgs)(nil)).Elem()
}

type NodePatchInput interface {
	pulumi.Input

	ToNodePatchOutput() NodePatchOutput
	ToNodePatchOutputWithContext(ctx context.Context) NodePatchOutput
}

func (*NodePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePatch)(nil)).Elem()
}

func (i *NodePatch) ToNodePatchOutput() NodePatchOutput {
	return i.ToNodePatchOutputWithContext(context.Background())
}

func (i *NodePatch) ToNodePatchOutputWithContext(ctx context.Context) NodePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePatchOutput)
}

// NodePatchArrayInput is an input type that accepts NodePatchArray and NodePatchArrayOutput values.
// You can construct a concrete instance of `NodePatchArrayInput` via:
//
//          NodePatchArray{ NodePatchArgs{...} }
type NodePatchArrayInput interface {
	pulumi.Input

	ToNodePatchArrayOutput() NodePatchArrayOutput
	ToNodePatchArrayOutputWithContext(context.Context) NodePatchArrayOutput
}

type NodePatchArray []NodePatchInput

func (NodePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodePatch)(nil)).Elem()
}

func (i NodePatchArray) ToNodePatchArrayOutput() NodePatchArrayOutput {
	return i.ToNodePatchArrayOutputWithContext(context.Background())
}

func (i NodePatchArray) ToNodePatchArrayOutputWithContext(ctx context.Context) NodePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePatchArrayOutput)
}

// NodePatchMapInput is an input type that accepts NodePatchMap and NodePatchMapOutput values.
// You can construct a concrete instance of `NodePatchMapInput` via:
//
//          NodePatchMap{ "key": NodePatchArgs{...} }
type NodePatchMapInput interface {
	pulumi.Input

	ToNodePatchMapOutput() NodePatchMapOutput
	ToNodePatchMapOutputWithContext(context.Context) NodePatchMapOutput
}

type NodePatchMap map[string]NodePatchInput

func (NodePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodePatch)(nil)).Elem()
}

func (i NodePatchMap) ToNodePatchMapOutput() NodePatchMapOutput {
	return i.ToNodePatchMapOutputWithContext(context.Background())
}

func (i NodePatchMap) ToNodePatchMapOutputWithContext(ctx context.Context) NodePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePatchMapOutput)
}

type NodePatchOutput struct{ *pulumi.OutputState }

func (NodePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePatch)(nil)).Elem()
}

func (o NodePatchOutput) ToNodePatchOutput() NodePatchOutput {
	return o
}

func (o NodePatchOutput) ToNodePatchOutputWithContext(ctx context.Context) NodePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o NodePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o NodePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o NodePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *NodePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec defines the behavior of a node. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o NodePatchOutput) Spec() NodeSpecPatchPtrOutput {
	return o.ApplyT(func(v *NodePatch) NodeSpecPatchPtrOutput { return v.Spec }).(NodeSpecPatchPtrOutput)
}

// Most recently observed status of the node. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o NodePatchOutput) Status() NodeStatusPatchPtrOutput {
	return o.ApplyT(func(v *NodePatch) NodeStatusPatchPtrOutput { return v.Status }).(NodeStatusPatchPtrOutput)
}

type NodePatchArrayOutput struct{ *pulumi.OutputState }

func (NodePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodePatch)(nil)).Elem()
}

func (o NodePatchArrayOutput) ToNodePatchArrayOutput() NodePatchArrayOutput {
	return o
}

func (o NodePatchArrayOutput) ToNodePatchArrayOutputWithContext(ctx context.Context) NodePatchArrayOutput {
	return o
}

func (o NodePatchArrayOutput) Index(i pulumi.IntInput) NodePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodePatch {
		return vs[0].([]*NodePatch)[vs[1].(int)]
	}).(NodePatchOutput)
}

type NodePatchMapOutput struct{ *pulumi.OutputState }

func (NodePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodePatch)(nil)).Elem()
}

func (o NodePatchMapOutput) ToNodePatchMapOutput() NodePatchMapOutput {
	return o
}

func (o NodePatchMapOutput) ToNodePatchMapOutputWithContext(ctx context.Context) NodePatchMapOutput {
	return o
}

func (o NodePatchMapOutput) MapIndex(k pulumi.StringInput) NodePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodePatch {
		return vs[0].(map[string]*NodePatch)[vs[1].(string)]
	}).(NodePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodePatchInput)(nil)).Elem(), &NodePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodePatchArrayInput)(nil)).Elem(), NodePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodePatchMapInput)(nil)).Elem(), NodePatchMap{})
	pulumi.RegisterOutputType(NodePatchOutput{})
	pulumi.RegisterOutputType(NodePatchArrayOutput{})
	pulumi.RegisterOutputType(NodePatchMapOutput{})
}
