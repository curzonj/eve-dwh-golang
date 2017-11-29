ALTER TABLE intraday_order_stats RENAME TO bulk_order_stats;
TRUNCATE TABLE bulk_order_stats;

ALTER TABLE bulk_order_stats DROP constraint if exists bulk_order_stats_pkey;

ALTER TABLE bulk_order_stats DROP column if exists date_of;
ALTER TABLE bulk_order_stats ADD column date_year integer NOT NULL;

ALTER TABLE ONLY bulk_order_stats
    ADD CONSTRAINT bulk_order_stats_pkey PRIMARY KEY (type_id, region_id, date_year);

ALTER TABLE bulk_order_stats drop column if exists buy_price_max;
ALTER TABLE bulk_order_stats drop column if exists sell_price_min;
ALTER TABLE bulk_order_stats drop column if exists buy_orders;
ALTER TABLE bulk_order_stats drop column if exists sell_orders;
