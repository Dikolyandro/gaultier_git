CREATE TABLE deliveries (
                            id SERIAL PRIMARY KEY,
                            order_id INT NOT NULL,
                            status VARCHAR(20) NOT NULL DEFAULT 'packed',
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
