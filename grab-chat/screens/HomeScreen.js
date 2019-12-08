import * as WebBrowser from 'expo-web-browser';
import React, { useRef, useEffect, useContext } from 'react';
import {
  Image,
  Platform,
  ScrollView,
  StyleSheet,
  TouchableOpacity,
  View,
  Button
} from 'react-native';

// import SlidingUpPanel from 'rn-sliding-up-panel';
import SelectMode from '../components/SelectMode';
import { InitialSelection } from '../constants/ModesList';
import ChatHead from '../components/ChatHead';
import Welcome from '../components/Welcome';
import ContactList from '../components/ContactList';
import Payment from '../components/PaymentRequest';
import { Context as ContactContext } from '../context/ContactContext';
import Vouching from '../components/Vouching';
import { Context as UserContext } from '../context/UserContext';
import {
  ActivityIndicator,
  Colors,
  Card,
  Text,
  Avatar
} from 'react-native-paper';
import { _removeData } from '../utils/localStorage';
import { theme } from '../theme';
import { Context as PaymentContext } from '../context/PaymentContext';
import { isEmpty, GetNestedValue } from '../utils';
import { FlatList } from 'react-native-gesture-handler';

export default function HomeScreen({ navigation }) {
  const slideRef = useRef();
  const scrollRef = useRef();

  const contactStore = useContext(ContactContext);
  const { checkIsUserLoggedIn, state: userState } = useContext(UserContext);
  const {
    getSteps,
    state: paymentStore,
    createConversation,
    getConversations
  } = useContext(PaymentContext);
  const {
    checkingLoginDetails,
    loggedIn,
    detailsExist,
    name: username,
    phone_number
  } = userState;
  const { getContactList } = contactStore;

  const { options, messages, conversations } = paymentStore;

  useEffect(() => {
    getContactList();
  }, []);

  useEffect(() => {
    if (loggedIn || detailsExist) return;
    // _removeData('userId').then(() => {
    //   checkIsUserLoggedIn();
    // });
    checkIsUserLoggedIn();
  }, [loggedIn, detailsExist]);

  useEffect(() => {
    if (!phone_number) return;
    // getConversations(phone_number);
  }, [phone_number]);

  useEffect(() => {
    if (!isEmpty(options)) return;
    if (loggedIn) {
      // getSteps();
      isEmpty(messages) && createConversation(phone_number);
    }
  }, [loggedIn, options, messages]);

  if (checkingLoginDetails)
    return (
      <View style={{ ...styles.container, justifyContent: 'center' }}>
        <ActivityIndicator animating={true} color={Colors.red800} />
      </View>
    );

  return (
    <View style={styles.container}>
      <ScrollView
        style={styles.container}
        contentContainerStyle={styles.contentContainer}
        ref={scrollRef}
        onContentSizeChange={(contentWidth, contentHeight) => {
          scrollRef.current.scrollToEnd({ animated: true });
        }}
      >
        <Welcome name={loggedIn ? username : null} />

        {loggedIn ? (
          <React.Fragment>
            <>
              {!isEmpty(conversations) &&
                conversations.map((msg, idx) => {
                  const requester = GetNestedValue(msg, 'messages', []);

                  const requesterOptions = GetNestedValue(msg, 'options', []);

                  return (
                    <View key={idx}>
                      {/* {!isEmpty(requestsOptions) && (
                    <SelectMode
                      data={requestsOptions}
                      onSelect={() => {
                        navigation.navigate('Payment');
                      }}
                      checked={null}
                      createConversation={createConversation}
                    />
                  )}
                  {!isEmpty(requests) && (
                    <>
                      {requests.map((item, index) => {
                        return (
                          <ChatHead
                            title={item.message}
                            isFromBot
                            key={index}
                          />
                        );
                      })}
                    </>
                  )} */}
                      {!isEmpty(requester) && !isEmpty(requesterOptions) && (
                        <>
                          {requester.map((item, index) => {
                            return (
                              <ChatHead
                                title={item.message}
                                isFromBot
                                key={index}
                              />
                            );
                          })}
                        </>
                      )}
                      {!isEmpty(requesterOptions) && (
                        <SelectMode
                          data={requesterOptions}
                          onSelect={() => {
                            navigation.navigate('Payment');
                          }}
                          checked={null}
                          createConversation={createConversation}
                          disable
                        />
                      )}
                      {!isEmpty(requester) && isEmpty(requesterOptions) && (
                        <>
                          {requester.map((item, index) => {
                            return (
                              <ChatHead
                                title={item.message}
                                isFromBot
                                key={index}
                              />
                            );
                          })}
                        </>
                      )}
                    </View>
                  );
                })}
            </>
            <>
              {messages.map((msg, idx) => {
                const requester = GetNestedValue(
                  msg,
                  'data.requester.messages',
                  []
                );

                const requesterOptions = GetNestedValue(
                  msg,
                  'data.requester.options',
                  []
                );

                return (
                  <View key={idx}>
                    {/* {!isEmpty(requestsOptions) && (
                    <SelectMode
                      data={requestsOptions}
                      onSelect={() => {
                        navigation.navigate('Payment');
                      }}
                      checked={null}
                      createConversation={createConversation}
                    />
                  )}
                  {!isEmpty(requests) && (
                    <>
                      {requests.map((item, index) => {
                        return (
                          <ChatHead
                            title={item.message}
                            isFromBot
                            key={index}
                          />
                        );
                      })}
                    </>
                  )} */}
                    {!isEmpty(requester) && !isEmpty(requesterOptions) && (
                      <>
                        {requester.map((item, index) => {
                          return (
                            <ChatHead
                              title={item.message}
                              isFromBot
                              key={index}
                            />
                          );
                        })}
                      </>
                    )}
                    {!isEmpty(requesterOptions) && (
                      <SelectMode
                        data={requesterOptions}
                        onSelect={() => {
                          navigation.navigate('Payment');
                        }}
                        checked={null}
                        createConversation={createConversation}
                      />
                    )}
                    {!isEmpty(requester) && isEmpty(requesterOptions) && (
                      <>
                        {requester.map((item, index) => {
                          return (
                            <ChatHead
                              title={item.message}
                              isFromBot
                              key={index}
                            />
                          );
                        })}
                      </>
                    )}
                  </View>
                );
              })}

              {/* <ChatHead title={'HELLO'} isFromBot />
            <ChatHead title={'HELLO'} isFromBot={false} />
            <ChatHead title={'HELLLO this is prudhvi'} isFromBot={false} />
            <ChatHead title={'HELLO'} isFromBot /> */}
              {/* <SelectMode
              data={InitialSelection}
              onSelect={() => {}}
              checked={null}
            /> */}
              {/* <Payment isLeft={true} />
            <Payment isLeft={true} />
            <Vouching isLeft /> */}
              <View style={{ paddingBottom: 40 }}></View>
            </>
          </React.Fragment>
        ) : (
          <LoginFlow navigation={navigation} />
        )}
      </ScrollView>
      {/* <ContactList /> */}
      {loggedIn && (
        <View
          style={{
            position: 'absolute',
            bottom: 20,
            left: 15
          }}
        >
          <TouchableOpacity onPress={() => navigation.navigate('Links')}>
            <Avatar.Icon icon={'settings'} size={30} />
          </TouchableOpacity>
        </View>
      )}
    </View>
  );
}

