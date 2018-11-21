import * as React from 'react';
import * as ReactDOM from 'react-dom';
import './myButton'

export default class App extends React.Component {
  render() {
    return (
      <div>
        hello world!
        {
          React.createElement('my-button', {
            click: 'hoge'
          })
        }
      </div>
    );
  }
}
