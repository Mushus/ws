import styled from "styled-components"
import Input from "@/containers/todo/input"
import List from "@/containers/todo/list"

import * as React from "react"

const PaneBox = styled.div`
	display: flex;
`
const Panel = styled.div`
	flex: 1;
`

export default () => (
	<PaneBox>
		<Panel>
			<Input />
		</Panel>
		<Panel>
			<List />
		</Panel>
		<Panel />
	</PaneBox>
)
