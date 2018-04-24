# This is a simple program with very basic opcodes with ABI

CALLDATA

# ABI
DUP
PUSH 1          # should be replaced with function signature hash
EQ
JMPIF addNums

DUP
PUSH 2          # should be replaced with function signature hash
EQ
JMPIF subNums

HALT

# Functions
addNums:        # function 1
POP
ADD
HALT

subNums:        # function 2
POP
SUB
HALT
