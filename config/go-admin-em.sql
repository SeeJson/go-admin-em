/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.43 : Database - go-admin-em
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`go-admin-em` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;

USE `go-admin-em`;

/*Table structure for table `em_class_info` */

DROP TABLE IF EXISTS `em_class_info`;

CREATE TABLE `em_class_info` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '班级名',
  `major` varchar(30) DEFAULT NULL COMMENT '专业',
  `college` varchar(30) DEFAULT NULL COMMENT '学院',
  `re_person` varchar(30) DEFAULT NULL COMMENT '负责人',
  `re_person_no` int(30) DEFAULT NULL COMMENT '负责人学号',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='班级信息';

/*Data for the table `em_class_info` */

insert  into `em_class_info`(`id`,`name`,`major`,`college`,`re_person`,`re_person_no`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (3,'18级通信一班','通信工程','电信分院','陈天',1111,0,0,'2024-03-10 17:54:07.139','2024-03-10 17:54:07.139',NULL);

/*Table structure for table `em_college_info` */

DROP TABLE IF EXISTS `em_college_info`;

CREATE TABLE `em_college_info` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '学院名',
  `bk` varchar(200) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '学院简介',
  `re_person` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '负责人姓名',
  `re_phone` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '负责人手机号',
  `re_no` int(20) DEFAULT NULL COMMENT '负责人工号',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='学院信息';

/*Data for the table `em_college_info` */

insert  into `em_college_info`(`id`,`name`,`bk`,`re_person`,`re_phone`,`re_no`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'电信分院','分院好的很','牛牛','18174010019',1002,3,0,'2024-03-10 21:17:23.524','2024-03-10 21:17:23.524',NULL);

/*Table structure for table `em_major_info` */

DROP TABLE IF EXISTS `em_major_info`;

CREATE TABLE `em_major_info` (
  `id` int(12) NOT NULL AUTO_INCREMENT,
  `major_name` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '所属专业',
  `major_bk` varchar(200) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '简介',
  `college_name` varchar(12) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '所属学院',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='专业信息';

/*Data for the table `em_major_info` */

insert  into `em_major_info`(`id`,`major_name`,`major_bk`,`college_name`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'通信工程','11','电信分院',3,1,'2024-03-10 21:18:56.837','2024-03-10 21:19:43.019',NULL);

/*Table structure for table `em_student` */

DROP TABLE IF EXISTS `em_student`;

CREATE TABLE `em_student` (
  `id` int(12) NOT NULL AUTO_INCREMENT,
  `sex` int(1) DEFAULT NULL COMMENT '性别',
  `name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '姓名',
  `number` int(20) DEFAULT NULL COMMENT '学号',
  `class_name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '班级名',
  `major` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '专业名',
  `college_name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '学院名',
  `phone` varchar(12) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号',
  `email` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '邮箱',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='学生信息';

/*Data for the table `em_student` */

insert  into `em_student`(`id`,`sex`,`name`,`number`,`class_name`,`major`,`college_name`,`phone`,`email`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (3,0,'学生1',1002,'通信一班','通信工程','电信分院','18174010019','1212',3,1,'2024-03-10 20:52:24.480','2024-03-10 20:58:24.638',NULL);

/*Table structure for table `em_textbook` */

DROP TABLE IF EXISTS `em_textbook`;

CREATE TABLE `em_textbook` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '教材名',
  `price` int(10) DEFAULT NULL COMMENT '价格',
  `sup` int(10) DEFAULT NULL COMMENT '教材供应商',
  `pro_data` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '出版日期',
  `author` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '出版作者',
  `press` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '出版社',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='教材信息';

/*Data for the table `em_textbook` */

insert  into `em_textbook`(`id`,`name`,`price`,`sup`,`pro_data`,`author`,`press`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'通信原理',56,1,'2024-03-19 00:00:00','奥巴马','清华出版社',1,1,'2024-03-10 19:59:19.502','2024-03-10 20:18:49.016',NULL),(2,'电路信息',65,3,'2024-03-10 20:14:08','刘婷婷','北京大学出版社',1,0,'2024-03-10 20:14:22.528','2024-03-10 20:14:22.528',NULL);

/*Table structure for table `em_textbook_in` */

DROP TABLE IF EXISTS `em_textbook_in`;

CREATE TABLE `em_textbook_in` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `textbook_id` int(10) DEFAULT NULL COMMENT '教材_id',
  `in_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '入库时间',
  `price` int(10) DEFAULT NULL COMMENT '定价',
  `num` int(10) DEFAULT NULL COMMENT '数量',
  `textbook_pro` int(10) DEFAULT NULL COMMENT '供应商',
  `textbook_unit` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '经办单位',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='教材入库信息';

/*Data for the table `em_textbook_in` */

insert  into `em_textbook_in`(`id`,`textbook_id`,`in_time`,`price`,`num`,`textbook_pro`,`textbook_unit`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,1,'2024-03-10 20:37:18',26,500,2,'清华出版社1',0,0,'2024-03-10 20:37:31.051','2024-03-10 20:37:35.683',NULL);

/*Table structure for table `em_textbook_on` */

DROP TABLE IF EXISTS `em_textbook_on`;

CREATE TABLE `em_textbook_on` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `textbook_id` int(10) DEFAULT NULL COMMENT '教材_id',
  `num` int(10) DEFAULT NULL COMMENT '数量',
  `textbook_unit` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '经办单位',
  `out_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '出库时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='教材出库信息';

/*Data for the table `em_textbook_on` */

insert  into `em_textbook_on`(`id`,`textbook_id`,`num`,`textbook_unit`,`out_time`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,1,20,'电信分院','2024-03-10 20:38:01',1,0,'2024-03-10 20:38:02.634','2024-03-10 20:38:02.634',NULL);

/*Table structure for table `em_textbook_rev` */

DROP TABLE IF EXISTS `em_textbook_rev`;

CREATE TABLE `em_textbook_rev` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `rev_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '领取书剑',
  `class_id` int(10) DEFAULT NULL COMMENT '班级_id',
  `textbook_id` int(10) DEFAULT NULL COMMENT '教材_id',
  `num` int(11) DEFAULT NULL COMMENT '数量',
  `class_leader` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '班级负责人',
  `class_leader_phone` int(12) DEFAULT NULL COMMENT '负责人电话',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='教材班级领取信息';

/*Data for the table `em_textbook_rev` */

/*Table structure for table `em_textbook_sup` */

DROP TABLE IF EXISTS `em_textbook_sup`;

CREATE TABLE `em_textbook_sup` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '供应商名',
  `phone` bigint(12) DEFAULT NULL COMMENT '供应商电话',
  `addr` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '供应商地址',
  `time` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '供应商开始供应时间',
  `status` int(1) DEFAULT NULL COMMENT '状态',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`,`time`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='供应商信息';

/*Data for the table `em_textbook_sup` */

insert  into `em_textbook_sup`(`id`,`name`,`phone`,`addr`,`time`,`status`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'1',1,'1','2024-03-10 12:19:44.193',NULL,NULL,NULL,NULL,NULL,'2024-03-10 20:19:44.193'),(2,'清华出版社',18174010019,'清华出版社','2024-03-10 20:08:14.000',2,1,1,'2024-03-10 20:08:16.810','2024-03-10 20:12:25.908',NULL),(3,'北京大学出版社',18174010019,'北京大学出版社','2024-03-10 20:12:48.000',2,1,0,'2024-03-10 20:12:57.587','2024-03-10 20:12:57.587',NULL),(4,'江西南昌大学出版社',10086,'江西南昌大学出版社','2024-03-10 20:19:53.000',1,1,1,'2024-03-10 20:19:55.910','2024-03-10 20:20:00.891',NULL);

/*Table structure for table `sys_api` */

DROP TABLE IF EXISTS `sys_api`;

CREATE TABLE `sys_api` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `handle` varchar(128) DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) DEFAULT NULL COMMENT '标题',
  `path` varchar(128) DEFAULT NULL COMMENT '地址',
  `type` varchar(16) DEFAULT NULL COMMENT '接口类型',
  `action` varchar(16) DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`),
  KEY `idx_sys_api_create_by` (`create_by`),
  KEY `idx_sys_api_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=210 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_api` */

insert  into `sys_api`(`id`,`handle`,`title`,`path`,`type`,`action`,`created_at`,`updated_at`,`deleted_at`,`create_by`,`update_by`) values (5,'go-admin/app/admin/apis.SysLoginLog.Get-fm','登录日志通过id获取','/api/v1/sys-login-log/:id','BUS','GET','2021-05-13 19:59:00.728','2021-06-17 11:37:12.331',NULL,0,0),(6,'go-admin/app/admin/apis.SysOperaLog.GetPage-fm','操作日志列表','/api/v1/sys-opera-log','BUS','GET','2021-05-13 19:59:00.778','2021-06-17 11:48:40.732',NULL,0,0),(7,'go-admin/app/admin/apis.SysOperaLog.Get-fm','操作日志通过id获取','/api/v1/sys-opera-log/:id','BUS','GET','2021-05-13 19:59:00.821','2021-06-16 21:49:48.598',NULL,0,0),(8,'go-admin/common/actions.IndexAction.func1','分类列表','/api/v1/syscategory','BUS','GET','2021-05-13 19:59:00.870','2021-06-13 20:53:47.883',NULL,0,0),(9,'go-admin/common/actions.ViewAction.func1','分类通过id获取','/api/v1/syscategory/:id','BUS','GET','2021-05-13 19:59:00.945','2021-06-13 20:53:47.926',NULL,0,0),(10,'go-admin/common/actions.IndexAction.func1','内容列表','/api/v1/syscontent','BUS','GET','2021-05-13 19:59:01.005','2021-06-13 20:53:47.966',NULL,0,0),(11,'go-admin/common/actions.ViewAction.func1','内容通过id获取','/api/v1/syscontent/:id','BUS','GET','2021-05-13 19:59:01.056','2021-06-13 20:53:48.005',NULL,0,0),(15,'go-admin/common/actions.IndexAction.func1','job列表','/api/v1/sysjob','BUS','GET','2021-05-13 19:59:01.248','2021-06-13 20:53:48.169',NULL,0,0),(16,'go-admin/common/actions.ViewAction.func1','job通过id获取','/api/v1/sysjob/:id','BUS','GET','2021-05-13 19:59:01.298','2021-06-13 20:53:48.214',NULL,0,0),(21,'go-admin/app/admin/apis.SysDictType.GetPage-fm','字典类型列表','/api/v1/dict/type','BUS','GET','2021-05-13 19:59:01.525','2021-06-17 11:48:40.732',NULL,0,0),(22,'go-admin/app/admin/apis.SysDictType.GetAll-fm','字典类型查询【代码生成】','/api/v1/dict/type-option-select','SYS','GET','2021-05-13 19:59:01.582','2021-06-13 20:53:48.388',NULL,0,0),(23,'go-admin/app/admin/apis.SysDictType.Get-fm','字典类型通过id获取','/api/v1/dict/type/:id','BUS','GET','2021-05-13 19:59:01.632','2021-06-17 11:48:40.732',NULL,0,0),(24,'go-admin/app/admin/apis.SysDictData.GetPage-fm','字典数据列表','/api/v1/dict/data','BUS','GET','2021-05-13 19:59:01.684','2021-06-17 11:48:40.732',NULL,0,0),(25,'go-admin/app/admin/apis.SysDictData.Get-fm','字典数据通过code获取','/api/v1/dict/data/:dictCode','BUS','GET','2021-05-13 19:59:01.732','2021-06-17 11:48:40.732',NULL,0,0),(26,'go-admin/app/admin/apis.SysDictData.GetSysDictDataAll-fm','数据字典根据key获取','/api/v1/dict-data/option-select','SYS','GET','2021-05-13 19:59:01.832','2021-06-17 11:48:40.732',NULL,0,0),(27,'go-admin/app/admin/apis.SysDept.GetPage-fm','部门列表','/api/v1/dept','BUS','GET','2021-05-13 19:59:01.940','2021-06-17 11:48:40.732',NULL,0,0),(28,'go-admin/app/admin/apis.SysDept.Get-fm','部门通过id获取','/api/v1/dept/:id','BUS','GET','2021-05-13 19:59:02.009','2021-06-17 11:48:40.732',NULL,0,0),(29,'go-admin/app/admin/apis.SysDept.Get2Tree-fm','查询部门下拉树【角色权限-自定权限】','/api/v1/deptTree','SYS','GET','2021-05-13 19:59:02.050','2021-06-17 11:48:40.732',NULL,0,0),(30,'go-admin/app/admin/apis/tools.(*Gen).GetDBTableList-fm','数据库表列表','/api/v1/db/tables/page','SYS','GET','2021-05-13 19:59:02.098','2021-06-13 20:53:48.730',NULL,0,0),(31,'go-admin/app/admin/apis/tools.(*Gen).GetDBColumnList-fm','数据表列列表','/api/v1/db/columns/page','SYS','GET','2021-05-13 19:59:02.140','2021-06-13 20:53:48.771',NULL,0,0),(32,'go-admin/app/admin/apis/tools.Gen.GenCode-fm','数据库表生成到项目','/api/v1/gen/toproject/:tableId','SYS','GET','2021-05-13 19:59:02.183','2021-06-13 20:53:48.812',NULL,0,0),(33,'go-admin/app/admin/apis/tools.Gen.GenMenuAndApi-fm','数据库表生成到DB','/api/v1/gen/todb/:tableId','SYS','GET','2021-05-13 19:59:02.227','2021-06-13 20:53:48.853',NULL,0,0),(34,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesTree-fm','关系表数据【代码生成】','/api/v1/gen/tabletree','SYS','GET','2021-05-13 19:59:02.271','2021-06-13 20:53:48.893',NULL,0,0),(35,'go-admin/app/admin/apis/tools.Gen.Preview-fm','生成预览通过id获取','/api/v1/gen/preview/:tableId','SYS','GET','2021-05-13 19:59:02.315','2021-06-13 20:53:48.935',NULL,0,0),(36,'go-admin/app/admin/apis/tools.Gen.GenApiToFile-fm','生成api带文件','/api/v1/gen/apitofile/:tableId','SYS','GET','2021-05-13 19:59:02.357','2021-06-13 20:53:48.977',NULL,0,0),(37,'go-admin/app/admin/apis.System.GenerateCaptchaHandler-fm','验证码获取','/api/v1/getCaptcha','SYS','GET','2021-05-13 19:59:02.405','2021-06-13 20:53:49.020',NULL,0,0),(38,'go-admin/app/admin/apis.SysUser.GetInfo-fm','用户信息获取','/api/v1/getinfo','SYS','GET','2021-05-13 19:59:02.447','2021-06-13 20:53:49.065',NULL,0,0),(39,'go-admin/app/admin/apis.SysMenu.GetPage-fm','菜单列表','/api/v1/menu','BUS','GET','2021-05-13 19:59:02.497','2021-06-17 11:48:40.732',NULL,0,0),(40,'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm','查询菜单下拉树结构【废弃】','/api/v1/menuTreeselect','SYS','GET','2021-05-13 19:59:02.542','2021-06-03 22:37:21.857',NULL,0,0),(41,'go-admin/app/admin/apis.SysMenu.Get-fm','菜单通过id获取','/api/v1/menu/:id','BUS','GET','2021-05-13 19:59:02.584','2021-06-17 11:48:40.732',NULL,0,0),(42,'go-admin/app/admin/apis.SysMenu.GetMenuRole-fm','角色菜单【顶部左侧菜单】','/api/v1/menurole','SYS','GET','2021-05-13 19:59:02.630','2021-06-13 20:53:49.574',NULL,0,0),(43,'go-admin/app/admin/apis.SysMenu.GetMenuIDS-fm','获取角色对应的菜单id数组【废弃】','/api/v1/menuids','SYS','GET','2021-05-13 19:59:02.675','2021-06-03 22:39:52.500',NULL,0,0),(44,'go-admin/app/admin/apis.SysRole.GetPage-fm','角色列表','/api/v1/role','BUS','GET','2021-05-13 19:59:02.720','2021-06-17 11:48:40.732',NULL,0,0),(45,'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm','菜单权限列表【角色配菜单使用】','/api/v1/roleMenuTreeselect/:roleId','SYS','GET','2021-05-13 19:59:02.762','2021-06-17 11:48:40.732',NULL,0,0),(46,'go-admin/app/admin/apis.SysDept.GetDeptTreeRoleSelect-fm','角色部门结构树【自定义数据权限】','/api/v1/roleDeptTreeselect/:roleId','SYS','GET','2021-05-13 19:59:02.809','2021-06-17 11:48:40.732',NULL,0,0),(47,'go-admin/app/admin/apis.SysRole.Get-fm','角色通过id获取','/api/v1/role/:id','BUS','GET','2021-05-13 19:59:02.850','2021-06-17 11:48:40.732',NULL,0,0),(48,'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).RefreshHandler-fm','刷新token','/api/v1/refresh_token','SYS','GET','2021-05-13 19:59:02.892','2021-06-13 20:53:49.278',NULL,0,0),(53,'go-admin/app/admin/apis.SysConfig.GetPage-fm','参数列表','/api/v1/config','BUS','GET','2021-05-13 19:59:03.116','2021-06-17 11:48:40.732',NULL,0,0),(54,'go-admin/app/admin/apis.SysConfig.Get-fm','参数通过id获取','/api/v1/config/:id','BUS','GET','2021-05-13 19:59:03.157','2021-06-17 11:48:40.732',NULL,0,0),(55,'go-admin/app/admin/apis.SysConfig.GetSysConfigByKEYForService-fm','参数通过键名搜索【基础默认配置】','/api/v1/configKey/:configKey','SYS','GET','2021-05-13 19:59:03.198','2021-06-13 20:53:49.745',NULL,0,0),(57,'go-admin/app/jobs/apis.SysJob.RemoveJobForService-fm','job移除','/api/v1/job/remove/:id','BUS','GET','2021-05-13 19:59:03.295','2021-06-13 20:53:49.786',NULL,0,0),(58,'go-admin/app/jobs/apis.SysJob.StartJobForService-fm','job启动','/api/v1/job/start/:id','BUS','GET','2021-05-13 19:59:03.339','2021-06-13 20:53:49.829',NULL,0,0),(59,'go-admin/app/admin/apis.SysPost.GetPage-fm','岗位列表','/api/v1/post','BUS','GET','2021-05-13 19:59:03.381','2021-06-17 11:48:40.732',NULL,0,0),(60,'go-admin/app/admin/apis.SysPost.Get-fm','岗位通过id获取','/api/v1/post/:id','BUS','GET','2021-05-13 19:59:03.433','2021-06-17 11:48:40.732',NULL,0,0),(62,'go-admin/app/admin/apis.SysConfig.GetSysConfigBySysApp-fm','系统前端参数','/api/v1/app-config','SYS','GET','2021-05-13 19:59:03.526','2021-06-13 20:53:49.994',NULL,0,0),(63,'go-admin/app/admin/apis.SysUser.GetProfile-fm','*用户信息获取','/api/v1/user/profile','SYS','GET','2021-05-13 19:59:03.567','2021-06-13 20:53:50.038',NULL,0,0),(66,'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).WsClient-fm','链接ws【定时任务log】','/ws/:id/:channel','BUS','GET','2021-05-13 19:59:03.705','2021-06-13 20:53:50.167',NULL,0,0),(67,'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).UnWsClient-fm','退出ws【定时任务log】','/wslogout/:id/:channel','BUS','GET','2021-05-13 19:59:03.756','2021-06-13 20:53:50.209',NULL,0,0),(68,'go-admin/common/middleware/handler.Ping','*用户基本信息','/info','SYS','GET','2021-05-13 19:59:03.800','2021-06-13 20:53:50.251',NULL,0,0),(72,'go-admin/common/actions.CreateAction.func1','分类创建','/api/v1/syscategory','BUS','POST','2021-05-13 19:59:03.982','2021-06-13 20:53:50.336',NULL,0,0),(73,'go-admin/common/actions.CreateAction.func1','内容创建','/api/v1/syscontent','BUS','POST','2021-05-13 19:59:04.027','2021-06-13 20:53:50.375',NULL,0,0),(76,'go-admin/common/actions.CreateAction.func1','job创建','/api/v1/sysjob','BUS','POST','2021-05-13 19:59:04.164','2021-06-13 20:53:50.500',NULL,0,0),(80,'go-admin/app/admin/apis.SysDictData.Insert-fm','字典数据创建','/api/v1/dict/data','BUS','POST','2021-05-13 19:59:04.347','2021-06-17 11:48:40.732',NULL,0,0),(81,'go-admin/app/admin/apis.SysDictType.Insert-fm','字典类型创建','/api/v1/dict/type','BUS','POST','2021-05-13 19:59:04.391','2021-06-17 11:48:40.732',NULL,0,0),(82,'go-admin/app/admin/apis.SysDept.Insert-fm','部门创建','/api/v1/dept','BUS','POST','2021-05-13 19:59:04.435','2021-06-17 11:48:40.732',NULL,0,0),(85,'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).LoginHandler-fm','*登录','/api/v1/login','SYS','POST','2021-05-13 19:59:04.597','2021-06-13 20:53:50.784',NULL,0,0),(86,'go-admin/common/middleware/handler.LogOut','*退出','/api/v1/logout','SYS','POST','2021-05-13 19:59:04.642','2021-06-13 20:53:50.824',NULL,0,0),(87,'go-admin/app/admin/apis.SysConfig.Insert-fm','参数创建','/api/v1/config','BUS','POST','2021-05-13 19:59:04.685','2021-06-17 11:48:40.732',NULL,0,0),(88,'go-admin/app/admin/apis.SysMenu.Insert-fm','菜单创建','/api/v1/menu','BUS','POST','2021-05-13 19:59:04.777','2021-06-17 11:48:40.732',NULL,0,0),(89,'go-admin/app/admin/apis.SysPost.Insert-fm','岗位创建','/api/v1/post','BUS','POST','2021-05-13 19:59:04.886','2021-06-17 11:48:40.732',NULL,0,0),(90,'go-admin/app/admin/apis.SysRole.Insert-fm','角色创建','/api/v1/role','BUS','POST','2021-05-13 19:59:04.975','2021-06-17 11:48:40.732',NULL,0,0),(91,'go-admin/app/admin/apis.SysUser.InsetAvatar-fm','*用户头像编辑','/api/v1/user/avatar','SYS','POST','2021-05-13 19:59:05.058','2021-06-13 20:53:51.079',NULL,0,0),(92,'go-admin/app/admin/apis.SysApi.Update-fm','接口编辑','/api/v1/sys-api/:id','BUS','PUT','2021-05-13 19:59:05.122','2021-06-17 11:48:40.732',NULL,0,0),(95,'go-admin/common/actions.UpdateAction.func1','分类编辑','/api/v1/syscategory/:id','BUS','PUT','2021-05-13 19:59:05.255','2021-06-13 20:53:51.247',NULL,0,0),(96,'go-admin/common/actions.UpdateAction.func1','内容编辑','/api/v1/syscontent/:id','BUS','PUT','2021-05-13 19:59:05.299','2021-06-13 20:53:51.289',NULL,0,0),(97,'go-admin/common/actions.UpdateAction.func1','job编辑','/api/v1/sysjob','BUS','PUT','2021-05-13 19:59:05.343','2021-06-13 20:53:51.331',NULL,0,0),(101,'go-admin/app/admin/apis.SysDictData.Update-fm','字典数据编辑','/api/v1/dict/data/:dictCode','BUS','PUT','2021-05-13 19:59:05.519','2021-06-17 11:48:40.732',NULL,0,0),(102,'go-admin/app/admin/apis.SysDictType.Update-fm','字典类型编辑','/api/v1/dict/type/:id','BUS','PUT','2021-05-13 19:59:05.569','2021-06-17 11:48:40.732',NULL,0,0),(103,'go-admin/app/admin/apis.SysDept.Update-fm','部门编辑','/api/v1/dept/:id','BUS','PUT','2021-05-13 19:59:05.613','2021-06-17 11:48:40.732',NULL,0,0),(104,'go-admin/app/other/apis.SysFileDir.Update-fm','文件夹编辑','/api/v1/file-dir/:id','BUS','PUT','2021-05-13 19:59:05.662','2021-06-13 20:53:51.847',NULL,0,0),(105,'go-admin/app/other/apis.SysFileInfo.Update-fm','文件编辑','/api/v1/file-info/:id','BUS','PUT','2021-05-13 19:59:05.709','2021-06-13 20:53:51.892',NULL,0,0),(106,'go-admin/app/admin/apis.SysRole.Update-fm','角色编辑','/api/v1/role/:id','BUS','PUT','2021-05-13 19:59:05.752','2021-06-17 11:48:40.732',NULL,0,0),(107,'go-admin/app/admin/apis.SysRole.Update2DataScope-fm','角色数据权限修改','/api/v1/roledatascope','BUS','PUT','2021-05-13 19:59:05.803','2021-06-17 11:48:40.732',NULL,0,0),(108,'go-admin/app/admin/apis.SysConfig.Update-fm','参数编辑','/api/v1/config/:id','BUS','PUT','2021-05-13 19:59:05.848','2021-06-17 11:48:40.732',NULL,0,0),(109,'go-admin/app/admin/apis.SysMenu.Update-fm','编辑菜单','/api/v1/menu/:id','BUS','PUT','2021-05-13 19:59:05.891','2021-06-17 11:48:40.732',NULL,0,0),(110,'go-admin/app/admin/apis.SysPost.Update-fm','岗位编辑','/api/v1/post/:id','BUS','PUT','2021-05-13 19:59:05.934','2021-06-17 11:48:40.732',NULL,0,0),(111,'go-admin/app/admin/apis.SysUser.UpdatePwd-fm','*用户修改密码','/api/v1/user/pwd','SYS','PUT','2021-05-13 19:59:05.987','2021-06-13 20:53:51.724',NULL,0,0),(112,'go-admin/common/actions.DeleteAction.func1','分类删除','/api/v1/syscategory','BUS','DELETE','2021-05-13 19:59:06.030','2021-06-13 20:53:52.237',NULL,0,0),(113,'go-admin/common/actions.DeleteAction.func1','内容删除','/api/v1/syscontent','BUS','DELETE','2021-05-13 19:59:06.076','2021-06-13 20:53:52.278',NULL,0,0),(114,'go-admin/app/admin/apis.SysLoginLog.Delete-fm','登录日志删除','/api/v1/sys-login-log','BUS','DELETE','2021-05-13 19:59:06.118','2021-06-17 11:48:40.732',NULL,0,0),(115,'go-admin/app/admin/apis.SysOperaLog.Delete-fm','操作日志删除','/api/v1/sys-opera-log','BUS','DELETE','2021-05-13 19:59:06.162','2021-06-17 11:48:40.732',NULL,0,0),(116,'go-admin/common/actions.DeleteAction.func1','job删除','/api/v1/sysjob','BUS','DELETE','2021-05-13 19:59:06.206','2021-06-13 20:53:52.323',NULL,0,0),(117,'go-admin/app/other/apis.SysChinaAreaData.Delete-fm','行政区删除','/api/v1/sys-area-data','BUS','DELETE','2021-05-13 19:59:06.249','2021-06-13 20:53:52.061',NULL,0,0),(120,'go-admin/app/admin/apis.SysDictData.Delete-fm','字典数据删除','/api/v1/dict/data','BUS','DELETE','2021-05-13 19:59:06.387','2021-06-17 11:48:40.732',NULL,0,0),(121,'go-admin/app/admin/apis.SysDictType.Delete-fm','字典类型删除','/api/v1/dict/type','BUS','DELETE','2021-05-13 19:59:06.432','2021-06-17 11:48:40.732',NULL,0,0),(122,'go-admin/app/admin/apis.SysDept.Delete-fm','部门删除','/api/v1/dept/:id','BUS','DELETE','2021-05-13 19:59:06.475','2021-06-17 11:48:40.732',NULL,0,0),(123,'go-admin/app/other/apis.SysFileDir.Delete-fm','文件夹删除','/api/v1/file-dir/:id','BUS','DELETE','2021-05-13 19:59:06.520','2021-06-13 20:53:52.539',NULL,0,0),(124,'go-admin/app/other/apis.SysFileInfo.Delete-fm','文件删除','/api/v1/file-info/:id','BUS','DELETE','2021-05-13 19:59:06.566','2021-06-13 20:53:52.580',NULL,0,0),(125,'go-admin/app/admin/apis.SysConfig.Delete-fm','参数删除','/api/v1/config','BUS','DELETE','2021-05-13 19:59:06.612','2021-06-17 11:48:40.732',NULL,0,0),(126,'go-admin/app/admin/apis.SysMenu.Delete-fm','删除菜单','/api/v1/menu','BUS','DELETE','2021-05-13 19:59:06.654','2021-06-17 11:48:40.732',NULL,0,0),(127,'go-admin/app/admin/apis.SysPost.Delete-fm','岗位删除','/api/v1/post/:id','BUS','DELETE','2021-05-13 19:59:06.702','2021-06-17 11:48:40.732',NULL,0,0),(128,'go-admin/app/admin/apis.SysRole.Delete-fm','角色删除','/api/v1/role','BUS','DELETE','2021-05-13 19:59:06.746','2021-06-17 11:48:40.732',NULL,0,0),(131,'github.com/go-admin-team/go-admin-core/tools/transfer.Handler.func1','系统指标','/api/v1/metrics','SYS','GET','2021-05-17 22:13:55.933','2021-06-13 20:53:49.614',NULL,0,0),(132,'go-admin/app/other/router.registerMonitorRouter.func1','健康状态','/api/v1/health','SYS','GET','2021-05-17 22:13:56.285','2021-06-13 20:53:49.951',NULL,0,0),(133,'go-admin/app/admin/apis.HelloWorld','项目默认接口','/','SYS','GET','2021-05-24 10:30:44.553','2021-06-13 20:53:47.406',NULL,0,0),(134,'go-admin/app/other/apis.ServerMonitor.ServerInfo-fm','服务器基本状态','/api/v1/server-monitor','SYS','GET','2021-05-24 10:30:44.937','2021-06-13 20:53:48.255',NULL,0,0),(135,'go-admin/app/admin/apis.SysApi.GetPage-fm','接口列表','/api/v1/sys-api','BUS','GET','2021-05-24 11:37:53.303','2021-06-17 11:48:40.732',NULL,0,0),(136,'go-admin/app/admin/apis.SysApi.Get-fm','接口通过id获取','/api/v1/sys-api/:id','BUS','GET','2021-05-24 11:37:53.359','2021-06-17 11:48:40.732',NULL,0,0),(137,'go-admin/app/admin/apis.SysLoginLog.GetPage-fm','登录日志列表','/api/v1/sys-login-log','BUS','GET','2021-05-24 11:47:30.397','2021-06-17 11:48:40.732',NULL,0,0),(138,'go-admin/app/other/apis.File.UploadFile-fm','文件上传','/api/v1/public/uploadFile','SYS','POST','2021-05-25 19:16:18.493','2021-06-13 20:53:50.866',NULL,0,0),(139,'go-admin/app/admin/apis.SysConfig.Update2Set-fm','参数信息修改【参数配置】','/api/v1/set-config','BUS','PUT','2021-05-27 09:45:14.853','2021-06-17 11:48:40.732',NULL,0,0),(140,'go-admin/app/admin/apis.SysConfig.Get2Set-fm','参数获取全部【配置使用】','/api/v1/set-config','BUS','GET','2021-05-27 11:54:14.384','2021-06-17 11:48:40.732',NULL,0,0),(141,'go-admin/app/admin/apis.SysUser.GetPage-fm','管理员列表','/api/v1/sys-user','BUS','GET','2021-06-13 19:24:57.111','2021-06-17 20:31:14.318',NULL,0,0),(142,'go-admin/app/admin/apis.SysUser.Get-fm','管理员通过id获取','/api/v1/sys-user/:id','BUS','GET','2021-06-13 19:24:57.188','2021-06-17 20:31:14.318',NULL,0,0),(143,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesInfo-fm','','/api/v1/sys/tables/info','','GET','2021-06-13 19:24:57.437','2021-06-13 20:53:48.047',NULL,0,0),(144,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTables-fm','','/api/v1/sys/tables/info/:tableId','','GET','2021-06-13 19:24:57.510','2021-06-13 20:53:48.088',NULL,0,0),(145,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTableList-fm','','/api/v1/sys/tables/page','','GET','2021-06-13 19:24:57.582','2021-06-13 20:53:48.128',NULL,0,0),(146,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/static/*filepath','','GET','2021-06-13 19:24:59.641','2021-06-13 20:53:50.081',NULL,0,0),(147,'github.com/swaggo/gin-swagger.CustomWrapHandler.func1','','/swagger/*any','','GET','2021-06-13 19:24:59.713','2021-06-13 20:53:50.123',NULL,0,0),(148,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/form-generator/*filepath','','GET','2021-06-13 19:24:59.914','2021-06-13 20:53:50.295',NULL,0,0),(149,'go-admin/app/admin/apis/tools.(*SysTable).InsertSysTable-fm','','/api/v1/sys/tables/info','','POST','2021-06-13 19:25:00.163','2021-06-13 20:53:50.539',NULL,0,0),(150,'go-admin/app/admin/apis.SysUser.Insert-fm','管理员创建','/api/v1/sys-user','BUS','POST','2021-06-13 19:25:00.233','2021-06-17 20:31:14.318',NULL,0,0),(151,'go-admin/app/admin/apis.SysUser.Update-fm','管理员编辑','/api/v1/sys-user','BUS','PUT','2021-06-13 19:25:00.986','2021-06-17 20:31:14.318',NULL,0,0),(152,'go-admin/app/admin/apis/tools.(*SysTable).UpdateSysTable-fm','','/api/v1/sys/tables/info','','PUT','2021-06-13 19:25:01.149','2021-06-13 20:53:51.375',NULL,0,0),(153,'go-admin/app/admin/apis.SysRole.Update2Status-fm','','/api/v1/role-status','','PUT','2021-06-13 19:25:01.446','2021-06-13 20:53:51.636',NULL,0,0),(154,'go-admin/app/admin/apis.SysUser.ResetPwd-fm','','/api/v1/user/pwd/reset','','PUT','2021-06-13 19:25:01.601','2021-06-13 20:53:51.764',NULL,0,0),(155,'go-admin/app/admin/apis.SysUser.UpdateStatus-fm','','/api/v1/user/status','','PUT','2021-06-13 19:25:01.671','2021-06-13 20:53:51.806',NULL,0,0),(156,'go-admin/app/admin/apis.SysUser.Delete-fm','管理员删除','/api/v1/sys-user','BUS','DELETE','2021-06-13 19:25:02.043','2021-06-17 20:31:14.318',NULL,0,0),(157,'go-admin/app/admin/apis/tools.(*SysTable).DeleteSysTables-fm','','/api/v1/sys/tables/info/:tableId','','DELETE','2021-06-13 19:25:02.283','2021-06-13 20:53:52.367',NULL,0,0),(158,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/static/*filepath','','HEAD','2021-06-13 19:25:02.734','2021-06-13 20:53:52.791',NULL,0,0),(159,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/form-generator/*filepath','','HEAD','2021-06-13 19:25:02.808','2021-06-13 20:53:52.838',NULL,0,0),(160,'go-admin/app/admin/apis.EmClassInfo.GetPage-fm','班级信息列表','/api/v1/em-class-info','BUS','GET','2024-03-10 17:40:08.833','2024-03-10 20:54:37.248',NULL,0,0),(161,'go-admin/app/admin/apis.EmClassInfo.Get-fm','班级信息详情','/api/v1/em-class-info/:id','BUS','GET','2024-03-10 17:40:08.837','2024-03-10 17:42:42.539',NULL,0,0),(162,'go-admin/app/admin/apis.EmStudent.GetPage-fm','学生信息-列表','/api/v1/em-student','BUS','GET','2024-03-10 17:40:08.842','2024-03-10 20:54:37.248',NULL,0,0),(163,'go-admin/app/admin/apis.EmStudent.Get-fm','学生信息--详情','/api/v1/em-student/:id','BUS','GET','2024-03-10 17:40:08.846','2024-03-10 17:45:34.722',NULL,0,0),(164,'go-admin/app/admin/apis.System.GenerateCaptchaHandler-fm','','/api/v1/captcha','','GET','2024-03-10 17:40:08.855','2024-03-10 17:40:08.855',NULL,0,0),(165,'github.com/swaggo/gin-swagger.CustomWrapHandler.func1','','/swagger/admin/*any','','GET','2024-03-10 17:40:08.874','2024-03-10 17:40:08.874',NULL,0,0),(166,'go-admin/app/admin/apis.EmClassInfo.Insert-fm','班级信息-新增','/api/v1/em-class-info','BUS','POST','2024-03-10 17:40:08.894','2024-03-10 20:54:37.248',NULL,0,0),(167,'go-admin/app/admin/apis.EmStudent.Insert-fm','学生信息--新增','/api/v1/em-student','BUS','POST','2024-03-10 17:40:08.898','2024-03-10 20:54:37.248',NULL,0,0),(168,'go-admin/app/admin/apis.SysUser.UpdatePwd-fm','','/api/v1/user/pwd/set','','PUT','2024-03-10 17:40:08.924','2024-03-10 17:40:08.924',NULL,0,0),(169,'go-admin/app/admin/apis.EmClassInfo.Update-fm','班级信息-修改','/api/v1/em-class-info/:id','BUS','PUT','2024-03-10 17:40:08.931','2024-03-10 20:54:37.248',NULL,0,0),(170,'go-admin/app/admin/apis.EmStudent.Update-fm','学生信息---修改1','/api/v1/em-student/:id','BUS','PUT','2024-03-10 17:40:08.935','2024-03-10 20:54:37.248',NULL,0,0),(171,'go-admin/app/admin/apis.SysDept.Delete-fm','','/api/v1/dept','','DELETE','2024-03-10 17:40:08.952','2024-03-10 17:40:08.952',NULL,0,0),(172,'go-admin/app/admin/apis.EmClassInfo.Delete-fm','班级信息删除','/api/v1/em-class-info','BUS','DELETE','2024-03-10 17:40:08.956','2024-03-10 20:54:37.248',NULL,0,0),(173,'go-admin/app/admin/apis.EmStudent.Delete-fm','学生信息--删除','/api/v1/em-student','BUS','DELETE','2024-03-10 17:40:08.960','2024-03-10 20:54:37.248',NULL,0,0),(174,'go-admin/app/admin/apis.SysPost.Delete-fm','','/api/v1/post','','DELETE','2024-03-10 17:40:08.967','2024-03-10 17:40:08.967',NULL,0,0),(175,'go-admin/app/admin/apis.EmCollegeInfo.GetPage-fm','学院信息--列表','/api/v1/em-college-info','BUS','GET','2024-03-10 18:01:24.269','2024-03-10 20:54:37.248',NULL,0,0),(176,'go-admin/app/admin/apis.EmCollegeInfo.Get-fm','学院信息--详情','/api/v1/em-college-info/:id','BUS','GET','2024-03-10 18:01:24.273','2024-03-10 18:03:20.484',NULL,0,0),(177,'go-admin/app/admin/apis.EmCollegeInfo.Insert-fm','学院信息--新增','/api/v1/em-college-info','BUS','POST','2024-03-10 18:01:24.323','2024-03-10 20:54:37.248',NULL,0,0),(178,'go-admin/app/admin/apis.EmCollegeInfo.Update-fm','学院信息--修改','/api/v1/em-college-info/:id','BUS','PUT','2024-03-10 18:01:24.354','2024-03-10 20:54:37.248',NULL,0,0),(179,'go-admin/app/admin/apis.EmCollegeInfo.Delete-fm','学院信息--删除','/api/v1/em-college-info','BUS','DELETE','2024-03-10 18:01:24.384','2024-03-10 20:54:37.248',NULL,0,0),(180,'go-admin/app/admin/apis.EmMajorInfo.GetPage-fm','专业信息--列表','/api/v1/em-major-info','BUS','GET','2024-03-10 18:02:28.245','2024-03-10 20:54:37.248',NULL,0,0),(181,'go-admin/app/admin/apis.EmMajorInfo.Get-fm','专业信息--详情','/api/v1/em-major-info/:id','BUS','GET','2024-03-10 18:02:28.250','2024-03-10 18:04:19.559',NULL,0,0),(182,'go-admin/app/admin/apis.EmMajorInfo.Insert-fm','专业信息--新增','/api/v1/em-major-info','BUS','POST','2024-03-10 18:02:28.301','2024-03-10 20:54:37.248',NULL,0,0),(183,'go-admin/app/admin/apis.EmMajorInfo.Update-fm','专业信息--修改','/api/v1/em-major-info/:id','BUS','PUT','2024-03-10 18:02:28.335','2024-03-10 20:54:37.248',NULL,0,0),(184,'go-admin/app/admin/apis.EmMajorInfo.Delete-fm','专业信息--删除','/api/v1/em-major-info','BUS','DELETE','2024-03-10 18:02:28.366','2024-03-10 20:54:37.248',NULL,0,0),(185,'go-admin/app/admin/apis.EmTextbook.GetPage-fm','教材信息--列表','/api/v1/em-textbook','BUS','GET','2024-03-10 19:55:05.026','2024-03-10 20:54:37.248',NULL,0,0),(186,'go-admin/app/admin/apis.EmTextbook.Get-fm','教材信息--详情','/api/v1/em-textbook/:id','BUS','GET','2024-03-10 19:55:05.043','2024-03-10 19:55:52.174',NULL,0,0),(187,'go-admin/app/admin/apis.EmTextbook.Insert-fm','教材信息---新增','/api/v1/em-textbook','BUS','POST','2024-03-10 19:55:05.093','2024-03-10 20:54:37.248',NULL,0,0),(188,'go-admin/app/admin/apis.EmTextbook.Update-fm','教材信息--修改','/api/v1/em-textbook/:id','BUS','PUT','2024-03-10 19:55:05.120','2024-03-10 20:54:37.248',NULL,0,0),(189,'go-admin/app/admin/apis.EmTextbook.Delete-fm','教材信息--删除','/api/v1/em-textbook','BUS','DELETE','2024-03-10 19:55:05.152','2024-03-10 20:54:37.248',NULL,0,0),(190,'go-admin/app/admin/apis.EmTextbookSup.GetPage-fm','供应商--列表','/api/v1/em-textbook-sup','BUS','GET','2024-03-10 20:01:58.415','2024-03-10 20:54:37.248',NULL,0,0),(191,'go-admin/app/admin/apis.EmTextbookSup.Get-fm','供应商--详情','/api/v1/em-textbook-sup/:id','BUS','GET','2024-03-10 20:01:58.421','2024-03-10 20:02:49.322',NULL,0,0),(192,'go-admin/app/admin/apis.EmTextbookSup.Insert-fm','供应商--新增','/api/v1/em-textbook-sup','BUS','POST','2024-03-10 20:01:58.493','2024-03-10 20:54:37.248',NULL,0,0),(193,'go-admin/app/admin/apis.EmTextbookSup.Update-fm','供应商--修改','/api/v1/em-textbook-sup/:id','BUS','PUT','2024-03-10 20:01:58.524','2024-03-10 20:54:37.248',NULL,0,0),(194,'go-admin/app/admin/apis.EmTextbookSup.Delete-fm','供应商--删除','/api/v1/em-textbook-sup','BUS','DELETE','2024-03-10 20:01:58.559','2024-03-10 20:54:37.248',NULL,0,0),(195,'go-admin/app/admin/apis.EmTextbookIn.GetPage-fm','教学信息入库--列表','/api/v1/em-textbook-in','BUS','GET','2024-03-10 20:29:27.983','2024-03-10 20:54:37.248',NULL,0,0),(196,'go-admin/app/admin/apis.EmTextbookIn.Get-fm','教学信息入库--详情','/api/v1/em-textbook-in/:id','BUS','GET','2024-03-10 20:29:27.988','2024-03-10 20:33:26.268',NULL,0,0),(197,'go-admin/app/admin/apis.EmTextbookOn.GetPage-fm','教学信息出库--列表','/api/v1/em-textbook-on','BUS','GET','2024-03-10 20:29:27.996','2024-03-10 20:54:37.248',NULL,0,0),(198,'go-admin/app/admin/apis.EmTextbookOn.Get-fm','教学信息出库--详情','/api/v1/em-textbook-on/:id','BUS','GET','2024-03-10 20:29:28.001','2024-03-10 20:34:23.132',NULL,0,0),(199,'go-admin/app/admin/apis.EmTextbookIn.Insert-fm','教学信息入库--新增','/api/v1/em-textbook-in','BUS','POST','2024-03-10 20:29:28.102','2024-03-10 20:33:39.514',NULL,0,0),(200,'go-admin/app/admin/apis.EmTextbookOn.Insert-fm','教学信息出库--新增','/api/v1/em-textbook-on','BUS','POST','2024-03-10 20:29:28.107','2024-03-10 20:54:37.248',NULL,0,0),(201,'go-admin/app/admin/apis.EmTextbookIn.Update-fm','教学信息入库--修改','/api/v1/em-textbook-in/:id','BUS','PUT','2024-03-10 20:29:28.214','2024-03-10 20:33:49.036',NULL,0,0),(202,'go-admin/app/admin/apis.EmTextbookOn.Update-fm','教学信息出库--修改','/api/v1/em-textbook-on/:id','BUS','PUT','2024-03-10 20:29:28.220','2024-03-10 20:54:37.248',NULL,0,0),(203,'go-admin/app/admin/apis.EmTextbookIn.Delete-fm','教学信息入库--删除','/api/v1/em-textbook-in','BUS','DELETE','2024-03-10 20:29:28.259','2024-03-10 20:33:56.802',NULL,0,0),(204,'go-admin/app/admin/apis.EmTextbookOn.Delete-fm','教学信息出库--删除','/api/v1/em-textbook-on','BUS','DELETE','2024-03-10 20:29:28.263','2024-03-10 20:54:37.248',NULL,0,0),(205,'go-admin/app/admin/apis.EmTextbookRev.GetPage-fm','教材领取信息--list','/api/v1/em-textbook-rev','BUS','GET','2024-03-10 20:44:33.845','2024-03-10 20:54:37.248',NULL,0,0),(206,'go-admin/app/admin/apis.EmTextbookRev.Get-fm','教材领取信息--详情','/api/v1/em-textbook-rev/:id','BUS','GET','2024-03-10 20:44:33.850','2024-03-10 20:45:25.752',NULL,0,0),(207,'go-admin/app/admin/apis.EmTextbookRev.Insert-fm','教材领取信息--新增','/api/v1/em-textbook-rev','BUS','POST','2024-03-10 20:44:33.945','2024-03-10 20:54:37.248',NULL,0,0),(208,'go-admin/app/admin/apis.EmTextbookRev.Update-fm','教材领取信息--修改','/api/v1/em-textbook-rev/:id','BUS','PUT','2024-03-10 20:44:33.977','2024-03-10 20:54:37.248',NULL,0,0),(209,'go-admin/app/admin/apis.EmTextbookRev.Delete-fm','教材领取信息--删除','/api/v1/em-textbook-rev','BUS','DELETE','2024-03-10 20:44:34.015','2024-03-10 20:54:37.248',NULL,0,0);

/*Table structure for table `sys_casbin_rule` */

DROP TABLE IF EXISTS `sys_casbin_rule`;

CREATE TABLE `sys_casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v6` varchar(25) COLLATE utf8mb4_bin DEFAULT NULL,
  `v7` varchar(25) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sys_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

/*Data for the table `sys_casbin_rule` */

insert  into `sys_casbin_rule`(`id`,`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`) values (9,'p','class-manager','/api/v1/em-class-info','DELETE','','','','',''),(6,'p','class-manager','/api/v1/em-class-info','GET','','','','',''),(7,'p','class-manager','/api/v1/em-class-info','POST','','','','',''),(8,'p','class-manager','/api/v1/em-class-info/:id','PUT','','','','',''),(10,'p','class-manager','/api/v1/em-student','GET','','','','',''),(11,'p','class-manager','/api/v1/em-student/:id','PUT','','','','',''),(80,'p','college-manager','/api/v1/em-class-info','DELETE','','','','',''),(77,'p','college-manager','/api/v1/em-class-info','GET','','','','',''),(78,'p','college-manager','/api/v1/em-class-info','POST','','','','',''),(79,'p','college-manager','/api/v1/em-class-info/:id','PUT','','','','',''),(88,'p','college-manager','/api/v1/em-college-info','DELETE','','','','',''),(85,'p','college-manager','/api/v1/em-college-info','GET','','','','',''),(86,'p','college-manager','/api/v1/em-college-info','POST','','','','',''),(87,'p','college-manager','/api/v1/em-college-info/:id','PUT','','','','',''),(92,'p','college-manager','/api/v1/em-major-info','DELETE','','','','',''),(89,'p','college-manager','/api/v1/em-major-info','GET','','','','',''),(90,'p','college-manager','/api/v1/em-major-info','POST','','','','',''),(91,'p','college-manager','/api/v1/em-major-info/:id','PUT','','','','',''),(84,'p','college-manager','/api/v1/em-student','DELETE','','','','',''),(81,'p','college-manager','/api/v1/em-student','GET','','','','',''),(82,'p','college-manager','/api/v1/em-student','POST','','','','',''),(83,'p','college-manager','/api/v1/em-student/:id','PUT','','','','',''),(96,'p','college-manager','/api/v1/em-textbook','DELETE','','','','',''),(93,'p','college-manager','/api/v1/em-textbook','GET','','','','',''),(94,'p','college-manager','/api/v1/em-textbook','POST','','','','',''),(101,'p','college-manager','/api/v1/em-textbook-in','GET','','','','',''),(104,'p','college-manager','/api/v1/em-textbook-on','DELETE','','','','',''),(105,'p','college-manager','/api/v1/em-textbook-on','GET','','','','',''),(102,'p','college-manager','/api/v1/em-textbook-on','POST','','','','',''),(103,'p','college-manager','/api/v1/em-textbook-on/:id','PUT','','','','',''),(109,'p','college-manager','/api/v1/em-textbook-rev','DELETE','','','','',''),(106,'p','college-manager','/api/v1/em-textbook-rev','GET','','','','',''),(107,'p','college-manager','/api/v1/em-textbook-rev','POST','','','','',''),(108,'p','college-manager','/api/v1/em-textbook-rev/:id','PUT','','','','',''),(100,'p','college-manager','/api/v1/em-textbook-sup','DELETE','','','','',''),(97,'p','college-manager','/api/v1/em-textbook-sup','GET','','','','',''),(98,'p','college-manager','/api/v1/em-textbook-sup','POST','','','','',''),(99,'p','college-manager','/api/v1/em-textbook-sup/:id','PUT','','','','',''),(95,'p','college-manager','/api/v1/em-textbook/:id','PUT','','','','','');

/*Table structure for table `sys_columns` */

DROP TABLE IF EXISTS `sys_columns`;

CREATE TABLE `sys_columns` (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `table_id` bigint(20) DEFAULT NULL,
  `column_name` varchar(128) DEFAULT NULL,
  `column_comment` varchar(128) DEFAULT NULL,
  `column_type` varchar(128) DEFAULT NULL,
  `go_type` varchar(128) DEFAULT NULL,
  `go_field` varchar(128) DEFAULT NULL,
  `json_field` varchar(128) DEFAULT NULL,
  `is_pk` varchar(4) DEFAULT NULL,
  `is_increment` varchar(4) DEFAULT NULL,
  `is_required` varchar(4) DEFAULT NULL,
  `is_insert` varchar(4) DEFAULT NULL,
  `is_edit` varchar(4) DEFAULT NULL,
  `is_list` varchar(4) DEFAULT NULL,
  `is_query` varchar(4) DEFAULT NULL,
  `query_type` varchar(128) DEFAULT NULL,
  `html_type` varchar(128) DEFAULT NULL,
  `dict_type` varchar(128) DEFAULT NULL,
  `sort` bigint(20) DEFAULT NULL,
  `list` varchar(1) DEFAULT NULL,
  `pk` tinyint(1) DEFAULT NULL,
  `required` tinyint(1) DEFAULT NULL,
  `super_column` tinyint(1) DEFAULT NULL,
  `usable_column` tinyint(1) DEFAULT NULL,
  `increment` tinyint(1) DEFAULT NULL,
  `insert` tinyint(1) DEFAULT NULL,
  `edit` tinyint(1) DEFAULT NULL,
  `query` tinyint(1) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `fk_table_name` longtext,
  `fk_table_name_class` longtext,
  `fk_table_name_package` longtext,
  `fk_label_id` longtext,
  `fk_label_name` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`column_id`),
  KEY `idx_sys_columns_deleted_at` (`deleted_at`),
  KEY `idx_sys_columns_create_by` (`create_by`),
  KEY `idx_sys_columns_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_columns` */

insert  into `sys_columns`(`column_id`,`table_id`,`column_name`,`column_comment`,`column_type`,`go_type`,`go_field`,`json_field`,`is_pk`,`is_increment`,`is_required`,`is_insert`,`is_edit`,`is_list`,`is_query`,`query_type`,`html_type`,`dict_type`,`sort`,`list`,`pk`,`required`,`super_column`,`usable_column`,`increment`,`insert`,`edit`,`query`,`remark`,`fk_table_name`,`fk_table_name_class`,`fk_table_name_package`,`fk_label_id`,`fk_label_name`,`created_at`,`updated_at`,`deleted_at`,`create_by`,`update_by`) values (1,1,'id','','int(10)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.784','2024-03-10 17:40:28.784',NULL,0,0),(2,1,'name','班级名','varchar(50)','string','Name','name','0','','0','1','','1','1','LIKE','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.787','2024-03-10 17:41:03.854',NULL,0,0),(3,1,'major','专业','varchar(30)','string','Major','major','0','','0','1','','1','1','LIKE','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.791','2024-03-10 17:41:03.859',NULL,0,0),(4,1,'college','学院','varchar(30)','string','College','college','0','','0','1','','1','1','LIKE','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.795','2024-03-10 17:41:03.863',NULL,0,0),(5,1,'re_person','负责人','varchar(30)','string','RePerson','rePerson','0','','0','1','','1','1','LIKE','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.798','2024-03-10 17:41:03.868',NULL,0,0),(6,1,'re_person_no','负责人学号','int(30)','string','RePersonNo','rePersonNo','0','','0','1','','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.801','2024-03-10 17:41:03.872',NULL,0,0),(7,1,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.804','2024-03-10 17:40:28.804',NULL,0,0),(8,1,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.808','2024-03-10 17:40:28.808',NULL,0,0),(9,1,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.811','2024-03-10 17:40:28.811',NULL,0,0),(10,1,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.813','2024-03-10 17:40:28.813',NULL,0,0),(11,1,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.817','2024-03-10 17:40:28.817',NULL,0,0),(12,2,'id','','int(10)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.826','2024-03-10 17:40:28.826',NULL,0,0),(13,2,'name','学院名','varchar(50)','string','Name','name','0','','0','1','','1','1','LIKE','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.828','2024-03-10 18:00:15.946',NULL,0,0),(14,2,'bk','学院简介','varchar(200)','string','Bk','bk','0','','0','1','','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.831','2024-03-10 18:00:15.950',NULL,0,0),(15,2,'re_person','负责人姓名','varchar(50)','string','RePerson','rePerson','0','','0','1','','1','1','LIKE','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.835','2024-03-10 18:00:15.955',NULL,0,0),(16,2,'re_phone','负责人手机号','varchar(50)','string','RePhone','rePhone','0','','0','1','','1','1','LIKE','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.838','2024-03-10 18:00:15.959',NULL,0,0),(17,2,'re_no','负责人工号','int(20)','string','ReNo','reNo','0','','0','1','','1','1','LIKE','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.840','2024-03-10 18:00:15.964',NULL,0,0),(18,2,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.844','2024-03-10 17:40:28.844',NULL,0,0),(19,2,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.847','2024-03-10 17:40:28.847',NULL,0,0),(20,2,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.850','2024-03-10 17:40:28.850',NULL,0,0),(21,2,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.853','2024-03-10 17:40:28.853',NULL,0,0),(22,2,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.856','2024-03-10 17:40:28.856',NULL,0,0),(23,3,'id','','int(12)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.866','2024-03-10 17:40:28.866',NULL,0,0),(24,3,'major_name','所属专业','varchar(10)','string','MajorName','majorName','0','','0','1','','1','1','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.869','2024-03-10 18:02:07.597',NULL,0,0),(25,3,'major_bk','简介','varchar(200)','string','MajorBk','majorBk','0','','0','1','','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.873','2024-03-10 18:02:07.602',NULL,0,0),(26,3,'college_name','所属学院','varchar(12)','string','CollegeName','collegeName','0','','0','1','','1','1','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.876','2024-03-10 18:02:07.606',NULL,0,0),(27,3,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.879','2024-03-10 17:40:28.879',NULL,0,0),(28,3,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.882','2024-03-10 17:40:28.882',NULL,0,0),(29,3,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.885','2024-03-10 17:40:28.885',NULL,0,0),(30,3,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.888','2024-03-10 17:40:28.888',NULL,0,0),(31,3,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.891','2024-03-10 17:40:28.891',NULL,0,0),(32,4,'id','','int(12)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.901','2024-03-10 17:40:28.901',NULL,0,0),(33,4,'sex','性别','int(1)','string','Sex','sex','0','','0','1','','1','0','EQ','radio','sys_user_sex',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.903','2024-03-10 20:55:39.349',NULL,0,0),(34,4,'name','姓名','varchar(20)','string','Name','name','0','','0','1','','1','1','LIKE','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.907','2024-03-10 20:55:39.354',NULL,0,0),(35,4,'number','学号','int(20)','string','Number','number','0','','0','1','','1','1','LIKE','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.910','2024-03-10 20:55:39.358',NULL,0,0),(36,4,'class_name','班级名','varchar(50)','string','ClassName','className','0','','0','1','','1','1','LIKE','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.913','2024-03-10 20:55:39.362',NULL,0,0),(37,4,'major','专业名','varchar(50)','string','Major','major','0','','0','1','','1','1','LIKE','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.916','2024-03-10 20:55:39.366',NULL,0,0),(38,4,'college_name','学院名','varchar(50)','string','CollegeName','collegeName','0','','0','1','','1','1','LIKE','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.919','2024-03-10 20:55:39.371',NULL,0,0),(39,4,'phone','手机号','varchar(12)','string','Phone','phone','0','','0','1','','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.922','2024-03-10 20:55:39.375',NULL,0,0),(40,4,'email','邮箱','varchar(50)','string','Email','email','0','','0','1','','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.925','2024-03-10 20:55:39.379',NULL,0,0),(41,4,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.928','2024-03-10 17:40:28.928',NULL,0,0),(42,4,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.931','2024-03-10 17:40:28.931',NULL,0,0),(43,4,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.935','2024-03-10 17:40:28.935',NULL,0,0),(44,4,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.939','2024-03-10 17:40:28.939',NULL,0,0),(45,4,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',14,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 17:40:28.942','2024-03-10 17:40:28.942',NULL,0,0),(46,5,'id','','int(10)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.767','2024-03-10 19:51:41.767',NULL,0,0),(47,5,'name','教材名','varchar(50)','string','Name','name','0','','0','1','','1','1','LIKE','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.771','2024-03-10 20:06:39.562',NULL,0,0),(48,5,'price','价格','int(10)','string','Price','price','0','','0','1','','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.774','2024-03-10 20:06:39.567',NULL,0,0),(49,5,'sup','教材供应商','int(10)','string','Sup','sup','0','','0','1','','1','1','EQ','select','',4,'',0,0,0,0,0,1,0,0,'','em_textbook_sup','EmTextbookSup','em-textbook-sup','id','name','2024-03-10 19:51:41.777','2024-03-10 20:06:39.571',NULL,0,0),(50,5,'pro_data','出版日期','timestamp','time.Time','ProData','proData','0','','1','1','','1','1','EQ','datetime','',5,'',0,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.780','2024-03-10 20:06:39.576',NULL,0,0),(51,5,'author','出版作者','varchar(50)','string','Author','author','0','','0','1','','1','1','LIKE','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.783','2024-03-10 20:06:39.580',NULL,0,0),(52,5,'press','出版社','varchar(50)','string','Press','press','0','','0','1','','1','1','LIKE','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.787','2024-03-10 20:06:39.584',NULL,0,0),(53,5,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.790','2024-03-10 19:51:41.790',NULL,0,0),(54,5,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.793','2024-03-10 19:51:41.793',NULL,0,0),(55,5,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.796','2024-03-10 19:51:41.796',NULL,0,0),(56,5,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.799','2024-03-10 19:51:41.799',NULL,0,0),(57,5,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.802','2024-03-10 19:51:41.802',NULL,0,0),(58,6,'id','','int(10)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.812','2024-03-10 19:51:41.812',NULL,0,0),(59,6,'textbook_id','教材_id','int(10)','string','TextbookId','textbookId','0','','1','1','','1','1','EQ','select','',2,'',0,0,0,0,0,1,0,0,'','em_textbook','EmTextbook','em-textbook','id','name','2024-03-10 19:51:41.815','2024-03-10 20:36:54.046',NULL,0,0),(60,6,'in_time','入库时间','timestamp','time.Time','InTime','inTime','0','','1','1','','1','','EQ','datetime','',3,'',0,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.818','2024-03-10 20:36:54.050',NULL,0,0),(61,6,'price','定价','int(10)','string','Price','price','0','','0','1','','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.821','2024-03-10 20:36:54.055',NULL,0,0),(62,6,'num','数量','int(10)','string','Num','num','0','','0','1','','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.824','2024-03-10 20:36:54.059',NULL,0,0),(63,6,'textbook_pro','供应商','int(10)','string','TextbookPro','textbookPro','0','','1','1','','1','1','EQ','select','',6,'',0,0,0,0,0,1,0,0,'','em_textbook_sup','EmTextbookSup','em-textbook-sup','id','name','2024-03-10 19:51:41.827','2024-03-10 20:36:54.064',NULL,0,0),(64,6,'textbook_unit','经办单位','varchar(20)','string','TextbookUnit','textbookUnit','0','','1','1','','1','1','LIKE','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.830','2024-03-10 20:36:54.068',NULL,0,0),(65,7,'id','','int(10)','string','Id','id','0','','0','1','','','','EQ','input','',1,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.839','2024-03-10 19:51:41.839',NULL,0,0),(66,7,'textbook_id','教材_id','int(10)','string','TextbookId','textbookId','0','','1','1','','1','1','EQ','select','',2,'',0,0,0,0,0,1,0,0,'','em_textbook','EmTextbook','em-textbook','id','name','2024-03-10 19:51:41.842','2024-03-10 20:30:40.164',NULL,0,0),(67,7,'num','数量','int(10)','string','Num','num','0','','1','1','','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.845','2024-03-10 20:30:40.169',NULL,0,0),(68,7,'textbook_unit','经办单位','varchar(50)','string','TextbookUnit','textbookUnit','0','','1','1','','1','1','LIKE','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.849','2024-03-10 20:30:40.173',NULL,0,0),(69,7,'out_time','出库时间','timestamp','time.Time','OutTime','outTime','0','','1','1','','1','0','EQ','datetime','',5,'',0,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.851','2024-03-10 20:30:40.177',NULL,0,0),(70,7,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.854','2024-03-10 19:51:41.854',NULL,0,0),(71,7,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.858','2024-03-10 19:51:41.858',NULL,0,0),(72,7,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.861','2024-03-10 19:51:41.861',NULL,0,0),(73,7,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.864','2024-03-10 19:51:41.864',NULL,0,0),(74,7,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.867','2024-03-10 19:51:41.867',NULL,0,0),(75,8,'id','','int(10)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.876','2024-03-10 19:51:41.876',NULL,0,0),(76,8,'rev_time','领取日期','timestamp','time.Time','RevTime','revTime','0','','1','1','','1','1','EQ','datetime','',2,'',0,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.879','2024-03-10 20:44:20.671',NULL,0,0),(77,8,'class_id','班级名','int(10)','string','ClassId','classId','0','','0','1','','1','1','LIKE','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.882','2024-03-10 20:44:20.675',NULL,0,0),(78,8,'textbook_id','教材名','int(10)','string','TextbookId','textbookId','0','','0','1','','1','1','EQ','select','',4,'',0,0,0,0,0,1,0,0,'','em_textbook','EmTextbook','em-textbook','id','name','2024-03-10 19:51:41.885','2024-03-10 20:44:20.680',NULL,0,0),(79,8,'num','数量','int(11)','string','Num','num','0','','0','1','','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.888','2024-03-10 20:44:20.684',NULL,0,0),(80,8,'class_leader','班级负责人','varchar(50)','string','ClassLeader','classLeader','0','','0','1','','1','1','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.891','2024-03-10 20:44:20.688',NULL,0,0),(81,8,'class_leader_phone','负责人电话','int(12)','string','ClassLeaderPhone','classLeaderPhone','0','','0','1','','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.894','2024-03-10 20:44:20.692',NULL,0,0),(82,9,'id','','int(10)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.904','2024-03-10 19:51:41.904',NULL,0,0),(83,9,'name','供应商名','varchar(50)','string','Name','name','0','','0','1','','1','1','LIKE','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.907','2024-03-10 20:12:05.399',NULL,0,0),(84,9,'phone','供应商电话','int(12)','string','Phone','phone','0','','0','1','','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.909','2024-03-10 20:12:05.403',NULL,0,0),(85,9,'addr','供应商地址','varchar(50)','string','Addr','addr','0','','0','1','','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.913','2024-03-10 20:12:05.408',NULL,0,0),(86,9,'time','供应商开始供应时间','timestamp','time.Time','Time','time','0','','1','1','','1','','EQ','datetime','',5,'',0,1,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.916','2024-03-10 20:12:05.412',NULL,0,0),(87,9,'status','状态','int(1)','string','Status','status','0','','0','1','','1','1','EQ','radio','sys_job_status',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.919','2024-03-10 20:12:05.416',NULL,0,0),(88,9,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.922','2024-03-10 19:51:41.922',NULL,0,0),(89,9,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.925','2024-03-10 19:51:41.925',NULL,0,0),(90,9,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.928','2024-03-10 19:51:41.928',NULL,0,0),(91,9,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.931','2024-03-10 19:51:41.931',NULL,0,0),(92,9,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 19:51:41.934','2024-03-10 19:51:41.934',NULL,0,0),(93,10,'id','','int(10)','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.217','2024-03-10 20:31:04.217',NULL,0,0),(94,10,'textbook_id','教材_id','int(10)','string','TextbookId','textbookId','0','','0','1','','1','1','EQ','select','',2,'',0,0,0,0,0,1,0,0,'','em_textbook','EmTextbook','em-textbook','id','name','2024-03-10 20:31:04.219','2024-03-10 20:32:17.841',NULL,0,0),(95,10,'num','数量','int(10)','string','Num','num','0','','0','1','','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.224','2024-03-10 20:32:17.846',NULL,0,0),(96,10,'textbook_unit','经办单位','varchar(50)','string','TextbookUnit','textbookUnit','0','','0','1','','1','1','LIKE','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.227','2024-03-10 20:32:17.850',NULL,0,0),(97,10,'out_time','出库时间','timestamp','time.Time','OutTime','outTime','0','','1','1','','1','','EQ','datetime','',5,'',0,1,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.230','2024-03-10 20:32:17.854',NULL,0,0),(98,10,'create_by','创建者','bigint(20)','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.233','2024-03-10 20:31:04.233',NULL,0,0),(99,10,'update_by','更新者','bigint(20)','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.237','2024-03-10 20:31:04.237',NULL,0,0),(100,10,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.240','2024-03-10 20:31:04.240',NULL,0,0),(101,10,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.243','2024-03-10 20:31:04.243',NULL,0,0),(102,10,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-03-10 20:31:04.246','2024-03-10 20:31:04.246',NULL,0,0);

/*Table structure for table `sys_config` */

DROP TABLE IF EXISTS `sys_config`;

CREATE TABLE `sys_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) DEFAULT NULL COMMENT 'ConfigType',
  `is_frontend` varchar(64) DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) DEFAULT NULL COMMENT 'Remark',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_sys_config_deleted_at` (`deleted_at`),
  KEY `idx_sys_config_create_by` (`create_by`),
  KEY `idx_sys_config_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_config` */

insert  into `sys_config`(`id`,`config_name`,`config_key`,`config_value`,`config_type`,`is_frontend`,`remark`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'皮肤样式','sys_index_skinName','skin-green','Y','0','主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow',1,1,'2021-05-13 19:56:37.913','2021-06-05 13:50:13.123',NULL),(2,'初始密码','sys_user_initPassword','123456','Y','0','用户管理-账号初始密码:123456',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),(3,'侧栏主题','sys_index_sideTheme','theme-dark','Y','0','主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),(4,'系统名称','sys_app_name','go-admin管理系统','Y','1','',1,0,'2021-03-17 08:52:06.067','2021-05-28 10:08:25.248',NULL),(5,'系统logo','sys_app_logo','https://doc-image.zhangwj.com/img/go-admin.png','Y','1','',1,0,'2021-03-17 08:53:19.462','2021-03-17 08:53:19.462',NULL);

/*Table structure for table `sys_dept` */

DROP TABLE IF EXISTS `sys_dept`;

CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) DEFAULT NULL,
  `dept_path` varchar(255) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `leader` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`),
  KEY `idx_sys_dept_deleted_at` (`deleted_at`),
  KEY `idx_sys_dept_create_by` (`create_by`),
  KEY `idx_sys_dept_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_dept` */

insert  into `sys_dept`(`dept_id`,`parent_id`,`dept_path`,`dept_name`,`sort`,`leader`,`phone`,`email`,`status`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,0,'/0/1/','爱拓科技',0,'aituo','13782218188','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.913','2021-06-05 17:06:44.960',NULL),(7,1,'/0/1/7/','研发部',1,'aituo','13782218188','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.913','2021-06-16 21:35:00.109',NULL),(8,1,'/0/1/8/','运维部',0,'aituo','13782218188','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.913','2021-06-16 21:41:39.747',NULL),(9,1,'/0/1/9/','客服部',0,'aituo','13782218188','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.913','2021-06-05 17:07:05.993',NULL),(10,1,'/0/1/10/','人力资源',3,'aituo','13782218188','atuo@aituo.com',1,1,1,'2021-05-13 19:56:37.913','2021-06-05 17:07:08.503',NULL);

/*Table structure for table `sys_dict_data` */

DROP TABLE IF EXISTS `sys_dict_data`;

CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT,
  `dict_sort` bigint(20) DEFAULT NULL,
  `dict_label` varchar(128) DEFAULT NULL,
  `dict_value` varchar(255) DEFAULT NULL,
  `dict_type` varchar(64) DEFAULT NULL,
  `css_class` varchar(128) DEFAULT NULL,
  `list_class` varchar(128) DEFAULT NULL,
  `is_default` varchar(8) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `default` varchar(8) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`),
  KEY `idx_sys_dict_data_create_by` (`create_by`),
  KEY `idx_sys_dict_data_update_by` (`update_by`),
  KEY `idx_sys_dict_data_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_dict_data` */

