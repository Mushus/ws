import * as React from 'react';
import styled from 'styled-components';
import { connect } from 'react-redux';
import { IState } from '@/modules';

interface StateProp {
	result: any;
}

interface ActionProps {}

type Props = StateProp & ActionProps;

const mapStateToProps = ({ main: { result } }: IState): StateProp => ({
	result,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({});

const OutputArea = styled.div`
	width: 100%;
	height: 100%;
	color: #ccc;
	background-color: #1d1f21;
	overflow: auto;
	white-space: pre-wrap;
`;

const component = ({ result }: Props) => {
	return <OutputArea>{result}</OutputArea>;
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;
