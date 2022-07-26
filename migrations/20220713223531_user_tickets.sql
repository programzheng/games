-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_tickets (
    id int UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id int UNSIGNED,
    ticket_id int UNSIGNED NOT NULL,
    code VARCHAR(100) NOT NULL,
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    delete_at DATETIME,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ticket_id) REFERENCES tickets(id),
    UNIQUE (code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_tickets`;
-- +goose StatementEnd
