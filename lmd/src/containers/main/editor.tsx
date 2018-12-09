import * as React from 'react';
import styled from 'styled-components';
import { connect } from 'react-redux';
import AceEditor from 'react-ace';
import { IState } from '@/modules';
import { InputFormulaAction, CalcResultAction } from '@/modules/main';

import 'brace/mode/java';
import 'brace/theme/tomorrow_night';

interface StateProp {
	formula: string;
}

interface ActionProps {
	inputFormula: (formula: string) => void;
	calcResult: () => void;
}

type Props = StateProp & ActionProps;

const mapStateToProps = ({ main: { formula } }: IState): StateProp => ({
	formula,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	inputFormula: formula => dispatch(InputFormulaAction(formula)),
	calcResult: () => dispatch(CalcResultAction()),
});

const component = ({ formula, inputFormula, calcResult }: Props) => {
	return (
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
	);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;
