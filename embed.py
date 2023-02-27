from ctypes import *
vm=CDLL('./libgo.vm.so')

GoInt = c_longlong

class GoSlice(Structure):
    _fields_ = [('data',c_char_p),
                ('len',GoInt),
                ('cap',GoInt)]


def run_code(code):
    #code = b'0\x00\r\x00Hello, World!1\x00'
    size = len(code)
    prog = c_char_p(code)
    program = GoSlice(prog,size,size)
    vm.RunCode(program)

