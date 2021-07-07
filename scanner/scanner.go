package scanner

import (
	"Glox/token"
	"fmt"
	"unicode"
)

type Scanner struct {
	source  []rune
	tokens  []token.Token
	errors  []string
	start   int
	current int
	line    int
	col     int
}

func NewScanner(source string) *Scanner {
	s := &Scanner{
		source:  []rune(source),
		tokens:  []token.Token{},
		errors:  []string{},
		start:   0,
		current: 0,
		line:    1,
		col:     0,
	}
	return s
}

func (s *Scanner) Errors() []string {
	return s.errors
}

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		// We are at the beginning og the next lexeme.
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, token.Token{Type: token.EOF, Lexeme: ""})
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case rune('('):
		s.addToken(token.LEFT_PAREN, "(")
	case rune(')'):
		s.addToken(token.RIGHT_PAREN, ")")
	case rune('{'):
		s.addToken(token.LEFT_BRACE, "{")
	case rune('}'):
		s.addToken(token.RIGHT_BRACE, "}")
	case rune(','):
		s.addToken(token.COMMA, ",")
	case rune('.'):
		s.addToken(token.DOT, ".")
	case rune('-'):
		s.addToken(token.MINUS, ".")
	case rune('+'):
		s.addToken(token.PLUS, "+")
	case rune(';'):
		s.addToken(token.SEMICOLON, ";")
	case rune('*'):
		s.addToken(token.STAR, "*")
	case rune('!'):
		if s.match('=') {
			s.addToken(token.BANG_EQUAL, "!=")
		} else {
			s.addToken(token.BANG, "!")
		}
	case rune('='):
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL, "==")
		} else {
			s.addToken(token.EQUAL, "=")
		}
	case rune('<'):
		if s.match('=') {
			s.addToken(token.LESS_EQUAL, "<=")
		} else {
			s.addToken(token.LESS, "<")
		}
	case rune('>'):
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL, ">=")
		} else {
			s.addToken(token.GREATER, ">")
		}
	case rune('/'):
		if s.match('/') {
			// A comment goes until the end of the line.
			for !s.isAtEnd() && s.peek() != rune('\n') {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH, "/")
		}
	case rune(' '), rune('\r'), rune('\t'):
		s.col += 1
	case rune('\n'):
		s.line += 1
		s.col = 0
	case rune('"'):
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			s.errors = append(s.errors, fmt.Sprintf("Ln %d, Col %d Unexpected character: '%c'", s.line, s.col, c))
		}
	}
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	if t, ok := keywords[string(text)]; ok {
		s.addToken(t, string(text))
	} else {
		s.addToken(token.IDENTIFIER, string(text))
	}
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	// Look for a fractional part.
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	s.addToken(token.NUMBER, string(s.source[s.start:s.current]))
}

func (s *Scanner) string() {
	for !s.isAtEnd() && s.peek() != '"' {
		if s.peek() == '\n' {
			s.line += 1
			s.col = 0
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.errors = append(s.errors, fmt.Sprintf("Ln %d, Col %d Unterminated string.", s.line, s.col))
		return
	}
	// The closing ".
	s.advance()

	// Trim the surrounding quotes.
	value := s.source[s.start+1 : s.current-1]
	s.addToken(token.STRING, string(value))
}

func (s *Scanner) match(c rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != c {
		return false
	}
	s.current += 1
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return rune(0)
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return rune(0)
	}
	return s.source[s.current+1]
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c) || c == rune('_')
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

func isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

func (s *Scanner) advance() rune {
	c := s.source[s.current]
	s.current += 1
	s.col += 1
	return c
}

func (s *Scanner) addToken(t token.TokenType, lexeme string) {
	tok := token.Token{
		Type:   t,
		Lexeme: lexeme,
		Line:   s.line,
		Col:    s.col,
	}
	s.tokens = append(s.tokens, tok)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
