-- name: InsertBDTrades :copyfrom
INSERT INTO bd_trades(
    id,
    bd_dt_date,
    bd_dt_order,
    bd_scrip_name,
    bd_client_name,
    bd_buy_sell,
    bd_qty_trd,
    bd_tp_watp
) VALUES (
     @id,
     @bd_dt_date,
     @bd_dt_order,
     @bd_scrip_name,
     @bd_client_name,
     @bd_buy_sell,
     @bd_qty_trd,
     @bd_tp_watp
);


-- name: GetBulkData :many
SELECT * FROM bd_trades;


