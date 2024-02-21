# Introduction

`jtb-totp` is a simple TOTP program that stores keys and secrets in a keystore and generates OTPs for use in all the places you used to use Authy for.

## Supported Operating Systems
It should support all major operating systems as it was written to be OS-agnostic.

## Installation

### Binaries
Binaries for Linux, Windows, and macOS are available for download.  Once downloaded, just put the binary into your user/system path and enjoy.  For Linux, it would look like:

```console
$ mkdir -p ~/.local/bin
$ cp jtb-totp ~/.local/bin
$ chmod 755 ~/.local/bin/jtb-totp
```

In this example, ensure that `~/.local/bin` is in your PATH.

### Compilation
If you have `golang` installed, you can generate the binaries by cloning the repository and building the program.

```console
$ git clone https://github.com/JamesTheBard/jtb-totp
$ cd jtb-totp
$ go build jtb-totp.go
```

Again, ensure that the `jtb-totp` binary is in your user/system path so you can use it.

## Initial Setup

Once everything is installed, the keystore and initial configuration files need to be created.  Simply run:

```console
$ jtb-totp init
```

If you want to use a custom known password, you can set that up prior to the init command.  This also means that you can delete the config file so that the password is no longer on the file system.  However, you _will_ have to ensure that the `JTB_TOTP_SECRET` environment variable is set with the correct password.

```console
$ JTB_TOTP_SECRET=putpasswordhere jtb-totp init
```

## Commands

```
JTB-TOTP is a quick-and-dirty program that generates TOTP tokens and manages TOTP keys via the command-line.

Usage:
  jtb-totp [flags]
  jtb-totp [command]

Available Commands:
  add         Add key to keystore
  get         Generate TOTP code from key
  help        Help about any command
  import      Import keystore from YAML/JSON file
  init        Initialize keystore and settings
  list        List all names in keystore
  remove      Remove key from keystore

Flags:
  -h, --help   help for jtb-totp

Use "jtb-totp [command] --help" for more information about a command.
```

### `add` Command

```
Add key to keystore

Usage:
  jtb-totp add [flags]

Flags:
  -h, --help               help for add
  -n, --key-name string    name of the key (required)
  -v, --key-value string   value of the key (required)
```

To add a key to the keystore, you can use the `add` command.  For example, to add a key named `Google` with the secret of `JBSWY3DPEHPK3PXP`, you would:

```console
$ jtb-totp add -n Google -v JBSWY3DPEHPK3PXP
Updated keystore with new/changed data.
Added key 'Google' to keystore successfully!
```

### `get` Command

```console
Generate TOTP code from key

Usage:
  jtb-totp get [key name] [flags]

Flags:
  -h, --help   help for get
```

To generate an OTP, simply type the following command.  The `key name` argument uses fuzzy search, so it doesn't need to be exact.  It will return the best match in the keystore.

```console
$ jtb-totp get goog
Google -> 123456
```

### `import` Command

```console
Import keystore from YAML/JSON file

Usage:
  jtb-totp import [file to import] [flags]

Flags:
  -h, --help        help for import
  -o, --overwrite   overwrite the keys in current keystore
```

This command will import either a YAML or JSON file into your keystore.

For JSON files, the format needed is:
```json
[
    {
      "name": "test_key_1",
      "key": "ABCDEF"
    },
    {
      "name": "test_key_2",
      "key": "QWERTY"
    },
    {
      "name": "test_key_3",
      "key": "ZXCVBN"
    }
]
```

For YAML files, the format is very, very similar:
```yaml
- name: test_key_1
  key: ABCDEF
- name: test_key_2
  key: QWERTY
- name: test_key_3
  key: ZXCVBN
```

If provided in either format, the keys will be imported into the keystore.

There is one flag `-o/--overwrite`.  When set, this flag will overwrite keys currently in the datastore with the ones in the import file.  By default, keys that already exist in the datastore will not be overwritten.

### `init` Command

```console
Initialize keystore and settings

Usage:
  jtb-totp init [flags]

Flags:
  -h, --help   help for init
```

This command creates the initial encrypted keystore and configuration file, then outputs where they were created.  During the creation process, the program will generate a 32 character random password to encrypt the keystore with.  To override this, you can use the `JTB_TOTP_SECRET` environment variable.  Setting this prior to running the init will replace the generated password with the value of the environment variable.

```console
$ jtb-totp init
Initialized keystore and configuration files!
- Config file:   /home/jweatherly/.config/jtb-totp/jtb-totp.conf
- Keystore file: /home/jweatherly/.local/share/jtb-totp/keystore.enc
```

### `list` Command

```console
List all keys by name in keystore

Usage:
  jtb-totp list [flags]

Flags:
  -h, --help   help for list
```

This command lists all of the keys by name that exist in the keystore.

```console
$ jtb-totp list
Google
Test Key
Super Secret Key
bash.org
```

### `remove` Command

IN PROGRESS
