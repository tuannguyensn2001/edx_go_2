-- +goose Up
-- +goose StatementBegin
CREATE TABLE `chapters` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255),
    `course_id` int(11),
    `order` int(11),
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `chapters`
-- +goose StatementEnd
