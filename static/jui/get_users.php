<?php
require_once("conn.php");
$sql = "SELECT column_json(data) as dd  FROM `app` WHERE 1;";

$rs = mysql_query($sql );
$data = array();
$i=0;
while($r = mysql_fetch_array($rs))
{
    $d = json_decode($r['dd'] ,true );

   
     $d['id']=$i ;
     $i++;
$data['rows'][]=$d;

}

$data['total']=5;

echo json_encode($data);
 