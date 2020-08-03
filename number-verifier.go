package main

import (
	"flag"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gosuri/uilive"
	"github.com/upmasked/number-verifier/providers"
	"time"
)

func main() {
	var (
		selectedProvider = flag.String("provider", "upmasked", "a string")
		err              error
		number           string
		messages         []string
	)
	flag.Parse()

	provider := getProvider(selectedProvider)
	numbers, err := provider.GetNumbers()
	if err != nil {
		fmt.Println("error when getting numbers:", err)
		return
	}
	if len(numbers) == 0 {
		fmt.Println("error: no numbers were found for " + provider.GetProvider().Name)
		return
	}

	err = survey.AskOne(&survey.Select{
		Message: "Choose a number:",
		Options: numbers,
	}, &number)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return
	}

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for {
		messages, err = provider.GetMessages(number)
		if err != nil {
			fmt.Println("error getting messages: " + err.Error())
			return
		}

		for i := 0; i < len(messages); i++ {
			_, _ = fmt.Fprintf(writer, "%s\n", messages[i])
		}

		time.Sleep(time.Second * 5)
	}
}

// getProvider selects a provider based on the input string.
func getProvider(selectedProvider *string) providers.Providers {
	var provider providers.Providers

	switch *selectedProvider {
	case "smsreceivefree":
		provider = providers.SMSReceiveFree{}
	case "upmasked":
		provider = providers.Upmasked{}
	default:
		fmt.Println("Provider '" + *selectedProvider + "' not found, falling back to default provider SMSReceiveFree")
		provider = providers.SMSReceiveFree{}
	}

	return provider
}
