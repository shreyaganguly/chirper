package main

const clock = `
<html>

<head>
  <style>
    .backgroundOrange {
      background-color: orange !important;
    }
  </style>
  <title></title>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="/assets/css/flipclock.css">
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
  <script type="text/javascript" src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  <script src="http://1000hz.github.io/bootstrap-validator/dist/validator.min.js"></script>
  <script src="/assets/js/flipclock.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.9.0/moment-with-locales.js"></script>
  <script src="https://cdn.rawgit.com/Eonasdan/bootstrap-datetimepicker/e8bddc60e73c1ec2475f827be36e1957af72e2ea/src/js/bootstrap-datetimepicker.js"></script>
</head>

<body>
  <div class="container">
  </div>
  <div class="container">
    <h3>Tabs</h3>
    <ul class="nav nav-tabs nav-justified">
      <li class="active"><a data-toggle="tab" href="#alarmsection">Alarm</a></li>
      <li><a data-toggle="tab" href="#timersection">Timer</a></li>
    </ul>
  </div>
  <div class="tab-content">
    <div id="alarmsection" class="tab-pane fade in active">
      <div class="container">
        <div class="row well" style="margin:0px;">
          <div class='col-sm-offset-3 col-sm-6'>
            <h3>Set Alarm</h3>
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
                {{ range .Alarms}} {{ if eq $alarmedTimeStamp .TimeStamp }}
                <li class="list-group-item" id="alarmed">{{ .DateTime }}
                  <span class="pull-right" value={{ .TimeStamp }} onClick="deleteAlarm(this)"><i class="fa fa-times" id="clock-delete" aria-hidden="true"></i></span>
                  <span class="pull-right" value={{ .TimeStamp }} time={{ .DateTime }} onClick="snoozeAlarm(this)" style="margin-right:10px;"><i class="fa fa-clock-o" id="clock-snooze" aria-hidden="true"></i></span>
                </li>
                {{ else }}
                <li class="list-group-item">{{ .DateTime }}
                  <span class="pull-right" value={{ .TimeStamp }} onClick="deleteAlarm(this)"><i class="fa fa-times" id="clock-delete" aria-hidden="true"></i></span>
                  <span class="pull-right hidden" value={{ .TimeStamp }} onClick="snoozeAlarm(this)" style="margin-right:10px;"><i class="fa fa-clock-o" id="clock-snooze" aria-hidden="true"></i></span>
                </li>
                {{ end }} {{ end }}
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div id="timersection" class="tab-pane fade">
      <div class="container">
        <div class="row well" style="margin:0px;">
          <div class='col-sm-offset-3 col-sm-6'>
            <h3>Start Timer</h3>
              <form data-toggle="validator" role="form" novalidate="true">
                <div class="row">
                  <div class="form-group col-xs-3">
                    <label for="inputHours" class="control-label">Hours</label>
                    <input type="text" pattern="^[0-5]?[0-9]$" maxlength="2" class="form-control" id="hours" placeholder="0-59" data-error="Incorrect!">
                    <div class="help-block with-errors"></div>
                  </div>
                  <div class="form-group col-xs-3">
                    <label for="inputMinutes" class="control-label">Minutes</label>
                    <input type="text" pattern="^[0-5]?[0-9]$" maxlength="2" class="form-control" id="minutes" placeholder="0-59" data-error="Incorrect!">
                    <div class="help-block with-errors"></div>
                  </div>
                  <div class="form-group col-xs-3">
                    <label for="inputSeconds" class="control-label">Seconds</label>
                    <input type="text" pattern="^[0-5]?[0-9]$" maxlength="2" class="form-control" id="seconds" placeholder="0-59" data-error="Incorrect!">
                    <div class="help-block with-errors"></div>
                  </div>
                </div>
                <div class="form-group">
                  <button type="submit" id="submitbtntimer" class="btn btn-primary" style="margin:10;">Submit</button>
                </div>
              </form>
          </div>
        </div>
            <div class="row timer hidden well" style="margin-top:5px;margin-left:0px;margin-right:0px;margin-bottom:0px;">
              <div class="col-sm-offset-3 col-sm-7">
                <div class="clock"></div>
              </div>
              <div class="col-sm-2">
                <button type="button" class="btn btn-warning" style="margin-top:40px;" id="btnclose">Close</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
</body>
<script>
  var audio = new Audio({{ .SoundFile }});
  var clock;
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
      $('#datetimepicker').data('DateTimePicker').date(new Date())
      return false;

  });
  $("#submitbtntimer").on("click", function() {
    audio.pause();
    var hour = isNaN(parseInt($("#hours").val())) ? 0 : parseInt($("#hours").val()) * 60 * 60
    var minute = isNaN(parseInt($("#minutes").val())) ? 0 : parseInt($("#minutes").val()) * 60
    var second = isNaN(parseInt($("#seconds").val())) ? 0 : parseInt($("#seconds").val())
    $("#hours").val("")
    $("#minutes").val("")
    $("#seconds").val("")
    $(".timer").removeClass("hidden");
    clock = $(".clock").FlipClock({
      clockFace: "HourlyCounter",
      autoStart: false,
      callbacks: {
        stop: function() {
          if ($(".timer").is(":visible") === true) {
            audio.play();
          }
        },
        reset: function() {
          $(".timer").addClass("hidden");
        }
      }
    });
    clock.setTime(hour + minute + second);
    clock.setCountdown(true);
    clock.start();
    return false;
  });
  $("#btnclose").on("click", function() {
    audio.pause();
    clock.reset();
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
