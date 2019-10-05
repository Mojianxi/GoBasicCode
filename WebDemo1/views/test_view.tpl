<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
</head>
<body>
	{{if .IsDisplay}}
		<em>{{.Content}}</em>
	{{else}}
		<em>{{.Content2}}</em>
	{{end}}

	{{range .users}}
		{{.TName}}
	{{end}}
</body>
</html>