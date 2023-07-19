CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE CMONEY AS (
  amount BIGINT,
  currency VARCHAR(3)
);

CREATE TABLE account(
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  document_number VARCHAR(14) NOT NULL UNIQUE,
  balance CMONEY NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);
CREATE INDEX document_number_idx ON account (document_number);

