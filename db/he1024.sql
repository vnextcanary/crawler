/*
 Navicat Premium Data Transfer

 Source Server         : 223
 Source Server Type    : MySQL
 Source Server Version : 50727
 Source Host           : 192.168.200.223:3306
 Source Schema         : 1024

 Target Server Type    : MySQL
 Target Server Version : 50727
 File Encoding         : 65001

 Date: 01/04/2020 16:12:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for he1024
-- ----------------------------
DROP TABLE IF EXISTS `he1024`;
CREATE TABLE `he1024`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `pics` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `url` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'bbs',
  `err` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `add_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  INDEX `Auto_Increment_Key`(`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 22435 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
