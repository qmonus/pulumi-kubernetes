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

// ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBinding, and will no longer be served in v1.20.
type ClusterRoleBindingPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefPatchPtrOutput `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectPatchArrayOutput `pulumi:"subjects"`
}

// NewClusterRoleBindingPatch registers a new resource with the given unique name, arguments, and options.
func NewClusterRoleBindingPatch(ctx *pulumi.Context,
	name string, args *ClusterRoleBindingPatchArgs, opts ...pulumi.ResourceOption) (*ClusterRoleBindingPatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("ClusterRoleBinding")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBindingPatch"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBindingPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource ClusterRoleBindingPatch
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRoleBindingPatch gets an existing ClusterRoleBindingPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRoleBindingPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleBindingPatchState, opts ...pulumi.ResourceOption) (*ClusterRoleBindingPatch, error) {
	var resource ClusterRoleBindingPatch
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRoleBindingPatch resources.
type clusterRoleBindingPatchState struct {
}

type ClusterRoleBindingPatchState struct {
}

func (ClusterRoleBindingPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingPatchState)(nil)).Elem()
}

type clusterRoleBindingPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef *RoleRefPatch `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []SubjectPatch `pulumi:"subjects"`
}

// The set of arguments for constructing a ClusterRoleBindingPatch resource.
type ClusterRoleBindingPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchInput
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefPatchPtrInput
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectPatchArrayInput
}

func (ClusterRoleBindingPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingPatchArgs)(nil)).Elem()
}

type ClusterRoleBindingPatchInput interface {
	pulumi.Input

	ToClusterRoleBindingPatchOutput() ClusterRoleBindingPatchOutput
	ToClusterRoleBindingPatchOutputWithContext(ctx context.Context) ClusterRoleBindingPatchOutput
}

func (*ClusterRoleBindingPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleBindingPatch)(nil)).Elem()
}

func (i *ClusterRoleBindingPatch) ToClusterRoleBindingPatchOutput() ClusterRoleBindingPatchOutput {
	return i.ToClusterRoleBindingPatchOutputWithContext(context.Background())
}

func (i *ClusterRoleBindingPatch) ToClusterRoleBindingPatchOutputWithContext(ctx context.Context) ClusterRoleBindingPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleBindingPatchOutput)
}

// ClusterRoleBindingPatchArrayInput is an input type that accepts ClusterRoleBindingPatchArray and ClusterRoleBindingPatchArrayOutput values.
// You can construct a concrete instance of `ClusterRoleBindingPatchArrayInput` via:
//
//          ClusterRoleBindingPatchArray{ ClusterRoleBindingPatchArgs{...} }
type ClusterRoleBindingPatchArrayInput interface {
	pulumi.Input

	ToClusterRoleBindingPatchArrayOutput() ClusterRoleBindingPatchArrayOutput
	ToClusterRoleBindingPatchArrayOutputWithContext(context.Context) ClusterRoleBindingPatchArrayOutput
}

type ClusterRoleBindingPatchArray []ClusterRoleBindingPatchInput

func (ClusterRoleBindingPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRoleBindingPatch)(nil)).Elem()
}

func (i ClusterRoleBindingPatchArray) ToClusterRoleBindingPatchArrayOutput() ClusterRoleBindingPatchArrayOutput {
	return i.ToClusterRoleBindingPatchArrayOutputWithContext(context.Background())
}

func (i ClusterRoleBindingPatchArray) ToClusterRoleBindingPatchArrayOutputWithContext(ctx context.Context) ClusterRoleBindingPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleBindingPatchArrayOutput)
}

// ClusterRoleBindingPatchMapInput is an input type that accepts ClusterRoleBindingPatchMap and ClusterRoleBindingPatchMapOutput values.
// You can construct a concrete instance of `ClusterRoleBindingPatchMapInput` via:
//
//          ClusterRoleBindingPatchMap{ "key": ClusterRoleBindingPatchArgs{...} }
type ClusterRoleBindingPatchMapInput interface {
	pulumi.Input

	ToClusterRoleBindingPatchMapOutput() ClusterRoleBindingPatchMapOutput
	ToClusterRoleBindingPatchMapOutputWithContext(context.Context) ClusterRoleBindingPatchMapOutput
}

