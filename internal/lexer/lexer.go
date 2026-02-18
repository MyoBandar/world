/*
Package lexer adopts character-based lexing;

it reads one character at a time and advance through the source like a tape reader.
*/
package lexer

/*
Lexer struct

	This is the machine that moves through characters.
*/
type Lexer struct {
	// input is the entire program
	input string
	// pos is where you are currently
	pos int
	// pos is where the reader is at currently
	readPos int
	// ch is the current character
	ch byte
	// line tracks errors
	line int
}

// New constructs a new Lexer
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		line:  1,
	}
	// read the first character to set the starting point
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.ch == 0 {
		return Token{
			Type: EOF,
			Line: l.line,
		}
	}

	if isLetter(l.ch) {
		word := l.readWord()

		switch word {
		case "display":
			return Token{
				Type:  DISPLAY,
				Value: word,
				Line:  l.line,
			}
		default:
			return Token{
				Type:  STRING,
				Value: word,
				Line:  l.line,
			}
		}
	}

	illegal := l.ch
	l.readChar()

	return Token{
		Type:  ILLEGAL,
		Value: string(illegal),
		Line:  l.line,
	}
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos++

	if l.ch == '\n' {
		l.line++
	}
}

func (l *Lexer) readWord() string {
	start := l.pos

	for isLetter(l.ch) {
		// println("CHAR:", string(l.ch))
		l.readChar()
	}

	return l.input[start:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}
