create table cache_items (
  key varchar(255) primary key,
  content bytea not null,
  last_touched_at timestamp with time zone NOT NULL
);
