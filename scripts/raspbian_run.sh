# if you do not have screen installed
# $ sudo apt install screen -y
# this script will run the go binary as a daemon
# to kill the daemon
# $ ps aux | grep rpi_main
# find the PID for the non screen process
# should be the second one in the list!

screen -d -m ./builds/rpi_main
