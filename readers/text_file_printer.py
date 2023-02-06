import sys

"""
Simple Command Line script that prints the content of specified textfile to terminal.  
Ex: python3 text_file_printer.py <path_to_.txt_file> 
"""


def print_file_contents(file_path):
    try:
        with open(file_path, 'r') as file:
            print(file.read())
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' does not exist.")
        sys.exit(1)


if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Error: Incorrect number of arguments.")
        print("Usage: python print_file_contents.py <file_path>")
        sys.exit(1)
    print_file_contents(sys.argv[1])
