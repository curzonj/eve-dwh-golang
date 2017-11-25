CREATE TABLE planet_extraction_history (
  character_id integer not null,
  planet_id integer not null,
  observed_at date not null,
  extraction_type_id integer not null,
  qty_per_cycle integer not null,
  cycle_time integer not null,
  head_radius double precision not null,
  extractor_heads integer not null,
  extractors integer not null,
  basic_factories integer not null,
  upgrade_level integer not null
);

ALTER TABLE ONLY planet_extraction_history
    ADD CONSTRAINT planet_extraction_history_pk PRIMARY KEY (planet_id, extraction_type_id, character_id, observed_at);
