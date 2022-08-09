// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.core.v1.inputs.LocalObjectReferencePatchArgs;
import com.pulumi.kubernetes.core.v1.inputs.ObjectReferencePatchArgs;
import com.pulumi.kubernetes.meta.v1.inputs.ObjectMetaPatchArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceAccountPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceAccountPatchArgs Empty = new ServiceAccountPatchArgs();

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Import(name="apiVersion")
    private @Nullable Output<String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Optional<Output<String>> apiVersion() {
        return Optional.ofNullable(this.apiVersion);
    }

    /**
     * AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
     * 
     */
    @Import(name="automountServiceAccountToken")
    private @Nullable Output<Boolean> automountServiceAccountToken;

    /**
     * @return AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
     * 
     */
    public Optional<Output<Boolean>> automountServiceAccountToken() {
        return Optional.ofNullable(this.automountServiceAccountToken);
    }

    /**
     * ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
     * 
     */
    @Import(name="imagePullSecrets")
    private @Nullable Output<List<LocalObjectReferencePatchArgs>> imagePullSecrets;

    /**
     * @return ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
     * 
     */
    public Optional<Output<List<LocalObjectReferencePatchArgs>>> imagePullSecrets() {
        return Optional.ofNullable(this.imagePullSecrets);
    }

    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Import(name="kind")
    private @Nullable Output<String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Optional<Output<String>> kind() {
        return Optional.ofNullable(this.kind);
    }

    /**
     * Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<ObjectMetaPatchArgs> metadata;

    /**
     * @return Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Optional<Output<ObjectMetaPatchArgs>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a &#34;kubernetes.io/enforce-mountable-secrets&#34; annotation set to &#34;true&#34;. This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
     * 
     */
    @Import(name="secrets")
    private @Nullable Output<List<ObjectReferencePatchArgs>> secrets;

    /**
     * @return Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a &#34;kubernetes.io/enforce-mountable-secrets&#34; annotation set to &#34;true&#34;. This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
     * 
     */
    public Optional<Output<List<ObjectReferencePatchArgs>>> secrets() {
        return Optional.ofNullable(this.secrets);
    }

    private ServiceAccountPatchArgs() {}

    private ServiceAccountPatchArgs(ServiceAccountPatchArgs $) {
        this.apiVersion = $.apiVersion;
        this.automountServiceAccountToken = $.automountServiceAccountToken;
        this.imagePullSecrets = $.imagePullSecrets;
        this.kind = $.kind;
        this.metadata = $.metadata;
        this.secrets = $.secrets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceAccountPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceAccountPatchArgs $;

        public Builder() {
            $ = new ServiceAccountPatchArgs();
        }

        public Builder(ServiceAccountPatchArgs defaults) {
            $ = new ServiceAccountPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiVersion APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
         * 
         * @return builder
         * 
         */
        public Builder apiVersion(@Nullable Output<String> apiVersion) {
            $.apiVersion = apiVersion;
            return this;
        }

        /**
         * @param apiVersion APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
         * 
         * @return builder
         * 
         */
        public Builder apiVersion(String apiVersion) {
            return apiVersion(Output.of(apiVersion));
        }

        /**
         * @param automountServiceAccountToken AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
         * 
         * @return builder
         * 
         */
        public Builder automountServiceAccountToken(@Nullable Output<Boolean> automountServiceAccountToken) {
            $.automountServiceAccountToken = automountServiceAccountToken;
            return this;
        }

        /**
         * @param automountServiceAccountToken AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
         * 
         * @return builder
         * 
         */
        public Builder automountServiceAccountToken(Boolean automountServiceAccountToken) {
            return automountServiceAccountToken(Output.of(automountServiceAccountToken));
        }

        /**
         * @param imagePullSecrets ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
         * 
         * @return builder
         * 
         */
        public Builder imagePullSecrets(@Nullable Output<List<LocalObjectReferencePatchArgs>> imagePullSecrets) {
            $.imagePullSecrets = imagePullSecrets;
            return this;
        }

        /**
         * @param imagePullSecrets ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
         * 
         * @return builder
         * 
         */
        public Builder imagePullSecrets(List<LocalObjectReferencePatchArgs> imagePullSecrets) {
            return imagePullSecrets(Output.of(imagePullSecrets));
        }

        /**
         * @param imagePullSecrets ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
         * 
         * @return builder
         * 
         */
        public Builder imagePullSecrets(LocalObjectReferencePatchArgs... imagePullSecrets) {
            return imagePullSecrets(List.of(imagePullSecrets));
        }

        /**
         * @param kind Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
         * 
         * @return builder
         * 
         */
        public Builder kind(@Nullable Output<String> kind) {
            $.kind = kind;
            return this;
        }

        /**
         * @param kind Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
         * 
         * @return builder
         * 
         */
        public Builder kind(String kind) {
            return kind(Output.of(kind));
        }

        /**
         * @param metadata Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<ObjectMetaPatchArgs> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
         * 
         * @return builder
         * 
         */
        public Builder metadata(ObjectMetaPatchArgs metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param secrets Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a &#34;kubernetes.io/enforce-mountable-secrets&#34; annotation set to &#34;true&#34;. This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
         * 
         * @return builder
         * 
         */
        public Builder secrets(@Nullable Output<List<ObjectReferencePatchArgs>> secrets) {
            $.secrets = secrets;
            return this;
        }

        /**
         * @param secrets Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a &#34;kubernetes.io/enforce-mountable-secrets&#34; annotation set to &#34;true&#34;. This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
         * 
         * @return builder
         * 
         */
        public Builder secrets(List<ObjectReferencePatchArgs> secrets) {
            return secrets(Output.of(secrets));
        }

        /**
         * @param secrets Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a &#34;kubernetes.io/enforce-mountable-secrets&#34; annotation set to &#34;true&#34;. This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
         * 
         * @return builder
         * 
         */
        public Builder secrets(ObjectReferencePatchArgs... secrets) {
            return secrets(List.of(secrets));
        }

        public ServiceAccountPatchArgs build() {
            $.apiVersion = Codegen.stringProp("apiVersion").output().arg($.apiVersion).getNullable();
            $.kind = Codegen.stringProp("kind").output().arg($.kind).getNullable();
            return $;
        }
    }

}
