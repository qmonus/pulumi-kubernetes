// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Apps.V1
{
    /// <summary>
    /// Patch resources are used to modify existing Kubernetes resources by using
    /// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
    /// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
    /// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
    /// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/installation-configuration/#server-side-apply) for
    /// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
    /// Deployment enables declarative updates for Pods and ReplicaSets.
    /// 
    /// This resource waits until its status is ready before registering success
    /// for create/update, and populating output properties from the current state of the resource.
    /// The following conditions are used to determine whether the resource creation has
    /// succeeded or failed:
    /// 
    /// 1. The Deployment has begun to be updated by the Deployment controller. If the current
    ///    generation of the Deployment is &gt; 1, then this means that the current generation must
    ///    be different from the generation reported by the last outputs.
    /// 2. There exists a ReplicaSet whose revision is equal to the current revision of the
    ///    Deployment.
    /// 3. The Deployment's '.status.conditions' has a status of type 'Available' whose 'status'
    ///    member is set to 'True'.
    /// 4. If the Deployment has generation &gt; 1, then '.status.conditions' has a status of type
    ///    'Progressing', whose 'status' member is set to 'True', and whose 'reason' is
    ///    'NewReplicaSetAvailable'. For generation &lt;= 1, this status field does not exist,
    ///    because it doesn't do a rollout (i.e., it simply creates the Deployment and
    ///    corresponding ReplicaSet), and therefore there is no rollout to mark as 'Progressing'.
    /// 
    /// If the Deployment has not reached a Ready state after 10 minutes, it will
    /// time out and mark the resource update as Failed. You can override the default timeout value
    /// by setting the 'customTimeouts' option on the resource.
    /// </summary>
    [KubernetesResourceType("kubernetes:apps/v1:DeploymentPatch")]
    public partial class DeploymentPatch : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch> Metadata { get; private set; } = null!;

        /// <summary>
        /// Specification of the desired behavior of the Deployment.
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Apps.V1.DeploymentSpecPatch> Spec { get; private set; } = null!;

        /// <summary>
        /// Most recently observed status of the Deployment.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Apps.V1.DeploymentStatusPatch> Status { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentPatch(string name, Pulumi.Kubernetes.Types.Inputs.Apps.V1.DeploymentPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:apps/v1:DeploymentPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal DeploymentPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:apps/v1:DeploymentPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:apps/v1:DeploymentPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Apps.V1.DeploymentPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Apps.V1.DeploymentPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Apps.V1.DeploymentPatchArgs();
            args.ApiVersion = "apps/v1";
            args.Kind = "Deployment";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "kubernetes:apps/v1beta1:DeploymentPatch"},
                    new global::Pulumi.Alias { Type = "kubernetes:apps/v1beta2:DeploymentPatch"},
                    new global::Pulumi.Alias { Type = "kubernetes:extensions/v1beta1:DeploymentPatch"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DeploymentPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeploymentPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Apps.V1
{

    public class DeploymentPatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaPatchArgs>? Metadata { get; set; }

        /// <summary>
        /// Specification of the desired behavior of the Deployment.
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecPatchArgs>? Spec { get; set; }

        public DeploymentPatchArgs()
        {
        }
        public static new DeploymentPatchArgs Empty => new DeploymentPatchArgs();
    }
}
