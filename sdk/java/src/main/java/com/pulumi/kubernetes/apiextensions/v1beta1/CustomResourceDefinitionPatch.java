// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.apiextensions.v1beta1;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.apiextensions.v1beta1.CustomResourceDefinitionPatchArgs;
import com.pulumi.kubernetes.apiextensions.v1beta1.outputs.CustomResourceDefinitionSpecPatch;
import com.pulumi.kubernetes.apiextensions.v1beta1.outputs.CustomResourceDefinitionStatusPatch;
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
 * CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format &lt;.spec.name&gt;.&lt;.spec.group&gt;. Deprecated in v1.16, planned for removal in v1.19. Use apiextensions.k8s.io/v1 CustomResourceDefinition instead.
 * 
 */
@ResourceType(type="kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinitionPatch")
public class CustomResourceDefinitionPatch extends com.pulumi.resources.CustomResource {
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
     * spec describes how the user wants the resources to appear
     * 
     */
    @Export(name="spec", type=CustomResourceDefinitionSpecPatch.class, parameters={})
    private Output</* @Nullable */ CustomResourceDefinitionSpecPatch> spec;

    /**
     * @return spec describes how the user wants the resources to appear
     * 
     */
    public Output<Optional<CustomResourceDefinitionSpecPatch>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * status indicates the actual state of the CustomResourceDefinition
     * 
     */
    @Export(name="status", type=CustomResourceDefinitionStatusPatch.class, parameters={})
    private Output</* @Nullable */ CustomResourceDefinitionStatusPatch> status;

    /**
     * @return status indicates the actual state of the CustomResourceDefinition
     * 
     */
    public Output<Optional<CustomResourceDefinitionStatusPatch>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomResourceDefinitionPatch(String name) {
        this(name, CustomResourceDefinitionPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomResourceDefinitionPatch(String name, @Nullable CustomResourceDefinitionPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomResourceDefinitionPatch(String name, @Nullable CustomResourceDefinitionPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinitionPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private CustomResourceDefinitionPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinitionPatch", name, null, makeResourceOptions(options, id));
    }

    private static CustomResourceDefinitionPatchArgs makeArgs(@Nullable CustomResourceDefinitionPatchArgs args) {
        var builder = args == null ? CustomResourceDefinitionPatchArgs.builder() : CustomResourceDefinitionPatchArgs.builder(args);
        return builder
            .apiVersion("apiextensions.k8s.io/v1beta1")
            .kind("CustomResourceDefinition")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinitionPatch").build())
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
    public static CustomResourceDefinitionPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomResourceDefinitionPatch(name, id, options);
    }
}
