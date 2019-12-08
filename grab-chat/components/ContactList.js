import React, { useState, useEffect } from 'react';
import { View, StyleSheet, ScrollView } from 'react-native';
import { Text } from 'react-native-paper';
import { theme } from '../theme';
import { Portal, Modal } from 'react-native-paper';
import * as Permissions from 'expo-permissions';
import * as Contacts from 'expo-contacts';
import { FlatList } from 'react-native-gesture-handler';

const ContactList = ({ title, isFromBot }) => {
  const [results, setResults] = useState([]);

  const requestForContacts = async () => {
    try {
      const response = await Permissions.askAsync(Permissions.CONTACTS);
      return response;
    } catch (e) {
      throw e;
    }
  };
  useEffect(() => {
    requestForContacts().then(async res => {
      const { data } = await Contacts.getContactsAsync({
        fields: [Contacts.Fields.PhoneNumbers]
      });

      setResults(data.filter((item) => Array.isArray(item.phoneNumbers)));
    });
  }, []);

  return (
    <Portal>
      <Modal visible>
        <View style={{ backgroundColor: '#fff', margin: 30 }}>
          <ScrollView
            style={{
              height: 600
            }}
          >
            <FlatList
              data={results}
              keyExtractor={(data) => data.id}
              renderItem={({ item }) => {
                const phoneDetails = item.phoneNumbers && item.phoneNumbers[0];
                return <Text>{phoneDetails.number}</Text>;
              }}
            />
          </ScrollView>
        </View>
      </Modal>
    </Portal>
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

export default ContactList;
