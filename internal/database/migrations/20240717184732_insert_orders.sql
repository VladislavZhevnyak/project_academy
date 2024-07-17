-- +goose Up
-- +goose StatementBegin
INSERT INTO orders
VALUES
	( 1, '2024-07-18', 'Vlad', 'TV', 'Minsk', 'not delivered' ),
	( 2, '2024-07-20', 'Stas', 'fridge', 'Minsk', 'delivered' ),
	( 3, '2024-07-25', 'Kirill', 'washing machine', 'Minsk', 'failure' );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
