# Templ Tailwind CSS Example

An example of how to use the tailwindcss standalone CLI to generate CSS for template files.

## Tasks

### install-deps

```
go install github.com/a-h/templ/cmd/templ@main
npm install -g tailwindcss
```

### generate-css

```
tailwindcss -o css/styles.css --minify
```

### generate-templates

```
templ generate
```

### run

Requires: generate-css, generate-templates

```
go run .
```
