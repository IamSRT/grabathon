import React, { useContext, useReducer } from 'react';
import { View, StyleSheet } from 'react-native';

import { Button, Card, TextInput, Text, Avatar } from 'react-native-paper';

import { theme } from '../theme';
import { Context as UserContext } from '../context/UserContext';
import user from '../api/user';
import { GetNestedValue } from '../utils';
import { ScrollView, TouchableOpacity } from 'react-native-gesture-handler';

const reducer = (prevState, state) => ({ ...prevState, ...state });

const initialState = {
  name: '',
  phone_number: '',
  city: '',
  email: '',
  auto_pay: false,

  name: 'Prudhvi',
  phone_number: '6366903399',
  city: 'Bangalore',
  email: 'akira.prudhvi@gmail.com',
  auto_pay: false
};

const CreateUserScreen = ({ navigation }) => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { getUserDetails } = useContext(UserContext);

  const createUserApi = () => {
    const isValid = ['name', 'phone_number', 'city', 'email'].every(
      text => !!text
    );
    if (!isValid) return;
    user
      .post('/create', { ...state })
      .then(res => {
        const data = GetNestedValue(res, 'data.data', {});
        getUserDetails(data.phone_number, () => {
          navigation.navigate('Home');
        });
      })
      .catch(err => {});
  };

  return (
    <View
      style={{
        flex: 1,
        backgroundColor: '#fff',
        ...StyleSheet.absoluteFillObject
      }}
    >
      <ScrollView>
        <View
          style={{
            alignItems: 'center',
            marginTop: 20
          }}
        >
          <Text style={styles.titleStyle}>PayMate</Text>
          <Avatar.Image
            size={150}
            style={styles.avatarStyle}
            source={require('../assets/images/download.png')}
          />
        </View>

        <Card
          style={{
            justifyContent: 'center',
            margin: 30,
            padding: 10
          }}
        >
          <Input
            label="Name"
            value={state.name}
            onChangeText={name => dispatch({ name })}
          />
          <Input
            label="Phone Number"
            value={state.phone_number}
            onChangeText={phone_number => dispatch({ phone_number })}
            keyboardType={'number-pad'}
          />
          <Input
            label="Email"
            value={state.email}
            onChangeText={email => dispatch({ email })}
          />
          <Input
            label="City"
            value={state.city}
            onChangeText={city => dispatch({ city })}
          />
        </Card>

        <View
          style={{
            marginTop: 0,
            margin: 40
          }}
        >
          <Button
            onPress={() => createUserApi()}
            style={styles.buttonStyle}
            mode="contained"
          >
            Create Wallet
          </Button>

          <TouchableOpacity onPress={() => navigation.navigate('Login')}>
            <View style={{ marginTop: 30, alignItems: 'center' }}>
              <Text style={{ color: theme.colors.helperText }}>
                Already have an account? Sign In
              </Text>
            </View>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </View>
  );
};

const Input = ({ label, value, onChangeText, ...props }) => {
  return (
    <View style={styles.inputContainer}>
      <TextInput
        label={label}
        value={value}
        onChangeText={onChangeText}
        mode="outlined"
        style={{
          borderRadius: 0,
          borderColor: theme.colors.background
        }}
        {...props}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  cardContainer: {
    borderBottomWidth: 1,
    borderBottomColor: theme.colors.background,
    paddingVertical: 15,
    paddingHorizontal: 15,
    flexDirection: 'row',
    justifyContent: 'space-between'
  },
  nameContainer: {
    flexDirection: 'row'
  },
  nameStyle: {
    fontSize: 16,
    fontWeight: '500'
  },
  actionContainer: {
    // position: 'absolute',
    // bottom: 0,
    // flexDirection: 'row',
    // backgroundColor: '#fff'
  },
  buttonStyle: {
    borderRadius: 50
  },
  inputContainer: {
    marginBottom: 10
  },
  titleStyle: {
    fontSize: 20,
    fontWeight: '600',
    marginBottom: 10
  },
  avatarStyle: {
    backgroundColor: '#fff',
    borderColor: theme.colors.background,
    borderWidth: 1,
    marginBottom: 10
  }
});

export default CreateUserScreen;
