CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'cmoney') THEN
        CREATE TYPE CMONEY AS (
          amount BIGINT,
          currency VARCHAR(3)
        );
    END IF;
END;
$$;


CREATE TABLE operation_type(
  id SERIAL PRIMARY KEY NOT NULL,
  description VARCHAR(50) NOT NULL
);
INSERT INTO operation_type(description) VALUES ('COMPRA A VISTA');
INSERT INTO operation_type(description) VALUES ('COMPRA PARCELADA');
INSERT INTO operation_type(description) VALUES ('SAQUE');
INSERT INTO operation_type(description) VALUES ('PAGAMENTO');

CREATE TABLE transaction(
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  account_id UUID NOT NULL,
  operation_type_id INTEGER REFERENCES operation_type(id),
  amount CMONEY NOT NULL,
  event_date TIMESTAMPTZ NOT NULL
);

