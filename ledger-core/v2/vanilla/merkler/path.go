//
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package merkler

import (
	"math"
	"math/bits"

	"github.com/insolar/assured-ledger/ledger-core/v2/vanilla/args"
)

type PathBuilder struct {
	count   uint
	levels  []pathLevel
	stubbed bool
}

type pathLevel struct {
	bitMask    uint
	nodeValueL uint
	nodeValueR uint
}

func NewPathBuilder(leafCount uint, stubbed bool) PathBuilder {
	switch leafCount {
	case 0:
		panic("illegal value")
	case 1, 2:
		return PathBuilder{leafCount, nil, stubbed}
	}

	depth := uint8(bits.Len(leafCount - 1))
	levels := make([]pathLevel, depth-1)

	for i := range levels {
		bitMask := uint(2) << uint8(i)
		levels[i].bitMask = bitMask
	}

	switch {
	case args.IsPowerOfTwo(leafCount):
		perfectPath(leafCount, levels)
	case stubbed:
		stubbedPath(leafCount, levels)
	default:
		unstubbedPath(leafCount, levels)
	}

	return PathBuilder{leafCount, levels, stubbed}
}

func perfectPath(leafCount uint, levels []pathLevel) {
	nodeCount := leafCount
	for i := 1; i < len(levels); i++ {
		levels[i].nodeValueR = nodeCount
		nodeCount++
	}
}

func unstubbedPath(leafCount uint, levels []pathLevel) {
	lastLeafIndex := leafCount - 1

	invertMask := ^(uint(math.MaxUint64&^1) << UnbalancedBitCount(lastLeafIndex, uint8(len(levels)+1)))
	invertedLastBalanced := lastLeafIndex ^ invertMask // | invertMask>>1
	//if invertedLastBalanced >= leafCount {
	//	panic("illegal state")
	//}

	lastRightmost := lastLeafIndex
	nodeIndex := leafCount

	for i, level := range levels {
		if lastLeafIndex&level.bitMask == 0 {
			levels[i].nodeValueL = lastRightmost

			leftGapSibling := invertedLastBalanced ^ level.bitMask
			//if leftGapSibling >= leafCount {
			//	panic("illegal state")
			//}
			if indexR := siblingIndex(leftGapSibling, level.bitMask); indexR >= leafCount {
				levels[i].nodeValueR = nodeIndex
				nodeIndex++
			}
			continue
		}

		indexL := siblingIndex(lastLeafIndex, level.bitMask)
		if indexL >= leafCount {
			levels[i].nodeValueL = nodeIndex
			nodeIndex++
		}

		levels[i].nodeValueR = lastRightmost
		lastRightmost = nodeIndex
		nodeIndex++
	}

	if nodeCount := leafCount + uint(StackSequenceUnused(leafCount)) - 1; nodeCount != nodeIndex {
		panic("illegal state")
	}
}

func stubbedPath(leafCount uint, levels []pathLevel) {
	lastLeafIndex := leafCount - 1

	invertMask := ^(uint(math.MaxUint64&^1) << UnbalancedBitCount(lastLeafIndex, uint8(len(levels)+1)))
	invertedLastBalanced := lastLeafIndex ^ invertMask // | invertMask>>1
	//if invertedLastBalanced >= leafCount {
	//	panic("illegal state")
	//}

	nodeIndex := leafCount

	for i, level := range levels {
		if i == 0 && leafCount&1 == 0 {
			continue
		}

		if lastLeafIndex&level.bitMask == 0 {
			leftGapSibling := invertedLastBalanced ^ level.bitMask
			//if leftGapSibling >= leafCount {
			//	panic("illegal state")
			//}
			if indexR := siblingIndex(leftGapSibling, level.bitMask); indexR >= leafCount {
				levels[i].nodeValueR = nodeIndex
				nodeIndex++
			}

			levels[i].nodeValueL = nodeIndex
			nodeIndex++
			continue
		}

		indexL := siblingIndex(lastLeafIndex, level.bitMask)
		if indexL >= leafCount {
			levels[i].nodeValueL = nodeIndex
			nodeIndex++
		}

		levels[i].nodeValueR = nodeIndex
		nodeIndex++
	}

	// TODO check count
}

const StubNodeIndex = 0

type PathEntryFunc func(index uint, isLeaf, isRight bool)

func (v PathBuilder) WalkFor(index uint, nodeFn PathEntryFunc) {
	if v.count <= index {
		panic("illegal value")
	}

	// sibling leaf
	switch sibling := index ^ 1; {
	case sibling < v.count:
		nodeFn(sibling, true, index&1 == 0)
	case v.stubbed:
		nodeFn(StubNodeIndex, false, true)
	}

	for _, level := range v.levels {
		tIndex := siblingIndex(index, level.bitMask)

		isRightSibling := index&level.bitMask == 0

		switch {
		case tIndex < v.count:
			// the tree node was calculated as a part of the normal flow
			nodeFn(tIndex, false, isRightSibling)
			continue
		case isRightSibling:
			lastOfLeftBranch := index | (level.bitMask - 1)
			if lastOfLeftBranch >= v.count-1 {
				// right branch doesn't exist
				if v.stubbed {
					nodeFn(StubNodeIndex, false, true)
				}
				continue
			}
			// right branch is delayed
			tIndex = level.nodeValueR
		default:
			// left branch is delayed
			tIndex = level.nodeValueL
		}

		if tIndex == 0 {
			panic("illegal state")
		}
		nodeFn(tIndex, tIndex == v.count-1 && v.count&1 != 0, isRightSibling)
	}
}

func siblingIndex(index uint, bitMask uint) uint {
	tIndex := index ^ bitMask
	bitMask--
	tIndex |= bitMask
	tIndex += bitMask >> 1
	return tIndex
}