
CREATE TABLE `survey_info_qd` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `survey_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'survey_id',
  `survey_info` varchar(10240) NOT NULL DEFAULT '' COMMENT '问卷内容信息',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='快决策平台问卷信息';


CREATE TABLE `survey_user_record_qd` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT 'uuid',
  `terminal_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'terminal-id',
  `identity_id` varchar(128) NOT NULL DEFAULT '' COMMENT '身份标识',
  `survey_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'survey_id',
  `status_data` varchar(4096) NOT NULL DEFAULT '' COMMENT '答题状态数据',
  `answer_data` varchar(10240) NOT NULL DEFAULT '' COMMENT '答题内容数据',
  `sync_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '同步时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sid` (`survey_id`,`identity_id`) ,
  KEY `idx_hash` (`identity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='快决策平台问卷数据';


CREATE TABLE `user_identity_hash` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `hash` varchar(128) NOT NULL DEFAULT '' COMMENT 'uuid的hash值',
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT 'uuid',
  `terminal_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'terminal-id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_hash` (`hash`) ,
  KEY `idx_uuid` (`uuid`) ,
  KEY `idx_tid` (`terminal_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户标识 hash表';
