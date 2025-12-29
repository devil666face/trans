# Google Translate CLI (`trans`)

A small command‑line client for Google Translate. It lets you translate text directly from your terminal (from an argument or from standard input) and list all supported languages.

## What this project does

`trans` is designed to:

- translate text from the terminal without opening a browser;
- be used in shell scripts and pipelines (via stdin);
- quickly show which languages are supported as source and target.

Under the hood it uses the unofficial HTTP endpoint `https://translate.googleapis.com/translate_a/single`.

## Installation

From the repository root:

```bash
go build -o trans ./cmd/trans
# or
go install ./cmd/trans
```

After building, run `./trans` (or just `trans` if the binary is on your `$PATH`).

## Command overview

```bash
trans [command] [--flags]
```

Available commands:

- `translate` (`t`) — translate text from an argument or stdin;
- `list` (`l`) — list available languages;
- `help` — show help for commands;
- `completion` — generate a shell completion script.

Global flags (provided by `cobra` / `fang`):

- `-h`, `--help` — help for a command;
- `-v`, `--version` — show `trans` version.

## Translating text

Command syntax:

```bash
trans translate [text] [--flags]
# or shorter
trans t [text] [--flags]
```

Defaults:

- target language (`--target`, `-t`) — `ru` (Russian);
- source language (`--source`, `-s`) — `auto` (auto‑detect).

### Basic examples

Translate a string passed as an argument:

```bash
trans translate -t ru "Hello, world"
# same using the alias
trans t -t ru "Hello, world"
```

Explicitly set source and target languages:

```bash
# short flags
trans translate -s en -t de "Good morning"

# long flags
trans translate --source en --target de "Good morning"
```

Translate text from stdin (source language is auto‑detected):

```bash
echo "How are you?" | trans translate -t ru
```

If you do not provide the `[text]` argument and there is no data in stdin, the command returns an error: `no input text provided`.

## Listing languages

Command:

```bash
trans list
# or
trans l
```

Prints all languages that can be used as source or target in the format `code  name`.

## Supported languages

List of supported languages in the current version (see `internal/trans/trans.go`):

| Code  | Language              |
| ----- | --------------------- |
| af    | Afrikaans             |
| sq    | Albanian              |
| am    | Amharic               |
| ar    | Arabic                |
| hy    | Armenian              |
| az    | Azerbaijani           |
| eu    | Basque                |
| be    | Belarusian            |
| bn    | Bengali               |
| bs    | Bosnian               |
| bg    | Bulgarian             |
| ca    | Catalan               |
| ceb   | Cebuano               |
| zh-CN | Chinese (Simplified)  |
| zh-TW | Chinese (Traditional) |
| co    | Corsican              |
| hr    | Croatian              |
| cs    | Czech                 |
| da    | Danish                |
| nl    | Dutch                 |
| en    | English               |
| eo    | Esperanto             |
| et    | Estonian              |
| fi    | Finnish               |
| fr    | French                |
| fy    | Frisian               |
| gl    | Galician              |
| ka    | Georgian              |
| de    | German                |
| el    | Greek                 |
| gu    | Gujarati              |
| ht    | Haitian Creole        |
| ha    | Hausa                 |
| haw   | Hawaiian              |
| he    | Hebrew                |
| hi    | Hindi                 |
| hmn   | Hmong                 |
| hu    | Hungarian             |
| is    | Icelandic             |
| ig    | Igbo                  |
| id    | Indonesian            |
| ga    | Irish                 |
| it    | Italian               |
| ja    | Japanese              |
| jv    | Javanese              |
| kn    | Kannada               |
| kk    | Kazakh                |
| km    | Khmer                 |
| ko    | Korean                |
| ku    | Kurdish               |
| ky    | Kyrgyz                |
| lo    | Lao                   |
| la    | Latin                 |
| lv    | Latvian               |
| lt    | Lithuanian            |
| lb    | Luxembourgish         |
| mk    | Macedonian            |
| mg    | Malagasy              |
| ms    | Malay                 |
| ml    | Malayalam             |
| mt    | Maltese               |
| mi    | Maori                 |
| mr    | Marathi               |
| mn    | Mongolian             |
| my    | Myanmar (Burmese)     |
| ne    | Nepali                |
| no    | Norwegian             |
| ny    | Nyanja (Chichewa)     |
| or    | Odia (Oriya)          |
| ps    | Pashto                |
| fa    | Persian               |
| pl    | Polish                |
| pt    | Portuguese            |
| pa    | Punjabi               |
| ro    | Romanian              |
| ru    | Russian               |
| sm    | Samoan                |
| gd    | Scots Gaelic          |
| sr    | Serbian               |
| st    | Sesotho               |
| sn    | Shona                 |
| sd    | Sindhi                |
| si    | Sinhala (Sinhalese)   |
| sk    | Slovak                |
| sl    | Slovenian             |
| so    | Somali                |
| es    | Spanish               |
| su    | Sundanese             |
| sw    | Swahili               |
| sv    | Swedish               |
| tl    | Tagalog (Filipino)    |
| tg    | Tajik                 |
| ta    | Tamil                 |
| tt    | Tatar                 |
| te    | Telugu                |
| th    | Thai                  |
| tr    | Turkish               |
| tk    | Turkmen               |
| uk    | Ukrainian             |
| ur    | Urdu                  |
| ug    | Uyghur                |
| uz    | Uzbek                 |
| vi    | Vietnamese            |
| cy    | Welsh                 |
| xh    | Xhosa                 |
| yi    | Yiddish               |
| yo    | Yoruba                |
| zu    | Zulu                  |

Use the language codes from this table with the `--source` / `-s` and `--target` / `-t` flags.
