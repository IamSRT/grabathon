import createDataContext from './createDataContext';
import user from '../api/user';
import { GetNestedValue } from '../utils';

import { _storeData, _retrieveData, _removeData } from '../utils/localStorage';

const initialState = {
  name: '',
  phone_number: '',
  city: '',
  email: '',
  auto_pay: false,
  loggedIn: false,
  checkingLoginDetails: true,
  detailsExist: false
};

const contactReducer = (state, action) => {
  switch (action.type) {
    case 'setUserDetails':
      return {
        ...state,
        ...(action.payload || {}),
        loading: false,
        error: null,
        detailsExist: true
      };
    case 'error':
      return { ...initialState };
    case 'loginCheck':
      return {
        ...state,
        checkingLoginDetails: action.payload
      };
    case 'setLogIn':
      return {
        ...state,
        loggedIn: action.payload,
        checkingLoginDetails: false
      };
    case 'reset':
      return { ...initialState };
    default:
      return state;
  }
};

const getUserDetails = dispatch => {
  return (id, _callback) => {
    return user
      .get(`/${id}`)
      .then(res => {
        const data = GetNestedValue(res, 'data.data', {});
        dispatch({
          type: 'setUserDetails',
          payload: data
        });
        _storeData('userId', id).then(() => {
          dispatch({
            type: 'setLogIn',
            payload: true
          });
          _callback && _callback();
        });
      })
      .catch(err => {
        dispatch({
          type: 'error'
        });
      });
  };
};

const checkIsUserLoggedIn = (dispatch, state) => {
  return () => {
    dispatch({
      type: 'loginCheck',
      payload: true
    });
    _retrieveData('userId')
      .then(val => {
        if (!val) throw 'error';
        return getUserDetails(dispatch)(val);
      })
      .catch(e => {})
      .finally(() => {
        dispatch({
          type: 'loginCheck',
          payload: false
        });
      });
  };
};

const logout = dispatch => {
  return _callback => {
    _removeData('userId').then(() => {
      dispatch({
        type: 'reset'
      });
      _callback && _callback();
    });
  };
};

export const { Context, Provider } = createDataContext(
  contactReducer,
  {
    getUserDetails,
    checkIsUserLoggedIn,
    logout
  },
  initialState
);
