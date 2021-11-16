/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 100509
 Source Host           : localhost:3306
 Source Schema         : oburn

 Target Server Type    : MySQL
 Target Server Version : 100509
 File Encoding         : 65001

 Date: 26/10/2021 14:38:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for burn_config
-- ----------------------------
DROP TABLE IF EXISTS `burn_config`;
CREATE TABLE `burn_config`  (
  `id` bigint NOT NULL,
  `created_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `disc_label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `disc_passwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fs_type` int NULL DEFAULT NULL,
  `record_mode` int NULL DEFAULT NULL,
  `verify` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for disc_info
-- ----------------------------
DROP TABLE IF EXISTS `disc_info`;
CREATE TABLE `disc_info`  (
  `id` bigint NOT NULL,
  `created_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `existed` int NULL DEFAULT NULL,
  `blank` int NULL DEFAULT NULL,
  `complete` int NULL DEFAULT NULL,
  `serial_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `m_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `disc_type` int NULL DEFAULT NULL,
  `total_size` int NULL DEFAULT NULL,
  `free_size` int NULL DEFAULT NULL,
  `used_size` int NULL DEFAULT NULL,
  `track_num` int NULL DEFAULT NULL,
  `user_defined_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fs_type` int NULL DEFAULT NULL,
  `disc_lable` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `media_status` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for file_part_info
-- ----------------------------
DROP TABLE IF EXISTS `file_part_info`;
CREATE TABLE `file_part_info`  (
  `id` bigint NOT NULL,
  `created_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `task_id` int NULL DEFAULT NULL,
  `uuid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `file_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `file_name_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `file_type` int NULL DEFAULT NULL,
  `status` int NULL DEFAULT NULL,
  `crc32` int NULL DEFAULT NULL,
  `offset_start` int NULL DEFAULT NULL,
  `offset_end` int NULL DEFAULT NULL,
  `exists` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for task_info
-- ----------------------------
DROP TABLE IF EXISTS `task_info`;
CREATE TABLE `task_info`  (
  `id` bigint NOT NULL,
  `created_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_at` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `expired` int NULL DEFAULT NULL,
  `uuid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `file_type` int NULL DEFAULT NULL,
  `storage_type` int NULL DEFAULT NULL,
  `file_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `object_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `bucket_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `disc_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `off_set_start` int NULL DEFAULT NULL,
  `off_set_end` int NULL DEFAULT NULL,
  `retries` int NULL DEFAULT NULL,
  `file_size` int NULL DEFAULT NULL,
  `task_mode` int NULL DEFAULT NULL,
  `disc_mode` int NULL DEFAULT NULL,
  `parent_id` int NULL DEFAULT NULL,
  `status` int NULL DEFAULT NULL,
  `burn_progress` int NULL DEFAULT NULL,
  `verify_progress` int NULL DEFAULT NULL,
  `file_crc32` int NULL DEFAULT NULL,
  `error_code` int NULL DEFAULT NULL,
  `error_message` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `error` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
