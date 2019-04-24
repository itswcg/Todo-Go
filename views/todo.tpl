{{template "base.tpl" .}}
<main>
<div class = "sidebar">
<div style = "color: #ccc; text-align: left">
<div style = "text-align: right; float: right">
<a href = "/edit"><span class ="glyphicon glyphicon-pencil"></span></a></div>
备忘
</div><br>
<textarea class = "form-control" rows ="20" placeholder = "{{ .task}}" readonly></textarea>

</div>
<div class = "container">
<button style = "text-align: right" type = 'button' class = "btn btn-primary " data-toggle = "modal"
data-target = "#myModal">添加新的任务
</button>
<div class = "modal fade" id = "myModal" tabindex= "-1" role = "dialog" aria-labelledby = "myModalLabel"
aria-hidden = "true">
<div class= "modal-dialog ">
<div class = "modal-content">
<div class = "modal-header">
<button type = "button" class = "close" data-dismiss = "modal" aria-hidden = "true">&times; </button>
<h4 class = "modal-title" id = "myModalLabel">添加</h4>
</div>
<div class = "modal-body">
<form role = "form" method = "post" action = "/add">
<div class= "form-group">
<textarea class = "form-control" id ="todo" name = "content" required
placeholder = "新的任务"></textarea>
<div class = "modal-footer">
<button type = "submit" class = "btn btn-primary">提交</button>
</div>
</div>
</form>
</div>
</div>
</div>
</div>
<div style = "float: right">
<div class = "dropdown">
<a href = "#" class= "dropdown-toggle" data-toggle = "dropdown" role = "button" aria-haspopup = "true"
aria-expanded ="false"><span class = "glyphicon glyphicon-option-vertical"></span> </a>
<ul class ="dropdown-menu dropdown-menu-right" aria-labelledby = "dLabel">
<li><a href = "/edit">编辑</a></li>
<li><a href = "#">总结</a></li>
</ul>
</div>
</div>
<div style = "color: #ccc"><h2>今天</h2></div>
{{if or .todoList_Today  .doList_Today}}
<div class = "list-group" style ="margin-top: 10px">
{{range .todoList_Today}}
{{if .is_do}}
<a class = "list-group-item list-group-item-danger" href = "/do/{{ .id}}">{{ .content }}</a>
{{ else }}
<a class = "list-group-item disabled" href = "/undo/{{ .id }}">{{ .content }}</a>
{{ end }}
{{ else}}
{{end}}
{{range .doList_Today}}
{{if not .is_do}}
<a class = "list-group-item list-group-item-danger" href = "/do/{{ .id}}">{{ .content }}</a>
{{ else }}
<a class = "list-group-item disabled" href = "/undo/{{ .id }}">{{ .content }}</a>
{{ end }}
{{else}}
{{end}}
</div>
{{else}}
<div style = "color: #ccc">
<center>请添加你的第一个任务</center>
</div>
{{end}}

<div style = "color: #ccc;"><h2>往日</h2></div>
<div class = "list-group" style = "margin-top: 10px">
{{range .todoList}}
{{if not.is_do}}
<a class = "list-group-item list-group-item-danger" href = "/do/{{ .id }}">
{{ .content }}
<div style = "text-align: right; float: right">{{substr .create_date 0 10 }}</div>
</a>
{% else %}
<a class = "list-group-item disabled" href = "/undo/{{ .id }}">
{{ .content }}
<div style = "text-align: right; float: right">{{substr .create_date 0 10 }}</div>
</a>
{{end}}
{{else}}
{{end}}
{{range .doList}}
{{if not.is_do}}
<a class = "list-group-item list-group-item-danger" href = "/do/{{ .id }}">
{{ .content }}
<div style = "text-align: right; float: right">{{substr .create_date 0 10 }}</div>
</a>
{% else %}
<a class = "list-group-item disabled" href = "/undo/{{ .id }}">
{{ .content }}
<div style = "text-align: right; float: right">{{substr .create_date 0 10 }}</div>
</a>
{{end}}
{{else}}
{{end}}
</div>

</div>
</main>
</body>
</html>