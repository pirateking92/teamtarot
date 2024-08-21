import { d as createEventDispatcher } from "./ssr.js";
const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
const rng = (min, max) => Math.floor(Math.random() * (max - min) + min);
const runOnEveryParentUntil = async (element, parent, callback) => {
  if (!parent) {
    console.error("The specified parent element does not exists!");
    return;
  }
  let currentElement = element;
  do {
    if (currentElement === parent)
      return;
    callback(currentElement);
    currentElement = currentElement.parentElement || currentElement.parentNode;
  } while (currentElement !== null && currentElement.nodeType === 1);
};
const hasSingleTextNode = (el) => el.childNodes.length === 1 && el.childNodes[0].nodeType === 3;
const createElement = (text, elementTag) => {
  const element = document.createElement(elementTag);
  element.textContent = text;
  return element;
};
const filterOutStaticElements = (child) => child.dataset.static === void 0;
const getElements = (node, { parentElement }) => {
  if (hasSingleTextNode(parentElement)) {
    const text = parentElement.textContent;
    const childNode = createElement(parentElement.textContent, "p");
    parentElement.textContent = "";
    parentElement.appendChild(childNode);
    return [{ currentNode: childNode, text }];
  }
  if (hasSingleTextNode(node)) {
    const textWithFilteredAmpersand = node.innerHTML.replaceAll("&amp;", "&");
    return [{ currentNode: node, text: textWithFilteredAmpersand }];
  } else {
    const children = [...node.children].filter(filterOutStaticElements);
    const allChildren = children.flatMap((child) => getElements(child, { parentElement }));
    return allChildren;
  }
};
const makeNestedStaticElementsVisible = (parentElement) => {
  const staticElements = [...parentElement.querySelectorAll("[data-static]")];
  for (const staticElement of staticElements) {
    runOnEveryParentUntil(staticElement, parentElement, (currentStaticElement) => {
      const isParentElement = currentStaticElement !== staticElement;
      isParentElement && currentStaticElement.classList.add("finished-typing");
    });
  }
};
const animationSetup = (node, props) => {
  const dispatch = createEventDispatcher();
  const options = { parentElement: node, dispatch, ...props };
  const elements = getElements(node, options);
  makeNestedStaticElementsVisible(node);
  return { options, elements };
};
export {
  animationSetup as a,
  rng as b,
  makeNestedStaticElementsVisible as m,
  runOnEveryParentUntil as r,
  sleep as s
};
