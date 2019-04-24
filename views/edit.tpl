{{template "base.tpl" .}}
<main>
<div class = "sidebar">
<form action = "/edit" method = "post" role = "form">
<div style = "color: #ccc; text-align: left">
<div style ="text-align: right; float: right">
<button type = "submit" class = "btn btn-success btn-sm">提交</button>
</div>
备忘
</div>
<br>

<textarea class = "form-control" id = "task" name = "content" rows = "20">{{ .task }}</textarea>
</form>
</div>
<div class ="container">
<div style = "float: right; color: blue"><a href = "/"><span class = "glyphicon glyphicon-ok
glyphicon glyphicon-"></span></a></div>
<div style = "overflow: hidden; margin-bottom: 20px">
<ul class = "nav nav-pills">
<li role= "presentation" class = "active"><a href ="#">编辑</a></li>
<li role = "presentation"><a href= "#">总结</a></li>
</ul>
</div>

<ul class = "list-group">
{{range .todoList}}
<li todo-id = "{{ .id }}" class = "list-group-item" style = "color: #8590a6">{{substr .create_date 0 10 }}
&nbsp; |&nbsp; {{ .content }}
<span class= "glyphicon glyphicon-remove remove-todo" style="display:none"></span>
&nbsp; <span class = "glyphicon glyphicon-edit edit-todo" style="display:none" data-toggle = "modal" data-target = "#{{ .id }}"></span>
<div class = "modal fade" id = "{{ .id }}" tabindex ="-1" role = "dialog" aria-labelledby = "myModalLabel"
aria-hidden = "true">
<div class ="modal-dialog ">
<div class = "modal-content">
<div class = "modal-header">
<button type = "button" class= "close" data-dismiss = "modal" aria-label= "Close"><span
aria-hidden ="true">&times; </span></button>
<h4 class = "modal-title" id ="myModalLabel">编辑</h4>
</div>
<div class = "modal-body">
<form role= "form" method = "post" action = "/update/{{ .id }}">
<div class = "form-group">
<textarea class = "form-control" id = "todo" name = "content" required>{{ .content }}</textarea>
<div class = "modal-footer">
<button type = "submit" class = "btn btn-primary">修改</button>
</div>
</div>
</form>
</div>
</div>
</div>
</div>
</li>
{{else}}
{{end}}
</ul>
</div>
</main>
</body>
</html>