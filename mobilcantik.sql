/*
SQLyog Ultimate v11.33 (64 bit)
MySQL - 5.7.29 : Database - mobilcantik
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`mobilcantik` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `mobilcantik`;

/*Table structure for table `APP_Ad` */

DROP TABLE IF EXISTS `APP_Ad`;

CREATE TABLE `APP_Ad` (
  `intAdId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szEmail` varchar(255) NOT NULL,
  `szPassword` varchar(255) NOT NULL,
  `szTitle` varchar(255) NOT NULL,
  `szSender` varchar(255) NOT NULL,
  `szCityId` varchar(255) NOT NULL,
  `szPhoneNumber` varchar(255) NOT NULL,
  `szAnnotation` text,
  `isActive` tinyint(1) NOT NULL DEFAULT '0',
  `intAdTypeId` bigint(20) unsigned NOT NULL,
  `dtmPosted` datetime DEFAULT CURRENT_TIMESTAMP,
  `dtmConfirmation` datetime DEFAULT NULL,
  `dtmExpired` datetime DEFAULT NULL,
  `intAdStatusTypeId` bigint(20) unsigned NOT NULL,
  `intHits` int(10) unsigned NOT NULL,
  `isExpired` tinyint(1) NOT NULL DEFAULT '0',
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdId`),
  KEY `szCityId` (`szCityId`),
  KEY `intAdStatusId` (`intAdStatusTypeId`),
  KEY `intAdTypeId` (`intAdTypeId`),
  CONSTRAINT `APP_Ad_ibfk_1` FOREIGN KEY (`szCityId`) REFERENCES `GEN_City` (`szCityId`) ON UPDATE CASCADE,
  CONSTRAINT `APP_Ad_ibfk_2` FOREIGN KEY (`intAdStatusTypeId`) REFERENCES `GEN_AdStatusType` (`intAdStatusTypeId`) ON UPDATE CASCADE,
  CONSTRAINT `APP_Ad_ibfk_3` FOREIGN KEY (`intAdTypeId`) REFERENCES `GEN_AdType` (`intAdTypeId`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=latin1;

/*Data for the table `APP_Ad` */

insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (33,'naufal.kotori@outlook.com','73c66d48086be92d','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3276','089620595294','DIJUAL HONDA CRV MURAH ',1,1,'2020-06-30 23:52:25','2020-06-30 23:52:25','2020-08-29 23:52:25',3,14,0,'2020-06-10 21:50:30','2020-07-01 06:52:25');
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (34,'naufal.kotori@outlook.com','c0020c5f49f31600','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH ',1,1,'2020-06-10 20:24:57','2020-06-10 20:24:57','2020-08-09 20:24:57',3,5,0,'2020-06-10 23:06:09','2020-07-01 06:41:19');
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (36,'naufal.kotori@outlook.com','c0020c5f49f31600','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH ',0,1,'2020-06-10 21:39:28','2020-06-10 21:39:28','2020-08-09 21:39:28',3,0,0,'2020-06-10 23:06:13','2020-06-11 07:27:08');
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (37,'naufal.kotori@outlook.com','49ca185a6553a424','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH ',1,1,'2020-06-13 17:07:28','2020-06-13 17:07:28','2020-08-12 17:07:28',3,0,0,'2020-06-14 00:04:48','2020-06-14 00:07:28');
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (38,'naufal.kotori@outlook.com','b757bf0c3013504c','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH ',0,1,NULL,NULL,NULL,1,0,0,'2020-06-14 00:10:13',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (39,'naufal.kotori@outlook.com','3128c8edaaa43599','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 09:37:17',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (40,'naufal.kotori@outlook.com','1aa7b0c477722c34','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 09:38:46',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (41,'naufal.kotori@outlook.com','039eaa5bf1e18500','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 09:42:30',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (42,'naufal.kotori@outlook.com','f203ee5b1b17d4f3','Honda CRV 2.4 AT 2016','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 09:44:28',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (43,'naufal.kotori@outlook.com','0d3548613fce2e3f','Honda CRV 2.4 AT 2019','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 09:58:11',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (44,'naufal.kotori@outlook.com','4d641b2c48ea1af7','Honda CRV 2.4 AT 2019','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 09:59:11',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (45,'naufal.kotori@outlook.com','95b23ad5be9ba773','Honda CRV 2.4 AT 2019','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 10:00:44',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (46,'naufal.kotori@outlook.com','a14c75fb8b3042f1','Honda CRV 2.4 AT 2019','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,2,0,0,'2020-06-14 10:01:26',NULL);
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (47,'naufal.kotori@outlook.com','00dc255d4c04da87','Honda CRV 2.4 AT 2019','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',1,1,'2020-06-14 03:10:24','2020-06-14 03:10:24','2020-08-13 03:10:24',3,0,0,'2020-06-14 10:02:50','2020-06-14 10:10:23');
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (48,'naufal.kotori@outlook.com','3f50ee5243e31123','Honda CRV x 2.4 AT 2019','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',1,1,'2020-06-14 03:14:40','2020-06-14 03:14:40','2020-08-13 03:14:40',3,9,0,'2020-06-14 10:06:29','2020-07-01 06:35:22');
insert  into `APP_Ad`(`intAdId`,`szEmail`,`szPassword`,`szTitle`,`szSender`,`szCityId`,`szPhoneNumber`,`szAnnotation`,`isActive`,`intAdTypeId`,`dtmPosted`,`dtmConfirmation`,`dtmExpired`,`intAdStatusTypeId`,`intHits`,`isExpired`,`dtmCreatedDate`,`dtmModifiedDate`) values (49,'naufal.kotori@outlook.com','ab3fff37923e30d3','Honda CRV x 2.4 AT 2019','Muhammad Naufal Fadhillah','3275','089620595294','DIJUAL HONDA CRV MURAH  2222',0,1,NULL,NULL,NULL,1,0,0,'2020-07-01 06:31:10',NULL);

/*Table structure for table `APP_AdBill` */

DROP TABLE IF EXISTS `APP_AdBill`;

CREATE TABLE `APP_AdBill` (
  `intAdBillId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdId` bigint(20) unsigned NOT NULL,
  `decAmount` decimal(20,4) NOT NULL,
  `decBalance` decimal(20,4) NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdBillId`),
  KEY `intAdvertisementId` (`intAdId`),
  CONSTRAINT `APP_AdBill_ibfk_1` FOREIGN KEY (`intAdId`) REFERENCES `APP_Ad` (`intAdId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdBill` */

insert  into `APP_AdBill`(`intAdBillId`,`intAdId`,`decAmount`,`decBalance`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,34,'50000.0000','0.0000','2020-06-10 23:06:11','2020-06-11 03:24:57');
insert  into `APP_AdBill`(`intAdBillId`,`intAdId`,`decAmount`,`decBalance`,`dtmCreatedDate`,`dtmModifiedDate`) values (5,36,'50000.0000','0.0000','2020-06-10 23:06:13','2020-06-11 04:39:28');
insert  into `APP_AdBill`(`intAdBillId`,`intAdId`,`decAmount`,`decBalance`,`dtmCreatedDate`,`dtmModifiedDate`) values (6,37,'50000.0000','0.0000','2020-06-14 00:04:48','2020-06-14 00:07:28');
insert  into `APP_AdBill`(`intAdBillId`,`intAdId`,`decAmount`,`decBalance`,`dtmCreatedDate`,`dtmModifiedDate`) values (7,38,'50000.0000','50000.0000','2020-06-14 00:10:14',NULL);
insert  into `APP_AdBill`(`intAdBillId`,`intAdId`,`decAmount`,`decBalance`,`dtmCreatedDate`,`dtmModifiedDate`) values (8,48,'50000.0000','0.0000','2020-06-14 10:06:31','2020-06-14 10:14:40');
insert  into `APP_AdBill`(`intAdBillId`,`intAdId`,`decAmount`,`decBalance`,`dtmCreatedDate`,`dtmModifiedDate`) values (9,49,'50000.0000','50000.0000','2020-07-01 06:31:10',NULL);

/*Table structure for table `APP_AdClassified` */

DROP TABLE IF EXISTS `APP_AdClassified`;

CREATE TABLE `APP_AdClassified` (
  `intAdClassifiedId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szSender` varchar(255) NOT NULL,
  `szEmail` varchar(255) NOT NULL,
  `szPhoneNumber` varchar(255) NOT NULL,
  `intVehicleTypeId` bigint(20) unsigned NOT NULL,
  `intVehicleBrandId` bigint(20) unsigned NOT NULL,
  `intAdTypeId` bigint(20) unsigned NOT NULL,
  `szModel` varchar(255) NOT NULL,
  `intProductYear` year(4) NOT NULL,
  `decPrice` decimal(20,4) NOT NULL,
  `dtmPosted` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdClassifiedId`),
  KEY `intVehicleTypeId` (`intVehicleTypeId`),
  KEY `intVehicleBrandId` (`intVehicleBrandId`),
  KEY `intAdTypeId` (`intAdTypeId`),
  CONSTRAINT `APP_AdClassified_ibfk_1` FOREIGN KEY (`intVehicleTypeId`) REFERENCES `GEN_VehicleType` (`intVehicleTypeId`) ON UPDATE CASCADE,
  CONSTRAINT `APP_AdClassified_ibfk_2` FOREIGN KEY (`intVehicleBrandId`) REFERENCES `GEN_VehicleBrand` (`intVehicleBrandId`) ON UPDATE CASCADE,
  CONSTRAINT `APP_AdClassified_ibfk_3` FOREIGN KEY (`intAdTypeId`) REFERENCES `GEN_AdType` (`intAdTypeId`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdClassified` */

/*Table structure for table `APP_AdConfirmation` */

DROP TABLE IF EXISTS `APP_AdConfirmation`;

CREATE TABLE `APP_AdConfirmation` (
  `intAdConfirmationId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdId` bigint(20) unsigned NOT NULL,
  `intAdBillId` bigint(20) unsigned NOT NULL,
  `intBankAccountId` bigint(20) unsigned NOT NULL,
  `szBankAccountNumber` varchar(255) NOT NULL,
  `szBankName` varchar(255) NOT NULL,
  `szBankBeneficiaryName` varchar(255) NOT NULL,
  `szAnnotation` text NOT NULL,
  `dtmDate` datetime NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdConfirmationId`),
  KEY `intAdvertisementId` (`intAdId`),
  KEY `intAdvertisementBillId` (`intAdBillId`),
  KEY `intBankAccountId` (`intBankAccountId`),
  CONSTRAINT `APP_AdConfirmation_ibfk_1` FOREIGN KEY (`intAdId`) REFERENCES `APP_Ad` (`intAdId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `APP_AdConfirmation_ibfk_2` FOREIGN KEY (`intAdBillId`) REFERENCES `APP_AdBill` (`intAdBillId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `APP_AdConfirmation_ibfk_3` FOREIGN KEY (`intBankAccountId`) REFERENCES `GEN_BankAccount` (`intBankAccountId`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdConfirmation` */

insert  into `APP_AdConfirmation`(`intAdConfirmationId`,`intAdId`,`intAdBillId`,`intBankAccountId`,`szBankAccountNumber`,`szBankName`,`szBankBeneficiaryName`,`szAnnotation`,`dtmDate`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,34,3,1,'AAA','AA','AA','dasdadsadas','2020-06-11 03:02:59','2020-06-11 03:02:55',NULL);
insert  into `APP_AdConfirmation`(`intAdConfirmationId`,`intAdId`,`intAdBillId`,`intBankAccountId`,`szBankAccountNumber`,`szBankName`,`szBankBeneficiaryName`,`szAnnotation`,`dtmDate`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,36,5,1,'555555','CIMB Niaga','Penjual 1','pembayaran iklan mobilcantik','2020-06-11 04:38:40','2020-06-11 04:38:40',NULL);
insert  into `APP_AdConfirmation`(`intAdConfirmationId`,`intAdId`,`intAdBillId`,`intBankAccountId`,`szBankAccountNumber`,`szBankName`,`szBankBeneficiaryName`,`szAnnotation`,`dtmDate`,`dtmCreatedDate`,`dtmModifiedDate`) values (4,37,6,1,'555555','CIMB Niaga','Penjual 1','pembayaran iklan mobilcantik','2020-06-14 00:06:47','2020-06-14 00:06:47',NULL);
insert  into `APP_AdConfirmation`(`intAdConfirmationId`,`intAdId`,`intAdBillId`,`intBankAccountId`,`szBankAccountNumber`,`szBankName`,`szBankBeneficiaryName`,`szAnnotation`,`dtmDate`,`dtmCreatedDate`,`dtmModifiedDate`) values (5,38,7,1,'555555','CIMB Niaga','Penjual 1','pembayaran iklan mobilcantik','2020-06-14 10:19:06','2020-06-14 10:19:06',NULL);

/*Table structure for table `APP_AdConfirmationImage` */

DROP TABLE IF EXISTS `APP_AdConfirmationImage`;

CREATE TABLE `APP_AdConfirmationImage` (
  `intAdConfirmationImageId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdConfirmationId` bigint(20) unsigned NOT NULL,
  `szDirectory` varchar(255) NOT NULL,
  `szFilename` varchar(255) NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdConfirmationImageId`),
  KEY `intAdConfirmationId` (`intAdConfirmationId`),
  CONSTRAINT `APP_AdConfirmationImage_ibfk_1` FOREIGN KEY (`intAdConfirmationId`) REFERENCES `APP_AdConfirmation` (`intAdConfirmationId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdConfirmationImage` */

insert  into `APP_AdConfirmationImage`(`intAdConfirmationImageId`,`intAdConfirmationId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,1,'/sada','dsada.png','2020-06-11 03:03:13',NULL);
insert  into `APP_AdConfirmationImage`(`intAdConfirmationImageId`,`intAdConfirmationId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,3,'/2020/6/11','c95c0cf7f0ec5729.png','2020-06-11 04:38:40',NULL);
insert  into `APP_AdConfirmationImage`(`intAdConfirmationImageId`,`intAdConfirmationId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,4,'/2020/6/14','340f52c2ceb2542c.png','2020-06-14 00:06:47',NULL);
insert  into `APP_AdConfirmationImage`(`intAdConfirmationImageId`,`intAdConfirmationId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (4,5,'2020/6/14/','6dbeba4a26a90639.png','2020-06-14 10:19:06',NULL);

/*Table structure for table `APP_AdFacility` */

DROP TABLE IF EXISTS `APP_AdFacility`;

CREATE TABLE `APP_AdFacility` (
  `intAdFacilityId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdId` bigint(20) unsigned NOT NULL,
  `intFacilityId` bigint(20) unsigned NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdFacilityId`),
  KEY `APP_Advertisement_Facility_ibfk_1` (`intAdId`),
  KEY `APP_Advertisement_Facility_ibfk_2` (`intFacilityId`),
  CONSTRAINT `APP_AdFacility_ibfk_1` FOREIGN KEY (`intAdId`) REFERENCES `APP_Ad` (`intAdId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `APP_AdFacility_ibfk_2` FOREIGN KEY (`intFacilityId`) REFERENCES `GEN_Facility` (`intFacilityId`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdFacility` */

insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (57,33,1,'2020-06-10 21:50:30',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (58,33,2,'2020-06-10 21:50:30',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (59,34,1,'2020-06-10 23:06:11',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (60,34,2,'2020-06-10 23:06:11',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (63,36,1,'2020-06-10 23:06:13',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (64,36,2,'2020-06-10 23:06:13',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (65,37,1,'2020-06-14 00:04:48',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (66,37,2,'2020-06-14 00:04:48',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (67,38,1,'2020-06-14 00:10:14',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (68,38,2,'2020-06-14 00:10:14',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (69,39,1,'2020-06-14 09:37:17',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (70,39,2,'2020-06-14 09:37:17',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (71,40,1,'2020-06-14 09:38:46',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (72,40,2,'2020-06-14 09:38:46',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (73,41,1,'2020-06-14 09:42:30',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (74,41,2,'2020-06-14 09:42:30',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (75,42,1,'2020-06-14 09:44:28',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (76,42,2,'2020-06-14 09:44:28',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (77,43,1,'2020-06-14 09:58:11',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (78,43,2,'2020-06-14 09:58:11',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (79,44,1,'2020-06-14 09:59:11',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (80,44,2,'2020-06-14 09:59:11',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (81,45,1,'2020-06-14 10:00:44',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (82,45,2,'2020-06-14 10:00:44',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (83,46,1,'2020-06-14 10:01:26',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (84,46,2,'2020-06-14 10:01:26',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (85,47,1,'2020-06-14 10:02:50',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (86,47,2,'2020-06-14 10:02:50',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (87,48,1,'2020-06-14 10:06:31',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (88,48,2,'2020-06-14 10:06:31',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (89,49,1,'2020-07-01 06:31:10',NULL);
insert  into `APP_AdFacility`(`intAdFacilityId`,`intAdId`,`intFacilityId`,`dtmCreatedDate`,`dtmModifiedDate`) values (90,49,2,'2020-07-01 06:31:10',NULL);

/*Table structure for table `APP_AdImage` */

DROP TABLE IF EXISTS `APP_AdImage`;

CREATE TABLE `APP_AdImage` (
  `intAdImageId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdId` bigint(20) unsigned NOT NULL,
  `szDirectory` varchar(255) NOT NULL,
  `szFilename` varchar(255) NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdImageId`),
  KEY `APP_Advertisement_Image_ibfk_1` (`intAdId`),
  CONSTRAINT `APP_AdImage_ibfk_1` FOREIGN KEY (`intAdId`) REFERENCES `APP_Ad` (`intAdId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdImage` */

insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (33,33,'2020/6/10/','6703e939b2ca9809.png','2020-06-10 21:50:30',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (34,33,'2020/6/10/','584657da37fe8349.png','2020-06-10 21:50:30',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (35,34,'2020/6/10/','f775814f5a3ae701.png','2020-06-10 23:06:11',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (36,34,'2020/6/10/','c7b30f7d9028d7c0.png','2020-06-10 23:06:11',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (39,36,'2020/6/10/','c157de3610fe5511.png','2020-06-10 23:06:13',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (40,36,'2020/6/10/','73177e8b07109ff8.png','2020-06-10 23:06:13',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (41,37,'2020/6/14/','ffb8cff814531861.png','2020-06-14 00:04:48',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (42,37,'2020/6/14/','984dd82f78934c54.png','2020-06-14 00:04:48',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (43,38,'2020/6/14/','c7083ed74a905a54.png','2020-06-14 00:10:14',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (44,38,'2020/6/14/','32ac95425daea15b.png','2020-06-14 00:10:14',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (45,39,'2020/6/14/','3d4b5d1bd60864b9.png','2020-06-14 09:37:17',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (46,39,'2020/6/14/','8b557115c070eabe.png','2020-06-14 09:37:17',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (47,40,'2020/6/14/','1a6abd3bfe34b3b1.png','2020-06-14 09:38:46',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (48,40,'2020/6/14/','5c853c3965ac73fc.png','2020-06-14 09:38:46',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (49,41,'2020/6/14/','0cb58f04849e1ac2.png','2020-06-14 09:42:30',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (50,41,'2020/6/14/','f07fd88deb2fb3f5.png','2020-06-14 09:42:30',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (51,42,'2020/6/14/','9c52813e68ee8c5b.png','2020-06-14 09:44:28',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (52,42,'2020/6/14/','d84830f8119d9d58.png','2020-06-14 09:44:28',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (53,43,'2020/6/14/','67e8a75bed7cd784.png','2020-06-14 09:58:11',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (54,43,'2020/6/14/','1bf004ffd606aa4a.png','2020-06-14 09:58:11',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (55,44,'2020/6/14/','144cd86402222eb3.png','2020-06-14 09:59:11',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (56,44,'2020/6/14/','1ffcd26a80e1177d.png','2020-06-14 09:59:11',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (57,45,'2020/6/14/','fa669d95ff5edf43.png','2020-06-14 10:00:44',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (58,45,'2020/6/14/','36494a112a9a6050.png','2020-06-14 10:00:44',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (59,46,'2020/6/14/','d6f1e3de8e386da2.png','2020-06-14 10:01:26',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (60,46,'2020/6/14/','ad4afa7ccddbaa66.png','2020-06-14 10:01:26',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (61,47,'2020/6/14/','e1cf63e4625e8b64.png','2020-06-14 10:02:51',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (62,47,'2020/6/14/','1f0468105676c4ba.png','2020-06-14 10:02:51',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (63,48,'2020/6/14/','d603c655a1db0103.png','2020-06-14 10:06:31',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (64,48,'2020/6/14/','c579032f735f4ae8.png','2020-06-14 10:06:31',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (65,49,'2020/7/1/','ac3caee386a75006.png','2020-07-01 06:31:10',NULL);
insert  into `APP_AdImage`(`intAdImageId`,`intAdId`,`szDirectory`,`szFilename`,`dtmCreatedDate`,`dtmModifiedDate`) values (66,49,'2020/7/1/','ff41592c231f676e.png','2020-07-01 06:31:10',NULL);

/*Table structure for table `APP_AdReport` */

DROP TABLE IF EXISTS `APP_AdReport`;

CREATE TABLE `APP_AdReport` (
  `intAdReportId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdId` bigint(20) unsigned NOT NULL,
  `intReportTypeId` bigint(20) unsigned NOT NULL,
  `szAnnotation` text,
  `dtmReportTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `szReporterName` varchar(255) NOT NULL,
  `szReporterEmail` varchar(255) NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdReportId`),
  KEY `intAdId` (`intAdId`),
  KEY `intReportTypeId` (`intReportTypeId`),
  CONSTRAINT `APP_AdReport_ibfk_1` FOREIGN KEY (`intAdId`) REFERENCES `APP_Ad` (`intAdId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `APP_AdReport_ibfk_2` FOREIGN KEY (`intReportTypeId`) REFERENCES `GEN_ReportType` (`intReportTypeId`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdReport` */

insert  into `APP_AdReport`(`intAdReportId`,`intAdId`,`intReportTypeId`,`szAnnotation`,`dtmReportTime`,`szReporterName`,`szReporterEmail`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,33,4,'SPAM SPAM SPAM SPAM SPAM SPAM SPAM','2020-06-11 05:55:50','Naufal','naufal.kotori@outlook.com','2020-06-11 05:55:50',NULL);
insert  into `APP_AdReport`(`intAdReportId`,`intAdId`,`intReportTypeId`,`szAnnotation`,`dtmReportTime`,`szReporterName`,`szReporterEmail`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,33,4,'SPAM SPAM SPAM SPAM SPAM SPAM SPAM','2020-06-11 05:58:58','Naufal','naufal.kotori@outlook.com','2020-06-11 05:58:58',NULL);

/*Table structure for table `APP_AdVehicle` */

DROP TABLE IF EXISTS `APP_AdVehicle`;

CREATE TABLE `APP_AdVehicle` (
  `intAdVehicleId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdId` bigint(20) unsigned NOT NULL,
  `intVehicleTypeId` bigint(20) unsigned NOT NULL,
  `intVehicleBrandId` bigint(20) unsigned NOT NULL,
  `szType` varchar(255) NOT NULL,
  `szModel` varchar(255) NOT NULL,
  `intConditionId` bigint(20) unsigned NOT NULL,
  `intProductYear` year(4) NOT NULL,
  `intFuelId` bigint(20) unsigned NOT NULL,
  `szColor` varchar(255) NOT NULL,
  `intKilometer` int(10) unsigned NOT NULL,
  `decPrice` decimal(20,4) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdVehicleId`),
  KEY `intAdvertisementId` (`intAdId`),
  KEY `intVehicleBrandId` (`intVehicleBrandId`),
  KEY `intConditionId` (`intConditionId`),
  KEY `intVehicleTypeId` (`intVehicleTypeId`),
  KEY `APP_AdVehicle_ibfk_4` (`intFuelId`,`intVehicleTypeId`),
  CONSTRAINT `APP_AdVehicle_ibfk_1` FOREIGN KEY (`intAdId`) REFERENCES `APP_Ad` (`intAdId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `APP_AdVehicle_ibfk_2` FOREIGN KEY (`intVehicleBrandId`) REFERENCES `GEN_VehicleBrand` (`intVehicleBrandId`) ON UPDATE CASCADE,
  CONSTRAINT `APP_AdVehicle_ibfk_3` FOREIGN KEY (`intConditionId`) REFERENCES `GEN_Condition` (`intConditionId`) ON UPDATE CASCADE,
  CONSTRAINT `APP_AdVehicle_ibfk_4` FOREIGN KEY (`intFuelId`) REFERENCES `GEN_VehicleFuel` (`intFuelId`) ON UPDATE CASCADE,
  CONSTRAINT `APP_AdVehicle_ibfk_5` FOREIGN KEY (`intVehicleTypeId`) REFERENCES `GEN_VehicleType` (`intVehicleTypeId`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=latin1;

/*Data for the table `APP_AdVehicle` */

insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (32,33,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH ','2020-06-10 21:50:30',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (33,34,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH ','2020-06-10 23:06:09',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (35,36,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH ','2020-06-10 23:06:13',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (36,37,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH ','2020-06-14 00:04:48',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (37,38,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH ','2020-06-14 00:10:14',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (38,39,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 09:37:17',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (39,40,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 09:38:46',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (40,41,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 09:42:30',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (41,42,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 09:44:28',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (42,43,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 09:58:11',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (43,44,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 09:59:11',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (44,45,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 10:00:44',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (45,46,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 10:01:26',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (46,47,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 10:02:50',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (47,48,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-06-14 10:06:31',NULL);
insert  into `APP_AdVehicle`(`intAdVehicleId`,`intAdId`,`intVehicleTypeId`,`intVehicleBrandId`,`szType`,`szModel`,`intConditionId`,`intProductYear`,`intFuelId`,`szColor`,`intKilometer`,`decPrice`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (48,49,1,1,'CRV','A.T',2,2016,2,'Black',10000,'300000000.0000','DIJUAL HONDA CRV MURAH  2222','2020-07-01 06:31:10',NULL);

/*Table structure for table `APP_Testimonial` */

DROP TABLE IF EXISTS `APP_Testimonial`;

CREATE TABLE `APP_Testimonial` (
  `intTestimonialId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intAdId` bigint(20) unsigned NOT NULL,
  `szComment` varchar(255) NOT NULL,
  `intAdClosingReasonId` bigint(20) unsigned NOT NULL,
  `dtmDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `isActive` tinyint(1) NOT NULL DEFAULT '0',
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intTestimonialId`),
  KEY `intAdClosingReasonId` (`intAdClosingReasonId`),
  KEY `APP_Testimonial_ibfk_1` (`intAdId`),
  CONSTRAINT `APP_Testimonial_ibfk_1` FOREIGN KEY (`intAdId`) REFERENCES `APP_Ad` (`intAdId`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `APP_Testimonial_ibfk_2` FOREIGN KEY (`intAdClosingReasonId`) REFERENCES `GEN_AdClosingReason` (`intAdClosingReasonId`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `APP_Testimonial` */

insert  into `APP_Testimonial`(`intTestimonialId`,`intAdId`,`szComment`,`intAdClosingReasonId`,`dtmDate`,`isActive`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,33,'Alhamdulillah Laku',1,'2020-06-11 07:25:27',1,'2020-06-11 07:25:27',NULL);
insert  into `APP_Testimonial`(`intTestimonialId`,`intAdId`,`szComment`,`intAdClosingReasonId`,`dtmDate`,`isActive`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,36,'Alhamdulillah Laku 2',1,'2020-06-11 07:27:08',1,'2020-06-11 07:27:08',NULL);

/*Table structure for table `GEN_AdClosingReason` */

DROP TABLE IF EXISTS `GEN_AdClosingReason`;

CREATE TABLE `GEN_AdClosingReason` (
  `intAdClosingReasonId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szReasonName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdClosingReasonId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_AdClosingReason` */

insert  into `GEN_AdClosingReason`(`intAdClosingReasonId`,`szReasonName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'Laku Dijual',NULL,'2020-06-11 05:33:47',NULL);
insert  into `GEN_AdClosingReason`(`intAdClosingReasonId`,`szReasonName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,'Laku Disewa',NULL,'2020-06-01 04:05:50',NULL);
insert  into `GEN_AdClosingReason`(`intAdClosingReasonId`,`szReasonName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,'Salah Input',NULL,'2020-06-11 07:16:15',NULL);

/*Table structure for table `GEN_AdStatusType` */

DROP TABLE IF EXISTS `GEN_AdStatusType`;

CREATE TABLE `GEN_AdStatusType` (
  `intAdStatusTypeId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szAdStatusTypeName` varchar(255) NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdStatusTypeId`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_AdStatusType` */

insert  into `GEN_AdStatusType`(`intAdStatusTypeId`,`szAdStatusTypeName`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'WAITING FOR CONFIRM','2020-04-27 22:27:29',NULL);
insert  into `GEN_AdStatusType`(`intAdStatusTypeId`,`szAdStatusTypeName`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,'WAITING FOR TO BE SHOWN','2020-04-27 22:27:29',NULL);
insert  into `GEN_AdStatusType`(`intAdStatusTypeId`,`szAdStatusTypeName`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,'SHOWN','2020-04-27 22:27:29',NULL);
insert  into `GEN_AdStatusType`(`intAdStatusTypeId`,`szAdStatusTypeName`,`dtmCreatedDate`,`dtmModifiedDate`) values (4,'CLOSED','2020-04-27 22:27:29',NULL);

/*Table structure for table `GEN_AdType` */

DROP TABLE IF EXISTS `GEN_AdType`;

CREATE TABLE `GEN_AdType` (
  `intAdTypeId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szAdTypeName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intAdTypeId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_AdType` */

insert  into `GEN_AdType`(`intAdTypeId`,`szAdTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'Sell',NULL,'2020-06-01 04:07:06',NULL);
insert  into `GEN_AdType`(`intAdTypeId`,`szAdTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,'Rent',NULL,'2020-06-01 04:07:08',NULL);

/*Table structure for table `GEN_BankAccount` */

DROP TABLE IF EXISTS `GEN_BankAccount`;

CREATE TABLE `GEN_BankAccount` (
  `intBankAccountId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szBankAccountNumber` varchar(255) NOT NULL,
  `szBankCode` varchar(10) NOT NULL,
  `szBankAccountName` varchar(255) NOT NULL,
  `isActive` tinyint(3) unsigned NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intBankAccountId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_BankAccount` */

insert  into `GEN_BankAccount`(`intBankAccountId`,`szBankAccountNumber`,`szBankCode`,`szBankAccountName`,`isActive`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'534231311123','008','Mandiri a/n M Naufal Fadhillah',0,'2020-06-09 02:43:18','2020-06-09 02:55:17');

/*Table structure for table `GEN_City` */

DROP TABLE IF EXISTS `GEN_City`;

CREATE TABLE `GEN_City` (
  `szCityId` varchar(255) NOT NULL,
  `szProvinceId` varchar(255) NOT NULL,
  `szCityName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`szCityId`),
  KEY `szProvinceId` (`szProvinceId`),
  CONSTRAINT `GEN_City_ibfk_1` FOREIGN KEY (`szProvinceId`) REFERENCES `GEN_Province` (`szProvinceId`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `GEN_City` */

insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1101','11','KABUPATEN SIMEULUE',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1102','11','KABUPATEN ACEH SINGKIL',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1103','11','KABUPATEN ACEH SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1104','11','KABUPATEN ACEH TENGGARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1105','11','KABUPATEN ACEH TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1106','11','KABUPATEN ACEH TENGAH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1107','11','KABUPATEN ACEH BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1108','11','KABUPATEN ACEH BESAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1109','11','KABUPATEN PIDIE',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1110','11','KABUPATEN BIREUEN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1111','11','KABUPATEN ACEH UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1112','11','KABUPATEN ACEH BARAT DAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1113','11','KABUPATEN GAYO LUES',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1114','11','KABUPATEN ACEH TAMIANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1115','11','KABUPATEN NAGAN RAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1116','11','KABUPATEN ACEH JAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1117','11','KABUPATEN BENER MERIAH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1118','11','KABUPATEN PIDIE JAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1171','11','KOTA BANDA ACEH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1172','11','KOTA SABANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1173','11','KOTA LANGSA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1174','11','KOTA LHOKSEUMAWE',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1175','11','KOTA SUBULUSSALAM',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1201','12','KABUPATEN NIAS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1202','12','KABUPATEN MANDAILING NATAL',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1203','12','KABUPATEN TAPANULI SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1204','12','KABUPATEN TAPANULI TENGAH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1205','12','KABUPATEN TAPANULI UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1206','12','KABUPATEN TOBA SAMOSIR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1207','12','KABUPATEN LABUHAN BATU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1208','12','KABUPATEN ASAHAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1209','12','KABUPATEN SIMALUNGUN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1210','12','KABUPATEN DAIRI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1211','12','KABUPATEN KARO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1212','12','KABUPATEN DELI SERDANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1213','12','KABUPATEN LANGKAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1214','12','KABUPATEN NIAS SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1215','12','KABUPATEN HUMBANG HASUNDUTAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1216','12','KABUPATEN PAKPAK BHARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1217','12','KABUPATEN SAMOSIR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1218','12','KABUPATEN SERDANG BEDAGAI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1219','12','KABUPATEN BATU BARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1220','12','KABUPATEN PADANG LAWAS UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1221','12','KABUPATEN PADANG LAWAS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1222','12','KABUPATEN LABUHAN BATU SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1223','12','KABUPATEN LABUHAN BATU UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1224','12','KABUPATEN NIAS UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1225','12','KABUPATEN NIAS BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1271','12','KOTA SIBOLGA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1272','12','KOTA TANJUNG BALAI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1273','12','KOTA PEMATANG SIANTAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1274','12','KOTA TEBING TINGGI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1275','12','KOTA MEDAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1276','12','KOTA BINJAI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1277','12','KOTA PADANGSIDIMPUAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1278','12','KOTA GUNUNGSITOLI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1301','13','KABUPATEN KEPULAUAN MENTAWAI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1302','13','KABUPATEN PESISIR SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1303','13','KABUPATEN SOLOK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1304','13','KABUPATEN SIJUNJUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1305','13','KABUPATEN TANAH DATAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1306','13','KABUPATEN PADANG PARIAMAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1307','13','KABUPATEN AGAM',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1308','13','KABUPATEN LIMA PULUH KOTA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1309','13','KABUPATEN PASAMAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1310','13','KABUPATEN SOLOK SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1311','13','KABUPATEN DHARMASRAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1312','13','KABUPATEN PASAMAN BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1371','13','KOTA PADANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1372','13','KOTA SOLOK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1373','13','KOTA SAWAH LUNTO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1374','13','KOTA PADANG PANJANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1375','13','KOTA BUKITTINGGI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1376','13','KOTA PAYAKUMBUH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1377','13','KOTA PARIAMAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1401','14','KABUPATEN KUANTAN SINGINGI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1402','14','KABUPATEN INDRAGIRI HULU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1403','14','KABUPATEN INDRAGIRI HILIR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1404','14','KABUPATEN PELALAWAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1405','14','KABUPATEN S I A K',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1406','14','KABUPATEN KAMPAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1407','14','KABUPATEN ROKAN HULU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1408','14','KABUPATEN BENGKALIS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1409','14','KABUPATEN ROKAN HILIR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1410','14','KABUPATEN KEPULAUAN MERANTI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1471','14','KOTA PEKANBARU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1473','14','KOTA D U M A I',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1501','15','KABUPATEN KERINCI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1502','15','KABUPATEN MERANGIN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1503','15','KABUPATEN SAROLANGUN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1504','15','KABUPATEN BATANG HARI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1505','15','KABUPATEN MUARO JAMBI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1506','15','KABUPATEN TANJUNG JABUNG TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1507','15','KABUPATEN TANJUNG JABUNG BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1508','15','KABUPATEN TEBO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1509','15','KABUPATEN BUNGO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1571','15','KOTA JAMBI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1572','15','KOTA SUNGAI PENUH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1601','16','KABUPATEN OGAN KOMERING ULU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1602','16','KABUPATEN OGAN KOMERING ILIR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1603','16','KABUPATEN MUARA ENIM',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1604','16','KABUPATEN LAHAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1605','16','KABUPATEN MUSI RAWAS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1606','16','KABUPATEN MUSI BANYUASIN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1607','16','KABUPATEN BANYU ASIN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1608','16','KABUPATEN OGAN KOMERING ULU SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1609','16','KABUPATEN OGAN KOMERING ULU TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1610','16','KABUPATEN OGAN ILIR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1611','16','KABUPATEN EMPAT LAWANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1612','16','KABUPATEN PENUKAL ABAB LEMATANG ILIR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1613','16','KABUPATEN MUSI RAWAS UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1671','16','KOTA PALEMBANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1672','16','KOTA PRABUMULIH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1673','16','KOTA PAGAR ALAM',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1674','16','KOTA LUBUKLINGGAU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1701','17','KABUPATEN BENGKULU SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1702','17','KABUPATEN REJANG LEBONG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1703','17','KABUPATEN BENGKULU UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1704','17','KABUPATEN KAUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1705','17','KABUPATEN SELUMA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1706','17','KABUPATEN MUKOMUKO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1707','17','KABUPATEN LEBONG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1708','17','KABUPATEN KEPAHIANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1709','17','KABUPATEN BENGKULU TENGAH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1771','17','KOTA BENGKULU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1801','18','KABUPATEN LAMPUNG BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1802','18','KABUPATEN TANGGAMUS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1803','18','KABUPATEN LAMPUNG SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1804','18','KABUPATEN LAMPUNG TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1805','18','KABUPATEN LAMPUNG TENGAH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1806','18','KABUPATEN LAMPUNG UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1807','18','KABUPATEN WAY KANAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1808','18','KABUPATEN TULANGBAWANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1809','18','KABUPATEN PESAWARAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1810','18','KABUPATEN PRINGSEWU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1811','18','KABUPATEN MESUJI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1812','18','KABUPATEN TULANG BAWANG BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1813','18','KABUPATEN PESISIR BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1871','18','KOTA BANDAR LAMPUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1872','18','KOTA METRO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1901','19','KABUPATEN BANGKA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1902','19','KABUPATEN BELITUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1903','19','KABUPATEN BANGKA BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1904','19','KABUPATEN BANGKA TENGAH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1905','19','KABUPATEN BANGKA SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1906','19','KABUPATEN BELITUNG TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('1971','19','KOTA PANGKAL PINANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('2101','21','KABUPATEN KARIMUN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('2102','21','KABUPATEN BINTAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('2103','21','KABUPATEN NATUNA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('2104','21','KABUPATEN LINGGA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('2105','21','KABUPATEN KEPULAUAN ANAMBAS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('2171','21','KOTA B A T A M',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('2172','21','KOTA TANJUNG PINANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3101','31','KABUPATEN KEPULAUAN SERIBU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3171','31','KOTA JAKARTA SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3172','31','KOTA JAKARTA TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3173','31','KOTA JAKARTA PUSAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3174','31','KOTA JAKARTA BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3175','31','KOTA JAKARTA UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3201','32','KABUPATEN BOGOR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3202','32','KABUPATEN SUKABUMI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3203','32','KABUPATEN CIANJUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3204','32','KABUPATEN BANDUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3205','32','KABUPATEN GARUT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3206','32','KABUPATEN TASIKMALAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3207','32','KABUPATEN CIAMIS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3208','32','KABUPATEN KUNINGAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3209','32','KABUPATEN CIREBON',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3210','32','KABUPATEN MAJALENGKA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3211','32','KABUPATEN SUMEDANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3212','32','KABUPATEN INDRAMAYU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3213','32','KABUPATEN SUBANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3214','32','KABUPATEN PURWAKARTA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3215','32','KABUPATEN KARAWANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3216','32','KABUPATEN BEKASI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3217','32','KABUPATEN BANDUNG BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3218','32','KABUPATEN PANGANDARAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3271','32','KOTA BOGOR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3272','32','KOTA SUKABUMI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3273','32','KOTA BANDUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3274','32','KOTA CIREBON',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3275','32','KOTA BEKASI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3276','32','KOTA DEPOK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3277','32','KOTA CIMAHI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3278','32','KOTA TASIKMALAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3279','32','KOTA BANJAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3301','33','KABUPATEN CILACAP',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3302','33','KABUPATEN BANYUMAS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3303','33','KABUPATEN PURBALINGGA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3304','33','KABUPATEN BANJARNEGARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3305','33','KABUPATEN KEBUMEN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3306','33','KABUPATEN PURWOREJO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3307','33','KABUPATEN WONOSOBO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3308','33','KABUPATEN MAGELANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3309','33','KABUPATEN BOYOLALI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3310','33','KABUPATEN KLATEN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3311','33','KABUPATEN SUKOHARJO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3312','33','KABUPATEN WONOGIRI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3313','33','KABUPATEN KARANGANYAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3314','33','KABUPATEN SRAGEN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3315','33','KABUPATEN GROBOGAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3316','33','KABUPATEN BLORA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3317','33','KABUPATEN REMBANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3318','33','KABUPATEN PATI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3319','33','KABUPATEN KUDUS',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3320','33','KABUPATEN JEPARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3321','33','KABUPATEN DEMAK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3322','33','KABUPATEN SEMARANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3323','33','KABUPATEN TEMANGGUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3324','33','KABUPATEN KENDAL',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3325','33','KABUPATEN BATANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3326','33','KABUPATEN PEKALONGAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3327','33','KABUPATEN PEMALANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3328','33','KABUPATEN TEGAL',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3329','33','KABUPATEN BREBES',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3371','33','KOTA MAGELANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3372','33','KOTA SURAKARTA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3373','33','KOTA SALATIGA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3374','33','KOTA SEMARANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3375','33','KOTA PEKALONGAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3376','33','KOTA TEGAL',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3401','34','KABUPATEN KULON PROGO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3402','34','KABUPATEN BANTUL',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3403','34','KABUPATEN GUNUNG KIDUL',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3404','34','KABUPATEN SLEMAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3471','34','KOTA YOGYAKARTA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3501','35','KABUPATEN PACITAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3502','35','KABUPATEN PONOROGO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3503','35','KABUPATEN TRENGGALEK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3504','35','KABUPATEN TULUNGAGUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3505','35','KABUPATEN BLITAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3506','35','KABUPATEN KEDIRI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3507','35','KABUPATEN MALANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3508','35','KABUPATEN LUMAJANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3509','35','KABUPATEN JEMBER',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3510','35','KABUPATEN BANYUWANGI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3511','35','KABUPATEN BONDOWOSO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3512','35','KABUPATEN SITUBONDO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3513','35','KABUPATEN PROBOLINGGO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3514','35','KABUPATEN PASURUAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3515','35','KABUPATEN SIDOARJO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3516','35','KABUPATEN MOJOKERTO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3517','35','KABUPATEN JOMBANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3518','35','KABUPATEN NGANJUK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3519','35','KABUPATEN MADIUN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3520','35','KABUPATEN MAGETAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3521','35','KABUPATEN NGAWI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3522','35','KABUPATEN BOJONEGORO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3523','35','KABUPATEN TUBAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3524','35','KABUPATEN LAMONGAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3525','35','KABUPATEN GRESIK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3526','35','KABUPATEN BANGKALAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3527','35','KABUPATEN SAMPANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3528','35','KABUPATEN PAMEKASAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3529','35','KABUPATEN SUMENEP',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3571','35','KOTA KEDIRI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3572','35','KOTA BLITAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3573','35','KOTA MALANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3574','35','KOTA PROBOLINGGO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3575','35','KOTA PASURUAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3576','35','KOTA MOJOKERTO',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3577','35','KOTA MADIUN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3578','35','KOTA SURABAYA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3579','35','KOTA BATU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3601','36','KABUPATEN PANDEGLANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3602','36','KABUPATEN LEBAK',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3603','36','KABUPATEN TANGERANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3604','36','KABUPATEN SERANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3671','36','KOTA TANGERANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3672','36','KOTA CILEGON',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3673','36','KOTA SERANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('3674','36','KOTA TANGERANG SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5101','51','KABUPATEN JEMBRANA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5102','51','KABUPATEN TABANAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5103','51','KABUPATEN BADUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5104','51','KABUPATEN GIANYAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5105','51','KABUPATEN KLUNGKUNG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5106','51','KABUPATEN BANGLI',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5107','51','KABUPATEN KARANG ASEM',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5108','51','KABUPATEN BULELENG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5171','51','KOTA DENPASAR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5201','52','KABUPATEN LOMBOK BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5202','52','KABUPATEN LOMBOK TENGAH',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5203','52','KABUPATEN LOMBOK TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5204','52','KABUPATEN SUMBAWA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5205','52','KABUPATEN DOMPU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5206','52','KABUPATEN BIMA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5207','52','KABUPATEN SUMBAWA BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5208','52','KABUPATEN LOMBOK UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5271','52','KOTA MATARAM',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5272','52','KOTA BIMA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5301','53','KABUPATEN SUMBA BARAT',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5302','53','KABUPATEN SUMBA TIMUR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5303','53','KABUPATEN KUPANG',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5304','53','KABUPATEN TIMOR TENGAH SELATAN',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5305','53','KABUPATEN TIMOR TENGAH UTARA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5306','53','KABUPATEN BELU',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5307','53','KABUPATEN ALOR',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5308','53','KABUPATEN LEMBATA',NULL,'2020-06-09 01:00:24',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5309','53','KABUPATEN FLORES TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5310','53','KABUPATEN SIKKA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5311','53','KABUPATEN ENDE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5312','53','KABUPATEN NGADA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5313','53','KABUPATEN MANGGARAI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5314','53','KABUPATEN ROTE NDAO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5315','53','KABUPATEN MANGGARAI BARAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5316','53','KABUPATEN SUMBA TENGAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5317','53','KABUPATEN SUMBA BARAT DAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5318','53','KABUPATEN NAGEKEO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5319','53','KABUPATEN MANGGARAI TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5320','53','KABUPATEN SABU RAIJUA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5321','53','KABUPATEN MALAKA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('5371','53','KOTA KUPANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6101','61','KABUPATEN SAMBAS',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6102','61','KABUPATEN BENGKAYANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6103','61','KABUPATEN LANDAK',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6104','61','KABUPATEN MEMPAWAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6105','61','KABUPATEN SANGGAU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6106','61','KABUPATEN KETAPANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6107','61','KABUPATEN SINTANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6108','61','KABUPATEN KAPUAS HULU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6109','61','KABUPATEN SEKADAU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6110','61','KABUPATEN MELAWI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6111','61','KABUPATEN KAYONG UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6112','61','KABUPATEN KUBU RAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6171','61','KOTA PONTIANAK',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6172','61','KOTA SINGKAWANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6201','62','KABUPATEN KOTAWARINGIN BARAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6202','62','KABUPATEN KOTAWARINGIN TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6203','62','KABUPATEN KAPUAS',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6204','62','KABUPATEN BARITO SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6205','62','KABUPATEN BARITO UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6206','62','KABUPATEN SUKAMARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6207','62','KABUPATEN LAMANDAU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6208','62','KABUPATEN SERUYAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6209','62','KABUPATEN KATINGAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6210','62','KABUPATEN PULANG PISAU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6211','62','KABUPATEN GUNUNG MAS',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6212','62','KABUPATEN BARITO TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6213','62','KABUPATEN MURUNG RAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6271','62','KOTA PALANGKA RAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6301','63','KABUPATEN TANAH LAUT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6302','63','KABUPATEN KOTA BARU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6303','63','KABUPATEN BANJAR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6304','63','KABUPATEN BARITO KUALA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6305','63','KABUPATEN TAPIN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6306','63','KABUPATEN HULU SUNGAI SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6307','63','KABUPATEN HULU SUNGAI TENGAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6308','63','KABUPATEN HULU SUNGAI UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6309','63','KABUPATEN TABALONG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6310','63','KABUPATEN TANAH BUMBU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6311','63','KABUPATEN BALANGAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6371','63','KOTA BANJARMASIN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6372','63','KOTA BANJAR BARU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6401','64','KABUPATEN PASER',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6402','64','KABUPATEN KUTAI BARAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6403','64','KABUPATEN KUTAI KARTANEGARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6404','64','KABUPATEN KUTAI TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6405','64','KABUPATEN BERAU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6409','64','KABUPATEN PENAJAM PASER UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6411','64','KABUPATEN MAHAKAM HULU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6471','64','KOTA BALIKPAPAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6472','64','KOTA SAMARINDA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6474','64','KOTA BONTANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6501','65','KABUPATEN MALINAU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6502','65','KABUPATEN BULUNGAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6503','65','KABUPATEN TANA TIDUNG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6504','65','KABUPATEN NUNUKAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6571','65','KOTA TARAKAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7101','71','KABUPATEN BOLAANG MONGONDOW',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7102','71','KABUPATEN MINAHASA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7103','71','KABUPATEN KEPULAUAN SANGIHE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7104','71','KABUPATEN KEPULAUAN TALAUD',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7105','71','KABUPATEN MINAHASA SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7106','71','KABUPATEN MINAHASA UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7107','71','KABUPATEN BOLAANG MONGONDOW UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7108','71','KABUPATEN SIAU TAGULANDANG BIARO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7109','71','KABUPATEN MINAHASA TENGGARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7110','71','KABUPATEN BOLAANG MONGONDOW SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7111','71','KABUPATEN BOLAANG MONGONDOW TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7171','71','KOTA MANADO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7172','71','KOTA BITUNG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7173','71','KOTA TOMOHON',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7174','71','KOTA KOTAMOBAGU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7201','72','KABUPATEN BANGGAI KEPULAUAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7202','72','KABUPATEN BANGGAI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7203','72','KABUPATEN MOROWALI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7204','72','KABUPATEN POSO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7205','72','KABUPATEN DONGGALA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7206','72','KABUPATEN TOLI-TOLI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7207','72','KABUPATEN BUOL',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7208','72','KABUPATEN PARIGI MOUTONG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7209','72','KABUPATEN TOJO UNA-UNA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7210','72','KABUPATEN SIGI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7211','72','KABUPATEN BANGGAI LAUT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7212','72','KABUPATEN MOROWALI UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7271','72','KOTA PALU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7301','73','KABUPATEN KEPULAUAN SELAYAR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7302','73','KABUPATEN BULUKUMBA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7303','73','KABUPATEN BANTAENG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7304','73','KABUPATEN JENEPONTO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7305','73','KABUPATEN TAKALAR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7306','73','KABUPATEN GOWA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7307','73','KABUPATEN SINJAI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7308','73','KABUPATEN MAROS',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7309','73','KABUPATEN PANGKAJENE DAN KEPULAUAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7310','73','KABUPATEN BARRU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7311','73','KABUPATEN BONE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7312','73','KABUPATEN SOPPENG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7313','73','KABUPATEN WAJO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7314','73','KABUPATEN SIDENRENG RAPPANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7315','73','KABUPATEN PINRANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7316','73','KABUPATEN ENREKANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7317','73','KABUPATEN LUWU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7318','73','KABUPATEN TANA TORAJA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7322','73','KABUPATEN LUWU UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7325','73','KABUPATEN LUWU TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7326','73','KABUPATEN TORAJA UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7371','73','KOTA MAKASSAR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7372','73','KOTA PAREPARE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7373','73','KOTA PALOPO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7401','74','KABUPATEN BUTON',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7402','74','KABUPATEN MUNA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7403','74','KABUPATEN KONAWE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7404','74','KABUPATEN KOLAKA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7405','74','KABUPATEN KONAWE SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7406','74','KABUPATEN BOMBANA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7407','74','KABUPATEN WAKATOBI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7408','74','KABUPATEN KOLAKA UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7409','74','KABUPATEN BUTON UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7410','74','KABUPATEN KONAWE UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7411','74','KABUPATEN KOLAKA TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7412','74','KABUPATEN KONAWE KEPULAUAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7413','74','KABUPATEN MUNA BARAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7414','74','KABUPATEN BUTON TENGAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7415','74','KABUPATEN BUTON SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7471','74','KOTA KENDARI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7472','74','KOTA BAUBAU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7501','75','KABUPATEN BOALEMO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7502','75','KABUPATEN GORONTALO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7503','75','KABUPATEN POHUWATO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7504','75','KABUPATEN BONE BOLANGO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7505','75','KABUPATEN GORONTALO UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7571','75','KOTA GORONTALO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7601','76','KABUPATEN MAJENE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7602','76','KABUPATEN POLEWALI MANDAR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7603','76','KABUPATEN MAMASA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7604','76','KABUPATEN MAMUJU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7605','76','KABUPATEN MAMUJU UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('7606','76','KABUPATEN MAMUJU TENGAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8101','81','KABUPATEN MALUKU TENGGARA BARAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8102','81','KABUPATEN MALUKU TENGGARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8103','81','KABUPATEN MALUKU TENGAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8104','81','KABUPATEN BURU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8105','81','KABUPATEN KEPULAUAN ARU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8106','81','KABUPATEN SERAM BAGIAN BARAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8107','81','KABUPATEN SERAM BAGIAN TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8108','81','KABUPATEN MALUKU BARAT DAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8109','81','KABUPATEN BURU SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8171','81','KOTA AMBON',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8172','81','KOTA TUAL',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8201','82','KABUPATEN HALMAHERA BARAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8202','82','KABUPATEN HALMAHERA TENGAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8203','82','KABUPATEN KEPULAUAN SULA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8204','82','KABUPATEN HALMAHERA SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8205','82','KABUPATEN HALMAHERA UTARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8206','82','KABUPATEN HALMAHERA TIMUR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8207','82','KABUPATEN PULAU MOROTAI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8208','82','KABUPATEN PULAU TALIABU',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8271','82','KOTA TERNATE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('8272','82','KOTA TIDORE KEPULAUAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9101','91','KABUPATEN FAKFAK',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9102','91','KABUPATEN KAIMANA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9103','91','KABUPATEN TELUK WONDAMA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9104','91','KABUPATEN TELUK BINTUNI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9105','91','KABUPATEN MANOKWARI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9106','91','KABUPATEN SORONG SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9107','91','KABUPATEN SORONG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9108','91','KABUPATEN RAJA AMPAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9109','91','KABUPATEN TAMBRAUW',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9110','91','KABUPATEN MAYBRAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9111','91','KABUPATEN MANOKWARI SELATAN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9112','91','KABUPATEN PEGUNUNGAN ARFAK',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9171','91','KOTA SORONG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9401','94','KABUPATEN MERAUKE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9402','94','KABUPATEN JAYAWIJAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9403','94','KABUPATEN JAYAPURA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9404','94','KABUPATEN NABIRE',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9408','94','KABUPATEN KEPULAUAN YAPEN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9409','94','KABUPATEN BIAK NUMFOR',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9410','94','KABUPATEN PANIAI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9411','94','KABUPATEN PUNCAK JAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9412','94','KABUPATEN MIMIKA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9413','94','KABUPATEN BOVEN DIGOEL',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9414','94','KABUPATEN MAPPI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9415','94','KABUPATEN ASMAT',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9416','94','KABUPATEN YAHUKIMO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9417','94','KABUPATEN PEGUNUNGAN BINTANG',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9418','94','KABUPATEN TOLIKARA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9419','94','KABUPATEN SARMI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9420','94','KABUPATEN KEEROM',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9426','94','KABUPATEN WAROPEN',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9427','94','KABUPATEN SUPIORI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9428','94','KABUPATEN MAMBERAMO RAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9429','94','KABUPATEN NDUGA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9430','94','KABUPATEN LANNY JAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9431','94','KABUPATEN MAMBERAMO TENGAH',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9432','94','KABUPATEN YALIMO',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9433','94','KABUPATEN PUNCAK',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9434','94','KABUPATEN DOGIYAI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9435','94','KABUPATEN INTAN JAYA',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9436','94','KABUPATEN DEIYAI',NULL,'2020-06-09 01:00:25',NULL);
insert  into `GEN_City`(`szCityId`,`szProvinceId`,`szCityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('9471','94','KOTA JAYAPURA',NULL,'2020-06-09 01:00:25',NULL);

/*Table structure for table `GEN_Condition` */

DROP TABLE IF EXISTS `GEN_Condition`;

CREATE TABLE `GEN_Condition` (
  `intConditionId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szConditionName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intConditionId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_Condition` */

insert  into `GEN_Condition`(`intConditionId`,`szConditionName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'New',NULL,'2020-04-27 22:29:40',NULL);
insert  into `GEN_Condition`(`intConditionId`,`szConditionName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,'Used','','2020-06-09 00:08:48',NULL);

/*Table structure for table `GEN_Facility` */

DROP TABLE IF EXISTS `GEN_Facility`;

CREATE TABLE `GEN_Facility` (
  `intFacilityId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intVehicleTypeId` bigint(20) unsigned NOT NULL,
  `szFacilityName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intFacilityId`),
  KEY `intVehicleTypeId` (`intVehicleTypeId`),
  CONSTRAINT `GEN_Facility_ibfk_1` FOREIGN KEY (`intVehicleTypeId`) REFERENCES `GEN_VehicleType` (`intVehicleTypeId`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_Facility` */

insert  into `GEN_Facility`(`intFacilityId`,`intVehicleTypeId`,`szFacilityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,1,'A/C','','2020-06-09 00:48:15','2020-06-09 00:51:23');
insert  into `GEN_Facility`(`intFacilityId`,`intVehicleTypeId`,`szFacilityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,1,'Keyless Ignition System','','2020-06-09 00:49:29',NULL);
insert  into `GEN_Facility`(`intFacilityId`,`intVehicleTypeId`,`szFacilityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,2,'Keyless Ignition System','','2020-06-09 00:49:35',NULL);
insert  into `GEN_Facility`(`intFacilityId`,`intVehicleTypeId`,`szFacilityName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (4,2,'Shockbreaker Upside Down','','2020-06-09 00:50:07',NULL);

/*Table structure for table `GEN_Fuel` */

DROP TABLE IF EXISTS `GEN_Fuel`;

CREATE TABLE `GEN_Fuel` (
  `intFuelId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szFuelName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intFuelId`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_Fuel` */

insert  into `GEN_Fuel`(`intFuelId`,`szFuelName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'Premium','Si Kuning','2020-04-27 22:30:37','2020-06-08 22:44:06');
insert  into `GEN_Fuel`(`intFuelId`,`szFuelName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,'Pertamax',NULL,'2020-04-27 22:30:37',NULL);
insert  into `GEN_Fuel`(`intFuelId`,`szFuelName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,'Solar',NULL,'2020-04-27 22:30:37',NULL);
insert  into `GEN_Fuel`(`intFuelId`,`szFuelName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (4,'Pertalite',NULL,'2020-04-27 22:30:37',NULL);
insert  into `GEN_Fuel`(`intFuelId`,`szFuelName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (5,'Pertamax Turbo',NULL,'2020-04-27 22:30:37',NULL);

/*Table structure for table `GEN_Partner` */

DROP TABLE IF EXISTS `GEN_Partner`;

CREATE TABLE `GEN_Partner` (
  `intPartnerId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szTitle` varchar(255) NOT NULL,
  `szUrl` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intPartnerId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_Partner` */

insert  into `GEN_Partner`(`intPartnerId`,`szTitle`,`szUrl`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'LyndaERP','lyndaerp.id','','2020-06-09 04:19:45','2020-06-14 01:16:59');

/*Table structure for table `GEN_Province` */

DROP TABLE IF EXISTS `GEN_Province`;

CREATE TABLE `GEN_Province` (
  `szProvinceId` varchar(255) NOT NULL,
  `szProvinceName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`szProvinceId`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `GEN_Province` */

insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('11','ACEH',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('12','SUMATERA UTARA',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('13','SUMATERA BARAT',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('14','RIAU',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('15','JAMBI',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('16','SUMATERA SELATAN',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('17','BENGKULU',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('18','LAMPUNG',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('19','KEPULAUAN BANGKA BELITUNG',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('21','KEPULAUAN RIAU',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('31','DKI JAKARTA',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('32','JAWA BARAT',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('33','JAWA TENGAH',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('34','DI YOGYAKARTA',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('35','JAWA TIMUR',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('36','BANTEN',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('51','BALI',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('52','NUSA TENGGARA BARAT',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('53','NUSA TENGGARA TIMUR',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('61','KALIMANTAN BARAT',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('62','KALIMANTAN TENGAH',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('63','KALIMANTAN SELATAN',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('64','KALIMANTAN TIMUR',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('65','KALIMANTAN UTARA',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('71','SULAWESI UTARA',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('72','SULAWESI TENGAH',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('73','SULAWESI SELATAN',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('74','SULAWESI TENGGARA',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('75','GORONTALO',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('76','SULAWESI BARAT',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('81','MALUKU',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('82','MALUKU UTARA',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('91','PAPUA BARAT',NULL,'2020-06-09 00:55:32',NULL);
insert  into `GEN_Province`(`szProvinceId`,`szProvinceName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('94','PAPUA',NULL,'2020-06-09 00:55:32',NULL);

/*Table structure for table `GEN_ReportGroup` */

DROP TABLE IF EXISTS `GEN_ReportGroup`;

CREATE TABLE `GEN_ReportGroup` (
  `intReportGroupId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szReportGroupName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intReportGroupId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_ReportGroup` */

insert  into `GEN_ReportGroup`(`intReportGroupId`,`szReportGroupName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'Indikasi Penipuan',NULL,'2020-06-02 18:59:20',NULL);
insert  into `GEN_ReportGroup`(`intReportGroupId`,`szReportGroupName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,'Spam Produk',NULL,'2020-06-02 18:59:49',NULL);
insert  into `GEN_ReportGroup`(`intReportGroupId`,`szReportGroupName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,'Produk Dewasa / Pornografi',NULL,'2020-06-02 19:00:19',NULL);

/*Table structure for table `GEN_ReportType` */

DROP TABLE IF EXISTS `GEN_ReportType`;

CREATE TABLE `GEN_ReportType` (
  `intReportTypeId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intReportGroupId` bigint(20) unsigned NOT NULL,
  `szReportTypeName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intReportTypeId`),
  KEY `intReportGroupId` (`intReportGroupId`),
  CONSTRAINT `GEN_ReportType_ibfk_1` FOREIGN KEY (`intReportGroupId`) REFERENCES `GEN_ReportGroup` (`intReportGroupId`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_ReportType` */

insert  into `GEN_ReportType`(`intReportTypeId`,`intReportGroupId`,`szReportTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,1,'Harga Tidak Wajar',NULL,'2020-06-02 19:08:25',NULL);
insert  into `GEN_ReportType`(`intReportTypeId`,`intReportGroupId`,`szReportTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,1,'Indikasi Manipulasi',NULL,'2020-06-02 19:08:37',NULL);
insert  into `GEN_ReportType`(`intReportTypeId`,`intReportGroupId`,`szReportTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,1,'Harga tidak sesuai pasaran',NULL,'2020-06-02 19:08:44',NULL);
insert  into `GEN_ReportType`(`intReportTypeId`,`intReportGroupId`,`szReportTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (4,2,'Spam produk yang sama',NULL,'2020-06-02 19:09:09',NULL);
insert  into `GEN_ReportType`(`intReportTypeId`,`intReportGroupId`,`szReportTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (5,3,'Gambar produk sensual',NULL,'2020-06-02 19:09:23',NULL);
insert  into `GEN_ReportType`(`intReportTypeId`,`intReportGroupId`,`szReportTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (6,3,'Mengandung unsur pornografi',NULL,'2020-06-02 19:09:32',NULL);

/*Table structure for table `GEN_VehicleBrand` */

DROP TABLE IF EXISTS `GEN_VehicleBrand`;

CREATE TABLE `GEN_VehicleBrand` (
  `intVehicleBrandId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `intVehicleTypeId` bigint(20) unsigned NOT NULL,
  `szBrandName` varchar(255) NOT NULL,
  `intCount` bigint(20) NOT NULL DEFAULT '0',
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intVehicleBrandId`),
  KEY `intVehicleTypeId` (`intVehicleTypeId`),
  CONSTRAINT `GEN_VehicleBrand_ibfk_1` FOREIGN KEY (`intVehicleTypeId`) REFERENCES `GEN_VehicleType` (`intVehicleTypeId`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_VehicleBrand` */

insert  into `GEN_VehicleBrand`(`intVehicleBrandId`,`intVehicleTypeId`,`szBrandName`,`intCount`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,1,'Honda',0,NULL,'2020-06-08 18:54:36',NULL);
insert  into `GEN_VehicleBrand`(`intVehicleBrandId`,`intVehicleTypeId`,`szBrandName`,`intCount`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,1,'Mercedes Benz',0,NULL,'2020-06-08 18:54:54',NULL);
insert  into `GEN_VehicleBrand`(`intVehicleBrandId`,`intVehicleTypeId`,`szBrandName`,`intCount`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (3,2,'Yamaha',0,'Motor Yamaha','2020-06-08 18:54:57','2020-06-08 20:26:52');
insert  into `GEN_VehicleBrand`(`intVehicleBrandId`,`intVehicleTypeId`,`szBrandName`,`intCount`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (4,2,'Honda',0,NULL,'2020-06-08 18:55:01',NULL);
insert  into `GEN_VehicleBrand`(`intVehicleBrandId`,`intVehicleTypeId`,`szBrandName`,`intCount`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (10,1,'Daihatsu',0,'','2020-06-08 20:25:32',NULL);
insert  into `GEN_VehicleBrand`(`intVehicleBrandId`,`intVehicleTypeId`,`szBrandName`,`intCount`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (12,1,'Toyota',0,'','2020-06-08 20:37:11',NULL);
insert  into `GEN_VehicleBrand`(`intVehicleBrandId`,`intVehicleTypeId`,`szBrandName`,`intCount`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (14,1,'Wuling',0,'','2020-06-08 20:39:25',NULL);

/*Table structure for table `GEN_VehicleFuel` */

DROP TABLE IF EXISTS `GEN_VehicleFuel`;

CREATE TABLE `GEN_VehicleFuel` (
  `intVehicleTypeId` bigint(20) unsigned NOT NULL,
  `intFuelId` bigint(20) unsigned NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intVehicleTypeId`,`intFuelId`),
  KEY `intFuelTypeId` (`intFuelId`),
  CONSTRAINT `GEN_VehicleFuel_ibfk_1` FOREIGN KEY (`intVehicleTypeId`) REFERENCES `GEN_VehicleType` (`intVehicleTypeId`) ON UPDATE CASCADE,
  CONSTRAINT `GEN_VehicleFuel_ibfk_2` FOREIGN KEY (`intFuelId`) REFERENCES `GEN_Fuel` (`intFuelId`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `GEN_VehicleFuel` */

insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,1,'2020-06-08 23:09:20',NULL);
insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,2,'2020-06-08 23:09:20',NULL);
insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,3,'2020-06-08 23:09:20',NULL);
insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,4,'2020-06-08 23:09:20',NULL);
insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,1,'2020-04-27 22:31:47',NULL);
insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,2,'2020-04-27 22:31:47',NULL);
insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,4,'2020-04-27 22:31:47',NULL);
insert  into `GEN_VehicleFuel`(`intVehicleTypeId`,`intFuelId`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,5,'2020-04-27 22:31:47',NULL);

/*Table structure for table `GEN_VehicleType` */

DROP TABLE IF EXISTS `GEN_VehicleType`;

CREATE TABLE `GEN_VehicleType` (
  `intVehicleTypeId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szVehicleTypeName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intVehicleTypeId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `GEN_VehicleType` */

insert  into `GEN_VehicleType`(`intVehicleTypeId`,`szVehicleTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'Car',NULL,'2020-04-27 22:32:10','2020-06-08 20:17:42');
insert  into `GEN_VehicleType`(`intVehicleTypeId`,`szVehicleTypeName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values (2,'Motorcycle',NULL,'2020-04-27 22:32:10',NULL);

/*Table structure for table `SYS_ApiKey` */

DROP TABLE IF EXISTS `SYS_ApiKey`;

CREATE TABLE `SYS_ApiKey` (
  `szApiKeyId` varchar(255) NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`szApiKeyId`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `SYS_ApiKey` */

insert  into `SYS_ApiKey`(`szApiKeyId`,`dtmCreatedDate`,`dtmModifiedDate`) values ('6a4354349fb911eaadcb08002722efb6','2020-05-27 08:28:57',NULL);
insert  into `SYS_ApiKey`(`szApiKeyId`,`dtmCreatedDate`,`dtmModifiedDate`) values ('dsfsd','2020-06-06 18:33:52',NULL);

/*Table structure for table `SYS_Role` */

DROP TABLE IF EXISTS `SYS_Role`;

CREATE TABLE `SYS_Role` (
  `szRoleId` varchar(100) NOT NULL,
  `szRoleName` varchar(255) NOT NULL,
  `szAnnotation` text,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`szRoleId`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `SYS_Role` */

insert  into `SYS_Role`(`szRoleId`,`szRoleName`,`szAnnotation`,`dtmCreatedDate`,`dtmModifiedDate`) values ('admin','Admin',NULL,'2020-05-27 08:27:14',NULL);

/*Table structure for table `SYS_User` */

DROP TABLE IF EXISTS `SYS_User`;

CREATE TABLE `SYS_User` (
  `intUserId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szUsername` varchar(255) NOT NULL,
  `szFullname` varchar(255) NOT NULL,
  `szEmail` varchar(255) NOT NULL,
  `szPassword` varchar(255) NOT NULL,
  `isDefault` tinyint(1) NOT NULL DEFAULT '0',
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intUserId`),
  UNIQUE KEY `szUsername` (`szUsername`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `SYS_User` */

insert  into `SYS_User`(`intUserId`,`szUsername`,`szFullname`,`szEmail`,`szPassword`,`isDefault`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'admin','Admin','naufal.kotori@outlook.com','$2a$10$NXciFtl4a.Ru2GCbC5ZNvOtPfd8Ogx8.z4zIMHHenPzwLNpECrPk6',1,'2020-05-27 08:30:06',NULL);

/*Table structure for table `SYS_UserRole` */

DROP TABLE IF EXISTS `SYS_UserRole`;

CREATE TABLE `SYS_UserRole` (
  `intUserId` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `szRoleId` varchar(100) NOT NULL,
  `dtmCreatedDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `dtmModifiedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`intUserId`,`szRoleId`),
  KEY `szRoleId` (`szRoleId`),
  CONSTRAINT `SYS_UserRole_ibfk_1` FOREIGN KEY (`intUserId`) REFERENCES `SYS_User` (`intUserId`) ON UPDATE CASCADE,
  CONSTRAINT `SYS_UserRole_ibfk_2` FOREIGN KEY (`szRoleId`) REFERENCES `SYS_Role` (`szRoleId`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `SYS_UserRole` */

insert  into `SYS_UserRole`(`intUserId`,`szRoleId`,`dtmCreatedDate`,`dtmModifiedDate`) values (1,'admin','2020-05-28 05:23:40',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
