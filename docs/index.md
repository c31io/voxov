# VOxOV

## Core Services

### gate

API gateway with session control.

- Redis

### auth

SMS and trusted devices.

- authdb: PostgreSQL

### pay

Add balance.

- Upstreams: Alipay, Paypal
- paydb: PostgreSQL

### plan

Subscription and privacy control.

- plandb: Finance

### meme

Data backend and access control.

- Meta: MongoDB
- File: Object Storage

### gene

Logic backend.

- Stateless
- Stateful

### censor

Public content filter and private content spotlight.

- MQ: rabbitmq
- censordb: MongoDB

### staff

Privileged management.

## Arch

### Monolithic

### Istio
