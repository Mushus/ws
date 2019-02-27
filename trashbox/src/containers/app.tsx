import * as React from 'react';
import Wysiwyg from '~/components/Wysiwyg';

interface StateProps {}

interface ActionProps {}

type Props = StateProps & ActionProps;

const App = () => (
	<React.Fragment>
		<Wysiwyg />
	</React.Fragment>
);

export default App;
