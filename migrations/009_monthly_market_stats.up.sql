TRUNCATE TABLE bulk_order_stats;

ALTER TABLE bulk_order_stats DROP constraint if exists bulk_order_stats_pkey;

ALTER TABLE bulk_order_stats add column date_of date NOT NULL;
ALTER TABLE bulk_order_stats DROP column if exists date_year;

ALTER TABLE ONLY bulk_order_stats
    ADD CONSTRAINT bulk_order_stats_pkey PRIMARY KEY (type_id, region_id, date_of);

ALTER TABLE bulk_order_stats add column buy_price_max bigint[] NOT NULL;
ALTER TABLE bulk_order_stats add column sell_price_min bigint[] NOT NULL;
ALTER TABLE bulk_order_stats add column buy_orders integer[] NOT NULL;
ALTER TABLE bulk_order_stats add column sell_orders integer[] NOT NULL;

ALTER TABLE bulk_order_stats RENAME TO intraday_order_stats;
