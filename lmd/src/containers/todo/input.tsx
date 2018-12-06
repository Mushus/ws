import * as React from "react"
import { connect } from "react-redux"
import styled from "styled-components"
import { IState } from "@/modules"
import { AddAction, InputAction } from "@/modules/todo"

interface StateProp {
	text: string
}

interface ActionProps {
	input: (text: string) => void
	add: () => void
}

type Props = StateProp & ActionProps

const mapStateToProps = ({ todo: { text } }: IState): StateProp => ({
	text,
})

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	input: text => dispatch(InputAction(text)),
	add: () => dispatch(AddAction()),
})

const InputArea = styled.textarea`
	box-sizing: border-box;
	width: 100%;
	height: 100%;
`

const component = ({ add, input, text }: Props) => {
	return (
		<div>
			<h2>input</h2>
			<div>
				<h3>直接入力</h3>
				<InputArea value={text} onChange={e => input(e.target.value)} />
			</div>
		</div>
	)
}

const container = connect(
	mapStateToProps,
	mapDispatchToProps
)(component)
export default container
