# Web page details

This is a simple scraper to check information from webpages.

---

### TODOs:
- Add unit testcases
- Add gorounties for channels

### Solution details:
- Time spent : 2 hours
- Focused upon: readability, working solution, easy setup, manual testing

### Features

- Checks HTML version
- Checks page title
- Headings count by level
- Amount of internal and external Links
- Amount of inacessible links
- If a page contains a login form

---

### Get Started

##### Docker

If you have docker installed, clone this repository anywhere in your machine and in terminal, on this project folder, type the following comands:

```bash
 docker build -t webPageDetailer .
 docker run -it webPageDetailer
```

##### Golang


If you do not have docker installed but have golang in your machine, clone this repo inside your \$GOPATH work directory, them run the application in the terminal on this project folder:

```bash
go run .
```

---

### Extras

install docker: https://docs.docker.com/get-docker/ <br/>
install golang: https://golang.org/doc/install/
