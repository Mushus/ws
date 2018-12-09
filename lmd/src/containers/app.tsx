import styled from 'styled-components';
import Input from '@/containers/main/input';
import Editor from '@/containers/main/editor';
import Output from '@/containers/main/output';
import ProcessButton from '@/containers/main/processBtn';

import * as React from 'react';

const PaneBox = styled.div`
	width: 100%;
	height: 100%;
	display: grid;
	grid-template:
		'left c-line r-top' 1fr
		'left c-line r-bottom' 1fr
		/ 1fr 10px 1fr;
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
const CLineArea = styled.div`
	grid-area: c-line;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
`;
const CenterPane = styled.div`
	position: absolute;
	opacity: 0.5;
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
			<Output />
		</RBottomArea>
		<CLineArea>
			<CenterPane>
				<ProcessButton />
			</CenterPane>
		</CLineArea>
	</PaneBox>
);
