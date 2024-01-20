-- Table: campaign_details

CREATE TABLE `campaign_details` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `unix_id` CHAR(12),
  `user_campaign_id` CHAR(12),
  `name_campaign` VARCHAR(255),
  `category` VARCHAR(255),
  `description` TEXT,
  `address_campaign` VARCHAR(255),
  `goal_amount` BIGINT,
  `current_amount` BIGINT,
  `minimum_invest` BIGINT,
  `interest_rate` INT(11),
  `tenor_period` INT(11),
  `deadline_campaign` DATETIME,
  `business_proposal` VARCHAR(255),
  `perks` TEXT,
  `link_bisnis` VARCHAR(255),
  `status_campaign` VARCHAR(10),
  `slug` VARCHAR(255),
  `rating` TINYINT(1),
  `done_campaign` VARCHAR(10),
  `refund` VARCHAR(10),
  `status_refund` VARCHAR(10),
  `updateId_reviewer` CHAR(12),
  `updateId_admin` VARCHAR(11),
  `updateAt_admin` DATETIME,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Table: campaign_images

CREATE TABLE `campaign_images` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `campaign_id` CHAR(12),
  `file_name` VARCHAR(255),
  `is_primary` TINYINT(4),
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Table: application_funds

CREATE TABLE `application_funds` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `campaign_id` CHAR(12),
  `status_submission` VARCHAR(10),
  `status_adminId` CHAR(12),
  `submission_at` DATETIME,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- insert data

-- Indexes for table `users`
--
-- ALTER TABLE `users`
--   ADD PRIMARY KEY (`id`);

-- Remove token from table users
-- DELIMITER //

-- CREATE EVENT delete_expired_tokens
-- ON SCHEDULE EVERY 1 HOUR
-- DO
-- BEGIN
--     DELETE FROM users
--     WHERE token IS NOT NULL
--     AND created_at < NOW() - INTERVAL 2 DAY;
-- END //

-- DELIMITER ;

-- Backup database
-- SELECT *
-- INTO OUTFILE '/path/to/backup/users_backup.csv'
-- FIELDS TERMINATED BY ','
-- ENCLOSED BY '"'
-- LINES TERMINATED BY '\n'
-- FROM users;
