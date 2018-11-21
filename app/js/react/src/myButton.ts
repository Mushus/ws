
class MyButton extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    console.log(this.shadowRoot)
    this.shadowRoot.innerHTML = `<button type="button"><script src="hoge.js"></script></script><slot>hello</slot></button>`;
    this.shadowRoot.querySelector('button').onclick = () => {
      console.log(this.getAttribute('click'));
      console.log(window)
    };
  }
}

customElements.define("my-button", MyButton);
console.log("hoge")

export default function() {}
