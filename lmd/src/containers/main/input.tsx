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

const Root = styled.div`
	display: grid;
	width: 100%;
	height: 100%;
	background-color: #1d1f21;
	grid-template:
		'top' 20px
		'bottom' 1fr
		/ 1fr;
`;
const InputArea = styled.textarea`
	display: block;
	color: #ccc;
	border: 0 none;
	overflow: auto;

	:focus {
		outline: 0;
		border-color: orange;
	}
`;
const ControllBox = styled.div`
	width
`;

const component = ({ input, text }: Props) => {
	return (
		<Root>
			<ControllBox>
				<select>
					<option value="text">テキスト全体</option>
					<option value="">一行ごと</option>
					<option value="json">JSON</option>
					<option value="csv">CSV</option>
					<option value="tsv">TSV</option>
				</select>
				<button type="button">入力から</button>
				<button type="button">ファイルから</button>
			</ControllBox>
			<InputArea
				placeholder="input here"
				value={text}
				onChange={e => {
					console.log(e.target.value);
					input(e.target.value);
				}}
			/>
		</Root>
	);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;
