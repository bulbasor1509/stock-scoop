-- name: InsertBulkData :exec
INSERT INTO bulk_deals (
    date,
    symbol,
    security_name,
    client_name,
    trade_type,
    quantity_traded,
    wt_price
)
SELECT
    UNNEST(@dates::DATE[]),
    UNNEST(@symbols::TEXT[]),
    UNNEST(@security_names::TEXT[]),
    UNNEST(@client_names::text[]),
    UNNEST(@trade_types::trade[]),
    UNNEST(@quantities::TEXT[]),
    UNNEST(@prices::TEXT[]);


-- name: GetBulkData :many
SELECT * FROM bulk_deals;


