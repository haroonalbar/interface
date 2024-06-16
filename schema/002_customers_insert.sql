-- +goose Up
INSERT INTO customers (id, name) VALUES
(1,'john'),
(2,'doe'),
(3,'di'),
(4,'dey'),
(5,'domm'),
(6,'larry');

-- +goose Down
TRUNCATE TABLE customers;
