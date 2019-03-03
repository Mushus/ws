import React, { useState, useCallback, useMemo } from 'react';
import {
	Editor as DraftEditor,
	EditorState,
	CompositeDecorator,
	ContentBlock,
} from 'draft-js';
import Context from '~/components/Editor/Conotext';

interface DecoratorProp {
	children: React.ReactChildren;
	offsetKey: string;
}

type DecoratorCallback = (start: number, end: number) => void;

const backQuoteDecorator = ({ children }: DecoratorProp) => <b>{children}</b>;
const CodeDecorator = ({ children, offsetKey }: DecoratorProp) => (
	<Context.Consumer>
		{({ selected }) => {
			const editThisLine = selected.some(key =>
				offsetKey.startsWith(key),
			);
			return (
				<b
					style={{
						fontSize: '1.5rem',
						color: editThisLine ? 'blue' : 'black',
					}}
				>
					{children}
				</b>
			);
		}}
	</Context.Consumer>
);
const HashTagDecorator = ({ children, offsetKey }: DecoratorProp) => (
	<Context.Consumer>
		{v => <b style={{ color: 'blue' }}>{children}</b>}
	</Context.Consumer>
);

const decorator = new CompositeDecorator([
	{
		strategy: (block, callback, content) => {
			const blockKey = block.getKey();
			console.log(blockKey);
			console.log(String(content.getSelectionBefore()));
			console.log(content.getSelectionAfter());
			content.getBlockMap;

			findWithRegex(/`[^`]+`/g, block, callback);
		},
		component: CodeDecorator,
	},
	{
		strategy: (block, callback) => {
			findWithRegex(/`/g, block, callback);
		},
		component: backQuoteDecorator,
	},
	{
		strategy: (block, callback) => {
			findWithRegex(/#[\w\u0590-\u05ff]+/g, block, callback);
		},
		component: HashTagDecorator,
	},
]);

const InitialState = EditorState.createEmpty(decorator);

const findWithRegex = (
	regex: RegExp,
	contentBlock: ContentBlock,
	callback: DecoratorCallback,
) => {
	const text = contentBlock.getText();
	let matchArr, start;
	while ((matchArr = regex.exec(text)) !== null) {
		start = matchArr.index;
		callback(start, start + matchArr[0].length);
	}
};

/**
 * @description 見たまんまのエディタ
 */
const Editor = () => {
	const [input, setInput] = useState(InitialState);
	const handleChange = useCallback(
		(input: EditorState) => {
			setInput(input);
		},
		[input],
	);

	const selection = input.getSelection();
	const anchor = selection.getAnchorKey();
	const focus = selection.getFocusKey();
	/** @description 選択している行 */
	const selected = useMemo(() => {
		const currentContent = input.getCurrentContent();
		const blockKeys = currentContent.getBlockMap().keys();

		let isSelect = false;
		let isAnkerStart = false;
		let blockResult: { value: string; done: boolean };
		const selected: string[] = [];
		while (!(blockResult = blockKeys.next()).done) {
			const block = blockResult.value;
			console.log(block);
			if (block == anchor) {
				isAnkerStart = true;
				isSelect = true;
			}
			if (block == focus) {
				isAnkerStart = false;
				isSelect = true;
			}
			if (isSelect) {
				selected.push(block);
				if (isAnkerStart && block == focus) {
					isSelect = false;
				} else if (!isAnkerStart && block == anchor) {
					isSelect = false;
				}
			}
		}
		return selected;
	}, [anchor, focus]);

	return (
		<Context.Provider value={{ selected }}>
			<DraftEditor editorState={input} onChange={handleChange} />
		</Context.Provider>
	);
};

export default Editor;
