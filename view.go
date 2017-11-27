package main

const clock = `
<html>

<head>
<style>
.backgroundOrange{
        background-color: orange !important;
    }
</style>
  <title></title>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" type="text/css" media="screen" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css" />
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css">
  <link href="https://eonasdan.github.io/bootstrap-datetimepicker/css/prettify-1.0.css" rel="stylesheet">
  <link href="https://eonasdan.github.io/bootstrap-datetimepicker/css/base.css" rel="stylesheet">
  <link href="https://cdn.rawgit.com/Eonasdan/bootstrap-datetimepicker/e8bddc60e73c1ec2475f827be36e1957af72e2ea/build/css/bootstrap-datetimepicker.css" rel="stylesheet">

  <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->
  <!--[if lt IE 9]>
            <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
            <script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
        <![endif]-->
  <script src="http://localhost:35729/livereload.js"></script>
  <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
  <script type="text/javascript" src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.9.0/moment-with-locales.js"></script>
  <script src="https://cdn.rawgit.com/Eonasdan/bootstrap-datetimepicker/e8bddc60e73c1ec2475f827be36e1957af72e2ea/src/js/bootstrap-datetimepicker.js"></script>
</head>

<body>
  <div class="container">
    <div class="row">
      <div class='col-sm-6'>
        <div class="form-group">
          <form id="datetimepickerform" method="post">

            <label for="dateandtime">Set Date And Time For Alarm</label>
            <div class='input-group date' id='datetimepicker'>
              <input type='text' class="form-control" name="datetime" value="" />
              <span class="input-group-addon">
            <span class="glyphicon glyphicon-calendar"></span>
              </span>
            </div>
        </div>
        <button type="submit" id="submitbtn" class="btn btn-primary" style="margin:10;">Submit</button>
        </form>
        {{ $alarmedTimeStamp := .TimeStamp }}
        <div id="alarmview">
        <ul class="list-group">
          {{ range .Alarms}}
            {{ if eq $alarmedTimeStamp .TimeStamp }}
              <li class="list-group-item" id="alarmed">{{ .DateTime }}
                <span class="pull-right" value={{ .TimeStamp }} onClick="deleteAlarm(this)"><i class="fa fa-times" id="clock-delete" aria-hidden="true"></i></span>
                <span class="pull-right" value={{ .TimeStamp }} time={{ .DateTime }} onClick="snoozeAlarm(this)" style="margin-right:10px;"><i class="fa fa-clock-o" id="clock-snooze" aria-hidden="true"></i></span>
              </li>
            {{ else }}
              <li class="list-group-item">{{ .DateTime }}
                <span class="pull-right" value={{ .TimeStamp }} onClick="deleteAlarm(this)"><i class="fa fa-times" id="clock-delete" aria-hidden="true"></i></span>
                <span class="pull-right hidden" value={{ .TimeStamp }} onClick="snoozeAlarm(this)" style="margin-right:10px;"><i class="fa fa-clock-o" id="clock-snooze" aria-hidden="true"></i></span>
              </li>
            {{ end }}
          {{ end }}
        </ul>
        </div>
      </div>
    </div>
  </div>
</body>
<script>
  var audio = new Audio({{ .SoundFile }});
  {{ if .Playing }}
    audio = new Audio({{ .SoundFile }});
    audio.play();
  {{ end }}
  var currentDate = new Date();
  $(function() {
    $('#datetimepicker').datetimepicker({
      minDate: currentDate,
      defaultDate: currentDate,
      format: "DD/MM/YYYY hh:mm A"
    });
  });
    $("#submitbtn").on("click", function(){
      $.ajax({
              type: "post",
              url: "/set",
              dataType: 'html',
              data: $("#datetimepickerform").serialize(),
              success: function(result){
                $("#alarmview").html(result);
              },
      });
  });
    function deleteAlarm(e){
      $.ajax({
              type: "post",
              url: "/delete",
              dataType: 'html',
              data: {timestamp: e.getAttribute('value')},
              success: function(result){
                $("#alarmview").html(result);
                if (e.getAttribute('value') == {{ .TimeStamp }} && {{ .Playing }} === true) {
                  audio.pause();
                }
              },
      });
    };
    function snoozeAlarm(e){
      $.ajax({
              type: "post",
              url: "/snooze",
              dataType: 'html',
              data: {timestamp: e.getAttribute('value'), time: moment.utc(e.getAttribute('time'),'DD/MM/YYYY hh:mm A').add(5,'minutes').format('DD/MM/YYYY hh:mm A')},
              success: function(result){
                $("#alarmview").html(result);
                if (e.getAttribute('value') == {{ .TimeStamp }} && {{ .Playing }} === true) {
                  audio.pause();
                }
              },
      });
    };
    {{ if .Playing }}
      setInterval(function(){
        $("#alarmed").toggleClass("backgroundOrange");
        },1000)
    {{ end }}
</script>
</body>

</html>`
