CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  plaid_item_id TEXT NOT NULL,
  name TEXT NOT NULL
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  account_id INTEGER REFERENCES accounts(id),
  date DATE NOT NULL,
  amount_cents BIGINT NOT NULL,
  category TEXT,
  description TEXT
);
