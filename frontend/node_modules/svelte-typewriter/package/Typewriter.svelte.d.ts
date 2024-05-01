/** @typedef {typeof __propDef.props}  TypewriterProps */
/** @typedef {typeof __propDef.events}  TypewriterEvents */
/** @typedef {typeof __propDef.slots}  TypewriterSlots */
export default class Typewriter extends SvelteComponentTyped<{
    [x: string]: any;
    unwriteInterval?: number | undefined;
    wordInterval?: number | undefined;
    scrambleDuration?: number | undefined;
    scrambleSlowdown?: boolean | undefined;
    mode?: string | undefined;
    interval?: number | undefined;
    cursor?: boolean | undefined;
    keepCursorOnFinish?: boolean | undefined;
    delay?: number | undefined;
    showCursorOnDelay?: boolean | undefined;
    disabled?: boolean | undefined;
    element?: string | undefined;
}, {
    [evt: string]: CustomEvent<any>;
}, {
    default: {};
}> {
}
export type TypewriterProps = typeof __propDef.props;
export type TypewriterEvents = typeof __propDef.events;
export type TypewriterSlots = typeof __propDef.slots;
import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        unwriteInterval?: number | undefined;
        wordInterval?: number | undefined;
        scrambleDuration?: number | undefined;
        scrambleSlowdown?: boolean | undefined;
        mode?: string | undefined;
        interval?: number | undefined;
        cursor?: boolean | undefined;
        keepCursorOnFinish?: boolean | undefined;
        delay?: number | undefined;
        showCursorOnDelay?: boolean | undefined;
        disabled?: boolean | undefined;
        element?: string | undefined;
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {
        default: {};
    };
};
export {};
