// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Patch resources are used to modify existing Kubernetes resources by using
 * Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
 * one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
 * Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
 * [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/installation-configuration/#server-side-apply) for
 * additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
 * Status is a return value for calls that don't return other objects.
 */
export class StatusPatch extends pulumi.CustomResource {
    /**
     * Get an existing StatusPatch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StatusPatch {
        return new StatusPatch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:meta/v1:StatusPatch';

    /**
     * Returns true if the given object is an instance of StatusPatch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StatusPatch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StatusPatch.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"v1">;
    /**
     * Suggested HTTP return code for this status, 0 if not set.
     */
    public readonly code!: pulumi.Output<number>;
    /**
     * Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
     */
    public readonly details!: pulumi.Output<outputs.meta.v1.StatusDetailsPatch>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"Status">;
    /**
     * A human-readable description of the status of this operation.
     */
    public readonly message!: pulumi.Output<string>;
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ListMetaPatch>;
    /**
     * A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
     */
    public readonly reason!: pulumi.Output<string>;
    /**
     * Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a StatusPatch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StatusPatchArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["apiVersion"] = "v1";
            resourceInputs["code"] = args ? args.code : undefined;
            resourceInputs["details"] = args ? args.details : undefined;
            resourceInputs["kind"] = "Status";
            resourceInputs["message"] = args ? args.message : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["reason"] = args ? args.reason : undefined;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["code"] = undefined /*out*/;
            resourceInputs["details"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["message"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["reason"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StatusPatch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a StatusPatch resource.
 */
export interface StatusPatchArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"v1">;
    /**
     * Suggested HTTP return code for this status, 0 if not set.
     */
    code?: pulumi.Input<number>;
    /**
     * Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
     */
    details?: pulumi.Input<inputs.meta.v1.StatusDetailsPatch>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"Status">;
    /**
     * A human-readable description of the status of this operation.
     */
    message?: pulumi.Input<string>;
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: pulumi.Input<inputs.meta.v1.ListMetaPatch>;
    /**
     * A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
     */
    reason?: pulumi.Input<string>;
}
