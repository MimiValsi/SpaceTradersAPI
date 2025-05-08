-- +goose Up
CREATE TABLE agent (
  accountID TEXT NOT NULL PRIMARY KEY,
  symbol VARCHAR(14) NOT NULL,
  headquarters TEXT NOT NULL,
  credits BIGINT NOT NULL,
  startingFaction TEXT NOT NULL,
  shipCount INTEGER NOT NULL
);
--

-- +goose Down
DROP TABLE agent;
--
