-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_agents (
    id int UNSIGNED NOT NULL AUTO_INCREMENT,
    agent_id int UNSIGNED NOT NULL,
    user_id int UNSIGNED NOT NULL,
    third_party_id VARCHAR(255),
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    delete_at DATETIME,
    PRIMARY KEY(id),
    FOREIGN KEY (agent_id) REFERENCES agents(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_agents`
-- +goose StatementEnd
