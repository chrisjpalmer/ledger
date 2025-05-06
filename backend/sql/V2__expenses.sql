
CREATE TABLE expenses (
    amount FLOAT NOT NULL,
    "date" DATE NOT NULL,
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "month" INTEGER NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "paid" BOOLEAN NOT NULL
);
