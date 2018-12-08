import styled from 'styled-components';
import Input from '@/containers/main/input';
import Editor from '@/containers/main/editor';
import Result from '@/containers/main/result';

import * as React from 'react';

const PaneBox = styled.div`
	width: 100%;
	height: 100%;
	display: grid;
	grid-template:
		'left r-top' 1fr
		'left r-bottom' 1fr
		/ 1fr 1fr;
	grid-gap: 10px;
`;
const LeftArea = styled.div`
	grid-area: left;
`;
const RTopArea = styled.div`
	grid-area: r-top;
`;
const RBottomArea = styled.div`
	grid-area: r-bottom;
`;

export default () => (
	<PaneBox>
		<LeftArea>
			<Editor />
		</LeftArea>
		<RTopArea>
			<Input />
		</RTopArea>
		<RBottomArea>
			<Result />
		</RBottomArea>
	</PaneBox>
);
