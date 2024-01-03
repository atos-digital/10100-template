# 10.10.0 App

## Getting started

### Running the app

You must be on MacOS/Linux in order to use Make. Windows users can install [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) to use Make.

### Quick start

Install:

- [go](https://go.dev/doc/install) *Note: ensure you have the gobin on your path*
- [node](https://nodejs.org/en/download)

*N.B.* These are just the official install pages, you can use your package manager of choice.

Once installed, we need a few tools to be available on the command line.

Run the following command:

We use [Tailwind](https://tailwindcss.com/) for styling, and [Prettier](https://prettier.io/) for formatting.

```bash
npm install -g prettier tailwindcss
```

We use [templ](https://templ.guide/) to generate the HTML pages from Go.

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

For building and auto rebuild we user [air](https://github.com/cosmtrek/air)

```bash
go install github.com/cosmtrek/air@latest
```

To simply run the app, run the following command from the root of the project:

```bash
make
```

**Note:** You may need to clone the `.env.example` file to `.env` first, although the make command will do this for you.

## How this works

The basic premise is that we use HTMX to make the app interactive. This means that we can use Go to render the initial page, and then use HTMX to make the app interactive.

Golang serves the pages by using templ to generate pages and serving them via standard handlers. Tailwind is used to style the components made in templ.

Air is simply used to reload on change and run the tailwind and templ commands on change.

The key command that does all of this is in the `.air.toml` file:

```toml
cmd = "npx tailwindcss -i ./ui/styles.css -o ./internal/server/assets/css/styles.css && templ generate && go build -o ./tmp/main cmd/server/main.go"
```

Here we can see that we are running the tailwind command, then the templ command, and then building the app.
