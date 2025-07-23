package web

const IndexHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>Weather Ball</title>
    <link rel="stylesheet" href = "style.css">
</head>

<body>
    <h1>Weather Ball</h1>
    <input id="location" placeholder="Enter location (e.g. San Diego, California)" size="40" />
    <button onclick="fetchCoolTimes()">Search</button>

    <h2>Results:</h2>
    <pre id="results">No results yet.</pre>

    <script src="script.js"></script>
</body>


</html>
`

// func renderIndex(w http.ResponseWriter) error {
// 	t, err := template.New("index").Parse(IndexHTML)
// 	if err != nil {
// 		return err
// 	}
// type TemplateData struct {
// 	Id         string
// 	UploadTime string
// 	EscapedId  string
// }

// // Convert videos to template data
// templateData := make([]TemplateData, len(videos))
// for i, video := range videos {
// 	templateData[i] = TemplateData{
// 		Id:         video.Id,
// 		UploadTime: video.UploadedAt.Format("2006-01-02 15:04:05"),
// 		EscapedId:  url.PathEscape(video.Id),
// 	}
// }
// 	return t.Execute(w, nil)
// }

// const videoHTML = `
// <!DOCTYPE html>
// <html>
//   <head>
//     <meta charset="UTF-8" />
//     <title>{{.Id}} - TritonTube</title>
//     <script src="https://cdn.dashjs.org/latest/dash.all.min.js"></script>
//   </head>
//   <body>
//     <h1>{{.Id}}</h1>
// 	  <p>Uploaded at: {{.UploadedAt}}</p>

//     <video id="dashPlayer" controls style="width: 640px; height: 360px"></video>
//     <script>
//       var url = "/content/{{.Id}}/manifest.mpd";
//       var player = dashjs.MediaPlayer().create();
//       player.initialize(document.querySelector("#dashPlayer"), url, false);
//     </script>

//     <p><a href="/">Back to Home</a></p>
//   </body>
// </html>
// `
