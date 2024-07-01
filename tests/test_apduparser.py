
import unittest
from apdu_interpreter.parser import ApduCommandParser, ParserException
from io import StringIO

class TestApduParser(unittest.TestCase):

   def test_get_next_apdu_command(self):
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


   def test_get_next_apdu_command_twice(self):
      input = "00 01  02 03 04 11 22 33 44\n 05 06070809 112233445566778899"
      input_stream = StringIO(input)
      parser  = ApduCommandParser(input_stream= input_stream)
      
      result = parser.get_next_apdu_command()
      self.assertEqual(result.get_cla(), 0)
      self.assertEqual(result.get_ins(), 1)
      self.assertEqual(result.get_p1(), 2)
      self.assertEqual(result.get_p2(), 3)
      self.assertEqual(result.get_p3(), 4)
      self.assertEqual(result.get_data(), bytearray.fromhex("11223344"))

      result = parser.get_next_apdu_command()
      self.assertEqual(result.get_cla(), 5)
      self.assertEqual(result.get_ins(), 6)
      self.assertEqual(result.get_p1(), 7)
      self.assertEqual(result.get_p2(), 8)
      self.assertEqual(result.get_p3(), 9)
      self.assertEqual(result.get_data(), bytearray.fromhex("112233445566778899"))


   def test_get_next_apdu_command_eof(self):
      input = ""
      input_stream = StringIO(input)
      parser  = ApduCommandParser(input_stream= input_stream)
      
      result = parser.get_next_apdu_command()
      self.assertIsNone(result)


   def test_get_next_apdu_command_invalid_char(self):
      input = "00 01  0X 03 04 11 22 33 44"
      input_stream = StringIO(input)
      parser  = ApduCommandParser(input_stream= input_stream)
      
      with self.assertRaises(ParserException):
         result = parser.get_next_apdu_command()


   def test_get_next_apdu_command_odd_number_of_chars(self):
      input = "00 01  02 03 04 11 22 33 4"
      input_stream = StringIO(input)
      parser  = ApduCommandParser(input_stream= input_stream)
      
      with self.assertRaises(ParserException):
         result = parser.get_next_apdu_command()
