CREATE TABLE IF NOT EXISTS ReservationOrders( 
    id UUID NOT NULL,
    reservation_id UUID NOT NULL,
    menu_item_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (reservation_id) REFERENCES Reservations(id),
    FOREIGN KEY (menu_item_id) REFERENCES Menu(id)
);