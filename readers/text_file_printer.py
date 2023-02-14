import sys
import os

"""
Command Line script that prints or opens files of various types based on their extension.  
Ex: python3 file_opener.py <file_path> 
"""


def open_file(file_path):
    file_extension = os.path.splitext(file_path)[1]
    if file_extension == '.txt':
        with open(file_path, 'r') as file:
            print(file.read())
    elif file_extension == '.pdf':
        os.system(f'open "{file_path}"')
    elif file_extension == '.csv':
        os.system(f'pandas "{file_path}"')
    else:
        print(f"Error: The file extension '{file_extension}' is not supported.")
        sys.exit(1)


if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Error: Incorrect number of arguments.")
        print("Usage: python file_opener.py <file_path>")
        sys.exit(1)
    open_file(sys.argv[1])
