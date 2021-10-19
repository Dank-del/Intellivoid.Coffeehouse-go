<!--
	This file is part of Intellivoid.Coffeehouse-go (https://github.com/intellivoid/Intellivoid.Coffeehouse-go).
	Copyright (c) 2021 Sayan Biswas, ALiwoto.

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, version 3.

	This program is distributed in the hope that it will be useful, but
	WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
	General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program. If not, see <http://www.gnu.org/licenses/>.

	
-->

# <h1 align="middle"> <img src="https://coffeehouse.intellivoid.net/assets/favicon/favicon-194x194.png" width="35px" align="left"></img> Coffeehouse-go <img src="https://intellivoid.net/assets/favicon/android-chrome-192x192.png" width="35px" align="right"></img> </h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/intellivoid/Intellivoid.Coffeehouse-go.svg)](https://pkg.go.dev/github.com/intellivoid/Intellivoid.Coffeehouse-go)

Coffeehouse-go is a small, simple and official [golang](https://go.dev) library for
using [Coffeehouse](coffeehouse.intellivoid.net) API services. You can find JavaScript version of this
library [here](https://github.com/intellivoid/CoffeeHouse-JavaScript-API-Wrapper). It's designed to support all features
of Coffeehouse in one piece, so by using this library, you don't have to install anything else. It has simple and
special tools to work with complicated parts of the API, like Generalization.

<hr/>

### Table of contents:
 * [What is Coffeehouse?](#What-is-Coffeehouse)
 * [Library dependencies](#library-dependencies)
 * [Features](#features)
 * [Getting started](#getting-started)
 * [How to use](#how-to-use)
 * [Support and Contributions](#support-and-contributions)
 * [Links](#links)
 * [License](#license)


 <img align="middle" src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/798246901916499998.gif" width="350px" height="4px" ></img>

<hr/>

## What is Coffeehouse?
[CoffeeHouse](https://coffeehouse.intellivoid.net) is one of [Intellivoid](intellivoid.net)'s productiona. It's a one stop solution for cloud-based artificial intelligence and machine learning processing, overtime more features are added and improved on CoffeeHouse.

### Priorities of Coffeehouse
- Community and Official Support: It has a great community where you can drop by in and ask any question your heart desires.

- Open Platform: It has open source API Wrappers and its Documentations are available to all and it's free to contribute!

- Affordable and Free services: Coffeehouse provides its services for free if you would like to use it for personal uses. Need more? It's also affordable with monthly subscriptions!

- Independent Technologies: CoffeeHouse is not dependent upon third party services or libraries such as Tensorflow to function, everything was build from scratch.

- Always More: We are constantly working hard to add and improve these services for everyone!

- Truly Simple: No need for complicated configurations or setups, start using our services with little to no effort as a software developer.

<hr/>

## Library dependencies
This library doesn't depend on any other third party libraries.
It's written in pure Go and there is no need to install any kind of plugin.
It needs at least Go 1.16.

<hr/>

## Features

- Uses official Coffeehouse APIs, which makes this library:
   - Easy to update
   - Guaranteed to match the docs
   - No third party endpoints
   - Self-documenting (Contains all pre-existing Intellivoid's docs)
- Additional tools for using all Coffeehouse features, like Generalization and Media Proccessing.
- Serialization tools, which make you able to easily convert data and structures to Json formats and if you would like you can save them to your local storage as files with only one method.
- No third party library bloat; only uses standard library.
- Type safe; using separated structs for separated APIs, and at the same time, using only one type for Generalization and errors.
- No panicking; you don't need to be worry about your program bein stopped by panicking of this library. It uses standard way of error handling in Golang. For more information, read [this file](CODESTYLE.md).

<hr/>

## Getting started
Simply use only one command to download the library in less than ten seconds in your local storage:
It's the standard `go get` command which is available in all platforms:

```bash
go get github.com/intellivoid/Intellivoid.Coffeehouse-go
```
or
```bash
dep ensure -add github.com/intellivoid/Intellivoid.Coffeehouse-go
```

Also if you need full documentation of this library, you can find
it [here](https://pkg.go.dev/github.com/intellivoid/Intellivoid.Coffeehouse-go).

<hr/>

## How to use?

Every package has its own ReadMe file, please read them carefully in order to be able to use the library.
You can see a list of preview here:
- [Coffeehouse](coffeehouse/README.md)
   - [Generalization](coffeehouse/generalization/README.md)
   - [MediaProcessing](coffeehouse/mediaProcessing/README.md)
      - [NSFW Classification](coffeehouse/classificationNSFW/README.md)
   - [NLP (Natural Language Processing)](coffeehouse/nlp/README.md)
       - [Emotion Analysis](coffeehouse/emotionAnalysis/README.md)
       - [language Detection](coffeehouse/langDetect/README.md)
       - [Part of Speech Tagging](coffeehouse/posTagging/README.md)
       - [Sentence Split](coffeehouse/sentenceSplit/README.md)
       - [Sentiment Analysis](coffeehouse/sentimentAnalysis/README.md)
       - [Spam Prediction](coffeehouse/spamPrediction/README.md)

<hr/>

## Support and Contributions
 * Contributions are welcome! If you are interested in contributing fixes or features to Coffeehouse-go, please read our [contributors guide](CONTRIBUTING.md) first.

 * Join [Intellivoid updates channel](https://t.me/Intellivoid), if you want to be aware of most recent changes!

 * Have a problem with API? Servers are down? Something went wrong? Ensure that you are joined at our [Support chat](https://t.me/IntellivoidDiscussions)!

 * Having a problem with library? Wanna talk with repository's owner? Contact the Maintainers:
    - [Sayan Biswas](https://t.me/dank_as_fuck)
    - [ALiwoto](https://t.me/Falling_inside_the_Black)

 * Still don't know what's going on? Not sure about how API works? Be sure to
   read [Introduction](https://docs.intellivoid.net/coffeehouse/v1/introduction).

 * If you think you have found a bug or have a feature request, feel free to use
   our [issue tracker](https://github.com/intellivoid/Intellivoid.Coffeehouse-go/issues). Before opening a new issue,
   please search to see if your problem has already been reported or not. Try to be as detailed as possible in your
   issue reports.

 * If you need help using Intellivoid's APIs or have other questions we suggest you to join
   our [telegram community](https://t.me/IntellivoidCommunity). Please do not use the GitHub issue tracker for personal
   support requests.

<hr/>

## Links

 * [Official website](https://intellivoid.net)
 * [Coffeehouse](https://coffeehouse.intellivoid.net)
 * [Official channel](https://t.me/Intellivoid)
 * [Support chat](https://t.me/IntellivoidDiscussions)
 * [Intellivoid Community](https://t.me/IntellivoidCommunity)
 * [Discord server](https://discord.gg/euNkxEKPJb)
 * [API Introduction](https://docs.intellivoid.net/spamprotection/introduction)
 * [API documents](https://docs.intellivoid.net/spamprotection/v1/lookup)
 * [SpamProtection bot](https://t.me/SpamProtectionBot)

<hr/>

## License

<img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/Something_that_looks_like_Diamond.png" width="25px"></img> The Intellivoid.SpamProtection-go project is under the [GPL-3.0](https://opensource.org/licenses/GPL-3.0) license. You can find the license file [here](LICENSE).


<img src="https://intellivoid.net/assets/media/TextLogo2.svg" width="195px" style="background-color:#a0b1c1">
