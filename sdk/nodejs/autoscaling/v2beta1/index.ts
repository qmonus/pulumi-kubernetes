// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { HorizontalPodAutoscalerArgs } from "./horizontalPodAutoscaler";
export type HorizontalPodAutoscaler = import("./horizontalPodAutoscaler").HorizontalPodAutoscaler;
export const HorizontalPodAutoscaler: typeof import("./horizontalPodAutoscaler").HorizontalPodAutoscaler = null as any;

export { HorizontalPodAutoscalerListArgs } from "./horizontalPodAutoscalerList";
export type HorizontalPodAutoscalerList = import("./horizontalPodAutoscalerList").HorizontalPodAutoscalerList;
export const HorizontalPodAutoscalerList: typeof import("./horizontalPodAutoscalerList").HorizontalPodAutoscalerList = null as any;

export { HorizontalPodAutoscalerPatchArgs } from "./horizontalPodAutoscalerPatch";
export type HorizontalPodAutoscalerPatch = import("./horizontalPodAutoscalerPatch").HorizontalPodAutoscalerPatch;
export const HorizontalPodAutoscalerPatch: typeof import("./horizontalPodAutoscalerPatch").HorizontalPodAutoscalerPatch = null as any;

utilities.lazyLoad(exports, ["HorizontalPodAutoscaler"], () => require("./horizontalPodAutoscaler"));
utilities.lazyLoad(exports, ["HorizontalPodAutoscalerList"], () => require("./horizontalPodAutoscalerList"));
utilities.lazyLoad(exports, ["HorizontalPodAutoscalerPatch"], () => require("./horizontalPodAutoscalerPatch"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler":
                return new HorizontalPodAutoscaler(name, <any>undefined, { urn })
            case "kubernetes:autoscaling/v2beta1:HorizontalPodAutoscalerList":
                return new HorizontalPodAutoscalerList(name, <any>undefined, { urn })
            case "kubernetes:autoscaling/v2beta1:HorizontalPodAutoscalerPatch":
                return new HorizontalPodAutoscalerPatch(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "autoscaling/v2beta1", _module)
