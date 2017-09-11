INSERT INTO bulk_order_stats (
  date_year,
  type_id,
  region_id,

  stats_timestamp,
  buy_units,
  sell_units
) VALUES (
  date_part('year', current_timestamp),
  $1,
  $2,

  ARRAY[ ( $3 )::integer ],
  ARRAY[ ( $4 )::bigint ],
  ARRAY[ ( $5 )::bigint ]
) ON CONFLICT (type_id, region_id, date_year) DO UPDATE SET
  stats_timestamp = array_append( bulk_order_stats.stats_timestamp, ( $3 )::integer),
  buy_units       = array_append( bulk_order_stats.buy_units,       ( $4 )::bigint),
  sell_units      = array_append( bulk_order_stats.sell_units,      ( $5 )::bigint)
;
