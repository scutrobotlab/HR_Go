CREATE DATABASE IF NOT EXISTS `hr`;
USE `hr`;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE = @@TIME_ZONE */;
/*!40103 SET TIME_ZONE = '+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;

DROP TABLE IF EXISTS `admins`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admins`
(
    `id`                  bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`                varchar(32)         NOT NULL DEFAULT '' COMMENT '管理员姓名',
    `default_standard_id` bigint(20)          NULL     DEFAULT NULL COMMENT '默认评分标准ID',
    `profile`             json                NULL     DEFAULT NULL COMMENT '管理员信息',
    `sms_enabled`         bool                NOT NULL DEFAULT false COMMENT '是否开启短信',
    `room_id`             bigint(20)          NOT NULL DEFAULT '0' COMMENT '房间ID',
    `deleted_at`          timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`          timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '管理员';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `admits`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admits`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `group`        varchar(32)         NOT NULL DEFAULT '' COMMENT '录取组别',
    `admin_id`     bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '操作录取的管理员',
    `applicant_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '申请ID',
    `deleted_at`   timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `group_admin_applicant` (`group`, `admin_id`, `applicant_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '录取';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `announce_configs`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `announce_configs`
(
    `id`     bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `status` varchar(255)        NOT NULL DEFAULT '' COMMENT '状态名称',
    `body`   text                NOT NULL DEFAULT '' COMMENT '状态内容',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '发布配置';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `applicant_questions`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `applicant_questions`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `applicant_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '申请ID',
    `question_id`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '问题ID',
    `answer`       char(1)             NULL     DEFAULT NULL COMMENT '答案',
    `deleted_at`   timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `applicant_question` (`applicant_id`, `question_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '面试者回答问题';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `applicant_times`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `applicant_times`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `applicant_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '申请ID',
    `group`        varchar(255)        NOT NULL DEFAULT '' COMMENT '申请组别',
    `time_id`      bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '时间ID',
    `room_id`      bigint(20) unsigned NULL     DEFAULT NULL COMMENT '房间ID',
    `extend`       json                NULL     DEFAULT NULL COMMENT '扩展信息',
    `deleted_at`   timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `applicant_group` (`applicant_id`, `group`),
    UNIQUE KEY `applicant_time` (`applicant_id`, `time_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '面试者选择的面试时间';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `applicants`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `applicants`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `wechat_id`  varchar(32)         NOT NULL DEFAULT '' COMMENT '微信OpenID',
    `name`       varchar(32)         NOT NULL DEFAULT '' COMMENT '姓名',
    `gender`     tinyint             NOT NULL DEFAULT 0 COMMENT '性别 0-未知 1-男性 2-女性 9-未说明',
    `phone`      varchar(16)         NOT NULL DEFAULT '' COMMENT '手机号码',
    `avatar`     varchar(255)        NOT NULL DEFAULT '' COMMENT '头像URL',
    `profile`    json                NOT NULL COMMENT '微信简介',
    `form`       json                NOT NULL COMMENT '申请表单',
    `deleted_at` timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `applicants_wechat_id_unique` (`wechat_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '面试者信息';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `evaluation_standards`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `evaluation_standards`
(
    `id`                 bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`               varchar(32)         NOT NULL DEFAULT '' COMMENT '标准名称',
    `standard`           json                NOT NULL COMMENT '评估标准',
    `last_edit_admin_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '最后编辑的管理员ID',
    `deleted_at`         timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`         timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`         timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '评估标准';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `forms`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `forms`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`       varchar(32)         NOT NULL DEFAULT '' COMMENT '显示名称',
    `key`        varchar(32)         NOT NULL DEFAULT '' COMMENT '存储键',
    `type`       varchar(32)         NOT NULL DEFAULT 'text-field' COMMENT '表单类型 {text-field,select,textarea}',
    `required`   tinyint(1)          NOT NULL DEFAULT 0 COMMENT '是否必填 0-否 1-是',
    `options`    json                NULL     DEFAULT NULL COMMENT '选项列表（若type=select）',
    `regexp`     varchar(64)         NULL     DEFAULT NULL COMMENT '正则校验',
    `max_length` int(10) unsigned    NULL     DEFAULT NULL COMMENT '最大长度',
    `deleted_at` timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `forms_key_unique` (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '申请表单';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `intents`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `intents`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `rank`         int(10) unsigned    NOT NULL DEFAULT 0 COMMENT '意愿顺序',
    `group`        varchar(255)        NOT NULL DEFAULT '' COMMENT '意愿组别',
    `parallel`     tinyint(1)          NOT NULL DEFAULT 0 COMMENT '是否平行志愿',
    `applicant_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '申请ID',
    `deleted_at`   timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '申请意向';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `migrations`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `migration` varchar(255)     NOT NULL DEFAULT '', # TODO COMMENT
    `batch`     int(11)          NOT NULL DEFAULT 0,  # TODO COMMENT
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = ''; # TODO COMMENT
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `questions`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `questions`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `question`   text                NOT NULL COMMENT '问题内容',
    `group`      varchar(255)        NOT NULL DEFAULT '' COMMENT '问题所属组别',
    `deleted_at` timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '问题';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `remarks`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `remarks`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `admin_id`     bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '管理员ID',
    `applicant_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '申请ID',
    `remark`       text                NOT NULL COMMENT '评价',
    `deleted_at`   timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `admin_id_applicant_id` (`admin_id`, `applicant_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '评价';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `scores`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `scores`
(
    `id`                 bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `admin_id`           bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '管理员ID',
    `applicant_id`       bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '申请ID',
    `group`              varchar(255)        NOT NULL DEFAULT '' COMMENT '组别',
    `score`              double(8, 2)        NOT NULL DEFAULT 0.0 COMMENT '打分',
    `standard_id`        bigint(20) unsigned NULL     DEFAULT NULL COMMENT '打分标准',
    `evaluation_details` json                NULL     DEFAULT NULL COMMENT '评价详情',
    `deleted_at`         timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`         timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`         timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `admin_id_applicant_id_group` (`admin_id`, `applicant_id`, `group`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '评分';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `time_configs`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `time_configs`
(
    `key`   varchar(255) NOT NULL DEFAULT '' COMMENT '键',
    `value` datetime     NOT NULL DEFAULT NOW() COMMENT '值',
    PRIMARY KEY (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '时间配置';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `times`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `times`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `group`      varchar(255)        NOT NULL DEFAULT '' COMMENT '组别',
    `time`       datetime            NOT NULL DEFAULT (CURTIME()) COMMENT '时间',
    `rank`       int(10) unsigned    NOT NULL DEFAULT 0 COMMENT '顺序',
    `location`   varchar(255)        NULL     DEFAULT NULL COMMENT '地点',
    `total_cnt`  int(10) unsigned    NOT NULL DEFAULT 1 COMMENT '总数',
    `campus`     varchar(255)        NULL     DEFAULT NULL COMMENT '校区',
    `grade`      varchar(255)        NULL     DEFAULT NULL COMMENT '年级',
    `meeting_id` varchar(64)         NULL     DEFAULT NULL COMMENT '会议ID',
    `deleted_at` timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '可选择的面试时间';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `applicant_sms`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `applicant_sms`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `applicant_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '申请者ID',
    `typ`          varchar(32)         NOT NULL DEFAULT '' COMMENT '类型',
    `status`       int unsigned        NOT NULL DEFAULT '0' COMMENT '状态',
    `args`         json                NOT NULL COMMENT '变量',
    `content`      text                NOT NULL COMMENT '内容',
    `deleted_at`   timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_by`   bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建者',
    PRIMARY KEY (`id`),
    KEY `applicant_sms_applicant_id_typ_status_idx` (`applicant_id`, `typ`, `status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '面试者短信';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `receive_sms`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `receive_sms`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `phone`      varchar(16)         NOT NULL DEFAULT '' COMMENT '手机号码',
    `content`    text                NOT NULL COMMENT '内容',
    `deleted_at` timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `receive_sms_phone_idx` (`phone`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '接收短信';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `configuration`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `configuration`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `key`        varchar(255)        NOT NULL DEFAULT '' COMMENT '键',
    `value`      varchar(255)        NOT NULL DEFAULT '' COMMENT '值',
    `deleted_at` timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `configuration_key_unique` (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '配置';

DROP TABLE IF EXISTS `guide`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `guide`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `group`      varchar(255)        NOT NULL DEFAULT '' COMMENT '组别',
    `guide`      json                NOT NULL COMMENT '指南',
    `deleted_at` timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `guide_group_unique` (`group`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '指南';

DROP TABLE IF EXISTS `rooms`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rooms`
(
    `id`                  bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`                varchar(64)         NOT NULL DEFAULT '' COMMENT '房间名称',
    `location`            varchar(64)         NOT NULL DEFAULT '' COMMENT '房间位置',
    `status`              tinyint unsigned    NOT NULL DEFAULT '0' COMMENT '状态 0-已停用 1-休息中 2-等待中 3-已占用',
    `status_updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '状态更新时间',
    `applicant_time_id`   bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '申请时间ID',
    `interviewer_comment` text                NOT NULL COMMENT '面试官留言',
    `receiver_comment`    text                NOT NULL COMMENT '接待者留言',
    `group_label`         varchar(64)         NOT NULL DEFAULT '' COMMENT '组别标签',
    `deleted_at`          timestamp           NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`          timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by`          bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
    PRIMARY KEY (`id`),
    UNIQUE KEY `rooms_name_unique` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '面试房间';


UNLOCK
    TABLES;
/*!40103 SET TIME_ZONE = @OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE = @OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS = @OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION = @OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES = @OLD_SQL_NOTES */;
