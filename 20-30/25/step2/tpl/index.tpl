<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Meta.Title}}</title>
</head>
<body>
<h1>{{.Meta.WelcomeText}}</h1>
<ul>
    {{range $index, $element := .Data}}
    <li>{{$index}} - {{$element.UserName}} is in {{$element.Department}}</li>
    {{end}}
</ul>
</body>
</html>