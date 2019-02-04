import * as React from "react"
import { render } from "react-dom"
import { createStore } from "redux"
import { Provider } from "react-redux"
import App from "@/containers/app"
import { reducer } from "@/modules"
import '../assets/css/main.scss';

const store = createStore(reducer)

render(
	<Provider store={store}>
		<App />
	</Provider>,
	document.querySelector(".app")
)
