# SMS Number Verifier - Simple, free SMS verification tool

Number Verifier is a **SMS verification tool** that makes it easy to get a disposable SMS number and bypass SMS number verifications on any site, for example Facebook, Discord, Twitter, Uber, WeChat, Google & more.

This is a fork of the original project developed by [upmasked](upmasked/number-verifier) featuring additional SMS provider support and command line options

- Written in Golang
- Support for multiple free SMS providers (**PRs are open!**)
- Automatically refreshing of latest messages (5s delay)
- Easy switching between providers

## Usage
To use this tool first download the [latest release](https://github.com/jmtsantos/number-verifier/releases/latest). If you're using Windows/Mac (darwin)/Linux, and not sure which one to select, you'll probably want to choose the **'amd64'** (=64-bit) variant. After that, extract the package and run it.

## Providers
We currently support the following providers:

- [sms-online.co](https://sms-online.co)
- [receive-smss.com](https://receive-smss.com)
- [receive-sms-online/](https://receive-sms-online.info/)
- [smsreceivefree](https://smsreceivefree.com/)
- [Upmasked](https://upmasked.com/temporary-phone-number/fake-sms)

## TODO
- [ ] Support more free providers
- [ ] Add documentation on how to write new provider support modules
- [ ] Allow easy switching between numbers
- [ ] Make amount of messsages shown a parameter
- [ ] Improve message list output

## Disclaimer
Using this software to violate the terms and conditions of any third-party service is strictly against the intent of this software. By using this software, you are acknowledging this fact and absolving the author or any potential liability or wrongdoing it may cause. This software is meant for testing and experimental purposes only, so please act responsibly.

## License
MIT &copy; [Upmasked](https://upmasked.com)
