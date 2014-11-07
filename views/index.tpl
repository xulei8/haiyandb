<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>NewSQL 数据管理</title>
    <link rel="stylesheet" type="text/css" href="static/jui/themes/default/easyui.css">
    <link rel="stylesheet" type="text/css" href="static/jui/themes/icon.css">
  
    <script type="text/javascript" src="static/jui/jquery.min.js"></script>
    <script type="text/javascript" src="static/jui/jquery.easyui.min.js"></script>
	    <script type="text/javascript" src="static/jui/locale/easyui-lang-zh_CN.js"></script>
 
	  <style type="text/css">
        #fm{
            margin:0;
            padding:10px 30px;
        }
        .ftitle{
            font-size:14px;
            font-weight:bold;
            padding:5px 0;
            margin-bottom:10px;
            border-bottom:1px solid #ccc;
        }
        .fitem{
            margin-bottom:5px;
        }
        .fitem label{
            display:inline-block;
            width:80px;
        }
        .fitem input{
            width:160px;
        }
		.nav{ clear:both  ; display:block ; height:25px; }
		.nav  li a { font-size:12px; }
		.nav li{ list-style-type: none ; float:left ; width:80px;  }
    </style>
	
	
	
</head>
<body>
   
		<div class=nav>
<li><a href="/data_app?app=contact">客户管理</a></li>
<li><a href="/data_app?app=config">配置管理</a></li>
<li><a href="/data_app?app=modsconfig">模块管理</a></li>
<li><a href="/data_app?app=fsconfig">字段管理</a></li>



<li><a href="/demo">Demo</a></li>
				</div> 	
 	
    <table id="dg" title="{{.AppName}}管理" class="easyui-datagrid" style="width:900px;height:450px"
            url="/data_get/?modName={{.ModName}}"
            toolbar="#toolbar" pagination="true"
            rownumbers="true" fitColumns="true" singleSelect="true">
        <thead>
            <tr>
			
			 {{range .FS}}
			{{ if .HideInList}}
			
			{{else}}
	                  <th field="{{.Name}}" 
					{{if .ListWidth}}
					style=" width:{{.ListWidth}};  "
					{{else}}
					style=" min-width:50px;  "
					{{end}}
					>{{.Label}}</th>
			 {{end}}
            {{end}}
             
            </tr>
        </thead>
    </table>
	
    <div id="toolbar">
	 {{if .ToolNew}} 
        <a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-add" plain="true" onclick="newUser()">新建</a>
         {{end}}
		{{if .ToolEdit}} 
		<a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-edit" plain="true" onclick="editUser()">编辑</a>
       {{end}}
		{{if .ToolDelete}} 
		   <a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-remove" plain="true" onclick="destroyUser()">删除</a>
	   {{end}}  
 </div>
    
    <div id="dlg" class="easyui-dialog" style="width:400px;height:280px;padding:10px 20px"
            closed="true" buttons="#dlg-buttons">
        <div class="ftitle">{{.AppName}}信息</div>
        <form id="fm" method="post" novalidate>
             {{range .FS}}
	                
						 <div class="fitem {{if .HideInNew}} filedHideInNew {{end}}   {{if .HideInEdit}} filedHideInEdit {{end}}" hideInNew="{{.HideInNew}}"  HideInEdit="{{.HideInEdit}}">
			                <label>{{.Label}}</label>
			             
							    <input name="{{.Name}}" class="{{.JuiType}}"  
 id="{{.Name}}" 	 
							{{if .AddRequire }}
							required="true"
							{{end}}
							 	validType="{{.AddType}}"
							>
							</div>
            {{end}}
			
			
         
        </form>
    </div>
    <div id="dlg-buttons">
        <a href="javascript:void(0)" class="easyui-linkbutton c6" iconCls="icon-ok" onclick="saveUser()" style="width:90px">保存</a>
        <a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-cancel" onclick="javascript:$('#dlg').dialog('close')" style="width:90px">取消</a>
    </div>
	
				
		
    <script type="text/javascript">
	//$("#lang").attr("data-options",js_init_stroption1);
	
	{{range .FS}}
		{{ if .FDataOption}}
	$("#{{.Name}}").attr("data-options", {{.FDataOption}} );  
		{{end}}
	{{end}}
        var url;
        function newUser(){
            $('#dlg').dialog('open').dialog('setTitle','新建');
		
            $('#fm').form('clear');
				$('.filedHideInEdit').show()
			$('.filedHideInNew').hide()
            url = '/data_save/?modName={{.ModName}}';
        }
        function editUser(){
            var row = $('#dg').datagrid('getSelected');
            if (row){
				
                $('#dlg').dialog('open').dialog('setTitle','编辑');
                $('#fm').form('load',row);
				
				$('.filedHideInNew').show()
			$('.filedHideInEdit').hide()
                url = '/data_update/?id='+row.id + '&modName={{.ModName}}';
            }
        }
        function saveUser(){
            $('#fm').form('submit',{
                url: url,
                onSubmit: function(){
                    return $(this).form('validate');
                },
                success: function(result){
                    var result = eval('('+result+')');
                    if (result.errorMsg){
                        $.messager.show({
                            title: 'Error',
                            msg: result.errorMsg
                        });
                    } else {
                        $('#dlg').dialog('close');        // close the dialog
                        $('#dg').datagrid('reload');    // reload the user data
                    }
                }
            });
        }
        function destroyUser(){
            var row = $('#dg').datagrid('getSelected');
            if (row){
                $.messager.confirm('请确认删除','确认要删除这条记录吗？',function(r){
                    if (r){
                        $.post('/data_delete/',{id:row.id},function(result){
                            if (result.success){
                                $('#dg').datagrid('reload');    // reload the user data
                            } else {
                                $.messager.show({    // show error message
                                    title: 'Error',
                                    msg: result.errorMsg
                                });
                            }
                        },'json');
                    }
                });
            }
        }
    </script>
  
	

				
</body>
</html>
 