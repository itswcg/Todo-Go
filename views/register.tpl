{{template "base.tpl" .}}
<main>
    <div class = "sidebar">
    <center><h3>Todo注册</h3></center>
    </div>
    <div class= "container">
    <form class = "form-signin" action ="/register" method = "post" role = "form">
    <div class = "form-group">
    <label for = "username">用户名</label>
    <input type = "text" class = "form-control" name = "username">
    </div>
    <div class = "form-group">
    <label for= "password">密码</label>
    <input type = "password" class = "form-control" name = "password">
    </div>
    <div class = "form-group">
    <label for = "confirm_password">确认密码</label>
    <input type = "confirm_password" class = "form-control" name= "confirm_password">
    </div>
    <button type = "submit" class = "btn btn-lg btn-primary btn-block">注册</button>
    </form>
    </div>
</main>
</body>
</html>