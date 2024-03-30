// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("credentials");

/**
 * The (entirely uncryptographic) hash function used to encode the "password".
 */
export declare const hash: enums.HashKind | undefined;
Object.defineProperty(exports, "hash", {
    get() {
        return __config.getObject<enums.HashKind>("hash");
    },
    enumerable: true,
});

/**
 * The password. It is very secret.
 */
export declare const password: string;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password") ?? (utilities.getEnv("FOO") || "");
    },
    enumerable: true,
});

export declare const shared: outputs.Shared | undefined;
Object.defineProperty(exports, "shared", {
    get() {
        return __config.getObject<outputs.Shared>("shared");
    },
    enumerable: true,
});

/**
 * The username. Its important but not secret.
 */
export declare const user: string | undefined;
Object.defineProperty(exports, "user", {
    get() {
        return __config.get("user");
    },
    enumerable: true,
});