type ClusterRoleBindingPatchMap map[string]ClusterRoleBindingPatchInput

func (ClusterRoleBindingPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRoleBindingPatch)(nil)).Elem()
}

func (i ClusterRoleBindingPatchMap) ToClusterRoleBindingPatchMapOutput() ClusterRoleBindingPatchMapOutput {
	return i.ToClusterRoleBindingPatchMapOutputWithContext(context.Background())
}

func (i ClusterRoleBindingPatchMap) ToClusterRoleBindingPatchMapOutputWithContext(ctx context.Context) ClusterRoleBindingPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleBindingPatchMapOutput)
}

type ClusterRoleBindingPatchOutput struct{ *pulumi.OutputState }

func (ClusterRoleBindingPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleBindingPatch)(nil)).Elem()
}

func (o ClusterRoleBindingPatchOutput) ToClusterRoleBindingPatchOutput() ClusterRoleBindingPatchOutput {
	return o
}

func (o ClusterRoleBindingPatchOutput) ToClusterRoleBindingPatchOutputWithContext(ctx context.Context) ClusterRoleBindingPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterRoleBindingPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterRoleBindingPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterRoleBindingPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterRoleBindingPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata.
func (o ClusterRoleBindingPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ClusterRoleBindingPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
func (o ClusterRoleBindingPatchOutput) RoleRef() RoleRefPatchPtrOutput {
	return o.ApplyT(func(v *ClusterRoleBindingPatch) RoleRefPatchPtrOutput { return v.RoleRef }).(RoleRefPatchPtrOutput)
}

// Subjects holds references to the objects the role applies to.
func (o ClusterRoleBindingPatchOutput) Subjects() SubjectPatchArrayOutput {
	return o.ApplyT(func(v *ClusterRoleBindingPatch) SubjectPatchArrayOutput { return v.Subjects }).(SubjectPatchArrayOutput)
}

type ClusterRoleBindingPatchArrayOutput struct{ *pulumi.OutputState }

func (ClusterRoleBindingPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRoleBindingPatch)(nil)).Elem()
}

func (o ClusterRoleBindingPatchArrayOutput) ToClusterRoleBindingPatchArrayOutput() ClusterRoleBindingPatchArrayOutput {
	return o
}

func (o ClusterRoleBindingPatchArrayOutput) ToClusterRoleBindingPatchArrayOutputWithContext(ctx context.Context) ClusterRoleBindingPatchArrayOutput {
	return o
}

func (o ClusterRoleBindingPatchArrayOutput) Index(i pulumi.IntInput) ClusterRoleBindingPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterRoleBindingPatch {
		return vs[0].([]*ClusterRoleBindingPatch)[vs[1].(int)]
	}).(ClusterRoleBindingPatchOutput)
}

type ClusterRoleBindingPatchMapOutput struct{ *pulumi.OutputState }

func (ClusterRoleBindingPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRoleBindingPatch)(nil)).Elem()
}

func (o ClusterRoleBindingPatchMapOutput) ToClusterRoleBindingPatchMapOutput() ClusterRoleBindingPatchMapOutput {
	return o
}

func (o ClusterRoleBindingPatchMapOutput) ToClusterRoleBindingPatchMapOutputWithContext(ctx context.Context) ClusterRoleBindingPatchMapOutput {
	return o
}

func (o ClusterRoleBindingPatchMapOutput) MapIndex(k pulumi.StringInput) ClusterRoleBindingPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterRoleBindingPatch {
		return vs[0].(map[string]*ClusterRoleBindingPatch)[vs[1].(string)]
	}).(ClusterRoleBindingPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleBindingPatchInput)(nil)).Elem(), &ClusterRoleBindingPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleBindingPatchArrayInput)(nil)).Elem(), ClusterRoleBindingPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleBindingPatchMapInput)(nil)).Elem(), ClusterRoleBindingPatchMap{})
	pulumi.RegisterOutputType(ClusterRoleBindingPatchOutput{})
	pulumi.RegisterOutputType(ClusterRoleBindingPatchArrayOutput{})
	pulumi.RegisterOutputType(ClusterRoleBindingPatchMapOutput{})
}
