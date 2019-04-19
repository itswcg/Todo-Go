<!DOCTYPE html>
<html lang = "zh-CN">
    <head>
    <meta charset = "utf-8">
    <meta http-equiv = "X-UA-Compatible" content = "IE=edge">
    <meta name = "viewport" content = "width=device-width, initial-scale=1">
    <title>Todo</title>
    <link rel= "icon" type = "image/x-icon" href = "/static/img/fav.jpg">
    <link href = "/static/css/todo.css" rel = "stylesheet">
    <link href = "/static/css/bootstrap.css" rel = "stylesheet">
    <script src= "https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
    <script src= "/static/js/bootstrap.js"></script>
    </head>
    <body>
        <header>
            <nav class= "navbar navbar-default navbar-fixed-top" role = "navigation">
            <div class ="container-fluid">
            <div class = "navbar-header">
            <button type = "button" class = "navbar-toggle" data-toggle ="collapse" data-target = "#menu">
            <span class= "sr-only">Toggle navigation</span>
            <span class = "icon-bar"></span>
            <span class = "icon-bar"></span>
            <span class = "icon-bar"></span></button>
            <a class = "navbar-brand" href ="/">Todo</a></div>

            <div class = "collapse navbar-collapse" id = "menu">
                <form class = "navbar-form navbar-left" role = "search" action ="/search" method = "post">
                <input type = "text" class = "form-control" name = "value" placeholder = "">
                <button type = "submit"><span class = "glyphicon glyphicon-search "></span></button>
                </form>

            <ul class = "nav navbar-nav navbar-right">
            {{ if.isLogin }}

            <li class = "dropdown">
            <a href = "#" class = "dropdown-toggle" data-toggle = "dropdown" role = "button"
            aria-haspopup = "true" aria-expanded = "false"><span
            class = "glyphicon glyphicon-plus"></span> <span class = "caret"></span></a>
            <ul class = "dropdown-menu" aria-labelledby = "dLabel">
            <li><a href = "#" data-toggle = "modal" data-target = "#myModal">新的任务</a>
            <div class = "modal fade" id = "#myModal" tabindex = "-1" role = "dialog"
            aria-labelledby = "myModalLabel" aria-hidden = "true">
            <div class = "modal-dialog ">
            <div class = "modal-content">
            <div class = "modal-header">
            <button type = "button" class = "close" data-dismiss = "modal"
            aria-hidden = "true">&times;
            </button>
            <h4 class = "modal-title" id = "myModalLabel">添加</h4>
            </div>
            <div class = "modal-body">
            <form role = "form" method = "post" action = "/add">
            <div class = "form-group">
            <textarea class= "form-control" id = "todo" name = "todo"
            required placeholder = "新的任务"></textarea>
            <div class = "modal-footer">
            <button type = "submit" class= "btn btn-primary">提交
            </button>
            </div>
            </div>
            </form>
            </div>
            </div>
            </div>
            </div>
            </li>
            </ul>

            <li class = "dropdown">
            <a href = "#" class = "dropdown-toggle" data-toggle = "dropdown"><img
            src ="/static/img/user.png" height = "22" width = "22"><span
            class = "caret"></span></a>
            <ul class = "dropdown-menu dropdown-menu-right" aria-labelledby = "dLabel">
            <li><a href = "/">我的任务</a></li>
            <li role = "separator" class = "divider"></li>
            <li><a href = "/setting">设置</a></li>
            <li><a href = "/logout">退出</a></li>
            </ul>
            </li>
            {{ else }}
            <li><a href = "/register">注册</a></li>
            <li><a href = "/login">登录</a></li>
            {{ end }}
            </ul>
            </div>
            </div>
            </nav>
            </header>