const LoginFlow = ({ navigation }) => {
  return (
    <View>
      <ChatHead
        title={
          "Hey, looks like you are not loggedIn or haven't Created a wallet yet 🧐"
        }
        isFromBot
      />
      <ChatHead
        title={
          "Don't Worry!! You can choose any of the options below to continue 😊"
        }
        isFromBot
        hideAvatar
      />
      <Card style={styles.loginCardContainer} elevation={5}>
        <View style={styles.loginActions}>
          <TouchableOpacity onPress={() => navigation.navigate('CreateUser')}>
            <Text style={styles.loginActionText}>
              I want to create a wallet
            </Text>
          </TouchableOpacity>
        </View>
        <View
          style={{
            ...styles.loginActions,
            borderBottomWidth: 0
          }}
        >
          <TouchableOpacity onPress={() => navigation.navigate('Login')}>
            <Text style={styles.loginActionText}>I have an account</Text>
          </TouchableOpacity>
        </View>
      </Card>
    </View>
  );
};

HomeScreen.navigationOptions = {
  header: null
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff'
  },
  developmentModeText: {
    marginBottom: 20,
    color: 'rgba(0,0,0,0.4)',
    fontSize: 14,
    lineHeight: 19,
    textAlign: 'center'
  },
  contentContainer: {
    paddingTop: 40
  },
  welcomeContainer: {
    alignItems: 'center',
    marginTop: 10,
    marginBottom: 20
  },
  loginCardContainer: {
    padding: 15,
    margin: 15
  },
  loginActions: {
    // marginBottom: 10,
    paddingVertical: 15,
    alignItems: 'center',
    borderBottomColor: theme.colors.background,
    borderBottomWidth: 1
  },
  loginActionText: {
    color: '#5d8ed5'
  }
});
