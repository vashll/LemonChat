package msg

import (
	"sync"
	"container/list"
	"errors"
)

type MessageQueue struct {
	lock       sync.Mutex
	maxSize    int
	linkedList *list.List
}

func NewMsgQueue(maxSize int) *MessageQueue {
	mq := new(MessageQueue)
	mq.maxSize = maxSize
	mq.linkedList = list.New()
	return mq
}

func (mq *MessageQueue) Put(v interface{}) error {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if !mq.IsFull() {
		mq.linkedList.PushBack(v)
	} else {
		return errors.New("Message queue is full.")
	}
	return errors.New("Put element to message queue failed.")
}

func (mq *MessageQueue) Get() (interface{}, error) {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if !mq.IsEmpty() {
		e := mq.linkedList.Front()
		mq.linkedList.Remove(e)
		return e, nil
	}
	return nil, errors.New("Message queue is empyt.")
}

func (mq *MessageQueue) IsFull() bool {
	return mq.linkedList.Len() >= mq.maxSize
}

func (mq *MessageQueue) IsEmpty() bool {
	return mq.linkedList.Len() == 0
}
