-- +goose Up
CREATE TABLE sales (
  id INT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE sales;
