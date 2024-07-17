-- +goose Up
-- +goose StatementBegin
CREATE TABLE  Orders (
order_id INT NOT NULL PRIMARY KEY,
order_date VARCHAR ( 30 ) NOT NULL DEFAULT '0000-00-00',
customers_name VARCHAR ( 30 ) NOT NULL DEFAULT '-',
product VARCHAR ( 255 ) NOT NULL,
adress VARCHAR ( 255 ) NOT NULL,
order_status VARCHAR ( 30 ) DEFAULT 'not delivered' 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
