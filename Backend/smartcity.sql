-- MySQL dump 10.13  Distrib 5.7.44, for Win64 (x86_64)
--
-- Host: localhost    Database: smartcity
-- ------------------------------------------------------
-- Server version	5.7.44-log

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
-- Table structure for table `accident_history`
--

DROP TABLE IF EXISTS `accident_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `accident_history` (
  `V_ID` int(11) NOT NULL,
  `I_ID` int(11) NOT NULL,
  PRIMARY KEY (`V_ID`,`I_ID`),
  KEY `I_ID` (`I_ID`),
  CONSTRAINT `accident_history_ibfk_1` FOREIGN KEY (`V_ID`) REFERENCES `vehicle` (`VEHICLE_ID`),
  CONSTRAINT `accident_history_ibfk_2` FOREIGN KEY (`I_ID`) REFERENCES `incident` (`INCIDENT_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accident_history`
--

LOCK TABLES `accident_history` WRITE;
/*!40000 ALTER TABLE `accident_history` DISABLE KEYS */;
INSERT INTO `accident_history` VALUES (1,1),(2,2),(3,3),(4,4),(5,5),(6,6),(7,7),(8,8),(9,9),(10,10),(11,11),(12,12),(13,13),(14,14),(15,15),(16,16),(17,17),(18,18),(19,19),(20,20),(21,21),(22,22),(23,23),(24,24),(25,25),(26,26),(27,27),(28,28),(29,29),(30,30),(31,31),(32,32),(33,33),(34,34),(35,35),(36,36),(37,37),(38,38),(39,39),(40,40),(41,41),(42,42),(43,43),(44,44),(45,45),(46,46),(47,47),(48,48),(49,49),(11,50),(50,50),(5,51);
/*!40000 ALTER TABLE `accident_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bookings`
--

DROP TABLE IF EXISTS `bookings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bookings` (
  `vehicle_id` int(11) DEFAULT NULL,
  `booking_count` int(11) DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookings`
--

