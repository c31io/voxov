# PostgreSQL Database

## Person

- int64 pid: Person ID
- string phone: Phone number

- [](bytes dtoken, string name, string info) devices

- int64 hid: Human ID, zero value if not a human
- []string name(s): a list of used names
- []string identity: passport, id card, driver license, etc.

- int64 balance, binds to local currency per Joule
- (gene, limit, repeat) plan

## Genes

- int64 gid: Gene ID
- int64 pid: Gene owner ID
