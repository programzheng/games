-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets (
    id int UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(100),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `tickets`;
-- +goose StatementEnd