insert  into `sys_dict_data`(`dict_code`,`dict_sort`,`dict_label`,`dict_value`,`dict_type`,`css_class`,`list_class`,`is_default`,`status`,`default`,`remark`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,0,'正常','2','sys_normal_disable','','','',2,'','系统正常',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:40.168',NULL),(2,0,'停用','1','sys_normal_disable','','','',2,'','系统停用',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(3,0,'男','0','sys_user_sex','','','',2,'','性别男',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(4,0,'女','1','sys_user_sex','','','',2,'','性别女',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(5,0,'未知','2','sys_user_sex','','','',2,'','性别未知',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(6,0,'显示','0','sys_show_hide','','','',2,'','显示菜单',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(7,0,'隐藏','1','sys_show_hide','','','',2,'','隐藏菜单',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(8,0,'是','Y','sys_yes_no','','','',2,'','系统默认是',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(9,0,'否','N','sys_yes_no','','','',2,'','系统默认否',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(10,0,'正常','2','sys_job_status','','','',2,'','正常状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(11,0,'停用','1','sys_job_status','','','',2,'','停用状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(12,0,'默认','DEFAULT','sys_job_group','','','',2,'','默认分组',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(13,0,'系统','SYSTEM','sys_job_group','','','',2,'','系统分组',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(14,0,'通知','1','sys_notice_type','','','',2,'','通知',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(15,0,'公告','2','sys_notice_type','','','',2,'','公告',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(16,0,'正常','2','sys_common_status','','','',2,'','正常状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(17,0,'关闭','1','sys_common_status','','','',2,'','关闭状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(18,0,'新增','1','sys_oper_type','','','',2,'','新增操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(19,0,'修改','2','sys_oper_type','','','',2,'','修改操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(20,0,'删除','3','sys_oper_type','','','',2,'','删除操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(21,0,'授权','4','sys_oper_type','','','',2,'','授权操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(22,0,'导出','5','sys_oper_type','','','',2,'','导出操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(23,0,'导入','6','sys_oper_type','','','',2,'','导入操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(24,0,'强退','7','sys_oper_type','','','',2,'','强退操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(25,0,'生成代码','8','sys_oper_type','','','',2,'','生成操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(26,0,'清空数据','9','sys_oper_type','','','',2,'','清空操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(27,0,'成功','0','sys_notice_status','','','',2,'','成功状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(28,0,'失败','1','sys_notice_status','','','',2,'','失败状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(29,0,'登录','10','sys_oper_type','','','',2,'','登录操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(30,0,'退出','11','sys_oper_type','','','',2,'','',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(31,0,'获取验证码','12','sys_oper_type','','','',2,'','获取验证码',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(32,0,'正常','1','sys_content_status','','','',1,'','',1,1,'2021-05-13 19:56:40.845','2021-05-13 19:56:40.845',NULL),(33,1,'禁用','2','sys_content_status','','','',1,'','',1,1,'2021-05-13 19:56:40.845','2021-05-13 19:56:40.845',NULL);

/*Table structure for table `sys_dict_type` */

DROP TABLE IF EXISTS `sys_dict_type`;

CREATE TABLE `sys_dict_type` (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(128) DEFAULT NULL,
  `dict_type` varchar(128) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_id`),
  KEY `idx_sys_dict_type_update_by` (`update_by`),
  KEY `idx_sys_dict_type_deleted_at` (`deleted_at`),
  KEY `idx_sys_dict_type_create_by` (`create_by`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_dict_type` */

insert  into `sys_dict_type`(`dict_id`,`dict_name`,`dict_type`,`status`,`remark`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'系统开关','sys_normal_disable',2,'系统开关列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(2,'用户性别','sys_user_sex',2,'用户性别列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(3,'菜单状态','sys_show_hide',2,'菜单状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(4,'系统是否','sys_yes_no',2,'系统是否列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(5,'任务状态','sys_job_status',2,'任务状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(6,'任务分组','sys_job_group',2,'任务分组列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(7,'通知类型','sys_notice_type',2,'通知类型列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(8,'系统状态','sys_common_status',2,'登录状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(9,'操作类型','sys_oper_type',2,'操作类型列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(10,'通知状态','sys_notice_status',2,'通知状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),(11,'内容状态','sys_content_status',2,'',1,1,'2021-05-13 19:56:40.813','2021-05-13 19:56:40.813',NULL);

/*Table structure for table `sys_job` */

DROP TABLE IF EXISTS `sys_job`;

CREATE TABLE `sys_job` (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `job_name` varchar(255) DEFAULT NULL,
  `job_group` varchar(255) DEFAULT NULL,
  `job_type` tinyint(4) DEFAULT NULL,
  `cron_expression` varchar(255) DEFAULT NULL,
  `invoke_target` varchar(255) DEFAULT NULL,
  `args` varchar(255) DEFAULT NULL,
  `misfire_policy` bigint(20) DEFAULT NULL,
  `concurrent` tinyint(4) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `entry_id` smallint(6) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`job_id`),
  KEY `idx_sys_job_deleted_at` (`deleted_at`),
  KEY `idx_sys_job_create_by` (`create_by`),
  KEY `idx_sys_job_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_job` */

insert  into `sys_job`(`job_id`,`job_name`,`job_group`,`job_type`,`cron_expression`,`invoke_target`,`args`,`misfire_policy`,`concurrent`,`status`,`entry_id`,`created_at`,`updated_at`,`deleted_at`,`create_by`,`update_by`) values (1,'接口测试','DEFAULT',1,'0/5 * * * * ','http://localhost:8000','',1,1,1,0,'2021-05-13 19:56:37.914','2021-06-14 20:59:55.417',NULL,1,1),(2,'函数测试','DEFAULT',2,'0/5 * * * * ','ExamplesOne','参数',1,1,1,0,'2021-05-13 19:56:37.914','2021-05-31 23:55:37.221',NULL,1,1);

/*Table structure for table `sys_login_log` */

DROP TABLE IF EXISTS `sys_login_log`;

CREATE TABLE `sys_login_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `username` varchar(128) DEFAULT NULL COMMENT '用户名',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) DEFAULT NULL COMMENT '系统',
  `platform` varchar(255) DEFAULT NULL COMMENT '固件',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '登录时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `msg` varchar(255) DEFAULT NULL COMMENT '信息',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_login_log_create_by` (`create_by`),
  KEY `idx_sys_login_log_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_login_log` */

/*Table structure for table `sys_menu` */

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL,
  `title` varchar(128) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL,
  `parent_id` smallint(6) DEFAULT NULL,
  `no_cache` tinyint(1) DEFAULT NULL,
  `breadcrumb` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `visible` varchar(1) DEFAULT NULL,
  `is_frame` varchar(1) DEFAULT '0',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`menu_id`),
  KEY `idx_sys_menu_create_by` (`create_by`),
  KEY `idx_sys_menu_update_by` (`update_by`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=597 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_menu` */

insert  into `sys_menu`(`menu_id`,`menu_name`,`title`,`icon`,`path`,`paths`,`menu_type`,`action`,`permission`,`parent_id`,`no_cache`,`breadcrumb`,`component`,`sort`,`visible`,`is_frame`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (2,'Admin','系统管理','api-server','/admin','/0/2','M','无','',0,1,'','Layout',10,'0','1',0,1,'2021-05-20 21:58:45.679','2024-03-10 17:33:47.230',NULL),(3,'SysUserManage','用户管理','user','/admin/sys-user','/0/2/3','C','无','admin:sysUser:list',2,0,'','/admin/sys-user/index',10,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.236',NULL),(43,'','新增管理员','app-group-fill','','/0/2/3/43','F','POST','admin:sysUser:add',3,0,'','',10,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.263',NULL),(44,'','查询管理员','app-group-fill','','/0/2/3/44','F','GET','admin:sysUser:query',3,0,'','',40,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.265',NULL),(45,'','修改管理员','app-group-fill','','/0/2/3/45','F','PUT','admin:sysUser:edit',3,0,'','',30,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.267',NULL),(46,'','删除管理员','app-group-fill','','/0/2/3/46','F','DELETE','admin:sysUser:remove',3,0,'','',20,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.270',NULL),(51,'SysMenuManage','菜单管理','tree-table','/admin/sys-menu','/0/2/51','C','无','admin:sysMenu:list',2,1,'','/admin/sys-menu/index',30,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.238',NULL),(52,'SysRoleManage','角色管理','peoples','/admin/sys-role','/0/2/52','C','无','admin:sysRole:list',2,1,'','/admin/sys-role/index',20,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.240',NULL),(56,'SysDeptManage','部门管理','tree','/admin/sys-dept','/0/2/56','C','无','admin:sysDept:list',2,0,'','/admin/sys-dept/index',40,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.243',NULL),(57,'SysPostManage','岗位管理','pass','/admin/sys-post','/0/2/57','C','无','admin:sysPost:list',2,0,'','/admin/sys-post/index',50,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.246',NULL),(58,'Dict','字典管理','education','/admin/dict','/0/2/58','C','无','admin:sysDictType:list',2,0,'','/admin/dict/index',60,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.248',NULL),(59,'SysDictDataManage','字典数据','education','/admin/dict/data/:dictId','/0/2/59','C','无','admin:sysDictData:list',2,0,'','/admin/dict/data',100,'1','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.250',NULL),(60,'Tools','开发工具','dev-tools','/dev-tools','/0/60','M','无','',0,0,'','Layout',40,'1','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 21:20:23.861',NULL),(61,'Swagger','系统接口','guide','/dev-tools/swagger','/0/60/61','C','无','',60,0,'','/dev-tools/swagger/index',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 21:20:23.862',NULL),(62,'SysConfigManage','参数管理','swagger','/admin/sys-config','/0/2/62','C','无','admin:sysConfig:list',2,0,'','/admin/sys-config/index',70,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.252',NULL),(211,'Log','日志管理','log','/log','/0/2/211','M','','',2,0,'','/log/index',80,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.255',NULL),(212,'SysLoginLogManage','登录日志','logininfor','/admin/sys-login-log','/0/2/211/212','C','','admin:sysLoginLog:list',211,0,'','/admin/sys-login-log/index',1,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.347',NULL),(216,'OperLog','操作日志','skill','/admin/sys-oper-log','/0/2/211/216','C','','admin:sysOperLog:list',211,0,'','/admin/sys-oper-log/index',1,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.349',NULL),(220,'','新增菜单','app-group-fill','','/0/2/51/220','F','','admin:sysMenu:add',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.272',NULL),(221,'','修改菜单','app-group-fill','','/0/2/51/221','F','','admin:sysMenu:edit',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.274',NULL),(222,'','查询菜单','app-group-fill','','/0/2/51/222','F','','admin:sysMenu:query',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.277',NULL),(223,'','删除菜单','app-group-fill','','/0/2/51/223','F','','admin:sysMenu:remove',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.279',NULL),(224,'','新增角色','app-group-fill','','/0/2/52/224','F','','admin:sysRole:add',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.281',NULL),(225,'','查询角色','app-group-fill','','/0/2/52/225','F','','admin:sysRole:query',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.283',NULL),(226,'','修改角色','app-group-fill','','/0/2/52/226','F','','admin:sysRole:update',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.286',NULL),(227,'','删除角色','app-group-fill','','/0/2/52/227','F','','admin:sysRole:remove',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.288',NULL),(228,'','查询部门','app-group-fill','','/0/2/56/228','F','','admin:sysDept:query',56,0,'','',40,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.290',NULL),(229,'','新增部门','app-group-fill','','/0/2/56/229','F','','admin:sysDept:add',56,0,'','',10,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.293',NULL),(230,'','修改部门','app-group-fill','','/0/2/56/230','F','','admin:sysDept:edit',56,0,'','',30,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.295',NULL),(231,'','删除部门','app-group-fill','','/0/2/56/231','F','','admin:sysDept:remove',56,0,'','',20,'0','1',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.298',NULL),(232,'','查询岗位','app-group-fill','','/0/2/57/232','F','','admin:sysPost:query',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.300',NULL),(233,'','新增岗位','app-group-fill','','/0/2/57/233','F','','admin:sysPost:add',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.303',NULL),(234,'','修改岗位','app-group-fill','','/0/2/57/234','F','','admin:sysPost:edit',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.305',NULL),(235,'','删除岗位','app-group-fill','','/0/2/57/235','F','','admin:sysPost:remove',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.307',NULL),(236,'','查询字典','app-group-fill','','/0/2/58/236','F','','admin:sysDictType:query',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.310',NULL),(237,'','新增类型','app-group-fill','','/0/2/58/237','F','','admin:sysDictType:add',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.312',NULL),(238,'','修改类型','app-group-fill','','/0/2/58/238','F','','admin:sysDictType:edit',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.314',NULL),(239,'','删除类型','app-group-fill','','/0/2/58/239','F','','system:sysdicttype:remove',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.316',NULL),(240,'','查询数据','app-group-fill','','/0/2/59/240','F','','admin:sysDictData:query',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.319',NULL),(241,'','新增数据','app-group-fill','','/0/2/59/241','F','','admin:sysDictData:add',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.321',NULL),(242,'','修改数据','app-group-fill','','/0/2/59/242','F','','admin:sysDictData:edit',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.323',NULL),(243,'','删除数据','app-group-fill','','/0/2/59/243','F','','admin:sysDictData:remove',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.326',NULL),(244,'','查询参数','app-group-fill','','/0/2/62/244','F','','admin:sysConfig:query',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.338',NULL),(245,'','新增参数','app-group-fill','','/0/2/62/245','F','','admin:sysConfig:add',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.340',NULL),(246,'','修改参数','app-group-fill','','/0/2/62/246','F','','admin:sysConfig:edit',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.343',NULL),(247,'','删除参数','app-group-fill','','/0/2/62/247','F','','admin:sysConfig:remove',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.345',NULL),(248,'','查询登录日志','app-group-fill','','/0/2/211/212/248','F','','admin:sysLoginLog:query',212,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.352',NULL),(249,'','删除登录日志','app-group-fill','','/0/2/211/212/249','F','','admin:sysLoginLog:remove',212,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.354',NULL),(250,'','查询操作日志','app-group-fill','','/0/2/211/216/250','F','','admin:sysOperLog:query',216,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.357',NULL),(251,'','删除操作日志','app-group-fill','','/0/2/211/216/251','F','','admin:sysOperLog:remove',216,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 17:33:47.359',NULL),(261,'Gen','代码生成','code','/dev-tools/gen','/0/60/261','C','','',60,0,'','/dev-tools/gen/index',2,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 21:20:23.863',NULL),(262,'EditTable','代码生成修改','build','/dev-tools/editTable','/0/60/262','C','','',60,0,'','/dev-tools/gen/editTable',100,'1','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 21:20:23.864',NULL),(264,'Build','表单构建','build','/dev-tools/build','/0/60/264','C','','',60,0,'','/dev-tools/build/index',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-03-10 21:20:23.866',NULL),(269,'ServerMonitor','服务监控','druid','/sys-tools/monitor','/0/537/269','C','','sysTools:serverMonitor:list',537,0,'','/sys-tools/monitor',0,'0','1',1,1,'2020-04-14 00:28:19.000','2024-03-10 21:01:17.408',NULL),(459,'Schedule','定时任务','time-range','/schedule','/0/459','M','无','',0,0,'','Layout',20,'1','1',1,1,'2020-08-03 09:17:37.000','2024-03-10 21:20:06.669',NULL),(460,'ScheduleManage','Schedule','job','/schedule/manage','/0/459/460','C','无','job:sysJob:list',459,0,'','/schedule/index',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-03-10 21:20:06.670',NULL),(461,'sys_job','分页获取定时任务','app-group-fill','','/0/459/460/461','F','无','job:sysJob:query',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-03-10 21:20:06.671',NULL),(462,'sys_job','创建定时任务','app-group-fill','','/0/459/460/462','F','无','job:sysJob:add',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-03-10 21:20:06.672',NULL),(463,'sys_job','修改定时任务','app-group-fill','','/0/459/460/463','F','无','job:sysJob:edit',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-03-10 21:20:06.673',NULL),(464,'sys_job','删除定时任务','app-group-fill','','/0/459/460/464','F','无','job:sysJob:remove',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-03-10 21:20:06.674',NULL),(471,'JobLog','日志','bug','/schedule/log','/0/459/471','C','','',459,0,'','/schedule/log',0,'1','1',1,1,'2020-08-05 21:24:46.000','2024-03-10 21:20:06.676',NULL),(528,'SysApiManage','接口管理','api-doc','/admin/sys-api','/0/2/528','C','无','admin:sysApi:list',2,0,'','/admin/sys-api/index',0,'0','0',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.258',NULL),(529,'','查询接口','app-group-fill','','/0/2/528/529','F','无','admin:sysApi:query',528,0,'','',40,'0','0',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.376',NULL),(531,'','修改接口','app-group-fill','','/0/2/528/531','F','无','admin:sysApi:edit',528,0,'','',30,'0','0',0,1,'2021-05-20 22:08:44.526','2024-03-10 17:33:47.378',NULL),(537,'SysTools','系统工具','system-tools','/sys-tools','/0/537','M','','',0,0,'','Layout',30,'1','1',1,1,'2021-05-21 11:13:32.166','2024-03-10 21:01:17.409',NULL),(540,'SysConfigSet','参数设置','system-tools','/admin/sys-config/set','/0/2/540','C','','admin:sysConfigSet:list',2,0,'','/admin/sys-config/set',0,'0','1',1,1,'2021-05-25 16:06:52.560','2024-03-10 17:33:47.260',NULL),(542,'','修改','upload','','/0/2/540/542','F','','admin:sysConfigSet:update',540,0,'','',0,'0','1',1,1,'2021-06-13 11:45:48.670','2024-03-10 17:33:47.382',NULL),(543,'','班级信息','pass','/em-class-info','/0/543','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 17:41:09.157','2024-03-10 20:54:37.245',NULL),(544,'EmClassInfoManage','班级列表','pass','/admin/em-class-info','/0/543/544','C','无','admin:emClassInfo:list',543,0,'','/admin/em-class-info/index',0,'0','0',1,1,'2024-03-10 17:41:09.162','2024-03-10 20:54:37.245',NULL),(545,'','分页获取班级信息','','em_class_info','/0/543/544/545','F','无','admin:emClassInfo:query',544,0,'','',0,'0','0',1,1,'2024-03-10 17:41:09.168','2024-03-10 20:54:37.245',NULL),(546,'','创建班级信息','','em_class_info','/0/543/544/546','F','无','admin:emClassInfo:add',544,0,'','',0,'0','0',1,1,'2024-03-10 17:41:09.175','2024-03-10 20:54:37.245',NULL),(547,'','修改班级信息','','em_class_info','/0/543/544/547','F','无','admin:emClassInfo:edit',544,0,'','',0,'0','0',1,1,'2024-03-10 17:41:09.181','2024-03-10 20:54:37.245',NULL),(548,'','删除班级信息','','em_class_info','/0/543/544/548','F','无','admin:emClassInfo:remove',544,0,'','',0,'0','0',1,1,'2024-03-10 17:41:09.187','2024-03-10 20:54:37.245',NULL),(549,'','学生信息','pass','/em-student','/0/549','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 17:44:30.803','2024-03-10 17:44:30.804','2024-03-10 17:47:23.605'),(550,'EmStudentManage','学生信息','pass','/admin/em-student','/0/549/550','C','无','admin:emStudent:list',543,0,'','/admin/em-student/index',0,'0','0',1,1,'2024-03-10 17:44:30.808','2024-03-10 20:54:37.245',NULL),(551,'','分页获取学生信息','','em_student','/0/549/550/551','F','无','admin:emStudent:query',550,0,'','',0,'0','0',1,1,'2024-03-10 17:44:30.815','2024-03-10 20:54:37.245',NULL),(552,'','创建学生信息','','em_student','/0/549/550/552','F','无','admin:emStudent:add',550,0,'','',0,'0','0',1,1,'2024-03-10 17:44:30.821','2024-03-10 20:54:37.245',NULL),(553,'','修改学生信息','','em_student','/0/549/550/553','F','无','admin:emStudent:edit',550,0,'','',0,'0','0',1,1,'2024-03-10 17:44:30.827','2024-03-10 20:54:37.245',NULL),(554,'','删除学生信息','','em_student','/0/549/550/554','F','无','admin:emStudent:remove',550,0,'','',0,'0','0',1,1,'2024-03-10 17:44:30.833','2024-03-10 20:54:37.245',NULL),(555,'','学院信息','pass','/em-college-info','/0/555','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 18:00:20.665','2024-03-10 20:54:37.245',NULL),(556,'EmCollegeInfoManage','学院列表','pass','/admin/em-college-info','/0/555/556','C','无','admin:emCollegeInfo:list',555,0,'','/admin/em-college-info/index',0,'0','0',1,1,'2024-03-10 18:00:20.670','2024-03-10 20:54:37.245',NULL),(557,'','分页获取学院信息','','em_college_info','/0/555/556/557','F','无','admin:emCollegeInfo:query',556,0,'','',0,'0','0',1,1,'2024-03-10 18:00:20.676','2024-03-10 20:54:37.245',NULL),(558,'','创建学院信息','','em_college_info','/0/555/556/558','F','无','admin:emCollegeInfo:add',556,0,'','',0,'0','0',1,1,'2024-03-10 18:00:20.684','2024-03-10 20:54:37.245',NULL),(559,'','修改学院信息','','em_college_info','/0/555/556/559','F','无','admin:emCollegeInfo:edit',556,0,'','',0,'0','0',1,1,'2024-03-10 18:00:20.690','2024-03-10 20:54:37.245',NULL),(560,'','删除学院信息','','em_college_info','/0/555/556/560','F','无','admin:emCollegeInfo:remove',556,0,'','',0,'0','0',1,1,'2024-03-10 18:00:20.697','2024-03-10 20:54:37.245',NULL),(561,'','专业信息','pass','/em-major-info','/0/561','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 18:02:14.417','2024-03-10 18:02:14.418','2024-03-10 18:07:16.582'),(562,'EmMajorInfoManage','专业列表','pass','/admin/em-major-info','/0/561/562','C','无','admin:emMajorInfo:list',555,0,'','/admin/em-major-info/index',0,'0','0',1,1,'2024-03-10 18:02:14.422','2024-03-10 21:20:01.451',NULL),(563,'','分页获取专业信息','','em_major_info','/0/561/562/563','F','无','admin:emMajorInfo:query',562,0,'','',0,'0','0',1,1,'2024-03-10 18:02:14.434','2024-03-10 21:20:01.452',NULL),(564,'','创建专业信息','','em_major_info','/0/561/562/564','F','无','admin:emMajorInfo:add',562,0,'','',0,'0','0',1,1,'2024-03-10 18:02:14.442','2024-03-10 21:20:01.453',NULL),(565,'','修改专业信息','','em_major_info','/0/561/562/565','F','无','admin:emMajorInfo:edit',562,0,'','',0,'0','0',1,1,'2024-03-10 18:02:14.450','2024-03-10 21:20:01.454',NULL),(566,'','删除专业信息','','em_major_info','/0/561/562/566','F','无','admin:emMajorInfo:remove',562,0,'','',0,'0','0',1,1,'2024-03-10 18:02:14.457','2024-03-10 21:20:01.456',NULL),(567,'','教材信息','pass','/em-textbook','/0/567','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 19:54:50.043','2024-03-10 20:54:37.245',NULL),(568,'EmTextbookManage','教材信息','pass','/admin/em-textbook','/0/567/568','C','无','admin:emTextbook:list',567,0,'','/admin/em-textbook/index',0,'0','0',1,1,'2024-03-10 19:54:50.048','2024-03-10 20:54:37.245',NULL),(569,'','分页获取教材信息','','em_textbook','/0/567/568/569','F','无','admin:emTextbook:query',568,0,'','',0,'0','0',1,1,'2024-03-10 19:54:50.054','2024-03-10 20:54:37.245',NULL),(570,'','创建教材信息','','em_textbook','/0/567/568/570','F','无','admin:emTextbook:add',568,0,'','',0,'0','0',1,1,'2024-03-10 19:54:50.060','2024-03-10 20:54:37.245',NULL),(571,'','修改教材信息','','em_textbook','/0/567/568/571','F','无','admin:emTextbook:edit',568,0,'','',0,'0','0',1,1,'2024-03-10 19:54:50.068','2024-03-10 20:54:37.245',NULL),(572,'','删除教材信息','','em_textbook','/0/567/568/572','F','无','admin:emTextbook:remove',568,0,'','',0,'0','0',1,1,'2024-03-10 19:54:50.075','2024-03-10 20:54:37.245',NULL),(573,'','供应商信息','pass','/em-textbook-sup','/0/573','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 20:01:48.747','2024-03-10 20:01:48.749','2024-03-10 20:05:27.288'),(574,'EmTextbookSupManage','教材供应商信息','pass','/admin/em-textbook-sup','/0/573/574','C','无','admin:emTextbookSup:list',567,0,'','/admin/em-textbook-sup/index',0,'0','0',1,1,'2024-03-10 20:01:48.753','2024-03-10 20:54:37.245',NULL),(575,'','分页获取供应商信息','','em_textbook_sup','/0/573/574/575','F','无','admin:emTextbookSup:query',574,0,'','',0,'0','0',1,1,'2024-03-10 20:01:48.759','2024-03-10 20:54:37.245',NULL),(576,'','创建供应商信息','','em_textbook_sup','/0/573/574/576','F','无','admin:emTextbookSup:add',574,0,'','',0,'0','0',1,1,'2024-03-10 20:01:48.765','2024-03-10 20:54:37.245',NULL),(577,'','修改供应商信息','','em_textbook_sup','/0/573/574/577','F','无','admin:emTextbookSup:edit',574,0,'','',0,'0','0',1,1,'2024-03-10 20:01:48.771','2024-03-10 20:54:37.245',NULL),(578,'','删除供应商信息','','em_textbook_sup','/0/573/574/578','F','无','admin:emTextbookSup:remove',574,0,'','',0,'0','0',1,1,'2024-03-10 20:01:48.777','2024-03-10 20:54:37.245',NULL),(579,'','教材出入库信息','pass','/em-textbook-in','/0/579','M','无','',567,0,'','Layout',0,'0','0',1,1,'2024-03-10 20:25:39.606','2024-03-10 20:39:07.288','2024-03-10 20:40:30.213'),(580,'EmTextbookInManage','教材入库信息','pass','/admin/em-textbook-in','/0/579/580','C','无','admin:emTextbookIn:list',567,0,'','/admin/em-textbook-in/index',0,'0','0',1,1,'2024-03-10 20:25:39.611','2024-03-10 20:54:37.245',NULL),(581,'','分页获取教材入库信息','','em_textbook_in','/0/579/580/581','F','无','admin:emTextbookIn:query',580,0,'','',0,'0','0',1,1,'2024-03-10 20:25:39.618','2024-03-10 20:54:37.245',NULL),(582,'','创建教材入库信息','','em_textbook_in','/0/579/580/582','F','无','admin:emTextbookIn:add',580,0,'','',0,'0','0',1,1,'2024-03-10 20:25:39.625','2024-03-10 20:54:37.245',NULL),(583,'','修改教材入库信息','','em_textbook_in','/0/579/580/583','F','无','admin:emTextbookIn:edit',580,0,'','',0,'0','0',1,1,'2024-03-10 20:25:39.631','2024-03-10 20:54:37.245',NULL),(584,'','删除教材入库信息','','em_textbook_in','/0/579/580/584','F','无','admin:emTextbookIn:remove',580,0,'','',0,'0','0',1,1,'2024-03-10 20:25:39.637','2024-03-10 20:54:37.245',NULL),(585,'','教材出库信息','pass','/em-textbook-on','/0/585','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 20:25:46.024','2024-03-10 20:25:46.025','2024-03-10 20:39:23.825'),(586,'EmTextbookOnManage','教材出库信息','pass','/admin/em-textbook-on','/0/585/586','C','无','admin:emTextbookOn:list',567,0,'','/admin/em-textbook-on/index',0,'0','0',1,1,'2024-03-10 20:25:46.029','2024-03-10 20:54:37.245',NULL),(587,'','分页获取教材出库信息','','em_textbook_on','/0/585/586/587','F','无','admin:emTextbookOn:query',586,0,'','',0,'0','0',1,1,'2024-03-10 20:25:46.036','2024-03-10 20:54:37.245',NULL),(588,'','创建教材出库信息','','em_textbook_on','/0/585/586/588','F','无','admin:emTextbookOn:add',586,0,'','',0,'0','0',1,1,'2024-03-10 20:25:46.042','2024-03-10 20:54:37.245',NULL),(589,'','修改教材出库信息','','em_textbook_on','/0/585/586/589','F','无','admin:emTextbookOn:edit',586,0,'','',0,'0','0',1,1,'2024-03-10 20:25:46.049','2024-03-10 20:54:37.245',NULL),(590,'','删除教材出库信息','','em_textbook_on','/0/585/586/590','F','无','admin:emTextbookOn:remove',586,0,'','',0,'0','0',1,1,'2024-03-10 20:25:46.055','2024-03-10 20:54:37.245',NULL),(591,'','教材班级领取信息','pass','/em-textbook-rev','/0/591','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-03-10 20:44:24.924','2024-03-10 20:44:24.925','2024-03-10 20:46:52.789'),(592,'EmTextbookRevManage','教材领取信息','pass','/admin/em-textbook-rev','/0/591/592','C','无','admin:emTextbookRev:list',567,0,'','/admin/em-textbook-rev/index',0,'0','0',1,1,'2024-03-10 20:44:24.932','2024-03-10 20:54:37.245',NULL),(593,'','分页获取教材班级领取信息','','em_textbook_rev','/0/591/592/593','F','无','admin:emTextbookRev:query',592,0,'','',0,'0','0',1,1,'2024-03-10 20:44:24.940','2024-03-10 20:54:37.245',NULL),(594,'','创建教材班级领取信息','','em_textbook_rev','/0/591/592/594','F','无','admin:emTextbookRev:add',592,0,'','',0,'0','0',1,1,'2024-03-10 20:44:24.947','2024-03-10 20:54:37.245',NULL),(595,'','修改教材班级领取信息','','em_textbook_rev','/0/591/592/595','F','无','admin:emTextbookRev:edit',592,0,'','',0,'0','0',1,1,'2024-03-10 20:44:24.955','2024-03-10 20:54:37.245',NULL),(596,'','删除教材班级领取信息','','em_textbook_rev','/0/591/592/596','F','无','admin:emTextbookRev:remove',592,0,'','',0,'0','0',1,1,'2024-03-10 20:44:24.962','2024-03-10 20:54:37.245',NULL);

/*Table structure for table `sys_menu_api_rule` */

DROP TABLE IF EXISTS `sys_menu_api_rule`;

CREATE TABLE `sys_menu_api_rule` (
  `sys_menu_menu_id` bigint(20) NOT NULL,
  `sys_api_id` bigint(20) NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_menu_id`,`sys_api_id`),
  KEY `fk_sys_menu_api_rule_sys_api` (`sys_api_id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_menu_api_rule` */

insert  into `sys_menu_api_rule`(`sys_menu_menu_id`,`sys_api_id`) values (216,6),(250,6),(58,21),(236,21),(238,23),(59,24),(240,24),(242,25),(58,26),(236,26),(56,27),(228,27),(230,28),(226,29),(51,39),(222,39),(221,41),(52,44),(225,44),(226,45),(226,46),(226,47),(62,53),(244,53),(246,54),(57,59),(232,59),(234,60),(241,80),(237,81),(229,82),(245,87),(220,88),(233,89),(224,90),(531,92),(242,101),(238,102),(230,103),(226,106),(226,107),(246,108),(221,109),(234,110),(249,114),(251,115),(243,120),(239,121),(231,122),(247,125),(223,126),(235,127),(227,128),(51,135),(528,135),(529,135),(531,136),(212,137),(248,137),(542,139),(540,140),(3,141),(44,141),(45,142),(43,150),(45,151),(46,156),(545,160),(551,162),(546,166),(552,167),(547,169),(553,170),(548,172),(554,173),(557,175),(558,177),(559,178),(560,179),(563,180),(564,182),(565,183),(566,184),(569,185),(570,187),(571,188),(572,189),(575,190),(576,192),(577,193),(578,194),(581,195),(587,197),(582,200),(588,200),(583,202),(589,202),(584,204),(590,204),(593,205),(594,207),(595,208),(596,209);

/*Table structure for table `sys_migration` */

DROP TABLE IF EXISTS `sys_migration`;

CREATE TABLE `sys_migration` (
  `version` varchar(191) COLLATE utf8mb4_bin NOT NULL,
  `apply_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

/*Data for the table `sys_migration` */

insert  into `sys_migration`(`version`,`apply_time`) values ('1599190683659','2024-03-10 17:33:47.224'),('1653638869132','2024-03-10 17:33:47.383');

/*Table structure for table `sys_opera_log` */

DROP TABLE IF EXISTS `sys_opera_log`;

CREATE TABLE `sys_opera_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(255) DEFAULT NULL COMMENT '操作模块',
  `business_type` varchar(128) DEFAULT NULL COMMENT '操作类型',
  `business_types` varchar(128) DEFAULT NULL COMMENT 'BusinessTypes',
  `method` varchar(128) DEFAULT NULL COMMENT '函数',
  `request_method` varchar(128) DEFAULT NULL COMMENT '请求方式: GET POST PUT DELETE',
  `operator_type` varchar(128) DEFAULT NULL COMMENT '操作类型',
  `oper_name` varchar(128) DEFAULT NULL COMMENT '操作者',
  `dept_name` varchar(128) DEFAULT NULL COMMENT '部门名称',
  `oper_url` varchar(255) DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) DEFAULT NULL COMMENT '访问位置',
  `oper_param` text COMMENT '请求参数',
  `status` varchar(4) DEFAULT NULL COMMENT '操作状态 1:正常 2:关闭',
  `oper_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '操作时间',
  `json_result` varchar(255) DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `latency_time` varchar(128) DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) DEFAULT NULL COMMENT 'ua',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_opera_log_create_by` (`create_by`),
  KEY `idx_sys_opera_log_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_opera_log` */

/*Table structure for table `sys_post` */

DROP TABLE IF EXISTS `sys_post`;

CREATE TABLE `sys_post` (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) DEFAULT NULL,
  `post_code` varchar(128) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`),
  KEY `idx_sys_post_create_by` (`create_by`),
  KEY `idx_sys_post_update_by` (`update_by`),
  KEY `idx_sys_post_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_post` */

insert  into `sys_post`(`post_id`,`post_name`,`post_code`,`sort`,`status`,`remark`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'首席执行官','CEO',0,2,'首席执行官',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),(2,'首席技术执行官','CTO',2,2,'首席技术执行官',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),(3,'首席运营官','COO',3,2,'测试工程师',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL);

/*Table structure for table `sys_role` */

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL,
  `status` varchar(4) DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` bigint(20) DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_update_by` (`update_by`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`),
  KEY `idx_sys_role_create_by` (`create_by`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_role` */

insert  into `sys_role`(`role_id`,`role_name`,`status`,`role_key`,`role_sort`,`flag`,`remark`,`admin`,`data_scope`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'系统管理员','2','admin',1,'','',1,'',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),(2,'班级管理人员','2','class-manager',0,'','',0,'',0,0,'2024-03-10 17:49:16.015','2024-03-10 17:50:34.239',NULL),(3,'学院管理人员','2','college-manager',0,'','',0,'',0,0,'2024-03-10 20:47:51.214','2024-03-10 20:54:37.244',NULL);

/*Table structure for table `sys_role_dept` */

DROP TABLE IF EXISTS `sys_role_dept`;

CREATE TABLE `sys_role_dept` (
  `role_id` smallint(6) NOT NULL,
  `dept_id` smallint(6) NOT NULL,
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_role_dept` */

/*Table structure for table `sys_role_menu` */

DROP TABLE IF EXISTS `sys_role_menu`;

CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL,
  `menu_id` bigint(20) NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `fk_sys_role_menu_sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_role_menu` */

insert  into `sys_role_menu`(`role_id`,`menu_id`) values (3,543),(2,544),(3,544),(2,545),(3,545),(2,546),(3,546),(2,547),(3,547),(2,548),(3,548),(3,550),(2,551),(3,551),(3,552),(2,553),(3,553),(3,554),(3,555),(3,556),(3,557),(3,558),(3,559),(3,560),(3,562),(3,563),(3,564),(3,565),(3,566),(3,567),(3,568),(3,569),(3,570),(3,571),(3,572),(3,574),(3,575),(3,576),(3,577),(3,578),(3,580),(3,581),(3,582),(3,583),(3,584),(3,586),(3,587),(3,588),(3,589),(3,590),(3,592),(3,593),(3,594),(3,595),(3,596);

/*Table structure for table `sys_tables` */

DROP TABLE IF EXISTS `sys_tables`;

CREATE TABLE `sys_tables` (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) DEFAULT NULL,
  `table_comment` varchar(255) DEFAULT NULL,
  `class_name` varchar(255) DEFAULT NULL,
  `tpl_category` varchar(255) DEFAULT NULL,
  `package_name` varchar(255) DEFAULT NULL,
  `module_name` varchar(255) DEFAULT NULL,
  `module_front_name` varchar(255) DEFAULT NULL COMMENT '前端文件名',
  `business_name` varchar(255) DEFAULT NULL,
  `function_name` varchar(255) DEFAULT NULL,
  `function_author` varchar(255) DEFAULT NULL,
  `pk_column` varchar(255) DEFAULT NULL,
  `pk_go_field` varchar(255) DEFAULT NULL,
  `pk_json_field` varchar(255) DEFAULT NULL,
  `options` varchar(255) DEFAULT NULL,
  `tree_code` varchar(255) DEFAULT NULL,
  `tree_parent_code` varchar(255) DEFAULT NULL,
  `tree_name` varchar(255) DEFAULT NULL,
  `tree` tinyint(1) DEFAULT '0',
  `crud` tinyint(1) DEFAULT '1',
  `remark` varchar(255) DEFAULT NULL,
  `is_data_scope` tinyint(4) DEFAULT NULL,
  `is_actions` tinyint(4) DEFAULT NULL,
  `is_auth` tinyint(4) DEFAULT NULL,
  `is_logical_delete` varchar(1) DEFAULT NULL,
  `logical_delete` tinyint(1) DEFAULT NULL,
  `logical_delete_column` varchar(128) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`table_id`),
  KEY `idx_sys_tables_deleted_at` (`deleted_at`),
  KEY `idx_sys_tables_create_by` (`create_by`),
  KEY `idx_sys_tables_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_tables` */

insert  into `sys_tables`(`table_id`,`table_name`,`table_comment`,`class_name`,`tpl_category`,`package_name`,`module_name`,`module_front_name`,`business_name`,`function_name`,`function_author`,`pk_column`,`pk_go_field`,`pk_json_field`,`options`,`tree_code`,`tree_parent_code`,`tree_name`,`tree`,`crud`,`remark`,`is_data_scope`,`is_actions`,`is_auth`,`is_logical_delete`,`logical_delete`,`logical_delete_column`,`created_at`,`updated_at`,`deleted_at`,`create_by`,`update_by`) values (1,'em_class_info','班级信息','EmClassInfo','crud','admin','em-class-info','','emClassInfo','班级信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 17:40:28.780','2024-03-10 17:41:03.849',NULL,0,0),(2,'em_college_info','学院信息','EmCollegeInfo','crud','admin','em-college-info','','emCollegeInfo','学院信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 17:40:28.822','2024-03-10 18:00:15.941',NULL,0,0),(3,'em_major_info','专业信息','EmMajorInfo','crud','admin','em-major-info','','emMajorInfo','专业信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 17:40:28.863','2024-03-10 18:02:07.592',NULL,0,0),(4,'em_student','学生信息','EmStudent','crud','admin','em-student','','emStudent','学生信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 17:40:28.897','2024-03-10 20:55:39.344',NULL,0,0),(5,'em_textbook','教材信息','EmTextbook','crud','admin','em-textbook','','emTextbook','教材信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 19:51:41.764','2024-03-10 20:06:39.556',NULL,0,0),(6,'em_textbook_in','教材入库信息','EmTextbookIn','crud','admin','em-textbook-in','','emTextbookIn','教材入库信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 19:51:41.808','2024-03-10 20:36:54.038',NULL,0,0),(8,'em_textbook_rev','教材班级领取信息','EmTextbookRev','crud','admin','em-textbook-rev','','emTextbookRev','教材班级领取信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 19:51:41.873','2024-03-10 20:44:20.664',NULL,0,0),(9,'em_textbook_sup','供应商信息','EmTextbookSup','crud','admin','em-textbook-sup','','emTextbookSup','供应商信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 19:51:41.900','2024-03-10 20:12:05.394',NULL,0,0),(10,'em_textbook_on','教材出库信息','EmTextbookOn','crud','admin','em-textbook-on','','emTextbookOn','教材出库信息','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-03-10 20:31:04.212','2024-03-10 20:32:17.835',NULL,0,0);

/*Table structure for table `sys_user` */

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `role_id` bigint(20) DEFAULT NULL COMMENT '角色ID',
  `salt` varchar(255) DEFAULT NULL COMMENT '加盐',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `sex` varchar(255) DEFAULT NULL COMMENT '性别',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门',
  `post_id` bigint(20) DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_user_create_by` (`create_by`),
  KEY `idx_sys_user_update_by` (`update_by`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_user` */

insert  into `sys_user`(`user_id`,`username`,`password`,`nick_name`,`phone`,`role_id`,`salt`,`avatar`,`sex`,`email`,`dept_id`,`post_id`,`remark`,`status`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`) values (1,'admin','$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.','zhangwj','13818888888',1,'','','1','1@qq.com',1,1,'','2',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:40.205',NULL),(2,'xiaomaolv','$2a$10$lDL8Ut9CXCANPVYzgseJkuCxLBJXP6uohfJt..ZcfdzhIK4Cb14ee','xml','18174010019',2,'','','0','aa@qq.com',1,1,'','2',1,0,'2024-03-10 17:49:45.912','2024-03-10 17:49:45.912',NULL),(3,'test','$2a$10$FNMs01DFuwwg8x0qlf4yZeVIeReVJxAoJYUldWcS6cjJ.fT33Noq6','test','18174010019',3,'','','0','124443@qq.com',1,1,'','2',1,0,'2024-03-10 20:48:38.791','2024-03-10 20:48:38.791',NULL);

/*Table structure for table `tb_demo` */

DROP TABLE IF EXISTS `tb_demo`;

CREATE TABLE `tb_demo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(128) DEFAULT NULL COMMENT '名称',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_tb_demo_deleted_at` (`deleted_at`),
  KEY `idx_tb_demo_create_by` (`create_by`),
  KEY `idx_tb_demo_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `tb_demo` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
