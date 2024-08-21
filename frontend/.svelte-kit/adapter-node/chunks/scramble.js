import { b as rng, a as animationSetup, r as runOnEveryParentUntil, s as sleep } from "./animationSetup.js";
const getLettersTimeout = (textLetters, timeout) => {
  const minimumTimeoutPossible = timeout / 3;
  const lettersTimeout = textLetters.map(() => rng(minimumTimeoutPossible, timeout - 100));
  return lettersTimeout;
};
const getRandomLetter = () => {
  const possibleLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".split(
    ""
  );
  const letterIndexLimit = possibleLetters.length;
  const randomLetterIndex = rng(0, letterIndexLimit);
  const randomLetter = possibleLetters[randomLetterIndex];
  return randomLetter;
};
const scramble = async (node, props) => {
  const { options, elements } = animationSetup(node, props);
  const timeout = options.scrambleDuration;
  await new Promise((resolve) => {
    elements.forEach(async ({ currentNode, text }) => {
      let wordLetters = text.split("");
      const lettersTimeout = getLettersTimeout(wordLetters, timeout);
      const startingTime = Date.now();
      runOnEveryParentUntil(currentNode, options.parentElement, (element) => {
        element.classList.add("finished-typing");
      });
      while (Date.now() - startingTime < timeout) {
        const randomLetterIndex = rng(0, wordLetters.length);
        const randomLetterTimeout = lettersTimeout[randomLetterIndex];
        const isRandomLetterWhitespace = wordLetters[randomLetterIndex] === " ";
        const timeEllapsed = () => Date.now() - startingTime;
        const didRandomLetterReachTimeout = () => timeEllapsed() >= randomLetterTimeout;
        if (didRandomLetterReachTimeout() || isRandomLetterWhitespace) {
          const letterFinishedAnimation = wordLetters[randomLetterIndex] === text[randomLetterIndex];
          if (!letterFinishedAnimation)
            wordLetters[randomLetterIndex] = text[randomLetterIndex];
          else
            continue;
        } else {
          wordLetters[randomLetterIndex] = getRandomLetter();
        }
        const scrambledText = wordLetters.join("");
        currentNode.innerHTML = scrambledText;
        const finishedScrambling = scrambledText === text;
        const letterInterval = options.scrambleSlowdown ? Math.round(timeEllapsed() / 100) : 1;
        await sleep(letterInterval);
        if (finishedScrambling) {
          resolve();
          break;
        }
      }
      currentNode.innerHTML = text;
    });
  });
  options.dispatch("done");
  return {
    update() {
    },
    destroy() {
    }
  };
};
export {
  scramble as default
};
