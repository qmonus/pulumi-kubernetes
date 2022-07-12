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

// ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a principal that can be authenticated and authorized * a set of secrets
type ServiceAccountPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	AutomountServiceAccountToken pulumi.BoolPtrOutput `pulumi:"automountServiceAccountToken"`
	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets LocalObjectReferencePatchArrayOutput `pulumi:"imagePullSecrets"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Secrets ObjectReferencePatchArrayOutput `pulumi:"secrets"`
}

// NewServiceAccountPatch registers a new resource with the given unique name, arguments, and options.
func NewServiceAccountPatch(ctx *pulumi.Context,
	name string, args *ServiceAccountPatchArgs, opts ...pulumi.ResourceOption) (*ServiceAccountPatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ServiceAccount")
	var resource ServiceAccountPatch
	err := ctx.RegisterResource("kubernetes:core/v1:ServiceAccountPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccountPatch gets an existing ServiceAccountPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccountPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountPatchState, opts ...pulumi.ResourceOption) (*ServiceAccountPatch, error) {
	var resource ServiceAccountPatch
	err := ctx.ReadResource("kubernetes:core/v1:ServiceAccountPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccountPatch resources.
type serviceAccountPatchState struct {
}

type ServiceAccountPatchState struct {
}

func (ServiceAccountPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountPatchState)(nil)).Elem()
}

type serviceAccountPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	AutomountServiceAccountToken *bool `pulumi:"automountServiceAccountToken"`
	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []LocalObjectReferencePatch `pulumi:"imagePullSecrets"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Secrets []ObjectReferencePatch `pulumi:"secrets"`
}

// The set of arguments for constructing a ServiceAccountPatch resource.
type ServiceAccountPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	AutomountServiceAccountToken pulumi.BoolPtrInput
	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets LocalObjectReferencePatchArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchInput
	// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Secrets ObjectReferencePatchArrayInput
}

func (ServiceAccountPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountPatchArgs)(nil)).Elem()
}

type ServiceAccountPatchInput interface {
	pulumi.Input

	ToServiceAccountPatchOutput() ServiceAccountPatchOutput
	ToServiceAccountPatchOutputWithContext(ctx context.Context) ServiceAccountPatchOutput
}

func (*ServiceAccountPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccountPatch)(nil)).Elem()
}

func (i *ServiceAccountPatch) ToServiceAccountPatchOutput() ServiceAccountPatchOutput {
	return i.ToServiceAccountPatchOutputWithContext(context.Background())
}

func (i *ServiceAccountPatch) ToServiceAccountPatchOutputWithContext(ctx context.Context) ServiceAccountPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountPatchOutput)
}

// ServiceAccountPatchArrayInput is an input type that accepts ServiceAccountPatchArray and ServiceAccountPatchArrayOutput values.
// You can construct a concrete instance of `ServiceAccountPatchArrayInput` via:
//
//          ServiceAccountPatchArray{ ServiceAccountPatchArgs{...} }
type ServiceAccountPatchArrayInput interface {
	pulumi.Input

	ToServiceAccountPatchArrayOutput() ServiceAccountPatchArrayOutput
	ToServiceAccountPatchArrayOutputWithContext(context.Context) ServiceAccountPatchArrayOutput
}

type ServiceAccountPatchArray []ServiceAccountPatchInput

func (ServiceAccountPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccountPatch)(nil)).Elem()
}

func (i ServiceAccountPatchArray) ToServiceAccountPatchArrayOutput() ServiceAccountPatchArrayOutput {
	return i.ToServiceAccountPatchArrayOutputWithContext(context.Background())
}

func (i ServiceAccountPatchArray) ToServiceAccountPatchArrayOutputWithContext(ctx context.Context) ServiceAccountPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountPatchArrayOutput)
}

