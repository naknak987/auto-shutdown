# Installation:
1. Move the executable to `/usr/local/bin`
1. Make sure the user has executable permissions on the executable using `chmod u+x /usr/local/bin/auto-shutdown`
1. Customize the service file with the target IP address and any options you want
1. Move the service file to `/etc/systemd/system/`
1. Reload the systemd daemon with `systemctl daemon-reload`
1. Enable the service with `systemctl enable auto-shutdown`
1. Start the service with `systemctl start auto-shutdown`

# TODO:
* Create "installer" script
  * A shell script that just moves the files and sets special permissions on the executable
    * setcap cap_net_raw=+ep /path/to/your/compiled/binary
      * May not need to do this. Seemed to work without it
    * chmod u+x /path/to/your/compiled/binary