import subprocess
import sys
import os


def main():
    if len(sys.argv) <= 1:
        print("Please provide a valid file name from the current dir. ")
        return
    if not os.path.isfile(sys.argv[1]):
        print("No such a file exists: ", sys.argv[1])
        return
    i = open("./input.txt", "+r")
    o = open("./output.txt", "+w")
    subprocess.run(
        ["chmod", "-R", "777", "/Users/itgelt/Desktop/Leetcode/codeforce/py"])
    subprocess.Popen(sys.argv[1], stdin=i,
                     stdout=o, stderr=subprocess.STDOUT)
    i.close()
    o.close()


main()
