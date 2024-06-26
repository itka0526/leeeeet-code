import string
import subprocess
import sys
import os


class bcolors:
    HEADER = "\033[95m"
    OKBLUE = "\033[94m"
    OKCYAN = "\033[96m"
    OKGREEN = "\033[92m"
    WARNING = "\033[93m"
    FAIL = "\033[91m"
    ENDC = "\033[0m"
    BOLD = "\033[1m"
    UNDERLINE = "\033[4m"


signals = {
    1: "SIGHUP    terminal line hangup",
    2: "SIGINT    interrupt program",
    3: "SIGQUIT    quit program",
    4: "SIGILL    illegal instruction",
    5: "SIGTRAP    trace trap",
    6: "SIGABRT    abort program (formerly SIGIOT)",
    7: "SIGEMT    emulate instruction executed",
    8: "SIGFPE    floating-point exception",
    9: "SIGKILL    kill program",
    10: "SIGBUS    bus error",
    11: "SIGSEGV    segmentation violation",
    12: "SIGSYS    non-existent system call invoked",
    13: "SIGPIPE    write on a pipe with no reader",
    14: "SIGALRM    real-time timer expired",
    15: "SIGTERM    software termination signal",
    16: "SIGURG    urgent condition present on socket",
    17: "SIGSTOP    stop (cannot be caught or ignored)",
    18: "SIGTSTP    stop signal generated from keyboard",
    19: "SIGCONT    continue after stop",
    20: "SIGCHLD    child status has changed",
    21: "SIGTTIN    background read attempted from control terminal",
    22: "SIGTTOU    background write attempted to control terminal",
    23: "SIGIO     I/O is possible on a descriptor (see fcntl(2))",
    24: "SIGXCPU    cpu time limit exceeded (see setrlimit(2))",
    25: "SIGXFSZ    file size limit exceeded (see setrlimit(2))",
    26: "SIGVTALRM    virtual time alarm (see setitimer(2))",
    27: "SIGPROF    profiling timer alarm (see setitimer(2))",
    28: "SIGWINCH    Window size change",
    29: "SIGINFO    status request from keyboard",
    30: "SIGUSR1    User defined signal 1",
    31: "SIGUSR2    User defined signal 2",
}


def main():
    if len(sys.argv) <= 1:
        print(bcolors.FAIL + "Please provide a valid file name from the current dir. ")
    basename = sys.argv[1]
    basename_no_ext = sys.argv[1].removesuffix(".cpp")
    print(basename, basename_no_ext)
    if not os.path.isfile(basename):
        print(bcolors.FAIL + "No such a file exists: ", bcolors.UNDERLINE + basename)
        return
    i = open("./input.txt", "+r")
    o = open("./output.txt", "+w")
    o.write("")
    if os.system(f"g++-13 -ld_classic -g {sys.argv[1]} -o {basename_no_ext}") == 0:
        print(bcolors.OKGREEN + "Successfully compiled. ")
    else:
        print(bcolors.FAIL + "Failed to compile. ")
        return
    proc = subprocess.Popen("./" + basename_no_ext, stdin=i, stdout=o, stderr=subprocess.STDOUT, shell=True)
    proc.communicate()
    if proc.returncode != 0:
        err = signals[abs(proc.returncode)].split("    ")
        o.write(f"'{basename_no_ext}'" + " got: " + err[0] + f" ({err[1]})")
    i.close()
    o.close()


main()
