package tree

import (
	"github.com/ecodeclub/ekit/internal/iterator"
	"github.com/ecodeclub/ekit/tuple/pair"
)

type rbTreeIterator[K any, V any] struct {
	rbTree   *RBTree[K, V]
	currNode *rbNode[K, V]
	nxtNode  *rbNode[K, V]
	err      error
}

func (iter *rbTreeIterator[K, V]) isValidIterator() bool {
	return iter.err == nil
}

func (iter *rbTreeIterator[K, V]) Next() bool {
	if !iter.isValidIterator() {
		return false
	}
	iter.currNode = iter.nxtNode
	if iter.currNode != nil {
		iter.nxtNode = iter.currNode.getNext()
		return true
	}
	iter.err = ErrRBTreeIteratorNoNext
	return false
}

func (iter *rbTreeIterator[K, V]) Get() (kvPair pair.Pair[K, V]) {
	if !iter.isValidIterator() {
		return
	}
	if iter.currNode == nil {
		iter.err = ErrRBTreeIteratorInvalid
		return
	}
	kvPair = pair.NewPair(iter.currNode.key, iter.currNode.value)
	return
}

func (iter *rbTreeIterator[K, V]) Err() error {
	return iter.err
}

func (iter *rbTreeIterator[K, V]) Valid() bool {
	return iter.currNode != nil
}

func (iter *rbTreeIterator[K, V]) Delete() {
	if !iter.isValidIterator() {
		return
	}
	if !iter.currNode.isValidNode() {
		iter.err = ErrRBTreeIteratorInvalid
		return
	}
	iter.rbTree.deleteNode(iter.currNode)
}

func newRBTreeIterator[K any, V any](rbTree *RBTree[K, V], rbNode *rbNode[K, V]) iterator.Iterator[pair.Pair[K, V]] {
	iter := &rbTreeIterator[K, V]{
		rbTree:   rbTree,
		currNode: rbNode,
	}
	iter.nxtNode = iter.currNode.getNext()
	return iter
}
