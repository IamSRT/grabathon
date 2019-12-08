import React, { useContext } from 'react';
import { View, StyleSheet, FlatList } from 'react-native';
import { RadioButton, Card, TouchableRipple, Text } from 'react-native-paper';
import { withNavigation } from 'react-navigation';

import { theme } from '../theme';
import { Context } from '../context/UserContext';

const PAGES = {
  MULTIPLE_CONTACT_CARD: 'Contact',
  LOAD_MONEY: 'Payment'
};

const SelectMode = ({
  data,
  onSelect,
  checked,
  navigation,
  createConversation,
  disable
}) => {
  const { state: userState } = useContext(Context);
  const { phone_number } = userState;

  return (
    <View>
      <View style={styles.cardTitleContainer}>
        {/* <Text style={styles.cardTitle}>Select Mode</Text> */}
      </View>
      <Card style={styles.mainContainer} elevation={6}>
        <FlatList
          data={data}
          keyExtractor={data => data.title}
          renderItem={({ item, index }) => {
            return (
              <ItemDetail
                item={item}
                onSelect={onSelect}
                checked={checked}
                lastIndex={index === data.length - 1}
                navigation={navigation}
                createConversation={createConversation}
                phone_number={phone_number}
                disable={disable}
              />
            );
          }}
        />
      </Card>
    </View>
  );
};

const ItemDetail = ({
  onSelect,
  item,
  checked,
  lastIndex,
  navigation,
  phone_number,
  createConversation,
  disable
}) => {
  return (
    <TouchableRipple
      onPress={() => {
        if (disable) return;
        if (['BALANCE'].includes(item.renderType)) {
          createConversation(phone_number, item.id);
        } else {
          navigation.navigate(PAGES[item.renderType] || 'Payment', {
            renderType: item.renderType,
            id: item.id,
            submitNow: item.renderType === 'LOAD_MONEY'
          });
        }
      }}
      rippleColor={theme.colors.primary}
    >
      <View
        style={{
          ...styles.cardContainer,
          borderBottomWidth: lastIndex ? 0 : 1,
          flex: 1
        }}
      >
        <View>
          <Text>{item.title}</Text>
          <Text
            style={{
              fontSize: 12,
              color: '#bebfc5'
            }}
            // numberOfLines={2}
          >
            {item.description}
          </Text>
        </View>
      </View>
    </TouchableRipple>
  );
};

const styles = StyleSheet.create({
  mainContainer: {
    backgroundColor: '#fff',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.8,
    shadowRadius: 2,
    marginHorizontal: 15,
    marginVertical: 10,
    marginBottom: 15,
    overflow: 'hidden'
  },
  cardTitleContainer: {
    marginLeft: 15,
    marginVertical: 5
  },
  cardTitle: {
    fontSize: 16,
    fontWeight: '600'
  },
  cardContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    borderBottomWidth: 1,
    borderBottomColor: theme.colors.background,
    paddingVertical: 10,
    paddingHorizontal: 10
  }
});

export default withNavigation(SelectMode);
