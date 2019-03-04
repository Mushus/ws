import * as React from 'react';
import Editor from '~/components/Editor';

interface StateProps {}

interface ActionProps {}

type Props = StateProps & ActionProps;

const App = () => (
	<React.Fragment>
		<Editor />
	</React.Fragment>
);

export default App;
