/*
 * Copyright 2020 The SealEVM Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package instructions

import (
	"SealEVM/opcodes"
)

func loadComparision() {
	instructionTable[opcodes.LT] = opCodeInstruction{
		doAction: ltAction,
		minStackDepth: 2,
		enabled: true,
	}

	instructionTable[opcodes.GT] = opCodeInstruction {
		doAction: gtAction,
		minStackDepth: 2,
		enabled: true,
	}

	instructionTable[opcodes.SLT] = opCodeInstruction{
		doAction: sltAction,
		minStackDepth: 2,
		enabled: true,
	}

	instructionTable[opcodes.SGT] = opCodeInstruction{
		doAction: sgtAction,
		minStackDepth: 2,
		enabled: true,
	}

	instructionTable[opcodes.EQ] = opCodeInstruction{
		doAction: eqAction,
		minStackDepth: 2,
		enabled: true,
	}

	instructionTable[opcodes.ISZERO] = opCodeInstruction{
		doAction: isZeroAction,
		minStackDepth: 1,
		enabled: true,
	}
}

func ltAction(setting *instructionsSetting) ([]byte, error) {
	x, _ := setting.stack.Pop()
	y := setting.stack.Peek()

	if x.LT(y) {
		y.SetUint64(1)
	} else {
		y.SetUint64(0)
	}
	return nil, nil
}

func gtAction(setting *instructionsSetting) ([]byte, error) {
	x, _ := setting.stack.Pop()
	y := setting.stack.Peek()

	if x.GT(y) {
		y.SetUint64(1)
	} else {
		y.SetUint64(0)
	}
	return nil, nil
}

func sltAction(setting *instructionsSetting) ([]byte, error) {
	x, _ := setting.stack.Pop()
	y := setting.stack.Peek()

	if x.SLT(y) {
		y.SetUint64(1)
	} else {
		y.SetUint64(0)
	}
	return nil, nil
}

func sgtAction(setting *instructionsSetting) ([]byte, error) {
	x, _ := setting.stack.Pop()
	y := setting.stack.Peek()

	if x.SGT(y) {
		y.SetUint64(1)
	} else {
		y.SetUint64(0)
	}
	return nil, nil
}

func eqAction(setting *instructionsSetting) ([]byte, error) {
	x, _ := setting.stack.Pop()
	y := setting.stack.Peek()

	if x.EQ(y) {
		y.SetUint64(1)
	} else {
		y.SetUint64(0)
	}
	return nil, nil
}

func isZeroAction(setting *instructionsSetting) ([]byte, error) {
	x := setting.stack.Peek()

	x.IsZero()
	return nil, nil
}
