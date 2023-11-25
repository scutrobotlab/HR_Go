create database if not exists hr;

create table hr.admins
(
    id                  bigint auto_increment
        primary key,
    name                longtext    not null comment '管理员姓名',
    default_standard_id bigint      null comment '默认评分标准ID',
    profile             longtext    null comment '管理员信息',
    sms_enabled         tinyint(1)  not null comment '是否开启短信',
    room_id             bigint      not null comment '房间ID',
    deleted_at          datetime(3) null,
    created_at          datetime(3) null,
    updated_at          datetime(3) null
);

create table hr.admits
(
    id           bigint auto_increment
        primary key,
    `group`      longtext    not null comment '录取组别',
    admin_id     bigint      not null comment '操作录取的管理员',
    applicant_id bigint      not null comment '申请ID',
    deleted_at   datetime(3) null,
    created_at   datetime(3) null,
    updated_at   datetime(3) null
);

create table hr.announce_configs
(
    id     bigint auto_increment
        primary key,
    status longtext not null comment '状态名称',
    body   longtext not null comment '状态内容'
);

create table hr.applicant_questions
(
    id           bigint auto_increment
        primary key,
    applicant_id bigint      not null comment '申请ID',
    question_id  bigint      not null comment '问题ID',
    answer       longtext    null comment '答案',
    deleted_at   datetime(3) null,
    created_at   datetime(3) null,
    updated_at   datetime(3) null
);

create table hr.applicant_sms
(
    id           bigint auto_increment
        primary key,
    applicant_id bigint      not null comment '申请者ID',
    typ          longtext    not null comment '类型',
    status       int         not null comment '状态',
    args         longtext    not null comment '变量',
    content      longtext    not null comment '内容',
    created_by   bigint      not null comment '创建者',
    deleted_at   datetime(3) null,
    created_at   datetime(3) null,
    updated_at   datetime(3) null
);

create table hr.applicant_times
(
    id           bigint auto_increment
        primary key,
    applicant_id bigint      not null comment '申请ID',
    `group`      longtext    not null comment '申请组别',
    time_id      bigint      not null comment '时间ID',
    room_id      bigint      not null comment '房间ID',
    extend       longtext    null comment '扩展信息',
    deleted_at   datetime(3) null,
    created_at   datetime(3) null,
    updated_at   datetime(3) null
);

create table hr.applicants
(
    id         bigint auto_increment
        primary key,
    wechat_id  longtext    not null comment '微信OpenID',
    name       longtext    not null comment '姓名',
    gender     int         not null comment '性别 0-未知 1-男性 2-女性 9-未说明',
    phone      longtext    not null comment '手机号码',
    avatar     longtext    not null comment '头像URL',
    profile    longtext    not null comment '微信简介',
    form       longtext    not null comment '申请表单',
    deleted_at datetime(3) null,
    created_at datetime(3) null,
    updated_at datetime(3) null
);

create table hr.configuration
(
    id         bigint auto_increment
        primary key,
    `key`      longtext    not null comment '键',
    value      longtext    not null comment '值',
    deleted_at datetime(3) null,
    created_at datetime(3) null,
    updated_at datetime(3) null
);

create table hr.evaluation_standards
(
    id                 bigint auto_increment
        primary key,
    name               longtext    not null comment '标准名称',
    standard           longtext    not null comment '评估标准',
    last_edit_admin_id bigint      not null comment '最后编辑的管理员ID',
    deleted_at         datetime(3) null,
    created_at         datetime(3) null,
    updated_at         datetime(3) null
);

create table hr.forms
(
    id         bigint auto_increment
        primary key,
    name       longtext                          not null comment '显示名称',
    `key`      longtext                          not null comment '存储键',
    type       varchar(191) default 'text-field' not null comment '表单类型 {text-field,select,textarea}',
    required   tinyint(1)                        not null comment '是否必填 0-否 1-是',
    options    longtext                          null comment '选项列表（若type=select）',
    `regexp`   longtext                          null comment '正则校验',
    max_length int                               null comment '最大长度',
    deleted_at datetime(3)                       null,
    created_at datetime(3)                       null,
    updated_at datetime(3)                       null
);

create table hr.guide
(
    id         bigint auto_increment
        primary key,
    `group`    longtext    not null comment '组别',
    guide      longtext    not null comment '指南',
    deleted_at datetime(3) null,
    created_at datetime(3) null,
    updated_at datetime(3) null
);

create table hr.intents
(
    id           bigint auto_increment
        primary key,
    `rank`       int         not null comment '意愿顺序',
    `group`      longtext    not null comment '意愿组别',
    parallel     tinyint(1)  not null comment '是否平行志愿',
    applicant_id bigint      not null comment '申请ID',
    deleted_at   datetime(3) null,
    created_at   datetime(3) null,
    updated_at   datetime(3) null
);

create table hr.questions
(
    id         bigint auto_increment
        primary key,
    question   longtext    not null comment '问题内容',
    `group`    longtext    not null comment '问题所属组别',
    deleted_at datetime(3) null,
    created_at datetime(3) null,
    updated_at datetime(3) null
);

create table hr.receive_sms
(
    id         bigint auto_increment
        primary key,
    phone      longtext    not null comment '手机号码',
    content    longtext    not null comment '内容',
    deleted_at datetime(3) null,
    created_at datetime(3) null,
    updated_at datetime(3) null
);

create table hr.remarks
(
    id           bigint auto_increment
        primary key,
    admin_id     bigint      not null comment '管理员ID',
    applicant_id bigint      not null comment '申请ID',
    remark       longtext    not null comment '评价',
    deleted_at   datetime(3) null,
    created_at   datetime(3) null,
    updated_at   datetime(3) null
);

create table hr.rooms
(
    id                  bigint auto_increment
        primary key,
    name                longtext    not null comment '房间名称',
    location            longtext    not null comment '房间位置',
    status              int         not null comment '状态 0-已停用 1-休息中 2-等待中 3-已占用',
    status_updated_at   datetime(3) not null comment '状态更新时间',
    applicant_time_id   bigint      not null comment '申请时间ID',
    interviewer_comment longtext    null comment '面试官留言',
    receiver_comment    longtext    null comment '接待者留言',
    group_label         longtext    not null comment '组别标签',
    updated_by          bigint      not null comment '更新者',
    deleted_at          datetime(3) null,
    created_at          datetime(3) null,
    updated_at          datetime(3) null
);

create table hr.scores
(
    id                 bigint auto_increment
        primary key,
    admin_id           bigint      not null comment '管理员ID',
    applicant_id       bigint      not null comment '申请ID',
    `group`            longtext    not null comment '组别',
    score              double      not null comment '打分',
    standard_id        bigint      null comment '打分标准',
    evaluation_details longtext    null comment '评价详情',
    deleted_at         datetime(3) null,
    created_at         datetime(3) null,
    updated_at         datetime(3) null
);

create table hr.time_configs
(
    `key` varchar(191) not null comment '键'
        primary key,
    value datetime(3)  not null comment '值'
);

create table hr.times
(
    id         bigint auto_increment
        primary key,
    `group`    longtext      not null comment '组别',
    time       datetime(3)   not null comment '时间',
    `rank`     int           not null comment '顺序',
    location   longtext      null comment '地点',
    total_cnt  int default 1 not null comment '总数',
    campus     longtext      null comment '校区',
    grade      longtext      null comment '年级',
    meeting_id longtext      null comment '会议ID',
    deleted_at datetime(3)   null,
    created_at datetime(3)   null,
    updated_at datetime(3)   null
);
