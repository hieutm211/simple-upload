<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Upload File</title>
</head>
<body>
	<fieldset>
		<legend>Upload File</legend>
		<form enctype="multipart/form-data" action="/upload" method="post">
			<input type="file" name="userFile">
			<input type="submit" value="Upload">
		</form>
	</fieldset>
	<br/>
	<p>{{.}}<p>
</body>
</html>
