
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
