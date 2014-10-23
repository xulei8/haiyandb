<!DOCTYPE html>
<html> 
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>文档管理 - 会员中心  </title>
<link href="/static/css/main.css" rel="stylesheet" type="text/css" />
<style type="text/css">
	center{font-size:14px;font-weight:bold;line-height:32px;text-align:left;background:#efefef;border-bottom:1px solid #ddd;text-indent:8px;color:#006600;margin-bottom:8px;}
	form tr{height:28px;line-height:28px;padding:0px 5px;}
	input[type='submit']{
	height:28px;
	width:80px;
	}
</style> 
<script type="text/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
<script >
function setnote(aa)
{
 a= $("#selectmod").val() ;
 
$("#note").val(a);
}
 
</script>
</head>
<body> 
 
 <center>业务记录</center>
 <form  action="gocrm_agentcallaction" method="post">
 <input type="hidden" name="id" value="{{.Cid}}" />
 <table>  
 	<tr>
 	
 	<td> 结果 </td><td>
 	 
 <select name="statuid">
	 
 	 <option value='1'>1</option> 
			<option value='2'>2</option> 
			<option value='3'>3</option> 

 		</select> 
 	</td>	<td>类型</td><td>
 		<select name="actiontype">
 			 
			<option value='1'>1</option> 
			<option value='2'>2</option> 
			<option value='3'>3</option> 
 	 
 		</select>  </td>
 	</tr>
 	
 	
 	
 	<td>客户分类</td><td>
  <input  type='hidden'   name='cji' value=11> 
 		<select name="cji">
			 <option value='1' >22</option> 
			 
 		</select>
		<!--隐藏字段记录当前的客户状态-->
		<input type="hidden" name="oldstatu" value="<?php echo $stus;?>"> 
		</td>
 	
 <td>成交额</td> <td><input type="text" name="nv" value="0" /></td>
			 
    
   
    <td>
 
 	</td>
 	</tr>
 	
 	
 	<tr>
 		<td>笔记</td><td colspan="3"   ><textarea id="note" name="note" style="width:280px; height:40px ;"> </textarea></td>
  
 	</tr>
 	
 	<tr>
 		<td > 选择</td><td colspan="3"   > 
 		<select  name="selectmod" id="selectmod"  onchange="setnote();">
  <option value="ddddddddsxxxxxxxxxxxxxxdddd">sssssssssasdasdssssssss</option>
 		<option value="dddddddddddd">sssssssssssssss</option>
  	</select>  </td>
  
 	</tr>
 	<tr>
 		<td> </td><td colspan="3"   ><input  type="submit" value="提交笔记"></td>
  
 	</tr>
 </table>
 
 </form>
  
  
  </body>
</html>