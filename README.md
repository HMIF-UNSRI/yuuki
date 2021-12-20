# Yuuki

A backend application on the [HMIF UNSRI](hmifunsri.org) website.

## Getting Started

1. Start with cloning this repo on your local machine :

```
$ git clone git@github.com:HMIF-UNSRI/yuuki.git
$ cd yuuki
```

2. Install MySQL, then create database named `yuuki`

3. Run database migration using `migrate` library, read
   documentation [here](https://github.com/golang-migrate/migrate) :

```
$ migrate -database "mysql://root:@tcp(localhost:3306)/yuuki?parseTime=true" -path migrations up
```

You can customize your database url with your own `.env` files.

4. Start server

```
$ go run main.go
```

Connect to https://localhost:8080/api.

## Contributors

<table>
    <tr>
        <td align="center"><a href="https://www.linkedin.com/in/arya-yunanta-255424174/"><img src="https://avatars.githubusercontent.com/u/77351340?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Arya Yunanta</b></sub></a></td>
    </tr>
</table>