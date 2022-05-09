-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.6.5-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table sdcd_app.device
CREATE TABLE IF NOT EXISTS `devices` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE armscii8_bin DEFAULT NULL,
  `ip` varchar(50) COLLATE armscii8_bin DEFAULT NULL,
  `onDeman` tinyint(4) DEFAULT NULL,
  `protocol` varchar(50) COLLATE armscii8_bin DEFAULT NULL,
  `port` varchar(50) COLLATE armscii8_bin DEFAULT NULL,
  `username` varchar(50) COLLATE armscii8_bin DEFAULT NULL,
  `password` varchar(50) COLLATE armscii8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=armscii8 COLLATE=armscii8_bin;

-- Dumping data for table sdcd_app.device: ~2 rows (approximately)
DELETE FROM `devices`;
/*!40000 ALTER TABLE `device` DISABLE KEYS */;
INSERT INTO `devices` (`id`, `name`, `ip`, `onDeman`, `protocol`, `port`, `username`, `password`) VALUES
	(1, 'CAM 1', '192.168.1.21', 1, 'rtsp', '554', 'admin', '123456'),
	(2, 'CAM 2', '192.168.1.22', 1, 'rtsp', '554', 'admin', '123456');
/*!40000 ALTER TABLE `device` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
