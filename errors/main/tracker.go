package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(),"test")
	_ = tr.Event(context.Background(),"test")
	_ = tr.Event(context.Background(),"test")
	ctx, cancel := context.WithDeadline(context.Background(),time.Now().Add(5*time.Second))
	defer cancel()
	tr.ShutDown(ctx)
}

func NewTracker() *Tracker{
	return &Tracker{
		ch: make(chan string,10),
	}
}


type Tracker struct{
	ch chan string
	stop chan struct{}
}

func (t *Tracker) Event(ctx context.Context,data string) error{
	select {
		case t.ch <- data:
			return nil
		case <-ctx.Done():
			return ctx.Err()
	}
}


func (t *Tracker) Run(){
	for data := range t.ch{
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) ShutDown(ctx context.Context){
	close(t.ch)
	select {
		case <-t.stop:
		case <-ctx.Done():
	}
}
