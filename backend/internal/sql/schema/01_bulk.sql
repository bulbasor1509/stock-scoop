-- +goose Up
CREATE TYPE trade AS ENUM ('BUY', 'SELL');

CREATE TABLE bulk(
    id UUID PRIMARY KEY,
    date DATE NOT NULL ,
    symbol TEXT NOT NULL,
    security_name TEXT NOT NULL,
    client_name TEXT NOT NULL,
    trade_type  trade NOT NULL,
    quantity_traded BIGINT NOT NULL,
    wt_price NUMERIC(18,6) NOT NULL
);

-- +goose Down
DROP TABLE bulk;
DROP TYPE trade;