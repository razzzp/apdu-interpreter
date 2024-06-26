
import unittest
from apdu_interpreter.parser import ApduCommandParser
from io import StringIO

class TestApduParser(unittest.TestCase):

   def test_get_apdu(self):
      input = "00 01  02 03 04 11 22 33 44"
      input_stream = StringIO(input)
      parser  = ApduCommandParser(input_stream= input_stream)
      
      result = parser.get_next_apdu_command()
      self.assertEqual(result.get_cla(), 0)
      self.assertEqual(result.get_ins(), 1)
      self.assertEqual(result.get_p1(), 2)
      self.assertEqual(result.get_p2(), 3)
      self.assertEqual(result.get_p3(), 4)
      self.assertEqual(result.get_data(), bytearray.fromhex("11223344"))