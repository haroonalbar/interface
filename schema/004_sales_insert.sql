-- +goose Up
INSERT INTO sales (id, name) VALUES
(1,'juice'),
(2,'sandwitch'),
(3,'burger'),
(4,'ice cream');

-- +goose Down
TRUNCATE TABLE sales;
