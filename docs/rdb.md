# Redis Database

## Single instance for authentication layer

This reduces complexity. Plus, exploitation of any one of the authenticating databases means exploitation of all, so there is no security reason for multiple instances.

## Prefixes

The first letter of a key is the namespace, so make sure it is 1 byte in length.

- s: session
- m: sms to wait
- M: sms received
- u: a user's sessions

## Data Structures

session: token -> person, person is a zero value if not authenticated.

m: (tel, msg) -> token, expire after 10min

M: (tel, msg) -> phone, expire after 1min

user sessions: person -> (tokenA, tokenB, ...)
