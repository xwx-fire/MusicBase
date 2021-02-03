package queue_test

import (
	"fmt"
	"strconv"
	"testing"
)

const  MAXQSIZE = 10

type Queue struct {
	Base *[MAXQSIZE]int
	Front int
	Rear int
}

func QueueInit() *Queue{
	q := Queue{
		Base:  new([MAXQSIZE]int),
		Front: 0,
		Rear:  0,
	}
	return &q
}

func (q *Queue)EnQueue(a int){
	fmt.Printf("队列中已有元素%d\n",q.QueueLength())
	if (q.Rear + 1)%MAXQSIZE == q.Front{
		fmt.Println("队列已满")
		return
	}
	q.Base[q.Rear] = a
	q.Rear = (q.Rear + 1)%MAXQSIZE
}

func (q *Queue)DeQueue(){
	fmt.Printf("队列中已有元素%d\n",q.QueueLength())
	if q.Rear == q.Front{
		fmt.Println("队列为空")
		return
	}
	q.Front = (q.Front + 1)%MAXQSIZE
}

func (q *Queue)QueueLength()int{
	return (q.Rear - q.Front + MAXQSIZE)%MAXQSIZE
}

func (q *Queue)QueuePrint() string{
	str := ""
	if q.QueueLength() > 0{
		for i := q.Front; i <= q.QueueLength();i++{
			str += strconv.Itoa(q.Base[i])+"->"
		}
	}
	fmt.Println(q.Base)
	return str
}

func TestQueue(t *testing.T){
	q := QueueInit()
	for i := 1 ; i < 10; i++{
		q.EnQueue(i)
	}
	q.DeQueue()
	q.DeQueue()
	t.Log(q.QueuePrint())
}
