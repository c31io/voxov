--People
CREATE TABLE people (
	pid     bigserial PRIMARY KEY,
	hid     bigint UNIQUE,
	balance bigint NOT NULL,
	phone   varchar(15) UNIQUE NOT NULL,
	pname   varchar(1023),
	id_doc  varchar(1023) UNIQUE,
	created TIMESTAMP NOT NULL,
    last_in TIMESTAMP
);
CREATE INDEX hid_index   ON people (hid);
CREATE INDEX phone_index ON people (phone);

--Devices
CREATE TABLE devices (
    did     bigserial PRIMARY KEY,
	dtoken  bytea NOT NULL,
    pid     bigint NOT NULL,
	created TIMESTAMP NOT NULL,
    last_in TIMESTAMP
);
CREATE INDEX dtoken_index ON devices (dtoken);
CREATE INDEX pid_index    ON devices (pid);

--Plans
CREATE TABLE plans (
    plan_id bigserial PRIMARY KEY,
    pid     bigint NOT NULL,
    gid     bigint NOT NULL,
    vlimit  bigint NOT NULL,
	billing TIMESTAMP NOT NULL
);
CREATE INDEX pid_index ON plans (pid);
CREATE INDEX gid_index ON plans (gid);

--Transfers
CREATE TABLE transfers (
	tid      bigserial PRIMARY KEY,
	from_pid bigint NOT NULL,
	to_pid   bigint NOT NULL,
	volume   bigint NOT NULL,
	note     varchar(1023),
	ttime    TIMESTAMP NOT NULL
);
CREATE INDEX from_pid_index ON transfers (from_pid);
CREATE INDEX to_pid_index   ON transfers (to_pid);

--Genes
CREATE TABLE genes (
	gid     bigserial PRIMARY KEY,
	pid     bigint NOT NULL
);
CREATE INDEX pid_index ON genes (pid);
