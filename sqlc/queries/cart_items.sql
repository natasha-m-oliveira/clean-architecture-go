-- name: CreateCartItems :exec
INSERT INTO cart_items (id, cart_id, product_id, quantity, created_at)
VALUES (
  unnest(@ids::text[]),
  @cart_id::text,
  unnest(@product_ids::text[]),
  unnest(@quantities::int[]),
  unnest(@created_ats::timestamptz[])
);

-- name: DeleteCartItemsByCartId :exec
DELETE FROM cart_items
WHERE cart_id = @cart_id::text;