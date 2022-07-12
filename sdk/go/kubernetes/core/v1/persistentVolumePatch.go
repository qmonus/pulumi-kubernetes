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

// PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
type PersistentVolumePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	Spec PersistentVolumeSpecPatchPtrOutput `pulumi:"spec"`
	// status represents the current information/status for the persistent volume. Populated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	Status PersistentVolumeStatusPatchPtrOutput `pulumi:"status"`
}

// NewPersistentVolumePatch registers a new resource with the given unique name, arguments, and options.
func NewPersistentVolumePatch(ctx *pulumi.Context,
	name string, args *PersistentVolumePatchArgs, opts ...pulumi.ResourceOption) (*PersistentVolumePatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("PersistentVolume")
	var resource PersistentVolumePatch
	err := ctx.RegisterResource("kubernetes:core/v1:PersistentVolumePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistentVolumePatch gets an existing PersistentVolumePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistentVolumePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistentVolumePatchState, opts ...pulumi.ResourceOption) (*PersistentVolumePatch, error) {
	var resource PersistentVolumePatch
	err := ctx.ReadResource("kubernetes:core/v1:PersistentVolumePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistentVolumePatch resources.
type persistentVolumePatchState struct {
}

type PersistentVolumePatchState struct {
}

func (PersistentVolumePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumePatchState)(nil)).Elem()
}

type persistentVolumePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	Spec *PersistentVolumeSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a PersistentVolumePatch resource.
type PersistentVolumePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchInput
	// spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	Spec PersistentVolumeSpecPatchPtrInput
}

func (PersistentVolumePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumePatchArgs)(nil)).Elem()
}

type PersistentVolumePatchInput interface {
	pulumi.Input

	ToPersistentVolumePatchOutput() PersistentVolumePatchOutput
	ToPersistentVolumePatchOutputWithContext(ctx context.Context) PersistentVolumePatchOutput
}

func (*PersistentVolumePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumePatch)(nil)).Elem()
}

func (i *PersistentVolumePatch) ToPersistentVolumePatchOutput() PersistentVolumePatchOutput {
	return i.ToPersistentVolumePatchOutputWithContext(context.Background())
}

func (i *PersistentVolumePatch) ToPersistentVolumePatchOutputWithContext(ctx context.Context) PersistentVolumePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumePatchOutput)
}

// PersistentVolumePatchArrayInput is an input type that accepts PersistentVolumePatchArray and PersistentVolumePatchArrayOutput values.
// You can construct a concrete instance of `PersistentVolumePatchArrayInput` via:
//
//          PersistentVolumePatchArray{ PersistentVolumePatchArgs{...} }
type PersistentVolumePatchArrayInput interface {
	pulumi.Input

	ToPersistentVolumePatchArrayOutput() PersistentVolumePatchArrayOutput
	ToPersistentVolumePatchArrayOutputWithContext(context.Context) PersistentVolumePatchArrayOutput
}

type PersistentVolumePatchArray []PersistentVolumePatchInput

func (PersistentVolumePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistentVolumePatch)(nil)).Elem()
}

func (i PersistentVolumePatchArray) ToPersistentVolumePatchArrayOutput() PersistentVolumePatchArrayOutput {
	return i.ToPersistentVolumePatchArrayOutputWithContext(context.Background())
}

func (i PersistentVolumePatchArray) ToPersistentVolumePatchArrayOutputWithContext(ctx context.Context) PersistentVolumePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumePatchArrayOutput)
}

// PersistentVolumePatchMapInput is an input type that accepts PersistentVolumePatchMap and PersistentVolumePatchMapOutput values.
// You can construct a concrete instance of `PersistentVolumePatchMapInput` via:
//
//          PersistentVolumePatchMap{ "key": PersistentVolumePatchArgs{...} }
type PersistentVolumePatchMapInput interface {
	pulumi.Input

	ToPersistentVolumePatchMapOutput() PersistentVolumePatchMapOutput
	ToPersistentVolumePatchMapOutputWithContext(context.Context) PersistentVolumePatchMapOutput
}

