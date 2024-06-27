package expressions

import (
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
)

func expMultiply(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	lv, rv, err := validateNumericalDataTypes(tree, left, right)
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Number,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Number, lv*rv),
	})
}

func expDivide(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	lv, rv, err := validateNumericalDataTypes(tree, left, right)
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Number,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Number, lv/rv),
	})
}
