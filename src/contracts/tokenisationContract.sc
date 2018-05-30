# How to define a contract variable?
# How can we do something like a constructor
NEWMAP
# Minter



mint:
SSLOAD 1
EQ
JUMPIF isMinter
RET

isMinter:
#Get receiver to tos
MAPGETVAL
ADD
AMOUNT
MAPSETVAL
RET


send:
CALLER
MAPGETVAL
#Get the amount to tos
LTE
JMPIF

CALLER
MAPGETVAL
# amount
SUB
MAPSETVAL

# receiver
MAPGETVAL
# amount
ADD
MAPSETVAL
RET






PUSH 8
PUSH 5
ADD
HALT
