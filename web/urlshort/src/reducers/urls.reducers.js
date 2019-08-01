import types from '../actions/types';

const reducers = {
  [types.URLS_SET](state, action) {
    return action.payload;
  }
};

export default reducers;
