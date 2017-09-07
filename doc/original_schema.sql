--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.4
-- Dumped by pg_dump version 9.5.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: _final_median(anyarray); Type: FUNCTION; Schema: public; Owner: jcurzon
--

CREATE FUNCTION _final_median(anyarray) RETURNS double precision
    LANGUAGE sql IMMUTABLE
    AS $_$
  WITH q AS
  (
     SELECT val
     FROM unnest($1) val
     WHERE VAL IS NOT NULL
     ORDER BY 1
  ),
  cnt AS
  (
    SELECT COUNT(*) AS c FROM q
  )
  SELECT AVG(val)::float8
  FROM
  (
    SELECT val FROM q
    LIMIT  2 - MOD((SELECT c FROM cnt), 2)
    OFFSET GREATEST(CEIL((SELECT c FROM cnt) / 2.0) - 1,0)
  ) q2;
$_$;


ALTER FUNCTION public._final_median(anyarray) OWNER TO jcurzon;

--
-- Name: est_market_share(integer, numeric, numeric, numeric); Type: FUNCTION; Schema: public; Owner: jcurzon
--

CREATE FUNCTION est_market_share(max_isking integer, orders numeric, isking numeric, units numeric) RETURNS numeric
    LANGUAGE sql IMMUTABLE
    AS $$
  select floor(orders * least(max_isking, greatest(1, isking)) / greatest(1, isking)) * units
$$;


ALTER FUNCTION public.est_market_share(max_isking integer, orders numeric, isking numeric, units numeric) OWNER TO jcurzon;

--
-- Name: median(anyelement); Type: AGGREGATE; Schema: public; Owner: jcurzon
--

CREATE AGGREGATE median(anyelement) (
    SFUNC = array_append,
    STYPE = anyarray,
    INITCOND = '{}',
    FINALFUNC = _final_median
);


ALTER AGGREGATE public.median(anyelement) OWNER TO jcurzon;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: market_history; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_history (
    type_id integer NOT NULL,
    region_id integer NOT NULL,
    history_date date NOT NULL,
    orders integer NOT NULL,
    quantity bigint NOT NULL,
    low numeric(16,2) NOT NULL,
    high numeric(16,2) NOT NULL,
    average numeric(16,2) NOT NULL
);


ALTER TABLE market_history OWNER TO jcurzon;

--
-- Name: market_order_stats_ts; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_order_stats_ts (
    calculated_at timestamp with time zone NOT NULL,
    last_updated_at timestamp with time zone NOT NULL,
    type_id integer NOT NULL,
    station_id integer NOT NULL,
    region_id integer NOT NULL,
    buy_price_max numeric(16,2),
    buy_price_wavg numeric(16,2),
    buy_price_5pct numeric(16,2),
    buy_price_median numeric(16,2),
    buy_units bigint,
    buy_orders_price_chg integer,
    buy_orders_vol_chg integer,
    buy_orders_disappeared integer,
    buy_units_vol_chg bigint,
    buy_units_disappeared bigint,
    buy_price_wavg_sold numeric(16,2),
    buy_price_min_sold numeric(16,2),
    buy_price_max_sold numeric(16,2),
    sell_price_min numeric(16,2),
    sell_price_wavg numeric(16,2),
    sell_price_5pct numeric(16,2),
    sell_price_median numeric(16,2),
    sell_units bigint,
    sell_orders_price_chg integer,
    sell_orders_vol_chg integer,
    sell_orders_disappeared integer,
    sell_units_vol_chg bigint,
    sell_units_disappeared bigint,
    sell_price_wavg_sold numeric(16,2),
    sell_price_min_sold numeric(16,2),
    sell_price_max_sold numeric(16,2),
    new_sell_orders integer,
    new_buy_orders integer,
    buy_orders integer,
    sell_orders integer,
    new_buy_order_units bigint,
    new_sell_order_units bigint
);


ALTER TABLE market_order_stats_ts OWNER TO jcurzon;

--
-- Name: observed_history; Type: MATERIALIZED VIEW; Schema: public; Owner: jcurzon
--

