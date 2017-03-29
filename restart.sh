#!bin/bash

killall ps_web_0

if [ "$?" = "0" ]; then
	echo "Closing down server..."
else
	echo "Restarting server..."
fi


rm -f ps_web_0
echo Restarting server...
go install
go build
./ps_web_0
