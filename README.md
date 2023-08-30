A little test app for htmx and echo framework using go

templates are inside `renderer/` which has all of the html assets and templates

assets are embedded into the go binary

# Setup

```
go install github.com/a-h/templ/cmd/templ@latest # for code generation
cd renderer/npm
npm install
```

# Runing the server

```
task run-server
```
