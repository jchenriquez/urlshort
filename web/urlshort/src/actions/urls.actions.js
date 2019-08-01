import axios from 'axios';
import types from './types';

const setUrl = (short, long) => dispatch => {
  axios.post('/admin/urls', { short, long }).then(function(response) {
    dispatch({
      type: types.URLS_SET,
      payload: response.data
    });
  });
};

const removeUrl = (short, long) => dispatch => {
  axios
    .delete('admin/urls', { params: { short, long } })
    .then(function(response) {
      dispatch({
        type: types.URLS_SET,
        payload: response.data
      });
    });
};

const getUrls = dispatch => {
  axios.get('admin/urls').then(function(response) {
    dispatch({
      type: types.URLS_SET,
      payload: response.data
    });
  });
};

export default {
  setUrl,
  removeUrl,
  getUrls
};
