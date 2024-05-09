import { w as writeAndUnwriteText } from './writeAndUnwriteText-CECu9zuZ.js';
import { a as animationSetup, m as makeNestedStaticElementsVisible } from './animationSetup-DLp9F-jH.js';
import './writeEffect-Cxwxs-XH.js';
import './unwriteEffect-CfV5mm0r.js';
import './ssr-tGUPulqZ.js';

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

export { loop as default };
//# sourceMappingURL=loop-N_SdXHTJ.js.map
