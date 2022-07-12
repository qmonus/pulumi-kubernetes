// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.FlowControl.V1Alpha1
{
    /// <summary>
    /// FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a "flow distinguisher".
    /// </summary>
    [KubernetesResourceType("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaPatch")]
    public partial class FlowSchemaPatch : KubernetesResource
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
        /// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch> Metadata { get; private set; } = null!;

        /// <summary>
        /// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.FlowSchemaSpecPatch> Spec { get; private set; } = null!;

        /// <summary>
        /// `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.FlowSchemaStatusPatch> Status { get; private set; } = null!;


        /// <summary>
        /// Create a FlowSchemaPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlowSchemaPatch(string name, Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Alpha1.FlowSchemaPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal FlowSchemaPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private FlowSchemaPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Alpha1.FlowSchemaPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Alpha1.FlowSchemaPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Alpha1.FlowSchemaPatchArgs();
            args.ApiVersion = "flowcontrol.apiserver.k8s.io/v1alpha1";
            args.Kind = "FlowSchema";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:flowcontrol.apiserver.k8s.io/v1beta1:FlowSchemaPatch"},
                    new Pulumi.Alias { Type = "kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:FlowSchemaPatch"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FlowSchemaPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlowSchemaPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FlowSchemaPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Alpha1
{

    public class FlowSchemaPatchArgs : Pulumi.ResourceArgs
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
        /// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaPatchArgs> Metadata { get; set; } = null!;

        /// <summary>
        /// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Alpha1.FlowSchemaSpecPatchArgs>? Spec { get; set; }

        public FlowSchemaPatchArgs()
        {
        }
    }
}
