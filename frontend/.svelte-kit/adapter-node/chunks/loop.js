import { w as writeAndUnwriteText } from "./writeAndUnwriteText.js";
import { a as animationSetup, m as makeNestedStaticElementsVisible } from "./animationSetup.js";
const loop = async (node, props) => {
  const { options, elements } = animationSetup(node, props);
  while (true) {
    makeNestedStaticElementsVisible(node);
    for (const element of elements)
      await writeAndUnwriteText(element, options);
  }
  return {
    update() {
    },
    destroy() {
    }
  };
};
export {
  loop as default
};
