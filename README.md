# Grabathon #
## Service Components ##


### Payment Middleware ###


Every request through has to be routed through this middleware which stores all conversations and messages exchanged with Bot, assuming a one-to-one relationship between user and the conversation. This service converts user actions in to a consolidated conversation
### Payment Assistant Bot ###

Smart reply service which replies to user’s requests and provides the next set of actions to be performed by the user. Initially we can have programmed workflows but in future this can be replaced with an AI Bot. 

### Wallet Service ###

Wallet Service owns responsibility of User Account Management. It owns following entities:

User

Wallet

Transactions

Vouches

It provides CRUD apis for above entities with validations. It provides APIs to manipulate User Account Settings and Perform Transactions.
### Interactive Mobile App ###


Currently Mobile App acts as the link between users and PayMate. App has rich features to perform intelligent actions returned from Payment Assistant Bot. It can have augmented features like text to voice and vice versa. 

## Tools used ##
### Languages/Frameworks ###
Payment Middleware:3 NodeJS/Javascript

Wallet Service: GoLang/Chi Framework

Payment Assistant Bot: Java / Spring Boot

App: React-Native 

### Databases ###
MySQL

MongoDB
