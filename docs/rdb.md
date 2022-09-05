# Redis Database

## Single instance for authentication layer

This reduces complexity. Plus, exploitation of any one of the authenticating databases means exploitation of all, so there is no security reason for multiple instances.

## Prefixes

The first letter of a key is the namespace, so make sure it is 1 byte in length.

- s: session
- m: sms
- u: user sessions

## Data Structures

session: token -> person, person is a zero value if not authenticated.

sms: (tel, msg) -> phone, phone is a zero value if sms not received.

user sessions: person -> (tokenA, tokenB, ...)
