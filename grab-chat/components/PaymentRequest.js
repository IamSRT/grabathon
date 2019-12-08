import React from 'react';
import { View, StyleSheet } from 'react-native';
import {
  Card,
  Text,
  TouchableRipple,
  Button,
  Avatar
} from 'react-native-paper';
import { theme } from '../theme';
import AvatarHead from './AvatarHead';

const Payment = ({ isLeft }) => {
  return (
    <View
      style={{
        alignItems: isLeft ? 'flex-start' : 'flex-end',
        marginHorizontal: 15,
        marginVertical: 5,
        marginBottom: 15
      }}
    >
      <AvatarHead />

      <Card style={styles.cardStyle} elevation={0}>
        <View
          style={{
            marginBottom: 10
          }}
        >
          <Text style={styles.amount}>345</Text>
          <Text>Requested</Text>
        </View>

        <View style={styles.buttonContainer}>
          <Button
            mode="contained"
            onPress={() => console.log('Pressed')}
            style={styles.buttonStyle}
          >
            Accept
          </Button>
          <Button
            mode="Outlined"
            onPress={() => console.log('Pressed')}
            style={{...styles.buttonStyle, marginLeft: 10}}
          >
            Reject
          </Button>
        </View>
      </Card>
    </View>
  );
};

const styles = StyleSheet.create({
  cardStyle: {
    padding: 10,
    borderColor: theme.colors.primary,
    borderWidth: 1,
    borderStyle: 'dashed',
    width: 250
  },
  amount: {
    fontSize: 60,
    fontWeight: '400'
  },
  buttonContainer: {
    flexDirection: 'row',
    alignItems: 'center'
  },
  buttonStyle: {
    borderRadius: 50
  },
  avatarStyle: {
    marginBottom: 10
  }
});

export default Payment;