CREATE MATERIALIZED VIEW observed_history AS
 SELECT s1.type_id,
    s1.station_id,
    s1.total_observations,
    s1.sell_isking,
    s1.buy_isking,
    s1.buy_units_tx,
    s1.sell_units_tx,
    s1.duration,
    ((s1.buy_isking)::double precision / s1.duration) AS daily_buy_isking,
    ((s1.sell_isking)::double precision / s1.duration) AS daily_sell_isking,
    ((s1.buy_units_tx)::double precision / s1.duration) AS daily_buy_units,
    ((s1.sell_units_tx)::double precision / s1.duration) AS daily_sell_units,
    ((LEAST(s1.buy_units_tx, s1.sell_units_tx))::double precision / s1.duration) AS trade_units_tx
   FROM ( SELECT s2.type_id,
            s2.station_id,
            count(*) AS total_observations,
            sum(s2.sell_orders_price_chg) AS sell_isking,
            sum(s2.buy_orders_price_chg) AS buy_isking,
            sum((s2.buy_units_vol_chg + s2.buy_units_disappeared)) AS buy_units_tx,
            sum((s2.sell_units_vol_chg + s2.sell_units_disappeared)) AS sell_units_tx,
            date_part('day'::text, (max(s2.last_updated_at) - min(s2.last_updated_at))) AS duration
           FROM ( SELECT market_order_stats_ts.calculated_at,
                    market_order_stats_ts.last_updated_at,
                    market_order_stats_ts.type_id,
                    market_order_stats_ts.station_id,
                    market_order_stats_ts.region_id,
                    market_order_stats_ts.buy_price_max,
                    market_order_stats_ts.buy_price_wavg,
                    market_order_stats_ts.buy_price_5pct,
                    market_order_stats_ts.buy_price_median,
                    market_order_stats_ts.buy_units,
                    market_order_stats_ts.buy_orders_price_chg,
                    market_order_stats_ts.buy_orders_vol_chg,
                    market_order_stats_ts.buy_orders_disappeared,
                    market_order_stats_ts.buy_units_vol_chg,
                    market_order_stats_ts.buy_units_disappeared,
                    market_order_stats_ts.buy_price_wavg_sold,
                    market_order_stats_ts.buy_price_min_sold,
                    market_order_stats_ts.buy_price_max_sold,
                    market_order_stats_ts.sell_price_min,
                    market_order_stats_ts.sell_price_wavg,
                    market_order_stats_ts.sell_price_5pct,
                    market_order_stats_ts.sell_price_median,
                    market_order_stats_ts.sell_units,
                    market_order_stats_ts.sell_orders_price_chg,
                    market_order_stats_ts.sell_orders_vol_chg,
                    market_order_stats_ts.sell_orders_disappeared,
                    market_order_stats_ts.sell_units_vol_chg,
                    market_order_stats_ts.sell_units_disappeared,
                    market_order_stats_ts.sell_price_wavg_sold,
                    market_order_stats_ts.sell_price_min_sold,
                    market_order_stats_ts.sell_price_max_sold,
                    market_order_stats_ts.new_sell_orders,
                    market_order_stats_ts.new_buy_orders,
                    market_order_stats_ts.buy_orders,
                    market_order_stats_ts.sell_orders,
                    market_order_stats_ts.new_buy_order_units,
                    market_order_stats_ts.new_sell_order_units
                   FROM market_order_stats_ts
                  WHERE (market_order_stats_ts.last_updated_at > '2000-01-01 00:00:01-08'::timestamp with time zone)) s2
          GROUP BY s2.type_id, s2.station_id) s1
  WHERE (s1.duration > (0)::double precision)
  WITH NO DATA;


ALTER TABLE observed_history OWNER TO jcurzon;

--
-- Name: order_frequencies; Type: MATERIALIZED VIEW; Schema: public; Owner: jcurzon
--

CREATE MATERIALIZED VIEW order_frequencies AS
 SELECT market_history.type_id,
    market_history.region_id,
    trunc((trunc((((max(market_history.history_date) - min(market_history.history_date)) + 1))::numeric, 2) / (count(*))::numeric), 4) AS ratio,
    trunc(avg(market_history.quantity), 4) AS avg_quantity,
    trunc(avg(market_history.orders), 4) AS avg_orders
   FROM market_history
  GROUP BY market_history.region_id, market_history.type_id
  WITH NO DATA;


ALTER TABLE order_frequencies OWNER TO jcurzon;

--
-- Name: recent_observed_history; Type: MATERIALIZED VIEW; Schema: public; Owner: jcurzon
--

CREATE MATERIALIZED VIEW recent_observed_history AS
 SELECT s2.type_id,
    s2.station_id,
    count(*) AS recent_observations,
    avg(s2.buy_price_wavg_sold) AS buy_price_wavg_sold,
    avg(s2.sell_price_wavg_sold) AS sell_price_wavg_sold
   FROM ( SELECT market_order_stats_ts.calculated_at,
            market_order_stats_ts.last_updated_at,
            market_order_stats_ts.type_id,
            market_order_stats_ts.station_id,
            market_order_stats_ts.region_id,
            market_order_stats_ts.buy_price_max,
            market_order_stats_ts.buy_price_wavg,
            market_order_stats_ts.buy_price_5pct,
            market_order_stats_ts.buy_price_median,
            market_order_stats_ts.buy_units,
            market_order_stats_ts.buy_orders_price_chg,
            market_order_stats_ts.buy_orders_vol_chg,
            market_order_stats_ts.buy_orders_disappeared,
            market_order_stats_ts.buy_units_vol_chg,
            market_order_stats_ts.buy_units_disappeared,
            market_order_stats_ts.buy_price_wavg_sold,
            market_order_stats_ts.buy_price_min_sold,
            market_order_stats_ts.buy_price_max_sold,
            market_order_stats_ts.sell_price_min,
            market_order_stats_ts.sell_price_wavg,
            market_order_stats_ts.sell_price_5pct,
            market_order_stats_ts.sell_price_median,
            market_order_stats_ts.sell_units,
            market_order_stats_ts.sell_orders_price_chg,
            market_order_stats_ts.sell_orders_vol_chg,
            market_order_stats_ts.sell_orders_disappeared,
            market_order_stats_ts.sell_units_vol_chg,
            market_order_stats_ts.sell_units_disappeared,
            market_order_stats_ts.sell_price_wavg_sold,
            market_order_stats_ts.sell_price_min_sold,
            market_order_stats_ts.sell_price_max_sold,
            market_order_stats_ts.new_sell_orders,
            market_order_stats_ts.new_buy_orders,
            market_order_stats_ts.buy_orders,
            market_order_stats_ts.sell_orders,
            market_order_stats_ts.new_buy_order_units,
            market_order_stats_ts.new_sell_order_units
           FROM market_order_stats_ts
          WHERE ((now() - market_order_stats_ts.last_updated_at) < '5 days'::interval)) s2
  GROUP BY s2.type_id, s2.station_id
  WITH NO DATA;


ALTER TABLE recent_observed_history OWNER TO jcurzon;

