<div align="center">
	<img src="assets/images/gooey-gopher.png" width="200" height="200">
	<h1>gooey</h1>
	<p>
		<b>Declarative UI framework for Go</b>
	</p>
	<br>
	<br>
	<br>
</div>

#### 🚨 This project is in its beginning and any APIs are subject to change without a notice 🚨
> If you have some ideas, suggestions or generally would like to help, please submit a GitHub Issue and/or PR!

## 🌈 Features

* **Declarative syntax** - think [SwiftUI](https://developer.apple.com/xcode/swiftui/) or [React](https://reactjs.org/) but in Go
* **Component-based** - split your UI into small, fully encapsulated components
* **Powered by GTK3** - leverage a established, stable cross-platform widget toolkit that should work almost everywhere Go can be compiled on

## 📙 Usage

### Basic example

```go
// You need to init GTK before anything else
gtk.Init(nil)

// Declare the window
window := Window{
	Title: "Hello go-gooey!",
	// Hook the destroy signal to a callback
	Destroy: func() {
		gtk.MainQuit()
	},
	// Set the default size, otherwise the window will be as big as the initial contents
	DefaultSize: &Size{Width: 400, Height: 200},
	// Describe the window children components
	Children: []Widgetable{
		&Label{
			Text: "👋",
		},
	},
}

// Try opening the window
err := window.Open()
if err != nil {
	log.Fatal("Unable to open window:", err)
}

// Start the GTK main loop
gtk.Main()

```

It should look like this:

![](assets/images/example-simple.png)

## 🔍 GTK Inspector

To make the debugging easier, you can use the [GTK Inspector](https://blog.gtk.org/2017/04/05/the-gtk-inspector/) to peek inside of the running app.

## Attributions
* Logo composed from icons by [Hugo Arganda](http://about.me/argandas), Alex Muravev and Ben Davis.