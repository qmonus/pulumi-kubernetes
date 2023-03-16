// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.storage.v1beta1;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.core.v1.outputs.TopologySelectorTermPatch;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMetaPatch;
import com.pulumi.kubernetes.storage.v1beta1.StorageClassPatchArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Patch resources are used to modify existing Kubernetes resources by using
 * Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
 * one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
 * Conflicts will result in an error by default, but can be forced using the &#34;pulumi.com/patchForce&#34; annotation. See the
 * [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
 * additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
 * StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
 * 
 * StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
 * 
 */
@ResourceType(type="kubernetes:storage.k8s.io/v1beta1:StorageClassPatch")
public class StorageClassPatch extends com.pulumi.resources.CustomResource {
    /**
     * AllowVolumeExpansion shows whether the storage class allow volume expand
     * 
     */
    @Export(name="allowVolumeExpansion", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowVolumeExpansion;

    /**
     * @return AllowVolumeExpansion shows whether the storage class allow volume expand
     * 
     */
    public Output<Optional<Boolean>> allowVolumeExpansion() {
        return Codegen.optional(this.allowVolumeExpansion);
    }
    /**
     * Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    @Export(name="allowedTopologies", refs={List.class,TopologySelectorTermPatch.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TopologySelectorTermPatch>> allowedTopologies;

    /**
     * @return Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    public Output<Optional<List<TopologySelectorTermPatch>>> allowedTopologies() {
        return Codegen.optional(this.allowedTopologies);
    }
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
    @Export(name="metadata", refs={ObjectMetaPatch.class}, tree="[0]")
    private Output</* @Nullable */ ObjectMetaPatch> metadata;

    /**
     * @return Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<Optional<ObjectMetaPatch>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount of the PVs will simply fail if one is invalid.
     * 
     */
    @Export(name="mountOptions", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> mountOptions;

    /**
     * @return Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount of the PVs will simply fail if one is invalid.
     * 
     */
    public Output<Optional<List<String>>> mountOptions() {
        return Codegen.optional(this.mountOptions);
    }
    /**
     * Parameters holds the parameters for the provisioner that should create volumes of this storage class.
     * 
     */
    @Export(name="parameters", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> parameters;

    /**
     * @return Parameters holds the parameters for the provisioner that should create volumes of this storage class.
     * 
     */
    public Output<Optional<Map<String,String>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * Provisioner indicates the type of the provisioner.
     * 
     */
    @Export(name="provisioner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> provisioner;

    /**
     * @return Provisioner indicates the type of the provisioner.
     * 
     */
    public Output<Optional<String>> provisioner() {
        return Codegen.optional(this.provisioner);
    }
    /**
     * Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
     * 
     */
    @Export(name="reclaimPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> reclaimPolicy;

    /**
     * @return Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
     * 
     */
    public Output<Optional<String>> reclaimPolicy() {
        return Codegen.optional(this.reclaimPolicy);
    }
    /**
     * VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    @Export(name="volumeBindingMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> volumeBindingMode;

    /**
     * @return VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    public Output<Optional<String>> volumeBindingMode() {
        return Codegen.optional(this.volumeBindingMode);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StorageClassPatch(String name) {
        this(name, StorageClassPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StorageClassPatch(String name, @Nullable StorageClassPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StorageClassPatch(String name, @Nullable StorageClassPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:storage.k8s.io/v1beta1:StorageClassPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private StorageClassPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:storage.k8s.io/v1beta1:StorageClassPatch", name, null, makeResourceOptions(options, id));
    }

    private static StorageClassPatchArgs makeArgs(@Nullable StorageClassPatchArgs args) {
        var builder = args == null ? StorageClassPatchArgs.builder() : StorageClassPatchArgs.builder(args);
        return builder
            .apiVersion("storage.k8s.io/v1beta1")
            .kind("StorageClass")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:storage.k8s.io/v1:StorageClassPatch").build())
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
    public static StorageClassPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new StorageClassPatch(name, id, options);
    }
}
