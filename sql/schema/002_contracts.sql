-- +goose Up
CREATE TABLE payment (
  id SERIAL PRIMARY KEY,
  onAccept INTEGER NOT NULL,
  onFulfilled INTEGER NOT NULL
);

CREATE TABLE deliver (
  id SERIAL PRIMARY KEY,
  trade_symbol TEXT NOT NULL,
  destination_symbol TEXT NOT NULL,
  units_required INTEGER NOT NULL,
  units_fulfilled INTEGER NOT NULL
);

CREATE TABLE terms (
  id SERIAL PRIMARY KEY,
  deadline TIMESTAMPTZ NOT NULL,
  payment_id INTEGER REFERENCES payment(id) ON DELETE CASCADE NOT NULL,
  deliver_id INTEGER REFERENCES deliver(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE contract (
  id TEXT PRIMARY KEY,
  factionSymbol TEXT NOT NULL,
  type TEXT NOT NULL,
  terms_id INTEGER REFERENCES terms(id) ON DELETE CASCADE NOT NULL,
  accepted BOOLEAN NOT NULL DEFAULT false,
  fulfilled BOOLEAN NOT NULL DEFAULT false,
  deadline_to_accept TIMESTAMPTZ NOT NULL
);
--

-- +goose Down
DROP TABLE contract;
DROP TABLE terms;
DROP TABLE deliver;
DROP TABLE payment;
--
