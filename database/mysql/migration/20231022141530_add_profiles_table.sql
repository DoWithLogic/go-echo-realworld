-- +goose Up
-- +goose StatementBegin
CREATE TABLE `profiles` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `follow_user_id` INT NOT NULL,
  `is_active` BOOLEAN NOT NULL,
  `created_at` timestamp DEFAULT now(),
  `created_by` varchar(255) NOT NULL,
  `updated_at` timestamp DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `profiles`;
-- +goose StatementEnd
