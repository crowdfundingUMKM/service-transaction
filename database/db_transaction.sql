-- Table: campaign_details

CREATE TABLE `transactions` (
  `id` VARCHAR(24) NOT NULL PRIMARY KEY,
  `unix_id` VARCHAR(12),
  `campaign_id` VARCHAR(255),
  `user_investor_id` INT,
  `order_id` VARCHAR(255),
  `payment_type` VARCHAR(255),
  `amount` BIGINT,
  `code` VARCHAR(255),
  `status_payment` VARCHAR(255),
  `expiry_time` DATETIME,
  `fraud` VARCHAR(255),
  `url_payment` VARCHAR(255),
  `token` VARCHAR(255),
  `created_at` DATETIME,
  `updated_at` DATETIME
);



-- insert data

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
