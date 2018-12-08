import * as React from 'react';
import styled from 'styled-components';
import { connect } from 'react-redux';
import AceEditor from 'react-ace';
import { IState } from '@/modules';
import { InputLambda } from '@/modules/todo';

import 'brace/mode/java';
import 'brace/theme/tomorrow_night';

interface StateProp {
	lambda: string;
}

interface ActionProps {
	inputLambda: (text: string) => void;
}

type Props = StateProp & ActionProps;

const mapStateToProps = ({ todo: { lambda } }: IState): StateProp => ({
	lambda,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	inputLambda: text => dispatch(InputLambda(text)),
});

const component = ({ lambda, inputLambda }: Props) => {
	return (
		<AceEditor
			mode="javascript"
			theme="tomorrow_night"
			width="100%"
			height="100%"
			value={lambda}
			onChange={text => inputLambda(text)}
		/>
	);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;
