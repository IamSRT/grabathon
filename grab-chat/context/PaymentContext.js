import createDataContext from './createDataContext';
import * as Permissions from 'expo-permissions';
import * as Contacts from 'expo-contacts';
import steps from '../api/steps';
import { GetNestedValue, ParseJSON } from '../utils';
import user, { user_url } from '../api/user';
import conversations from '../api/conversations';
import { _storeData, _retrieveData } from '../utils/localStorage';

const initialState = {
  conversations: [],
  messages: []
};

const PaymentReducer = (state, action) => {
  switch (action.type) {
    case 'SET_STEPS':
      return {
        ...state,
        loading: false,
        error: null,
        messages: [...state.messages, action.data]
      };
    case 'setConversations':
      return {
        ...state,
        conversations: action.data
      };
    case 'error':
      return state;
  }
  return state;
};

const createVouche = dispatch => {
  return (list, amount, user_number) => {
    return fetch(`${user_url}/vouch`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        vouches: list
      })
    })
      .then(response => response.json())
      .then(responseJson => {
        console.log(responseJson);
      })
      .catch(error => {
        console.error(error);
      });
  };
};

const createConversation = dispatch => {
  return (user_number, modeId, data) => {
    const request = {
      title: 'Create Wallet',
      modeId: modeId || '0',
      messages: {
        options: [],
        messages: [{ message: '50', messageType: 'text' }]
      },
      ...(data || {})
    };
    console.log(request);
    conversations
      .post(
        '/conversations',
        { ...request },
        {
          headers: {
            userid: user_number,
            'Content-Type': 'application/json'
          }
        }
      )
      .then(res => {
        console.log(res.data);
        const data = GetNestedValue(res, 'data', []);
        dispatch({
          type: 'SET_STEPS',
          data
        });
        _retrieveData('messages').then(res => {
          const storedData = ParseJSON(res);
          if (storedData) {
            _storeData('messages', JSON.stringify([...storedData, data]));
          }
        });
        _storeData('messages');
      })
      .catch(e => {
        console.error(e);
      });
  };
};

const getConversations = dispatch => {
  return user_number => {
    console.log('user number', user_number);
    conversations
      .get(`/conversations/${user_number}`, {
        headers: {
          userid: user_number,
          'Content-Type': 'application/json'
        }
      })
      .then(res => {
        console.log('>>>> conversations', res.data);
        console.log(
          '>>> conversion',
          GetNestedValue(res, 'data.data.messageHistory', [])
        );
        dispatch({
          type: 'setConversations',
          data: GetNestedValue(res, 'data.data.messageHistory', [])
        });
      })
      .catch(e => {
        console.error(e);
      });
  };
};

const getAutoPayEnableStatus = () => {
  user
    .get('/is-auto-pay-enabled')
    .then(() => {})
    .catch(() => {});
};

export const { Context, Provider } = createDataContext(
  PaymentReducer,
  {
    createVouche,
    createConversation,
    getConversations
  },
  initialState
);
