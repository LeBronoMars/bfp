-- MySQL dump 10.13  Distrib 5.7.11, for osx10.10 (x86_64)
--
-- Host: localhost    Database: bfp
-- ------------------------------------------------------
-- Server version	5.7.11

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `fire_stations`
--

DROP TABLE IF EXISTS `fire_stations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fire_stations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `station_name` varchar(255) DEFAULT NULL,
  `station_code` varchar(255) DEFAULT NULL,
  `latitude` double DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `contact_no` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fire_stations`
--

LOCK TABLES `fire_stations` WRITE;
/*!40000 ALTER TABLE `fire_stations` DISABLE KEYS */;
INSERT INTO `fire_stations` VALUES (1,'2016-06-12 01:33:42','2016-06-12 01:40:14','Caloocan City Fire Station','caloocan001',14.72035,121.004435,'Caloocan City, Metro Manila','632 324-6527','Not available');
/*!40000 ALTER TABLE `fire_stations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fire_statuses`
--

DROP TABLE IF EXISTS `fire_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fire_statuses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `incident_id` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `reported_by` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fire_statuses`
--

LOCK TABLES `fire_statuses` WRITE;
/*!40000 ALTER TABLE `fire_statuses` DISABLE KEYS */;
INSERT INTO `fire_statuses` VALUES (18,'2016-06-09 11:52:50','2016-06-09 11:52:50',6,'1st alarm',1),(19,'2016-06-12 02:11:58','2016-06-12 02:11:58',8,'1st Level',1);
/*!40000 ALTER TABLE `fire_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `incidents`
--

DROP TABLE IF EXISTS `incidents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `incidents` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `latitude` double DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `remarks` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incidents`
--

LOCK TABLES `incidents` WRITE;
/*!40000 ALTER TABLE `incidents` DISABLE KEYS */;
INSERT INTO `incidents` VALUES (6,'2016-06-09 11:52:50','2016-06-09 11:52:50',14.592120170593262,121.05896759033203,'Robinsons Galleria',''),(8,'2016-06-12 02:11:58','2016-06-12 02:11:58',18.4655394,-66.1057355,'San Juan, Puerto Rico','1st Level');
/*!40000 ALTER TABLE `incidents` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `qry_incidents`
--

DROP TABLE IF EXISTS `qry_incidents`;
/*!50001 DROP VIEW IF EXISTS `qry_incidents`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `qry_incidents` AS SELECT 
 1 AS `incident_id`,
 1 AS `date_reported`,
 1 AS `latitude`,
 1 AS `longitude`,
 1 AS `address`,
 1 AS `remarks`,
 1 AS `fire_status_id`,
 1 AS `fire_status_reported`,
 1 AS `fire_status`,
 1 AS `reporter_id`,
 1 AS `reporter_first_name`,
 1 AS `reporter_last_name`,
 1 AS `reporter_role`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `contact_no` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `userrole` varchar(255) DEFAULT NULL,
  `userlevel` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `is_password_default` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2016-06-12 04:32:08','2016-06-12 04:32:08','Ned','Flanders','nedflanders@gmail.com','09274502976','active','admin','NHR','24HwWlvRAbv2Aq-GaCS-pdQVTQ==',1),(2,'2016-06-14 12:33:24','2016-06-14 12:33:24','Russel','Bulanon','rsbulanon@gmail.com','09321622825','active','admin','NHR','DRREzCkmKH7vvcRKhzKscwI3KA==',1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `qry_incidents`
--

/*!50001 DROP VIEW IF EXISTS `qry_incidents`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `qry_incidents` AS select `i`.`id` AS `incident_id`,`i`.`created_at` AS `date_reported`,`i`.`latitude` AS `latitude`,`i`.`longitude` AS `longitude`,`i`.`address` AS `address`,`i`.`remarks` AS `remarks`,`f`.`id` AS `fire_status_id`,`f`.`created_at` AS `fire_status_reported`,`f`.`status` AS `fire_status`,`u`.`id` AS `reporter_id`,`u`.`first_name` AS `reporter_first_name`,`u`.`last_name` AS `reporter_last_name`,`u`.`userrole` AS `reporter_role` from ((`fire_statuses` `f` join `incidents` `i` on((`i`.`id` = `f`.`incident_id`))) join `users` `u` on((`f`.`reported_by` = `u`.`id`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-06-14 20:46:46
