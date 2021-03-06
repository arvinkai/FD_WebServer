/*
SQLyog Ultimate v11.4 (64 bit)
MySQL - 5.7.3-m13 : Database - fd_webdb
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`fd_webdb` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `fd_webdb`;

/*Table structure for table `address` */

DROP TABLE IF EXISTS `address`;

CREATE TABLE `address` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `addr` varchar(1000) DEFAULT NULL,
  `postcode` varchar(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `address` */

/*Table structure for table `book` */

DROP TABLE IF EXISTS `book`;

CREATE TABLE `book` (
  `bkid` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `goodsid` bigint(20) NOT NULL DEFAULT '0',
  `name` varchar(50) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT '0',
  `addr` varchar(100) NOT NULL DEFAULT '',
  `statues` smallint(6) NOT NULL DEFAULT '0',
  `price` int(11) NOT NULL DEFAULT '0',
  `postage` int(11) NOT NULL DEFAULT '0',
  `createtime` varchar(255) NOT NULL DEFAULT '',
  `endtime` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`bkid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `book` */

insert  into `book`(`bkid`,`uid`,`goodsid`,`name`,`count`,`addr`,`statues`,`price`,`postage`,`createtime`,`endtime`) values (1,1,1,'超级无敌-甜甜圈1',1,'美丽小兔窝1',0,5,0,'2016-05-26 23:27:08',NULL),(2,1,2,'超级无敌-甜甜圈2',1,'美丽小兔窝2',1,10,0,'2016-05-26 23:37:08',NULL),(3,1,3,'超级无敌-甜甜圈3',1,'美丽小兔窝3',2,15,0,'2016-05-26 23:47:08','2016-05-26 23:55:08');

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `categoryid` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `category` */

insert  into `category`(`id`,`categoryid`,`name`) values (1,1,'糕点'),(2,2,'饼干'),(3,3,'披萨'),(4,4,'面包'),(5,5,'蛋糕'),(6,6,'饮料');

/*Table structure for table `character` */

DROP TABLE IF EXISTS `character`;

CREATE TABLE `character` (
  `uid` bigint(20) NOT NULL AUTO_INCREMENT,
  `uname` varchar(255) NOT NULL DEFAULT '',
  `nickname` varchar(20) NOT NULL DEFAULT 'default',
  `pw` varchar(255) NOT NULL DEFAULT '',
  `icon` varchar(255) DEFAULT NULL,
  `phone` varchar(255) NOT NULL DEFAULT '',
  `point` int(11) NOT NULL DEFAULT '0',
  `money` int(11) NOT NULL DEFAULT '0',
  `email` varchar(255) DEFAULT NULL,
  `id_card` varchar(255) NOT NULL DEFAULT '',
  `local` varchar(255) DEFAULT NULL,
  `create_time` varchar(255) NOT NULL DEFAULT '',
  `lastlogin_time` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT '',
  `qqnum` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

/*Data for the table `character` */

insert  into `character`(`uid`,`uname`,`nickname`,`pw`,`icon`,`phone`,`point`,`money`,`email`,`id_card`,`local`,`create_time`,`lastlogin_time`,`token`,`qqnum`) values (1,'arvin','default','123',NULL,'18217435475',0,0,'ywy_sky@foxmail.com','111111111111111111','shanghai','1463580479','1463580479',NULL,0),(2,'11111111111','Arvinkai','qqqqqqqq','','11111111111',0,0,'','','','','','',0),(5,'8E1492CF3094FF7FF738639BC4559122','蜗牛灬凯','6ED16EDE245C89F5675582AA4DC78521','','',0,0,'','','','','','',0),(6,'oRrdQt-3HGnwgBlrqmE1OUotHbfs','蜗牛凯','o7PsWfUg5OVTC_ev1eRULcooDahW9o2t0UfhcXGA-PPmfZKeaq_n7oR1tOmCK2m3tZ449hZ4bze_w4TTtLm5oxlWEqEDMwcIOZ3_nMB9UY0','','',0,0,'','','','','','',0),(7,'Ywy@qq.com','Ywysb','111111111','','',0,0,'Ywy@qq.com','','','','','',0);

/*Table structure for table `collect` */

DROP TABLE IF EXISTS `collect`;

CREATE TABLE `collect` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `goodsid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `collect` */

/*Table structure for table `discount_coupon` */

DROP TABLE IF EXISTS `discount_coupon`;

CREATE TABLE `discount_coupon` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `couponid` bigint(20) NOT NULL DEFAULT '0',
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `count` int(11) NOT NULL DEFAULT '0',
  `starttime` varchar(255) NOT NULL DEFAULT '',
  `endtime` varchar(255) NOT NULL DEFAULT '',
  `coupon_type` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `discount_coupon` */

/*Table structure for table `evaluate` */

DROP TABLE IF EXISTS `evaluate`;

CREATE TABLE `evaluate` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `goodsid` bigint(20) NOT NULL DEFAULT '0',
  `uname` varchar(20) DEFAULT NULL,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `picture_src` varchar(1000) DEFAULT NULL,
  `createtime` varchar(255) NOT NULL DEFAULT '',
  `good_star` tinyint(4) NOT NULL DEFAULT '0',
  `post_star` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `evaluate` */

/*Table structure for table `everyday_sales` */

DROP TABLE IF EXISTS `everyday_sales`;

CREATE TABLE `everyday_sales` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `goodsid` bigint(20) NOT NULL DEFAULT '0',
  `value` int(11) NOT NULL DEFAULT '0',
  `date` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `everyday_sales` */

/*Table structure for table `expense_log` */

DROP TABLE IF EXISTS `expense_log`;

CREATE TABLE `expense_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `money` bigint(20) NOT NULL DEFAULT '0',
  `date` varchar(255) NOT NULL DEFAULT '',
  `bkid` bigint(20) NOT NULL DEFAULT '0',
  `point` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `expense_log` */

/*Table structure for table `gift_log` */

DROP TABLE IF EXISTS `gift_log`;

CREATE TABLE `gift_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `posterid` int(11) NOT NULL DEFAULT '0',
  `giftid` int(11) NOT NULL DEFAULT '0',
  `giftname` varchar(255) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT '0',
  `point` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `gift_log` */

/*Table structure for table `goodsinfo` */

DROP TABLE IF EXISTS `goodsinfo`;

CREATE TABLE `goodsinfo` (
  `goodsid` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `price` int(11) NOT NULL DEFAULT '0',
  `discount` int(11) NOT NULL DEFAULT '0',
  `barcode` bigint(20) NOT NULL DEFAULT '0',
  `updatetime` varchar(255) DEFAULT NULL,
  `createtime` varchar(255) NOT NULL DEFAULT '',
  `editer` varchar(255) NOT NULL DEFAULT '',
  `ishot` int(11) NOT NULL DEFAULT '0',
  `ishave` int(11) NOT NULL DEFAULT '0',
  `count` int(11) NOT NULL DEFAULT '0',
  `content` varchar(200) DEFAULT NULL,
  `discount_endtime` varchar(255) DEFAULT NULL,
  `categoryid` int(11) DEFAULT NULL,
  `category_name` varchar(255) NOT NULL DEFAULT '',
  `couponid` bigint(20) DEFAULT NULL,
  `tourl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`goodsid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `goodsinfo` */

insert  into `goodsinfo`(`goodsid`,`name`,`price`,`discount`,`barcode`,`updatetime`,`createtime`,`editer`,`ishot`,`ishave`,`count`,`content`,`discount_endtime`,`categoryid`,`category_name`,`couponid`,`tourl`) values (1,'超级无敌-甜甜圈1',5,0,0,'1463580479','1463580479','',0,1,100,'新鲜美味 看的见的健康 看不见的信任 全进口原料 健康是吃出来的 ',NULL,1,'糕点',NULL,NULL),(2,'超级无敌-甜甜圈2',10,0,0,'1463580479','1463580479','',0,1,50,'超级好吃的超级无敌甜甜圈',NULL,2,'饼干',NULL,NULL),(3,'超级无敌-甜甜圈3',20,0,0,'1463580479','1463580479','',0,1,10,'超级好吃的超级无敌甜甜圈',NULL,3,'披萨',NULL,NULL);

/*Table structure for table `online` */

DROP TABLE IF EXISTS `online`;

CREATE TABLE `online` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `uname` varchar(255) NOT NULL DEFAULT '',
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `statues` tinyint(4) NOT NULL DEFAULT '0',
  `session_key` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `online` */

/*Table structure for table `picture` */

DROP TABLE IF EXISTS `picture`;

CREATE TABLE `picture` (
  `imgid` bigint(20) NOT NULL AUTO_INCREMENT,
  `goodsid` bigint(20) DEFAULT NULL,
  `imgsrc` varchar(1000) NOT NULL DEFAULT '',
  `type` varchar(20) DEFAULT NULL,
  `isshow` smallint(6) NOT NULL DEFAULT '0',
  `updatetime` varchar(255) DEFAULT NULL,
  `createtime` varchar(255) NOT NULL DEFAULT '',
  `tourl` varchar(255) DEFAULT NULL,
  `where` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`imgid`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

/*Data for the table `picture` */

insert  into `picture`(`imgid`,`goodsid`,`imgsrc`,`type`,`isshow`,`updatetime`,`createtime`,`tourl`,`where`) values (1,1,'../static/img/images/slider-0.JPG','slider',1,'1462889103','1462889103',NULL,'homehead'),(2,2,'../static/img/images/slider-1.JPG','slider',1,'1462889113','1462889113',NULL,'homehead'),(3,3,'../static/img/images/slider-2.JPG','slider',1,'1462889123','1462889123',NULL,'homehead'),(4,4,'../static/img/images/slider-1.JPG','slider',0,'1462889133','1462889133',NULL,'homehead'),(5,5,'../static/img/images/carousel0.jpg','carousel',1,'1462889133','1462889133',NULL,'homecarousel'),(6,1,'../static/img/images/showlist-0.JPG','showlist',1,'1462889133','1462889133','/buypage','homeshowlist'),(7,2,'../static/img/images/showlist-0.JPG','showlist',1,'1462889133','1462889133','/buypage','homeshowlist'),(8,3,'../static/img/images/showlist-0.JPG','showlist',1,'1462889133','1462889133','/buypage','homeshowlist'),(9,9,'../static/img/images/showlist-0.JPG','showlist',1,'1462889133','1462889133','/buypage','homeshowlist'),(10,1,'../static/img/images/showlist-0.JPG','shopcarlist',1,'1462889133','1462889133',NULL,'shopcar'),(11,2,'../static/img/images/showlist-0.JPG','shopcarlist',1,'1462889133','1462889133',NULL,'shopcar'),(12,3,'../static/img/images/showlist-0.JPG','shopcarlist',1,'1462889133','1462889133',NULL,'shopcar'),(13,1,'../static/img/images/slider-1.JPG','slider',1,'1462889133','1462889133',NULL,'buypage'),(14,1,'../static/img/images/slider-0.JPG','slider',1,'1462889133','1462889133',NULL,'buypage'),(15,1,'../static/img/images/showlist-0.JPG','goodsshow',1,'1462889133','1462889133','/buypage','shoppage'),(16,2,'../static/img/images/showlist-0.JPG','goodsshow',1,'1462889133','1462889133','/buypage','shoppage'),(17,3,'../static/img/images/showlist-0.JPG','goodsshow',1,'1462889133','1462889133','/buypage','shoppage');

/*Table structure for table `poster` */

DROP TABLE IF EXISTS `poster`;

CREATE TABLE `poster` (
  `posterid` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `postgift` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`posterid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `poster` */

/*Table structure for table `postergift` */

DROP TABLE IF EXISTS `postergift`;

CREATE TABLE `postergift` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `posterid` int(11) NOT NULL DEFAULT '0',
  `giftid` int(11) NOT NULL DEFAULT '0',
  `giftname` varchar(255) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT '0',
  `point` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `postergift` */

/*Table structure for table `rechage` */

DROP TABLE IF EXISTS `rechage`;

CREATE TABLE `rechage` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `orderid` bigint(20) NOT NULL DEFAULT '0',
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `value` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rechage` */

/*Table structure for table `shopcar` */

DROP TABLE IF EXISTS `shopcar`;

CREATE TABLE `shopcar` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `goodsid` bigint(20) NOT NULL DEFAULT '0',
  `count` int(11) NOT NULL DEFAULT '0',
  `create_date` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

/*Data for the table `shopcar` */

insert  into `shopcar`(`id`,`uid`,`goodsid`,`count`,`create_date`) values (17,1,1,1,'1466531269'),(18,1,2,1,'1467302039');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
