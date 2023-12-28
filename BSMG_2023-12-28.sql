# ************************************************************
# Sequel Ace SQL dump
# Version 20062
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: localhost (MySQL 11.2.2-MariaDB)
# Database: BSMG
# Generation Time: 2023-12-28 14:39:24 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table bsmg_member_infos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bsmg_member_infos`;

CREATE TABLE `bsmg_member_infos` (
  `mem_idx` int(11) NOT NULL,
  `mem_id` varchar(20) DEFAULT NULL,
  `mem_password` varchar(50) DEFAULT NULL,
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
	(1,'admin','admin','테스트유저',1,1),
	(2,'JJaturi','0000','짜투리팀장',3,11),
	(3,'부소장','0000','부소장',2,1),
	(4,'팀장1','0000','광학기구팀장',3,10),
	(5,'팀장2','0000','디자인팀장',3,9),
	(6,'김팀장2','0000','SW1팀장',3,2),
	(7,'한팀장','0000','FW2팀장',3,5),
	(8,'고팀장','0000','HW1팀장',3,6),
	(9,'정팀장','0000','HW2팀장',3,7),
	(10,'한팀장2','0000','Mobile팀장',3,8),
	(11,'문주영','0000','문주영이름',4,3),
	(12,'김지환','0000','김지환',4,2),
	(13,'최윤지','0000','최윤지',4,5),
	(14,'김서경','0000','김서경',4,4),
	(15,'박소은','0000','박소은',4,3),
	(16,'김근우','0000','김근우',4,6),
	(17,'김진우','0000','김진우이름',4,7),
	(18,'박준언','0000','박준언이름',4,3),
	(19,'오팀장','0000','SW2팀장',3,3),
	(20,'fw1team','0000','FW1팀장',3,4);

/*!40000 ALTER TABLE `bsmg_member_infos` ENABLE KEYS */;
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
	(0,'문주영','20231224024452','테스트유저','박소은','2023년 12월 24일 뀨뀨 일일 업무보고','123\n\ndho\n\n안ㅂ수자',1,2,'123123',0),
	(1,'문주영','20231224035109','부소장','박준언','2023년 12월 24일 뀨뀨 일일 업무보고','123213asdas\nasd\nasd\nasdasdd',2,11,'2313',0),
	(23,'문주영','20231224044033','부소장','박준언','2023년 12월 24일 뀨뀨 일일 업무보고','123213',1,3,'123123',1),
	(25,'문주영','20231224044515','연구소장','김서경','2023년 12월 24일 뀨뀨 일일 업무보고','하이',1,3,'바이',0),
	(26,'문주영','20231224045146','김팀장님','박준언','2023년 12월 24일 뀨뀨 일일 업무보고','고쳐야합니다\n\n그래야 삽니다',2,11,'',0),
	(27,'문주영','20231224045503','박준언','고팀장','2023년 12월 24일 뀨뀨 일일 업무보고','123',2,10,'',0),
	(29,'부소장','20231225011321','부소장','박소은','2023년 12월 25일 부소장 일일 업무보고','보고 확인테스트',2,11,'ㄴㄴ',0),
	(31,'admin','20231226002809','부소장','김서경','2023년 12월 26일 테스트유저 일일 업무보고','11',2,10,'',0),
	(32,'admin','20231226022929','문주영이름','박준언이름','2023년 12월 26일 테스트유저 일일 업무보고','qwe',2,12,'',0),
	(34,'박준언','20231227234919','부소장','오팀장','2023년 12월 27일 박준언이름 일일 업무보고','ㅜㅏ',2,11,'',0),
	(35,'문주영','20231228232144','SW2팀장','테스트유저,부소장,짜투리팀장','2023년 12월 28일 문주영이름 일일 업무보고','1. BSMG 리팩토링\n- WeekReport API 구현\n- SingleTon 구현 (클라이언트단에서 직급, 부서, 업무속성 저장하는 DataManager 구현)\n\nTODO : \n- 사용한 디자인 패턴 정리 및 \n- WeekReport 생성을 Cron을 통한 자동화\n- 코드 모듈화\n- Session -> JWT\n- 사용자 정보 암호화\n- DB 전달을 채널링을 통해 구현 (고민중)\n\n',1,3,'echo를 통한 리팩토링이 너무 길었음..',0);

/*!40000 ALTER TABLE `bsmg_report_infos` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
