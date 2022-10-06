CREATE DATABASE  IF NOT EXISTS `developers` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `developers`;
-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: developers
-- ------------------------------------------------------
-- Server version	5.7.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cities`
--

DROP TABLE IF EXISTS `cities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cities` (
  `id_ciudad` int(11) NOT NULL AUTO_INCREMENT,
  `ciudad` varchar(150) NOT NULL,
  `fk_pais` int(11) DEFAULT NULL,
  `fecha_creacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_ciudad`),
  KEY `fk_pais` (`fk_pais`),
  CONSTRAINT `cities_ibfk_1` FOREIGN KEY (`fk_pais`) REFERENCES `countries` (`id_pais`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cities`
--

LOCK TABLES `cities` WRITE;
/*!40000 ALTER TABLE `cities` DISABLE KEYS */;
INSERT INTO `cities` VALUES (1,'Buenos Aires',1,'2022-10-03 22:59:17'),(2,'Santa fe',1,'2022-10-03 22:59:17'),(3,'Mendoza',1,'2022-10-03 22:59:17'),(4,'Patagonia',1,'2022-10-03 22:59:17'),(5,'New York',2,'2022-10-03 22:59:17'),(6,'California',2,'2022-10-03 22:59:17'),(7,'Texas',2,'2022-10-03 22:59:17'),(8,'Vegas',2,'2022-10-03 22:59:17'),(9,'Ontario',3,'2022-10-03 22:59:17'),(10,'Vancouver',3,'2022-10-03 23:43:35'),(11,'Salta',1,'2022-10-04 05:24:02'),(12,'Tehotihuacan',5,'2022-10-05 00:40:41'),(13,'Bocas de toro',4,'2022-10-05 00:50:16');
/*!40000 ALTER TABLE `cities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `countries`
--

DROP TABLE IF EXISTS `countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `countries` (
  `id_pais` int(11) NOT NULL AUTO_INCREMENT,
  `pais` varchar(150) NOT NULL,
  `fecha_creacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_pais`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `countries`
--

LOCK TABLES `countries` WRITE;
/*!40000 ALTER TABLE `countries` DISABLE KEYS */;
INSERT INTO `countries` VALUES (1,'Argentina','2022-10-03 22:57:39'),(2,'Usa','2022-10-03 22:57:39'),(3,'Canada','2022-10-03 22:57:39'),(4,'Panama','2022-10-03 22:57:39'),(5,'Mexico','2022-10-03 22:57:39'),(6,'Curacao','2022-10-03 22:57:39');
/*!40000 ALTER TABLE `countries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `genders`
--

DROP TABLE IF EXISTS `genders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `genders` (
  `id_gender` int(11) NOT NULL AUTO_INCREMENT,
  `gender` varchar(25) NOT NULL,
  `fecha_creacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_gender`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genders`
--

LOCK TABLES `genders` WRITE;
/*!40000 ALTER TABLE `genders` DISABLE KEYS */;
INSERT INTO `genders` VALUES (1,'Femenino','2022-10-03 22:57:06'),(2,'Masculino','2022-10-03 22:57:06');
/*!40000 ALTER TABLE `genders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lenguajes`
--

DROP TABLE IF EXISTS `lenguajes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lenguajes` (
  `id_lenguaje` int(11) NOT NULL AUTO_INCREMENT,
  `lenguaje` varchar(100) NOT NULL,
  `fecha_creacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_lenguaje`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lenguajes`
--

LOCK TABLES `lenguajes` WRITE;
/*!40000 ALTER TABLE `lenguajes` DISABLE KEYS */;
INSERT INTO `lenguajes` VALUES (1,'Python','2022-10-03 23:00:23'),(2,'Java','2022-10-03 23:00:23'),(3,'Java Script','2022-10-03 23:00:23'),(4,'.Net','2022-10-03 23:00:23'),(5,'C#','2022-10-03 23:00:23'),(6,'Go','2022-10-03 23:00:23'),(7,'Golang','2022-10-03 23:00:23'),(8,'Php','2022-10-03 23:00:23'),(9,'Ruby','2022-10-03 23:00:23'),(10,'Cobol','2022-10-05 04:09:13'),(11,'Kotlin','2022-10-05 04:23:12'),(12,'C++','2022-10-05 04:46:11');
/*!40000 ALTER TABLE `lenguajes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id_user` int(11) NOT NULL AUTO_INCREMENT,
  `nombres` varchar(80) NOT NULL,
  `apellidos` varchar(100) DEFAULT NULL,
  `age` int(11) NOT NULL,
  `fk_ciudad` int(11) DEFAULT NULL,
  `fk_genero` int(11) DEFAULT NULL,
  `fecha_creacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_user`),
  KEY `fk_ciudad` (`fk_ciudad`),
  KEY `fk_genero` (`fk_genero`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`fk_ciudad`) REFERENCES `cities` (`id_ciudad`),
  CONSTRAINT `users_ibfk_2` FOREIGN KEY (`fk_genero`) REFERENCES `genders` (`id_gender`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Daniel Esteban','Ramirez',30,1,2,'2022-10-03 23:03:08'),(2,'Daniela','Estrada',22,2,1,'2022-10-03 23:03:08'),(3,'Fabian','Barbon',25,3,2,'2022-10-03 23:03:08'),(4,'Fabiana','Cantillo',42,4,1,'2022-10-03 23:03:08'),(5,'Julieta','Chavez',34,5,1,'2022-10-03 23:03:08'),(6,'Camilo','Pico ',27,6,2,'2022-10-03 23:03:08'),(7,'Pablo Daniel','Osvaldo',35,2,2,'2022-10-05 05:48:20'),(10,'Steven','Osorio Tipan',29,5,2,'2022-10-05 23:39:00');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userxlenguaje`
--

DROP TABLE IF EXISTS `userxlenguaje`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userxlenguaje` (
  `id_usrLenguaje` int(11) NOT NULL AUTO_INCREMENT,
  `fk_lenguaje` int(11) DEFAULT NULL,
  `fk_usuario` int(11) DEFAULT NULL,
  `nivel_lenguaje` int(11) NOT NULL,
  `fecha_creacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_usrLenguaje`),
  KEY `fk_lenguaje` (`fk_lenguaje`),
  KEY `fk_usuario` (`fk_usuario`),
  CONSTRAINT `userxlenguaje_ibfk_1` FOREIGN KEY (`fk_lenguaje`) REFERENCES `lenguajes` (`id_lenguaje`),
  CONSTRAINT `userxlenguaje_ibfk_2` FOREIGN KEY (`fk_usuario`) REFERENCES `users` (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userxlenguaje`
--

LOCK TABLES `userxlenguaje` WRITE;
/*!40000 ALTER TABLE `userxlenguaje` DISABLE KEYS */;
INSERT INTO `userxlenguaje` VALUES (1,1,1,5,'2022-10-03 23:04:33'),(2,2,1,4,'2022-10-03 23:04:33'),(3,3,1,4,'2022-10-03 23:04:33'),(4,4,1,5,'2022-10-03 23:04:33'),(5,5,1,5,'2022-10-03 23:04:33'),(6,6,1,4,'2022-10-03 23:04:33'),(7,1,2,3,'2022-10-03 23:04:33'),(8,2,2,4,'2022-10-03 23:04:33'),(9,3,2,3,'2022-10-03 23:04:33'),(10,4,2,5,'2022-10-03 23:04:33'),(11,5,2,4,'2022-10-03 23:04:33'),(12,6,2,2,'2022-10-03 23:04:33'),(13,1,3,3,'2022-10-03 23:04:33'),(14,2,3,5,'2022-10-03 23:04:33'),(15,3,3,5,'2022-10-03 23:04:33'),(16,4,3,4,'2022-10-03 23:04:33'),(17,5,3,3,'2022-10-03 23:04:33'),(18,6,3,4,'2022-10-03 23:04:33'),(19,1,4,3,'2022-10-03 23:04:33'),(20,2,4,4,'2022-10-03 23:04:33'),(21,3,4,4,'2022-10-03 23:04:33'),(22,4,4,5,'2022-10-03 23:04:33'),(23,5,4,5,'2022-10-03 23:04:33'),(24,6,4,3,'2022-10-03 23:04:33');
/*!40000 ALTER TABLE `userxlenguaje` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'developers'
--

--
-- Dumping routines for database 'developers'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-05 21:05:17
