/*
 Navicat Premium Data Transfer

 Source Server         : local docker
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : dtalk_record

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 10/03/2023 17:46:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS dtalk_record DEFAULT CHARACTER SET = utf8mb4;
Use dtalk_record;

-- ----------------------------
-- Table structure for dtalk_group_msg_content
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_msg_content`;
CREATE TABLE `dtalk_group_msg_content` (
  `mid` varchar(255) NOT NULL COMMENT '消息服务端序号',
  `cid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '消息客户端序号',
  `sender_id` varchar(40) NOT NULL COMMENT '发送者',
  `receiver_id` varchar(40) NOT NULL COMMENT '接收者',
  `msg_type` tinyint unsigned NOT NULL COMMENT '消息类型',
  `content` longtext NOT NULL COMMENT '消息内容',
  `create_time` bigint NOT NULL COMMENT '创建时间',
  `source` varchar(1024) DEFAULT NULL COMMENT '转发来源',
  `reference` varchar(255) DEFAULT NULL COMMENT '引用信息',
  PRIMARY KEY (`mid`) USING BTREE,
  KEY `idx_sender_id_cid` (`sender_id`,`cid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_group_msg_relation
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_msg_relation`;
CREATE TABLE `dtalk_group_msg_relation` (
  `mid` varchar(255) NOT NULL COMMENT '消息id',
  `owner_uid` varchar(40) NOT NULL COMMENT '索引用户',
  `other_uid` varchar(40) NOT NULL COMMENT '\n\n另一方用户\n',
  `type` tinyint unsigned NOT NULL COMMENT '0->发件箱；1->收件箱',
  `create_time` bigint NOT NULL COMMENT '\n\n创建时间\n',
  PRIMARY KEY (`mid`,`owner_uid`) USING BTREE,
  KEY `idx_owneruid_otheruid` (`owner_uid`,`other_uid`) USING BTREE,
  KEY `idx_owneruid_type` (`owner_uid`,`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_msg_content
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_msg_content`;
CREATE TABLE `dtalk_msg_content` (
  `mid` varchar(255) NOT NULL COMMENT '消息服务端序列号',
  `cid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '消息客户端序列号',
  `sender_id` varchar(40) NOT NULL COMMENT '发送者',
  `receiver_id` varchar(40) NOT NULL COMMENT '接收者',
  `msg_type` tinyint unsigned NOT NULL COMMENT '消息类型',
  `content` longtext NOT NULL COMMENT '消息内容',
  `create_time` bigint NOT NULL COMMENT '创建时间',
  `source` varchar(1024) DEFAULT NULL COMMENT '转发来源',
  `reference` varchar(255) DEFAULT NULL COMMENT '引用信息',
  PRIMARY KEY (`mid`) USING BTREE,
  KEY `idx_sender_id_cid` (`sender_id`,`cid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_msg_relation
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_msg_relation`;
CREATE TABLE `dtalk_msg_relation` (
  `mid` varchar(255) NOT NULL COMMENT '消息id',
  `owner_uid` varchar(40) NOT NULL COMMENT '索引用户',
  `other_uid` varchar(40) NOT NULL COMMENT '\n\n另一方用户\n',
  `type` tinyint unsigned NOT NULL COMMENT '0->发件箱；1->收件箱',
  `create_time` bigint NOT NULL COMMENT '\n\n创建时间\n',
  PRIMARY KEY (`mid`,`owner_uid`) USING BTREE,
  KEY `idx_owneruid_otheruid` (`owner_uid`,`other_uid`) USING BTREE,
  KEY `idx_owneruid_type` (`owner_uid`,`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_msg_version
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_msg_version`;
CREATE TABLE `dtalk_msg_version` (
  `uid` varchar(40) NOT NULL COMMENT '\n\n用户id\n',
  `version` bigint DEFAULT NULL COMMENT '版本号',
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_signal_content
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_signal_content`;
CREATE TABLE `dtalk_signal_content` (
  `uid` varchar(40) NOT NULL COMMENT '接收者',
  `seq` bigint NOT NULL COMMENT '消息id',
  `type` tinyint DEFAULT NULL COMMENT '通知类型',
  `content` varchar(1024) DEFAULT NULL COMMENT '通知内容',
  `create_time` bigint DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`uid`,`seq`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
