package lexer

import (
	"fmt"
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexCardinality2(l *Lexer) LexFn {
	l.SkipWhitespace()

	if(strings.HasPrefix(l.InputToEnd(), lexertoken.LEFT_BRACKET)){
		l.Pos += len(lexertoken.LEFT_BRACKET)
		l.Emit(lexertoken.TOKEN_LEFT_BRACKET, true)
		fmt.Println("TAG: LEFT_BRACKET")
	}

	l.SkipWhitespace()

	if (strings.HasPrefix(l.InputToEnd(), "M")){
		l.Pos += len("M")
		l.Emit(lexertoken.TOKEN_CARDINALITY, true)
		fmt.Println("TAG: CARDINALITY2")
	} else if (strings.HasPrefix(l.InputToEnd(), "N")){
		l.Pos += len("N")
		l.Emit(lexertoken.TOKEN_CARDINALITY, true)
		fmt.Println("TAG: CARDINALITY2")
	} else {
		var ch rune

		for {
			ch = rune(l.Input[l.Pos])
			if !unicode.IsDigit(ch) {
				l.SkipWhitespace()
				if (strings.HasPrefix(l.InputToEnd(), lexertoken.COMMA)){
					l.Emit(lexertoken.TOKEN_CARDINALITY, true)
					l.Emit(lexertoken.TOKEN_COMMA, true)
					break
				}else{
					fmt.Println("comma no there")
					return nil
				}
			}

			if l.IsEOF() {
				return nil
			}
			l.Inc()
		}
	}

	l.SkipWhitespace()
	if (strings.HasPrefix(l.InputToEnd(), "1")){
		l.Pos += len("1")
		l.Emit(lexertoken.TOKEN_PARTICIPATION, true)
		fmt.Println("TAG: PARTICIPATION2")
	} else if (strings.HasPrefix(l.InputToEnd(), "0")){
		l.Pos += len("0")
		l.Emit(lexertoken.TOKEN_PARTICIPATION, true)
		fmt.Println("TAG: PARTICIPATION2")
	} else{
		fmt.Println("invalid participation")
		return nil
	}

	l.SkipWhitespace()

	if (strings.HasPrefix(l.InputToEnd(), lexertoken.RIGHT_BRACKET)){
		l.Pos += len(lexertoken.RIGHT_BRACKET)
		l.Emit(lexertoken.TOKEN_RIGHT_BRACKET, true)
		fmt.Println("TAG: RIGHT_BRACKET")
	}

	return LexTable2
}