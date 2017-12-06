CREATE TABLE cogs_entries (
  source_id bigint not null,
  source_type text not null, -- job|contract|transaction

  type_id integer not null,
  quantity bigint not null,
  unit_cost bigint not null,
  acquired_on date not null,
  units_remaining bigint not null
);

CREATE INDEX cogs_entries_open ON cogs_entries (type_id) where units_remaining > 0;
ALTER TABLE ONLY cogs_entries
    ADD CONSTRAINT cogs_entries_pkey PRIMARY KEY (source_id, source_type);
