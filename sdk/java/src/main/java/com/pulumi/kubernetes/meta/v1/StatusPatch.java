// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.meta.v1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.meta.v1.StatusPatchArgs;
import com.pulumi.kubernetes.meta.v1.outputs.ListMetaPatch;
import com.pulumi.kubernetes.meta.v1.outputs.StatusDetailsPatch;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Patch resources are used to modify existing Kubernetes resources by using
 * Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
 * one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
 * Conflicts will result in an error by default, but can be forced using the &#34;pulumi.com/patchForce&#34; annotation. See the
 * [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
 * additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
 * Status is a return value for calls that don&#39;t return other objects.
 * 
 */
@ResourceType(type="kubernetes:meta/v1:StatusPatch")
public class StatusPatch extends com.pulumi.resources.CustomResource {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Export(name="apiVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Output<Optional<String>> apiVersion() {
        return Codegen.optional(this.apiVersion);
    }
    /**
     * Suggested HTTP return code for this status, 0 if not set.
     * 
     */
    @Export(name="code", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> code;

    /**
     * @return Suggested HTTP return code for this status, 0 if not set.
     * 
     */
    public Output<Optional<Integer>> code() {
        return Codegen.optional(this.code);
    }
    /**
     * Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
     * 
     */
    @Export(name="details", refs={StatusDetailsPatch.class}, tree="[0]")
    private Output</* @Nullable */ StatusDetailsPatch> details;

    /**
     * @return Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
     * 
     */
    public Output<Optional<StatusDetailsPatch>> details() {
        return Codegen.optional(this.details);
    }
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<Optional<String>> kind() {
        return Codegen.optional(this.kind);
    }
    /**
     * A human-readable description of the status of this operation.
     * 
     */
    @Export(name="message", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> message;

    /**
     * @return A human-readable description of the status of this operation.
     * 
     */
    public Output<Optional<String>> message() {
        return Codegen.optional(this.message);
    }
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="metadata", refs={ListMetaPatch.class}, tree="[0]")
    private Output</* @Nullable */ ListMetaPatch> metadata;

    /**
     * @return Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<Optional<ListMetaPatch>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * A machine-readable description of why this operation is in the &#34;Failure&#34; status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
     * 
     */
    @Export(name="reason", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> reason;

    /**
     * @return A machine-readable description of why this operation is in the &#34;Failure&#34; status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
     * 
     */
    public Output<Optional<String>> reason() {
        return Codegen.optional(this.reason);
    }
    /**
     * Status of the operation. One of: &#34;Success&#34; or &#34;Failure&#34;. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    /**
     * @return Status of the operation. One of: &#34;Success&#34; or &#34;Failure&#34;. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StatusPatch(String name) {
        this(name, StatusPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StatusPatch(String name, @Nullable StatusPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StatusPatch(String name, @Nullable StatusPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:meta/v1:StatusPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private StatusPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:meta/v1:StatusPatch", name, null, makeResourceOptions(options, id));
    }

    private static StatusPatchArgs makeArgs(@Nullable StatusPatchArgs args) {
        var builder = args == null ? StatusPatchArgs.builder() : StatusPatchArgs.builder(args);
        return builder
            .apiVersion("v1")
            .kind("Status")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static StatusPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new StatusPatch(name, id, options);
    }
}
