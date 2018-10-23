# LogBucket 

a simple web service to keep important logs in mongodb 


For recent versions of Ubuntu create `/etc/systemd/system/myservice.service` with:
```
    [Unit]
    Description=my amazing service
    
    [Service]
    Restart=always
    RestartSec=3
    ExecStart=/usr/bin/path/to/my/service some args
    
    [Install]
    WantedBy=multi-user.target
```
Run this one time:

   ``` systemctl enable myservice```

You can then start and stop your service with:
```
    sudo systemctl start myservice
    sudo systemctl stop myservice
```
It will also start automatically when the server boots, and restart if it crashes.

More here:

https://wiki.ubuntu.com/SystemdForUpstartUsers#Example_Systemd_service

There are many other service management solutions, but systemd comes with Ubuntu 16 so it's probably the easiest to use.