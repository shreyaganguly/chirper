package main

const clock = `<html>
    <head>
    <title></title>
    <script src="http://localhost:35729/livereload.js"></script>
    </head>
    <body>
        <form action="/set" method="post">
            Date:<input type="text" name="date">
            Time:<input type="text" name="time">
            <input type="submit" value="Set Alarm">
        </form>
        <script>
          var a = new Audio({{ .SoundFile }})
          {{ range .Alarms }}
            {{ if .Playing }}
              a.play()
            {{ end }}
          {{ end }}

        </script>
    </body>
</html>`
