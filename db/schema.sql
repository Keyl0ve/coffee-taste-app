CREATE TABLE
    `user` (
            `user_id` VARCHAR(100) NOT NULL COMMENT 'ユーザー ID',
            `user_name` VARCHAR(100) NOT NULL COMMENT 'ユーザー名',
            `password` VARCHAR(100) NOT NULL COMMENT 'パスワード',
            `coffee_table_id` VARCHAR(100) COMMENT 'コーヒーテーブル ID',
            `created_at` DATETIME NOT NULL COMMENT '作成日時',
            `updated_at` DATETIME NOT NULL COMMENT '更新日時',
            PRIMARY KEY (`user_id`)
) COMMENT = 'ユーザー';

CREATE TABLE
    `coffee` (
            `coffee_id` VARCHAR(100) NOT NULL COMMENT 'コーヒー ID',
            `coffee_name` VARCHAR(100) COMMENT 'コーヒー名',
            `roast_value` VARCHAR(100) COMMENT '焙煎度合い',
            `detail` VARCHAR(1000) COMMENT '詳細',
            `process` VARCHAR(100) COMMENT '生成法',
            `created_at` DATETIME NOT NULL COMMENT '作成日時',
            `updated_at` DATETIME NOT NULL COMMENT '更新日時',
            PRIMARY KEY (`coffee_id`)
) COMMENT = 'コーヒー';

# testUser1, testUser2, testUser3 の 3 ユーザーを作成

INSERT INTO `user`(user_id, user_name ,password, coffee_table_id)
VALUES ('UserID_testUser1_12345', 'testUser1', 'password', 'UserID_testUser1_12345');

INSERT INTO `user`(user_id, user_name ,password)
VALUES ('UserID_testUser2_24567', 'testUser2', 'password', 'UserID_testUser2_24567');

INSERT INTO `user`(user_id, user_name ,password)
VALUES ('UserID_testUser3_35635', 'testUser3', 'password', 'UserID_testUser3_35635');

# testCoffee1, testCoffee2 の 2 つのコーヒーを作成

INSERT INTO `coffee`(coffee_id, coffee_name, roast_value, detail, process)
VALUES ('CoffeeID_testCoffee1_12345', 'Ethiopia', 'Light', 'エチオピアのコーヒーです。', 'Washed');

INSERT INTO `coffee`(coffee_id, coffee_name, roasted_value, detail, process)
VALUES ('CoffeeID_testCoffee2_12645', 'Guatemala', 'Medium', 'グアテマラのコーヒーです。', 'Washed');





CREATE TABLE
    `joinCoffeeToUser` (
                `user_id` VARCHAR(100) NOT NULL COMMENT 'ユーザー ID',
                `user_name` VARCHAR(100) COMMENT 'ユーザー名',
                `coffee_id` VARCHAR(100) NOT NULL COMMENT 'チャンネル ID',
                `coffee_name` VARCHAR(100) COMMENT 'チャンネル名',
                PRIMARY KEY (`user_id`, `coffee_id`),
                FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                FOREIGN KEY (`coffee_id`) REFERENCES `coffee` (`coffee_id`),
) COMMENT = 'コーヒーとユーザーの中間テーブル';

# testUser1, testUser2, testUser3 の 3 ユーザーを作成

INSERT INTO `user`(user_id, user_name ,password,created_at,updated_at)
VALUES ('UserID_testUser1_12345', 'testUser1','password',NOW(),NOW());

INSERT INTO `user`(user_id, user_name ,password,created_at,updated_at)
VALUES ('UserID_testUser2_24567', 'testUser2','password',NOW(),NOW());

INSERT INTO `user`(user_id, user_name ,password,created_at,updated_at)
VALUES ('UserID_testUser3_35635', 'testUser3','password',NOW(),NOW());

# testCoffee1, testCoffee2 の 2 チャンネルを作成

INSERT INTO `coffee`(coffee_id, coffee_name,created_at,updated_at)
VALUES ('CoffeeID_testCoffee1_12345', 'testCoffee1',NOW(),NOW());

INSERT INTO `coffee`(coffee_id, coffee_name,created_at,updated_at)
VALUES ('CoffeeID_testCoffee2_12645', 'testCoffee2',NOW(),NOW());

# testUser1 は testCoffee1, testCoffee2 に所属
# testUser2 は testCoffee1 に所属
# testUser3 は testCoffee2 に所属

INSERT INTO `joinCoffeeToUser`(`user_id`,`user_name`,`coffee_id`, `coffee_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser1_12345', 'testUser1','CoffeeID_testCoffee1_12345' ,'testCoffee1', NOW(),NOW());

INSERT INTO `joinCoffeeToUser`(`user_id`,`user_name`,`coffee_id`, `coffee_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser1_12345', 'testUser1','CoffeeID_testCoffee2_12645' ,'testCoffee2', NOW(),NOW());

INSERT INTO `joinCoffeeToUser`(`user_id`,`user_name`,`coffee_id`, `coffee_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser2_24567', 'testUser2','CoffeeID_testCoffee1_12345' ,'testCoffee1', NOW(),NOW());

INSERT INTO `joinCoffeeToUser`(`user_id`,`user_name`,`coffee_id`, `coffee_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser3_35635', 'testUser3','CoffeeID_testCoffee2_12645' ,'testCoffee2', NOW(),NOW());

# testUser1 から送信されるメッセージ

INSERT INTO `message`(`message_id`, `message_body`,`author`,`coffee_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_123123', 'こんにちは！ 私は testUser1 です。','UserID_testUser1_12345' ,'CoffeeID_testCoffee1_12345', true,Now(),NOW(),NOW());

INSERT INTO `message`(`message_id`, `message_body`,`author`,`coffee_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_122313', 'テストメッセージです。','UserID_testUser1_12345' ,'CoffeeID_testCoffee1_12345', true,Now(),NOW(),NOW());

INSERT INTO `message`(`message_id`, `message_body`,`author`,`coffee_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_9123', '未送信のテストメッセージです！','UserID_testUser1_12345' ,'CoffeeID_testCoffee1_12345', false ,NOW() + INTERVAL 100 DAY ,NOW(),NOW());


INSERT INTO `message`(`message_id`, `message_body`,`author`,`coffee_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_869123', 'チャンネル2 のテストメッセージです。','UserID_testUser1_12345' ,'CoffeeID_testCoffee2_12645', true,Now(),NOW(),NOW());
