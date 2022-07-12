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

// DEPRECATED 1.9 - This group version of NetworkPolicy is deprecated by networking/v1/NetworkPolicy. NetworkPolicy describes what network traffic is allowed for a set of Pods
type NetworkPolicyPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior for this NetworkPolicy.
	Spec NetworkPolicySpecPatchPtrOutput `pulumi:"spec"`
}

// NewNetworkPolicyPatch registers a new resource with the given unique name, arguments, and options.
func NewNetworkPolicyPatch(ctx *pulumi.Context,
	name string, args *NetworkPolicyPatchArgs, opts ...pulumi.ResourceOption) (*NetworkPolicyPatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("NetworkPolicy")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:networking.k8s.io/v1:NetworkPolicyPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkPolicyPatch
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:NetworkPolicyPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPolicyPatch gets an existing NetworkPolicyPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPolicyPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPolicyPatchState, opts ...pulumi.ResourceOption) (*NetworkPolicyPatch, error) {
	var resource NetworkPolicyPatch
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:NetworkPolicyPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPolicyPatch resources.
type networkPolicyPatchState struct {
}

type NetworkPolicyPatchState struct {
}

func (NetworkPolicyPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPolicyPatchState)(nil)).Elem()
}

type networkPolicyPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Specification of the desired behavior for this NetworkPolicy.
	Spec *NetworkPolicySpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a NetworkPolicyPatch resource.
type NetworkPolicyPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchInput
	// Specification of the desired behavior for this NetworkPolicy.
	Spec NetworkPolicySpecPatchPtrInput
}

func (NetworkPolicyPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPolicyPatchArgs)(nil)).Elem()
}

type NetworkPolicyPatchInput interface {
	pulumi.Input

	ToNetworkPolicyPatchOutput() NetworkPolicyPatchOutput
	ToNetworkPolicyPatchOutputWithContext(ctx context.Context) NetworkPolicyPatchOutput
}

func (*NetworkPolicyPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicyPatch)(nil)).Elem()
}

func (i *NetworkPolicyPatch) ToNetworkPolicyPatchOutput() NetworkPolicyPatchOutput {
	return i.ToNetworkPolicyPatchOutputWithContext(context.Background())
}

func (i *NetworkPolicyPatch) ToNetworkPolicyPatchOutputWithContext(ctx context.Context) NetworkPolicyPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyPatchOutput)
}

// NetworkPolicyPatchArrayInput is an input type that accepts NetworkPolicyPatchArray and NetworkPolicyPatchArrayOutput values.
// You can construct a concrete instance of `NetworkPolicyPatchArrayInput` via:
//
//          NetworkPolicyPatchArray{ NetworkPolicyPatchArgs{...} }
type NetworkPolicyPatchArrayInput interface {
	pulumi.Input

	ToNetworkPolicyPatchArrayOutput() NetworkPolicyPatchArrayOutput
	ToNetworkPolicyPatchArrayOutputWithContext(context.Context) NetworkPolicyPatchArrayOutput
}

type NetworkPolicyPatchArray []NetworkPolicyPatchInput

func (NetworkPolicyPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPolicyPatch)(nil)).Elem()
}

func (i NetworkPolicyPatchArray) ToNetworkPolicyPatchArrayOutput() NetworkPolicyPatchArrayOutput {
	return i.ToNetworkPolicyPatchArrayOutputWithContext(context.Background())
}

func (i NetworkPolicyPatchArray) ToNetworkPolicyPatchArrayOutputWithContext(ctx context.Context) NetworkPolicyPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyPatchArrayOutput)
}

// NetworkPolicyPatchMapInput is an input type that accepts NetworkPolicyPatchMap and NetworkPolicyPatchMapOutput values.
// You can construct a concrete instance of `NetworkPolicyPatchMapInput` via:
//
//          NetworkPolicyPatchMap{ "key": NetworkPolicyPatchArgs{...} }
type NetworkPolicyPatchMapInput interface {
	pulumi.Input

	ToNetworkPolicyPatchMapOutput() NetworkPolicyPatchMapOutput
	ToNetworkPolicyPatchMapOutputWithContext(context.Context) NetworkPolicyPatchMapOutput
}

