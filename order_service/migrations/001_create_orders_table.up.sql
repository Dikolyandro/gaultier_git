CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INT NOT NULL,
                        items JSONB NOT NULL,
                        status VARCHAR(20) DEFAULT 'created',
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
