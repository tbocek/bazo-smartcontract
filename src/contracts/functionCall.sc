# This is a simple program which calls a function
PUSH 55780
PUSH 5
CALL addNums 2
HALT

addNums:
LOAD 0
LOAD 1
ADD
RET