type NetworkPolicyPatchMap map[string]NetworkPolicyPatchInput

func (NetworkPolicyPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPolicyPatch)(nil)).Elem()
}

func (i NetworkPolicyPatchMap) ToNetworkPolicyPatchMapOutput() NetworkPolicyPatchMapOutput {
	return i.ToNetworkPolicyPatchMapOutputWithContext(context.Background())
}

func (i NetworkPolicyPatchMap) ToNetworkPolicyPatchMapOutputWithContext(ctx context.Context) NetworkPolicyPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyPatchMapOutput)
}

type NetworkPolicyPatchOutput struct{ *pulumi.OutputState }

func (NetworkPolicyPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicyPatch)(nil)).Elem()
}

func (o NetworkPolicyPatchOutput) ToNetworkPolicyPatchOutput() NetworkPolicyPatchOutput {
	return o
}

func (o NetworkPolicyPatchOutput) ToNetworkPolicyPatchOutputWithContext(ctx context.Context) NetworkPolicyPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o NetworkPolicyPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkPolicyPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o NetworkPolicyPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkPolicyPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o NetworkPolicyPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *NetworkPolicyPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Specification of the desired behavior for this NetworkPolicy.
func (o NetworkPolicyPatchOutput) Spec() NetworkPolicySpecPatchPtrOutput {
	return o.ApplyT(func(v *NetworkPolicyPatch) NetworkPolicySpecPatchPtrOutput { return v.Spec }).(NetworkPolicySpecPatchPtrOutput)
}

type NetworkPolicyPatchArrayOutput struct{ *pulumi.OutputState }

func (NetworkPolicyPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPolicyPatch)(nil)).Elem()
}

func (o NetworkPolicyPatchArrayOutput) ToNetworkPolicyPatchArrayOutput() NetworkPolicyPatchArrayOutput {
	return o
}

func (o NetworkPolicyPatchArrayOutput) ToNetworkPolicyPatchArrayOutputWithContext(ctx context.Context) NetworkPolicyPatchArrayOutput {
	return o
}

func (o NetworkPolicyPatchArrayOutput) Index(i pulumi.IntInput) NetworkPolicyPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkPolicyPatch {
		return vs[0].([]*NetworkPolicyPatch)[vs[1].(int)]
	}).(NetworkPolicyPatchOutput)
}

type NetworkPolicyPatchMapOutput struct{ *pulumi.OutputState }

func (NetworkPolicyPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPolicyPatch)(nil)).Elem()
}

func (o NetworkPolicyPatchMapOutput) ToNetworkPolicyPatchMapOutput() NetworkPolicyPatchMapOutput {
	return o
}

func (o NetworkPolicyPatchMapOutput) ToNetworkPolicyPatchMapOutputWithContext(ctx context.Context) NetworkPolicyPatchMapOutput {
	return o
}

func (o NetworkPolicyPatchMapOutput) MapIndex(k pulumi.StringInput) NetworkPolicyPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkPolicyPatch {
		return vs[0].(map[string]*NetworkPolicyPatch)[vs[1].(string)]
	}).(NetworkPolicyPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyPatchInput)(nil)).Elem(), &NetworkPolicyPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyPatchArrayInput)(nil)).Elem(), NetworkPolicyPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyPatchMapInput)(nil)).Elem(), NetworkPolicyPatchMap{})
	pulumi.RegisterOutputType(NetworkPolicyPatchOutput{})
	pulumi.RegisterOutputType(NetworkPolicyPatchArrayOutput{})
	pulumi.RegisterOutputType(NetworkPolicyPatchMapOutput{})
}
