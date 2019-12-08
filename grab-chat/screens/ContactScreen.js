import React, { useContext, useEffect, useState } from 'react';
import {
  View,
  FlatList,
  ScrollView,
  StyleSheet,
  TextInput
} from 'react-native';
import { Context as ContactContext } from '../context/ContactContext';
import {
  Text,
  TouchableRipple,
  Avatar,
  Button,
  Divider
} from 'react-native-paper';
import { Checkbox } from 'react-native-paper';

import { theme } from '../theme';
import { Context as PaymentContext } from '../context/PaymentContext';
import { isEmpty, GetNestedValue } from '../utils';
import { Context as UserContext } from '../context/UserContext';

const ContactScreen = ({ navigation }) => {
  const [phoneNumber, setPhoneNumber] = useState('');
  const [isMultiSelect, setSelectionType] = useState(() => {
    return navigation.getParam('renderType') === 'MULTIPLE_CONTACT_CARD';
  });
  const [selectedList, setSelectedList] = useState({});

  const { state, getContactList } = useContext(ContactContext);
  const { data } = state;

  const { state: paymentState, createVouche, createConversation } = useContext(
    PaymentContext
  );
  const { state: userState } = useContext(UserContext);

  useEffect(() => {
    if (data && data.length) return;
    getContactList();
    return () => {};
  }, [data]);

  const createVoucheApi = () => {
    if (!isMultiSelect) {
      let number = null;
      if (phoneNumber) {
        number = phoneNumber;
      } else if (!isEmpty(selectedList)) {
        number = GetNestedValue(
          Object.values(selectedList),
          '0.phoneNumbers0.number',
          null
        );
      }
      console.log(number);
      if (!number) return;

      const request = {
        amount: Number(navigation.getParam('amount')),
        payment: {
          type: 'PAYMENT',
          status: 'COMPLETED',
          transactions: [
            {
              source_id: userState.phone_number,
              destination_id: number,
              transaction_type: 'PAYMENT',
              status: 'COMPLETED',
              amount: Number(navigation.getParam('amount'))
            }
          ]
        }
      };

      createConversation(
        userState.phone_number,
        navigation.getParam('id'),
        request
      );
      navigation.navigate('Home');

      return;
    }

    if (isEmpty(selectedList)) return;

    const data = Object.entries(selectedList).map(([key, data]) => {
      if (!data) return null;
      return {
        voucher: data.phoneNumbers[0].number,
        amount: navigation.getParam('amount'),
        vouchee_id: userState.phone_number
      };
    });
    console.log(data);
    // createVouche(data);
    createConversation(userState.phone_number, navigation.getParam('id'), {
      vouches: data
    });
    navigation.navigate('Home');
  };

  const onSelect = item => {
    if (!isMultiSelect) {
      setSelectedList({ [item.id]: selectedList[item.id] ? null : item });
      return;
    }
    if (selectedList[item.id]) {
      setSelectedList({ ...selectedList, [item.id]: null });
      return;
    }
    setSelectedList({ ...selectedList, [item.id]: item });
  };

  onNumberChange = val => {
    setPhoneNumber(val);
    !isEmpty(selectedList) && setSelectedList({});
  };

  if (!data) return null;
  return (
    <View
      style={{
        flex: 1,
        backgroundColor: '#fff',
        ...StyleSheet.absoluteFillObject
      }}
    >
      {!isMultiSelect && (
        <>
          <TextInput
            onChangeText={onNumberChange}
            value={phoneNumber}
            placeholder="Enter Phone Number"
            keyboardType="number-pad"
            style={{
              backgroundColor: '#fff',
              margin: 10,
              marginLeft: 15,
              marginBottom: 15
            }}
          />
          <Divider />
        </>
      )}

      <Text
        style={{
          fontSize: 24,
          fontWeight: '900',
          margin: 15
        }}
      >
        Select Contact
      </Text>
      <ScrollView
        style={{
          height: 600
        }}
      >
        <FlatList
          windowSize={2}
          initialNumToRender={2}
          maxToRenderPerBatch={10}
          data={data}
          keyExtractor={data => data.id}
          extraData={selectedList}
          renderItem={({ item }) => {
            return (
              <ContactCard
                item={item}
                checked={!!selectedList[item.id]}
                onSelect={onSelect}
                isMultiSelect={isMultiSelect}
              />
            );
          }}
        />
      </ScrollView>
      <View style={styles.actionContainer}>
        <Button
          mode="Outlined"
          onPress={() => navigation.navigate('Home')}
          style={styles.buttonStyle}
        >
          Cancel
        </Button>
        <Button
          mode="contained"
          onPress={() => {
            createVoucheApi();
            // navigation.navigate('Home');
          }}
          style={styles.buttonStyle}
        >
          Continue
        </Button>
      </View>
    </View>
  );
};

const ContactCard = React.memo(({ item, checked, onSelect, isMultiSelect }) => {
  const phoneDetails = item.phoneNumbers && item.phoneNumbers[0];
  return (
    <TouchableRipple
      onPress={() => {
        onSelect(item);
      }}
    >
      <View style={styles.cardContainer}>
        <View style={styles.nameContainer}>
          <Avatar.Icon
            size={30}
            icon="account"
            style={{ marginRight: 10, alignSelf: 'center' }}
          />
          <View style={{ justifyContent: 'center' }}>
            <Text style={styles.nameStyle}>{item.name}</Text>
            <Text>{phoneDetails.number}</Text>
          </View>
        </View>
        {isMultiSelect ? (
          <Checkbox
            status={checked ? 'checked' : 'unchecked'}
            onPress={() => {
              // this.setState({ checked: !checked });
            }}
          />
        ) : (
          <Avatar.Icon
            icon="check"
            size={24}
            style={{
              backgroundColor: checked
                ? theme.colors.primary
                : theme.colors.background
            }}
          />
        )}
      </View>
    </TouchableRipple>
  );
});

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
    position: 'absolute',
    bottom: 0,
    flexDirection: 'row',
    backgroundColor: '#fff'
  },
  buttonStyle: {
    flex: 1,
    paddingVertical: 10,
    // margin: 10,
    borderRadius: 0
  }
});

export default ContactScreen;
