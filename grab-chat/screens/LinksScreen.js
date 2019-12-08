import React, { useContext } from 'react';
import { ScrollView, StyleSheet } from 'react-native';
import { ExpoLinksView } from '@expo/samples';
import { Context as UserContext } from '../context/UserContext';
import { Button } from 'react-native-paper';

export default function LinksScreen({ navigation }) {
  const { logout } = useContext(UserContext);
  return (
    <ScrollView style={styles.container}>
      <Button onPress={() => logout(() => navigation.navigate('Home'))}>
        Logout
      </Button>
    </ScrollView>
  );
}

LinksScreen.navigationOptions = {
  title: 'Account   '
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff'
  }
});
