import actionCreatorFactory from "typescript-fsa"
import { reducerWithInitialState } from "typescript-fsa-reducers"

export interface IState {
	text: string
	inProgress: string[]
	finished: string[]
}

const initalState: IState = {
	finished: [],
	inProgress: [],
	text: "",
}

const actionCreator = actionCreatorFactory()

export const InputAction = actionCreator<string>("Input")
export const AddAction = actionCreator<void>("ADD")
export const FinishAction = actionCreator<number>("FINISH")

export const reducer = reducerWithInitialState<IState>(initalState)
	.case(InputAction, (state, text) => ({
		...state,
		text,
	}))
	.case(AddAction, state => ({
		...state,
		inProgress: [...state.inProgress, state.text],
		text: "",
	}))
	.case(FinishAction, (state, index) => {
		const inProgress = [...state.inProgress]
		const finished = [...state.finished, state.inProgress[index]]
		inProgress.splice(index, 1)
		return {
			...state,
			finished,
			inProgress,
		}
	})
	.default(state => state)
