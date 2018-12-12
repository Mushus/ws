import * as React from 'react';
import styled from 'styled-components';
import { connect } from 'react-redux';
import AceEditor from 'react-ace';
import { IState } from '@/modules';
import { InputFormulaAction, CalcResultAction, IFile } from '@/modules/main';

import 'brace/mode/java';
import 'brace/theme/tomorrow_night';

interface StateProp {
	fileList: Array<IFile>;
	formula: string;
}

interface ActionProps {
	inputFormula: (formula: string) => void;
	calcResult: () => void;
}

type Props = StateProp & ActionProps;

const mapStateToProps = ({
	main: { formula, fileList },
}: IState): StateProp => ({
	formula,
	fileList,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	inputFormula: formula => dispatch(InputFormulaAction(formula)),
	calcResult: () => dispatch(CalcResultAction()),
});

const component = ({ formula, fileList, inputFormula, calcResult }: Props) => {
	return (
		<Root>
			<ControllBox>
				<select>
					{fileList.map((file, index) => (
						<option key={index} value={file.name}>
							{file.name}
						</option>
					))}
				</select>
				<button type="button">+</button>
				<button type="button">-</button>
			</ControllBox>

			<EditArea>
				<AceEditor
					mode="javascript"
					theme="tomorrow_night"
					width="100%"
					height="100%"
					value={formula}
					onChange={formula => {
						inputFormula(formula);
						calcResult();
					}}
				/>
			</EditArea>
		</Root>
	);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;

const controlHeight = '25px';

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
	background-color: #34383b;
`;

const EditArea = styled.div`
	display: block;
`;
