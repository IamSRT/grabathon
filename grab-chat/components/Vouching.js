import React from 'react';
import { View, StyleSheet } from 'react-native';
import { Avatar, Card, Text } from 'react-native-paper';
import { theme } from '../theme';
import { FlatList } from 'react-native-gesture-handler';
import AvatarHead from './AvatarHead';

const usersList = [
  {
    name: 'Prudhvi',
    id: 1
  },
  {
    name: 'Adarsh',
    id: 2
  },
  {
    name: 'Swayam',
    id: 3
  },
  {
    name: 'Sai',
    id: 4
  }
];

const Vouching = ({ isLeft }) => {
  const list = React.useMemo(() => {
    return usersList.map(item => {
      return {
        ...item,
        bgColor: randomRgb()
      };
    });
  }, [usersList]);

  return (
    <View
      style={{
        flex: 1,
        alignItems: isLeft ? 'flex-start' : 'flex-end',
        marginHorizontal: 15,
        marginVertical: 5,
        marginBottom: 15
      }}
    >
      <AvatarHead />
      <Card elevation={0} style={styles.cardStyle}>
        <View style={styles.avatarContainer}>
          {list.map((item, index) => {
            return (
              <Avatar.Text
                key={index}
                size={24}
                label={item.name.substr(0, 2)}
                color={'#fff'}
                style={{ ...styles.avatarStyle, backgroundColor: item.bgColor }}
              />
            );
          })}
        </View>
        <View style={styles.textContainer}>
          <Text style={styles.amount}>350</Text>
          <Text>Vouche Requested Amount</Text>
        </View>
      </Card>
    </View>
  );
};

const randomRgb = () => {
  const red = Math.floor(Math.random() * 256);
  const green = Math.floor(Math.random() * 256);
  const blue = Math.floor(Math.random() * 256);
  return `rgb(${red},${green},${blue})`;
};

const styles = StyleSheet.create({
  cardStyle: {
    borderColor: theme.colors.background,
    borderWidth: 1,
    borderStyle: 'solid',
    width: 250,
    overflow: 'hidden'
  },
  avatarContainer: {
    flex: 1,
    flexDirection: 'row',
    flexWrap: 'wrap',
    backgroundColor: theme.colors.background,
    padding: 10
  },
  avatarStyle: {
    marginRight: 5,
    marginBottom: 5
  },
  textContainer: {
    paddingHorizontal: 10,
    paddingBottom: 5
  },
  amount: {
    fontSize: 60,
    fontWeight: '400'
  }
});

export default Vouching;
