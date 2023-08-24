USE app;

CREATE TABLE
    `user` (
            `user_id` VARCHAR(100) NOT NULL COMMENT 'ユーザー ID',
            `user_name` VARCHAR(100) NOT NULL COMMENT 'ユーザー名',
            `password` VARCHAR(100) NOT NULL COMMENT 'パスワード',
            `created_at` DATETIME NOT NULL COMMENT '作成日時',
            `updated_at` DATETIME NOT NULL COMMENT '更新日時',
            PRIMARY KEY (`user_id`)
) COMMENT = 'ユーザー';

CREATE TABLE
    `coffee` (
            `coffee_id` VARCHAR(100) NOT NULL COMMENT 'コーヒー ID',
            `coffee_name` VARCHAR(100) COMMENT 'コーヒー名',
            `created_at` DATETIME NOT NULL COMMENT '作成日時',
            `updated_at` DATETIME NOT NULL COMMENT '更新日時',
            PRIMARY KEY (`coffee_id`)
) COMMENT = 'コーヒー';


INSERT INTO `user`(user_id, user_name, password, created_at, updated_at)
VALUES ('UserID_testUser1_12345', 'testUser1', 'password', NOW(), NOW());

INSERT INTO `user`(user_id, user_name, password, created_at, updated_at)
VALUES ('UserID_testUser2_24567', 'testUser2', 'password', NOW(), NOW());

INSERT INTO `user`(user_id, user_name, password, created_at, updated_at)
VALUES ('UserID_testUser3_35635', 'testUser3', 'password', NOW(), NOW());

INSERT INTO `coffee` (coffee_id, coffee_name, created_at, updated_at)
VALUES ('CoffeeID_testCoffee1_12345', 'Ethiopia', NOW(), NOW());

INSERT INTO `coffee` (coffee_id, coffee_name, created_at, updated_at)
VALUES ('CoffeeID_testCoffee2_12645', 'Guatemala', NOW(), NOW());