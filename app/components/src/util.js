// html と css をコンポーネントに変換する
export function bundle(html, css) {
  return class extends HTMLElement {
    constructor() {
      super();
      this.attachShadow({ mode: 'open' });
      inject(this.shadowRoot, html, css);
    }
  };
}

// コンポーネント定義
// HTML ファイルでも ESModule でもどちらでもインポートできます
export function component(name, def) {
  let component = null;
  // 一応 def に html ファイルをそのまま入れるのも対応している
  if (typeof def === 'string') {
    const elem = bundle(def);
    return window.customElements.define(name, elem);
  }
  return window.customElements.define(name, def);
}

// 要素 elem に html と css を注入する
export function inject(elem, html, css) {
  console.log;
  if (css == null) {
    elem.innerHTML = html;
  }
  elem.innerHTML = `<style>${css}</style>${html}`;
}