LOCK TABLES `bookings` WRITE;
/*!40000 ALTER TABLE `bookings` DISABLE KEYS */;
INSERT INTO `bookings` VALUES (1,0),(2,1);
/*!40000 ALTER TABLE `bookings` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = cp850 */ ;
/*!50003 SET character_set_results = cp850 */ ;
/*!50003 SET collation_connection  = cp850_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER update_vehicle_status_on_booking_count
AFTER UPDATE ON bookings
FOR EACH ROW
BEGIN
    IF NEW.booking_count > 10 THEN
        
        INSERT INTO vehicle_status_change (vehicle_id, status_change)
        VALUES (NEW.vehicle_id, TRUE)
        ON DUPLICATE KEY UPDATE status_change = TRUE, last_update = NOW();
    END IF;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `human`
--

DROP TABLE IF EXISTS `human`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `human` (
  `ID_NO` int(11) NOT NULL,
  `FNAME` varchar(50) DEFAULT NULL,
  `LNAME` varchar(50) DEFAULT NULL,
  `DOB` date DEFAULT NULL,
  `AGE` int(11) DEFAULT NULL,
  `V_ID` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID_NO`),
  KEY `V_ID` (`V_ID`),
  CONSTRAINT `human_ibfk_1` FOREIGN KEY (`V_ID`) REFERENCES `vehicle` (`VEHICLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `human`
--

LOCK TABLES `human` WRITE;
/*!40000 ALTER TABLE `human` DISABLE KEYS */;
INSERT INTO `human` VALUES (0,'Suryansh','Rohil','2005-11-11',19,1),(1,'Gagan','Jyoti','1975-12-10',50,10),(2,'Alice','Smith','1990-03-22',34,2),(3,'Suryansh','Rohil','2005-11-11',19,5),(4,'Emily','Davis','1995-11-25',29,4),(5,'Michael','Brown','1987-05-30',37,5),(6,'Sarah','Miller','1993-07-12',31,6),(7,'David','Wilson','1980-09-18',44,7),(8,'Jessica','Moore','1998-12-05',26,8),(9,'Daniel','Taylor','1983-04-29',41,9),(10,'Laura','Anderson','1992-06-20',32,10),(11,'Chris','Thomas','1986-10-14',38,11),(12,'Sophia','Jackson','1997-01-07',27,12),(13,'Matthew','White','1989-02-17',35,13),(14,'Olivia','Harris','1994-08-30',30,14),(15,'James','Martin','1981-07-21',43,15),(16,'Emma','Thompson','2000-05-15',24,16),(17,'Ethan','Garcia','1991-09-05',33,17),(18,'Ava','Martinez','1984-12-10',40,18),(19,'Logan','Robinson','1999-04-28',25,19),(20,'Mia','Clark','1996-11-11',28,20),(21,'Alexander','Rodriguez','1988-03-09',36,21),(22,'Isabella','Lewis','1985-07-23',39,22),(23,'Benjamin','Walker','1990-10-01',34,23),(24,'Charlotte','Hall','2001-06-14',23,24),(25,'Henry','Allen','1983-08-02',41,25),(26,'Amelia','Young','1998-01-18',26,26),(27,'Lucas','King','1982-05-25',42,27),(28,'Harper','Wright','1995-12-31',29,28),(29,'Mason','Scott','1992-07-07',32,29),(30,'Ella','Green','1986-04-19',38,30),(31,'Elijah','Adams','1993-10-22',31,31),(32,'Scarlett','Baker','1980-03-16',44,32),(33,'Jacob','Gonzalez','1987-09-09',37,33),(34,'Grace','Nelson','1999-05-30',25,34),(35,'William','Carter','1984-11-14',40,35),(36,'Sophie','Mitchell','1997-06-08',27,36),(37,'Daniel','Perez','1991-02-05',33,37),(38,'Chloe','Roberts','1994-09-27',30,38),(39,'Jack','Turner','1989-12-21',35,39),(40,'Lily','Phillips','2002-03-11',22,40),(41,'Owen','Campbell','1985-06-04',39,41),(42,'Zoey','Parker','1996-08-16',28,42),(43,'Sebastian','Evans','1983-07-25',41,43),(44,'Hannah','Edwards','1995-01-29',29,44),(45,'Dylan','Collins','1990-11-13',34,45),(46,'Madison','Stewart','1988-04-07',36,46),(47,'Caleb','Sanchez','1992-12-17',32,47),(48,'Victoria','Morris','1993-05-03',31,48),(49,'Nathan','Rogers','1987-10-26',37,49),(50,'Eleanor','Reed','2000-07-12',24,50),(53,'Suryansh','Rohil','2005-11-11',19,52),(54,'Temp','Bro','2025-08-10',-1,10),(55,'dbms','test#1','2005-11-11',19,10),(56,'first','person','2005-11-11',19,0),(57,'PataNahi','IDK','2005-11-11',19,9);
/*!40000 ALTER TABLE `human` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `incident`
--

DROP TABLE IF EXISTS `incident`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `incident` (
  `INCIDENT_ID` int(11) NOT NULL,
  `V_ID` int(11) DEFAULT NULL,
  `DESCRIPTION` text,
  `REPORT_TIME_DATE` datetime DEFAULT NULL,
  PRIMARY KEY (`INCIDENT_ID`),
  KEY `V_ID` (`V_ID`),
  CONSTRAINT `incident_ibfk_1` FOREIGN KEY (`V_ID`) REFERENCES `vehicle` (`VEHICLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incident`
--

LOCK TABLES `incident` WRITE;
/*!40000 ALTER TABLE `incident` DISABLE KEYS */;
INSERT INTO `incident` VALUES (1,1,'Engine failure','2025-03-20 10:00:00'),(2,3,'Brake malfunction','2025-03-21 11:30:00'),(3,5,'Overheating','2025-04-13 08:47:58'),(4,2,'Transmission issue','2025-03-21 14:00:00'),(5,10,'Tire Puncture','2025-04-16 14:06:27'),(6,6,'Fuel leakage','2025-03-22 10:30:00'),(7,7,'Battery drain','2025-03-22 13:00:00'),(8,8,'Headlight malfunction','2025-03-22 18:20:00'),(9,9,'Alternator failure','2025-03-23 08:00:00'),(10,10,'Steering issue','2025-03-23 10:15:00'),(11,11,'Suspension problem','2025-03-23 14:30:00'),(12,12,'Oil leakage','2025-03-23 16:45:00'),(13,13,'Radiator failure','2025-03-23 17:00:00'),(14,14,'Exhaust system fault','2025-03-24 07:20:00'),(15,15,'Tire puncture','2025-03-24 09:10:00'),(16,16,'Clutch failure','2025-03-24 11:25:00'),(17,17,'Power window malfunction','2025-03-24 12:50:00'),(18,18,'Horn failure','2025-03-24 15:30:00'),(19,19,'Airbag fault','2025-03-24 17:40:00'),(20,20,'Dashboard warning lights','2025-03-24 20:00:00'),(21,21,'Wiper malfunction','2025-03-25 07:45:00'),(22,22,'Speedometer issue','2025-03-25 09:20:00'),(23,23,'Starter motor failure','2025-03-25 11:15:00'),(24,24,'Water pump issue','2025-03-25 12:30:00'),(25,25,'Fuel pump failure','2025-03-25 14:10:00'),(26,26,'Door lock failure','2025-03-25 16:30:00'),(27,27,'Shock absorber issue','2025-03-25 18:45:00'),(28,28,'Ignition coil failure','2025-03-25 20:00:00'),(29,29,'Seatbelt malfunction','2025-03-25 21:30:00'),(30,30,'Head gasket issue','2025-03-25 23:15:00'),(31,31,'Coolant leakage','2025-03-26 06:30:00'),(32,32,'Differential failure','2025-03-26 08:45:00'),(33,33,'Timing belt issue','2025-03-26 10:10:00'),(34,34,'Fan belt failure','2025-03-26 11:55:00'),(35,35,'Air conditioning problem','2025-03-26 13:40:00'),(36,36,'Car alarm malfunction','2025-03-26 15:00:00'),(37,37,'Windshield crack','2025-03-26 16:30:00'),(38,38,'Parking brake failure','2025-03-26 18:15:00'),(39,39,'Fog light malfunction','2025-03-26 19:45:00'),(40,40,'Reverse gear issue','2025-03-26 21:00:00'),(41,41,'Turbocharger failure','2025-03-26 22:30:00'),(42,42,'Engine misfire','2025-03-27 07:00:00'),(43,4,'GPS gaya re baaba','2025-04-17 17:01:02'),(44,44,'Catalytic converter issue','2025-03-27 10:00:00'),(45,45,'Throttle body malfunction','2025-03-27 11:15:00'),(46,46,'AC compressor issue','2025-03-27 12:45:00'),(47,47,'Car key not detected','2025-03-27 14:00:00'),(48,48,'Brake pad wear','2025-03-27 15:30:00'),(49,49,'Fuel injector issue','2025-03-27 17:00:00'),(50,50,'Flat tire','2025-03-27 18:45:00'),(51,5,'Engine Fail','2025-04-11 18:48:51');
/*!40000 ALTER TABLE `incident` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maintenance`
--

DROP TABLE IF EXISTS `maintenance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `maintenance` (
  `MAINTENANCE_ID` int(11) NOT NULL,
  `V_ID` int(11) DEFAULT NULL,
  `ISSUE_REPORTED` text,
  `REPAIR_STATUS` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`MAINTENANCE_ID`),
  KEY `V_ID` (`V_ID`),
  CONSTRAINT `maintenance_ibfk_1` FOREIGN KEY (`V_ID`) REFERENCES `vehicle` (`VEHICLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maintenance`
--

LOCK TABLES `maintenance` WRITE;
/*!40000 ALTER TABLE `maintenance` DISABLE KEYS */;
INSERT INTO `maintenance` VALUES (1,1,'Oil Change Required','Completed'),(2,3,'Battery Replacement','Completed'),(3,5,'Brake Pad Replacement','Completed'),(4,7,'Engine Check','Completed'),(5,9,'Transmission Fluid Change','In Progress'),(6,11,'Suspension Check','Pending'),(7,10,'Khatam sab khatam','Pending'),(8,15,'Air Filter Change','In Progress'),(9,17,'AC Repair','Pending'),(10,19,'Tire Rotation','Completed'),(11,21,'Brake Fluid Change','In Progress'),(12,23,'Fuel Pump Replacement','Pending'),(13,25,'Steering Alignment','Completed'),(14,27,'Exhaust System Repair','In Progress'),(15,29,'Headlight Replacement','Pending'),(16,31,'Windshield Crack Repair','Completed'),(17,33,'Battery Charge Check','In Progress'),(18,35,'Radiator Flush','Pending'),(19,37,'Door Lock Fix','Completed'),(20,39,'Fuel Injector Cleaning','In Progress'),(21,41,'Wiper Blade Replacement','Pending'),(22,43,'Seatbelt Repair','Completed'),(23,45,'Horn Malfunction Fix','In Progress'),(24,47,'Head Gasket Replacement','Pending'),(25,49,'Dashboard Light Fix','Completed'),(26,50,'Brake Disc Resurfacing','In Progress'),(27,48,'Spark Plug Replacement','Pending'),(28,46,'Throttle Body Cleaning','Completed'),(29,44,'Differential Fluid Change','In Progress'),(30,42,'Oil Leak Repair','Pending'),(31,40,'ECU Software Update','Completed'),(32,38,'Clutch Replacement','In Progress'),(33,36,'Power Steering Fluid Change','Pending'),(34,34,'Fuel Filter Change','Completed'),(35,32,'Catalytic Converter Check','In Progress'),(36,30,'Drive Belt Replacement','Pending'),(37,28,'Oxygen Sensor Replacement','Completed'),(38,26,'Alternator Check','In Progress'),(39,24,'Shock Absorber Replacement','Pending'),(40,22,'Starter Motor Check','Completed'),(41,20,'Timing Belt Replacement','In Progress'),(42,18,'Tire Pressure Monitoring','Pending'),(43,16,'Glow Plug Replacement','Completed'),(44,14,'Heater Core Flush','In Progress'),(45,12,'Turbocharger Inspection','Pending'),(46,10,'Tut gayi re baabbababa','Completed'),(47,8,'CV Joint Replacement','In Progress'),(48,6,'Battery Terminal Cleaning','Pending'),(49,4,'Water Pump Replacement','Completed'),(50,2,'Tire Replacement','In Progress'),(51,5,'Mar gaya','Completed');
/*!40000 ALTER TABLE `maintenance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maintenance_history`
--

DROP TABLE IF EXISTS `maintenance_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `maintenance_history` (
  `M_ID` int(11) NOT NULL,
  `V_ID` int(11) NOT NULL,
  PRIMARY KEY (`M_ID`,`V_ID`),
  KEY `V_ID` (`V_ID`),
  CONSTRAINT `maintenance_history_ibfk_1` FOREIGN KEY (`M_ID`) REFERENCES `maintenance` (`MAINTENANCE_ID`),
  CONSTRAINT `maintenance_history_ibfk_2` FOREIGN KEY (`V_ID`) REFERENCES `vehicle` (`VEHICLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maintenance_history`
--

LOCK TABLES `maintenance_history` WRITE;
/*!40000 ALTER TABLE `maintenance_history` DISABLE KEYS */;
INSERT INTO `maintenance_history` VALUES (1,1),(2,2),(3,3),(4,4),(5,5),(50,5),(6,6),(7,7),(8,8),(9,9),(10,10),(11,11),(12,12),(13,13),(14,14),(15,15),(16,16),(17,17),(18,18),(19,19),(20,20),(21,21),(22,22),(23,23),(24,24),(25,25),(26,26),(27,27),(28,28),(29,29),(30,30),(31,31),(32,32),(33,33),(34,34),(35,35),(36,36),(37,37),(38,38),(39,39),(40,40),(41,41),(42,42),(43,43),(44,44),(45,45),(46,46),(47,47),(48,48),(49,49);
/*!40000 ALTER TABLE `maintenance_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operates_on`
--

DROP TABLE IF EXISTS `operates_on`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `operates_on` (
  `V_ID` int(11) NOT NULL,
  `S_ID` int(11) NOT NULL,
  PRIMARY KEY (`V_ID`,`S_ID`),
  KEY `S_ID` (`S_ID`),
  CONSTRAINT `operates_on_ibfk_1` FOREIGN KEY (`V_ID`) REFERENCES `vehicle` (`VEHICLE_ID`),
  CONSTRAINT `operates_on_ibfk_2` FOREIGN KEY (`S_ID`) REFERENCES `schedule` (`SCHEDULE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operates_on`
--

LOCK TABLES `operates_on` WRITE;
/*!40000 ALTER TABLE `operates_on` DISABLE KEYS */;
INSERT INTO `operates_on` VALUES (1,1),(2,2),(5,2),(6,2),(3,3),(4,4),(6,6),(7,7),(8,8),(9,9),(10,10),(11,11),(12,12),(13,13),(14,14),(15,15),(16,16),(17,17),(18,18),(19,19),(20,20),(21,21),(22,22),(23,23),(24,24),(25,25),(26,26),(27,27),(28,28),(29,29),(30,30),(31,31),(32,32),(33,33),(34,34),(35,35),(36,36),(37,37),(38,38),(39,39),(40,40),(41,41),(42,42),(43,43),(44,44),(45,45),(46,46),(47,47),(48,48),(49,49),(50,50);
/*!40000 ALTER TABLE `operates_on` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment`
--

DROP TABLE IF EXISTS `payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment` (
  `PAYMENT_ID` int(11) NOT NULL,
  `PASSENGER_ID` int(11) DEFAULT NULL,
  `AMOUNT` decimal(10,2) DEFAULT NULL,
  `METHOD` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`PAYMENT_ID`),
  KEY `PASSENGER_ID` (`PASSENGER_ID`),
  CONSTRAINT `payment_ibfk_1` FOREIGN KEY (`PASSENGER_ID`) REFERENCES `human` (`ID_NO`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment`
--

LOCK TABLES `payment` WRITE;
/*!40000 ALTER TABLE `payment` DISABLE KEYS */;
INSERT INTO `payment` VALUES (1,2,345678.00,'Cash'),(2,5,20.00,'Cash'),(3,7,18.75,'Debit Card'),(4,10,22.00,'Mobile Wallet'),(5,12,19.50,'Credit Card'),(6,15,17.25,'Cash'),(7,18,25.00,'Debit Card'),(8,20,30.75,'Mobile Wallet'),(9,22,12.00,'Credit Card'),(10,25,16.50,'Cash'),(11,27,14.75,'Debit Card'),(12,30,21.00,'Mobile Wallet'),(13,32,23.25,'Credit Card'),(14,35,19.00,'Cash'),(15,38,15.50,'Debit Card'),(16,40,26.00,'Mobile Wallet'),(17,42,28.75,'Credit Card'),(18,45,11.50,'Cash'),(19,47,22.50,'Debit Card'),(20,49,20.25,'Mobile Wallet'),(21,3,213.00,'Credit Card'),(22,6,24.00,'Cash'),(23,9,16.75,'Debit Card'),(24,11,12.50,'Mobile Wallet'),(25,13,29.00,'Credit Card'),(26,16,14.25,'Cash'),(27,19,27.50,'Debit Card'),(28,21,25.25,'Mobile Wallet'),(29,23,21.75,'Credit Card'),(30,26,13.50,'Cash'),(31,28,17.25,'Debit Card'),(32,31,20.00,'Mobile Wallet'),(33,33,15.75,'Credit Card'),(34,36,18.50,'Cash'),(35,39,24.75,'Debit Card'),(36,41,29.50,'Mobile Wallet'),(37,43,22.00,'Credit Card'),(38,46,16.00,'Cash'),(39,48,14.50,'Debit Card'),(40,50,19.25,'Mobile Wallet'),(41,1,27.75,'Credit Card'),(42,4,13.00,'Cash'),(43,8,15.25,'Debit Card'),(44,14,17.00,'Mobile Wallet'),(45,17,23.50,'Credit Card'),(46,29,12.75,'Cash');
/*!40000 ALTER TABLE `payment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `performs_maintenance`
--

DROP TABLE IF EXISTS `performs_maintenance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `performs_maintenance` (
  `M_ID` int(11) NOT NULL,
  `STAFF_ID` int(11) NOT NULL,
  PRIMARY KEY (`M_ID`,`STAFF_ID`),
  KEY `STAFF_ID` (`STAFF_ID`),
  CONSTRAINT `performs_maintenance_ibfk_1` FOREIGN KEY (`M_ID`) REFERENCES `maintenance` (`MAINTENANCE_ID`),
  CONSTRAINT `performs_maintenance_ibfk_2` FOREIGN KEY (`STAFF_ID`) REFERENCES `human` (`ID_NO`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `performs_maintenance`
--

LOCK TABLES `performs_maintenance` WRITE;
/*!40000 ALTER TABLE `performs_maintenance` DISABLE KEYS */;
INSERT INTO `performs_maintenance` VALUES (1,1),(2,2),(3,3),(4,4),(3,5),(5,5),(6,6),(7,7),(8,8),(9,9),(10,10),(11,11),(12,12),(13,13),(14,14),(15,15),(16,16),(17,17),(18,18),(19,19),(20,20),(21,21),(22,22),(23,23),(24,24),(25,25),(26,26),(27,27),(28,28),(29,29),(30,30),(31,31),(32,32),(33,33),(34,34),(35,35),(36,36),(37,37),(38,38),(39,39),(40,40),(41,41),(42,42),(43,43),(44,44),(45,45),(46,46),(47,47),(48,48),(49,49),(50,50);
/*!40000 ALTER TABLE `performs_maintenance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route`
--

DROP TABLE IF EXISTS `route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `route` (
  `R_ID` int(11) NOT NULL,
  `JOURNEY_TIME` int(11) DEFAULT NULL,
  `START_POINT` varchar(100) DEFAULT NULL,
  `END_POINT` varchar(100) DEFAULT NULL,
  `DISTANCE` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`R_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route`
--

LOCK TABLES `route` WRITE;
/*!40000 ALTER TABLE `route` DISABLE KEYS */;
INSERT INTO `route` VALUES (0,3,'Radaur','Greater Noida',33.50),(1,5,'Room502','Room1001',2.00),(2,20,'D Block','A Block',1.00),(3,105,'Chicago','Milwaukee',148.30),(4,310,'Houston','Dallas',385.60),(5,865,'Dadri 4','Radaur H',300.00),(6,405,'Denver','Salt Lake City',835.50),(7,245,'Seattle','Portland',280.20),(8,175,'Atlanta','Charlotte',365.10),(9,450,'San Diego','Las Vegas',530.75),(10,220,'Detroit','Cleveland',275.80),(11,495,'Phoenix','Albuquerque',750.90),(12,290,'Philadelphia','Washington DC',225.50),(13,140,'San Antonio','Austin',125.70),(14,390,'Minneapolis','Kansas City',700.30),(15,350,'Indianapolis','Louisville',185.40),(16,190,'Columbus','Pittsburgh',290.60),(17,160,'Buffalo','Rochester',120.90),(18,440,'Las Vegas','Phoenix',500.80),(19,360,'Memphis','Nashville',380.70),(20,325,'St. Louis','Oklahoma City',630.50),(21,215,'Baltimore','Richmond',250.40),(22,170,'Charlotte','Raleigh',260.30),(23,250,'Portland','Boise',670.20),(24,315,'Milwaukee','Indianapolis',410.90),(25,400,'New Orleans','Baton Rouge',170.80),(26,470,'Salt Lake City','Boise',660.70),(27,225,'Omaha','Des Moines',210.60),(28,295,'Tampa','Jacksonville',360.50),(29,155,'Hartford','Providence',110.90),(30,480,'Anchorage','Fairbanks',860.40),(31,185,'Reno','Sacramento',220.30),(32,430,'Tulsa','Little Rock',380.20),(33,305,'Colorado Springs','Albuquerque',560.70),(34,145,'Albany','Syracuse',150.90),(35,380,'San Francisco','Reno',350.60),(36,270,'El Paso','Tucson',430.50),(37,205,'Knoxville','Lexington',240.80),(38,460,'Wichita','Kansas City',320.70),(39,410,'Birmingham','Montgomery',155.40),(40,285,'Augusta','Savannah',250.30),(41,335,'Charleston','Columbia',210.20),(42,135,'Manchester','Boston',90.70),(43,370,'Boise','Spokane',590.80),(44,235,'Lubbock','Amarillo',200.50),(45,420,'Corpus Christi','San Antonio',250.90),(46,500,'Honolulu','Hilo',420.60),(47,275,'Springfield','St. Louis',310.40),(48,345,'Greensboro','Raleigh',280.50),(49,165,'Sioux Falls','Omaha',270.30),(50,415,'Fargo','Minneapolis',395.20),(51,50,'Dheradun','Radaur',100.00),(52,120,'Radaur Bus Stand','ISBT Delhi',400.00),(53,31,'Greater Noida','Delhi',5.00);
/*!40000 ALTER TABLE `route` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_followed`
--

DROP TABLE IF EXISTS `route_followed`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `route_followed` (
  `VEHICLE_ID` int(11) NOT NULL,
  `RT_ID` int(11) NOT NULL,
  PRIMARY KEY (`VEHICLE_ID`,`RT_ID`),
  KEY `RT_ID` (`RT_ID`),
  CONSTRAINT `route_followed_ibfk_1` FOREIGN KEY (`VEHICLE_ID`) REFERENCES `vehicle` (`VEHICLE_ID`),
  CONSTRAINT `route_followed_ibfk_2` FOREIGN KEY (`RT_ID`) REFERENCES `route` (`R_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_followed`
--

LOCK TABLES `route_followed` WRITE;
/*!40000 ALTER TABLE `route_followed` DISABLE KEYS */;
INSERT INTO `route_followed` VALUES (1,1),(2,2),(3,3),(4,4),(5,5),(6,6),(7,7),(8,8),(9,9),(10,10),(11,11),(12,12),(13,13),(14,14),(15,15),(16,16),(17,17),(18,18),(19,19),(20,20),(21,21),(22,22),(23,23),(24,24),(25,25),(26,26),(27,27),(28,28),(29,29),(30,30),(31,31),(32,32),(33,33),(34,34),(35,35),(36,36),(37,37),(38,38),(39,39),(40,40),(41,41),(42,42),(43,43),(44,44),(45,45),(46,46),(47,47),(48,48),(49,49),(50,50);
/*!40000 ALTER TABLE `route_followed` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedule`
--

DROP TABLE IF EXISTS `schedule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schedule` (
  `SCHEDULE_ID` int(11) NOT NULL,
  `R_ID` int(11) DEFAULT NULL,
  `V_ID` int(11) DEFAULT NULL,
  `DEPARTURE_TIME` time DEFAULT NULL,
  `ARRIVAL_TIME` time DEFAULT NULL,
  PRIMARY KEY (`SCHEDULE_ID`),
  KEY `R_ID` (`R_ID`),
  KEY `V_ID` (`V_ID`),
  CONSTRAINT `schedule_ibfk_1` FOREIGN KEY (`R_ID`) REFERENCES `route` (`R_ID`),
  CONSTRAINT `schedule_ibfk_2` FOREIGN KEY (`V_ID`) REFERENCES `vehicle` (`VEHICLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedule`
--

LOCK TABLES `schedule` WRITE;
/*!40000 ALTER TABLE `schedule` DISABLE KEYS */;
INSERT INTO `schedule` VALUES (1,1,1,'08:00:00','10:30:00'),(2,2,2,'09:00:00','12:45:00'),(3,3,3,'10:15:00','13:00:00'),(4,4,4,'11:00:00','14:20:00'),(5,5,5,'12:30:00','15:10:00'),(6,6,6,'13:00:00','16:30:00'),(7,7,7,'14:45:00','17:20:00'),(8,8,8,'15:30:00','18:00:00'),(9,9,9,'16:10:00','19:15:00'),(10,10,10,'17:00:00','20:00:00'),(11,11,11,'18:30:00','21:15:00'),(12,12,12,'19:15:00','22:00:00'),(13,13,13,'20:00:00','23:30:00'),(14,14,14,'21:00:00','00:30:00'),(15,15,15,'22:10:00','01:00:00'),(16,16,16,'23:30:00','02:15:00'),(17,17,17,'00:45:00','03:10:00'),(18,18,18,'01:30:00','04:00:00'),(19,19,19,'02:15:00','05:00:00'),(20,20,20,'03:00:00','06:30:00'),(21,21,21,'04:45:00','07:20:00'),(22,22,22,'05:30:00','08:00:00'),(23,23,23,'06:10:00','09:15:00'),(24,24,24,'07:00:00','10:00:00'),(25,25,25,'08:30:00','11:15:00'),(26,1,26,'09:10:00','12:00:00'),(27,2,27,'10:00:00','13:30:00'),(28,3,28,'11:15:00','14:45:00'),(29,4,29,'12:45:00','16:20:00'),(30,5,30,'14:00:00','17:30:00'),(31,6,31,'15:30:00','18:50:00'),(32,7,32,'16:10:00','19:15:00'),(33,8,33,'17:45:00','21:00:00'),(34,9,34,'18:30:00','21:45:00'),(35,10,35,'19:00:00','22:30:00'),(36,11,36,'20:30:00','23:50:00'),(37,12,37,'21:10:00','00:40:00'),(38,13,38,'22:30:00','02:00:00'),(39,14,39,'23:45:00','03:15:00'),(40,15,40,'00:30:00','04:00:00'),(41,16,41,'01:15:00','05:10:00'),(42,17,42,'02:00:00','06:00:00'),(43,18,43,'03:20:00','07:30:00'),(44,19,44,'04:45:00','08:15:00'),(45,20,45,'05:30:00','09:00:00'),(46,21,46,'06:15:00','10:20:00'),(47,22,47,'07:00:00','11:15:00'),(48,23,48,'08:30:00','12:45:00'),(49,24,49,'12:00:00','14:50:00'),(50,25,50,'14:00:00','15:20:00'),(51,5,5,'15:04:05','19:04:05'),(54,3,6,'00:54:00','01:21:00'),(55,7,5,'00:54:00','01:21:00');
/*!40000 ALTER TABLE `schedule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedule_followed`
--

DROP TABLE IF EXISTS `schedule_followed`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schedule_followed` (
  `R_ID` int(11) NOT NULL,
  `S_ID` int(11) NOT NULL,
  PRIMARY KEY (`R_ID`,`S_ID`),
  KEY `S_ID` (`S_ID`),
  CONSTRAINT `schedule_followed_ibfk_1` FOREIGN KEY (`R_ID`) REFERENCES `route` (`R_ID`),
  CONSTRAINT `schedule_followed_ibfk_2` FOREIGN KEY (`S_ID`) REFERENCES `schedule` (`SCHEDULE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedule_followed`
--

LOCK TABLES `schedule_followed` WRITE;
/*!40000 ALTER TABLE `schedule_followed` DISABLE KEYS */;
INSERT INTO `schedule_followed` VALUES (1,1),(2,2),(3,3),(4,4),(6,6),(7,7),(8,8),(9,9),(10,10),(11,11),(12,12),(13,13),(14,14),(15,15),(16,16),(17,17),(18,18),(19,19),(20,20),(21,21),(22,22),(23,23),(24,24),(25,25),(26,26),(27,27),(28,28),(29,29),(30,30),(31,31),(32,32),(33,33),(34,34),(35,35),(36,36),(37,37),(38,38),(39,39),(40,40),(41,41),(42,42),(43,43),(44,44),(45,45),(46,46),(47,47),(48,48),(49,49),(50,50);
/*!40000 ALTER TABLE `schedule_followed` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `ID_NO` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `roles` varchar(40) DEFAULT 'user',
  PRIMARY KEY (`ID_NO`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`ID_NO`) REFERENCES `human` (`ID_NO`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Randomly created#1','$2a$10$SVXLFw09eH8otH7sX8/lX.lHd7b/8AuN0jWZLsFuMd7lXltqYFYG.','Maintenance Manager'),(2,'Xenomorph','$2a$10$TZMcsBoqqXYfkRg/NxHjtO01/aAMGYds4GyDQRfcNLbglyrw3EpuS','user'),(3,'Flash_boi_05','$2a$10$h6G.dqiJqqfkLvvQ7gk.7u0jyKcqTCtVWkaNg6CSmK4oeMWI07qv.','Admin'),(4,'Incident Manager','$2a$10$4n2mg.eVdnhOJ/vyyhvznOOxiVMYhdjetaEc29LSs.WwwDVpjeAIC','Incident Manager'),(5,'Schedule Manager','$2a$10$ZOOnS44pzBeDLDuReiFGwu8VOyFZM1adQIHocg4hQcEiAFz1hMipy','Schedule Manager'),(6,'Pata nahi kon hein','$2a$10$mNtwpNa.5ByW/9mDfnJ4YuhFsT3aDk0lIZo.jDSoGuNKN.pMQ8vZ.','Vehicle Manager'),(7,'Route Manager','$2a$10$I588OfzrVUddsoKImG0O7uxzuSd2WskISu7YivNx3iQdj4sC.Dfte','Route Manager'),(8,'Accident Manager','$2a$10$jwwnUc3IMs9aJh3pvc0iLe38aXwOjemXQ2WcLNi2XwrW/tv0NpnQy','Accident Manager'),(17,'Maintenance Manager','$2a$10$GfF.JvKwJyaECplfYCHaD.MG24WYbPhWIX7hb7VKFyUoCUIv5mkQO','Maintenance Manager'),(20,'TestSubject','$2a$10$yKJJVHv0I5Fc/4GPb7A.NuDvfPt6OTQ8CwuF9TOMgEOATIRrgPYaS','user'),(40,'Payment Manager','$2a$10$IGaywG0UORCfJ8vVCNxkv.unxATGIaaJnEBa4pNUeWmAdl9l3prSq','Payment Manager'),(50,'random2','$2a$10$XBKaydnrnLwAMpka5qQD9.gYe0uCJGgjkJ5dWJFXLQQJkendSAEAG','user'),(53,'Tempo','$2a$10$320AFqG44J/3T8fzv5OE2OAp30OQKjOV8kDFY.WB2bOps0/hAVHEK','user'),(55,'test#1','$2a$10$9RmNWErCJB1fzsisaM36GuLBzt0C8WW/9/LHVDxionnVPNt.lwt1u','user'),(56,'FlashIsLive','$2a$10$r92UB1IXs9QlAeDh4KOzXusQgfj1ZBR4veVGLH9M.lLhZqBSSpyeu','user'),(57,'idk','$2a$10$LT2N9ry9BfJCf1B0PL1.RuXowE/1N11ESiao2LfUwTWq20CXhtpBm','user');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vehicle`
--

DROP TABLE IF EXISTS `vehicle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vehicle` (
  `VEHICLE_ID` int(11) NOT NULL,
  `CURRENT_LOCATION` varchar(100) DEFAULT NULL,
  `STATUS` varchar(50) DEFAULT NULL,
  `LAST_UPDATE` datetime DEFAULT NULL,
  `status_change` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`VEHICLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vehicle`
--

LOCK TABLES `vehicle` WRITE;
/*!40000 ALTER TABLE `vehicle` DISABLE KEYS */;
INSERT INTO `vehicle` VALUES (0,'','Booked','2025-04-18 19:18:15',0),(1,'New York','Available','2025-04-19 01:43:58',0),(2,'Dadri','Available','2025-04-18 19:33:04',0),(3,'Shiv Nadar University','Available','2025-04-16 13:20:36',0),(4,'Houston','Available','2025-03-30 10:15:00',0),(5,'New York','Available','2025-04-16 12:19:42',0),(6,'Philadelphia','Available','2025-03-30 10:25:00',0),(7,'San Antonio','Available','2025-03-30 10:30:00',0),(8,'San Diego','Available','2025-03-30 10:35:00',0),(9,'','Available','2025-04-18 14:19:53',0),(10,'Jhanum','Available','2025-04-17 14:05:40',0),(11,'Austin','Available','2025-03-30 10:50:00',0),(12,'Jacksonville','Available','2025-03-30 10:55:00',0),(13,'Fort Worth','Available','2025-03-30 11:00:00',0),(14,'Dadri','Available','2025-04-10 19:08:34',0),(15,'San Francisco','Available','2025-03-30 11:10:00',0),(16,'Charlotte','Available','2025-03-30 11:15:00',0),(17,'Indianapolis','Available','2025-03-30 11:20:00',0),(18,'Seattle','Available','2025-03-30 11:25:00',0),(19,'Denver','Available','2025-03-30 11:30:00',0),(20,'Dadri bc','Available','2025-04-10 14:07:50',0),(21,'Boston','Available','2025-03-30 11:40:00',0),(22,'El Paso','Available','2025-03-30 11:45:00',0),(23,'Nashville','Available','2025-03-30 11:50:00',0),(24,'Detroit','Available','2025-03-30 11:55:00',0),(25,'Oklahoma City','Available','2025-03-30 12:00:00',0),(26,'Portland','Available','2025-03-30 12:05:00',0),(27,'Las Vegas','Available','2025-03-30 12:10:00',0),(28,'Memphis','Available','2025-03-30 12:15:00',0),(29,'Louisville','Available','2025-03-30 12:20:00',0),(30,'Baltimore','Available','2025-03-30 12:25:00',0),(31,'Milwaukee','Available','2025-03-30 12:30:00',0),(32,'Albuquerque','Available','2025-03-30 12:35:00',0),(33,'Tucson','Available','2025-03-30 12:40:00',0),(34,'Fresno','Available','2025-03-30 12:45:00',0),(35,'Sacramento','Available','2025-03-30 12:50:00',0),(36,'Mesa','Available','2025-03-30 12:55:00',0),(37,'Kansas City','Available','2025-03-30 13:00:00',0),(38,'Atlanta','Available','2025-03-30 13:05:00',0),(39,'Long Beach','Available','2025-03-30 13:10:00',0),(40,'Omaha','Available','2025-03-30 13:15:00',0),(41,'Raleigh','Available','2025-03-30 13:20:00',0),(42,'Miami','Available','2025-03-30 13:25:00',0),(43,'Oakland','Available','2025-03-30 13:30:00',0),(44,'Minneapolis','Available','2025-03-30 13:35:00',0),(45,'Tulsa','Available','2025-03-30 13:40:00',0),(46,'Cleveland','Available','2025-03-30 13:45:00',0),(47,'Wichita','Available','2025-03-30 13:50:00',0),(48,'New Orleans','Available','2025-03-30 13:55:00',0),(49,'Arlington','Available','2025-03-30 14:00:00',0),(50,'Bakersfield','Available','2025-03-30 14:05:00',0),(51,'Sector 14, SNIOE','Available','2025-04-08 20:59:24',0),(52,'Asia','Available','2025-04-13 08:46:19',0),(53,'D Block','Available','2025-04-16 13:58:28',0),(54,'A Block','Available','2025-04-16 14:04:23',0),(55,'New York','Available','2025-04-16 15:56:06',0);
/*!40000 ALTER TABLE `vehicle` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = cp850 */ ;
/*!50003 SET character_set_results = cp850 */ ;
/*!50003 SET collation_connection  = cp850_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER vehicle_booked_trigger
AFTER UPDATE ON vehicle
FOR EACH ROW
BEGIN
  
  IF OLD.STATUS = 'Available' AND NEW.STATUS = 'Booked' THEN
    
    IF EXISTS (
      SELECT 1 FROM bookings WHERE vehicle_id = NEW.VEHICLE_ID
    ) THEN
      
      UPDATE bookings
      SET booking_count = booking_count + 1
      WHERE vehicle_id = NEW.VEHICLE_ID;
    ELSE
      
      INSERT INTO bookings (vehicle_id, booking_count)
      VALUES (NEW.VEHICLE_ID, 1);
    END IF;
  END IF;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `vehicle_status_change`
--

DROP TABLE IF EXISTS `vehicle_status_change`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vehicle_status_change` (
  `vehicle_id` int(11) NOT NULL,
  `status_change` tinyint(1) DEFAULT '0',
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`vehicle_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vehicle_status_change`
--

LOCK TABLES `vehicle_status_change` WRITE;
/*!40000 ALTER TABLE `vehicle_status_change` DISABLE KEYS */;
INSERT INTO `vehicle_status_change` VALUES (1,0,'2025-04-19 01:43:58');
/*!40000 ALTER TABLE `vehicle_status_change` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-04-19  3:40:02
