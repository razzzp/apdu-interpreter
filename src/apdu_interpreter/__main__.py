import argparse

def main():
   argparser = argparse.ArgumentParser("apdu-interpreter")
   argparser.add_argument("--input-file")
   argparser.add_argument("--output-file")
   args = argparser.parse_args()


if __name__ == "__main__":
   main()