import{t as r}from"./writeEffect.BJvMesFy.js";const f=async(n,i)=>{const t=n.innerHTML.replaceAll("&amp;","&");for(let e=t.length-1;e>=0;e--){const a=t[e];a===">"&&(e=t.lastIndexOf("<",e)),!(new Blob([a]).size>1)&&(n.innerHTML=t.slice(0,e),await r(i.unwriteInterval?i.unwriteInterval:i.interval))}n.innerHTML=t};export{f as u};
