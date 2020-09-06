# Barebones Viginere

See: [https://en.wikipedia.org/wiki/Vigenère_cipher](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher)

## Usage

`go run viginere.go <passphrase> <message>`

## Example

```
$ go run viginere.go lemon "attack at dawn"
Key: lemon
Phrase: attack at dawn
Ciphertext: LXFOPVEFRNHR
Plaintext: ATTACKATDAWN
```
