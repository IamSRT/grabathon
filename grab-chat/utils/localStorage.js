import { AsyncStorage } from 'react-native';

export const _storeData = async (key, data) => {
  try {
    await AsyncStorage.setItem(key, data);
  } catch (error) {
    // Error saving data
  }
};

export const _retrieveData = async key => {
  try {
    const value = await AsyncStorage.getItem(key);
    if (value !== null) {
      // We have data!!
      console.log(value);
    }
    return value;
  } catch (error) {
    // Error retrieving data
  }
};


export const _removeData = async (key) => {
    try {
      await AsyncStorage.removeItem(key);
    } catch (error) {
      // Error saving data
    }
  };