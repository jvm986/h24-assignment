(this["webpackJsonph24-frontend"]=this["webpackJsonph24-frontend"]||[]).push([[0],{93:function(e,t,n){"use strict";n.r(t);var a=n(0),r=n(19),c=n.n(r),s=n(115),i=n(114),l=n(16),u=n.n(l),o=n(23),j=n(17),b=n(116),d=n(109),h=n(120),p=n(112),O=n(121),x=n(113),f=n(46),g=n.n(f),v="http://127.0.0.1:8000";function m(e){return k.apply(this,arguments)}function k(){return(k=Object(o.a)(u.a.mark((function e(t){var n;return u.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,g.a.get("".concat(v,"/url?url=").concat(t));case 2:return n=e.sent,e.abrupt("return",n);case 4:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var y=n(110),S=n(118),w=n(117),z=n(111),H=n(51),T=n.n(H),C=n(3),L=Object(d.a)((function(e){return Object(h.a)({root:{padding:"2px 4px",display:"flex",alignItems:"center"},input:{marginLeft:e.spacing(1),flex:1},iconButton:{padding:10}})})),N=function(e){var t=e.url,n=e.handleChange,a=e.handleSubmit,r=e.validUrl,c=e.loading,s=L();return Object(C.jsxs)(y.a,{className:s.root,children:[Object(C.jsx)(S.a,{className:s.input,placeholder:"Enter URL",inputProps:{"aria-label":"enter url"},value:t,onChange:function(e){return n(e.target.value)}}),Object(C.jsx)(w.a,{className:s.iconButton,"aria-label":"search",onClick:a,disabled:!r,children:c?Object(C.jsx)(z.a,{size:24}):Object(C.jsx)(T.a,{})})]})},E=Object(d.a)((function(e){return Object(h.a)({root:{padding:e.spacing(3,2),display:"flex",justifyContent:"center"}})})),I=function(){var e=E(),t=Object(a.useState)(""),n=Object(j.a)(t,2),r=n[0],c=n[1],s=Object(a.useState)(!1),i=Object(j.a)(s,2),l=i[0],d=i[1],h=Object(a.useState)(),f=Object(j.a)(h,2),g=f[0],v=f[1],k=Object(a.useState)(!1),y=Object(j.a)(k,2),S=y[0],w=y[1],z=Object(a.useState)(!1),H=Object(j.a)(z,2),T=H[0],L=H[1],I=Object(a.useState)(0),U=Object(j.a)(I,2),_=U[0],B=U[1],J=Object(a.useState)(""),P=Object(j.a)(J,2),R=P[0],V=P[1];function q(e){var t=0;return Promise.all(e.map(function(){var n=Object(o.a)(u.a.mark((function n(a,r){var c;return u.a.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return n.prev=0,n.next=3,m(a.link);case 3:c=n.sent,a.accessible=200==c.status,t++,B(t/e.length*100),n.next=14;break;case 9:return n.prev=9,n.t0=n.catch(0),t++,B(t/e.length*100),n.abrupt("return",a);case 14:return n.abrupt("return",a);case 15:case"end":return n.stop()}}),n,null,[[0,9]])})));return function(e,t){return n.apply(this,arguments)}}()))}function A(){return(A=Object(o.a)(u.a.mark((function e(){var t;return u.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return B(0),V(""),L(!1),d(!0),e.prev=4,e.next=7,m(r);case 7:if(t=e.sent,v(t.data),d(!1),!t.data.links){e.next=14;break}return e.next=13,q(t.data.links);case 13:L(!0);case 14:e.next=20;break;case 16:e.prev=16,e.t0=e.catch(4),V(e.t0.toString()),d(!1);case 20:case"end":return e.stop()}}),e,null,[[4,16]])})))).apply(this,arguments)}return Object(C.jsx)(p.a,{container:!0,className:e.root,alignItems:"center",children:Object(C.jsxs)(p.a,{item:!0,xs:12,md:6,lg:4,children:[Object(C.jsx)(b.a,{children:Object(C.jsx)(N,{url:r,handleChange:function(e){var t;B(0),V(""),L(!1),v(void 0),d(!1),c(e),w((t=e,!!new RegExp("^(https?:\\/\\/)?((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|((\\d{1,3}\\.){3}\\d{1,3}))(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*(\\?[;&a-z\\d%_.~+=-]*)?(\\#[-a-z\\d_]*)?$","i").test(t)))},handleSubmit:function(){return A.apply(this,arguments)},validUrl:S,loading:l})}),g?Object(C.jsxs)(b.a,{children:[Object(C.jsx)(b.a,{mt:4,mb:2,children:Object(C.jsx)(O.a,{variant:"determinate",value:_})}),Object(C.jsx)(x.a,{variant:"h5",children:g.title.replace("&amp;","&")}),Object(C.jsxs)(x.a,{children:["HTML Version: ",g.htmlVersion]}),Object(C.jsx)(x.a,{children:g.login?"Has a login form":"Does not have a login form"}),Object(C.jsxs)(x.a,{children:[g.links?g.links.filter((function(e){return!e.internal})).length:0," ","External Links :"," ",T?g.links?g.links.filter((function(e){return e.accessible&&!e.internal})).length+" accessible":"":"checking accessibility..."]}),Object(C.jsxs)(x.a,{children:[g.links?g.links.filter((function(e){return e.internal})).length:0," ","Internal Links :"," ",T?g.links?g.links.filter((function(e){return e.accessible&&e.internal})).length+" accessible":"":"checking accessibility..."]}),Object(C.jsxs)(x.a,{variant:"h1",children:[g.headings[0].count," H1 Tags"]}),Object(C.jsxs)(x.a,{variant:"h2",children:[g.headings[1].count," H2 Tags"]}),Object(C.jsxs)(x.a,{variant:"h3",children:[g.headings[2].count," H3 Tags"]}),Object(C.jsxs)(x.a,{variant:"h4",children:[g.headings[3].count," H4 Tags"]}),Object(C.jsxs)(x.a,{variant:"h5",children:[g.headings[4].count," H5 Tags"]}),Object(C.jsxs)(x.a,{variant:"h6",children:[g.headings[5].count," H6 Tags"]})]}):R?Object(C.jsx)(b.a,{mt:5,children:Object(C.jsx)(x.a,{children:R})}):""]})})},U=n(52),_=n.n(U),B=n(53),J=Object(B.a)({palette:{primary:{main:"#556cd6"},secondary:{main:"#19857b"},error:{main:_.a.A400},background:{default:"#fff"}}});c.a.render(Object(C.jsxs)(i.a,{theme:J,children:[Object(C.jsx)(s.a,{}),Object(C.jsx)(I,{})]}),document.querySelector("#root"))}},[[93,1,2]]]);
//# sourceMappingURL=main.3ccad8e8.chunk.js.map