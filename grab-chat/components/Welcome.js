import React from 'react';
import { View, StyleSheet } from 'react-native';
import { Text, Avatar } from 'react-native-paper';
import { theme } from '../theme';

const Welcome = ({ name }) => {
  return (
    <View style={styles.container}>
      <Text style={styles.titleStyle}>What's up? {name ? name : ''}</Text>
      <Avatar.Image size={150} style={styles.avatarStyle} source={require('../assets/images/download.png')}/>
      <Text style={styles.subText}>No worries.</Text>
      <Text style={styles.subText}>
        I'll take care of all the payments for you!
      </Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
    marginBottom: 20
  },
  titleStyle: {
    fontSize: 25,
    fontWeight: '600',
    marginBottom: 10
  },
  avatarStyle: {
    backgroundColor: '#fff',
    // borderColor: theme.colors.background,
    // borderWidth: 1,
    marginBottom: 10
  },
  subText: {
    fontWeight: '300',
    color: theme.colors.backdrop
  }
});

export default Welcome;
