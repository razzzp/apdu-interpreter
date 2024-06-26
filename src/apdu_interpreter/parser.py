
from typing import TextIO, Tuple
import re

class ApduCommand:
   def __init__(self, raw : bytearray = bytearray()) -> None:
      self._raw = raw
      pass

   def append_byte(self, byte : int):
      self._raw.append(byte)

   def is_valid_index(self, index : int) -> bool:
      if(index < 0): return False
      return index < len(self._raw) -1

   def get_cla(self) -> int | None:
      if not self.is_valid_index(0): return None
      return self._raw[0]

   def get_ins(self) -> int | None:
      if not self.is_valid_index(1): return None
      return self._raw[1]

   def get_p1(self) -> int | None:
      if not self.is_valid_index(2): return None
      return self._raw[2]

   def get_p2(self) -> int | None:
      if not self.is_valid_index(3): return None
      return self._raw[3]

   def get_p3(self) -> int | None:
      if not self.is_valid_index(4): return None
      return self._raw[4]
   
   def get_data(self) -> bytearray | None:
      if not self.is_valid_index(5): return None
      return self._raw[5:]



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
      line = self._stream.readline()
      # remove whitespace
      whitespace_removed = self._regx_whitespace.sub("",line)
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

