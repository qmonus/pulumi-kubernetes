// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Job represents the configuration of a single job.
 *
 * This resource waits until its status is ready before registering success
 * for create/update, and populating output properties from the current state of the resource.
 * The following conditions are used to determine whether the resource creation has
 * succeeded or failed:
 *
 * 1. The Job's '.status.startTime' is set, which indicates that the Job has started running.
 * 2. The Job's '.status.conditions' has a status of type 'Complete', and a 'status' set
 *    to 'True'.
 * 3. The Job's '.status.conditions' do not have a status of type 'Failed', with a
 *     'status' set to 'True'. If this condition is set, we should fail the Job immediately.
 *
 * If the Job has not reached a Ready state after 10 minutes, it will
 * time out and mark the resource update as Failed. You can override the default timeout value
 * by setting the 'customTimeouts' option on the resource.
 *
 * By default, if a resource failed to become ready in a previous update,
 * Pulumi will continue to wait for readiness on the next update. If you would prefer
 * to schedule a replacement for an unready resource on the next update, you can add the
 * "pulumi.com/replaceUnready": "true" annotation to the resource definition.
 */
export class JobPatch extends pulumi.CustomResource {
    /**
     * Get an existing JobPatch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JobPatch {
        return new JobPatch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:batch/v1:JobPatch';

    /**
     * Returns true if the given object is an instance of JobPatch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobPatch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobPatch.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"batch/v1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"Job">;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMetaPatch>;
    /**
     * Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    public readonly spec!: pulumi.Output<outputs.batch.v1.JobSpecPatch>;
    /**
     * Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.batch.v1.JobStatusPatch>;

    /**
     * Create a JobPatch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: JobPatchArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.metadata === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadata'");
            }
            resourceInputs["apiVersion"] = "batch/v1";
            resourceInputs["kind"] = "Job";
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
        super(JobPatch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a JobPatch resource.
 */
export interface JobPatchArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"batch/v1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"Job">;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata: pulumi.Input<inputs.meta.v1.ObjectMetaPatch>;
    /**
     * Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: pulumi.Input<inputs.batch.v1.JobSpecPatch>;
}
