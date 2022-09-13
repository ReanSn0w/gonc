package nc_test

import (
	"testing"
	"time"

	"github.com/ReanSn0w/gonc/pkg/nc"
)

var (
	values       = []int{1, 19, 284, 92384, 2284, 2832, 1724}
	secondValues = []int{123, 23, 124, 365, 342, 1, 876}
)

func Test_SendNotification(t *testing.T) {
	counter := 0
	intTester := nc.NewSubsriber(func(i interface{}) {
		num, ok := i.(int)
		if !ok {
			t.Log("value is not a number")
			t.Fail()
		}

		if values[counter] != num {
			t.Log("compairing values failed")
			t.Fail()
		}

		counter++
	})

	nc.Default().Subscribe("int_tester", intTester)

	for i := 0; i < 7; i++ {
		time.Sleep(time.Second * time.Duration(i) / 5)
		nc.Default().Send("int_tester", values[i])
	}
}

func Test_DoneSubscriber(t *testing.T) {
	intTester := nc.NewSubsriber(func(i interface{}) {
		t.Log("channel is done")
		t.Fail()
	})

	nc.Default().Subscribe("int_tester", intTester)
	nc.Default().Unsubscribe("int_tester", intTester)

	for i := 0; i < 7; i++ {
		time.Sleep(time.Second * time.Duration(i) / 5)
		nc.Default().Send("int_tester", values[i])
	}
}

func Test_ManySubscribers(t *testing.T) {
	first := 0
	firstIntTester := nc.NewSubsriber(func(i interface{}) {
		num, ok := i.(int)
		if !ok {
			t.Log("value is not a number")
			t.Fail()
		}

		if values[first] != num {
			t.Log("compairing values failed")
			t.Fail()
		}

		first++
	})

	second := 0
	secondIntTester := nc.NewSubsriber(func(i interface{}) {
		num, ok := i.(int)
		if !ok {
			t.Log("value is not a number")
			t.Fail()
		}

		if values[second] != num {
			t.Log("compairing values failed")
			t.Fail()
		}

		second++
	})

	nc.Default().Subscribe("int_tester", firstIntTester)
	nc.Default().Subscribe("int_tester", secondIntTester)

	for i := 0; i < 7; i++ {
		time.Sleep(time.Second * time.Duration(i) / 5)
		nc.Default().Send("int_tester", values[i])
	}
}

func Test_ManuChannels(t *testing.T) {
	first := 0
	firstIntTester := nc.NewSubsriber(func(i interface{}) {
		num, ok := i.(int)
		if !ok {
			t.Log("value is not a number")
			t.Fail()
		}

		if values[first] != num {
			t.Log("compairing values failed")
			t.Fail()
		}

		first++
	})

	second := 0
	secondIntTester := nc.NewSubsriber(func(i interface{}) {
		num, ok := i.(int)
		if !ok {
			t.Log("value is not a number")
			t.Fail()
		}

		if secondValues[second] != num {
			t.Log("compairing values failed")
			t.Fail()
		}

		second++
	})

	nc.Default().Subscribe("int_tester", firstIntTester)
	nc.Default().Subscribe("second_tester", secondIntTester)

	for i := 0; i < 7; i++ {
		time.Sleep(time.Second * time.Duration(i) / 5)
		nc.Default().Send("int_tester", values[i])
		nc.Default().Send("second_tester", secondValues[i])
	}
}
