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

You can also use the `Makefile` provided to build all of the relevant binaries under Linux.  The builds will be put into the `builds` directory.

```console
$ make clean && make releases
```

Again, ensure that the `jtb-totp` binary is in your user/system path so you can use it.

## Initial Setup

Once everything is installed, the keystore and initial configuration files need to be created.  Simply run:

```console
$ jtb-totp init --initialize
```

If you want to use a custom known password or put the keystore in a location of your choosing, you can do that as well.

```console
$ jtb-totp init --initialize \
> --keystore /path/to/keystore.enc \
> --password thisisacoolpassword
```

## Commands

```
JTB-TOTP is a quick-and-dirty program that generates TOTP tokens and manages TOTP keys via the command-line.

Usage:
  jtb-totp [flags]
  jtb-totp [command]

Available Commands:
  add         Add key to keystore
  export      Export keystore to YAML file or standard out
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
  jtb-totp add [key name] [key secret] [flags]

Flags:
  -h, --help               help for add
```

To add a key to the keystore, you can use the `add` command.  For example, to add a key named `Google` with the secret of `JBSWY3DPEHPK3PXP`, you would:

```console
$ jtb-totp add Google JBSWY3DPEHPK3PXP
Updated keystore with new/changed data.
Added key 'Google' to keystore successfully!
```

For key names with spaces, you can enclose it in quotes.

```console
$ jtb-totp add "My Little Pony" JBSWY3DPEHPK3PXP
Updated keystore with new/changed data.
Added key 'My Little Pony' to keystore successfully!
```

### `export` Command

```
Export keystore to YAML file or standard out

Usage:
  jtb-totp export [file to save export to] [flags]

Flags:
  -h, --help   help for export
```

This command will export the keystore into a YAML file in a convenient format that will also allow you to re-import it later if need-be.  Since this is a pretty dangerous command, it is _required_ that you supply the password via environment variable (`JTB_TOTP_SECRET`).  If a filename is supplied, the contents of the keystore will be saved to the filename.  If no arguments are supplied, the contents of the keystore will be written to standard out (`STDOUT`).

```console
$ JTB_TOTP_SECRET=mypasswordhere jtb-totp export
- name: Google
  key: ABCDEF
- name: Ars Technica
  key: QWERTY
- name: Asta La Vista
  key: ZXCVBN
```

```console
$ JTB_TOTP_SECRET=mypasswordhere jtb-totp export test.yaml
Keystore exported to 'test.yaml'.
```

### `get` Command

```
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

```
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

```
Initialize keystore and settings

Usage:
  jtb-totp init [flags]

Flags:
  -f, --force             force re-initialization (required)
  -h, --help              help for init
  -i, --initialize        create new keystore/config file
  -k, --keystore string   location of new keystore path
  -p, --password string   encrypt datastore with user-defined password
```

This command creates the initial encrypted keystore and configuration file, then outputs where they were created.  During the creation process, if you do not specify a password (`--password/-p`) the program will generate a 32 character random password to encrypt the keystore with.

You can change where the keystore is created by using the `--keystore/-k` option.

```console
$ jtb-totp init --initialize
Initialized keystore and configuration files!
- Config file:   /home/jweatherly/.config/jtb-totp/jtb-totp.conf
- Keystore file: /home/jweatherly/.local/share/jtb-totp/keystore.enc
```

### `list` Command

```
List all keys by name in keystore

Usage:
  jtb-totp list [flags]

Flags:
  -h, --help   help for list
```

This command lists all of the keys by name that exist in the keystore in alphabetical order.

```console
$ jtb-totp list
bash.org
Google
Super Secret Key
Test Key
```

### `remove` Command

```
Remove key from keystore

Usage:
  jtb-totp remove [key name] [flags]

Flags:
  -h, --help   help for remove
```

The `remove` command deletes a key from the keystore if present, or complains a bit if the key is not present.  The key name must be an exact match to the one in the keystore.

```console
$ jtb-totp remove Google
Deleted key 'Google' from the keystore.
$ jtb-totp remove Google
Could not find key 'Google' in the keystore.
```