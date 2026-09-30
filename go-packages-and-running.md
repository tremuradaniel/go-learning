# Go: Packages and Running Code

Notes for `~/code/go/go-learning`. Covers the two things that trip people up early:
what a package actually *is*, and how to run a folder's worth of code.

---

## 1. The mental model: folder ≠ package

This is the single most important rule and the source of almost every early error.

- **A Go package is defined by the `package <name>` line at the top of the file — not by the folder name.**
- A folder *usually* holds one package, and by convention the folder is named after it. But the compiler does not care. `cards/deck.go` starting with `package main` is perfectly valid and compiles as a `main` package.
- **Every `.go` file in the same folder must declare the same package name.** Mix `package main` and `package deck` in one folder and the build fails immediately.
- Every folder in a module is its own package. Subfolders are *separate* packages, not part of the parent.

Two kinds of package matter at the start:

| Kind | Declaration | Result |
|---|---|---|
| Executable | `package main` + `func main()` | Builds/runs as a program |
| Library | any other name, e.g. `package deck` | Imported by others; `func main()` does nothing here |

> **Rule of thumb:** if you type `go run` on it, it needs `package main` and a `func main()`.

### The `newCard` incident, as a worked example

`cards/deck.go` and `cards/main.go` both start with `package main`, so they are **one package**. That means:

- `type deck []string` declared once in `deck.go` is visible in `main.go` — no import needed.
- `newCard()` declared in `deck.go` is callable from `main.go`.

So if you get `undefined: deck` or `undefined: newCard`, the declaration is **missing from the whole folder**, not "in the wrong file". One declaration, anywhere in the folder, satisfies it — and a second copy anywhere is a `redeclared in this block` error.

---

## 2. Creating a package / module folder

### Option A — start a fresh module

```bash
mkdir myapp && cd myapp
go mod init myapp          # creates go.mod; name is your module path
```

`go.mod` is the anchor: it marks the root of your module and everything under it belongs to this module.

```
module myapp    # module path — also the import prefix for subpackages
go 1.22.2       # the Go version you're targeting
```

The module name (`cards`, `myapp`) is **just a path label**. It has zero effect on the `package` clause inside your files. `module cards` + `package main` is normal and correct.

### Option B — add a package to an existing module

Make a subfolder, give its files a shared package name:

```bash
cd go-learning
mkdir deck                # the package lives here
```

`deck/deck.go`:

```go
package deck              // NOT main — this is a library package

type Deck []string        // exported: capital first letter = visible outside the package

func New() Deck {         // exported constructor (Go has no constructors, so New() is the idiom)
    return Deck{}
}

func (d Deck) Print() {   // exported method
    // ...
}
```

`main.go` at the module root imports it by **module path + subfolder**:

```go
package main

import (
    "fmt"
    "cards/deck"        // module-name/subfolder — the folder path, not the package name
)

func main() {
    d := deck.New()
    d.Print()
    fmt.Println(d)
}
```

**Import path = `modulepath/subfolder`.** If your `go.mod` says `module cards` and the folder is `deck/`, the import string is `"cards/deck"`, even though the package clause inside is `package deck`.

### Exported vs unexported — the capital-letter rule

- `New`, `Deck`, `Print` → **exported**, usable from other packages.
- `new`, `deck`, `print` → **unexported**, private to the declaring package.

This is why `deck` stays lowercase inside your `cards/` folder — everything is in one package there, so it can stay private. Move it to its own package and you'd rename the type to `Deck` and add a `New()` constructor.

---

## 3. Running "everything in the folder"

`go run .` is the one you want. The `.` means **"the package in the current directory"** — no filename required.

```bash
cd ~/code/go/go-learning/cards
go run .                   # compile + run the current folder's package
```

### The full command set

| Command | What it does |
|---|---|
| `go run .` | Compile and run the current folder's `main` package |
| `go run main.go` | Run a single file (only that file's build) |
| `go run main.go deck.go` | Run an explicit set of files, ignoring the rest of the folder |
| `go build .` | Compile to a binary in the current dir |
| `go build -o cards .` | Compile to a named binary (`cards`) |
| `go build ./...` | Build **every package** in the module subtree |
| `go test ./...` | Run **all tests** in the module subtree |
| `go vet ./...` | Static checks across all packages |
| `go fmt ./...` | Format every file in the subtree |
| `go install` | Build and put the binary in `$GOPATH/bin` |

**The `./...` pattern is how you run everything below you.** It recursively matches every package from the current folder downward — the idiomatic way to build, test, and vet a whole project at once.

### Personal preference: prefer `go run .` over naming files

`go run main.go` only compiles `main.go`. If `main.go` calls `newCard()` which lives in `deck.go`, that fails with `undefined: newCard` — because `deck.go` was never part of the build. `go run .` pulls in **the whole folder**, which is what you actually want 99% of the time.

Naming files explicitly is the escape hatch, not the default.

---

## 4. Gotchas that cost real time

### `go run .` compiles what's on **disk**, not your editor

The compiler never sees an unsaved buffer. Classic symptom: your editor shows the complete file, but `go run .` throws `undefined` errors against code you can see on screen.

**Diagnose it in one line** — compare the byte count to what you expect:

```bash
wc -c deck.go            # a full file is ~1300 bytes; 467 means it's stale
cat -n deck.go | head -20   # confirm the type declaration is actually there
pwd                      # confirm you're in the folder you think you are
```

Fix: save the file (`Ctrl+S` / `:w` / `Ctrl+O`), then re-run. If it still doesn't take, open the file by **absolute path** so you know the editor and terminal point at the same file — it's easy to have two `cards` folders open across the 8 subdirs under `go-learning`.

### `no Go files in ...`

You ran `go run .` in a folder with no `.go` files, or in a parent folder that only contains subfolders. Run it in the folder that actually holds the code.

### `package main is not a main package` / `function main is undeclared`

A `main` package **must** have exactly one `func main()` with no arguments and no return values.

### `imported and not used` / `declared and not used`

Go refuses to compile with unused imports or unused local variables. This is a feature, not a bug — it keeps code clean. Delete them, or use `_ = x` if you genuinely must keep one.

### `undefined: X` where X is clearly defined

Almost always one of:
1. The file isn't saved.
2. You ran `go run main.go` (single file) instead of `go run .` (whole folder).
3. The declaration has the wrong case — `deck` vs `Deck`, `print` vs `Print`. Cross-package access requires the capital.

### One folder, one package

A single `package` clause per folder. If you need a second package, it needs a second folder.

---

## 5. Cheat sheet

```bash
go mod init myapp      # start a module
mkdir sub && echo ...  # add a package folder (give its files a shared package name)

go run .               # run the current folder's package
go run main.go deck.go # run explicit files
go build ./...         # build every package below
go test ./...          # test every package below
go vet ./...           # vet every package below
go fmt ./...           # format everything below
```

Remember:
- **Package = the `package` line, not the folder name.**
- **One `package` clause per folder; every file must agree.**
- **Import path = `modulepath/folder`.**
- **Capital first letter = exported.**
- **`go run .` reads disk — save before you run.**
- **`./...` means every package from here down.**

