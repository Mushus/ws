import * as React from 'react';
import styled from 'styled-components';
import { connect } from 'react-redux';
import { IState } from '@/modules';

interface StateProp {
	result: any;
}

interface ActionProps {}

type Props = StateProp & ActionProps;

const mapStateToProps = ({ todo: { result } }: IState): StateProp => ({
	result,
});

const mapDispatchToProps = (dispatch: any): ActionProps => ({});

const InputArea = styled.textarea`
	box-sizing: border-box;
	width: 100%;
	height: 100%;
`;

const component = ({ result }: Props) => {
	console.log(result);
	return <div>result: {result}</div>;
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;
