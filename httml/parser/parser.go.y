%{
package parser
import (
	"fmt"
)

type Token struct {
    Token   int
    Literal string
}

type Elements []Element
type Element interface{}
type TagElement struct {
	Name Token
}
type TextElement struct {
	Text string
}
type CommentElement struct {
	Text string
}
%}
%union{
	elements Elements
	element Element
    token Token
    expr  Element
}
%type<elements> program elements
%type<element> element
%type<expr> tag_element text_element comment_element
%token<token> LEFT_COMMENT RIGHT_COMMENT COMMENT_TEXT TEXT_CONTENT TAG_NAME GT LT
%%
program
    : elements
    {
        $$ = $1
        yylex.(*lexer).result = $$
    }
elements
	: element elements
	{
		fmt.Printf("prev %#v %#v\n", $$, $1)
		$$ = append($$, $1)
	}
	| element
	{
		fmt.Printf("prev %#v %#v\n", $$, $1)
		$$ = append($$, $1)
	}
element
    : tag_element
    {
        $$ = $1
    }
    | text_element
    {
        $$ = $1
    }
	| comment_element
	{
		$$ = $1
	}
tag_element
	: LT TAG_NAME GT
	{
		$$ = TagElement{ Name: $2 }
	}
text_element
	: TEXT_CONTENT
	{
		$$ = TextElement{ Text: $1.Literal }
	}
comment_element
	: LEFT_COMMENT COMMENT_TEXT RIGHT_COMMENT
	{
		$$ = CommentElement{ Text: $2.Literal }
	}
	| LEFT_COMMENT RIGHT_COMMENT
	{
		$$ = CommentElement{}
	}
%%
