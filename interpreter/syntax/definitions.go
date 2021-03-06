package syntax

import "github.com/Guilherme-De-Marchi/orwells-language/interpreter/token"

var (
	VAR_MODEL_UNITARY = SyntaxModel{
		{token.VAR_TOKEN},
		{token.IDENTIFIER_TOKEN},
		{token.ASSIGNMENT_TOKEN},
		{token.IDENTIFIER_TOKEN, token.STRING_LITERAL_TOKEN, token.INTEGER_LITERAL_TOKEN, token.FLOATING_LITERAL_TOKEN},
		{token.INSTRUCTION_DELIMITER_TOKEN},
	}

	VAR_MODEL_BINARY = SyntaxModel{
		{token.VAR_TOKEN},
		{token.IDENTIFIER_TOKEN},
		{token.ASSIGNMENT_TOKEN},
		{token.IDENTIFIER_TOKEN, token.STRING_LITERAL_TOKEN, token.INTEGER_LITERAL_TOKEN, token.FLOATING_LITERAL_TOKEN},
		{token.SUM_TOKEN, token.SUBTRACTION_TOKEN},
		{token.IDENTIFIER_TOKEN, token.STRING_LITERAL_TOKEN, token.INTEGER_LITERAL_TOKEN, token.FLOATING_LITERAL_TOKEN},
		{token.INSTRUCTION_DELIMITER_TOKEN},
	}

	IF_MODEL = SyntaxModel{
		{token.IF_TOKEN},
		{token.IDENTIFIER_TOKEN, token.STRING_LITERAL_TOKEN, token.INTEGER_LITERAL_TOKEN, token.FLOATING_LITERAL_TOKEN},
		{token.EQUAL_TOKEN},
		{token.IDENTIFIER_TOKEN, token.STRING_LITERAL_TOKEN, token.INTEGER_LITERAL_TOKEN, token.FLOATING_LITERAL_TOKEN},
		{token.INSTRUCTION_DELIMITER_TOKEN},
	}

	ENDIF_MODEL = SyntaxModel{
		{token.ENDIF_TOKEN},
		{token.INSTRUCTION_DELIMITER_TOKEN},
	}

	EXEC_MODEL = SyntaxModel{
		{token.EXEC_TOKEN},
		{token.IDENTIFIER_TOKEN},
		{token.LEFT_PARENTESIS_TOKEN},
		{token.GENERAL_TOKEN},
		{token.RIGHT_PARENTESIS_TOKEN},
		{token.INSTRUCTION_DELIMITER_TOKEN},
	}

	EXEC_MODEL_NO_ARGUMENTS = SyntaxModel{
		{token.EXEC_TOKEN},
		{token.IDENTIFIER_TOKEN},
		{token.LEFT_PARENTESIS_TOKEN},
		{token.RIGHT_PARENTESIS_TOKEN},
		{token.INSTRUCTION_DELIMITER_TOKEN},
	}

	SyntaxModels = map[*token.Token][]SyntaxModel{
		token.VAR_TOKEN:   {VAR_MODEL_UNITARY, VAR_MODEL_BINARY},
		token.IF_TOKEN:    {IF_MODEL},
		token.ENDIF_TOKEN: {ENDIF_MODEL},
		token.EXEC_TOKEN:  {EXEC_MODEL, EXEC_MODEL_NO_ARGUMENTS},
	}
)
