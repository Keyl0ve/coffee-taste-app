
CREATE TABLE
    `user` (
            `user_id` VARCHAR(100) NOT NULL COMMENT 'ユーザー ID',
            `user_name` VARCHAR(100) NOT NULL COMMENT 'ユーザー名',
            `password` VARCHAR(100) NOT NULL COMMENT 'パスワード',
            `coffee_table_id` VARCHAR(100) COMMENT 'コーヒーテーブル ID',
            PRIMARY KEY (`user_id`),
) COMMENT = 'ユーザー';

CREATE TABLE
    `coffee` (
            `coffee_id` VARCHAR(100) NOT NULL COMMENT 'コーヒー ID',
            `coffee_name` VARCHAR(100) COMMENT 'コーヒー名',
            `roast_value` VARCHAR(100) COMMENT '焙煎度合い',
            `detail` VARCHAR(1000) COMMENT '詳細',
            `process` VARCHAR(100) COMMENT '生成法',
            PRIMARY KEY (`coffee_id`),
) COMMENT = 'コーヒー';

CREATE TABLE
    `user_coffee` (
            `user_id` VARCHAR(100) NOT NULL,
            `coffee_id` VARCHAR(100) NOT NULL,
            PRIMARY KEY (`user_id`, `coffee_id`),
            FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
            FOREIGN KEY (`coffee_id`) REFERENCES `coffee` (`coffee_id`)
) COMMENT = 'ユーザーとコーヒーを結ぶ中間テーブル';
