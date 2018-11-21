const html = `
<style>
p {
  color: #008000;
}
</style>
<div>
  <h2>sample component</h2>
  <p>hello world!</p>
  <button type="button" id="my-button">button</button>
</div>
`;

// デフォルト値としてコンポーネントをエクスポートします
export default class SimpleWidget extends HTMLElement {
  constructor() {
    super();

    // NOTE: DOMが干渉しないように
    this.root = this.attachShadow();
    this.root.innerHTML = html;

    // ボタンのイベントを設定
    const button = this.root.querySelector('#my-button');
    button.addEventListener("click", e => {
      window.alert('hello world!');
    });
  }
}


