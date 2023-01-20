# Essential Go Rest API Example
## apptimedev/go-rest-api-template

[![Apptime](https://apptime.dev/icons/icon-128x128.png)](https://apptime.dev)

---

Application `apptimedev/go-rest-api-template` helps you to create a Rest API project.

---

* [Clone the project](#clone-the-project)
* [Execute](#execute)
* [Used packages](#used-packages)
* [License](#license)

---

## Clone the project

```sh
$ git clone https://github.com/apptimedev/go-rest-api-template
$ cd go-rest-api-template
```

---

## Execute

First copy the `dev/.env.example` to the `dev/.env`

Sync dependencies:

```sh
$ go mod -sync
```

Then run the application:

```sh
$ go run main.go
```

## Used packages

* [gorilla/mux](https://github.com/gorilla/mux): A powerful HTTP router and URL matcher.
* [joho/godotenv](https://github.com/joho/godotenv): Loads environment variables from `.env`.

---

## License

BSD licensed. See the LICENSE file for details.
