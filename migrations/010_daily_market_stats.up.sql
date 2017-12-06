CREATE TABLE daily_market_stats (
  type_id integer not null,
  region_id integer not null,
  date_of date not null,

  hist_orders integer,
  hist_quantity bigint,
  hist_low bigint,
  hist_high bigint,
  hist_avg bigint,

  buy_units_avg bigint,
  sell_units_avg bigint,
  buy_price_max_avg bigint,
  sell_price_min_avg bigint
);

ALTER TABLE intraday_order_stats RENAME TO intraday_market_stats;
