import{s as ee,e as p,b as $,d as m,f as K,j as _,i as b,p as C,k as g,l as i,q as te,r as R,n as T,v as Q,w as Z,x as ne}from"../chunks/scheduler.CgzrBarl.js";import{S as se,i as le,c as L,b as P,m as S,t as V,a as O,d as W}from"../chunks/index.B7BUR2Ah.js";import{T as z}from"../chunks/Typewriter.2-Kk8hvz.js";function re(f){let e,l="Ah, I've been expecting you, seeker.";return{c(){e=p("h1"),e.textContent=l},l(n){e=m(n,"H1",{"data-svelte-h":!0}),R(e)!=="svelte-13nwx91"&&(e.textContent=l)},m(n,s){g(n,e,s)},p:T,d(n){n&&b(e)}}}function ae(f){let e,l="My name is Cassandra.";return{c(){e=p("h1"),e.textContent=l},l(n){e=m(n,"H1",{"data-svelte-h":!0}),R(e)!=="svelte-956bc1"&&(e.textContent=l)},m(n,s){g(n,e,s)},p:T,d(n){n&&b(e)}}}function ie(f){let e,l="What's yours?";return{c(){e=p("h1"),e.textContent=l},l(n){e=m(n,"H1",{"data-svelte-h":!0}),R(e)!=="svelte-1hn4ch9"&&(e.textContent=l)},m(n,s){g(n,e,s)},p:T,d(n){n&&b(e)}}}function X(f){let e,l,n;return{c(){e=p("input"),this.h()},l(s){e=m(s,"INPUT",{class:!0,placeholder:!0}),this.h()},h(){C(e,"class","input-field svelte-1st4oqi"),C(e,"placeholder","Enter your name")},m(s,t){g(s,e,t),Q(e,f[0]),l||(n=Z(e,"input",f[4]),l=!0)},p(s,t){t&1&&e.value!==s[0]&&Q(e,s[0])},d(s){s&&b(e),l=!1,n()}}}function Y(f){let e,l='<span class="svelte-1st4oqi">Follow Cassandra</span><i class="svelte-1st4oqi"></i>',n,s;return{c(){e=p("button"),e.innerHTML=l,this.h()},l(t){e=m(t,"BUTTON",{style:!0,class:!0,"data-svelte-h":!0}),R(e)!=="svelte-fu92jo"&&(e.innerHTML=l),this.h()},h(){ne(e,"--clr","#871616"),C(e,"class","svelte-1st4oqi")},m(t,d){g(t,e,d),n||(s=Z(e,"click",f[3]),n=!0)},p:T,d(t){t&&b(e),n=!1,s()}}}function oe(f){let e,l,n,s,t,d,w,h,k,q,H,v,I,M,x,E,y,F,N,U,j,A,D,B;d=new z({props:{delay:500,interval:60,$$slots:{default:[re]},$$scope:{ctx:f}}}),h=new z({props:{delay:4e3,interval:60,$$slots:{default:[ae]},$$scope:{ctx:f}}}),v=new z({props:{delay:6e3,interval:60,$$slots:{default:[ie]},$$scope:{ctx:f}}});let u=f[1]&&X(f),c=f[2]&&Y(f);return{c(){e=p("div"),l=p("br"),n=p("br"),s=$(),t=p("div"),L(d.$$.fragment),w=$(),L(h.$$.fragment),k=$(),q=p("br"),H=$(),L(v.$$.fragment),I=$(),M=p("br"),x=p("br"),E=$(),u&&u.c(),y=$(),F=p("br"),N=p("br"),U=p("br"),j=$(),c&&c.c(),A=$(),D=p("br"),this.h()},l(r){e=m(r,"DIV",{class:!0});var o=K(e);l=m(o,"BR",{}),n=m(o,"BR",{}),s=_(o),t=m(o,"DIV",{class:!0});var a=K(t);P(d.$$.fragment,a),w=_(a),P(h.$$.fragment,a),k=_(a),q=m(a,"BR",{}),H=_(a),P(v.$$.fragment,a),I=_(a),M=m(a,"BR",{}),x=m(a,"BR",{}),E=_(a),u&&u.l(a),y=_(a),F=m(a,"BR",{}),N=m(a,"BR",{}),U=m(a,"BR",{}),j=_(a),c&&c.l(a),a.forEach(b),A=_(o),D=m(o,"BR",{}),o.forEach(b),this.h()},h(){C(t,"class","intro-text svelte-1st4oqi"),C(e,"class","container svelte-1st4oqi")},m(r,o){g(r,e,o),i(e,l),i(e,n),i(e,s),i(e,t),S(d,t,null),i(t,w),S(h,t,null),i(t,k),i(t,q),i(t,H),S(v,t,null),i(t,I),i(t,M),i(t,x),i(t,E),u&&u.m(t,null),i(t,y),i(t,F),i(t,N),i(t,U),i(t,j),c&&c.m(t,null),i(e,A),i(e,D),B=!0},p(r,[o]){const a={};o&32&&(a.$$scope={dirty:o,ctx:r}),d.$set(a);const G={};o&32&&(G.$$scope={dirty:o,ctx:r}),h.$set(G);const J={};o&32&&(J.$$scope={dirty:o,ctx:r}),v.$set(J),r[1]?u?u.p(r,o):(u=X(r),u.c(),u.m(t,y)):u&&(u.d(1),u=null),r[2]?c?c.p(r,o):(c=Y(r),c.c(),c.m(t,null)):c&&(c.d(1),c=null)},i(r){B||(V(d.$$.fragment,r),V(h.$$.fragment,r),V(v.$$.fragment,r),B=!0)},o(r){O(d.$$.fragment,r),O(h.$$.fragment,r),O(v.$$.fragment,r),B=!1},d(r){r&&b(e),W(d),W(h),W(v),u&&u.d(),c&&c.d()}}}function fe(f,e,l){let n="",s=!1,t=!1;te(()=>{setTimeout(()=>{l(1,s=!0)},9e3),setTimeout(()=>{l(2,t=!0)},13e3)});function d(){const h=`/cassandra?name=${encodeURIComponent(n)}`;window.location.href=h}function w(){n=this.value,l(0,n)}return[n,s,t,d,w]}class me extends se{constructor(e){super(),le(this,e,fe,oe,ee,{})}}export{me as component};
