
CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "first_name" VARCHAR(50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "phone_number" VARCHAR(24) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "password" VARCHAR(24) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
CREATE TABLE "car"(

    "id" UUID PRIMARY KEY,
    "model" VARCHAR(24) NOT NULL,
    "brand" VARCHAR(24) NOT NULL,
    "state_number" VARCHAR(24) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
CREATE TABLE "branch"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(24) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
CREATE TABLE "driver"(
    "id" UUID PRIMARY KEY,
    "first_name" VARCHAR(24) NOT NULL,
    "last_name" VARCHAR(24) NOT NULL,
    "car_id" UUID ,
    "phone_number" VARCHAR(24) NOT NULL,
    "email"  VARCHAR(24) NOT NULL,
    "password" VARCHAR(24) NOT NULL,
    "branch_id" UUID,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
);
CREATE TABLE "order"(
    "id" UUID PRIMARY KEY, 
    "driver_id" UUID ,
    "date" VARCHAR(24) NOT NULL,
    "discount" VARCHAR(24) NOT NULL,
    "discount_price" INT NOT NULL,
    "branch_id" UUID ,
    "status" VARCHAR(24) NOT NULL,
    "first_client_id" UUID,
    "second_client_id" UUID ,
    "first_client_location" POINT NOT NULL, 
    "first_client_destination" POINT NOT NULL,
    "second_client_location" POINT NOT NULL,
    "second_client_destination" POINT NOT NULL,
    "millage_for_first_client" INT NOT NULL,
    "millage_for_second_client" INT NOT NULL,
    "price_for_millage" INT DEFAULT 3000,
    "total_price_for_first_client" INT NOT NULL,
    "total_price_for_second_client" INT NOT NULL,
    "payment_type" VARCHAR(24) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);