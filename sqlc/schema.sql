CREATE TYPE cart_status AS ENUM ('pending', 'ordered');

CREATE TABLE products (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(200),
    image VARCHAR(255),
    price INTEGER NOT NULL,
    discount INTEGER,
    created_at TIMESTAMP(0) DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP(0) DEFAULT NOW() NOT NULL
);

CREATE TABLE carts (
    id VARCHAR(36) PRIMARY KEY,
    status cart_status NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP(0) DEFAULT NOW() NOT NULL
);

CREATE TABLE cart_items (
    id VARCHAR(36) PRIMARY KEY,
    cart_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity SMALLINT NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW() NOT NULL,
    
    CONSTRAINT fk_cart FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
