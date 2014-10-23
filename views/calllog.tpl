 {{template "head.tpl"}}
 <table >
                        <thead>
                            <tr>
                                <th>Id</th>
                                <th>a</th>
                                <th>s</th>
                              
                            </tr>
                        </thead>
                        <tbody>
                            {{range $article := .Objects}}
                            <tr>
                                <td>{{$article.Id}}</a></td>
                    
                                <td><a  target=_blank href="gocrm_agentcall?cid={{$article.Id}}">{{$article.Uname }}</a></td>
                               <td>{{$article.StatuID }}</td>
							
							  <td>{{$article.Notes}}</td>
							 
							  <td>{{$article.Addtime}}</td>
                            </tr>
                            {{end}}
                        </tbody>
                    </table>
