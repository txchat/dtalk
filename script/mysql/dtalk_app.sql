/*
 Navicat Premium Data Transfer

 Source Server         : 172.16.101.107
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 172.16.101.107:3306
 Source Schema         : dtalk

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 18/11/2021 10:36:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dtalk_addr_backup
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_addr_backup`;
CREATE TABLE `dtalk_addr_backup` (
  `address` varchar(255) NOT NULL COMMENT '用户地址',
  `area` varchar(4) DEFAULT NULL COMMENT '区号',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `mnemonic` varchar(1020) DEFAULT NULL COMMENT '助记词',
  `private_key` varchar(1020) DEFAULT NULL COMMENT '加密私钥',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`address`),
  KEY `idx_phone` (`phone`) USING HASH,
  KEY `idx_email` (`email`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dtalk_addr_relate
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_addr_relate`;
CREATE TABLE `dtalk_addr_relate` (
  `address` varchar(255) NOT NULL COMMENT '用户地址',
  `area` varchar(4) DEFAULT NULL COMMENT '区号',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `mnemonic` varchar(1020) DEFAULT NULL COMMENT '助记词',
  `private_key` varchar(1020) DEFAULT NULL COMMENT '加密私钥',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`address`),
  KEY `idx_phone` (`phone`) USING HASH,
  KEY `idx_email` (`email`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dtalk_cdk_info
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_cdk_info`;
CREATE TABLE `dtalk_cdk_info` (
  `cdk_id` bigint(20) NOT NULL COMMENT '兑换码id',
  `cdk_name` varchar(255) NOT NULL COMMENT '兑换码名称',
  `cdk_info` varchar(255) DEFAULT NULL COMMENT '兑换码详情',
  `coin_name` varchar(255) NOT NULL COMMENT '票券名称',
  `exchange_rate` bigint(20) NOT NULL COMMENT '汇率（一个兑换码需要的票券数量）',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `delete_time` bigint(20) NOT NULL COMMENT '删除时间(大于零表示已删除)',
  PRIMARY KEY (`cdk_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for dtalk_cdk_list
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_cdk_list`;
CREATE TABLE `dtalk_cdk_list` (
  `id` bigint(20) NOT NULL COMMENT '记录id',
  `cdk_id` bigint(20) NOT NULL COMMENT 'cdk的id',
  `cdk_content` varchar(255) NOT NULL COMMENT 'cdk的内容',
  `user_id` varchar(255) DEFAULT NULL COMMENT '拥有用户id',
  `cdk_status` tinyint(4) NOT NULL COMMENT 'cdk的状态（0：未发放；1：冻结；2：已发放）',
  `order_id` bigint(20) DEFAULT NULL COMMENT '订单id',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `delete_time` bigint(20) NOT NULL COMMENT '删除时间',
  `exchange_time` bigint(20) NOT NULL COMMENT '兑换时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for dtalk_group_apply
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_apply`;
CREATE TABLE `dtalk_group_apply` (
  `id` bigint(20) NOT NULL COMMENT '审批 ID',
  `group_id` bigint(20) NOT NULL COMMENT '群 ID',
  `inviter_id` varchar(40) DEFAULT NULL COMMENT '邀请人 ID, 空表示是自己主动申请的',
  `member_id` varchar(40) NOT NULL COMMENT '申请加入人 ID',
  `apply_note` varchar(255) DEFAULT NULL COMMENT '申请备注',
  `operator_id` varchar(40) DEFAULT NULL COMMENT '审批人 ID',
  `apply_status` tinyint(4) NOT NULL COMMENT '0=待审批, 1=审批通过, 2=审批不通过, 10=审批忽略',
  `reject_reason` varchar(255) DEFAULT NULL COMMENT '拒绝原因',
  `create_time` bigint(20) DEFAULT NULL COMMENT '创建时间 ms',
  `update_time` bigint(20) DEFAULT NULL COMMENT '修改时间 ms',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dtalk_group_info
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_info`;
CREATE TABLE `dtalk_group_info` (
  `group_id` bigint(20) NOT NULL COMMENT '群id',
  `group_mark_id` varchar(40) NOT NULL COMMENT '群编号',
  `group_name` varchar(200) NOT NULL COMMENT '群名称',
  `group_avatar` varchar(1000) NOT NULL DEFAULT '' COMMENT '群头像 url',
  `group_member_num` int(11) NOT NULL COMMENT '群成员人数',
  `group_maximum` int(11) NOT NULL DEFAULT '200' COMMENT '群成员人数上限， 默认 200 人',
  `group_introduce` longtext NOT NULL COMMENT '群简介',
  `group_status` tinyint(4) NOT NULL COMMENT '群状态，0=正常 1=封禁 2=解散',
  `group_owner_id` varchar(40) NOT NULL COMMENT '群主 id',
  `group_create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `group_update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `group_join_type` tinyint(4) NOT NULL COMMENT '加群方式，0=无需审批（默认），1=禁止加群，群主和管理员邀请加群',
  `group_mute_type` tinyint(4) NOT NULL COMMENT '禁言， 0=全员可发言， 1=全员禁言(除群主和管理员)',
  `group_friend_type` tinyint(4) NOT NULL COMMENT '加好友限制， 0=群内可加好友，1=群内禁止加好友',
  `group_aes_key` varchar(255) DEFAULT NULL COMMENT 'aes key',
  `group_pub_name` varchar(255) DEFAULT NULL COMMENT '群公开名称',
  PRIMARY KEY (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dtalk_group_member
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_member`;
CREATE TABLE `dtalk_group_member` (
  `group_id` bigint(20) NOT NULL COMMENT '群 id',
  `group_member_id` varchar(40) NOT NULL COMMENT '用户 id',
  `group_member_name` varchar(40) NOT NULL COMMENT '用户群昵称',
  `group_member_type` tinyint(4) NOT NULL COMMENT '用户角色，2=群主，1=管理员，0=群员，3=退群',
  `group_member_join_time` bigint(20) NOT NULL COMMENT '用户加群时间',
  `group_member_update_time` bigint(20) NOT NULL COMMENT '用户更新时间',
  PRIMARY KEY (`group_id`,`group_member_id`),
  KEY `idx_userid_type` (`group_member_id`,`group_member_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dtalk_group_member_mute
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_group_member_mute`;
CREATE TABLE `dtalk_group_member_mute` (
  `group_id` bigint(20) NOT NULL COMMENT '群 id',
  `group_member_id` varchar(40) NOT NULL COMMENT '用户 id',
  `group_member_mute_time` bigint(20) NOT NULL COMMENT '用户禁言结束时间',
  `group_member_mute_update_time` bigint(20) NOT NULL COMMENT '用户上一次被禁言的时间',
  PRIMARY KEY (`group_id`,`group_member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dtalk_oss_config
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_oss_config`;
CREATE TABLE `dtalk_oss_config` (
  `app` varchar(20) NOT NULL COMMENT '应用类型',
  `oss_type` varchar(20) NOT NULL COMMENT '存储服务类型',
  `endpoint` varchar(255) DEFAULT NULL COMMENT '服务节点',
  `access_key_id` varchar(255) DEFAULT NULL,
  `access_key_secret` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `policy` varchar(255) DEFAULT NULL COMMENT '角色权限控制',
  `duration_seconds` int(11) DEFAULT NULL COMMENT '最大会话时间',
  PRIMARY KEY (`app`,`oss_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dtalk_ver_auth
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_ver_auth`;
CREATE TABLE `dtalk_ver_auth` (
  `app_id` varchar(40) NOT NULL COMMENT 'AppId',
  `app_config` text NOT NULL COMMENT '应用配置内容',
  `app_key` varchar(64) NOT NULL COMMENT 'key',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`app_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for dtalk_ver_backend
-- ----------------------------
DROP TABLE IF EXISTS `dtalk_ver_backend`;
CREATE TABLE `dtalk_ver_backend` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '版本编号',
  `platform` varchar(40) NOT NULL COMMENT '平台',
  `state` tinyint(4) NOT NULL COMMENT '线上状态',
  `device_type` varchar(40) NOT NULL COMMENT '终端',
  `version_code` bigint(20) NOT NULL COMMENT '版本号',
  `version_name` varchar(40) NOT NULL COMMENT '版本名字',
  `download_url` varchar(2083) NOT NULL COMMENT '下载地址',
  `size` bigint(20) NOT NULL COMMENT '包大小',
  `md5` varchar(40) NOT NULL COMMENT 'MD5',
  `force_update` tinyint(4) NOT NULL COMMENT '强制更新',
  `description` text NOT NULL COMMENT '描述信息',
  `ope_user` varchar(40) DEFAULT NULL COMMENT '操作者',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
