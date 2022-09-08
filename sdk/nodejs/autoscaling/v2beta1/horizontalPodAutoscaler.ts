// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
 */
export class HorizontalPodAutoscaler extends pulumi.CustomResource {
    /**
     * Get an existing HorizontalPodAutoscaler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HorizontalPodAutoscaler {
        return new HorizontalPodAutoscaler(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler';

    /**
     * Returns true if the given object is an instance of HorizontalPodAutoscaler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HorizontalPodAutoscaler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HorizontalPodAutoscaler.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"autoscaling/v2beta1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"HorizontalPodAutoscaler">;
    /**
     * metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
     */
    public readonly spec!: pulumi.Output<outputs.autoscaling.v2beta1.HorizontalPodAutoscalerSpec>;
    /**
     * status is the current information about the autoscaler.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.autoscaling.v2beta1.HorizontalPodAutoscalerStatus>;

    /**
     * Create a HorizontalPodAutoscaler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HorizontalPodAutoscalerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["apiVersion"] = "autoscaling/v2beta1";
            resourceInputs["kind"] = "HorizontalPodAutoscaler";
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["spec"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "kubernetes:autoscaling/v1:HorizontalPodAutoscaler" }, { type: "kubernetes:autoscaling/v2:HorizontalPodAutoscaler" }, { type: "kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(HorizontalPodAutoscaler.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a HorizontalPodAutoscaler resource.
 */
export interface HorizontalPodAutoscalerArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"autoscaling/v2beta1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"HorizontalPodAutoscaler">;
    /**
     * metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
     */
    spec?: pulumi.Input<inputs.autoscaling.v2beta1.HorizontalPodAutoscalerSpec>;
}
