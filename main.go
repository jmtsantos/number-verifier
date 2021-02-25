package main

import (
	"fmt"
	"number-verifier/providers"
	"number-verifier/util"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gosuri/uilive"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Command line params options
	app               = kingpin.New("number-verifier", "SMS verification tool")
	flagDebug         = app.Flag("debug", "Enable debug mode.").Bool()
	flagProviderFirst = app.Flag("provider", "Select number provider first").Bool()
	flagListProviders = app.Flag("list-providers", "List providers").Bool()
)

func init() {
	fmt.Println("[+] number-verifier - SMS verification tool")
	fmt.Println()

	// parse args
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func main() {

	// List providers
	if *flagListProviders {
		fmt.Print("[+] Supported providers\n\n")
		for _, p := range providers.SupportedProviders {
			fmt.Printf("[-] %s\n", p.GetName())
		}
		return
	}

	// Select number provider first
	if *flagProviderFirst {
		providerFirst()
		return
	}

	countryFirst()

}

// getProvider selects a provider based on the input string.
func getProvider(selectedProvider string) providers.Providers {
	var provider providers.Providers

	for _, p := range providers.SupportedProviders {
		if selectedProvider == p.GetName() {
			provider = p.GetProviders()
		}
	}

	return provider
}

// providerFirst execute the provider first flow
func providerFirst() {
	var (
		providerStr, number string
		provider            providers.Providers
		err                 error
	)

	var providersList []string
	for _, p := range providers.SupportedProviders {
		providersList = append(providersList, p.GetName())
	}

	// Ask for provider
	err = survey.AskOne(&survey.Select{
		Message: "Choose a provider:",
		Options: providersList,
	}, &providerStr)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return
	}
	provider = getProvider(providerStr)

	// Ask for number
	numbers, _ := provider.GetNumbers()
	var numbersStr []string
	for _, n := range numbers {
		numbersStr = append(numbersStr, n.String())
	}

	err = survey.AskOne(&survey.Select{
		Message: "Choose a number:",
		Options: numbersStr,
	}, &number, survey.WithPageSize(len(numbers)))
	if err != nil {
		fmt.Println("error: " + err.Error())
		return
	}
	numberParsed := util.NumberParser(number)

	// Print out the messages
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	var messages []string
	fmt.Print("\n\n")
	for {
		messages, err = provider.GetMessages(numberParsed)
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

// countryFirst execute the country first flow
func countryFirst() {
	var (
		err             error
		number, country string
		provider        providers.Providers
	)

	fmt.Printf("[-] getting phone numbers...\n\n")

	countryNumbMap := make(map[string]map[string]providers.Number)
	for _, p := range providers.SupportedProviders {

		var numbersSlc []providers.Number
		if numbersSlc, err = p.GetNumbers(); err != nil {
			fmt.Printf("[.] error fetching numbers for %s: %s\n", p.GetName(), err.Error())
			continue
		}

		for _, n := range numbersSlc {
			fmt.Print(".")
			if _, ok := countryNumbMap[n.Country]; !ok {
				countryNumbMap[n.Country] = make(map[string]providers.Number)
			}
			countryNumbMap[n.Country][n.PhoneNumber] = n
		}
	}
	fmt.Println("\n[-] done")

	var countryList []string
	for k := range countryNumbMap {
		countryList = append(countryList, k)
	}

	err = survey.AskOne(&survey.Select{
		Message: "Choose a country:",
		Options: countryList,
	}, &country, survey.WithPageSize(len(countryList)))
	if err != nil {
		fmt.Println("error: " + err.Error())
		return
	}

	var numberList []string
	for _, v := range countryNumbMap[country] {
		numberList = append(numberList, fmt.Sprintf("%s %s", v.Provider, v.PhoneNumber))
	}

	err = survey.AskOne(&survey.Select{
		Message: "Choose a number:",
		Options: numberList,
	}, &number, survey.WithPageSize(len(numberList)))
	if err != nil {
		fmt.Println("error: " + err.Error())
		return
	}

	numberParsed := util.NumberParser(number)

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	var messages []string
	provider = getProvider(countryNumbMap[country][numberParsed].Provider)
	fmt.Print("\n\n")
	for {
		messages, err = provider.GetMessages(numberParsed)
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
