-- auto-generated definition
create table users
(
  id     int auto_increment comment '自动编号'
    primary key,
  uid   varchar(100) default '' not null comment '用户编码',
  uname varchar(255) default '' not null comment '用户名',
  age    int          default 0  null comment '年龄'
)
  comment '用户表';

create index users_u_id_index
  on users (uid);
