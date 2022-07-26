-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id int UNSIGNED NOT NULL AUTO_INCREMENT,
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    delete_at DATETIME,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd
