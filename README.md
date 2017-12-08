## chirper

Chirper is a simple application that lets you set alarm and timer in your desktop. You can set multiple alarms, snooze your alarms. You can also set timers. So, you can get rid off your dependancies if setting your alarm or reminder on phone and instead let `chirper` chirp for you at your set time. ;)

### Run

```
$ go get github.com/shreyaganguly/chirper
$ chirper -b <ip-address-of-chirper-server>(default localhost) -p <port-number-of-chirper-server>(default 8080) -s <absolute-path-of-the-sound-file-to-play-for-alarms-and-timers> -snooze-interval <snooze-time-interval-in-minutes>(default 5)
```

Now hit `ip-address:port` in your browser and just start setting your tiem and let it chirp!

> P.S : Go Version should be from 1.8 onwards
