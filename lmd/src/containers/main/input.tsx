import * as React from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components';
import { IState } from '@/modules';
import { InputAction } from '@/modules/todo';

interface StateProp {
	text: string;
}

interface ActionProps {
	input: (text: string) => void;
}

type Props = StateProp & ActionProps;

const mapStateToProps = ({ todo: { text } }: IState): StateProp => ({
	text,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	input: text => dispatch(InputAction(text)),
});

const InputArea = styled.textarea`
	box-sizing: border-box;
	display: block;
	width: 100%;
	height: 100%;
	border: 0 none;
	color: #ccc;
	background-color: #1d1f21;
	overflow: auto;

	:focus {
		outline: 0;
		border-color: orange;
	}
`;

const component = ({ input, text }: Props) => {
	return (
		<InputArea
			value={text}
			onChange={e => {
				console.log(e.target.value);
				input(e.target.value);
			}}
		/>
	);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;
