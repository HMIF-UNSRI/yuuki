# Yuuki

A backend application on the [HMIF UNSRI](hmifunsri.org) website.

## Getting Started

1. Start with cloning this repo on your local machine :

```
$ git clone git@github.com:HMIF-UNSRI/yuuki.git
$ cd yuuki
```

2. Install MySQL, then create database named `yuuki`, and create table by copy-pasting the following code.

```mysql
CREATE TABLE IF NOT EXISTS categories
(
    id         INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    slug       VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

3. Start server

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