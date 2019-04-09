<!DOCTYPE html>

<html>
<head></head>
<body>
    <p>{{.content}}</p>
    <ul>
     {{range .todos}}
     <li>
     {{.id}}{{.content}}{{.is_do}}{{.authod_id}}{{.create_date}}
     </li>
     {{end}}
    </ul>
</body>
