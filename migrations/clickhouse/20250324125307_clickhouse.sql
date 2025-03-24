-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links
(
    link       String,
    fake_link  String,
    created_at DateTime DEFAULT now()
)
ENGINE = Kafka('localhost:29092', 'link_service', 'clickhouse_group', 'JSONEachRow');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
