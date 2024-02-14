# ************************************************************
# Sequel Ace SQL dump
# Version 20062
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: localhost (MySQL 11.2.2-MariaDB)
# Database: BSMG
# Generation Time: 2024-02-14 12:55:18 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table bsmg_attr1_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_attr1_infos`;

CREATE TABLE `bsmg_attr1_infos` (
  `attr1_idx` int(11) NOT NULL,
  `attr1_category` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`attr1_idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_attr1_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_attr1_infos` DISABLE KEYS */;

INSERT INTO `bsmg_attr1_infos` (`attr1_idx`, `attr1_category`)
VALUES
	(1,'솔루션'),
	(2,'제품');

/*!40000 ALTER TABLE `bsmg_attr1_infos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table bsmg_attr2_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_attr2_infos`;

CREATE TABLE `bsmg_attr2_infos` (
  `attr2_idx` int(11) NOT NULL,
  `attr1_idx` int(11) DEFAULT NULL,
  `attr2_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`attr2_idx`),
  KEY `bsmg_attr2_infos_attr1_idx_bsmg_attr1_infos_attr1_idx_foreign` (`attr1_idx`),
  CONSTRAINT `bsmg_attr2_infos_attr1_idx_bsmg_attr1_infos_attr1_idx_foreign` FOREIGN KEY (`attr1_idx`) REFERENCES `bsmg_attr1_infos` (`attr1_idx`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_attr2_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_attr2_infos` DISABLE KEYS */;

INSERT INTO `bsmg_attr2_infos` (`attr2_idx`, `attr1_idx`, `attr2_name`)
VALUES
	(1,1,'출입통제 솔루션'),
	(2,1,'발열감지 솔루션'),
	(3,1,'근태관리 솔루션'),
	(4,1,'식수관리 솔루션'),
	(5,1,'생체인증형 음주측정 솔루션'),
	(6,1,'비대면 방문자 및 행사관리 솔루션'),
	(7,1,'모바일 출입카드 시스템'),
	(8,1,'서버기반 생체인증 솔루션'),
	(9,2,'얼굴인식 장치'),
	(10,2,'홍채인식 장치'),
	(11,2,'지문인식 장치'),
	(12,2,'카드인식 장치'),
	(13,2,'라이브 스캐너'),
	(14,2,'지문 스캐너'),
	(15,2,'도장 스캐너'),
	(16,2,'지문인식 모듈'),
	(17,2,'컨트롤러'),
	(18,2,'발열감지 모듈'),
	(19,2,'단종 제품');

/*!40000 ALTER TABLE `bsmg_attr2_infos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table bsmg_include_name_reports
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_include_name_reports`;

CREATE TABLE `bsmg_include_name_reports` (
  `rpt_idx` int(11) NOT NULL,
  `rpt_reporter` varchar(20) DEFAULT NULL,
  `rpt_date` varchar(30) DEFAULT NULL,
  `rpt_to_rpt` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_ref` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_title` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_content` text DEFAULT NULL,
  `rpt_attr1` int(11) DEFAULT NULL,
  `rpt_attr2` int(11) DEFAULT NULL,
  `rpt_etc` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_confirm` tinyint(1) DEFAULT NULL,
  `reporter_name` varchar(255) DEFAULT NULL,
  `to_rpt_name` varchar(255) DEFAULT NULL,
  `ref_name` varchar(255) DEFAULT NULL,
  `rpt_reporter_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`rpt_idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table bsmg_include_name_week_reports
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_include_name_week_reports`;

CREATE TABLE `bsmg_include_name_week_reports` (
  `w_rpt_idx` int(11) NOT NULL,
  `w_rpt_reporter` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_date` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_to_rpt` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_title` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_content` text DEFAULT NULL,
  `w_rpt_part` int(11) DEFAULT NULL,
  `w_rpt_omission_date` varchar(50) DEFAULT NULL,
  `reporter_name` varchar(255) DEFAULT NULL,
  `to_rpt_name` varchar(255) DEFAULT NULL,
  `w_rpt_reporter_name` varchar(255) DEFAULT NULL,
  `w_rpt_to_rpt_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`w_rpt_idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table bsmg_member_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_member_infos`;

CREATE TABLE `bsmg_member_infos` (
  `mem_idx` int(11) NOT NULL,
  `mem_id` varchar(20) DEFAULT NULL,
  `mem_password` varchar(255) DEFAULT NULL,
  `mem_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `mem_rank` int(11) DEFAULT NULL,
  `mem_part` int(11) DEFAULT NULL,
  PRIMARY KEY (`mem_idx`),
  KEY `bsmg_member_infos_mem_rank_bsmg_rank_infos_rank_idx_foreign` (`mem_rank`),
  KEY `bsmg_member_infos_mem_part_bsmg_part_infos_part_idx_foreign` (`mem_part`),
  CONSTRAINT `bsmg_member_infos_mem_part_bsmg_part_infos_part_idx_foreign` FOREIGN KEY (`mem_part`) REFERENCES `bsmg_part_infos` (`part_idx`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `bsmg_member_infos_mem_rank_bsmg_rank_infos_rank_idx_foreign` FOREIGN KEY (`mem_rank`) REFERENCES `bsmg_rank_infos` (`rank_idx`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_member_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_member_infos` DISABLE KEYS */;

INSERT INTO `bsmg_member_infos` (`mem_idx`, `mem_id`, `mem_password`, `mem_name`, `mem_rank`, `mem_part`)
VALUES
	(1,'admin','$argon2id$v=19$m=65536,t=3,p=2$S+PKPEBoknDph6reKhBz8w$aefjZakGhtqlk40pQdrdK3MynfVKneX2DkaTIXf36uM','테스트유저',1,1),
	(2,'JJaturi','$argon2id$v=19$m=65536,t=3,p=2$K7Bf0ISZ8ebikFkGCah8GA$skESTtS0IpBfTplmLcgF5smx0gNmhjiSy17wT6MZEyg','짜투리팀장',3,11),
	(3,'부소장','$argon2id$v=19$m=65536,t=3,p=2$wtxM4QqyREsn0zbVM3zPPA$Et6Rb3J1zKBJR+QBx/v6lnX9EReHwMM+TmaCphUlVm4','부소장',2,1),
	(4,'팀장1','$argon2id$v=19$m=65536,t=3,p=2$8zg2+t+ngKnPQm+wYsehcw$S+PlRR/xpaaJfQhpEbVLpVOBkVlIulWkzpd3xtjHHWY','광학기구팀장',3,10),
	(5,'팀장2','$argon2id$v=19$m=65536,t=3,p=2$JhqzF5gZRrbb1PfOb0KQYw$lgvYc/BCpBDC5xPaLkbz7iHz2nXhGjU4Jtb/JtMvxDE','디자인팀장',3,9),
	(6,'김팀장2','$argon2id$v=19$m=65536,t=3,p=2$1xq9jlZ8worN3SiDoK0Now$Ijjsz1RSX2ovqulK7o+GqVBGjOZGCbtoJFK6bY0RQS4','SW1팀장',3,2),
	(7,'한팀장','$argon2id$v=19$m=65536,t=3,p=2$kaf0qVbRCwX4qD+51/KdnA$MVnFn6q8GvIoekJp7ou4QkLCxK/lM7y3uVuowhyght0','FW2팀장',3,5),
	(8,'고팀장','$argon2id$v=19$m=65536,t=3,p=2$SEIfTUJgxQ5RbrmxRuNrrg$DFv735Bl50p7ps26a+311nqq5yOpTv3tg1zYPo0YuGw','HW1팀장',3,6),
	(9,'정팀장','$argon2id$v=19$m=65536,t=3,p=2$iNJOB3NUw72SbwHcLZcIhA$+1BofqZfDoy3NzB4NZsx4Joy3VHAGUDXKGW26jg5qXc','HW2팀장',3,7),
	(10,'한팀장2','$argon2id$v=19$m=65536,t=3,p=2$5LMMqtGHXTqZmP4cc/+4ZQ$wg/HF/ZjbvFY4qBN73aCi3qjjcsR9sON5cZgz4swH3A','Mobile팀장',3,8),
	(11,'문주영','$argon2id$v=19$m=65536,t=3,p=2$1ZPaJWjGTPmuY3ZTWui9TQ$UlV6Q4XH/36ezxQpsw5f0+ea/0WzLbbHwalBEjtPEUU','문주영이름',4,3),
	(12,'김지환','$argon2id$v=19$m=65536,t=3,p=2$kscIvSQMZN6Va8vQsxQB7w$RjLuVdU8N2H4dy6v4tMSLprYlcEuquRlUOSQuQHOXjU','김지환',4,2),
	(13,'최윤지','$argon2id$v=19$m=65536,t=3,p=2$rxNMnAm4ypzVvA8rRqnwig$V92C0X4rxx/wVhCDOUBZ/ms6x96hR3i93RZMmtRLPd0','최윤지',4,5),
	(14,'김서경','$argon2id$v=19$m=65536,t=3,p=2$GbTtiiEtw0WLkdZ3LYAusw$ABjpfXQHgBAkojSakryqC0pLm2iI3HHfzXxMrTp/up8','김서경',4,4),
	(15,'박소은','$argon2id$v=19$m=65536,t=3,p=2$ZcO8ifI9udKLt+6k0yLCSg$GCDvEzfd7nV8aH0PU1pR++TPcF2LBVE5EymHSFjPTdk','박소은',4,3),
	(16,'김근우','$argon2id$v=19$m=65536,t=3,p=2$oxjvVpNDFpgDhZ5NzvYVWA$RiH9/PaR0QHkSv9PtgYKORn1OFrbOP9R3NlH59e81Ak','김근우',4,6),
	(17,'김진우','$argon2id$v=19$m=65536,t=3,p=2$OMP/foG9icDKPnEmkGeZew$LBaYSFJoPAi9a40Nfs3bLbKUdV3kRuTFmNZzlTRUHAc','김진우이름',4,7),
	(18,'박준언','$argon2id$v=19$m=65536,t=3,p=2$e2y6nKrHUwqTUviKDvW2MA$peOuvZ9oMati75wslUvARDlAg03844qV31BZ4KGS7bA','박준언이름',4,3),
	(19,'오팀장','$argon2id$v=19$m=65536,t=3,p=2$AZSKyEINGZ3ZJszRu5fqcw$Ik4Ky4NJrHd6bKF4vS0bl9WBI+JmkZAm99aj/24rPyk','SW2팀장',3,3),
	(20,'fw1team','$argon2id$v=19$m=65536,t=3,p=2$zcMGUnC0Dgto5S5+adMjzQ$yNJSS9QH6y6+e5MZKsGE0J/ouEKbyngb7r7xZDe0SKY','FW1팀장',3,4),
	(21,'argon2','$argon2id$v=19$m=65536,t=3,p=2$ORmcUpqZoxYWfiIOUyd31Q$cR07bbO5mwrEphMFtP1tYThQoFwucitoSwoox9ffVk0','암호화사용자',1,2),
	(22,'d','$argon2id$v=19$m=65536,t=3,p=2$qmrg4sgrBsidAWde/OT9LQ$q+W6KAHfunTolCQ/6GSilQYb89CZp9LQQxND0OTQLUA','d',2,2);

/*!40000 ALTER TABLE `bsmg_member_infos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table bsmg_part_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_part_infos`;

CREATE TABLE `bsmg_part_infos` (
  `part_idx` int(11) NOT NULL,
  `part_name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`part_idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_part_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_part_infos` DISABLE KEYS */;

INSERT INTO `bsmg_part_infos` (`part_idx`, `part_name`)
VALUES
	(1,'연구소'),
	(2,'SW1팀'),
	(3,'SW2팀'),
	(4,'FW1팀'),
	(5,'FW2팀'),
	(6,'HW1팀'),
	(7,'HW2팀'),
	(8,'Mobile팀'),
	(9,'디자인팀'),
	(10,'광학기구팀'),
	(11,'연구관리팀');

/*!40000 ALTER TABLE `bsmg_part_infos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table bsmg_rank_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_rank_infos`;

CREATE TABLE `bsmg_rank_infos` (
  `rank_idx` int(11) NOT NULL,
  `rank_name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`rank_idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_rank_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_rank_infos` DISABLE KEYS */;

INSERT INTO `bsmg_rank_infos` (`rank_idx`, `rank_name`)
VALUES
	(1,'연구소장'),
	(2,'부소장'),
	(3,'팀장'),
	(4,'Pro');

/*!40000 ALTER TABLE `bsmg_rank_infos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table bsmg_report_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_report_infos`;

CREATE TABLE `bsmg_report_infos` (
  `rpt_idx` int(11) NOT NULL AUTO_INCREMENT,
  `rpt_reporter` varchar(20) DEFAULT NULL,
  `rpt_date` varchar(30) DEFAULT NULL,
  `rpt_to_rpt` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_ref` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_title` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_content` text DEFAULT NULL,
  `rpt_attr1` int(11) DEFAULT NULL,
  `rpt_attr2` int(11) DEFAULT NULL,
  `rpt_etc` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `rpt_confirm` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`rpt_idx`),
  KEY `bsmg_report_infos_rpt_attr1_bsmg_attr1_infos_attr1_idx_foreign` (`rpt_attr1`),
  KEY `bsmg_report_infos_rpt_attr2_bsmg_attr2_infos_attr2_idx_foreign` (`rpt_attr2`),
  CONSTRAINT `bsmg_report_infos_rpt_attr1_bsmg_attr1_infos_attr1_idx_foreign` FOREIGN KEY (`rpt_attr1`) REFERENCES `bsmg_attr1_infos` (`attr1_idx`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `bsmg_report_infos_rpt_attr2_bsmg_attr2_infos_attr2_idx_foreign` FOREIGN KEY (`rpt_attr2`) REFERENCES `bsmg_attr2_infos` (`attr2_idx`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_report_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_report_infos` DISABLE KEYS */;

INSERT INTO `bsmg_report_infos` (`rpt_idx`, `rpt_reporter`, `rpt_date`, `rpt_to_rpt`, `rpt_ref`, `rpt_title`, `rpt_content`, `rpt_attr1`, `rpt_attr2`, `rpt_etc`, `rpt_confirm`)
VALUES
	(0,'문주영','20231224024452','admin','박소은','2023년 12월 24일 뀨뀨 일일 업무보고','123\n\ndho\n\n안ㅂ수자',1,2,'123123',0),
	(1,'문주영','20231224035109','부소장','박준언','2023년 12월 24일 뀨뀨 일일 업무보고','123213asdas\nasd\nasd\nasdasdd',2,11,'2313',0),
	(23,'문주영','20231224044033','부소장','박준언','2023년 12월 24일 뀨뀨 일일 업무보고','123213',1,3,'123123',1),
	(25,'문주영','20231224044515','한팀장2','김서경','2023년 12월 24일 뀨뀨 일일 업무보고','하이',1,3,'바이',0),
	(26,'문주영','20231224045146','김팀장2','박준언','2023년 12월 24일 뀨뀨 일일 업무보고','고쳐야합니다\n\n그래야 삽니다',2,11,'',0),
	(27,'문주영','20231224045503','박준언','고팀장','2023년 12월 24일 뀨뀨 일일 업무보고','123',2,10,'',0),
	(29,'부소장','20231225011321','부소장','박소은','2023년 12월 25일 부소장 일일 업무보고','보고 확인테스트',2,11,'ㄴㄴ',0),
	(31,'admin','20231226002809','부소장','김서경','2023년 12월 26일 테스트유저 일일 업무보고','11',2,10,'',0),
	(32,'admin','20231226022929','문주영','박준언이름','2023년 12월 26일 테스트유저 일일 업무보고','qwe',2,12,'',0),
	(34,'박준언','20231227234919','부소장','오팀장','2023년 12월 27일 박준언이름 일일 업무보고','ㅜㅏ',2,11,'',1),
	(35,'문주영','20231228232144','오팀장','테스트유저,부소장,짜투리팀장','2023년 12월 28일 문주영이름 일일 업무보고','1. BSMG 리팩토링\n- WeekReport API 구현\n- SingleTon 구현 (클라이언트단에서 직급, 부서, 업무속성 저장하는 DataManager 구현)\n\nTODO : \n- 사용한 디자인 패턴 정리 및 \n- WeekReport 생성을 Cron을 통한 자동화\n- 코드 모듈화\n- Session -> JWT\n- 사용자 정보 암호화\n- DB 전달을 채널링을 통해 구현 (고민중)\n\n',1,3,'echo를 통한 리팩토링이 너무 길었음..',0),
	(36,'문주영','20240103003202','오팀장','테스트유저,부소장,박준언이름','2024년 01월 03일 문주영이름 일일 업무보고','1. BSMG 리팩토링\n- 사용자 정보 암호화 (argon2)\n\nTODO : \n- 사용한 디자인 패턴 정리 및 \n- WeekReport 생성을 Cron을 통한 자동화\n- 코드 모듈화\n- Session -> JWT\n- DB 전달을 채널링을 통해 구현 (고민중)\n\n',1,3,'하고싶은건 많은데 그래도 차근차근 해나가야함',0),
	(37,'문주영','20240105234348','오팀장','테스트유저,부소장','2024년 01월 05일 문주영이름 일일 업무보고','1. BSMG 리팩토링\n- JWT 사용하여 로그인 및 인가처리 개발\n- Config.json 파일을 통해 서버 설정 제어하도록 개발\n- Cron을 이용하여 주간보고 자동화 개발\n\nTODO : \n- 사용한 디자인 패턴 정리 및 \n- 코드 모듈화\n\n\n',1,3,'',0),
	(38,'admin','20240118223502','김지환','문주영이름,박준언이름','2024년 01월 18일 테스트유저 일일 업무보고','테스트',1,3,'',0),
	(39,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);

/*!40000 ALTER TABLE `bsmg_report_infos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table bsmg_schedule_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_schedule_infos`;

CREATE TABLE `bsmg_schedule_infos` (
  `rpt_idx` int(11) DEFAULT NULL,
  `sc_content` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_schedule_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_schedule_infos` DISABLE KEYS */;

INSERT INTO `bsmg_schedule_infos` (`rpt_idx`, `sc_content`)
VALUES
	(0,'33'),
	(29,'일정'),
	(0,'엠시트'),
	(35,'업무보고 암호화'),
	(35,'업무보고 jwt'),
	(36,'bsmg 리팩토링'),
	(36,'디자인패턴 공부'),
	(36,'동시성,병렬성 개념 확실히'),
	(31,'스케쥴 추가'),
	(37,'배포');

/*!40000 ALTER TABLE `bsmg_schedule_infos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table bsmg_week_rpt_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_week_rpt_infos`;

CREATE TABLE `bsmg_week_rpt_infos` (
  `w_rpt_idx` int(11) NOT NULL AUTO_INCREMENT,
  `w_rpt_reporter` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_date` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_to_rpt` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_title` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `w_rpt_content` text DEFAULT NULL,
  `w_rpt_part` int(11) DEFAULT NULL,
  `w_rpt_omission_date` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`w_rpt_idx`),
  KEY `bsmg_week_rpt_infos_w_rpt_part_bsmg_part_infos_part_idx_foreign` (`w_rpt_part`),
  CONSTRAINT `bsmg_week_rpt_infos_w_rpt_part_bsmg_part_infos_part_idx_foreign` FOREIGN KEY (`w_rpt_part`) REFERENCES `bsmg_part_infos` (`part_idx`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `bsmg_week_rpt_infos` WRITE;
/*!40000 ALTER TABLE `bsmg_week_rpt_infos` DISABLE KEYS */;

INSERT INTO `bsmg_week_rpt_infos` (`w_rpt_idx`, `w_rpt_reporter`, `w_rpt_date`, `w_rpt_to_rpt`, `w_rpt_title`, `w_rpt_content`, `w_rpt_part`, `w_rpt_omission_date`)
VALUES
	(9,'부소장','20231231000000','JJaturi','12월 5주차 부소장 주간 업무보고','📆20231225011321\n보고 확인테스트\n\n\n수정이 되었을까요',1,'20231226, 20231227, 20231228, 20231229'),
	(10,'문주영','20231231000000','오팀장','12월 5주차 문주영이름 주간 업무보고','📆20231224024452\n123\n\ndho\n\n안ㅂ수자\n📆20231224035109\n123213asdas\nasd\nasd\nasdasdd\n📆20231224044033\n123213\n📆20231224044515\n하이\n📆20231224045146\n고쳐야합니다\n\n그래야 삽니다\n📆20231224045503\n123\n',3,'20231225, 20231226, 20231227, 20231228, 20231229'),
	(11,'박준언','20231231000000','오팀장','12월 5주차 박준언이름 주간 업무보고','📆20231227234919\nㅜㅏ\n',3,'20231225, 20231226, 20231228, 20231229'),
	(14,'문주영','20240109000000','오팀장','01월 1주차 문주영이름 주간 업무보고','📆20240103\n1. BSMG 리팩토링\n- 사용자 정보 암호화 (argon2)\n\nTODO : \n- 사용한 디자인 패턴 정리 및 \n- WeekReport 생성을 Cron을 통한 자동화\n- 코드 모듈화\n- Session -> JWT\n- DB 전달을 채널링을 통해 구현 (고민중)\n\n\n📆20240105\n1. BSMG 리팩토링\n- JWT 사용하여 로그인 및 인가처리 개발\n- Config.json 파일을 통해 서버 설정 제어하도록 개발\n- Cron을 이용하여 주간보고 자동화 개발\n\nTODO : \n- 사용한 디자인 패턴 정리 및 \n- 코드 모듈화\n\n\n\n',3,'20240102, 20240104, 20240108');

/*!40000 ALTER TABLE `bsmg_week_rpt_infos` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
