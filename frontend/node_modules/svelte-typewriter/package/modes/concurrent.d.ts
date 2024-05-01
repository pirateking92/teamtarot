export default concurrent;
export type Props = {
    interval?: number | undefined;
    cursor?: boolean | undefined;
    keepCursorOnFinish?: boolean | undefined;
    delay?: number | undefined;
    showCursorOnDelay?: boolean | undefined;
    disabled?: boolean | undefined;
    element?: string | undefined;
    scrambleDuration?: number | undefined;
    scrambleSlowdown?: number | undefined;
    unwriteInterval?: number | undefined;
    wordInterval?: number | undefined;
};
export type SvelteActionReturnType = {
    update: () => void;
    destroy: () => void;
};
export type TypewriterModeFn = (node: HTMLElement, props: Props) => SvelteActionReturnType;
/**
 * @typedef {object} Props
 * @property {number} [interval]
 * @property {boolean} [cursor]
 * @property {boolean} [keepCursorOnFinish]
 * @property {number} [delay]
 * @property {boolean} [showCursorOnDelay]
 * @property {boolean} [disabled]
 * @property {string} [element]
 * @property {number} [scrambleDuration]
 * @property {number} [scrambleSlowdown]
 * @property {number} [unwriteInterval]
 * @property {number} [wordInterval]
 */
/**
 * @typedef {{ update: () => void, destroy: () => void }} SvelteActionReturnType
 */
/**
 * @typedef {(node: HTMLElement, props: Props) => SvelteActionReturnType} TypewriterModeFn
 */
/**
 * @type {TypewriterModeFn}
 */
declare const concurrent: TypewriterModeFn;
