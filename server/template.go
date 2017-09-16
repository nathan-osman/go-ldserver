package server

import (
	"bytes"
	"html/template"
)

const pageHTML = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset='utf-8'>
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <title>ldserver</title>

        <style>
            body {
                background-color: #eee;
                font-family: sans-serif;
                font-size: 10pt;
                margin: 20px;
            }
            button {
                background-color: #ddd;
                border: 1px solid #777;
                padding: 4px 8px;
            }
            button:hover {
                background-color: #fff;
                cursor: pointer;
            }
            label {
                display: block;
                font-weight: bold;
            }
            p {
                color: #555;
            }
            .error {
                background-color: #fcc;
                border: 1px solid #f00;
                color: #f00;
                margin: 8px 0;
                padding: 4px;
            }
        </style>
    </head>
    <body>
        <h2>ldserver</h2>
        <p>Use the tools below to interact with the server.</p>
        <p>
            Status:
            <strong>
                {{if .Loaded}}
                    performance loaded
                {{else}}
                    awaiting upload
                {{end}}
            </strong>
        </p>
        <br>
        <form method="post" enctype="multipart/form-data">
            {{if .Error}}
                <div class="error">Error: {{.Error}}.</div>
            {{end}}
            <input type="file" name="file">
            <hr>
            <button type="submit" name="action" value="upload">Upload</button>
            {{if .Loaded}}
                <button type="submit" name="action" value="start">Start</button>
                <button type="submit" name="action" value="stop">Stop</button>
            {{end}}
        </form>
    </body>
</html>
`

var pageTemplate = template.Must(template.New("home").Parse(pageHTML))

func renderTemplate(data map[string]interface{}) []byte {
	b := bytes.NewBuffer(nil)
	pageTemplate.Execute(b, data)
	return b.Bytes()
}