type PersistentVolumePatchMap map[string]PersistentVolumePatchInput

func (PersistentVolumePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistentVolumePatch)(nil)).Elem()
}

func (i PersistentVolumePatchMap) ToPersistentVolumePatchMapOutput() PersistentVolumePatchMapOutput {
	return i.ToPersistentVolumePatchMapOutputWithContext(context.Background())
}

func (i PersistentVolumePatchMap) ToPersistentVolumePatchMapOutputWithContext(ctx context.Context) PersistentVolumePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumePatchMapOutput)
}

type PersistentVolumePatchOutput struct{ *pulumi.OutputState }

func (PersistentVolumePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumePatch)(nil)).Elem()
}

func (o PersistentVolumePatchOutput) ToPersistentVolumePatchOutput() PersistentVolumePatchOutput {
	return o
}

func (o PersistentVolumePatchOutput) ToPersistentVolumePatchOutputWithContext(ctx context.Context) PersistentVolumePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PersistentVolumePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentVolumePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PersistentVolumePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentVolumePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PersistentVolumePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PersistentVolumePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
func (o PersistentVolumePatchOutput) Spec() PersistentVolumeSpecPatchPtrOutput {
	return o.ApplyT(func(v *PersistentVolumePatch) PersistentVolumeSpecPatchPtrOutput { return v.Spec }).(PersistentVolumeSpecPatchPtrOutput)
}

// status represents the current information/status for the persistent volume. Populated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
func (o PersistentVolumePatchOutput) Status() PersistentVolumeStatusPatchPtrOutput {
	return o.ApplyT(func(v *PersistentVolumePatch) PersistentVolumeStatusPatchPtrOutput { return v.Status }).(PersistentVolumeStatusPatchPtrOutput)
}

type PersistentVolumePatchArrayOutput struct{ *pulumi.OutputState }

func (PersistentVolumePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistentVolumePatch)(nil)).Elem()
}

func (o PersistentVolumePatchArrayOutput) ToPersistentVolumePatchArrayOutput() PersistentVolumePatchArrayOutput {
	return o
}

func (o PersistentVolumePatchArrayOutput) ToPersistentVolumePatchArrayOutputWithContext(ctx context.Context) PersistentVolumePatchArrayOutput {
	return o
}

func (o PersistentVolumePatchArrayOutput) Index(i pulumi.IntInput) PersistentVolumePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersistentVolumePatch {
		return vs[0].([]*PersistentVolumePatch)[vs[1].(int)]
	}).(PersistentVolumePatchOutput)
}

type PersistentVolumePatchMapOutput struct{ *pulumi.OutputState }

func (PersistentVolumePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistentVolumePatch)(nil)).Elem()
}

func (o PersistentVolumePatchMapOutput) ToPersistentVolumePatchMapOutput() PersistentVolumePatchMapOutput {
	return o
}

func (o PersistentVolumePatchMapOutput) ToPersistentVolumePatchMapOutputWithContext(ctx context.Context) PersistentVolumePatchMapOutput {
	return o
}

func (o PersistentVolumePatchMapOutput) MapIndex(k pulumi.StringInput) PersistentVolumePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersistentVolumePatch {
		return vs[0].(map[string]*PersistentVolumePatch)[vs[1].(string)]
	}).(PersistentVolumePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumePatchInput)(nil)).Elem(), &PersistentVolumePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumePatchArrayInput)(nil)).Elem(), PersistentVolumePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumePatchMapInput)(nil)).Elem(), PersistentVolumePatchMap{})
	pulumi.RegisterOutputType(PersistentVolumePatchOutput{})
	pulumi.RegisterOutputType(PersistentVolumePatchArrayOutput{})
	pulumi.RegisterOutputType(PersistentVolumePatchMapOutput{})
}
