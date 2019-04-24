{{template "base.tpl" .}}
<main>
<div class = "sidebar">
<center><h3>搜索结果</h3></center>
</div>
<div class= "container">
<ul class = "list-group">
{{range .todos}}
<li todo-id = "{{ .id }}" class ="list-group-item" style = "color: #8590a6">{{substr .create_date 0 10 }}
&nbsp; |&nbsp; {{ .content }}
<span class = "glyphicon glyphicon-remove remove-todo"></span>
&nbsp; <span class = "glyphicon glyphicon-edit edit-todo" data-toggle = "modal" data-target = "#{{ .id }}"></span>
<div class = "modal fade" id = "{{ .id }}" tabindex = "-1" role = "dialog" aria-labelledby = "myModalLabel"
aria-hidden = "true">
<div class = "modal-dialog ">
<div class= "modal-content">
<div class = "modal-header">
<button type = "button" class = "close" data-dismiss= "modal" aria-label = "Close"><span
aria-hidden = "true">&times;</span></button>
<h4 class = "modal-title" id = "myModalLabel">编辑</h4>
</div>
<div class ="modal-body">
<form role = "form" method = "post" action = "/update/{{ .id }}">
<div class = "form-group">
<textarea class= "form-control" id = "todo" name = "content" required>{{ .content }}</textarea>
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
{{ end }}
</ul>
</div>
</main>
</body>
</html>
