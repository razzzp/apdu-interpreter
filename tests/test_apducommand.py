
import unittest
from apdu_interpreter.parser import ApduCommand
from io import StringIO

class TestApduParser(unittest.TestCase):

   def test_empty_apdu(self):
      apdu  = ApduCommand()
      
      self.assertEqual(apdu.get_cla(), None)
      self.assertEqual(apdu.get_ins(), None)
      self.assertEqual(apdu.get_p1(), None)
      self.assertEqual(apdu.get_p2(), None)
      self.assertEqual(apdu.get_p3(), None)
      self.assertEqual(apdu.get_data(), None)