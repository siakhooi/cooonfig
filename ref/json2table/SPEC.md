# json2table Spec Reference

This document describes the JSON spec format used by `json2table`.

## Top-level structure

```json
{
  "dataPath": "$.items",
  "columns": [
    {
      "path": "id",
      "title": "ID"
    }
  ]
}
```

- `dataPath` (optional, string)
  - JSONPath to select the array to render.
  - Default: `$`.
- `columns` (required, array, minimum 1 item)
  - Column definitions for table output.

## Column fields

Each item in `columns` supports:

- `path` (required, string or string[])
  - JSONPath(s) for the value.
  - If array is provided, the first non-null successful path is used.
- `title` (optional, string)
  - Header text.
  - Default: derived from the first `path` segment (last token after `.`).
- `minWidth` (optional, integer)
  - Minimum width constraint.
  - Must be `>= 0` and `<= maxWidth`.
- `maxWidth` (optional, integer)
  - Maximum width constraint.
  - Must be `>= 0` and `>= minWidth`.
  - Default: effectively unlimited.
- `align` (optional, string)
  - Allowed values: `left`, `right`, `center`.
  - Default: `left`.
- `urlPath` (optional, string)
  - JSONPath for a URL to render terminal links.
- `color` (optional)
  - Text color/style specification (see below).
  - Default: `"default"`.

## `color` formats

`color` supports 3 formats:

1. **Single string**

```json
"color": "red"
```

2. **Array of strings**

```json
"color": ["yellow", "italic", "underline"]
```

3. **Conditional object**

```json
"color": {
  "default": "bgRed",
  "conditions": [
    {
      "when": "Alice",
      "color": "green"
    },
    {
      "when": ["Bob", "Carol"],
      "color": ["yellow", "bold"]
    }
  ]
}
```

Conditional object fields:

- `default` (optional, string or string[])
  - Fallback color/style.
  - If omitted, defaults to `"default"`.
- `conditions` (optional, array)
  - Each condition has:
    - `when` (string or string[]): exact value match list.
    - `color` (string or string[]): style applied when matched.

## Supported color/style values

Foreground:

- `default`, `red`, `green`, `blue`, `yellow`, `magenta`, `cyan`, `white`, `black`
- `hiRed`, `hiGreen`, `hiBlue`, `hiYellow`, `hiMagenta`, `hiCyan`, `hiWhite`, `hiBlack`

Background:

- `bgRed`, `bgGreen`, `bgBlue`, `bgYellow`, `bgMagenta`, `bgCyan`, `bgWhite`, `bgBlack`
- `hiBgRed`, `hiBgGreen`, `hiBgBlue`, `hiBgYellow`, `hiBgMagenta`, `hiBgCyan`, `hiBgWhite`, `hiBgBlack`

Text styles:

- `bold`, `faint`, `italic`, `underline`, `blinkSlow`, `blinkRapid`, `reverseVideo`, `concealed`, `crossedOut`

## Full example

```json
{
  "dataPath": "$.data2",
  "columns": [
    {
      "path": "id",
      "title": "ID",
      "minWidth": 8,
      "align": "center",
      "urlPath": "url",
      "color": "red"
    },
    {
      "path": "url",
      "color": ["yellow", "italic", "underline"]
    },
    {
      "path": ["display.name", "display.name1"],
      "maxWidth": 8,
      "align": "left",
      "color": {
        "default": "bgRed",
        "conditions": [
          {
            "when": "ninnyhammer",
            "color": "cyan"
          },
          {
            "when": ["Alice", "baz"],
            "color": ["green", "bold"]
          }
        ]
      }
    }
  ]
}
```
