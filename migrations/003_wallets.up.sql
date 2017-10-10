CREATE TABLE wallet_journals (
    entity_id bigint NOT NULL,
    division_id bigint,

    journal_ref_id bigint NOT NULL,
    ref_type varchar(255) NOT NULL,
    occured_at timestamp with time zone NOT NULL,
    reason text,
    party_1_id integer NOT NULL,
    party_1_type varchar(255) not null,
    party_2_id integer NOT NULL,
    party_2_type varchar(255) NOT NULL,
    amount bigint NOT NULL,
    balance bigint NOT NULL,
    tax_collector_id integer,
    tax_amount bigint,
    extra_info jsonb
);

ALTER TABLE wallet_journals ADD CONSTRAINT wallet_journals_pkey PRIMARY KEY (entity_id, division_id, journal_ref_id);

CREATE TABLE wallet_transactions (
    entity_id bigint NOT NULL,
    division_id bigint,

    transaction_id bigint NOT NULL,
    character_id bigint NOT NULL,
    occured_at timestamp with time zone NOT NULL,
    quantity bigint NOT NULL,
    type_id integer NOT NULL,
    price bigint NOT NULL,
    client_id integer NOT NULL,
    location_id bigint NOT NULL,
    buy boolean NOT NULL,
    is_personal boolean NOT NULL,
    journal_ref_id bigint NOT NULL
);

ALTER TABLE wallet_transactions ADD CONSTRAINT wallet_transactions_pkey PRIMARY KEY (entity_id, division_id, transaction_id);
