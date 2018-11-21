import * as React from 'react';
import * as ReactDOM from 'react-dom'
import App from './App'

ReactDOM.render(
  <App />,
  document.querySelector('#app')
);

(window as any).cms = {
  customComponents: [],
  registerComponent: function (component: any) {
    this.customComponents.push(component);
  }
};