// ServiceAccountPatchMapInput is an input type that accepts ServiceAccountPatchMap and ServiceAccountPatchMapOutput values.
// You can construct a concrete instance of `ServiceAccountPatchMapInput` via:
//
//          ServiceAccountPatchMap{ "key": ServiceAccountPatchArgs{...} }
type ServiceAccountPatchMapInput interface {
	pulumi.Input

	ToServiceAccountPatchMapOutput() ServiceAccountPatchMapOutput
	ToServiceAccountPatchMapOutputWithContext(context.Context) ServiceAccountPatchMapOutput
}

type ServiceAccountPatchMap map[string]ServiceAccountPatchInput

func (ServiceAccountPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccountPatch)(nil)).Elem()
}

func (i ServiceAccountPatchMap) ToServiceAccountPatchMapOutput() ServiceAccountPatchMapOutput {
	return i.ToServiceAccountPatchMapOutputWithContext(context.Background())
}

func (i ServiceAccountPatchMap) ToServiceAccountPatchMapOutputWithContext(ctx context.Context) ServiceAccountPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountPatchMapOutput)
}

type ServiceAccountPatchOutput struct{ *pulumi.OutputState }

func (ServiceAccountPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccountPatch)(nil)).Elem()
}

func (o ServiceAccountPatchOutput) ToServiceAccountPatchOutput() ServiceAccountPatchOutput {
	return o
}

func (o ServiceAccountPatchOutput) ToServiceAccountPatchOutputWithContext(ctx context.Context) ServiceAccountPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ServiceAccountPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAccountPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
func (o ServiceAccountPatchOutput) AutomountServiceAccountToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAccountPatch) pulumi.BoolPtrOutput { return v.AutomountServiceAccountToken }).(pulumi.BoolPtrOutput)
}

// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
func (o ServiceAccountPatchOutput) ImagePullSecrets() LocalObjectReferencePatchArrayOutput {
	return o.ApplyT(func(v *ServiceAccountPatch) LocalObjectReferencePatchArrayOutput { return v.ImagePullSecrets }).(LocalObjectReferencePatchArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ServiceAccountPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAccountPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ServiceAccountPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ServiceAccountPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
func (o ServiceAccountPatchOutput) Secrets() ObjectReferencePatchArrayOutput {
	return o.ApplyT(func(v *ServiceAccountPatch) ObjectReferencePatchArrayOutput { return v.Secrets }).(ObjectReferencePatchArrayOutput)
}

type ServiceAccountPatchArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccountPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccountPatch)(nil)).Elem()
}

func (o ServiceAccountPatchArrayOutput) ToServiceAccountPatchArrayOutput() ServiceAccountPatchArrayOutput {
	return o
}

func (o ServiceAccountPatchArrayOutput) ToServiceAccountPatchArrayOutputWithContext(ctx context.Context) ServiceAccountPatchArrayOutput {
	return o
}

func (o ServiceAccountPatchArrayOutput) Index(i pulumi.IntInput) ServiceAccountPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceAccountPatch {
		return vs[0].([]*ServiceAccountPatch)[vs[1].(int)]
	}).(ServiceAccountPatchOutput)
}

type ServiceAccountPatchMapOutput struct{ *pulumi.OutputState }

func (ServiceAccountPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccountPatch)(nil)).Elem()
}

func (o ServiceAccountPatchMapOutput) ToServiceAccountPatchMapOutput() ServiceAccountPatchMapOutput {
	return o
}

func (o ServiceAccountPatchMapOutput) ToServiceAccountPatchMapOutputWithContext(ctx context.Context) ServiceAccountPatchMapOutput {
	return o
}

func (o ServiceAccountPatchMapOutput) MapIndex(k pulumi.StringInput) ServiceAccountPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceAccountPatch {
		return vs[0].(map[string]*ServiceAccountPatch)[vs[1].(string)]
	}).(ServiceAccountPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountPatchInput)(nil)).Elem(), &ServiceAccountPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountPatchArrayInput)(nil)).Elem(), ServiceAccountPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountPatchMapInput)(nil)).Elem(), ServiceAccountPatchMap{})
	pulumi.RegisterOutputType(ServiceAccountPatchOutput{})
	pulumi.RegisterOutputType(ServiceAccountPatchArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccountPatchMapOutput{})
}
