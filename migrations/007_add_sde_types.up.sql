CREATE TABLE sde_types (
  type_id integer PRIMARY KEY,
  group_id integer not null,
  market_group_id integer,
  volume double precision,
  name text not null,

  parent_type_id integer,
  meta_group_id integer,
  meta_level integer
);

CREATE TABLE sde_blueprints (
  blueprint_id integer PRIMARY KEY,
  activities jsonb not null,
  max_production_limit integer not null
);

CREATE TABLE sde_names (
  item_id integer PRIMARY KEY,
  item_name text not null
);

CREATE TABLE sde_planetary_schematics (
  schematic_id integer PRIMARY KEY,
  cycle_time integer,
  factory_type_ids integer[],
  inputs integer[],
  output integer,
  contents jsonb not null
);

CREATE TABLE sde_market_groups (
  market_group_id integer PRIMARY KEY,
  market_group_name text not null,
  has_types boolean not null,
  parent_group_id integer
);

CREATE TABLE sde_meta_groups (
  meta_group_id integer PRIMARY KEY,
  meta_group_name text not null
);

CREATE MATERIALIZED VIEW sde_market_group_arrays AS
  with recursive g(market_group_id, parent_group_id, group_name, id_list, name_list) as (
      select sde_market_groups.market_group_id, sde_market_groups.parent_group_id, sde_market_groups.market_group_name, ARRAY[sde_market_groups.market_group_id]::int[], ARRAY[sde_market_groups.market_group_name] from sde_market_groups where sde_market_groups.parent_group_id is null
        UNION
            select sde_market_groups.market_group_id, sde_market_groups.parent_group_id, sde_market_groups.market_group_name, id_list || sde_market_groups.market_group_id::int, (name_list || sde_market_groups.market_group_name)::varchar(100)[] from g, sde_market_groups where sde_market_groups.parent_group_id = g.market_group_id
  )
  select * from g;

CREATE MATERIALIZED VIEW sde_max_metas AS
  select parent_type_id, max(meta_level) max_meta from sde_types where meta_group_id = 1 and meta_level is not null group by parent_type_id;
