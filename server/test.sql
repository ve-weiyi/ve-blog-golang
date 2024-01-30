CREATE TABLE `qd_survey_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT 'uuid',
  `terminal-id` varchar(128) NOT NULL DEFAULT '' COMMENT 'terminal-id',
  `hash` varchar(128) NOT NULL DEFAULT '' COMMENT 'hash',
  `survey_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'survey_id',
  `source_data` varchar(10240) NOT NULL DEFAULT '' COMMENT '源数据',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_hash` (`hash`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='快决策平台问卷数据';

CREATE TABLE `user_uuid_hash` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT 'uuid',
  `hash` varchar(128) NOT NULL DEFAULT '' COMMENT 'uuid的hash值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户uuid hash表';
