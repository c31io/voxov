--People
CREATE TABLE people (
	pid     bigserial PRIMARY KEY,
	hid     bigint UNIQUE,
	balance bigint NOT NULL,
	phone   varchar(15) UNIQUE NOT NULL,
	pname   varchar(1023),
	id_doc  varchar(1023) UNIQUE,
	dlimit  int NOT NULL,
	created TIMESTAMP NOT NULL,
    last_in TIMESTAMP NOT NULL
);

--Devices
CREATE TABLE devices (
    did     bigserial PRIMARY KEY,
	dtoken  bytea UNIQUE NOT NULL,
	dname   varchar(1023),
	dinfo   varchar(1023),
    pid     bigint NOT NULL,
	created TIMESTAMP NOT NULL,
    last_in TIMESTAMP NOT NULL
);
CREATE INDEX devices_pid_index ON devices (pid);

--Plans
CREATE TABLE plans (
    plan_id bigserial PRIMARY KEY,
    pid     bigint NOT NULL,
    gid     bigint NOT NULL,
    vlimit  bigint NOT NULL,
	billing TIMESTAMP NOT NULL
);
CREATE INDEX plans_pid_index ON plans (pid);
CREATE INDEX plans_gid_index ON plans (gid);

--Transfers
CREATE TABLE transfers (
	tid      bigserial PRIMARY KEY,
	from_pid bigint NOT NULL,
	to_pid   bigint NOT NULL,
	volume   bigint NOT NULL,
	note     varchar(1023),
	ttime    TIMESTAMP NOT NULL
);
CREATE INDEX transfers_from_pid_index ON transfers (from_pid);
CREATE INDEX transfers_to_pid_index   ON transfers (to_pid);

--Genes
CREATE TABLE genes (
	gid     bigserial PRIMARY KEY,
	pid     bigint NOT NULL
);
CREATE INDEX genes_pid_index ON genes (pid);
