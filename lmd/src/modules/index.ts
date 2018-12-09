import { IState as MainState, reducer as mainReducer } from '@/modules/main';
import { combineReducers } from 'redux';

export interface IState {
	main: MainState;
}

export const reducer = combineReducers<IState>({
	main: mainReducer,
});
