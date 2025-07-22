-- +goose Up
CREATE TYPE trade AS ENUM ('BUY', 'SELL');

CREATE TABLE bd_trades(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bd_dt_date DATE NOT NULL,
    bd_dt_order TIMESTAMPTZ NOT NULL,
    bd_scrip_name VARCHAR(100) NOT NULL,
    bd_client_name VARCHAR(100) NOT NULL,
    bd_buy_sell trade NOT NULL,
    bd_qty_trd BIGINT NOT NULL,
    bd_tp_watp NUMERIC(10, 2) NOT NULL
);

-- +goose Down
DROP TABLE bd_trades;
