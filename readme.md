# TODO:
* Create "installer" script
  * A shell script that just moves the files and sets special permissions on the executable
    * setcap cap_net_raw=+ep /path/to/your/compiled/binary
    * chmod u+x /path/to/your/compiled/binary  