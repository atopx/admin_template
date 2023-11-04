create table user
(
    id          bigint unsigned primary key auto_increment comment '主键',
    username    varchar(64)     default '' not null comment '登录账号',
    name        varchar(64)     default '' not null comment '姓名',
    password    varchar(64)     default '' not null comment '登录密码',
    status      varchar(24)     default '' not null comment '用户状态',
    level       tinyint(4)      default 0  not null comment '用户级别',
    avatar      varchar(128)    default '' not null comment '头像',
    email       varchar(64)     default '' not null comment '邮箱',
    phone       varchar(12)     default '' not null comment '手机号',
    create_time bigint unsigned default 0  not null comment '创建时间 时间戳：秒',
    update_time bigint unsigned default 0  not null comment '更新时间 时间戳：秒',
    delete_time bigint unsigned default 0  not null comment '删除时间 时间戳：秒',
    constraint uname_index unique (username)
) comment '用户基础表';

create table user_token
(
    id            bigint unsigned primary key auto_increment comment '主键',
    user_id       bigint unsigned default 0  not null comment '用户ID',
    access_token  varchar(255)    default '' not null comment '登录token',
    refresh_token varchar(255)    default '' not null comment '登录token',
    issued_time   bigint unsigned default 0  not null comment '签发时间 时间戳：秒',
    expire_time   bigint unsigned default 0  not null comment '过期时间 时间戳：秒',
    create_time   bigint unsigned default 0  not null comment '创建时间 时间戳：秒',
    update_time   bigint unsigned default 0  not null comment '更新时间 时间戳：秒',
    delete_time   bigint unsigned default 0  not null comment '删除时间 时间戳：秒',
    constraint access_token_idx unique (access_token),
    constraint refresh_token_idx unique (refresh_token)
) comment '用户令牌表';
create index user_idx on user_token (user_id);


create table job
(
    id            bigint unsigned primary key auto_increment comment '主键',
    state         varchar(32) default '' not null comment 'job状态',
    mode          varchar(32) default '' not null comment 'job模式',
    creator       bigint unsigned default 0  not null comment '创建人 时间戳：秒',
    updater       bigint unsigned default 0  not null comment '更新人 时间戳：秒',
    create_time   bigint unsigned default 0  not null comment '创建时间 时间戳：秒',
    update_time   bigint unsigned default 0  not null comment '更新时间 时间戳：秒',
    delete_time   bigint unsigned default 0  not null comment '删除时间 时间戳：秒',

);

create table task
(
    id            bigint unsigned primary key auto_increment comment '主键',
    job_id        bigint unsigned default 0  not null comment 'JOB-ID',
    state         varchar(32) default '' not null comment 'TASK状态',
    create_time   bigint unsigned default 0  not null comment '创建时间 时间戳：秒',
    update_time   bigint unsigned default 0  not null comment '更新时间 时间戳：秒',
);

create table asset
(
    id     bigint unsigned primary key auto_increment comment '主键',
    create_time bigint unsigned default 0  not null comment '创建时间 时间戳：秒',
    update_time bigint unsigned default 0  not null comment '更新时间 时间戳：秒',
    delete_time bigint unsigned default 0  not null comment '删除时间 时间戳：秒',
);

create table label
(

);
