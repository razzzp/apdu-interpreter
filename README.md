# APDU Interpreter

Reads APDU command logs and provides an interpretation of the commands. Can define schemas to define how the commands are to be interpreted.

## Schema Defintion

1. Define exact byte to match
2. Define a byte pattern in hexadecimal, can use 'X' as a wildcard
3. Define per bit, the meaning, handle 1 on or off
4. Define bit pattern matching


## TODO

[ ] Schema reader
[ ] Schema parser
[ ] Log reader
[ ] Log parser
[ ] Interpreters
[ ] Output formatter