import actionCreatorFactory from 'typescript-fsa';
import { reducerWithInitialState } from 'typescript-fsa-reducers';

export interface IState {
	formula: string;
	input: string;
	tab: string;
	result: any;
}

const initalState: IState = {
	formula: '(v) => v',
	input: '',
	tab: 'text',
	result: null,
};

const actionCreator = actionCreatorFactory();

export const InputFormulaAction = actionCreator<string>('LAMBDA_INPUT_FORMULA');
export const TabChangeAction = actionCreator<string>('INPUT_TAB_CHANGE');
export const TextChangeAction = actionCreator<string>('INPUT_TEXT_CHANGE');
export const CalcResultAction = actionCreator('OUTPUT_CALC_RESULT');

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
	.case(InputFormulaAction, (state, formula) => {
		return {
			...state,
			formula,
		};
	})
	.case(TabChangeAction, (state, tab) => {
		return {
			...state,
			tab,
		};
	})
	.case(TextChangeAction, (state, input) => ({
		...state,
		input,
	}))
	.case(CalcResultAction, state => {
		return {
			...state,
			result: calcResult(state.result, state.formula, state.input),
		};
	})
	.default(state => state);
