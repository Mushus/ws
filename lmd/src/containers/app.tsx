import styled from 'styled-components';
import Input from '@/containers/main/input';
import Editor from '@/containers/main/editor';
import Result from '@/containers/main/result';

import * as React from 'react';

const PaneBox = styled.div`
	display: flex;
`;
const Panel = styled.div`
	flex: 1;
`;

export default () => (
	<PaneBox>
		<Panel>
			<Input />
		</Panel>
		<Panel>
			<Editor />
		</Panel>
		<Panel>
			<Result />
		</Panel>
	</PaneBox>
);
