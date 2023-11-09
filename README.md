## HTMX Trie

- Features
  - Golang
    - Air
    - Fiber HTTP
    - Go html templates
  - HTMX
  - Tailwind

### How to run

```bash
# clone project
$ git clone github.com/stneto1/htmx-trie

# install required packages
$ go mod install

# to run without reload
$ go run .

# to run with hot reload
$ air

# watch css changes
$ (npx|bunx) tailwindcss -i pkg/assets/input.css -o public/output.css --watc

$ open http://localhost:3000

```
