# ![logo](media/logo.svg) [![Godoc](https://godoc.org/github.com/qlova/seed?status.svg)](https://godoc.org/github.com/qlova/seed) [![Go Report Card](https://goreportcard.com/badge/github.com/qlova/seed)](https://goreportcard.com/report/github.com/qlova/seed) [![Build Status](https://travis-ci.org/qlova/seed.svg?branch=master)](https://travis-ci.org/qlova/seed)

The cross-platform Go framework for building apps.

## Usecases

*As a lightweight alternative to Electron*  
 Write your frontend and native code in Go, distribute native binaries of your app.
 Supported on Windows, Mac & Linux. Mobile support planned.
 
*Full-stack progressive webapp*  
 Write the complete app in Go, place binaries on public-facing web servers.
 Access these apps on Windows, Mac, Linux, IOS & Android.
 
*As a lightweight alternative to Phonegap* (WIP linux-only)  
 Write your app in Go, export the frontend as a native app.
 Android-only. IOS support planned.

![showcase](media/showcase.jpg)

## Installing

You can get Qlovaseed using go get.
```sh
go get -u -v github.com/qlova/seed/...
```
Qlovaseed apps can be developed on Android, Linux, Mac & Windows.

## Getting started

Create HelloWorld.go file and paste in the following contents:

```go
package main

import "github.com/qlova/seed"
import "github.com/qlova/seeds/text"
import "github.com/qlova/seeds/expander"

func main() {
	var App = seed.NewApp("Hello World")

	expander.AddTo(App)
	text.AddTo(App, "Hello World")
	expander.AddTo(App)

	App.Launch()
}

```

In the same folder, run 'go build' to create an executable for the app, run this to launch the app. By default, Qlovaseed will start a WebServer and open a browser window displaying your app.

## Widgets and Logic

Create a file called MyApp.go and paste in the following contents:

```go
package main

import "github.com/qlova/seed"

//Import a seed to use it, a list of seeds can be found [here](https://github.com/qlova/seeds).
import "github.com/qlova/seeds/button"

func main() {
	var App = seed.NewApp("My App")

	//In order to add a widget to your app, or container, use the package's AddTo method.
	var ClientPowered = button.AddTo(App, "My callback runs on the client")
	
		ClientPowered.OnClick(func(q seed.Script) {
			ClientPowered.Script(q).SetText(q.String("You clicked me!"))
		})
	
	
	var ServerPowered = button.AddTo(App, "My callback runs on the server")
	
		//You can style widgets with methods of the style package.
		ServerPowered.SetColor(seed.RGB(100, 100, 0))
	
		ServerPowered.OnClick(seed.Go(func(user seed.User) {
			ServerPowered.For(user).SetText("You clicked me!")
		}))

	App.Launch()
}
```

This example shows a quick glimpse on how powerful Qlovaseed is. You can find more widgets in the [seeds repository](https://github.com/qlova/seeds).

## Styles

All widgets/seeds can be styled with methods from the style package.
https://godoc.org/github.com/qlova/seed/style

```
var Text = text.AddTo(App, "Some syllable text)
Text.SetBold()
Text.Align(seed.Left)
Text.SetColor(seed.RGB(100, 0, 0)
Text.SetOuterSpacing(seed.Em, seed.Em)
```

## HTML/CSS/JS

Qlovaseed discourages the use of HTML, CSS and Javascript to build apps.
However, there may be good reasons to use these technologies to extend missing functionality. This is how:

* Seeds have a SetContent method for setting raw HTML.
* All seeds have a CSS method that returns a css.Style object with type-safe Set methods.
* When in doubt, seed.CSS().Set can be used to set css styles with strings,
* seed.Script has a Javascript method for raw Javascript.
* seed.Embed & seed.Seed.Require are useful for embedding Javascript and CSS files. Checkout the editor & swiper seeds.

## Themes

A work in progress material theme can be found [here](https://github.com/qlova/theme/tree/master/material).
Check the examples folder to learn how to use it.

**Please remember**, this framework is in development, it does not have a stable API and features are currently implemented as needed.
