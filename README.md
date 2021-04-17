# gin-boilerplate

## Directory Structure
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
├── main.go
└── templates
    ├── index
    │   └── index.tmpl
    ├── partials
    │   ├── head.tmpl
    │   ├── menu.tmpl
    │   └── scripts.tmpl
    ├── post
    │   └── page.tmpl
    ├── schedule
    │   └── index.tmpl
    └── user
        └── index.tmpl

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

***templates***\
This is for your template files.\
You must keep directory structure for this boilerplate.\
If you want to change the structure, you have to change below code.\
main.go
{: .small}
```go
	r.LoadHTMLGlob("templates/**/*")
```

