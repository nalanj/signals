# signals

signals is a silly test program that outputs and responds to signals. It's for
testing other programs that need to send signals to a subprocess.

## Installing

```
go get -u sr.ht/~nalanj/signals
```

## Usage

Run `signals` to start the process. It will output any signals it receives and
exit on most signals, but on USR1 and USR2 it will output the signal but not
exit.

