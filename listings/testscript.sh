# build swm, swmctl and test app
go build github.com/janbina/swm/cmd/swm
go build github.com/janbina/swm/cmd/swmctl
go build github.com/janbina/swm/test

# start xvfb
Xvfb :111 -screen 0 1024x768x16 &
XVFB_PID=$!

# start swm, discard output logs
./swm > /dev/null 2>&1 &

# run test and save the exit code
./test; EXIT_CODE=$?

# shut down nested X server and cleanup
kill -15 $XVFB_PID
rm swm swmctl test

exit $EXIT_CODE