-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL ,
    `password` varchar(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users` IF EXISTS
-- +goose StatementEnd
