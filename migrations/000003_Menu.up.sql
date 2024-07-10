CREATE TABLE IF NOT EXISTS Menu( 
    id UUID NOT NULL,
    restaurant_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price DECIMAL(8, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (restaurant_id) REFERENCES Restaurants(id)
);