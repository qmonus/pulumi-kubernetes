// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.
//
// VolumeAttachment objects are non-namespaced.
type VolumeAttachmentPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpecPatchPtrOutput `pulumi:"spec"`
	// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
	Status VolumeAttachmentStatusPatchPtrOutput `pulumi:"status"`
}

// NewVolumeAttachmentPatch registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachmentPatch(ctx *pulumi.Context,
	name string, args *VolumeAttachmentPatchArgs, opts ...pulumi.ResourceOption) (*VolumeAttachmentPatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("VolumeAttachment")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1:VolumeAttachmentPatch"),
		},
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1beta1:VolumeAttachmentPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource VolumeAttachmentPatch
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1alpha1:VolumeAttachmentPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachmentPatch gets an existing VolumeAttachmentPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachmentPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachmentPatchState, opts ...pulumi.ResourceOption) (*VolumeAttachmentPatch, error) {
	var resource VolumeAttachmentPatch
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1alpha1:VolumeAttachmentPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttachmentPatch resources.
type volumeAttachmentPatchState struct {
}

type VolumeAttachmentPatchState struct {
}

func (VolumeAttachmentPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentPatchState)(nil)).Elem()
}

type volumeAttachmentPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec *VolumeAttachmentSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a VolumeAttachmentPatch resource.
type VolumeAttachmentPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchInput
	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpecPatchPtrInput
}

func (VolumeAttachmentPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentPatchArgs)(nil)).Elem()
}

type VolumeAttachmentPatchInput interface {
	pulumi.Input

	ToVolumeAttachmentPatchOutput() VolumeAttachmentPatchOutput
	ToVolumeAttachmentPatchOutputWithContext(ctx context.Context) VolumeAttachmentPatchOutput
}

func (*VolumeAttachmentPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachmentPatch)(nil)).Elem()
}

func (i *VolumeAttachmentPatch) ToVolumeAttachmentPatchOutput() VolumeAttachmentPatchOutput {
	return i.ToVolumeAttachmentPatchOutputWithContext(context.Background())
}

func (i *VolumeAttachmentPatch) ToVolumeAttachmentPatchOutputWithContext(ctx context.Context) VolumeAttachmentPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentPatchOutput)
}

// VolumeAttachmentPatchArrayInput is an input type that accepts VolumeAttachmentPatchArray and VolumeAttachmentPatchArrayOutput values.
// You can construct a concrete instance of `VolumeAttachmentPatchArrayInput` via:
//
//          VolumeAttachmentPatchArray{ VolumeAttachmentPatchArgs{...} }
type VolumeAttachmentPatchArrayInput interface {
	pulumi.Input

	ToVolumeAttachmentPatchArrayOutput() VolumeAttachmentPatchArrayOutput
	ToVolumeAttachmentPatchArrayOutputWithContext(context.Context) VolumeAttachmentPatchArrayOutput
}

type VolumeAttachmentPatchArray []VolumeAttachmentPatchInput

func (VolumeAttachmentPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttachmentPatch)(nil)).Elem()
}

func (i VolumeAttachmentPatchArray) ToVolumeAttachmentPatchArrayOutput() VolumeAttachmentPatchArrayOutput {
	return i.ToVolumeAttachmentPatchArrayOutputWithContext(context.Background())
}

func (i VolumeAttachmentPatchArray) ToVolumeAttachmentPatchArrayOutputWithContext(ctx context.Context) VolumeAttachmentPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentPatchArrayOutput)
}

// VolumeAttachmentPatchMapInput is an input type that accepts VolumeAttachmentPatchMap and VolumeAttachmentPatchMapOutput values.
// You can construct a concrete instance of `VolumeAttachmentPatchMapInput` via:
//
//          VolumeAttachmentPatchMap{ "key": VolumeAttachmentPatchArgs{...} }
type VolumeAttachmentPatchMapInput interface {
	pulumi.Input

	ToVolumeAttachmentPatchMapOutput() VolumeAttachmentPatchMapOutput
	ToVolumeAttachmentPatchMapOutputWithContext(context.Context) VolumeAttachmentPatchMapOutput
}

type VolumeAttachmentPatchMap map[string]VolumeAttachmentPatchInput

func (VolumeAttachmentPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttachmentPatch)(nil)).Elem()
}

func (i VolumeAttachmentPatchMap) ToVolumeAttachmentPatchMapOutput() VolumeAttachmentPatchMapOutput {
	return i.ToVolumeAttachmentPatchMapOutputWithContext(context.Background())
}

func (i VolumeAttachmentPatchMap) ToVolumeAttachmentPatchMapOutputWithContext(ctx context.Context) VolumeAttachmentPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentPatchMapOutput)
}

type VolumeAttachmentPatchOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachmentPatch)(nil)).Elem()
}

func (o VolumeAttachmentPatchOutput) ToVolumeAttachmentPatchOutput() VolumeAttachmentPatchOutput {
	return o
}

func (o VolumeAttachmentPatchOutput) ToVolumeAttachmentPatchOutputWithContext(ctx context.Context) VolumeAttachmentPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VolumeAttachmentPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VolumeAttachmentPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VolumeAttachmentPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
func (o VolumeAttachmentPatchOutput) Spec() VolumeAttachmentSpecPatchPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentPatch) VolumeAttachmentSpecPatchPtrOutput { return v.Spec }).(VolumeAttachmentSpecPatchPtrOutput)
}

// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
func (o VolumeAttachmentPatchOutput) Status() VolumeAttachmentStatusPatchPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentPatch) VolumeAttachmentStatusPatchPtrOutput { return v.Status }).(VolumeAttachmentStatusPatchPtrOutput)
}

type VolumeAttachmentPatchArrayOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttachmentPatch)(nil)).Elem()
}

func (o VolumeAttachmentPatchArrayOutput) ToVolumeAttachmentPatchArrayOutput() VolumeAttachmentPatchArrayOutput {
	return o
}

func (o VolumeAttachmentPatchArrayOutput) ToVolumeAttachmentPatchArrayOutputWithContext(ctx context.Context) VolumeAttachmentPatchArrayOutput {
	return o
}

func (o VolumeAttachmentPatchArrayOutput) Index(i pulumi.IntInput) VolumeAttachmentPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VolumeAttachmentPatch {
		return vs[0].([]*VolumeAttachmentPatch)[vs[1].(int)]
	}).(VolumeAttachmentPatchOutput)
}

type VolumeAttachmentPatchMapOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttachmentPatch)(nil)).Elem()
}

func (o VolumeAttachmentPatchMapOutput) ToVolumeAttachmentPatchMapOutput() VolumeAttachmentPatchMapOutput {
	return o
}

func (o VolumeAttachmentPatchMapOutput) ToVolumeAttachmentPatchMapOutputWithContext(ctx context.Context) VolumeAttachmentPatchMapOutput {
	return o
}

func (o VolumeAttachmentPatchMapOutput) MapIndex(k pulumi.StringInput) VolumeAttachmentPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VolumeAttachmentPatch {
		return vs[0].(map[string]*VolumeAttachmentPatch)[vs[1].(string)]
	}).(VolumeAttachmentPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentPatchInput)(nil)).Elem(), &VolumeAttachmentPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentPatchArrayInput)(nil)).Elem(), VolumeAttachmentPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentPatchMapInput)(nil)).Elem(), VolumeAttachmentPatchMap{})
	pulumi.RegisterOutputType(VolumeAttachmentPatchOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentPatchArrayOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentPatchMapOutput{})
}
