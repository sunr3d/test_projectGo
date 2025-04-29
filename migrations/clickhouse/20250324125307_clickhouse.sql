-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links_from_kafka
(
    link       String,
    fake_link  String,
    created_at DateTime DEFAULT now()
)
ENGINE = Kafka('localhost:29092', 'link_service', 'clickhouse_group', 'JSONEachRow');

CREATE TABLE IF NOT EXISTS links
(
    link       String,
    fake_link  String,
    created_at DateTime DEFAULT now()
)
ENGINE = MergeTree()
ORDER BY created_at;

CREATE MATERIALIZED VIEW IF NOT EXISTS links_mv TO links AS
SELECT
    link,
    fake_link,
    created_at
FROM links_from_kafka;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS links_mv;
DROP TABLE IF EXISTS links_from_kafka;
-- +goose StatementEnd
