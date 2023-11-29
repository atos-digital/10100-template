# 10.10.0 App

## Development

### Getting started

#### Running the app

To simply run the app, run the following command from the root of the project:

```bash
make
```

**Note:** You may need to clone the `.env.exmaple` file to `.env` first, although the make command will do this for you.

**Note:** If you get an errors, see the [Prerequisites](#prerequisites) section below.

### Prerequisites

For a better developer experiance there are tools setup and configureed to allow automatic reload and regeneration of styles.

It is highly recommended to install all tools and use VSCode as the IDE. There are settings and tasks configured to make development easier.

First install `node` and a few `global` deps run:

```bash
npm install -g prettier tailwindcss
```

This will setup tailwind and prettier formatters.

#### Auto reload

[air](https://github.com/cosmtrek/air) is used to build and reload the server.

The included [air.toml](air.toml) file is configured to watch for changes. To use it, install air with:

```bash
go install github.com/cosmtrek/air@latest
```

Then run the following command from the root of the project:

```bash
make
```
