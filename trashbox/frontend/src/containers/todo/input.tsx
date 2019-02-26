import * as React from 'react';
import { connect } from 'react-redux';
import { IState } from '~/modules';
import { AddAction, InputAction } from '~/modules/todo';

interface StateProp {
	text: string;
}

interface ActionProps {
	input: (text: string) => void;
	add: () => void;
}

type Props = StateProp & ActionProps;

const mapStateToProps = ({ todo: { text } }: IState): StateProp => ({
	text,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	input: text => dispatch(InputAction(text)),
	add: () => dispatch(AddAction()),
});

const component = ({ add, input, text }: Props) => {
	return (
		<div>
			<input
				type="text"
				placeholder="Todo"
				value={text}
				onChange={e => input(text)}
			/>
			<button type="button" onClick={() => add()}>
				Add
			</button>
		</div>
	);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps
)(component);
export default container;
