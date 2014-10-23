{{template "head.tpl"}}

{{template "top.tpl"}}
 
 
		
		
      <form method="post"   action="?dosave=" target="dosave"> 
      <table width="880" id="custom-info" style="border-top:1px solid #ccc">
		<tr>
      		<th>姓名</td><td><input type="text"  name=uname value="{{.C.Uname}}"></td>
      		<th>性别</td>
			<td><select name="sexn">
			 
					 <option value='kk' >vv</option> 
				 
      			</select>
			</td>
			<th>地址</td><td><input type="text" name="addr" value="{{.C.Address}}"></td>
		</tr>
		<tr>
      		<th>手机</td><td>
			
			<input type="hidden" name="tel" value="{{.C.Tel}}">
			<a href="/dial.php?ext=&m={{.C.Tel}}"   target="dosave"  class="htmlButton"  style="width:140px ;">拨打{{.C.Tel}}</a>
			</td>
      		<th>电话</td><td>
				<a href="/dial.php?ext=<?php echo $cfg_ml->M_LoginID; ?>&m={{.C.Tel2}}"   target="dosave"  class="htmlButton"  style="width:140px ;">拨打{{.C.Tel2}}</a>
			<input type="text" name="tel2" style="width:80px ; float:left ;" value="{{.C.Tel2}}">
			
			
			</td>
			<th>客户状态</td><td><input type="text" name="statu" disabled value=""> </td> 
		</tr>
	 
        <tr>
      		<th>导入时间</td><td> </td>      		
			<th>初次拨打</td><td> </td>
			<th>最后拨打</td><td>{{.C.LastCalltime}}</td>
		</tr>
		<tr>
			 
			<th>预约时间</th>
			<td>
					<input   style="width:90px;" class="easyui-datebox" data-options="formatter:myformatter,parser:myparser" type="text" name="nextcall_time" value="">
					<select name="nextcall_h">
						 
					</select>
			</td>   
      <th>备注</td><td colspan="3"><input type="text" name="note" value="{{.C.Note}}" size=60></td>
		</tr>
		<tr>
			<td></td><td colspan="5">
			<input  type="submit"  class=htmlButton value="保存更改"><iframe frameborder="0" scrolling="no"  name="dosave" id="dosave"  src="" width="80"  height="20"></iframe>
			<a href="?called=called0" class=htmlButton>上一个</a>
			<a href="?next=1&lastcalltimes=<?php echo $r['lastcall_time']  ; ?>"  class=htmlButton>下一个</a>
			<a href="agent_contact.php?action=add" target="_blank"   class=htmlButton>添加客户</a>
			 
				<a 
				href="../../index.php?module=Accounts&parenttab=Sales&action=DetailView&record= " 
				target="_blank"  
				class=htmlButton>ERP查看</a>
			 
			<a href="agent_contact.php?addtocrm=<?php echo $r['id'] ;?>&goback=callout/member/agent_call.php" target="_blank"   class=htmlButton>添加到ERP</a>
			 
			 <a href="#"  class=htmlButton>发短信</a> 
			 </td>
		</tr>
      </table>
	  </form>

  <table width="880"   align="center"  width="880" id="custom-info" style="border-top:1px solid #ccc">
        <tr>
                <td width="420">
                        <iframe src="gocrm_agentcallaction?cid={{.C.Id}}" width="420" height="400" id="call_action" scrolling="no"  frameborder="0"   ></iframe>
                </td>
                        <td>
                                <iframe name="faction"  id="faction"  src="gocrm_agentcalllog?cid={{.C.Id}}" width="420"   frameborder="0"   height="400"></iframe>
                </td>
        </tr>
      </table>

{{template "foot.tpl"}}
