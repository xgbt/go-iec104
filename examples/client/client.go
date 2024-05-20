package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xgbt/go-iec104"
)

const (
	serverAddress = "127.0.0.1:2404"
)

type handler struct{}

func (h handler) GeneralInterrogationHandler(apdu *iec104.APDU) error {
	for _, signal := range apdu.Signals {
		fmt.Printf("%f ", signal.Value)
	}
	fmt.Println()
	return nil
}

func (h handler) CounterInterrogationHandler(apdu *iec104.APDU) error {
	for _, signal := range apdu.Signals {
		fmt.Printf("%f ", signal.Value)
	}
	fmt.Println()
	return nil
}

func (h handler) ReadCommandHandler(apdu *iec104.APDU) error {
	return nil
}

func (h handler) ClockSynchronizationHandler(apdu *iec104.APDU) error {
	return nil
}

func (h handler) TestCommandHandler(apdu *iec104.APDU) error {
	return nil
}

func (h handler) ResetProcessCommandHandler(apdu *iec104.APDU) error {
	return nil
}

func (h handler) DelayAcquisitionCommandHandler(apdu *iec104.APDU) error {
	return nil
}

func (h handler) APDUHandler(apdu *iec104.APDU) error {
	for _, signal := range apdu.Signals {
		fmt.Printf("%f ", signal.Value)
	}
	fmt.Println()
	return nil
}

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	iec104.SetLogger(logger)

	option, err := iec104.NewClientOption(serverAddress, &handler{})
	if err != nil {
		panic(any(err))
	}
	client := iec104.NewClient(option)
	if err := client.Connect(); err != nil {
		panic(any(err))
	}
	defer client.Close()

	client.SendSetPointShortFloat(10, 114.514)

	go func() {
		time.Sleep(5 * time.Second)
		client.SendTestFrame()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		client.SendGeneralInterrogation()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		client.SendCounterInterrogation()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		if err := client.SendSingleCommand(iec104.IOA(1), true /* close */); err != nil {
			panic(any(err))
		}
		if err := client.SendSingleCommand(iec104.IOA(1), false /* close */); err != nil {
			panic(any(err))
		}
		if err := client.SendDoubleCommand(iec104.IOA(1), true /* close */); err != nil {
			panic(any(err))
		}
		if err := client.SendDoubleCommand(iec104.IOA(1), false /* close */); err != nil {
			panic(any(err))
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Printf("Connected: %v\n", client.IsConnected())
	}()

	time.Sleep(30 * time.Minute)
}
