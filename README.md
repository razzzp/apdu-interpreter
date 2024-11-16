# APDU Interpreter

Reads APDU command logs and provides an interpretation of the commands. Can define schemas to define how the commands are to be interpreted.

## Schema Defintion

1. Define exact byte to match
2. Define a byte pattern in hexadecimal, can use 'X' as a wildcard
3. Define per bit, the meaning, handle 1 on or off
4. Define bit pattern matching

## Output


Each line of command response pair will have the following output

A0 A4 00 0C 02 3F 00
90 00
```json
{
   "hex": "A0A4000C023F00",
   "command_name": "SELECT",
   "descriptions": "...",
   "schema": {
      "name": "TS 102 221",
      "group": "ETSI",
      "version": "v18.2.0"
   },
   "cla": {
      "hex": "A0",
      "intp": []
   },
   "ins": {
      "hex": "A4",
      "intp": []
   },
   "p1": {
      "hex": "00",
      "intp": [
         "Select DF, EF or MF by file ID"
      ]
   },
   "p2": {
      "hex": "0C",
      "intp": [
         "No data returned"
         "First or only occurence"
      ]
   },
   "p3": {
      "hex": "02",
      "intp": [
         "Lc = 2 bytes"
      ]
   },
   "data": {
      "hex": "3F00",
      "intp": []
   }
}
```

## TODO

[ ] Schema reader
[ ] Schema parser
[ ] Log reader
[ ] Log parser
[ ] Interpreters
[ ] Output formatter