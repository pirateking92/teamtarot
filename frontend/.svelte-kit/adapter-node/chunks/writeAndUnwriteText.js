import { w as writeEffect, t as typingInterval } from "./writeEffect.js";
import { u as unwriteEffect } from "./unwriteEffect.js";
import { r as runOnEveryParentUntil } from "./animationSetup.js";
const writeAndUnwriteText = async ({ currentNode, text }, options) => {
  await writeEffect({ currentNode, text }, options);
  const textWithUnescapedAmpersands = text.replaceAll("&", "&amp;");
  const fullyWritten = currentNode.innerHTML === textWithUnescapedAmpersands;
  if (fullyWritten) {
    options.dispatch("done", "write");
    await typingInterval(options.wordInterval);
    await unwriteEffect(currentNode, options);
    options.dispatch("done", "unwrite");
  }
  runOnEveryParentUntil(currentNode, options.parentElement, (element) => {
    currentNode === element ? element.classList.remove("typing") : element.classList.remove("finished-typing");
  });
};
export {
  writeAndUnwriteText as w
};
