# Imagephrase

A CLI tool that allows you to hide your [BIP-39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) mnemonic in an image.

I worked on this project while attending the [ETHBerlin 2024](https://ethberlin.org) hackathon. Given that this is a hack, you shouldn't use the tool to hide your real mnemonic.

## How it works

An image is a 2-dimensional array of pixels. Every pixel has three color values (Red, Green and Blue). Each color is represented as 1 byte (8 bit), which means that it can have a value ranging from 0 to 255. "Mixing" the values of those three bytes together (one for Red, one for Green and one for Blue) results in the desired color for the pixel.

A minor change to the color's byte value results in a similar color. It's therefore possible to hide binary data in an image by setting the [least significant bit](https://en.wikipedia.org/wiki/Bit_numbering) of every color byte to the bits of the binary data to be hidden.

Given that a pixel has 3 color values, we can hide 3 bits of information per pixel.

This technique is called [Steganography](https://en.wikipedia.org/wiki/Steganography). It's important to note that **the information isn't encrypted** or otherwise protected. It's "hidden in plain sight" within the image.

In an upcoming version of this tool I plan to add support for a combination of encrypting the mnemonic and using [Shamir's Secret Sharing](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing) to hide and share the encryption key alongside the encrypted mnemonic via multiple images. This way, one would need $t$ out of $n$ images to reconstruct the encryption key and therefore recover the mnemonic.

## Setup

1. `git clone <url>`
2. `cd <name>`
3. `cat mnemonic.txt`
4. `go run ./cmd/imgp hide -image-path ./photo.png -mnemonic "$(cat mnemonic.txt)"`
5. `go run ./cmd/imgp reveal -image-path ./photo.modified.png`

## Useful Commands

```sh
go run <package-path>

go build [<package-path>]

go test [<package-path>][/...] [-v] [-parellel <number>]
go test -bench=. [<package-path>] [-count <number>] [-benchmem] [-benchtime 2s] [-memprofile <name>]
go test -cover [<package-path>]
go test -race [<package-path>] [-count <number>]
go test -shuffle on [<package-path>]
go test -cover
go test -coverprofile <name> [<package-path>]
go test -run FuncName/RunName

go tool cover -html <name>
go tool cover -func <name>

go tool pprof -list <name> <profile-name>

go fmt [<package-path>]

go vet [<package-path>]

go clean [<package-path>]

go help <command-name>

go mod init [<module-path>]
go mod tidy
go mod vendor
go mod download

go work init [<module-path-1> [<module-path-2>] [...]]
go work use [<module-path-1> [<module-path-2>] [...]]
go work sync

# Adjust dependencies in `go.mod`.
go get <package-path>[@<version>]

# Build and install commands.
go install <package-path>[@<version>]

go list -m [all]
```

## Useful Resources

- [Go - Learn](https://go.dev/learn)
- [Go - Documentation](https://go.dev/doc)
- [Go - A Tour of Go](https://go.dev/tour)
- [Go - Effective Go](https://go.dev/doc/effective_go)
- [Go - Playground](https://go.dev/play)
- [Go by Example](https://gobyexample.com)
