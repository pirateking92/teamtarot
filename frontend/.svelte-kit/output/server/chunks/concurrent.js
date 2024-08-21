import { w as writeEffect } from "./writeEffect.js";
import { a as animationSetup } from "./animationSetup.js";
const descendingSortFunction = (firstElement, secondElement) => secondElement.text.length - firstElement.text.length;
const getLongestTextElement = (elements) => {
  const descendingTextLengthOrder = elements.sort(descendingSortFunction);
  const longestTextElement = descendingTextLengthOrder[0].currentNode;
  return longestTextElement;
};
const onAnimationEnd = (element, callback) => {
  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      const elementAttributeChanged = mutation.type === "attributes";
      const elementFinishedTyping = mutation.target.classList.contains("finished-typing");
      if (elementAttributeChanged && elementFinishedTyping)
        callback();
    });
  });
  observer.observe(element, {
    attributes: true,
    childList: true,
    subtree: true
  });
};
const concurrent = (node, props) => {
  const { options, elements } = animationSetup(node, props);
  const lastElementToFinish = getLongestTextElement(elements);
  onAnimationEnd(lastElementToFinish, () => options.dispatch("done"));
  for (const element of elements) {
    writeEffect(element, options).then(() => {
      if (options.keepCursorOnFinish) {
        const isNotLongestElement = element.currentNode !== lastElementToFinish;
        isNotLongestElement ? element.currentNode.classList.replace("typing", "finished-typing") : options.dispatch("done");
      } else {
        element.currentNode.classList.replace("typing", "finished-typing");
      }
      const cursorHasTimeout = typeof options.keepCursorOnFinish === "number";
      cursorHasTimeout && setTimeout(() => {
        element.currentNode.classList.replace("typing", "finished-typing");
      }, options.keepCursorOnFinish);
    });
  }
  return {
    update() {
    },
    destroy() {
    }
  };
};
export {
  concurrent as default
};
