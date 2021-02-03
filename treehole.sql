/*
 Navicat MySQL Data Transfer
 
 Source Server         : 1
 Source Server Version : 50716
 Source Host           : localhost:3306
 Source Database       : treehole
 
 Target Server Type    : MYSQL
 Target Server Version : 50716
 File Encoding         : 65001
 
 Date: 2020-10-14 23:10:45
 */
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for resume
-- ----------------------------
DROP TABLE IF EXISTS `resume`;
CREATE TABLE `resume` (
    `id`         INT(10)      COMMENT  '简历信息id 每个用户对应多份简历', 
    `date`       VARCHAR(255) COMMENT  '自何年月至何年月', 
    `work_place` VARCHAR(255) COMMENT  '工作单位', 
    `job`        VARCHAR(255) COMMENT  '职务' 
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- ----------------------------
-- Records of resume
-- ----------------------------
INSERT INTO `resume`
VALUES ('16', '2009', '武汉', '教授');
INSERT INTO `resume`
VALUES ('16', '2006', '杭州', '副教授');
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `level` int(2) NOT NULL DEFAULT '0' COMMENT '用户权限等级 0--普通用户 1--普通管理员 2--超级管理员',
  `name` varchar(32) DEFAULT NULL COMMENT '姓名',
  `sex` varchar (32) DEFAULT NULL COMMENT '性别',
  `birth` varchar(50) DEFAULT NULL COMMENT '生日',
  `nation` varchar(32) DEFAULT NULL COMMENT '民族',
  `native_place` varchar(255) DEFAULT NULL COMMENT '籍贯',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `identity_number` varchar(255) DEFAULT NULL COMMENT '身份证号或护照证件号',
  `image_url` varchar(255) DEFAULT NULL COMMENT '头像',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `phone` varchar(32) NOT NULL COMMENT '注册时使用手机号码',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb4;
-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user`
VALUES (
    '1',
    '0',
    null,
    null,
    null,
    null,
    null,
    null,
    null,
    null,
    'hlyhahaha',
    '13476099548'
  );
INSERT INTO `user`
VALUES (
    '2',
    '0',
    null,
    null,
    null,
    null,
    null,
    null,
    null,
    null,
    'hlyhehehe',
    '1347609954'
  );
-- ----------------------------
-- Table structure for volunteer
-- ----------------------------
DROP TABLE IF EXISTS `volunteer`;
CREATE TABLE `volunteer` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL COMMENT '申请用户id',
  `ischeck` int(1) NOT NULL DEFAULT '1' COMMENT '志愿者申请表单进度 1--完善个人信息 2--提交申请表 3--审核申报人信息 4--志愿者协会审核 5--审核通过',
  /*`ispass` int(1) NOT NULL COMMENT '志愿者申请表单审核是否通过' */
  `reference` varchar(32) DEFAULT NULL COMMENT '推荐人',
  `name` varchar(32) NOT NULL COMMENT '志愿者姓名',
  `birth` varchar(50) NOT NULL COMMENT '出生日期',
  `political` varchar(32) NOT NULL COMMENT '政治面貌',
  `sex` varchar(32) NOT NULL COMMENT '性别',
  `nation` varchar(32) NOT NULL COMMENT '民族',
  `native_place` varchar(255) NOT NULL COMMENT '籍贯',
  `education` varchar(255) NOT NULL COMMENT '文化程度',
  `nationality` varchar(255) NOT NULL COMMENT '国籍',
  `identity_number` varchar(255) NOT NULL COMMENT '身份证号或护照证件号',
  `workphone` varchar(255) NOT NULL COMMENT '办公电话',
  `phone` varchar(255) NOT NULL COMMENT '手机电话',
  `email` varchar(255) NOT NULL COMMENT '邮箱',
  `job` varchar(255) NOT NULL COMMENT '所在部门及职务',
  `social_job` varchar(255) NOT NULL COMMENT '其他社会职务',
  `medical_history` varchar(255) NOT NULL COMMENT '病史',
  `treatment_history` varchar(255) NOT NULL COMMENT '治疗史',
  `medicine` varchar(255) NOT NULL COMMENT '现用药物',
  `reason` varchar(255) NOT NULL COMMENT '申请加入原因',
  `front_url` varchar(255) NOT NULL COMMENT '身份证正面',
  `contrary_url` varchar(255) NOT NULL COMMENT '身份证背面',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 17 DEFAULT CHARSET = utf8mb4;
-- ----------------------------
-- Records of volunteer
-- ----------------------------
INSERT INTO `volunteer`
VALUES (
    '16',
    '2',
    '2',
    '黄智生',
    '蒋兴鹏',
    '1979年7月6日',
    '中共党员',
    '男',
    '汉族',
    '山东省泰安市',
    '博士',
    '中国',
    '420102199810063714',
    '15071265377',
    '15071265377',
    'suoqi1998@mails.ccnu.edu.cn',
    '华中师范大学计算机学院,教授',
    '湖北省',
    '无',
    '无',
    '无',
    '加油',
    '\\img\\2020\\09\\24\\aae111cd-6c0b-4c64-9311-5d5daa5d2d1a.pdf',
    '\\img\\2020\\09\\24\\96fb8df3-8cea-4395-8b5c-24d4390035bc.pdf',
    '2020-09-24 12:04:39'
  );