--
-- Name: station_order_stats; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE station_order_stats (
    type_id integer NOT NULL,
    station_id integer NOT NULL,
    region_id integer NOT NULL,
    updated_at timestamp with time zone,
    buy_price_max numeric(16,2),
    buy_price_wavg numeric(16,2),
    buy_price_5pct numeric(16,2),
    buy_price_median numeric(16,2),
    buy_units bigint,
    sell_price_min numeric(16,2),
    sell_price_wavg numeric(16,2),
    sell_price_5pct numeric(16,2),
    sell_price_median numeric(16,2),
    sell_units bigint,
    buy_orders integer,
    sell_orders integer
);


ALTER TABLE station_order_stats OWNER TO jcurzon;

--
-- Name: agg_market_type_stats; Type: VIEW; Schema: public; Owner: jcurzon
--

CREATE VIEW agg_market_type_stats AS
 SELECT station_order_stats.type_id,
    station_order_stats.station_id,
    station_order_stats.region_id,
    station_order_stats.updated_at,
    station_order_stats.buy_price_max,
    station_order_stats.buy_price_wavg,
    station_order_stats.buy_price_5pct,
    station_order_stats.buy_price_median,
    station_order_stats.buy_units,
    station_order_stats.sell_price_min,
    station_order_stats.sell_price_wavg,
    station_order_stats.sell_price_5pct,
    station_order_stats.sell_price_median,
    station_order_stats.sell_units,
    station_order_stats.buy_orders,
    station_order_stats.sell_orders,
    order_frequencies.ratio,
    order_frequencies.avg_quantity,
    order_frequencies.avg_orders,
    observed_history.total_observations,
    observed_history.sell_isking,
    observed_history.buy_isking,
    observed_history.buy_units_tx,
    observed_history.sell_units_tx,
    observed_history.duration,
    observed_history.daily_buy_isking,
    observed_history.daily_sell_isking,
    observed_history.daily_buy_units,
    observed_history.daily_sell_units,
    observed_history.trade_units_tx,
    recent_observed_history.recent_observations,
    recent_observed_history.buy_price_wavg_sold,
    recent_observed_history.sell_price_wavg_sold,
    ((station_order_stats.sell_price_min * 0.985) - (station_order_stats.buy_price_max * 1.0075)) AS max_profit_per_unit,
    ((LEAST(station_order_stats.sell_price_min, recent_observed_history.sell_price_wavg_sold) * 0.985) - (GREATEST(station_order_stats.buy_price_max, recent_observed_history.buy_price_wavg_sold) * 1.0075)) AS wavg_profit_per_unit,
    (observed_history.trade_units_tx * ((10)::double precision / GREATEST((10)::double precision, observed_history.daily_buy_isking, observed_history.daily_sell_isking))) AS est_market_share
   FROM (((station_order_stats
     JOIN order_frequencies USING (type_id, region_id))
     JOIN observed_history USING (type_id, station_id))
     JOIN recent_observed_history USING (type_id, station_id));


ALTER TABLE agg_market_type_stats OWNER TO jcurzon;

