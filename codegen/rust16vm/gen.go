package rust16vm

import (
	"fmt"
	"math/bits"
	"stag/shunting_yard"
	"strings"
)

const Bits = 16

type Reg uint8

const (
	A Reg = iota
	B
	C
	M
	BP
	SP
	PC
	FLAGS
)

func (r Reg) String() string {
	switch r {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "M"
	case 4:
		return "BP"
	case 5:
		return "SP"
	case 6:
		return "PC"
	case 7:
		return "FLAGS"
	default:
		panic(fmt.Sprintf("unknown register %d", r))
	}
}

type vmCtx struct {
	errs []error
	//registers [8]uint16 // A B C M BP SP PC FLAGS
}

type regStack struct {
	arr []Reg
}

func (s *regStack) pop() Reg {
	value := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]
	return value
}

func (s *regStack) push(r Reg) {
	s.arr = append(s.arr, r)
}

func Generate(ast []shunting_yard.Statement) string {
	// // initialize the context will all the registers empty
	ctx := &vmCtx{}
	asm := strings.Builder{}

	usedRegs := &regStack{}

	for _, stmt := range ast {
		switch v := stmt.(type) {
		case *shunting_yard.BinaryOperation:
			resolveInnerBinOp(ctx, usedRegs, v, &asm)
		}
	}

	return asm.String()
}

func resolveInnerBinOp(ctx *vmCtx, s *regStack, v *shunting_yard.BinaryOperation, asm *strings.Builder) {
	lhsValue, lhsIsTerminal := v.Lhs.(*shunting_yard.Number)
	if !lhsIsTerminal {
		resolveInnerBinOp(ctx, s, v.Lhs.(*shunting_yard.BinaryOperation), asm)
	} else {
		emitMov(ctx, A, uint16(lhsValue.Value), asm)
		s.push(A)
	}

	rhsValue, rhsIsTerminal := v.Rhs.(*shunting_yard.Number)
	if !rhsIsTerminal {
		resolveInnerBinOp(ctx, s, v.Rhs.(*shunting_yard.BinaryOperation), asm)
	} else {
		emitMov(ctx, B, uint16(rhsValue.Value), asm)
		s.push(B)
	}

	rhs := s.pop()
	lhs := s.pop()

	switch v.Op {
	case shunting_yard.Add:
		emitArithRegReg("ADDR", C, lhs, rhs, asm)
	case shunting_yard.Sub:
		emitArithRegReg("ADDR", C, lhs, rhs, asm)
	case shunting_yard.Mul:
		emitArithRegReg("MULR", C, lhs, rhs, asm)
	case shunting_yard.Div:
		emitArithRegReg("DIVR", C, lhs, rhs, asm)
	}

	s.push(C)
}

func emitMov(vm *vmCtx, reg Reg, value uint16, asm *strings.Builder) {
	if !fitInOperation(value, 9) {
		vm.errs = append(vm.errs, fmt.Errorf("%w: max %d, got %d",
			errNumericValueOutOfBounds, 0b111111111, value))
	}

	asm.WriteString(fmt.Sprintf("MOV %s, #%d\n", reg.String(), value))
}

func emitArithRegReg(op string, dstReg Reg, fstReg Reg, sndReg Reg, asm *strings.Builder) {
	asm.WriteString(fmt.Sprintf("%s %s, %s, %s\n", op, dstReg.String(), fstReg.String(), sndReg.String()))
}

func fitInOperation(value uint16, bitAmount int) bool {
	return bits.LeadingZeros16(value) >= (Bits - bitAmount)
}
