CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `full_name` varchar(255),
  `gender` tinyint NOT NULL DEFAULT 0,
  `profile_url` varchar(255),
  `status` tinyint NOT NULL DEFAULT 0,
  `city_id` int,
  `is_verified` tinyint,
  `verified_until` timestamp,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `cities` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL
);

CREATE TABLE `swipes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `target_id` int NOT NULL,
  `type` tinyint NOT NULL,
  `is_matched` tinyint NOT NULL,
  `match_id` int,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `matches` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `target_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE `users` ADD FOREIGN KEY (`city_id`) REFERENCES `cities` (`id`);

ALTER TABLE `swipes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `swipes` ADD FOREIGN KEY (`target_id`) REFERENCES `users` (`id`);

ALTER TABLE `swipes` ADD FOREIGN KEY (`match_id`) REFERENCES `matches` (`id`);

ALTER TABLE `matches` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `matches` ADD FOREIGN KEY (`target_id`) REFERENCES `users` (`id`);
