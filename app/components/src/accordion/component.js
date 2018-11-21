import { inject } from '@/util.js';
import html from './template.html';
import css from './style.scss';

// 表示非表示を行う
function Update(isOpen, elm) {
  elm.style.display = isOpen ? 'block' : 'none';
}

export default class Accordion extends HTMLElement {
  constructor() {
    super();
    // 外部と衝突しないように
    this.attachShadow({ mode: 'open' });
    inject(this.shadowRoot, html, css);

    let open = false;
    const header = this.shadowRoot.querySelector('#header');
    const content = this.shadowRoot.querySelector('#content');
    // 開閉イベント
    header.addEventListener('click', e => {
      open = !open;
      Update(open, content);
    });
    Update(open, content);
  }
}
