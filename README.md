# GO4E

Go4E stands for Go for Education.


# Codebase information
- Uses Go 1.17.2
- Each folder in the base dir is a `module` and so we have multiple modules in this repo.
- `go.mod` is a file in each `module` which tells us about the module dependencies
- `go.sum` is _not_ like `package-lock.json` as in Node.js projects; it only contains the cryptographic hashes of the dependency
- In Go, `module` & `package` are different things.

# Running instructions

- Clone the repository: `https://github.com/shhivam/go4e.git`
- `cd` into the desired module which you want to build
- Make sure we have a `main` package in the module
- Run `go build` 