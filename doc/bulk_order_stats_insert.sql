INSERT INTO intraday_market_stats (
  date_of,
  type_id,
  region_id,

  stats_timestamp,
  buy_units,
  sell_units,
  buy_price_max,
  sell_price_min,
  buy_orders,
  sell_orders
) VALUES (
  current_timestamp,
  $1,
  $2,

  ARRAY[ ( $3 )::integer ],
  ARRAY[ ( $4 )::bigint ],
  ARRAY[ ( $5 )::bigint ],
  ARRAY[ ( $6 )::bigint ],
  ARRAY[ ( $7 )::bigint ],
  ARRAY[ ( $8 )::integer ],
  ARRAY[ ( $9 )::integer ]
) ON CONFLICT (type_id, region_id, date_of) DO UPDATE SET
  stats_timestamp = array_append( intraday_market_stats.stats_timestamp, ( $3 )::integer),
  buy_units       = array_append( intraday_market_stats.buy_units,       ( $4 )::bigint),
  sell_units      = array_append( intraday_market_stats.sell_units,      ( $5 )::bigint),
  buy_price_max   = array_append( intraday_market_stats.buy_price_max,   ( $6 )::bigint),
  sell_price_min  = array_append( intraday_market_stats.sell_price_min,  ( $7 )::bigint),
  buy_orders      = array_append( intraday_market_stats.buy_orders,      ( $8 )::integer),
  sell_orders     = array_append( intraday_market_stats.sell_orders,     ( $9 )::integer)
;
