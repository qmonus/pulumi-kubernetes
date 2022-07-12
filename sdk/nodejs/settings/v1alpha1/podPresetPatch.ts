// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * PodPreset is a policy resource that defines additional runtime requirements for a Pod.
 */
export class PodPresetPatch extends pulumi.CustomResource {
    /**
     * Get an existing PodPresetPatch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PodPresetPatch {
        return new PodPresetPatch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:settings.k8s.io/v1alpha1:PodPresetPatch';

    /**
     * Returns true if the given object is an instance of PodPresetPatch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PodPresetPatch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PodPresetPatch.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"settings.k8s.io/v1alpha1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"PodPreset">;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMetaPatch>;
    public readonly spec!: pulumi.Output<outputs.settings.v1alpha1.PodPresetSpecPatch>;

    /**
     * Create a PodPresetPatch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PodPresetPatchArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.metadata === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadata'");
            }
            resourceInputs["apiVersion"] = "settings.k8s.io/v1alpha1";
            resourceInputs["kind"] = "PodPreset";
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
        } else {
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["spec"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PodPresetPatch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PodPresetPatch resource.
 */
export interface PodPresetPatchArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"settings.k8s.io/v1alpha1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"PodPreset">;
    metadata: pulumi.Input<inputs.meta.v1.ObjectMetaPatch>;
    spec?: pulumi.Input<inputs.settings.v1alpha1.PodPresetSpecPatch>;
}
