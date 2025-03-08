CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE income (
    amount FLOAT NOT NULL,
    "date" DATE NOT NULL,
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "month" INTEGER NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "received" BOOLEAN NOT NULL
);