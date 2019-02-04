import * as React from "react"
import styled from "styled-components"
import { connect } from "react-redux"
import { IState } from "@/modules"
import { FinishAction } from "@/modules/todo"

interface StateProp {
	inProgress: string[]
	finished: string[]
}

interface ActionProps {
	finish: (index: number) => void
}

type Props = StateProp & ActionProps

const mapStateToProps = ({
	todo: { inProgress, finished },
}: IState): StateProp => ({
	inProgress,
	finished,
})

const mapDispatchToProps = (dispatch: any): ActionProps => ({
	finish: index => dispatch(FinishAction(index)),
})

const InProgress = styled.li`
	font-weight: bold;
`
const Finished = styled.li`
	text-decoration: line-through;
`

const component = ({ finish, finished, inProgress }: Props) => {
	return (
		<div>
			<ul>
				{inProgress.map((row, index) => (
					<InProgress key={index}>
						{row}{" "}
						<button type="button" onClick={() => finish(index)}>
							x
						</button>
					</InProgress>
				))}
			</ul>
			<ul>
				{finished.map((row, index) => (
					<Finished key={index}>{row}</Finished>
				))}
			</ul>
		</div>
	)
}

const container = connect(
	mapStateToProps,
	mapDispatchToProps
)(component)
export default container
