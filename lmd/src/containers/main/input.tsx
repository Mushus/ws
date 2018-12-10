import * as React from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components';
import { IState } from '@/modules';
import {
	TabChangeAction,
	TextChangeAction,
	CalcResultAction,
} from '@/modules/main';

interface StateProp {
	input: string;
	tab: string;
}

interface ActionProps {
	changeTab: (tab: string) => void;
	inputText: (input: string) => void;
	calcResult: () => void;
}

type Props = StateProp & ActionProps;

const mapStateToProps = ({ main: { input, tab } }: IState): StateProp => ({
	input,
	tab,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	changeTab: tab => dispatch(TabChangeAction(tab)),
	inputText: input => dispatch(TextChangeAction(input)),
	calcResult: () => dispatch(CalcResultAction()),
});

const component = ({ tab, input, changeTab, inputText, calcResult }: Props) => {
	return (
		<Root>
			<ControllBox>
				<TabArea>
					<label className={tab === 'text' ? 'active' : 'inactive'}>
						<input
							type="radio"
							name="input-tab"
							value="text"
							checked={tab === 'text'}
							onChange={e => changeTab(e.target.value)}
						/>
						テキストから
					</label>
					<label className={tab === 'file' ? 'active' : 'inactive'}>
						<input
							type="radio"
							name="input-tab"
							value="file"
							checked={tab === 'file'}
							onChange={e => changeTab(e.target.value)}
						/>
						ファイルから
					</label>
				</TabArea>
				<OptionArea>
					<select>
						<option value="text">テキスト全体</option>
						<option value="">一行ごと</option>
						<option value="json">JSON</option>
						<option value="csv">CSV</option>
						<option value="tsv">TSV</option>
					</select>
				</OptionArea>
			</ControllBox>
			<InputField className={tab === 'text' ? 'active' : 'inactive'}>
				<InputArea
					placeholder="input here"
					value={tab === 'text' ? input : ''}
					onChange={e => {
						inputText(e.target.value);
						calcResult();
					}}
				/>
			</InputField>
			<InputField className={tab === 'file' ? 'active' : 'inactive'} />
		</Root>
	);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;

const fontSize = '1.2rem';
const controlHeight = '25px';
const fomratInputWidth = '140px';

const Root = styled.div`
	display: grid;
	width: 100%;
	height: 100%;
	background-color: #1d1f21;
	grid-template:
		'top' ${controlHeight}
		'bottom' 1fr
		/ 1fr;
`;

const ControllBox = styled.div`
	background-color: #35393b;
	display: grid;
	grid-template:
		'tab option' 1fr
		/ 1fr ${fomratInputWidth};
`;

const TabArea = styled.div`
	grid-area: tab;
	display: flex;
	label {
		box-sizing: border-box;
		padding: 2px 5px;
		width: 120px;
		height: ${controlHeight};
		font-size: ${fontSize};
		background-color: #35393b;
		color: #c4c8c6;
	}
	label.active {
		background-color: #1d1f20;
	}
	label input[type='radio'] {
		display: none;
	}
`;
const OptionArea = styled.div`
	grid-area: option;
	width: ${fomratInputWidth};
	height: ${controlHeight};
	font-size: ${fontSize};
	select {
		display: block;
		box-sizing: border-box;
		width: ${fomratInputWidth};
		height: ${controlHeight};
	}
`;

const InputField = styled.div`
	display: none;
	&.active {
		display: grid;
	}
`;

const InputArea = styled.textarea`
	display: block;
	color: #ccc;
	background-color: #1d1f20;
	border: 0 none;
	overflow: auto;
	padding: 10px;
	font-size: ${fontSize};

	:focus {
		outline: 0;
		border-color: orange;
	}
`;
