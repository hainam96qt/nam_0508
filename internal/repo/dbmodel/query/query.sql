/* name: CreateWager :exec */
INSERT INTO wagers (
    id,
    total_wager_value,
    odds,
    selling_percentage,
    selling_price,
    current_selling_price,
    percentage_sold,
    amount_sold
) VALUES (?,?,?,?,?,?,?,?);

/* name: CreatePurchase :exec */
INSERT INTO purchases (
    wager_id,
    buying_price
) VALUES (?,?);

/* name: GetWager :one */
SELECT * FROM wagers WHERE id = ?;

/* name: GetWagerForUpdate :one */
SELECT * FROM wagers WHERE id = ? FOR UPDATE;

/* name: ListWagers :many */
SELECT * FROM wagers LIMIT ? OFFSET ?;

/* name: GetPurchase :one */
SELECT * FROM purchases where id = ? LIMIT 1;

/* name: LastInsertID :one */
select last_insert_id();

/* name: UpdatePurchaseWager :exec */
UPDATE wagers SET current_selling_price = ?, percentage_sold = ?, amount_sold = ? WHERE id = ?;