import argparse
import sys
from typing import TextIO, Tuple
import re

class ApduCommand:
   def __init__(self) -> None:
      self._raw = bytearray()
      pass

   def append_byte(self, byte : int):
      self._raw.append(byte)


class ParserException(Exception):
   def __init__(self, *args: object, line : int, charpos : int) -> None:
      super().__init__(*args)
      self.line = line
      self.charpos = charpos

class ApduCommandParser:
   def __init__(self, input_stream:  TextIO) -> None:
      self._cur_line = 1
      self._stream = input_stream
      self._regx_whitespace = re.compile(r'[\r\t ]')
      self._regx_invalid_hexchars = re.compile(r'[^a-f0-9]', re.IGNORECASE)

   def get_next_apdu_command(self) -> ApduCommand:
      line = self._stream.readline(1)
      whitespace_removed = self._regx_whitespace.sub(line)
      if whitespace_removed == "":
         return None

      invalid_char_match = self._regx_invalid_hexchars.search(whitespace_removed)
      if invalid_char_match != None:
         # found invalid hex chars
         raise ParserException("Invalid hex char", line= self._cur_line, charpos=invalid_char_match.pos)

      cur_hex = ""
      apdu = ApduCommand()
      for cur_char in whitespace_removed:
         if len(cur_hex) == 2:
            apdu.append_byte(bytearray.fromhex(cur_hex)[0])
            cur_hex =""

         cur_hex = cur_hex + cur_char

      # append last byte
      if len(cur_hex) == 2:
         apdu.append_byte(bytearray.fromhex(cur_hex)[0])
      elif len(cur_hex) == 1:
         raise ParserException("Odd numbered hex chars", line=self._cur_line, charpos=len(line)-1)
      
      return apdu




def main():
   argparser = argparse.ArgumentParser("apdu-interpreter")
   argparser.add_argument("--input-file")
   argparser.add_argument("--output-file")
   args = argparser.parse_args()




if __name__ == "__main__":
   main()