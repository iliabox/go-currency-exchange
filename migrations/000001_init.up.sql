CREATE TABLE current_exchange_rate (
  id UUID PRIMARY KEY,
  currency1 VARCHAR(3) NOT NULL,
  currency2 VARCHAR(3) NOT NULL,
  rate VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL
);
