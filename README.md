ngui
-------------------

Wu Haiyan MariaDB manager

I'll use jquery easy ui and mariadb dynamic columns to make a nosql data manager programe.

Tools:
-------------------
1. Jquery Easy UI 1.4.0 
2. MariaDB 10.0 
3. Beego 
4. Bee
5. LiteIDE

How to run:
-------------------
1. MariaDB user xulei, password, xulei .
2. Confige file : conf/app.config
3. Please create a table by follow SQL:

MariaDB SQL:
-------------------
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


Author:
-------------------
Email:xulei8@qq.com Beijing China .