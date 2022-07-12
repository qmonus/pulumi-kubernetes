// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type PodDisruptionBudgetPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpecPatchPtrOutput `pulumi:"spec"`
	// Most recently observed status of the PodDisruptionBudget.
	Status PodDisruptionBudgetStatusPatchPtrOutput `pulumi:"status"`
}

// NewPodDisruptionBudgetPatch registers a new resource with the given unique name, arguments, and options.
func NewPodDisruptionBudgetPatch(ctx *pulumi.Context,
	name string, args *PodDisruptionBudgetPatchArgs, opts ...pulumi.ResourceOption) (*PodDisruptionBudgetPatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("policy/v1beta1")
	args.Kind = pulumi.StringPtr("PodDisruptionBudget")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:policy/v1:PodDisruptionBudgetPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource PodDisruptionBudgetPatch
	err := ctx.RegisterResource("kubernetes:policy/v1beta1:PodDisruptionBudgetPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodDisruptionBudgetPatch gets an existing PodDisruptionBudgetPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodDisruptionBudgetPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodDisruptionBudgetPatchState, opts ...pulumi.ResourceOption) (*PodDisruptionBudgetPatch, error) {
	var resource PodDisruptionBudgetPatch
	err := ctx.ReadResource("kubernetes:policy/v1beta1:PodDisruptionBudgetPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodDisruptionBudgetPatch resources.
type podDisruptionBudgetPatchState struct {
}

type PodDisruptionBudgetPatchState struct {
}

func (PodDisruptionBudgetPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetPatchState)(nil)).Elem()
}

type podDisruptionBudgetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a PodDisruptionBudgetPatch resource.
type PodDisruptionBudgetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchInput
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpecPatchPtrInput
}

func (PodDisruptionBudgetPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetPatchArgs)(nil)).Elem()
}

type PodDisruptionBudgetPatchInput interface {
	pulumi.Input

	ToPodDisruptionBudgetPatchOutput() PodDisruptionBudgetPatchOutput
	ToPodDisruptionBudgetPatchOutputWithContext(ctx context.Context) PodDisruptionBudgetPatchOutput
}

func (*PodDisruptionBudgetPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PodDisruptionBudgetPatch)(nil)).Elem()
}

func (i *PodDisruptionBudgetPatch) ToPodDisruptionBudgetPatchOutput() PodDisruptionBudgetPatchOutput {
	return i.ToPodDisruptionBudgetPatchOutputWithContext(context.Background())
}

func (i *PodDisruptionBudgetPatch) ToPodDisruptionBudgetPatchOutputWithContext(ctx context.Context) PodDisruptionBudgetPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetPatchOutput)
}

// PodDisruptionBudgetPatchArrayInput is an input type that accepts PodDisruptionBudgetPatchArray and PodDisruptionBudgetPatchArrayOutput values.
// You can construct a concrete instance of `PodDisruptionBudgetPatchArrayInput` via:
//
//          PodDisruptionBudgetPatchArray{ PodDisruptionBudgetPatchArgs{...} }
type PodDisruptionBudgetPatchArrayInput interface {
	pulumi.Input

	ToPodDisruptionBudgetPatchArrayOutput() PodDisruptionBudgetPatchArrayOutput
	ToPodDisruptionBudgetPatchArrayOutputWithContext(context.Context) PodDisruptionBudgetPatchArrayOutput
}

type PodDisruptionBudgetPatchArray []PodDisruptionBudgetPatchInput

func (PodDisruptionBudgetPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodDisruptionBudgetPatch)(nil)).Elem()
}

func (i PodDisruptionBudgetPatchArray) ToPodDisruptionBudgetPatchArrayOutput() PodDisruptionBudgetPatchArrayOutput {
	return i.ToPodDisruptionBudgetPatchArrayOutputWithContext(context.Background())
}

func (i PodDisruptionBudgetPatchArray) ToPodDisruptionBudgetPatchArrayOutputWithContext(ctx context.Context) PodDisruptionBudgetPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetPatchArrayOutput)
}

// PodDisruptionBudgetPatchMapInput is an input type that accepts PodDisruptionBudgetPatchMap and PodDisruptionBudgetPatchMapOutput values.
// You can construct a concrete instance of `PodDisruptionBudgetPatchMapInput` via:
//
//          PodDisruptionBudgetPatchMap{ "key": PodDisruptionBudgetPatchArgs{...} }
type PodDisruptionBudgetPatchMapInput interface {
	pulumi.Input

	ToPodDisruptionBudgetPatchMapOutput() PodDisruptionBudgetPatchMapOutput
	ToPodDisruptionBudgetPatchMapOutputWithContext(context.Context) PodDisruptionBudgetPatchMapOutput
}

