--  Copyright © 2020. Drew Lee. All rights reserved.

create table contacts
(
    id int auto_increment comment 'contact_id'
        primary key,
    owner_id int null comment '归属者member_id',
    dst_id int null comment '目标(group/member)_id',
    cat int null comment '类型',
    memo varchar(150) null comment '备注',
    created_at datetime null comment '创建时间',
    updated_at datetime null comment '更新时间',
    deleted_at datetime null comment '删除时间'
);

create table `groups`
(
    id int auto_increment comment 'group_id'
        primary key,
    name varchar(30) not null comment '群名',
    owner_id int null comment '群主',
    icon varchar(250) null comment '群logo',
    cat int null comment '类型',
    memo varchar(150) null comment '备注',
    created_at datetime null comment '创建时间',
    updated_at datetime null comment '更新时间',
    deleted_at datetime null comment '删除时间'
);

create table members
(
    id int auto_increment comment 'member_id'
        primary key,
    phone_num varchar(20) default '' null comment '手机号',
    password varchar(40) not null comment '密码',
    avatar varchar(150) default '' null comment '头像',
    gender varchar(2) default 'U' null comment '性别',
    nickname varchar(20) default '' null comment '昵称',
    email varchar(20) default '' null comment '邮箱',
    salt varchar(10) null comment '密码盐值',
    online int null comment '是否在线',
    token varchar(40) null comment 'token',
    memo varchar(150) null comment '备注',
    created_at datetime null comment '创建时间',
    updated_at datetime null comment '更新时间',
    deleted_at datetime null comment '删除时间'
);
