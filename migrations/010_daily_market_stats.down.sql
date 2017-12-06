DROP TABLE IF EXISTS daily_market_stats;
ALTER TABLE intraday_market_stats RENAME TO intraday_order_stats;
