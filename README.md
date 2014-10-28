haiyandb
=====

Wu Haiyan MariaDB manager

I'll use jquery easy ui and mariadb dynamic columns to make a nosql data manager programe.

Jquery Easy UI 1.4.0 
MariaDB 10.0 
Beego 
Bee
LiteIDE

MariaDB SQL:
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `createtime` datetime DEFAULT NULL,
  `edittime` datetime DEFAULT NULL,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `type` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `data` mediumblob,
  `createorid` int(11) DEFAULT NULL,
  `ownerid` int(11) DEFAULT NULL,
  `deleted` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=latin1;
/* Oct 28th*/