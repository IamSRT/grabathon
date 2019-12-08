import React from 'react';
import { View, StyleSheet } from 'react-native';
import { Text, Avatar } from 'react-native-paper';
import { theme } from '../theme';
import LottieLoader from './LottieLoader';

const ChatHead = ({ title, isFromBot, hideAvatar }) => {
  return (
    <View
      style={{
        alignItems: isFromBot ? 'flex-start' : 'flex-end',
        marginHorizontal: 15,
        marginVertical: 5
      }}
    >
      {!hideAvatar && (
        <Avatar.Icon size={24} icon="robot" style={styles.avatarStyle} />
      )}
      <View style={styles.container}>
        <Text>{title}</Text>
      </View>
      {/* <LottieLoader animateFilePath={require('../assets/animators/chatLoader.json')}/> */}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    borderRadius: 50,
    borderWidth: 1,
    borderColor: theme.colors.background,
    paddingVertical: 10,
    paddingHorizontal: 15,
    maxWidth: 300
  },
  avatarStyle: {
    marginBottom: 10
  }
});

export default ChatHead;
