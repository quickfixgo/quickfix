package internal

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestWaitNormal(t *testing.T) {
	bucket := New(time.Second*1, 2)

	for i := 0; i < 10; i++ {
		bucket.Wait()
		t.Log(time.Now(), "|", i)
	}
}

func TestWaitForTimeout(t *testing.T) {
	bucket := New(time.Second*3, 3)

	for i := 0; i < 10; i++ {
		bucket.WaitForTimeout(time.Second)
		t.Log(time.Now(), "|", i)
	}

	for i := 0; i < 10; i++ {
		bucket.WaitForTimeout(time.Second * 5)
		t.Log(time.Now(), "|", i)
	}
}

func TestWaitRandom(t *testing.T) {
	bucket := New(time.Second*1, 3)
	for i := 0; i < 10; i++ {
		bucket.Wait()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		t.Log(time.Now(), "|", i)
	}
}

// func TestWaitConcurrenyOld(t *testing.T) {
// 	bucket := NewRateLimiter(8)

// 	for i := 0; i < 5; i++ {
// 		bucket.WaitRateLimit()
// 		t.Log("1:", time.Now(), "|", i)
// 	}
// 	time.Sleep(time.Millisecond * 900)
// 	go func() {
// 		for i := 0; i < 55; i++ {
// 			bucket.WaitRateLimit()
// 			t.Log("2:", time.Now(), "|", i)
// 		}
// 	}()

// 	time.Sleep(time.Second * 30)
// }

func TestWaitConcurreny1(t *testing.T) {
	bucket := New(time.Second, 8)

	for i := 0; i < 5; i++ {
		bucket.Wait()
		t.Log("1:", time.Now(), "|", i)
	}
	time.Sleep(time.Millisecond * 900)
	go func() {
		for i := 0; i < 55; i++ {
			bucket.Wait()
			t.Log("1:", time.Now(), "|", i)
		}
	}()

	time.Sleep(time.Second * 30)
}

func write(messageOut chan string) {
	for {
		text := <-messageOut
		fmt.Println(time.Now(), text)
		time.Sleep(time.Millisecond * time.Duration(100))
	}
}

func TestWaitConcurreny2(t *testing.T) {
	bucket := New(time.Second*2, 5)
	messageOut := make(chan string)
	go write(messageOut)
	go func() {
		for i := 0; i < 20; i++ {
			bucket.Wait()
			messageOut <- fmt.Sprintf("No.1 Worker, task:%v", i)
		}
	}()
	go func() {
		for i := 0; i < 20; i++ {
			bucket.Wait()
			messageOut <- fmt.Sprintf("No.2 Worker, task:%v", i)
		}
	}()
	time.Sleep(time.Second * 30)
}
