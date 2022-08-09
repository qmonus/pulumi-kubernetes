// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.authorization.v1beta1;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.authorization.v1beta1.SelfSubjectAccessReviewPatchArgs;
import com.pulumi.kubernetes.authorization.v1beta1.outputs.SelfSubjectAccessReviewSpecPatch;
import com.pulumi.kubernetes.authorization.v1beta1.outputs.SubjectAccessReviewStatusPatch;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMetaPatch;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Patch resources are used to modify existing Kubernetes resources by using
 * Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
 * one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
 * Conflicts will result in an error by default, but can be forced using the &#34;pulumi.com/patchForce&#34; annotation. See the
 * [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/installation-configuration/#server-side-apply) for
 * additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
 * SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means &#34;in all namespaces&#34;.  Self is a special case, because users should always be able to check whether they can perform an action
 * 
 */
@ResourceType(type="kubernetes:authorization.k8s.io/v1beta1:SelfSubjectAccessReviewPatch")
public class SelfSubjectAccessReviewPatch extends com.pulumi.resources.CustomResource {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Export(name="apiVersion", type=String.class, parameters={})
    private Output</* @Nullable */ String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Output<Optional<String>> apiVersion() {
        return Codegen.optional(this.apiVersion);
    }
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="kind", type=String.class, parameters={})
    private Output</* @Nullable */ String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<Optional<String>> kind() {
        return Codegen.optional(this.kind);
    }
    @Export(name="metadata", type=ObjectMetaPatch.class, parameters={})
    private Output</* @Nullable */ ObjectMetaPatch> metadata;

    public Output<Optional<ObjectMetaPatch>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * Spec holds information about the request being evaluated.  user and groups must be empty
     * 
     */
    @Export(name="spec", type=SelfSubjectAccessReviewSpecPatch.class, parameters={})
    private Output</* @Nullable */ SelfSubjectAccessReviewSpecPatch> spec;

    /**
     * @return Spec holds information about the request being evaluated.  user and groups must be empty
     * 
     */
    public Output<Optional<SelfSubjectAccessReviewSpecPatch>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * Status is filled in by the server and indicates whether the request is allowed or not
     * 
     */
    @Export(name="status", type=SubjectAccessReviewStatusPatch.class, parameters={})
    private Output</* @Nullable */ SubjectAccessReviewStatusPatch> status;

    /**
     * @return Status is filled in by the server and indicates whether the request is allowed or not
     * 
     */
    public Output<Optional<SubjectAccessReviewStatusPatch>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SelfSubjectAccessReviewPatch(String name) {
        this(name, SelfSubjectAccessReviewPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SelfSubjectAccessReviewPatch(String name, @Nullable SelfSubjectAccessReviewPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SelfSubjectAccessReviewPatch(String name, @Nullable SelfSubjectAccessReviewPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectAccessReviewPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private SelfSubjectAccessReviewPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectAccessReviewPatch", name, null, makeResourceOptions(options, id));
    }

    private static SelfSubjectAccessReviewPatchArgs makeArgs(@Nullable SelfSubjectAccessReviewPatchArgs args) {
        var builder = args == null ? SelfSubjectAccessReviewPatchArgs.builder() : SelfSubjectAccessReviewPatchArgs.builder(args);
        return builder
            .apiVersion("authorization.k8s.io/v1beta1")
            .kind("SelfSubjectAccessReview")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReviewPatch").build())
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static SelfSubjectAccessReviewPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SelfSubjectAccessReviewPatch(name, id, options);
    }
}
