// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.authentication.v1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.authentication.v1.TokenRequestArgs;
import com.pulumi.kubernetes.authentication.v1.outputs.TokenRequestSpec;
import com.pulumi.kubernetes.authentication.v1.outputs.TokenRequestStatus;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * TokenRequest requests a token for a given service account.
 * 
 */
@ResourceType(type="kubernetes:authentication.k8s.io/v1:TokenRequest")
public class TokenRequest extends com.pulumi.resources.CustomResource {
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
     * Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMeta.class}, tree="[0]")
    private Output</* @Nullable */ ObjectMeta> metadata;

    /**
     * @return Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<Optional<ObjectMeta>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * Spec holds information about the request being evaluated
     * 
     */
    @Export(name="spec", refs={TokenRequestSpec.class}, tree="[0]")
    private Output<TokenRequestSpec> spec;

    /**
     * @return Spec holds information about the request being evaluated
     * 
     */
    public Output<TokenRequestSpec> spec() {
        return this.spec;
    }
    /**
     * Status is filled in by the server and indicates whether the token can be authenticated.
     * 
     */
    @Export(name="status", refs={TokenRequestStatus.class}, tree="[0]")
    private Output</* @Nullable */ TokenRequestStatus> status;

    /**
     * @return Status is filled in by the server and indicates whether the token can be authenticated.
     * 
     */
    public Output<Optional<TokenRequestStatus>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TokenRequest(String name) {
        this(name, TokenRequestArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TokenRequest(String name, TokenRequestArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TokenRequest(String name, TokenRequestArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:authentication.k8s.io/v1:TokenRequest", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private TokenRequest(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:authentication.k8s.io/v1:TokenRequest", name, null, makeResourceOptions(options, id));
    }

    private static TokenRequestArgs makeArgs(TokenRequestArgs args) {
        var builder = args == null ? TokenRequestArgs.builder() : TokenRequestArgs.builder(args);
        return builder
            .apiVersion("authentication.k8s.io/v1")
            .kind("TokenRequest")
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
    public static TokenRequest get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TokenRequest(name, id, options);
    }
}
