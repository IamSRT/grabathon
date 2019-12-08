import React, { useState, useContext } from 'react';
import { View, StyleSheet, TextInput, TouchableOpacity } from 'react-native';
import { theme } from '../theme';
import { Avatar, TouchableRipple, IconButton, Text } from 'react-native-paper';
import { Context as PaymentContext } from '../context/PaymentContext';
import { Context as UserContext } from '../context/UserContext';

const MESSAGES = {
  GET_MONEY: 'Amount to Request',
  SEND_MONEY: 'Amount to Send'
};

const PaymentScreen = ({ navigation }) => {
  const [amount, setAmount] = useState('');
  const renderType = navigation.getParam('renderType');
  const id = navigation.getParam('id');
  const isSubmitNow = navigation.getParam('submitNow');
  const message = MESSAGES[renderType.toUpperCase()];

  const { createConversation } = useContext(PaymentContext);
  const { state: userState } = useContext(UserContext);

  const requestData = {
    amount: Number(amount),
    payment: {
      type: 'PAYMENT',
      status: 'COMPLETED',
      transactions: [
        {
          destination_id: userState.phone_number,
          transaction_type: 'LOAD',
          status: 'COMPLETED',
          amount: Number(amount)
        }
      ]
    }
  };

  return (
    <View
      style={{
        flex: 1,
        backgroundColor: theme.colors.primary,
        ...StyleSheet.absoluteFillObject
      }}
    >
      <View style={{ paddingTop: 40 }}></View>
      <IconButton
        icon="arrow-left"
        color={'#fff'}
        size={30}
        onPress={() => navigation.navigate('Home')}
      />

      <View style={styles.amountContainer}>
        {message && (
          <Text
            style={{
              fontSize: 15,
              fontWeight: '500',
              color: '#fff',
              textAlign: 'center'
            }}
          >
            {message}
          </Text>
        )}
        <TextInput
          //   selectionColor={'#fff'}
          placeholder="0"
          value={amount}
          onChangeText={val => setAmount(val)}
          style={styles.inputStyle}
          autoFocus
          autoCapitalize="none"
          autoCorrect={false}
          keyboardType={'number-pad'}
          maxLength={5}
          onEndEditing={() => {
            isSubmitNow
              ? (createConversation(userState.phone_number, id, {
                  ...requestData,
                  messages: {
                    options: [requestData],
                    messages: [{ message: '50', messageType: 'text' }]
                  }
                }),
                navigation.navigate('Home'))
              : navigation.navigate('Contact', {
                  amount: amount,
                  renderType,
                  id
                });
          }}
        />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  inputStyle: {
    backgroundColor: theme.colors.primary,
    // borderWidth: 12,
    color: '#fff',
    fontSize: 50,
    fontWeight: '600',
    textAlign: 'center'
  },
  amountContainer: {
    height: 400,
    justifyContent: 'center'
  }
});

export default PaymentScreen;
