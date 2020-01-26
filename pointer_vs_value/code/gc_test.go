package code

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
)

//const timeout = 10 * time.Second
const timeout = 100 * time.Second

func TestTeraValueReceiverRoutine(t *testing.T) {

	runRoutine(func(o tera) {
		calcTera(o)
	}, "TestTeraValueReceiverRoutine")
}

func TestTeraPointerReceiverRoutine(t *testing.T) {

	runRoutine(func(o tera) {
		calcTeraP(&o)
	}, "TestTeraPointerReceiverRoutine")

}

func TestTeraInterfaceReceiverRoutine(t *testing.T) {

	runRoutine(func(o tera) {
		calcIfc(&o)
	}, "TestTeraInterfaceReceiverRoutine")

}

func TestTeraValueReceiverNative(t *testing.T) {

	runNative(func(o tera) {
		calcTera(o)
	}, "TestTeraValueReceiverNative")
}

func TestTeraPointerReceiverNative(t *testing.T) {

	runNative(func(o tera) {
		calcTeraP(&o)
	}, "TestTeraPointerReceiverNative")

}
func TestTeraInterfaceReceiverNative(t *testing.T) {

	runNative(func(o tera) {
		calcIfc(&o)
	}, "TestTeraInterfaceReceiverNative")

}

func runRoutine(exec func(o tera), name string) {
	ch := make(chan tera, 1)

	ctx, cnl := context.WithCancel(context.Background())

	go func() {
		<-time.After(timeout)
		cnl()
	}()

	go func() {
		start := time.Now()
		for {
			select {
			case <-ctx.Done():
				duration := time.Now().Sub(start).Milliseconds()
				fmt.Println(fmt.Printf("Time passed [%v]", duration))
				close(ch)
				return
			default:
				ch <- createTera()
			}
		}
	}()

	i := 0

	for o := range ch {
		exec(o)
		i++
	}

	println(fmt.Sprintf("invocations = %v", i))
	trackGC(name)
}

// NOTE : results are more diverse with the runRoutine.
// keeping this as a reference
func runNative(exec func(o tera), name string) {

	ctx, cnl := context.WithCancel(context.Background())

	go func() {
		<-time.After(timeout)
		cnl()
	}()

	i := 0

	start := time.Now()
	for {
		select {
		case <-ctx.Done():
			duration := time.Now().Sub(start).Milliseconds()
			fmt.Println(fmt.Printf("Time passed [%v]", duration))
			println(fmt.Sprintf("invocations = %v", i))
			trackGC(name)
			return
		default:
			o := createTera()
			exec(o)
			i++
		}
	}

}

func trackGC(caller string) {
	runtime.GC()
	runtime.GC()
	stats := debug.GCStats{}
	debug.ReadGCStats(&stats)
	println(fmt.Sprintf("\n%s , %v , %v , %v , %v\n", caller, stats.NumGC, sum(stats.Pause), stats.PauseQuantiles, stats.PauseTotal))
	//debug.FreeOSMemory()
}

func sum(durations []time.Duration) int64 {
	var f int64 = 0
	for _, d := range durations {
		f += d.Microseconds()
	}
	return f
}
