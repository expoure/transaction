CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE account(
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  document_number VARCHAR(14) NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);
CREATE INDEX document_number_idx ON account (document_number);

-- -- CREATE TABLE operation_types
-- CREATE TABLE operation_types (id serial primary key NOT NULL, description text);
-- INSERT INTO operation_types(description) VALUES ('COMPRA A VISTA');
-- INSERT INTO operation_types(description) VALUES ('COMPRA PARCELADA');
-- INSERT INTO operation_types(description) VALUES ('SAQUE');
-- INSERT INTO operation_types(description) VALUES ('PAGAMENTO');

-- -- CREATE TABLE trasactions
-- CREATE TABLE transactions (id serial primary key NOT NULL, account_id INTEGER, operation_type_id INTEGER, amount FLOAT, event_date TIMESTAMP);
-- INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (1, 1, -12.5, '2022-10-05T13:33:29.386170889-03:00');
-- INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (2, 1, -50.5, '2022-10-05T13:33:29.386170889-03:00');
-- INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (3, 1, -25.5, '2022-10-05T13:33:29.386170889-03:00');
-- INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (4, 1, 1000.00, '2022-10-05T13:33:29.386170889-03:00');
