package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func (tree *ParserT) validateExpression() error {
	// compile data types and check for errors in the AST
	//
	// first walk to ensure we have a:
	// value, expression, value, expression, value....etc

	if tree.charPos < 0 {
		tree.charPos = 0
	}

	if len(tree.ast) == 0 {
		return fmt.Errorf("missing expression: '%s'", string(tree.expression[:tree.charPos]))
	}

	if len(tree.ast) == 1 &&
		(tree.ast[0].key == symbols.Bareword || tree.ast[0].key == symbols.SubExpressionBegin ||
			tree.ast[0].key == symbols.QuoteSingle || tree.ast[0].key == symbols.QuoteDouble) {
		if tree.charPos+1 > len(tree.expression) {
			tree.charPos--
		}
		return fmt.Errorf("not an expression: '%s'", string(tree.expression[:tree.charPos+1]))
	}

	var expectValue bool

	for tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		prev := tree.prevSymbol()
		node := tree.ast[tree.astPos]
		next := tree.nextSymbol()

		expectValue = !expectValue

		// check for errors raised by the parser
		if node.key < symbols.DataValues {
			return raiseError(tree.expression, node, 0, errMessage[node.key])
		}

		// check each operation has a left side and right side data value
		if expectValue {
			if (node.key < symbols.DataValues) ||
				node.key > symbols.Operations {
				return raiseError(tree.expression, node, 0, "expecting a data value")
			}

			if node.dt != nil {
				continue
			}
			primitive, err := node2primitive(node)
			if err != nil {
				return err
			}
			node.dt = primitive

		} else {

			switch {
			case node.key < symbols.Operations:
				return raiseError(tree.expression, node, 0, "expecting an operation")

			case node.key >= symbols.Add:
				if prev == nil || (prev.key != symbols.Number && prev.key != symbols.Calculated && prev.key != symbols.SubExpressionBegin && prev.key != symbols.Scalar) ||
					next == nil || (next.key != symbols.Number && next.key != symbols.Calculated && next.key != symbols.SubExpressionBegin && next.key != symbols.Scalar) {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("cannot %s non-numeric data types", node.key))
				}

			case node.key >= symbols.Elvis:
				if prev == nil || prev.key == symbols.Bareword ||
					next == nil || next.key == symbols.Bareword {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("cannot %s barewords", node.key))
				}
			}
		}

	}

	if !expectValue {
		return raiseError(tree.expression, tree.ast[len(tree.ast)-1], 0, "unexpected end of expression")
	}

	return nil
}
