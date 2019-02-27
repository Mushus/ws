import * as React from 'react';
import { useState, useCallback } from 'react';
import { Line, ILineText } from '~/components/Wysiwyg/Line';

interface ILine {
	key: number;
	text: ILineText;
}

const Wysiwig = () => {
	const [lineId, setLineId] = useState(0);
	const [input, setInput] = useState('');
	const [text, setText] = useState<ILine[]>([]);
	const changeInput = useCallback(
		({ target: { value } }) => setInput(value),
		[]
	);
	const keyPress = useCallback(
		e => {
			const { keyCode } = e;
			// Enter
			if (keyCode === 13) {
				setText([
					...text,
					{
						key: lineId,
						text: {
							text: input,
						},
					},
				]);
				setLineId(lineId + 1);
				setInput('');
			}
		},
		[text, input]
	);

	return (
		<div>
			{text.map(line => (
				<Line line={line.text} key={line.key} />
			))}
			<input
				type="text"
				onKeyDown={keyPress}
				onChange={changeInput}
				value={input}
			/>
		</div>
	);
};

export default Wysiwig;
