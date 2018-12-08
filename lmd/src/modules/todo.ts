import actionCreatorFactory from 'typescript-fsa';
import { reducerWithInitialState } from 'typescript-fsa-reducers';

export interface IState {
	lambda: string;
	text: string;
	result: any;
}

const initalState: IState = {
	lambda: '(v) => v',
	text: '',
	result: '',
};

const actionCreator = actionCreatorFactory();

export const InputLambda = actionCreator<string>('INPUT_LAMBDA');
export const InputAction = actionCreator<string>('Input');
export const AddAction = actionCreator<void>('ADD');
export const FinishAction = actionCreator<number>('FINISH');

/**
 * 計算する
 * @param result 以前の結果
 * @param lambda 計算式
 */
function calcResult(result: any, lambda: string, input: any): any {
	try {
		return eval(`${lambda}`)(input);
	} catch (e) {
		return result;
	}
}

export const reducer = reducerWithInitialState<IState>(initalState)
	.case(InputLambda, (state, lambda) => {
		return {
			...state,
			lambda,
			result: calcResult(state.result, lambda, state.text),
		};
	})
	.case(InputAction, (state, text) => ({
		...state,
		text,
		result: calcResult(state.result, state.lambda, text),
	}))
	.default(state => state);
