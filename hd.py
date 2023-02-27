import codon
import timeit

with open('j.raw','rb') as f:
    r = f.read()

def hexdumpp(r: bytes):
    return list('0'*(2-len(hex(x)[2:]))+hex(x)[2:] for x in list(r))
@codon.jit
def hexdumpc(r: str):
    return list('0'*(2-len(hex(x)[2:]))+hex(x)[2:] for x in list(r))

print("without:",timeit.timeit((lambda: hexdumpp(r)),number=5))
print("with:",timeit.timeit((lambda: hexdumpc(r)),number=5))
