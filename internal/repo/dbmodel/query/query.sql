/* name: CreateWager :exec */
INSERT INTO wagers (
    id,
    total_wager_value,
    odds,
    selling_percentage,
    selling_price,
    current_selling_price,
    percentage_sold,
    amount_sold,
    placed_at,
) VALUES (?,?,?,?,?,?,?,?,?);

/* name: CreatePurchase :exec */
INSERT INTO purchases (
    id,
    wager_id,
    buying_price
) VALUES (?,?,?);


/* name: GetWager :one */
SELECT * FROM wagers where id = ? LIMIT 1;

/* name: GetPurchase :one */
SELECT * FROM purchases where id = ? LIMIT 1;
