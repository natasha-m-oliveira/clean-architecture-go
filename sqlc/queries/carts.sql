-- name: GetCartByIdWithItems :one
SELECT
  carts.id,
  carts.status,
  carts.created_at,
  carts.updated_at,
 COALESCE(json_agg(json_build_object(
    'id', cart_items.id,
    'cart_id', cart_items.cart_id,
    'product_id', cart_items.product_id,
    'quantity', cart_items.quantity,
    'created_at', cart_items.created_at
  ) ORDER BY cart_items.created_at) FILTER (WHERE cart_items.id IS NOT NULL), '[]')::json AS items
FROM carts
LEFT JOIN cart_items ON carts.id = cart_items.cart_id
WHERE cart_id = @cart_id::text
GROUP BY carts.id
LIMIT 1;

-- name: ListCarts :many
SELECT
  id,
  status,
  created_at,
  updated_at
FROM carts
ORDER BY created_at;

-- name: CreateCart :exec
INSERT INTO carts (
  id, created_at, updated_at
) VALUES (@cart_id, @created_at, @updated_at);

-- name: UpdateCartStatus :exec
UPDATE carts
  set status = @status::text,
  updated_at = @updated_at::timestamptz
WHERE id = @id::text;

-- name: DeleteCartById :exec
DELETE FROM carts
WHERE id = @id::text;