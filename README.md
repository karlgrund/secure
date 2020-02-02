# secure
[![][languagego img]][languagego]

## Overview

`secure` is a CLI tool that enables sending files securly over an insecure connection (i.e. Slack). It's using public ssh-key for encrypting files and allows decryption using private key.

## Usage

To get help at any time for the tool or any command you can run:

```
secret --help            # help pertaining to the tool itself
secret encrypt --help    # help pertaining to the encrypt command
secret decrypt --help    # help pertaining to the decrypt command
```

### Encrypt

`secure encrypt --file my_secret_file.txt --publicKey recipient_key.pub`

Files to send to recipient.

| Filename           | Purpose                          |
| ------------------ | -------------------------------- |
| secret.txt.enc     | Encrypted file containing secret |
| secret.key.enc     | Encrypted symmetric key          |

### Decrypt

`secure decrypt --file secret.txt.enc --secretKey secret.key.enc`

| Filename           | Purpose                           |
| ------------------ | --------------------------------- |
| secret.txt         | Unencrypted file from sender      |


## Technical overview

1. Generating 32 bit symmetric key.
2. Encrypt file to transfer with the symmetric key.
3. Encrypt the symmetric key using recipients public ssh-key.

[languagego]:https://golang.org
[languagego img]:https://img.shields.io/badge/language-golang-77CDDD.svg?style=flat
