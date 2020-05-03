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

package common

import (
	"SealEVM/evmInt256"
)

const (
	hashLength = 32
)

func EVMIntToHashBytes(i *evmInt256.Int) [hashLength]byte {
	iBytes := i.Bytes()
	iLen := len(iBytes)

	var hash [hashLength]byte
	if iLen > hashLength {
		copy(hash[:], iBytes[iLen - hashLength:])
	} else {
		copy(hash[hashLength - iLen:], iBytes)
	}

	return hash
}

func HashBytesToEVMInt(hash [hashLength]byte) (*evmInt256.Int, error) {

	i := evmInt256.New(0)
	i.SetBytes(hash[:])

	return i, nil
}