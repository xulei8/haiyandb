<?php

mysql_connect("localhost","root","");
mysql_select_db("test");

$sql = "insert into t1 (name, type, price, dynstr) values
 ('Funny ssshirt', 'shirt', 10.0, COLUMN_CREATE(1, 'blue', 10, 'XL')) ";

 mysql_query($sql);