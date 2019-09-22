package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	// msg := Message  
	
	// failedMsg := FailedMessage {
	// 	ErrorMessage: "message intercepted by black rider",
	// 	OriginalMessage: Message{},
	// }

	// msgCh <- msg
	// errCh <- failedMsg 
	select {
	case receivedMsg := <- msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <- errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("no messages received at this time")
	}

	//original code - just to see that channels are working
	//above select block does the same thing 
	
	// msgCh <- msg
	// errCh <- failedMsg

	// fmt.Println(<- msgCh)
	// fmt.Println(<- errCh)
}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}