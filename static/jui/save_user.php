<?php

require_once("conn.php");

$data = json_encode($_POST);
$datas = $_POST;
$st = "";
$apptitle = "";
foreach ($datas as $k=>$v) {
	$st .= "'{$k}','{$v}',";
	if($apptitle=="") 
		$apptitle = $v ; 
}

$st  = substr($st, 0, strlen($st)-1);
  
$d = date('Y-m-d H:i:s');
  $sql = "INSERT INTO 
app (  `createtime`, `edittime`, `title`, `type`, `data`, `createorid`, `ownerid`)

VALUES (  '{$d}', NULL, '{$apptitle}', 'app1', COLUMN_CREATE({$st}), 1, 1);";

 mysql_query($sql);

 $s = true ;
 echo json_encode($s);