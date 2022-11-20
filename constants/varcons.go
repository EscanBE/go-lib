package constants

import (
	"os"
	"syscall"
)

var TrapExitSignal = []os.Signal{
	syscall.SIGTERM, // Terminate by another process
	syscall.SIGHUP,  // Terminal goes away
	syscall.SIGPIPE, // wrote to broken pipe/socket
	syscall.SIGKILL, // system force kill at kernel level
	syscall.SIGABRT, // normally, library calls abort when detect internal error
	syscall.SIGINT,  // Ctrl + C
	syscall.SIGQUIT, // Ctrl + \
	syscall.SIGTSTP, // Ctrl + Z
}
