-- +goose Up
CREATE TYPE payments_t AS (
  on_accepted INTEGER,
  on_fulfilled INTEGER
);

CREATE TYPE deliver_t AS (
  trade_symbol TEXT,
  destination_symbol TEXT,
  units_required INTEGER,
  units_fulfilled INTEGER
);

CREATE TYPE terms_t AS (
  deadline TIMESTAMP,
  payments payments_t,
  deliver deliver_t[]
);

CREATE TABLE contracts (
  id TEXT NOT NULL PRIMARY KEY,
  factionSymbol TEXT NOT NULL,
  type TEXT NOT NULL,
  terms terms_t NOT NULL,
  accepted BOOLEAN NOT NULL DEFAULT false,
  fulfilled BOOLEAN NOT NULL DEFAULT false,
  deadline_to_accept TIMESTAMP NOT NULL
);
--

-- +goose Down
DROP TABLE contracts;
DROP TYPE terms_t;
DROP TYPE deliver_t;
DROP TYPE payments_t;
--
