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
    UNNEST(@dates::date[]),
    UNNEST(@symbols::text[]),
    UNNEST(@security_names::text[]),
    UNNEST(@client_names::text[]),
    UNNEST(@trade_types::trade[]),
    UNNEST(@quantities::bigint[]),
    UNNEST(@prices::numeric(18,6)[]);


-- name: GetBulkData :many
SELECT * FROM bulk_deals;


