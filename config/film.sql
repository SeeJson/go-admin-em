DROP TABLE ev_film_info;

CREATE TABLE `ev_film_info` (
                                `id` INT(10) NOT NULL AUTO_INCREMENT,
                                `name` VARCHAR(50) DEFAULT NULL COMMENT '电影名',
                                `alias` VARCHAR(50) DEFAULT NULL COMMENT '别名',
                                `lead_actor` VARCHAR(100) DEFAULT NULL COMMENT '主演',
                                `director` VARCHAR(30) DEFAULT NULL COMMENT '导演',
                                `brief_introduction` TEXT DEFAULT NULL COMMENT '简介',
                                `cover` BIGINT(20) DEFAULT NULL COMMENT '封面',
                                `release_time` DATETgozeroIME(3) DEFAULT NULL COMMENT '上映时间',
                                `release_release` VARCHAR(50) DEFAULT NULL COMMENT '上映地区',
                                `country` INT(2) DEFAULT NULL COMMENT '国家',
                                `type` INT(1) DEFAULT NULL COMMENT '电影类型',
                                `status` INT(1) DEFAULT NULL COMMENT '状态 1：正在上映 2：即将上映',
                                `language` INT(1) DEFAULT NULL COMMENT '语言',
                                `film_length` INT(5) DEFAULT NULL COMMENT '片长',


                                `tpp_num` BIGINT(20) DEFAULT NULL COMMENT '淘票票想看数',
                                `my_num` BIGINT(20) DEFAULT NULL COMMENT '猫眼想看数',
                                `db_num` BIGINT(20) DEFAULT NULL COMMENT '豆瓣想看数',
                                `mt_num` BIGINT(20) DEFAULT NULL COMMENT '美团想看数',

                                `acc_release_days` INT(8) DEFAULT NULL COMMENT '累计上映天数',
                                `release_venne` BIGINT(20) DEFAULT NULL COMMENT '上映场次',
                                `release_num_people` BIGINT(20) DEFAULT NULL COMMENT '上映人次',
                                `occupancy_rate` BIGINT(20) DEFAULT NULL COMMENT '上座率',
                                `total_box_office` BIGINT(20) DEFAULT NULL COMMENT '总票房',
                                `box_office` BIGINT(20) DEFAULT NULL COMMENT '当前票房',
                                `diany_box_office` BIGINT(20) DEFAULT NULL COMMENT '点映票房',
                                `first_daye_box_office` BIGINT(20) DEFAULT NULL COMMENT '首日票房',
                                `first_week_box_office` BIGINT(20) DEFAULT NULL COMMENT '首周票房',
                                `two_week_box_office` BIGINT(20) DEFAULT NULL COMMENT '第二周票房',
                                `word_of_mouth` BIGINT(20) DEFAULT NULL COMMENT '口碑',
                                `score` FLOAT(20) DEFAULT NULL COMMENT '评分',

                                `create_by` BIGINT(20) DEFAULT NULL COMMENT '创建者',
                                `update_by` BIGINT(20) DEFAULT NULL COMMENT '更新者',
                                `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                                `updated_at` DATETIME(3) DEFAULT NULL COMMENT '最后更新时间',
                                `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                                PRIMARY KEY (`id`)
) ENGINE = INNODB AUTO_INCREMENT = 1 COMMENT = '电影信息';

DROP TABLE ev_area_film_info;

CREATE TABLE `ev_area_film_info` (
                                     `id` INT(10) NOT NULL AUTO_INCREMENT,
                                     `film_id` BIGINT(20) DEFAULT NULL COMMENT '电影id',
                                     `level` VARCHAR(50) DEFAULT NULL COMMENT '地区级别 1:一线城市 2：二线城市 3：三线城市 4：其他城市 5：北美 6：欧美 7：东亚 8：日韩 9：非洲' ,
                                     `num_people` BIGINT(20) DEFAULT NULL COMMENT '观看人次',

                                     `release_time` DATETIME(3) DEFAULT NULL COMMENT '上映时间',
                                     `acc_release_days` INT(8) DEFAULT NULL COMMENT '累计上映天数',
                                     `release_venne` BIGINT(20) DEFAULT NULL COMMENT '上映场次',
                                     `release_num_people` BIGINT(20) DEFAULT NULL COMMENT '上映人次',
                                     `occupancy_rate` BIGINT(20) DEFAULT NULL COMMENT '上座率',
                                     `box_office` BIGINT(20) DEFAULT NULL COMMENT '票房',
                                     `diany_box_office` BIGINT(20) DEFAULT NULL COMMENT '点映票房',
                                     `first_daye_box_office` BIGINT(20) DEFAULT NULL COMMENT '首日票房',
                                     `first_week_box_office` BIGINT(20) DEFAULT NULL COMMENT '首周票房',
                                     `two_week_box_office` BIGINT(20) DEFAULT NULL COMMENT '第二周票房',
                                     `word_of_mouth` BIGINT(20) DEFAULT NULL COMMENT '口碑',
                                     `score` FLOAT(20) DEFAULT NULL COMMENT '评分',

                                     `create_by` BIGINT(20) DEFAULT NULL COMMENT '创建者',
                                     `update_by` BIGINT(20) DEFAULT NULL COMMENT '更新者',
                                     `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                                     `updated_at` DATETIME(3) DEFAULT NULL COMMENT '最后更新时间',
                                     `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`)
) ENGINE = INNODB AUTO_INCREMENT = 1 COMMENT = '地区-影片票房信息';

DROP TABLE ev_film_score;

CREATE TABLE `ev_film_score` (
                                 `id` INT(10) NOT NULL AUTO_INCREMENT,
                                 `film_id` BIGINT(20) DEFAULT NULL COMMENT '电影id',
                                 `user_id` BIGINT(20) DEFAULT NULL COMMENT '用户id',
                                 `score` INT(2) DEFAULT NULL COMMENT '评分',
                                 `re_value` FLOAT(20) DEFAULT NULL COMMENT '推荐值',
                                 `create_by` BIGINT(20) DEFAULT NULL COMMENT '创建者',
                                 `update_by` BIGINT(20) DEFAULT NULL COMMENT '更新者',
                                 `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                                 `updated_at` DATETIME(3) DEFAULT NULL COMMENT '最后更新时间',
                                 `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                                 PRIMARY KEY (`id`)
) ENGINE = INNODB AUTO_INCREMENT = 1 COMMENT = '电影评分表';
