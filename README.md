# xlcat

**xlcat** is a lightweight, command-line tool that allows you to quickly browse and preview Excel files (`.xlsx`, `.xlsm`) directly from your terminal.

## 🚀 Features

- List all sheets in an Excel file with row and column counts
- Preview a specified number of rows from any sheet
- Automatically detect and list Excel files in directories
- Support for multiple Excel file extensions (`.xlsx`, `.xlsm`, `.xltx`, `.xltm`)

## 📥 Clone & build

```sh
git clone https://github.com/artemdorozhkin/xlcat.git
cd xlcat
```

```sh
go build
```

## 🛠 Usage

```
xlcat <file.xlsx> [options]
```

### Options

| Option    | Short | Description               | Default |
| --------- | ----- | ------------------------- | ------- |
| `--rows`  | `-r`  | Number of rows to preview | `5`     |
| `--sheet` | `-s`  | Sheet name to preview     | (none)  |
| `--help`  | `-h`  | Show help information     |         |

### Examples

Preview the first 10 rows of a specific sheet:

```sh
xlcat sales.xlsx --sheet "Q4" --rows 10
```

List all sheets in a file with their dimensions:

```sh
xlcat data.xlsx
```

List Excel files in a directory:

```sh
xlcat ./reports
```

## 📦 Dependencies

- [excelize](https://github.com/xuri/excelize) — Excel file handling
- [tablewriter](https://github.com/olekukonko/tablewriter) — Terminal tables

---

Made with ❤️ by [artemdorozhkin](https://github.com/artemdorozhkin).
