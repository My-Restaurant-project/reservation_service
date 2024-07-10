CREATE TYPE status AS ENUM ('active', 'inactive', 'pending');

CREATE TABLE IF NOT EXISTS "Reservations"(
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "restaurant_id" UUID,
    "reservation_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "status" status  DEFAULT 'pending',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("restaurant_id") REFERENCES "Restaurants"("id")
);
