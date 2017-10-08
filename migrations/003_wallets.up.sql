CREATE TABLE wallet_journals (
    journal_ref_id bigint NOT NULL,
    occured_at timestamp with time zone NOT NULL,
    ref_type_id integer NOT NULL,
    party_1_id integer NOT NULL,
    party_2_id integer NOT NULL,
    amount bigint NOT NULL,
    reason text,
    tax_collector_id integer,
    tax_amount bigint,
    optional_id bigint,
    optional_value text,
    entity_id bigint NOT NULL,
    entity_character boolean NOT NULL
);

CREATE TABLE wallet_transactions (
    transaction_id bigint NOT NULL,
    character_id bigint NOT NULL,
    occured_at timestamp with time zone NOT NULL,
    quantity bigint NOT NULL,
    type_id integer NOT NULL,
    price bigint NOT NULL,
    client_id integer NOT NULL,
    station_id bigint NOT NULL,
    buy boolean NOT NULL,
    corporate_order boolean NOT NULL,
    journal_ref_id bigint NOT NULL,
    corporation_id bigint
);
