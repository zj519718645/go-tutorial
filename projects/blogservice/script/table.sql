CREATE TABLE `blog_service`.`blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '创建人',
    `created_on` int(10) unsigned DEFAULT 0 COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT 0 COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除0为未删除、1为已删除',
    `state` tinyint(3) unsigned DEFAULT 1 COMMENT '状态0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '标签管理';

CREATE TABLE `blog_service`.`blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `description` varchar(255) DEFAULT '' COMMENT '文章简介',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
    `content` longtext COMMENT '文章内容',
    `created_on` int(10) unsigned DEFAULT 0 COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT 0 COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除0为未删除、1为已删除',
    `state` tinyint(3) unsigned DEFAULT 1 COMMENT '状态0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '文章管理';

CREATE TABLE `blog_service`.`blog_article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(10) unsigned NOT NULL COMMENT '文章ID',
    `tag_id` int(10) unsigned NOT NULL COMMENT '标签ID',
    `created_on` int(10) unsigned DEFAULT 0 COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT 0 COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除0为未删除、1为已删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '文章标签关联';