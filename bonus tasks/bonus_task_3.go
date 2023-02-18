package main
import "fmt"

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

func (l *linkedList) prepand (n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData(){
	toPrint := l.head
	for l.length != 0{
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	} 
	fmt.Println("\n")
}

func (l *linkedList) deleteWithValeu(value int){
	if l.length == 0{
		return
	}
	if l.head.data == value{
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value{
		if previousToDelete.next.next == nil{
			return
 		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}



type Stack struct{
	items []int
}

// push
func (s *Stack) Push(i int){
	s.items = append(s.items, i)
}

// peek
func (s *Stack) Peek() int{
	return s.items[0]
}

// pop
func (s *Stack) Pop() int{
	l := len(s.items)-1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// clean
func (s *Stack) Clean() {
	s.items = nil
}

// contains
func (s *Stack) Contains(e int) bool {
	ss := s.items
    for _, a := range ss {
        if a == e {
            return true
        }
    }
    return false
}

// Increment
func (s *Stack) Increment(){
	for i := 0; i < len(s.items); i++ {
		s.items[i]++
	}
}

// print
func (s *Stack) Print(){
	for i := 0; i < len(s.items); i++ {
		fmt.Print(s.items[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

// print reverse
func (s *Stack) PrintReverse(){
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Print(s.items[i])
		fmt.Print(" ")
	}
	fmt.Println()
}


func main(){
	myList := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:18}
	node3 := &node{data:16}
	node4 := &node{data:11}
	node5 := &node{data:7}
	node6 := &node{data:2}
	myList.prepand(node1)
	myList.prepand(node2)
	myList.prepand(node3)
	myList.prepand(node4)
	myList.prepand(node5)
	myList.prepand(node6)
	myList.printListData()
	myList.deleteWithValeu(16)
	myList.printListData()


	myStuck := Stack{}
	myStuck.Push(100)
	myStuck.Push(200)
	myStuck.Push(300)
	fmt.Println(myStuck)
	fmt.Println(myStuck.Pop())
	fmt.Println(myStuck)
	fmt.Println(myStuck.Peek())
	fmt.Println(myStuck.Contains(143))
	myStuck.Increment()
	fmt.Println(myStuck)
	myStuck.Print()
	myStuck.PrintReverse()
	myStuck.Clean()
	fmt.Println(myStuck)
}
