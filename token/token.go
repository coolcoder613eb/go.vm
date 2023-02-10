// Package token contains the list of token-types we accept/recognize.
package token

// Type is a string
type Type string

// Token struct represent the lexer token
type Token struct {
	Type    Type
	Literal string
}

// pre-defined Type
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	LABEL   = "LABEL"
	INT     = "INT"
	STRING  = "STRING"
	COMMA   = "COMMA"

	// math
	ADD = "ADD"
	AND = "AND"
	DEC = "DEC"
	DIV = "DIV"
	INC = "INC"
	MUL = "MUL"
	OR  = "OR"
	SUB = "SUB"
	XOR = "XOR"

	// control-flow
	CALL  = "CALL"
	JMP   = "JMP"
	JMPNZ = "JMPNZ"
	JMPZ  = "JMPZ"
	RET   = "RET"

	// stack
	PUSH = "PUSH"
	POP  = "POP"

	// types
	IS_STRING  = "IS_STRING"
	IS_INTEGER = "IS_INTEGER"
	STRING2INT = "STRING2INT"
	INT2STRING = "INT2STRING"

	// compare
	CMP         = "CMP"
	IS_MORE_REG = "IS_MORE"

	// store
	STORE = "STORE"

	// print
	PRINT_INT = "PRINT_INT"
	PRINT_STR = "PRINT_STR"

	// memory
	PEEK = "PEEK"
	POKE = "POKE"

	// Misc
	CONCAT = "CONCAT"
	DATA   = "DATA"
	DB     = "DB"
	EXIT   = "EXIT"
	MEMCPY = "MEMCPY"
	NOP    = "NOP"
	RANDOM = "RANDOM"
	SYSTEM = "SYSTEM"
	TRAP   = "TRAP"
)

// reversed keywords
var keywords = map[string]Type{

	// compare
	"cmp":     CMP,
	"is_more": IS_MORE_REG,

	// types
	"is_integer": IS_INTEGER,
	"is_string":  IS_STRING,
	"int2string": INT2STRING,
	"string2int": STRING2INT,

	// store
	"store": STORE,

	// print
	"print_int": PRINT_INT,
	"print_str": PRINT_STR,

	// math
	"add": ADD,
	"and": AND,
	"dec": DEC,
	"div": DIV,
	"inc": INC,
	"mul": MUL,
	"or":  OR,
	"sub": SUB,
	"xor": XOR,

	// control-flow
	"call":  CALL,
	"jmp":   JMP,
	"jmpnz": JMPNZ,
	"jmpz":  JMPZ,
	"ret":   RET,

	// stack
	"push": PUSH,
	"pop":  POP,

	// memory
	"peek": PEEK,
	"poke": POKE,

	// misc
	"exit":   EXIT,
	"concat": CONCAT,
	"DATA":   DATA,
	"DB":     DB,
	"int":    TRAP,
	"memcpy": MEMCPY,
	"nop":    NOP,
	"random": RANDOM,
	"system": SYSTEM,
}

// LookupIdentifier used to determinate whether identifier is keyword nor not
func LookupIdentifier(identifier string) Type {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
