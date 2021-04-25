# gin-boilerplate

## Structure

```shell
.
├── README.md
├── applicatoins
│   ├── index
│   │   ├── controllers.go
│   │   └── routes.go
│   ├── post
│   │   ├── controllers.go
│   │   └── routes.go
│   ├── schedule
│   │   ├── controllers.go
│   │   └── routes.go
│   └── user
│       ├── controllers.go
│       └── routes.go
├── assets
│   ├── cozyfex-logo.png
│   └── sub
│       └── avatar-background.png
├── config
│   └── config-directory
├── go.mod
├── go.sum
├── libs
│   ├── loadtemplates
│   │   └── loadtemplates.go
│   ├── mysqlstore
│   │   └── mysql.go
│   └── utils
│       └── utils.go
├── main.go
├── models
│   └── models.go
└── templates
    ├── includes
    │   ├── head.gohtml
    │   └── scripots.gohtml
    ├── index.gohtml
    ├── layouts
    │   ├── bottom.gohtml
    │   └── top.gohtml
    ├── post
    │   └── page.gohtml
    ├── schedule
    │   └── index.gohtml
    └── user
        └── index.gohtml
```

***applications***\
This is for your business logic and url mapping.\
The subdirectory like `post` is a name of your app.\
`controllers.go` is for your business logic.\
`routes.go` is for mapping url.

***assets***\
This is for your static files.

***config***\
This is just ready for your settings.

***libs***\
This is for custom packages in this project.

***models***\
This is for models for database.\
You can define a model for your service.\
In this boilerplate, I use [`GORM`](https://gorm.io) ORM.

***templates***\
This is for your template files.

## Installed Packages

```go
github.com/gin-contrib/sessions
github.com/gin-gonic/gin
github.com/go -sql-driver/mysql
github.com/gorilla/sessions
github.com/joho/godotenv
github.com/srinathgs/mysqlstore
gorm.io/driver/mysql
gorm.io/driver/sqlite
gorm.io/gorm
```

## Custom Packages in libs

***loadtemplates***\
Getting list of files in `templates` directory.

***mysqlstore***\
The package is for mixing with `github.com/gin-contrib/sessions` and `github.com/srinathgs/mysqlstore` packages for
saving session to MySQL.\
If you want to use another store for session,
visit [`github.com/gin-contrib/sessions`](https://github.com/gin-contrib/sessions)

***utils***\
Simple functions 
