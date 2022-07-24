-- +goose Up
-- +goose StatementBegin
CREATE TABLE agents (
    id int UNSIGNED NOT NULL AUTO_INCREMENT,
    code VARCHAR(100) NOT NULL,
    name VARCHAR(100),
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    delete_at DATETIME,
    PRIMARY KEY(id),
    UNIQUE (code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `agents`
-- +goose StatementEnd
