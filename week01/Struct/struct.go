package Struct

type LinkedLink struct {
	head string
}

func (linkedLink *LinkedLink) Head() string {
	return linkedLink.head
}

func (linkedLink *LinkedLink) SetHead(head string) {
	linkedLink.head = head
}
