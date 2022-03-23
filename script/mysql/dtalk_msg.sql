/*
 Navicat Premium Data Transfer

 Source Server         : 172.16.101.127
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : 172.16.101.127:3306
 Source Schema         : dtalk

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 18/11/2021 10:37:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dtalk_group_msg_content
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_msg_content`;
CREATE TABLE `dtalk_group_msg_content` (
  `mid` bigint(20) unsigned NOT NULL COMMENT '\n\n消息id\n',
  `seq` varchar(40) NOT NULL COMMENT '消息序列号',
  `sender_id` varchar(40) NOT NULL COMMENT '发送者',
  `receiver_id` varchar(40) NOT NULL COMMENT '接收者',
  `msg_type` tinyint(3) unsigned NOT NULL COMMENT '消息类型',
  `content` longtext NOT NULL COMMENT '消息内容',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `source` varchar(1024) DEFAULT NULL COMMENT '转发来源',
  PRIMARY KEY (`mid`) USING BTREE,
  KEY `idx_sender_id_seq` (`sender_id`,`seq`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_group_msg_relation
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_msg_relation`;
CREATE TABLE `dtalk_group_msg_relation` (
  `mid` bigint(20) unsigned NOT NULL COMMENT '消息id',
  `owner_uid` varchar(40) NOT NULL COMMENT '索引用户',
  `other_uid` varchar(40) NOT NULL COMMENT '\n\n另一方用户\n',
  `type` tinyint(3) unsigned NOT NULL COMMENT '0->发件箱；1->收件箱',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0->未接受；1->已接收',
  `create_time` bigint(20) NOT NULL COMMENT '\n\n创建时间\n',
  PRIMARY KEY (`mid`,`owner_uid`) USING BTREE,
  KEY `idx_owneruid_otheruid_msgid` (`owner_uid`,`other_uid`,`mid`) USING BTREE,
  KEY `idx_owneruid_type_state` (`owner_uid`,`type`,`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_msg_content
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_msg_content`;
CREATE TABLE `dtalk_msg_content` (
  `mid` bigint(20) unsigned NOT NULL COMMENT '\n\n消息id\n',
  `seq` varchar(40) NOT NULL COMMENT '消息序列号',
  `sender_id` varchar(40) NOT NULL COMMENT '发送者',
  `receiver_id` varchar(40) NOT NULL COMMENT '接收者',
  `msg_type` tinyint(3) unsigned NOT NULL COMMENT '消息类型',
  `content` longtext NOT NULL COMMENT '消息内容',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `source` varchar(1024) DEFAULT NULL COMMENT '转发来源',
  PRIMARY KEY (`mid`) USING BTREE,
  KEY `idx_sender_id_seq` (`sender_id`,`seq`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_msg_relation
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_msg_relation`;
CREATE TABLE `dtalk_msg_relation` (
  `mid` bigint(20) unsigned NOT NULL COMMENT '消息id',
  `owner_uid` varchar(40) NOT NULL COMMENT '索引用户',
  `other_uid` varchar(40) NOT NULL COMMENT '\n\n另一方用户\n',
  `type` tinyint(3) unsigned NOT NULL COMMENT '0->发件箱；1->收件箱',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0->未接受；1->已接收',
  `create_time` bigint(20) NOT NULL COMMENT '\n\n创建时间\n',
  PRIMARY KEY (`mid`,`owner_uid`) USING BTREE,
  KEY `idx_owneruid_otheruid_msgid` (`owner_uid`,`other_uid`,`mid`) USING BTREE,
  KEY `idx_owneruid_type_state` (`owner_uid`,`type`,`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_msg_version
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_msg_version`;
CREATE TABLE `dtalk_msg_version` (
  `uid` varchar(40) NOT NULL COMMENT '\n\n用户id\n',
  `version` bigint(20) DEFAULT NULL COMMENT '版本号',
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dtalk_notice_content
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_notice_content`;
CREATE TABLE `dtalk_notice_content` (
  `id` bigint(20) NOT NULL COMMENT '消息id',
  `uid` varchar(40) NOT NULL COMMENT '接收者',
  `type` tinyint(3) DEFAULT NULL COMMENT '通知类型',
  `state` tinyint(3) DEFAULT NULL COMMENT '0->未接收；1->已接收',
  `content` varchar(1024) DEFAULT NULL COMMENT '通知内容',
  `create_time` bigint(20) DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`uid`) USING BTREE,
  KEY `idx_uid_state` (`uid`,`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
