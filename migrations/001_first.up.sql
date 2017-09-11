CREATE TABLE bulk_order_stats (
  date_year integer NOT NULL,

  type_id integer NOT NULL,
  region_id integer NOT NULL,

  buy_units bigint[],

  sell_units bigint[],

  stats_timestamp integer[]
);

ALTER TABLE ONLY bulk_order_stats
    ADD CONSTRAINT bulk_order_stats_pkey PRIMARY KEY (type_id, region_id, date_year);

CREATE TABLE bulk_market_history (
  last_date date NOT NULL,
  type_id integer NOT NULL,
  region_id integer NOT NULL,

  orders integer[],
  quantity bigint[],
  low numeric(16,2)[],
  high numeric(16,2)[],
  average numeric(16,2)[]
);

ALTER TABLE ONLY bulk_market_history
    ADD CONSTRAINT bulk_market_history_pkey PRIMARY KEY (type_id, region_id);