type PodDisruptionBudgetPatchMap map[string]PodDisruptionBudgetPatchInput

func (PodDisruptionBudgetPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodDisruptionBudgetPatch)(nil)).Elem()
}

func (i PodDisruptionBudgetPatchMap) ToPodDisruptionBudgetPatchMapOutput() PodDisruptionBudgetPatchMapOutput {
	return i.ToPodDisruptionBudgetPatchMapOutputWithContext(context.Background())
}

func (i PodDisruptionBudgetPatchMap) ToPodDisruptionBudgetPatchMapOutputWithContext(ctx context.Context) PodDisruptionBudgetPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetPatchMapOutput)
}

type PodDisruptionBudgetPatchOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodDisruptionBudgetPatch)(nil)).Elem()
}

func (o PodDisruptionBudgetPatchOutput) ToPodDisruptionBudgetPatchOutput() PodDisruptionBudgetPatchOutput {
	return o
}

func (o PodDisruptionBudgetPatchOutput) ToPodDisruptionBudgetPatchOutputWithContext(ctx context.Context) PodDisruptionBudgetPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodDisruptionBudgetPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudgetPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodDisruptionBudgetPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudgetPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PodDisruptionBudgetPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudgetPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Specification of the desired behavior of the PodDisruptionBudget.
func (o PodDisruptionBudgetPatchOutput) Spec() PodDisruptionBudgetSpecPatchPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudgetPatch) PodDisruptionBudgetSpecPatchPtrOutput { return v.Spec }).(PodDisruptionBudgetSpecPatchPtrOutput)
}

// Most recently observed status of the PodDisruptionBudget.
func (o PodDisruptionBudgetPatchOutput) Status() PodDisruptionBudgetStatusPatchPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudgetPatch) PodDisruptionBudgetStatusPatchPtrOutput { return v.Status }).(PodDisruptionBudgetStatusPatchPtrOutput)
}

type PodDisruptionBudgetPatchArrayOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodDisruptionBudgetPatch)(nil)).Elem()
}

func (o PodDisruptionBudgetPatchArrayOutput) ToPodDisruptionBudgetPatchArrayOutput() PodDisruptionBudgetPatchArrayOutput {
	return o
}

func (o PodDisruptionBudgetPatchArrayOutput) ToPodDisruptionBudgetPatchArrayOutputWithContext(ctx context.Context) PodDisruptionBudgetPatchArrayOutput {
	return o
}

func (o PodDisruptionBudgetPatchArrayOutput) Index(i pulumi.IntInput) PodDisruptionBudgetPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PodDisruptionBudgetPatch {
		return vs[0].([]*PodDisruptionBudgetPatch)[vs[1].(int)]
	}).(PodDisruptionBudgetPatchOutput)
}

type PodDisruptionBudgetPatchMapOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodDisruptionBudgetPatch)(nil)).Elem()
}

func (o PodDisruptionBudgetPatchMapOutput) ToPodDisruptionBudgetPatchMapOutput() PodDisruptionBudgetPatchMapOutput {
	return o
}

func (o PodDisruptionBudgetPatchMapOutput) ToPodDisruptionBudgetPatchMapOutputWithContext(ctx context.Context) PodDisruptionBudgetPatchMapOutput {
	return o
}

func (o PodDisruptionBudgetPatchMapOutput) MapIndex(k pulumi.StringInput) PodDisruptionBudgetPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PodDisruptionBudgetPatch {
		return vs[0].(map[string]*PodDisruptionBudgetPatch)[vs[1].(string)]
	}).(PodDisruptionBudgetPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodDisruptionBudgetPatchInput)(nil)).Elem(), &PodDisruptionBudgetPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodDisruptionBudgetPatchArrayInput)(nil)).Elem(), PodDisruptionBudgetPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodDisruptionBudgetPatchMapInput)(nil)).Elem(), PodDisruptionBudgetPatchMap{})
	pulumi.RegisterOutputType(PodDisruptionBudgetPatchOutput{})
	pulumi.RegisterOutputType(PodDisruptionBudgetPatchArrayOutput{})
	pulumi.RegisterOutputType(PodDisruptionBudgetPatchMapOutput{})
}
