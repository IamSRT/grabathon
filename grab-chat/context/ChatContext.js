import createDataContext from './createDataContext';
import * as Permissions from 'expo-permissions';
import * as Contacts from 'expo-contacts';

const requestForContacts = async () => {
  try {
    const response = await Permissions.askAsync(Permissions.CONTACTS);
    return response;
  } catch (e) {
    throw e;
  }
};

const contactReducer = (state, action) => {
  switch (action.type) {
    case 'contact_list':
      return {
        data: action.payload,
        loading: false,
        error: null
      };
    case 'error':
      return state;
  }
  return state;
};

const getContactList = dispatch => {
  return () => {
    requestForContacts()
      .then(async res => {
        const { data } = await Contacts.getContactsAsync({
          fields: [Contacts.Fields.PhoneNumbers]
        });

        dispatch({
          type: 'contact_list',
          payload: data.filter(item => Array.isArray(item.phoneNumbers))
        });
      })
      .catch(err => {
        dispatch({
          type: 'error'
        });
      });
  };
};

export const { Context, Provider } = createDataContext(
  contactReducer,
  {
    getContactList
  },
  [
    {
      data: [],
      loading: false,
      error: null
    }
  ]
);
