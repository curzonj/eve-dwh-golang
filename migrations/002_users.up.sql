CREATE TABLE users (
  id uuid PRIMARY KEY
);

CREATE TABLE user_characters (
  user_id uuid NOT NULL,
  id bigint PRIMARY KEY,
  name text NOT NULL,

  owner_hash text NOT NULL,
  oauth_scopes text NOT NULL,
  oauth_token jsonb NOT NULL
);
