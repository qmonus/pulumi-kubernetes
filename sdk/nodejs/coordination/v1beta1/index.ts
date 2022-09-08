// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { LeaseArgs } from "./lease";
export type Lease = import("./lease").Lease;
export const Lease: typeof import("./lease").Lease = null as any;

export { LeaseListArgs } from "./leaseList";
export type LeaseList = import("./leaseList").LeaseList;
export const LeaseList: typeof import("./leaseList").LeaseList = null as any;

export { LeasePatchArgs } from "./leasePatch";
export type LeasePatch = import("./leasePatch").LeasePatch;
export const LeasePatch: typeof import("./leasePatch").LeasePatch = null as any;

utilities.lazyLoad(exports, ["Lease"], () => require("./lease"));
utilities.lazyLoad(exports, ["LeaseList"], () => require("./leaseList"));
utilities.lazyLoad(exports, ["LeasePatch"], () => require("./leasePatch"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:coordination.k8s.io/v1beta1:Lease":
                return new Lease(name, <any>undefined, { urn })
            case "kubernetes:coordination.k8s.io/v1beta1:LeaseList":
                return new LeaseList(name, <any>undefined, { urn })
            case "kubernetes:coordination.k8s.io/v1beta1:LeasePatch":
                return new LeasePatch(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "coordination.k8s.io/v1beta1", _module)
