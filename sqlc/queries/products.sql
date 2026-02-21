-- name: GetProductById :one
SELECT
  id, 
  name,
  description,
  image,
  price,
  discount,
  created_at,
  updated_at
FROM products
WHERE id = @id::text LIMIT 1;

-- name: GetProductByName :one
SELECT
  id, 
  name,
  description,
  image,
  price,
  discount,
  created_at,
  updated_at
FROM products
WHERE name = @name::text LIMIT 1;

-- name: ListProducts :many
SELECT
  id, 
  name,
  description,
  price,
  discount,
  created_at,
  updated_at
FROM products
ORDER BY created_at;

-- name: CreateProduct :exec
INSERT INTO products (
  id, name, description, price, discount, created_at, updated_at
) VALUES (
  @id::text, @name::text, @description::text, @price::int, @discount::int, @created_at::timestamptz, @updated_at::timestamptz
);

-- name: UpdateProduct :exec
UPDATE products
  set name = @name::text,
  description = @description::text,
  price = @price::int,
  discount = @discount::int,
  updated_at = @updated_at::timestamptz
WHERE id = @id::text;

-- name: UpdateProductImage :exec
UPDATE products
  set image = @image::text,
  updated_at = @updated_at::timestamptz
WHERE id = @id::text;

-- name: DeleteProductById :exec
DELETE FROM products
WHERE id = @id::text;