--
-- Name: assets; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE assets (
    station_id integer NOT NULL,
    type_id integer NOT NULL,
    quantity bigint NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE assets OWNER TO jcurzon;

--
-- Name: character_order_details; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE character_order_details (
    id bigint NOT NULL,
    character_id bigint NOT NULL,
    order_state integer NOT NULL,
    account_key integer NOT NULL,
    escrow numeric(16,2) NOT NULL,
    type_id integer NOT NULL,
    region_id integer NOT NULL,
    station_id integer NOT NULL,
    issued_at timestamp with time zone NOT NULL
);


ALTER TABLE character_order_details OWNER TO jcurzon;

--
-- Name: character_skills; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE character_skills (
    character_id bigint NOT NULL,
    type_id integer NOT NULL,
    skill_level integer NOT NULL
);


ALTER TABLE character_skills OWNER TO jcurzon;

--
-- Name: characters; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE characters (
    character_id bigint NOT NULL,
    corporation_id integer,
    name character varying(255) NOT NULL
);


ALTER TABLE characters OWNER TO jcurzon;

--
-- Name: corporations; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE corporations (
    corporation_id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE corporations OWNER TO jcurzon;

--
-- Name: eve_api_keys; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE eve_api_keys (
    key_id integer NOT NULL,
    vcode character(64),
    is_corporate boolean DEFAULT false NOT NULL,
    user_account_id integer
);


ALTER TABLE eve_api_keys OWNER TO jcurzon;

--
-- Name: eve_map_stats; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE eve_map_stats (
    region_id integer NOT NULL,
    system_id integer NOT NULL,
    date_of date NOT NULL,
    hour integer NOT NULL,
    ship_kills integer,
    pod_kills integer,
    npc_kills integer,
    jumps integer
);


ALTER TABLE eve_map_stats OWNER TO jcurzon;

--
-- Name: eve_sso; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE eve_sso (
    character_id bigint NOT NULL,
    user_account_id integer NOT NULL
);


ALTER TABLE eve_sso OWNER TO jcurzon;

--
-- Name: historical_orders; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE historical_orders (
    id bigint NOT NULL,
    first_observed_at timestamp with time zone NOT NULL,
    observed_at timestamp with time zone NOT NULL,
    price numeric(16,2) NOT NULL,
    volume_remaining bigint NOT NULL,
    volume_entered bigint NOT NULL,
    min_volume integer NOT NULL,
    buy boolean NOT NULL,
    issue_date timestamp with time zone NOT NULL,
    duration integer NOT NULL,
    range integer NOT NULL,
    type_id integer NOT NULL,
    station_id integer NOT NULL,
    region_id integer NOT NULL,
    disappeared_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE historical_orders OWNER TO jcurzon;

--
-- Name: industry_jobs; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE industry_jobs (
    job_id bigint NOT NULL,
    installer_id bigint NOT NULL,
    activity_id integer NOT NULL,
    blueprint_type_id bigint NOT NULL,
    start_date timestamp with time zone NOT NULL,
    completed_date timestamp with time zone,
    job_data jsonb NOT NULL
);


ALTER TABLE industry_jobs OWNER TO jcurzon;

--
-- Name: managed_characters; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE managed_characters (
    character_id bigint NOT NULL,
    key_id integer NOT NULL,
    total_orders integer DEFAULT 5 NOT NULL,
    broker_fee_rate numeric(8,8) DEFAULT 0.01 NOT NULL,
    sales_tax_rate numeric(8,8) DEFAULT 0.025 NOT NULL,
    poco_tax_rate numeric(2,2) DEFAULT 0.15 NOT NULL
);


ALTER TABLE managed_characters OWNER TO jcurzon;

--
-- Name: managed_corps; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE managed_corps (
    corporation_id integer NOT NULL,
    key_id integer NOT NULL
);


ALTER TABLE managed_corps OWNER TO jcurzon;

--
-- Name: market_change_conflicts; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_change_conflicts (
    observed_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    crest_order jsonb NOT NULL,
    conflicting_row jsonb NOT NULL,
    previous_row jsonb
);


ALTER TABLE market_change_conflicts OWNER TO jcurzon;

--
-- Name: market_daily_stats; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_daily_stats (
    date_of date NOT NULL,
    type_id integer NOT NULL,
    station_id integer NOT NULL,
    region_id integer NOT NULL,
    stats_timestamp integer[],
    buy_price_max numeric(16,2)[],
    buy_price_wavg numeric(16,2)[],
    buy_price_5pct numeric(16,2)[],
    buy_price_median numeric(16,2)[],
    buy_units bigint[],
    buy_orders_price_chg integer[],
    buy_orders_vol_chg integer[],
    buy_orders_disappeared integer[],
    buy_units_vol_chg bigint[],
    buy_units_disappeared bigint[],
    buy_price_wavg_sold numeric(16,2)[],
    buy_price_min_sold numeric(16,2)[],
    buy_price_max_sold numeric(16,2)[],
    sell_price_min numeric(16,2)[],
    sell_price_wavg numeric(16,2)[],
    sell_price_5pct numeric(16,2)[],
    sell_price_median numeric(16,2)[],
    sell_units bigint[],
    sell_orders_price_chg integer[],
    sell_orders_vol_chg integer[],
    sell_orders_disappeared integer[],
    sell_units_vol_chg bigint[],
    sell_units_disappeared bigint[],
    sell_price_wavg_sold numeric(16,2)[],
    sell_price_min_sold numeric(16,2)[],
    sell_price_max_sold numeric(16,2)[],
    new_sell_orders integer[],
    new_buy_orders integer[],
    new_sell_order_units bigint[],
    new_buy_order_units bigint[],
    hist_orders integer,
    hist_quantity bigint,
    hist_low numeric(16,2),
    hist_high numeric(16,2),
    hist_average numeric(16,2),
    day_buy_order_price_changes integer DEFAULT 0 NOT NULL,
    day_sell_order_price_changes integer DEFAULT 0 NOT NULL,
    day_buy_price_min_tx numeric(16,2),
    day_sell_price_min_tx numeric(16,2),
    day_buy_price_max_tx numeric(16,2),
    day_sell_price_max_tx numeric(16,2),
    day_buy_price_wavg_tx numeric(16,2),
    day_sell_price_wavg_tx numeric(16,2),
    day_new_buy_orders integer DEFAULT 0 NOT NULL,
    day_new_sell_orders integer DEFAULT 0 NOT NULL,
    day_buy_orders_tx integer DEFAULT 0 NOT NULL,
    day_sell_orders_tx integer DEFAULT 0 NOT NULL,
    day_buy_units_tx bigint DEFAULT 0 NOT NULL,
    day_sell_units_tx bigint DEFAULT 0 NOT NULL,
    day_avg_buy_units bigint DEFAULT 0 NOT NULL,
    day_avg_sell_units bigint DEFAULT 0 NOT NULL
);


ALTER TABLE market_daily_stats OWNER TO jcurzon;

--
-- Name: market_daily_stats_y2016m03; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_daily_stats_y2016m03 (
    CONSTRAINT market_daily_stats_y2016m03_date_of_check CHECK ((date_of < '2016-04-01'::date))
)
INHERITS (market_daily_stats);


ALTER TABLE market_daily_stats_y2016m03 OWNER TO jcurzon;

--
-- Name: market_daily_stats_y2016m04; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_daily_stats_y2016m04 (
    CONSTRAINT market_daily_stats_y2016m04_date_of_check CHECK (((date_of >= '2016-04-01'::date) AND (date_of < '2016-05-01'::date)))
)
INHERITS (market_daily_stats);


ALTER TABLE market_daily_stats_y2016m04 OWNER TO jcurzon;

--
-- Name: market_order_changes; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_order_changes (
    order_id bigint NOT NULL,
    type_id integer NOT NULL,
    station_id integer NOT NULL,
    region_id integer NOT NULL,
    observed_at timestamp with time zone NOT NULL,
    previously_observed_at timestamp with time zone,
    issue_date timestamp with time zone NOT NULL,
    previous_issue_date timestamp with time zone,
    volume_remaining bigint NOT NULL,
    volume_delta bigint,
    price numeric(16,2) NOT NULL,
    previous_price numeric(16,2),
    disappeared boolean DEFAULT false NOT NULL,
    buy boolean NOT NULL,
    canceled boolean,
    order_first_observed_at timestamp with time zone NOT NULL
);


ALTER TABLE market_order_changes OWNER TO jcurzon;

--
-- Name: market_order_changes_y2016m03; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_order_changes_y2016m03 (
    order_id bigint,
    type_id integer,
    station_id integer,
    region_id integer,
    observed_at timestamp with time zone,
    previously_observed_at timestamp with time zone,
    issue_date timestamp with time zone,
    previous_issue_date timestamp with time zone,
    volume_remaining bigint,
    volume_delta bigint,
    price numeric(16,2),
    previous_price numeric(16,2),
    disappeared boolean DEFAULT false,
    buy boolean,
    canceled boolean,
    order_first_observed_at timestamp with time zone,
    CONSTRAINT y2016m03 CHECK ((observed_at < '2016-04-01'::date))
)
INHERITS (market_order_changes);


ALTER TABLE market_order_changes_y2016m03 OWNER TO jcurzon;

--
-- Name: market_order_snaps; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_order_snaps (
    observed_at timestamp with time zone NOT NULL,
    region_id integer NOT NULL,
    type_id integer NOT NULL,
    buy_order_data jsonb NOT NULL,
    sell_order_data jsonb NOT NULL
);


ALTER TABLE market_order_snaps OWNER TO jcurzon;

--
-- Name: market_order_snaps_y2016m03; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_order_snaps_y2016m03 (
    CONSTRAINT market_order_snaps_y2016m03_observed_at_check CHECK ((observed_at < '2016-04-01'::date))
)
INHERITS (market_order_snaps);


ALTER TABLE market_order_snaps_y2016m03 OWNER TO jcurzon;

--
-- Name: market_order_stats_ts_y2016m03; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_order_stats_ts_y2016m03 (
    calculated_at timestamp with time zone,
    last_updated_at timestamp with time zone,
    type_id integer,
    station_id integer,
    region_id integer,
    buy_price_max numeric(16,2),
    buy_price_wavg numeric(16,2),
    buy_price_5pct numeric(16,2),
    buy_price_median numeric(16,2),
    buy_units bigint,
    buy_orders_price_chg integer,
    buy_orders_vol_chg integer,
    buy_orders_disappeared integer,
    buy_units_vol_chg bigint,
    buy_units_disappeared bigint,
    buy_price_wavg_sold numeric(16,2),
    buy_price_min_sold numeric(16,2),
    buy_price_max_sold numeric(16,2),
    sell_price_min numeric(16,2),
    sell_price_wavg numeric(16,2),
    sell_price_5pct numeric(16,2),
    sell_price_median numeric(16,2),
    sell_units bigint,
    sell_orders_price_chg integer,
    sell_orders_vol_chg integer,
    sell_orders_disappeared integer,
    sell_units_vol_chg bigint,
    sell_units_disappeared bigint,
    sell_price_wavg_sold numeric(16,2),
    sell_price_min_sold numeric(16,2),
    sell_price_max_sold numeric(16,2),
    new_sell_orders integer,
    new_buy_orders integer,
    buy_orders integer,
    sell_orders integer,
    new_buy_order_units bigint,
    new_sell_order_units bigint,
    CONSTRAINT y2016m03 CHECK ((calculated_at < '2016-04-01'::date))
)
INHERITS (market_order_stats_ts);


ALTER TABLE market_order_stats_ts_y2016m03 OWNER TO jcurzon;

--
-- Name: market_orders; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_orders (
    id bigint NOT NULL,
    first_observed_at timestamp with time zone NOT NULL,
    observed_at timestamp with time zone NOT NULL,
    price numeric(16,2) NOT NULL,
    volume_remaining bigint NOT NULL,
    volume_entered bigint NOT NULL,
    min_volume integer NOT NULL,
    buy boolean NOT NULL,
    issue_date timestamp with time zone NOT NULL,
    duration integer NOT NULL,
    range integer NOT NULL,
    type_id integer NOT NULL,
    station_id integer NOT NULL,
    region_id integer NOT NULL
);


ALTER TABLE market_orders OWNER TO jcurzon;

--
-- Name: market_polling; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE market_polling (
    type_id integer NOT NULL,
    region_id integer NOT NULL,
    orders_next_polling_at timestamp with time zone DEFAULT '2000-01-01 00:00:01-08'::timestamp with time zone NOT NULL,
    orders_polling_interval interval day to minute DEFAULT '06:00:00'::interval NOT NULL,
    history_next_polling_at timestamp with time zone DEFAULT '2000-01-01 00:00:01-08'::timestamp with time zone NOT NULL,
    history_polling_interval interval day to minute DEFAULT '3 days'::interval NOT NULL,
    orders_polling_override interval day to minute,
    order_polling_started_at timestamp with time zone DEFAULT '2000-01-01 00:00:01-08'::timestamp with time zone NOT NULL,
    history_polling_started_at timestamp with time zone DEFAULT '2000-01-01 00:00:01-08'::timestamp with time zone NOT NULL
);


ALTER TABLE market_polling OWNER TO jcurzon;

--
-- Name: metric_observations; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE metric_observations (
    hostname character varying(255) NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    proc_type character varying(255) NOT NULL,
    metrics jsonb NOT NULL
);


ALTER TABLE metric_observations OWNER TO jcurzon;

--
-- Name: neow_cache; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE neow_cache (
    sha1_hex character(41) NOT NULL,
    cache_until timestamp with time zone NOT NULL,
    json_data jsonb NOT NULL
);


ALTER TABLE neow_cache OWNER TO jcurzon;

--
-- Name: planetary_observations; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE planetary_observations (
    planet_id integer NOT NULL,
    character_id bigint NOT NULL,
    last_updated_at timestamp with time zone NOT NULL,
    observed_at timestamp with time zone,
    observation_data jsonb NOT NULL
);


ALTER TABLE planetary_observations OWNER TO jcurzon;

--
-- Name: wallet_transactions; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE wallet_transactions (
    transaction_id bigint NOT NULL,
    character_id bigint NOT NULL,
    occured_at timestamp with time zone NOT NULL,
    quantity bigint NOT NULL,
    type_id integer NOT NULL,
    price numeric(16,2) NOT NULL,
    client_id integer NOT NULL,
    station_id bigint NOT NULL,
    buy boolean NOT NULL,
    corporate_order boolean NOT NULL,
    journal_ref_id bigint NOT NULL,
    corporation_id bigint
);


ALTER TABLE wallet_transactions OWNER TO jcurzon;

--
-- Name: purchase_costs; Type: VIEW; Schema: public; Owner: jcurzon
--

CREATE VIEW purchase_costs AS
 SELECT a2.station_id,
    a2.type_id,
    a2.quantity,
    COALESCE(( SELECT s1.avg_price
           FROM ( SELECT t1.type_id,
                    t1.station_id,
                    t1.occured_at,
                    ( SELECT sum(t2.quantity) AS sum
                           FROM wallet_transactions t2
                          WHERE ((t1.type_id = t2.type_id) AND (t2.buy = true) AND (t2.occured_at >= t1.occured_at))) AS quantity_sum,
                    ( SELECT (sum((((t2.quantity)::numeric * t2.price) * 1.0075)) / sum(t2.quantity))
                           FROM wallet_transactions t2
                          WHERE ((t1.type_id = t2.type_id) AND (t2.buy = true) AND (t2.occured_at >= t1.occured_at))) AS avg_price
                   FROM wallet_transactions t1
                  WHERE (t1.buy = true)
                  ORDER BY t1.occured_at DESC) s1
          WHERE ((a2.station_id = s1.station_id) AND (a2.type_id = s1.type_id) AND (s1.quantity_sum >= a2.quantity))
         LIMIT 1), ( SELECT (max(m1.price) * 1.0075)
           FROM market_orders m1
          WHERE ((m1.type_id = a2.type_id) AND (m1.station_id = a2.station_id) AND (m1.buy = true)))) AS cost
   FROM ( SELECT a1.station_id,
            a1.type_id,
            ((((a1.quantity)::numeric + COALESCE(( SELECT sum(mo.volume_remaining) AS sum
                   FROM (character_order_details co
                     JOIN market_orders mo USING (id))
                  WHERE ((mo.buy = false) AND (co.type_id = a1.type_id) AND (co.station_id = a1.station_id))), (0)::numeric)) + COALESCE(( SELECT sum(wt.quantity) AS sum
                   FROM wallet_transactions wt
                  WHERE ((wt.occured_at > a1.updated_at) AND (wt.buy = true) AND (wt.type_id = a1.type_id) AND (wt.station_id = a1.station_id))), (0)::numeric)) - COALESCE(( SELECT sum(wt.quantity) AS sum
                   FROM wallet_transactions wt
                  WHERE ((wt.occured_at > a1.updated_at) AND (wt.buy = false) AND (wt.type_id = a1.type_id) AND (wt.station_id = a1.station_id))), (0)::numeric)) AS quantity
           FROM assets a1
        UNION
         SELECT co.station_id,
            co.type_id,
            mo.volume_remaining AS quantity
           FROM (character_order_details co
             JOIN market_orders mo USING (id))
          WHERE ((mo.buy = false) AND (NOT (EXISTS ( SELECT 1
                   FROM assets a3
                  WHERE ((a3.type_id = co.type_id) AND (a3.station_id = co.station_id) AND (a3.quantity > 0))))))) a2;


ALTER TABLE purchase_costs OWNER TO jcurzon;

--
-- Name: schemaversion; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE schemaversion (
    version bigint NOT NULL,
    name text DEFAULT ''::text,
    md5 text DEFAULT ''::text
);


ALTER TABLE schemaversion OWNER TO jcurzon;

--
-- Name: standings; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE standings (
    character_id bigint NOT NULL,
    corporation_id integer NOT NULL,
    standing numeric(4,2) DEFAULT 0.0 NOT NULL,
    broker_fee_rate numeric(8,8) DEFAULT 0.01 NOT NULL
);


ALTER TABLE standings OWNER TO jcurzon;

--
-- Name: trade_regions; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE trade_regions (
    character_id bigint NOT NULL,
    region_id integer NOT NULL
);


ALTER TABLE trade_regions OWNER TO jcurzon;

--
-- Name: user_accounts; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE user_accounts (
    id integer NOT NULL,
    character_id_list bigint[] DEFAULT ARRAY[]::bigint[] NOT NULL,
    corporation_id_list bigint[] DEFAULT ARRAY[]::bigint[] NOT NULL
);


ALTER TABLE user_accounts OWNER TO jcurzon;

--
-- Name: user_accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: jcurzon
--

CREATE SEQUENCE user_accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_accounts_id_seq OWNER TO jcurzon;

--
-- Name: user_accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jcurzon
--

ALTER SEQUENCE user_accounts_id_seq OWNED BY user_accounts.id;


--
-- Name: wallet_journal; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE wallet_journal (
    journal_ref_id bigint NOT NULL,
    occured_at timestamp with time zone NOT NULL,
    ref_type_id integer NOT NULL,
    party_1_id integer NOT NULL,
    party_2_id integer NOT NULL,
    amount numeric(18,2) NOT NULL,
    reason character varying(255),
    tax_collector_id integer,
    tax_amount numeric(16,2),
    optional_id bigint,
    optional_value character varying(255),
    entity_id bigint NOT NULL,
    entity_character boolean NOT NULL
);


ALTER TABLE wallet_journal OWNER TO jcurzon;

--
-- Name: zkillboard_data; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE zkillboard_data (
    kill_id bigint NOT NULL,
    system_id integer NOT NULL,
    kill_time timestamp with time zone NOT NULL,
    kill_data jsonb NOT NULL
);


ALTER TABLE zkillboard_data OWNER TO jcurzon;

--
-- Name: zkillboard_data_y2016m03; Type: TABLE; Schema: public; Owner: jcurzon
--

CREATE TABLE zkillboard_data_y2016m03 (
    kill_id bigint,
    system_id integer,
    kill_time timestamp with time zone,
    kill_data jsonb,
    CONSTRAINT y2016m03 CHECK ((kill_time < '2016-04-01'::date))
)
INHERITS (zkillboard_data);


ALTER TABLE zkillboard_data_y2016m03 OWNER TO jcurzon;

--
-- Name: day_buy_order_price_changes; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_buy_order_price_changes SET DEFAULT 0;


--
-- Name: day_sell_order_price_changes; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_sell_order_price_changes SET DEFAULT 0;


--
-- Name: day_new_buy_orders; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_new_buy_orders SET DEFAULT 0;


--
-- Name: day_new_sell_orders; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_new_sell_orders SET DEFAULT 0;


--
-- Name: day_buy_orders_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_buy_orders_tx SET DEFAULT 0;


--
-- Name: day_sell_orders_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_sell_orders_tx SET DEFAULT 0;


--
-- Name: day_buy_units_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_buy_units_tx SET DEFAULT 0;


--
-- Name: day_sell_units_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_sell_units_tx SET DEFAULT 0;


--
-- Name: day_avg_buy_units; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_avg_buy_units SET DEFAULT 0;


--
-- Name: day_avg_sell_units; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03 ALTER COLUMN day_avg_sell_units SET DEFAULT 0;


--
-- Name: day_buy_order_price_changes; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_buy_order_price_changes SET DEFAULT 0;


--
-- Name: day_sell_order_price_changes; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_sell_order_price_changes SET DEFAULT 0;


--
-- Name: day_new_buy_orders; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_new_buy_orders SET DEFAULT 0;


--
-- Name: day_new_sell_orders; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_new_sell_orders SET DEFAULT 0;


--
-- Name: day_buy_orders_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_buy_orders_tx SET DEFAULT 0;


--
-- Name: day_sell_orders_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_sell_orders_tx SET DEFAULT 0;


--
-- Name: day_buy_units_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_buy_units_tx SET DEFAULT 0;


--
-- Name: day_sell_units_tx; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_sell_units_tx SET DEFAULT 0;


--
-- Name: day_avg_buy_units; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_avg_buy_units SET DEFAULT 0;


--
-- Name: day_avg_sell_units; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04 ALTER COLUMN day_avg_sell_units SET DEFAULT 0;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY user_accounts ALTER COLUMN id SET DEFAULT nextval('user_accounts_id_seq'::regclass);


--
-- Name: assets_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY assets
    ADD CONSTRAINT assets_pkey PRIMARY KEY (type_id, station_id);


--
-- Name: character_order_details_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY character_order_details
    ADD CONSTRAINT character_order_details_pkey PRIMARY KEY (id);


--
-- Name: character_skills_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY character_skills
    ADD CONSTRAINT character_skills_pkey PRIMARY KEY (character_id, type_id);


--
-- Name: characters_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY characters
    ADD CONSTRAINT characters_pkey PRIMARY KEY (character_id);


--
-- Name: corporations_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY corporations
    ADD CONSTRAINT corporations_pkey PRIMARY KEY (corporation_id);


--
-- Name: eve_api_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY eve_api_keys
    ADD CONSTRAINT eve_api_keys_pkey PRIMARY KEY (key_id);


--
-- Name: eve_map_stats_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY eve_map_stats
    ADD CONSTRAINT eve_map_stats_pkey PRIMARY KEY (system_id, date_of, hour);


--
-- Name: historical_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY historical_orders
    ADD CONSTRAINT historical_orders_pkey PRIMARY KEY (id, first_observed_at);


--
-- Name: industry_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY industry_jobs
    ADD CONSTRAINT industry_jobs_pkey PRIMARY KEY (job_id);


--
-- Name: managed_characters_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY managed_characters
    ADD CONSTRAINT managed_characters_pkey PRIMARY KEY (character_id);


--
-- Name: managed_corps_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY managed_corps
    ADD CONSTRAINT managed_corps_pkey PRIMARY KEY (corporation_id);


--
-- Name: market_daily_stats_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats
    ADD CONSTRAINT market_daily_stats_pkey PRIMARY KEY (type_id, region_id, station_id, date_of);


--
-- Name: market_daily_stats_y2016m03_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m03
    ADD CONSTRAINT market_daily_stats_y2016m03_pkey PRIMARY KEY (type_id, region_id, station_id, date_of);


--
-- Name: market_daily_stats_y2016m04_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_daily_stats_y2016m04
    ADD CONSTRAINT market_daily_stats_y2016m04_pkey PRIMARY KEY (type_id, region_id, station_id, date_of);


--
-- Name: market_history_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_history
    ADD CONSTRAINT market_history_pkey PRIMARY KEY (type_id, region_id, history_date);


--
-- Name: market_order_changes_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_order_changes_y2016m03
    ADD CONSTRAINT market_order_changes_pkey PRIMARY KEY (order_id, issue_date, volume_remaining);


--
-- Name: market_order_snaps_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_order_snaps
    ADD CONSTRAINT market_order_snaps_pkey PRIMARY KEY (observed_at, region_id, type_id);


--
-- Name: market_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_orders
    ADD CONSTRAINT market_orders_pkey PRIMARY KEY (id);


--
-- Name: market_polling_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_polling
    ADD CONSTRAINT market_polling_pkey PRIMARY KEY (type_id, region_id);


--
-- Name: neow_cache_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY neow_cache
    ADD CONSTRAINT neow_cache_pkey PRIMARY KEY (sha1_hex);


--
-- Name: planetary_observations_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY planetary_observations
    ADD CONSTRAINT planetary_observations_pkey PRIMARY KEY (character_id, planet_id, last_updated_at);


--
-- Name: schemaversion_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY schemaversion
    ADD CONSTRAINT schemaversion_pkey PRIMARY KEY (version);


--
-- Name: standings_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY standings
    ADD CONSTRAINT standings_pkey PRIMARY KEY (character_id, corporation_id);


--
-- Name: station_order_stats_pkey2; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY station_order_stats
    ADD CONSTRAINT station_order_stats_pkey2 PRIMARY KEY (type_id, region_id, station_id);


--
-- Name: station_order_stats_ts_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY market_order_stats_ts_y2016m03
    ADD CONSTRAINT station_order_stats_ts_pkey PRIMARY KEY (type_id, region_id, station_id, calculated_at);


--
-- Name: trade_regions_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY trade_regions
    ADD CONSTRAINT trade_regions_pkey PRIMARY KEY (character_id, region_id);


--
-- Name: user_accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY user_accounts
    ADD CONSTRAINT user_accounts_pkey PRIMARY KEY (id);


--
-- Name: wallet_journal_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY wallet_journal
    ADD CONSTRAINT wallet_journal_pkey PRIMARY KEY (entity_character, entity_id, journal_ref_id);


--
-- Name: wallet_transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY wallet_transactions
    ADD CONSTRAINT wallet_transactions_pkey PRIMARY KEY (character_id, transaction_id);


--
-- Name: zkillboard_data_pkey; Type: CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY zkillboard_data_y2016m03
    ADD CONSTRAINT zkillboard_data_pkey PRIMARY KEY (kill_id);


--
-- Name: history_polling_interval_index; Type: INDEX; Schema: public; Owner: jcurzon
--

CREATE INDEX history_polling_interval_index ON market_polling USING btree (history_polling_interval);


--
-- Name: market_history_reverse_index; Type: INDEX; Schema: public; Owner: jcurzon
--

CREATE INDEX market_history_reverse_index ON market_history USING btree (region_id, type_id);


--
-- Name: market_order_changes_i2; Type: INDEX; Schema: public; Owner: jcurzon
--

CREATE INDEX market_order_changes_i2 ON market_order_changes_y2016m03 USING btree (type_id, region_id, observed_at, disappeared);


--
-- Name: market_order_changes_i3; Type: INDEX; Schema: public; Owner: jcurzon
--

CREATE INDEX market_order_changes_i3 ON market_order_changes_y2016m03 USING btree (observed_at, disappeared);


--
-- Name: pkey; Type: INDEX; Schema: public; Owner: jcurzon
--

CREATE INDEX pkey ON order_frequencies USING btree (region_id, type_id);


--
-- Name: character_order_details_character_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY character_order_details
    ADD CONSTRAINT character_order_details_character_id_fkey FOREIGN KEY (character_id) REFERENCES managed_characters(character_id);


--
-- Name: character_skills_character_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY character_skills
    ADD CONSTRAINT character_skills_character_id_fkey FOREIGN KEY (character_id) REFERENCES managed_characters(character_id);


--
-- Name: characters_corporation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY characters
    ADD CONSTRAINT characters_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES corporations(corporation_id);


--
-- Name: eve_sso_user_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY eve_sso
    ADD CONSTRAINT eve_sso_user_account_id_fkey FOREIGN KEY (user_account_id) REFERENCES user_accounts(id);


--
-- Name: managed_characters_character_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY managed_characters
    ADD CONSTRAINT managed_characters_character_id_fkey FOREIGN KEY (character_id) REFERENCES characters(character_id);


--
-- Name: managed_characters_key_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY managed_characters
    ADD CONSTRAINT managed_characters_key_id_fkey FOREIGN KEY (key_id) REFERENCES eve_api_keys(key_id);


--
-- Name: managed_corps_corporation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY managed_corps
    ADD CONSTRAINT managed_corps_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES corporations(corporation_id);


--
-- Name: managed_corps_key_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY managed_corps
    ADD CONSTRAINT managed_corps_key_id_fkey FOREIGN KEY (key_id) REFERENCES eve_api_keys(key_id);


--
-- Name: standings_character_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY standings
    ADD CONSTRAINT standings_character_id_fkey FOREIGN KEY (character_id) REFERENCES managed_characters(character_id);


--
-- Name: standings_corporation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY standings
    ADD CONSTRAINT standings_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES corporations(corporation_id);


--
-- Name: trade_regions_character_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jcurzon
--

ALTER TABLE ONLY trade_regions
    ADD CONSTRAINT trade_regions_character_id_fkey FOREIGN KEY (character_id) REFERENCES managed_characters(character_id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

