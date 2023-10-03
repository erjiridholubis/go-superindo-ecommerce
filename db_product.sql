-- Tabel User
CREATE TABLE IF NOT EXISTS users (
    id text PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Tabel Product
CREATE TABLE IF NOT EXISTS products (
    id text PRIMARY KEY,
    category_id text NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price INT NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Tabel Category
CREATE TABLE IF NOT EXISTS categories (
    id text PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Tabel CartItem
CREATE TABLE IF NOT EXISTS cart_items (
    id text PRIMARY KEY,
    user_id text NOT NULL,
    product_id text NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);
