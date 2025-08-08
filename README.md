# passgen-go
Diceware/XKCD-style password generator

No wordlists are included with this program. Provide a file or URL with a wordlist.

`passgen-go` automatically parses regular and Diceware-style wordlists.

# Usage



```sh
passgen-go https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt
passgen-go https://raw.githubusercontent.com/sts10/orchard-street-wordlists/refs/heads/main/lists/orchard-street-medium.txt

curl -o eff.txt https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt
passgen-go eff.txt
```

# Wordlist Resources
[EFF Large Wordlist](https://www.eff.org/document/passphrase-wordlists)
[sts10/orchard-street-wordlists](https://github.com/sts10/orchard-street-wordlists/tree/main)
