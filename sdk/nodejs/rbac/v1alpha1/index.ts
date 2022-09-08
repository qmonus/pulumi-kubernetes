// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ClusterRoleArgs } from "./clusterRole";
export type ClusterRole = import("./clusterRole").ClusterRole;
export const ClusterRole: typeof import("./clusterRole").ClusterRole = null as any;

export { ClusterRoleBindingArgs } from "./clusterRoleBinding";
export type ClusterRoleBinding = import("./clusterRoleBinding").ClusterRoleBinding;
export const ClusterRoleBinding: typeof import("./clusterRoleBinding").ClusterRoleBinding = null as any;

export { ClusterRoleBindingListArgs } from "./clusterRoleBindingList";
export type ClusterRoleBindingList = import("./clusterRoleBindingList").ClusterRoleBindingList;
export const ClusterRoleBindingList: typeof import("./clusterRoleBindingList").ClusterRoleBindingList = null as any;

export { ClusterRoleBindingPatchArgs } from "./clusterRoleBindingPatch";
export type ClusterRoleBindingPatch = import("./clusterRoleBindingPatch").ClusterRoleBindingPatch;
export const ClusterRoleBindingPatch: typeof import("./clusterRoleBindingPatch").ClusterRoleBindingPatch = null as any;

export { ClusterRoleListArgs } from "./clusterRoleList";
export type ClusterRoleList = import("./clusterRoleList").ClusterRoleList;
export const ClusterRoleList: typeof import("./clusterRoleList").ClusterRoleList = null as any;

export { ClusterRolePatchArgs } from "./clusterRolePatch";
export type ClusterRolePatch = import("./clusterRolePatch").ClusterRolePatch;
export const ClusterRolePatch: typeof import("./clusterRolePatch").ClusterRolePatch = null as any;

export { RoleArgs } from "./role";
export type Role = import("./role").Role;
export const Role: typeof import("./role").Role = null as any;

export { RoleBindingArgs } from "./roleBinding";
export type RoleBinding = import("./roleBinding").RoleBinding;
export const RoleBinding: typeof import("./roleBinding").RoleBinding = null as any;

export { RoleBindingListArgs } from "./roleBindingList";
export type RoleBindingList = import("./roleBindingList").RoleBindingList;
export const RoleBindingList: typeof import("./roleBindingList").RoleBindingList = null as any;

export { RoleBindingPatchArgs } from "./roleBindingPatch";
export type RoleBindingPatch = import("./roleBindingPatch").RoleBindingPatch;
export const RoleBindingPatch: typeof import("./roleBindingPatch").RoleBindingPatch = null as any;

export { RoleListArgs } from "./roleList";
export type RoleList = import("./roleList").RoleList;
export const RoleList: typeof import("./roleList").RoleList = null as any;

export { RolePatchArgs } from "./rolePatch";
export type RolePatch = import("./rolePatch").RolePatch;
export const RolePatch: typeof import("./rolePatch").RolePatch = null as any;

utilities.lazyLoad(exports, ["ClusterRole"], () => require("./clusterRole"));
utilities.lazyLoad(exports, ["ClusterRoleBinding"], () => require("./clusterRoleBinding"));
utilities.lazyLoad(exports, ["ClusterRoleBindingList"], () => require("./clusterRoleBindingList"));
utilities.lazyLoad(exports, ["ClusterRoleBindingPatch"], () => require("./clusterRoleBindingPatch"));
utilities.lazyLoad(exports, ["ClusterRoleList"], () => require("./clusterRoleList"));
utilities.lazyLoad(exports, ["ClusterRolePatch"], () => require("./clusterRolePatch"));
utilities.lazyLoad(exports, ["Role"], () => require("./role"));
utilities.lazyLoad(exports, ["RoleBinding"], () => require("./roleBinding"));
utilities.lazyLoad(exports, ["RoleBindingList"], () => require("./roleBindingList"));
utilities.lazyLoad(exports, ["RoleBindingPatch"], () => require("./roleBindingPatch"));
utilities.lazyLoad(exports, ["RoleList"], () => require("./roleList"));
utilities.lazyLoad(exports, ["RolePatch"], () => require("./rolePatch"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole":
                return new ClusterRole(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBinding":
                return new ClusterRoleBinding(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBindingList":
                return new ClusterRoleBindingList(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBindingPatch":
                return new ClusterRoleBindingPatch(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleList":
                return new ClusterRoleList(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRolePatch":
                return new ClusterRolePatch(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:Role":
                return new Role(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBinding":
                return new RoleBinding(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBindingList":
                return new RoleBindingList(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBindingPatch":
                return new RoleBindingPatch(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleList":
                return new RoleList(name, <any>undefined, { urn })
            case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RolePatch":
                return new RolePatch(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "rbac.authorization.k8s.io/v1alpha1", _module)
