import React from 'react';
import {
  createAppContainer,
  createSwitchNavigator,
  createStackNavigator
} from 'react-navigation';
import { Platform } from 'react-native';

import MainTabNavigator from './MainTabNavigator';
import HomeScreen from '../screens/HomeScreen';
import PaymentScreen from '../screens/PaymentScreen';
import ContactScreen from '../screens/ContactScreen';
import CreateUserScreen from '../screens/CreateUserScreen';
import LoginScreen from '../screens/LoginScreen';

const config = Platform.select({
  web: { headerMode: 'screen' },
  default: {}
});

export default createAppContainer(
  createSwitchNavigator({
    // You could add another route here for authentication.
    // Read more at https://reactnavigation.org/docs/en/auth-flow.html
    Home: HomeScreen,
    Main: MainTabNavigator,
    Payment: PaymentScreen,
    Contact: ContactScreen,
    Login: LoginScreen,
    CreateUser: CreateUserScreen
  })
);
