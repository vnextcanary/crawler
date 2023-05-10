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

 Date: 01/04/2020 16:12:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pic1024
-- ----------------------------
DROP TABLE IF EXISTS `pic1024`;
CREATE TABLE `pic1024`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `isdownloadok` int(11) NULL DEFAULT 0,
  `isok` int(11) NULL DEFAULT 0,
  `add_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
