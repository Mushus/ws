import * as React from 'react';

export interface ILineText {
	text: string;
}

export const Line = ({ line: { text } }: { line: ILineText }) => (
	<span>{text}</span>
);
