# **Spark**

Spark is a static-site generator that's intuitive and designed for simplicity. It allows you to put your focus on creating quality web content, rather than worrying about the complexities of web development.

----

## What can Spark do?
- Generate HTML files from Markdown files, making static-site creation simple.
- Write assets for your files, and they will get copied and applied to your converted files as well.
- Serve your generated files on a web server, with intuitive routing

----

## Installation
You can install Spark to your Go bin location by running:
```bash
$ go install gitlab.com/EndowTheGreat/spark/cmd/spark@latest
```
Or you can build from source:
```bash
$ git clone https://gitlab.com/EndowTheGreat/spark.git
$ cd spark
$ make && cd bin
$ ./spark
```

If you are using an Arch Linux based distribution, you can also install Spark directly from the AUR:
```bash
$ yay -S spark-git
```

## Converting Your Files
To convert files, it's recommended to create an input directory to house your Markdown files. However you structure this directory is how your routes will be structured if you choose to serve them. For example, if you have this directory tree:
```bash
├── markdown
    ├── assets
    │   ├── index.css
    │   └── login.js
    ├── account
    │   └── login.md
    └── index.md
```
Then the route setup would look like this:
```bash
/index # This is also assumed to be the home file, so it can be accessed simply at '/' as well
/account/login
```
You can choose whether or not to use the ".html" file extension when accessing routes with the server.

Alright, now onto the actual conversion. This can be done using the CLI's convert command:
```bash
$ spark convert -i markdown -o web
```
If you do not provide arguments, the input directory is assumed to be the current directory, and the output directory will be "output"

## Serving Converted Files
Now that you have some converted files, you can choose to serve them using a web server that Spark will spin up for you. It is more so built for development purposes than for actual production use, but it could definitely work for a smaller project that doesn't need too many bells and whistles.
```bash
$ spark serve -p 8080 --dir web
```
This will start an HTTP server on port 8080, and serve the files you just converted in the web directory.

## Contributing
Pull requests and contributions are absolutely welcome, feel free to fork or improve upon my work however you wish. To make things nice and easy, you can open a PR [here](https://gitlab.com/EndowTheGreat/spark/-/merge_requests/new).
