import * as React from 'react';

export default () => (
	<svg x="0px" y="0px" width="100px" height="200px" viewBox="0 0 100 200">
		<defs>
			<marker
				id="m_atr"
				markerUnits="strokeWidth"
				markerWidth="4"
				markerHeight="4"
				viewBox="0 0 10 10"
				refX="5"
				refY="5"
				orient="auto"
			>
				<polygon points="0,0 2,5 0,10 10,5 " fill="#fff" />
			</marker>
		</defs>
		<polyline
			points="70,5 20,5 20,70"
			style={{
				fill: 'transparent',
				stroke: '#fff',
				strokeWidth: 2,
				strokeLinecap: 'butt',
				strokeLinejoin: 'bevel',
				markerEnd: 'url(#m_atr)',
			}}
		/>
		<text
			x="20"
			y="100"
			style={{
				fontFamily: 'Verdana',
				fontSize: 15,
				fill: '#fff',
				textAnchor: 'middle',
				dominantBaseline: 'auto',
			}}
		>
			λ(x)
		</text>
		<polyline
			points="20,130 20,195 70,195"
			style={{
				fill: 'transparent',
				stroke: '#fff',
				strokeWidth: 2,
				strokeLinecap: 'butt',
				strokeLinejoin: 'bevel',
				markerStart: 'url(#m_atr)',
			}}
		/>
	</svg>
);
