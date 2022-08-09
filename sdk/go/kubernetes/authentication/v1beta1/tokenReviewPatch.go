// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
// TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.
type TokenReviewPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput          `pulumi:"kind"`
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec TokenReviewSpecPatchPtrOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request can be authenticated.
	Status TokenReviewStatusPatchPtrOutput `pulumi:"status"`
}

// NewTokenReviewPatch registers a new resource with the given unique name, arguments, and options.
func NewTokenReviewPatch(ctx *pulumi.Context,
	name string, args *TokenReviewPatchArgs, opts ...pulumi.ResourceOption) (*TokenReviewPatch, error) {
	if args == nil {
		args = &TokenReviewPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("authentication.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("TokenReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authentication.k8s.io/v1:TokenReviewPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource TokenReviewPatch
	err := ctx.RegisterResource("kubernetes:authentication.k8s.io/v1beta1:TokenReviewPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTokenReviewPatch gets an existing TokenReviewPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTokenReviewPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenReviewPatchState, opts ...pulumi.ResourceOption) (*TokenReviewPatch, error) {
	var resource TokenReviewPatch
	err := ctx.ReadResource("kubernetes:authentication.k8s.io/v1beta1:TokenReviewPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TokenReviewPatch resources.
type tokenReviewPatchState struct {
}

type TokenReviewPatchState struct {
}

func (TokenReviewPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenReviewPatchState)(nil)).Elem()
}

type tokenReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string                 `pulumi:"kind"`
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec *TokenReviewSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a TokenReviewPatch resource.
type TokenReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec holds information about the request being evaluated
	Spec TokenReviewSpecPatchPtrInput
}

func (TokenReviewPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenReviewPatchArgs)(nil)).Elem()
}

type TokenReviewPatchInput interface {
	pulumi.Input

	ToTokenReviewPatchOutput() TokenReviewPatchOutput
	ToTokenReviewPatchOutputWithContext(ctx context.Context) TokenReviewPatchOutput
}

func (*TokenReviewPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenReviewPatch)(nil)).Elem()
}

func (i *TokenReviewPatch) ToTokenReviewPatchOutput() TokenReviewPatchOutput {
	return i.ToTokenReviewPatchOutputWithContext(context.Background())
}

func (i *TokenReviewPatch) ToTokenReviewPatchOutputWithContext(ctx context.Context) TokenReviewPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewPatchOutput)
}

// TokenReviewPatchArrayInput is an input type that accepts TokenReviewPatchArray and TokenReviewPatchArrayOutput values.
// You can construct a concrete instance of `TokenReviewPatchArrayInput` via:
//
//          TokenReviewPatchArray{ TokenReviewPatchArgs{...} }
type TokenReviewPatchArrayInput interface {
	pulumi.Input

	ToTokenReviewPatchArrayOutput() TokenReviewPatchArrayOutput
	ToTokenReviewPatchArrayOutputWithContext(context.Context) TokenReviewPatchArrayOutput
}

type TokenReviewPatchArray []TokenReviewPatchInput

func (TokenReviewPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TokenReviewPatch)(nil)).Elem()
}

func (i TokenReviewPatchArray) ToTokenReviewPatchArrayOutput() TokenReviewPatchArrayOutput {
	return i.ToTokenReviewPatchArrayOutputWithContext(context.Background())
}

func (i TokenReviewPatchArray) ToTokenReviewPatchArrayOutputWithContext(ctx context.Context) TokenReviewPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewPatchArrayOutput)
}

// TokenReviewPatchMapInput is an input type that accepts TokenReviewPatchMap and TokenReviewPatchMapOutput values.
// You can construct a concrete instance of `TokenReviewPatchMapInput` via:
//
//          TokenReviewPatchMap{ "key": TokenReviewPatchArgs{...} }
type TokenReviewPatchMapInput interface {
	pulumi.Input

	ToTokenReviewPatchMapOutput() TokenReviewPatchMapOutput
	ToTokenReviewPatchMapOutputWithContext(context.Context) TokenReviewPatchMapOutput
}

type TokenReviewPatchMap map[string]TokenReviewPatchInput

func (TokenReviewPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TokenReviewPatch)(nil)).Elem()
}

func (i TokenReviewPatchMap) ToTokenReviewPatchMapOutput() TokenReviewPatchMapOutput {
	return i.ToTokenReviewPatchMapOutputWithContext(context.Background())
}

func (i TokenReviewPatchMap) ToTokenReviewPatchMapOutputWithContext(ctx context.Context) TokenReviewPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewPatchMapOutput)
}

type TokenReviewPatchOutput struct{ *pulumi.OutputState }

func (TokenReviewPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenReviewPatch)(nil)).Elem()
}

func (o TokenReviewPatchOutput) ToTokenReviewPatchOutput() TokenReviewPatchOutput {
	return o
}

func (o TokenReviewPatchOutput) ToTokenReviewPatchOutputWithContext(ctx context.Context) TokenReviewPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o TokenReviewPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TokenReviewPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o TokenReviewPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TokenReviewPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o TokenReviewPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *TokenReviewPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec holds information about the request being evaluated
func (o TokenReviewPatchOutput) Spec() TokenReviewSpecPatchPtrOutput {
	return o.ApplyT(func(v *TokenReviewPatch) TokenReviewSpecPatchPtrOutput { return v.Spec }).(TokenReviewSpecPatchPtrOutput)
}

// Status is filled in by the server and indicates whether the request can be authenticated.
func (o TokenReviewPatchOutput) Status() TokenReviewStatusPatchPtrOutput {
	return o.ApplyT(func(v *TokenReviewPatch) TokenReviewStatusPatchPtrOutput { return v.Status }).(TokenReviewStatusPatchPtrOutput)
}

type TokenReviewPatchArrayOutput struct{ *pulumi.OutputState }

func (TokenReviewPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TokenReviewPatch)(nil)).Elem()
}

func (o TokenReviewPatchArrayOutput) ToTokenReviewPatchArrayOutput() TokenReviewPatchArrayOutput {
	return o
}

func (o TokenReviewPatchArrayOutput) ToTokenReviewPatchArrayOutputWithContext(ctx context.Context) TokenReviewPatchArrayOutput {
	return o
}

func (o TokenReviewPatchArrayOutput) Index(i pulumi.IntInput) TokenReviewPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TokenReviewPatch {
		return vs[0].([]*TokenReviewPatch)[vs[1].(int)]
	}).(TokenReviewPatchOutput)
}

type TokenReviewPatchMapOutput struct{ *pulumi.OutputState }

func (TokenReviewPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TokenReviewPatch)(nil)).Elem()
}

func (o TokenReviewPatchMapOutput) ToTokenReviewPatchMapOutput() TokenReviewPatchMapOutput {
	return o
}

func (o TokenReviewPatchMapOutput) ToTokenReviewPatchMapOutputWithContext(ctx context.Context) TokenReviewPatchMapOutput {
	return o
}

func (o TokenReviewPatchMapOutput) MapIndex(k pulumi.StringInput) TokenReviewPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TokenReviewPatch {
		return vs[0].(map[string]*TokenReviewPatch)[vs[1].(string)]
	}).(TokenReviewPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TokenReviewPatchInput)(nil)).Elem(), &TokenReviewPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenReviewPatchArrayInput)(nil)).Elem(), TokenReviewPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenReviewPatchMapInput)(nil)).Elem(), TokenReviewPatchMap{})
	pulumi.RegisterOutputType(TokenReviewPatchOutput{})
	pulumi.RegisterOutputType(TokenReviewPatchArrayOutput{})
	pulumi.RegisterOutputType(TokenReviewPatchMapOutput{})
}
