import { combineReducers } from 'redux';
import reducers from './urls.reducers';
import * as _ from 'lodash';

function reducerRouter(state = [], action) {
  if (!_.has(reducers, action.type)) {
    return state;
  }
  return reducers[action.type](state, action);
}

export default combineReducers({ urls: reducerRouter });
