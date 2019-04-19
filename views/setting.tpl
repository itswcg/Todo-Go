{{template "base.tpl" .}}
<main>
    <div class = "sidebar">
    <center><h3>修改密码</h3></center>
    </div>
    <div class= "container">
        <form class = "form-signin" action ="/setting" method = "post" role = "form">
        <div class = "form-group">
        <label for = "old_password">旧密码</label>
        <input type = "password" class = "form-control" name = "old_password"></div>
        <div class = "form-group">
        <label for= "password">新密码</label>
        <input type = "password" class = "form-control" name = "new_password">
        </div>
        <div class = "form-group">
        <label for = "confirm_password">确认新密码</label>
        <input type = "confirm_password" class = "form-control" name= "confirm_password">
        </div>
        <button type = "submit" class = "btn btn-lg btn-primary btn-block">修改</button>
    </form>
    </div>
</main>
</body>
</html>