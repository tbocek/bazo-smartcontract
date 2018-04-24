# This is a program to benchmark the execution of a smart contract
# PUSH 4     # Base
# PUSH 13    # Exp
# PUSH 497   # Modulus

# For loop
PUSH 13
PUSH 0

loop:
PUSH 1
ADD
ROLL 0
LT
JMPIF loop

HALT