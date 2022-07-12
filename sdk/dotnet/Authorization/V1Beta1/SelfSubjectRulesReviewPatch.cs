// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Authorization.V1Beta1
{
    /// <summary>
    /// SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace. The returned list of actions may be incomplete depending on the server's authorization mode, and any errors experienced during the evaluation. SelfSubjectRulesReview should be used by UIs to show/hide actions, or to quickly let an end user reason about their permissions. It should NOT Be used by external systems to drive authorization decisions as this raises confused deputy, cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview are the correct way to defer authorization decisions to the API server.
    /// </summary>
    [KubernetesResourceType("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectRulesReviewPatch")]
    public partial class SelfSubjectRulesReviewPatch : KubernetesResource
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

        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch> Metadata { get; private set; } = null!;

        /// <summary>
        /// Spec holds information about the request being evaluated.
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1.SelfSubjectRulesReviewSpecPatch> Spec { get; private set; } = null!;

        /// <summary>
        /// Status is filled in by the server and indicates the set of actions a user can perform.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1.SubjectRulesReviewStatusPatch> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SelfSubjectRulesReviewPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SelfSubjectRulesReviewPatch(string name, Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1.SelfSubjectRulesReviewPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectRulesReviewPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal SelfSubjectRulesReviewPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectRulesReviewPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private SelfSubjectRulesReviewPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectRulesReviewPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1.SelfSubjectRulesReviewPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1.SelfSubjectRulesReviewPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1.SelfSubjectRulesReviewPatchArgs();
            args.ApiVersion = "authorization.k8s.io/v1beta1";
            args.Kind = "SelfSubjectRulesReview";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReviewPatch"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SelfSubjectRulesReviewPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SelfSubjectRulesReviewPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SelfSubjectRulesReviewPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1
{

    public class SelfSubjectRulesReviewPatchArgs : Pulumi.ResourceArgs
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

        [Input("metadata", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaPatchArgs> Metadata { get; set; } = null!;

        /// <summary>
        /// Spec holds information about the request being evaluated.
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1.SelfSubjectRulesReviewSpecPatchArgs>? Spec { get; set; }

        public SelfSubjectRulesReviewPatchArgs()
        {
        }
    }
}
