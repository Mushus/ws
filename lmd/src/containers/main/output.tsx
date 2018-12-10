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
	if (typeof result === 'object') {
		return <OutputArea>{createObjectDOM(result)}</OutputArea>;
	}
	return <OutputArea>{result}</OutputArea>;
};

const createObjectDOM = (result: any): JSX.Element | string | null => {
	if (result == null) {
		return null;
	} else if (typeof result === 'object') {
		const keys = Object.keys(result);
		return (
			<div>
				{'{'}
				{keys.map((key, index) => (
					<details key={key}>
						<summary>{JSON.stringify(key)}:</summary>
						{createObjectDOM(result[key])}
						{index !== keys.length - 1 && ','}
					</details>
				))}
				{'{'}
			</div>
		);
	}
	return JSON.stringify(result);
};

const container = connect(
	mapStateToProps,
	mapDispatchToProps,
)(component);
export default container;
