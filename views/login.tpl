{{template "base.tpl" .}}
<main>
    <div class = "sidebar">
    <center><h3>Todo登录</h3></center>
    </div>
    <div class= "container">
    <form class = "form-signin" action ="/login" method = "post" role = "form">
    <div class = "form-group">
    <label for = "username">用户名</label>
    <input type = "text" class = "form-control" name = "username">
    </div>
    <div class = "form-group">
    <label for= "password">密码</label>
    <input type = "password" class = "form-control" name = "password">
    </div>
    <button type = "submit" class = "btn btn-lg btn-primary btn-block">登录</button>
    </form>
    </div>
</main>
</body>
</html>