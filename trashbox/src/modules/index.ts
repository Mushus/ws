import { IState as todoState, reducer as todoReducer } from '~/modules/todo';
import { combineReducers } from 'redux';

export interface IState {
	todo: todoState;
}

export const reducer = combineReducers<IState>({
	todo: todoReducer,
});
