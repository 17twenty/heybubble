# Bubble Demo

A simple demo using [HTMX](https://htmx.org/docs/), [Gorilla](https://github.com/gorilla/mux) and [Slog](https://golang.org/x/exp/slog).

## Getting Started

```bash
go build && ./heybubble
```

I use the static tailwind binary and store it in `tools` but it's in .gitignore as it's sizey.
You can grab this and add it to the path with the following:

```bash
$ cd tools
$ curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64
$ chmod +x tailwindcss-macos-x64
$ mv tailwindcss-macos-x64 tailwindcss
$ PATH=$PATH:`pwd`/tools
...
```

Note: _You will likely want the right binary for your OS/arch_

## Building Tailwind Templates

With the following you can go fast and generate CSS output CSS to play with. The `templates/static` folder is ignored.

```bash
$ cd templates
$ tailwindcss -i ./static/css/tailwind.css -o ./static/css/main.css --watch

Rebuilding...

Done in 264ms.
```

When happy with templates, compile and minify the CSS for production
`tailwindcss -i ./static/css/tailwind.css -o ../static/css/main.css --minify`


```bash
$ go run *.go 
time=2023-05-30T11:16:15.487+10:00 level=INFO source=/Users/nickglynn/Projects/htmx-demo/server.go:41 msg="Starting server..." SERVER=http://localhost:8080
```

## TODO

- Blackfriday for markdown
- So many things
