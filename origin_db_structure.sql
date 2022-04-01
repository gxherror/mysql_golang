-- MySQL dump 10.16  Distrib 10.1.26-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: db
-- ------------------------------------------------------
-- Server version	10.1.26-MariaDB-0+deb9u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `advisor`
--

DROP TABLE IF EXISTS `advisor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `advisor` (
  `s_ID` mediumint(9) DEFAULT NULL,
  `i_ID` mediumint(9) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `advisor`
--

LOCK TABLES `advisor` WRITE;
/*!40000 ALTER TABLE `advisor` DISABLE KEYS */;
INSERT INTO `advisor` VALUES (128,45565),(12345,10101),(23121,76543),(44553,22222),(45678,22222),(76543,45565),(76653,98345),(98765,98345),(98988,76766);
/*!40000 ALTER TABLE `advisor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `classroom`
--

DROP TABLE IF EXISTS `classroom`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `classroom` (
  `building` varchar(7) DEFAULT NULL,
  `room_number` smallint(6) DEFAULT NULL,
  `capacity` smallint(6) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classroom`
--

LOCK TABLES `classroom` WRITE;
/*!40000 ALTER TABLE `classroom` DISABLE KEYS */;
INSERT INTO `classroom` VALUES ('Packard',101,500),('Painter',514,10),('Taylor',3128,70),('Watson',100,30),('Watson',120,50);
/*!40000 ALTER TABLE `classroom` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `course` (
  `course_id` varchar(7) DEFAULT NULL,
  `title` varchar(26) DEFAULT NULL,
  `dept_name` varchar(10) DEFAULT NULL,
  `credits` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES ('BIO-101','Intro. to Biology','Biology',4),('BIO-301','Genetics','Biology',4),('BIO-399','Computational Biology','Biology',3),('CS-101','Intro. to Computer Science','Comp. Sci.',4),('CS-190','Game Design','Comp. Sci.',4),('CS-315','Robotics','Comp. Sci.',3),('CS-319','Image Processing','Comp. Sci.',3),('CS-347','Database System Concepts','Comp. Sci.',3),('EE-181','Intro. to Digital Systems','Elec. Eng.',3),('FIN-201','Investment Banking','Finance',3),('HIS-351','World History','History',3),('MU-199','Music Video Production','Music',3),('PHY-101','Physical Principles','Physics',4);
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `department`
--

DROP TABLE IF EXISTS `department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `department` (
  `dept_name` varchar(10) DEFAULT NULL,
  `building` varchar(7) DEFAULT NULL,
  `budget` mediumint(9) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `department`
--

LOCK TABLES `department` WRITE;
/*!40000 ALTER TABLE `department` DISABLE KEYS */;
INSERT INTO `department` VALUES ('Biology','Watson',90000),('Comp. Sci.','Taylor',100000),('Elec. Eng.','Taylor',85000),('Finance','Painter',120000),('History','Painter',50000),('Music','Packard',80000),('Physics','Watson',70000);
/*!40000 ALTER TABLE `department` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `instructor`
--

DROP TABLE IF EXISTS `instructor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `instructor` (
  `ID` mediumint(9) DEFAULT NULL,
  `name` varchar(10) DEFAULT NULL,
  `dept_name` varchar(10) DEFAULT NULL,
  `salary` mediumint(9) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `instructor`
--

LOCK TABLES `instructor` WRITE;
/*!40000 ALTER TABLE `instructor` DISABLE KEYS */;
INSERT INTO `instructor` VALUES (10101,'Srinivasan','Comp. Sci.',65000),(12121,'Wu','Finance',90000),(15151,'Mozart','Music',40000),(22222,'Einstein','Physics',95000),(32343,'El Said','History',60000),(33456,'Gold','Physics',87000),(45565,'Katz','Comp. Sci.',75000),(58583,'Califieri','History',62000),(76543,'Singh','Finance',80000),(76766,'Crick','Biology',72000),(83821,'Brandt','Comp. Sci.',92000),(98345,'Kim','Elec. Eng.',80000);
/*!40000 ALTER TABLE `instructor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prereq`
--

DROP TABLE IF EXISTS `prereq`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `prereq` (
  `course_id` varchar(7) DEFAULT NULL,
  `prereq_id` varchar(7) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prereq`
--

LOCK TABLES `prereq` WRITE;
/*!40000 ALTER TABLE `prereq` DISABLE KEYS */;
INSERT INTO `prereq` VALUES ('BIO-301','BIO-101'),('BIO-399','BIO-101'),('CS-190','CS-101'),('CS-315','CS-101'),('CS-319','CS-101'),('CS-347','CS-101'),('EE-181','PHY-101');
/*!40000 ALTER TABLE `prereq` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `section`
--

DROP TABLE IF EXISTS `section`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `section` (
  `course_id` varchar(7) DEFAULT NULL,
  `sec_id` tinyint(4) DEFAULT NULL,
  `semester` varchar(6) DEFAULT NULL,
  `year` smallint(6) DEFAULT NULL,
  `building` varchar(7) DEFAULT NULL,
  `room_number` smallint(6) DEFAULT NULL,
  `time_slot_id` varchar(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `section`
--

LOCK TABLES `section` WRITE;
/*!40000 ALTER TABLE `section` DISABLE KEYS */;
INSERT INTO `section` VALUES ('BIO-101',1,'Summer',2017,'Painter',514,'B'),('BIO-301',1,'Summer',2018,'Painter',514,'A'),('CS-101',1,'Fall',2017,'Packard',101,'H'),('CS-101',1,'Spring',2018,'Packard',101,'F'),('CS-190',1,'Spring',2017,'Taylor',3128,'E'),('CS-190',2,'Spring',2017,'Taylor',3128,'A'),('CS-315',1,'Spring',2018,'Watson',120,'D'),('CS-319',1,'Spring',2018,'Watson',100,'B'),('CS-319',2,'Spring',2018,'Taylor',3128,'C'),('CS-347',1,'Fall',2017,'Taylor',3128,'A'),('EE-181',1,'Spring',2017,'Taylor',3128,'C'),('FIN-201',1,'Spring',2018,'Packard',101,'B'),('HIS-351',1,'Spring',2018,'Painter',514,'C'),('MU-199',1,'Spring',2018,'Packard',101,'D'),('PHY-101',1,'Fall',2017,'Watson',100,'A');
/*!40000 ALTER TABLE `section` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student` (
  `ID` mediumint(9) DEFAULT NULL,
  `name` varchar(8) DEFAULT NULL,
  `dept_name` varchar(10) DEFAULT NULL,
  `tot_cred` smallint(6) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--

LOCK TABLES `student` WRITE;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
INSERT INTO `student` VALUES (128,'Zhang','Comp. Sci.',102),(12345,'Shankar','Comp. Sci.',32),(19991,'Brandt','History',80),(23121,'Chavez','Finance',110),(44553,'Peltier','Physics',56),(45678,'Levy','Physics',46),(54321,'Williams','Comp. Sci.',54),(55739,'Sanchez','Music',38),(70557,'Snow','Physics',0),(76543,'Brown','Comp. Sci.',58),(76653,'Aoi','Elec. Eng.',60),(98765,'Bourikas','Elec. Eng.',98),(98988,'Tanaka','Biology',120);
/*!40000 ALTER TABLE `student` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `takes`
--

DROP TABLE IF EXISTS `takes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `takes` (
  `ID` mediumint(9) DEFAULT NULL,
  `course_id` varchar(7) DEFAULT NULL,
  `sec_id` tinyint(4) DEFAULT NULL,
  `semester` varchar(6) DEFAULT NULL,
  `year` smallint(6) DEFAULT NULL,
  `grade` varchar(2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `takes`
--

LOCK TABLES `takes` WRITE;
/*!40000 ALTER TABLE `takes` DISABLE KEYS */;
INSERT INTO `takes` VALUES (128,'CS-101',1,'Fall',2017,'A'),(128,'CS-347',1,'Fall',2017,'A-'),(12345,'CS-101',1,'Fall',2017,'C'),(12345,'CS-190',2,'Spring',2017,'A'),(12345,'CS-315',1,'Spring',2018,'A'),(12345,'CS-347',1,'Fall',2017,'A'),(19991,'HIS-351',1,'Spring',2018,'B'),(23121,'FIN-201',1,'Spring',2018,'C+'),(44553,'PHY-101',1,'Fall',2017,'B-'),(45678,'CS-101',1,'Fall',2017,'F'),(45678,'CS-101',1,'Spring',2018,'B+'),(45678,'CS-319',1,'Spring',2018,'B'),(54321,'CS-101',1,'Fall',2017,'A-'),(54321,'CS-190',2,'Spring',2017,'B+'),(55739,'MU-199',1,'Spring',2018,'A-'),(76543,'CS-101',1,'Fall',2017,'A'),(76543,'CS-319',2,'Spring',2018,'A'),(76653,'EE-181',1,'Spring',2017,'C'),(98765,'CS-101',1,'Fall',2017,'C-'),(98765,'CS-315',1,'Spring',2018,'B'),(98988,'BIO-101',1,'Summer',2017,'A'),(98988,'BIO-301',1,'Summer',2018,'');
/*!40000 ALTER TABLE `takes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teaches`
--

DROP TABLE IF EXISTS `teaches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teaches` (
  `ID` mediumint(9) DEFAULT NULL,
  `course_id` varchar(7) DEFAULT NULL,
  `sec_id` tinyint(4) DEFAULT NULL,
  `semester` varchar(6) DEFAULT NULL,
  `year` smallint(6) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teaches`
--

LOCK TABLES `teaches` WRITE;
/*!40000 ALTER TABLE `teaches` DISABLE KEYS */;
INSERT INTO `teaches` VALUES (10101,'CS-101',1,'Fall',2017),(10101,'CS-315',1,'Spring',2018),(10101,'CS-347',1,'Fall',2017),(12121,'FIN-201',1,'Spring',2018),(15151,'MU-199',1,'Spring',2018),(22222,'PHY-101',1,'Fall',2017),(32343,'HIS-351',1,'Spring',2018),(45565,'CS-101',1,'Spring',2018),(45565,'CS-319',1,'Spring',2018),(76766,'BIO-101',1,'Summer',2017),(76766,'BIO-301',1,'Summer',2018),(83821,'CS-190',1,'Spring',2017),(83821,'CS-190',2,'Spring',2017),(83821,'CS-319',2,'Spring',2018),(98345,'EE-181',1,'Spring',2017);
/*!40000 ALTER TABLE `teaches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `time_slot`
--

DROP TABLE IF EXISTS `time_slot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `time_slot` (
  `time_slot_id` varchar(1) DEFAULT NULL,
  `day` varchar(1) DEFAULT NULL,
  `start_hr` tinyint(4) DEFAULT NULL,
  `start_min` tinyint(4) DEFAULT NULL,
  `end_hr` tinyint(4) DEFAULT NULL,
  `end_min` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `time_slot`
--

LOCK TABLES `time_slot` WRITE;
/*!40000 ALTER TABLE `time_slot` DISABLE KEYS */;
INSERT INTO `time_slot` VALUES ('A','M',8,0,8,50),('A','W',8,0,8,50),('A','F',8,0,8,50),('B','M',9,0,9,50),('B','W',9,0,9,50),('B','F',9,0,9,50),('C','M',11,0,11,50),('C','W',11,0,11,50),('C','F',11,0,11,50),('D','M',13,0,13,50),('D','W',13,0,13,50),('D','F',13,0,13,50),('E','T',10,30,11,45),('E','R',10,30,11,45),('F','T',14,30,15,45),('F','R',14,30,15,45),('G','M',16,0,16,50),('G','W',16,0,16,50),('G','F',16,0,16,50),('H','W',10,0,12,30);
/*!40000 ALTER TABLE `time_slot` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-08-22 15:20:27
