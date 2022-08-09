// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/installation-configuration/#server-side-apply) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// Namespace provides a scope for Names. Use of multiple namespaces is optional.
type NamespacePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec defines the behavior of the Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec NamespaceSpecPatchPtrOutput `pulumi:"spec"`
	// Status describes the current status of a Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status NamespaceStatusPatchPtrOutput `pulumi:"status"`
}

// NewNamespacePatch registers a new resource with the given unique name, arguments, and options.
func NewNamespacePatch(ctx *pulumi.Context,
	name string, args *NamespacePatchArgs, opts ...pulumi.ResourceOption) (*NamespacePatch, error) {
	if args == nil {
		args = &NamespacePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Namespace")
	var resource NamespacePatch
	err := ctx.RegisterResource("kubernetes:core/v1:NamespacePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespacePatch gets an existing NamespacePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespacePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespacePatchState, opts ...pulumi.ResourceOption) (*NamespacePatch, error) {
	var resource NamespacePatch
	err := ctx.ReadResource("kubernetes:core/v1:NamespacePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespacePatch resources.
type namespacePatchState struct {
}

type NamespacePatchState struct {
}

func (NamespacePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespacePatchState)(nil)).Elem()
}

type namespacePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec defines the behavior of the Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *NamespaceSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a NamespacePatch resource.
type NamespacePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec defines the behavior of the Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec NamespaceSpecPatchPtrInput
}

func (NamespacePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespacePatchArgs)(nil)).Elem()
}

type NamespacePatchInput interface {
	pulumi.Input

	ToNamespacePatchOutput() NamespacePatchOutput
	ToNamespacePatchOutputWithContext(ctx context.Context) NamespacePatchOutput
}

func (*NamespacePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespacePatch)(nil)).Elem()
}

func (i *NamespacePatch) ToNamespacePatchOutput() NamespacePatchOutput {
	return i.ToNamespacePatchOutputWithContext(context.Background())
}

func (i *NamespacePatch) ToNamespacePatchOutputWithContext(ctx context.Context) NamespacePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePatchOutput)
}

// NamespacePatchArrayInput is an input type that accepts NamespacePatchArray and NamespacePatchArrayOutput values.
// You can construct a concrete instance of `NamespacePatchArrayInput` via:
//
//          NamespacePatchArray{ NamespacePatchArgs{...} }
type NamespacePatchArrayInput interface {
	pulumi.Input

	ToNamespacePatchArrayOutput() NamespacePatchArrayOutput
	ToNamespacePatchArrayOutputWithContext(context.Context) NamespacePatchArrayOutput
}

type NamespacePatchArray []NamespacePatchInput

func (NamespacePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespacePatch)(nil)).Elem()
}

func (i NamespacePatchArray) ToNamespacePatchArrayOutput() NamespacePatchArrayOutput {
	return i.ToNamespacePatchArrayOutputWithContext(context.Background())
}

func (i NamespacePatchArray) ToNamespacePatchArrayOutputWithContext(ctx context.Context) NamespacePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePatchArrayOutput)
}

// NamespacePatchMapInput is an input type that accepts NamespacePatchMap and NamespacePatchMapOutput values.
// You can construct a concrete instance of `NamespacePatchMapInput` via:
//
//          NamespacePatchMap{ "key": NamespacePatchArgs{...} }
type NamespacePatchMapInput interface {
	pulumi.Input

	ToNamespacePatchMapOutput() NamespacePatchMapOutput
	ToNamespacePatchMapOutputWithContext(context.Context) NamespacePatchMapOutput
}

type NamespacePatchMap map[string]NamespacePatchInput

func (NamespacePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespacePatch)(nil)).Elem()
}

func (i NamespacePatchMap) ToNamespacePatchMapOutput() NamespacePatchMapOutput {
	return i.ToNamespacePatchMapOutputWithContext(context.Background())
}

func (i NamespacePatchMap) ToNamespacePatchMapOutputWithContext(ctx context.Context) NamespacePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePatchMapOutput)
}

type NamespacePatchOutput struct{ *pulumi.OutputState }

func (NamespacePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespacePatch)(nil)).Elem()
}

func (o NamespacePatchOutput) ToNamespacePatchOutput() NamespacePatchOutput {
	return o
}

func (o NamespacePatchOutput) ToNamespacePatchOutputWithContext(ctx context.Context) NamespacePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o NamespacePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o NamespacePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o NamespacePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *NamespacePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec defines the behavior of the Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o NamespacePatchOutput) Spec() NamespaceSpecPatchPtrOutput {
	return o.ApplyT(func(v *NamespacePatch) NamespaceSpecPatchPtrOutput { return v.Spec }).(NamespaceSpecPatchPtrOutput)
}

// Status describes the current status of a Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o NamespacePatchOutput) Status() NamespaceStatusPatchPtrOutput {
	return o.ApplyT(func(v *NamespacePatch) NamespaceStatusPatchPtrOutput { return v.Status }).(NamespaceStatusPatchPtrOutput)
}

type NamespacePatchArrayOutput struct{ *pulumi.OutputState }

func (NamespacePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespacePatch)(nil)).Elem()
}

func (o NamespacePatchArrayOutput) ToNamespacePatchArrayOutput() NamespacePatchArrayOutput {
	return o
}

func (o NamespacePatchArrayOutput) ToNamespacePatchArrayOutputWithContext(ctx context.Context) NamespacePatchArrayOutput {
	return o
}

func (o NamespacePatchArrayOutput) Index(i pulumi.IntInput) NamespacePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NamespacePatch {
		return vs[0].([]*NamespacePatch)[vs[1].(int)]
	}).(NamespacePatchOutput)
}

type NamespacePatchMapOutput struct{ *pulumi.OutputState }

func (NamespacePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespacePatch)(nil)).Elem()
}

func (o NamespacePatchMapOutput) ToNamespacePatchMapOutput() NamespacePatchMapOutput {
	return o
}

func (o NamespacePatchMapOutput) ToNamespacePatchMapOutputWithContext(ctx context.Context) NamespacePatchMapOutput {
	return o
}

func (o NamespacePatchMapOutput) MapIndex(k pulumi.StringInput) NamespacePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NamespacePatch {
		return vs[0].(map[string]*NamespacePatch)[vs[1].(string)]
	}).(NamespacePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespacePatchInput)(nil)).Elem(), &NamespacePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespacePatchArrayInput)(nil)).Elem(), NamespacePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespacePatchMapInput)(nil)).Elem(), NamespacePatchMap{})
	pulumi.RegisterOutputType(NamespacePatchOutput{})
	pulumi.RegisterOutputType(NamespacePatchArrayOutput{})
	pulumi.RegisterOutputType(NamespacePatchMapOutput{})
}
