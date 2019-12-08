import React from 'react';
import { View, StyleSheet } from 'react-native';
import {
  Text,
  Avatar
} from 'react-native-paper';
import { theme } from '../theme';

const AvatarHead = () => {
  return (
    <View style={{ flexDirection: 'row', marginBottom: 10 }}>
      <Avatar.Icon size={24} icon="account" style={{ marginRight: 10 }} />
      <Text>Merchant</Text>
    </View>
  );
};

const styles = StyleSheet.create({});

export default AvatarHead;
