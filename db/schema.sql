CREATE TABLE
    `user` (
               `user_id` VARCHAR(100) NOT NULL COMMENT 'ユーザー ID',
               `user_name` VARCHAR(100) NOT NULL COMMENT 'ユーザー名',
               `password` VARCHAR(100) NOT NULL COMMENT 'パスワード',
               `created_at` DATETIME NOT NULL COMMENT '作成日時',
               `updated_at` DATETIME NOT NULL COMMENT '更新日時',
               PRIMARY KEY (`user_id`),
               INDEX `user_updated_at` (`updated_at`)
) COMMENT = 'ユーザー';

CREATE TABLE
    `channel` (
                  `channel_id` VARCHAR(100) NOT NULL COMMENT 'チャンネル ID',
                  `channel_name` VARCHAR(100) COMMENT 'チャンネル名',
                  `created_at` DATETIME NOT NULL COMMENT '作成日時',
                  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
                  PRIMARY KEY (`channel_id`),
                  INDEX `channel_updated_at` (`updated_at`)
) COMMENT = 'チャンネル';

CREATE TABLE
    `message` (
               `message_id` VARCHAR(100) NOT NULL COMMENT 'メッセージ ID',
               `message_body` VARCHAR(2000) NOT NULL COMMENT 'メッセージボディ',
               `author` VARCHAR(100) NOT NULL COMMENT 'メッセージ作者',
               `channel_id` VARCHAR(100) NOT NULL COMMENT 'チャンネル ID',
               `is_send` boolean NOT NULL COMMENT '送信状況',
               `send_at` DATETIME NOT NULL COMMENT '送信日時',
               `created_at` DATETIME NOT NULL COMMENT '作成日時',
               `updated_at` DATETIME NOT NULL COMMENT '更新日時',
               PRIMARY KEY (`message_id`),
               FOREIGN KEY (`author`) REFERENCES `user` (`user_id`),
               FOREIGN KEY (`channel_id`) REFERENCES `channel` (`channel_id`),
               INDEX `message_updated_at` (`updated_at`)
) COMMENT = 'メッセージ';

CREATE TABLE
    `joinChannelToUser` (
                  `user_id` VARCHAR(100) NOT NULL COMMENT 'ユーザー ID',
                  `user_name` VARCHAR(100) COMMENT 'ユーザー名',
                  `channel_id` VARCHAR(100) NOT NULL COMMENT 'チャンネル ID',
                  `channel_name` VARCHAR(100) COMMENT 'チャンネル名',
                  `created_at` DATETIME NOT NULL COMMENT '作成日時',
                  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
                  PRIMARY KEY (`user_id`, `channel_id`),
                  FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                  FOREIGN KEY (`channel_id`) REFERENCES `channel` (`channel_id`),
                  INDEX `join_updated_at` (`updated_at`)
) COMMENT = 'チャンネルとユーザーの中間テーブル';

# testUser1, testUser2, testUser3 の 3 ユーザーを作成

INSERT INTO `user`(user_id, user_name ,password,created_at,updated_at)
VALUES ('UserID_testUser1_12345', 'testUser1','password',NOW(),NOW());

INSERT INTO `user`(user_id, user_name ,password,created_at,updated_at)
VALUES ('UserID_testUser2_24567', 'testUser2','password',NOW(),NOW());

INSERT INTO `user`(user_id, user_name ,password,created_at,updated_at)
VALUES ('UserID_testUser3_35635', 'testUser3','password',NOW(),NOW());

# testChannel1, testChannel2 の 2 チャンネルを作成

INSERT INTO `channel`(channel_id, channel_name,created_at,updated_at)
VALUES ('ChannelID_testChannel1_12345', 'testChannel1',NOW(),NOW());

INSERT INTO `channel`(channel_id, channel_name,created_at,updated_at)
VALUES ('ChannelID_testChannel2_12645', 'testChannel2',NOW(),NOW());

# testUser1 は testChannel1, testChannel2 に所属
# testUser2 は testChannel1 に所属
# testUser3 は testChannel2 に所属

INSERT INTO `joinChannelToUser`(`user_id`,`user_name`,`channel_id`, `channel_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser1_12345', 'testUser1','ChannelID_testChannel1_12345' ,'testChannel1', NOW(),NOW());

INSERT INTO `joinChannelToUser`(`user_id`,`user_name`,`channel_id`, `channel_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser1_12345', 'testUser1','ChannelID_testChannel2_12645' ,'testChannel2', NOW(),NOW());

INSERT INTO `joinChannelToUser`(`user_id`,`user_name`,`channel_id`, `channel_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser2_24567', 'testUser2','ChannelID_testChannel1_12345' ,'testChannel1', NOW(),NOW());

INSERT INTO `joinChannelToUser`(`user_id`,`user_name`,`channel_id`, `channel_name`,`created_at`,`updated_at`)
VALUES ('UserID_testUser3_35635', 'testUser3','ChannelID_testChannel2_12645' ,'testChannel2', NOW(),NOW());

# testUser1 から送信されるメッセージ

INSERT INTO `message`(`message_id`, `message_body`,`author`,`channel_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_123123', 'こんにちは！ 私は testUser1 です。','UserID_testUser1_12345' ,'ChannelID_testChannel1_12345', true,Now(),NOW(),NOW());

INSERT INTO `message`(`message_id`, `message_body`,`author`,`channel_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_122313', 'テストメッセージです。','UserID_testUser1_12345' ,'ChannelID_testChannel1_12345', true,Now(),NOW(),NOW());

INSERT INTO `message`(`message_id`, `message_body`,`author`,`channel_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_9123', '未送信のテストメッセージです！','UserID_testUser1_12345' ,'ChannelID_testChannel1_12345', false ,NOW() + INTERVAL 100 DAY ,NOW(),NOW());


INSERT INTO `message`(`message_id`, `message_body`,`author`,`channel_id`, `is_send`, `send_at`,`created_at`,`updated_at`)
VALUES ('message_869123', 'チャンネル2 のテストメッセージです。','UserID_testUser1_12345' ,'ChannelID_testChannel2_12645', true,Now(),NOW(),NOW());
