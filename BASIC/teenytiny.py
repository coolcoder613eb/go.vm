from lex import *
from emit import *
from parse import *
import sys



def main():
    print("Teeny Tiny Compiler")

    if len(sys.argv) != 2:
        sys.exit("Error: Compiler needs source file as argument.")
    with open(sys.argv[1], 'r') as inputFile:
        input = inputFile.read()
    out = '.'.join(sys.argv[1].split('.')[0:-1])+'.in'


    # Initialize the lexer, emitter, and parser.
    lexer = Lexer(input)
    emitter = Emitter(out)
    parser = Parser(lexer, emitter)

    parser.program() # Start the parser.
    emitter.writeFile() # Write the output to file.
    print("Compiling completed.")

main()
