# This is an example how a simple smartcontract with a state variable looks like

# TransactionData
PUSH 5
PUSH incCounterBy       # Function hash

# ABI:
DUP
PUSH incCounterBy
EQ
JMPIF incCounterBy

DUP
PUSH dupFunc
EQ
JMPIF dupFunc

# Store counter with value 0 on account storage, initialisation
PUSH 0
SSTORE counter 1

incCounterBy:           # function
SLOAD counter
ADD
SSTORE counter
RET

dupFunc:
DUP
RET
