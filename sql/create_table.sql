--  Copyright © 2020. Drew Lee. All rights reserved.

-- BEGIN TABLE contacts
CREATE TABLE `contacts` (
                            `id` int NOT NULL AUTO_INCREMENT COMMENT 'contact_id',
                            `owner_id` int DEFAULT NULL COMMENT '归属者member_id',
                            `dst_id` int DEFAULT NULL COMMENT '目标(group/member)_id',
                            `cat` int DEFAULT NULL COMMENT '类型',
                            `memo` varchar(150) DEFAULT NULL COMMENT '备注',
                            `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                            `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                            `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- END TABLE contacts

-- BEGIN TABLE groups
CREATE TABLE `groups` (
                          `id` int NOT NULL AUTO_INCREMENT COMMENT 'group_id',
                          `name` varchar(30) NOT NULL COMMENT '群名',
                          `owner_id` int DEFAULT NULL COMMENT '群主',
                          `icon` varchar(250) DEFAULT NULL COMMENT '群logo',
                          `cat` int DEFAULT NULL COMMENT '类型',
                          `memo` varchar(150) DEFAULT NULL COMMENT '备注',
                          `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                          `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                          `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- END TABLE groups

-- BEGIN TABLE members
CREATE TABLE `members` (
                           `id` int NOT NULL AUTO_INCREMENT COMMENT 'member_id',
                           `phone_num` varchar(20) DEFAULT '' COMMENT '手机号',
                           `password` varchar(40) NOT NULL COMMENT '密码',
                           `avatar` varchar(150) DEFAULT '' COMMENT '头像',
                           `gender` varchar(2) DEFAULT 'U' COMMENT '性别',
                           `nickname` varchar(20) DEFAULT '' COMMENT '昵称',
                           `salt` varchar(10) DEFAULT NULL COMMENT '密码盐值',
                           `online` int DEFAULT NULL COMMENT '是否在线',
                           `token` varchar(40) DEFAULT NULL COMMENT 'token',
                           `memo` varchar(150) DEFAULT NULL COMMENT '备注',
                           `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                           `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                           `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- END TABLE members
