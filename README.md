# go_console_projects
Educational project - simple four console projects

# Smart Utilities

A collection of console-based utilities written in Go. This project demonstrates key programming skills like input parsing, basic algorithms, and working with Go's standard library. Each utility is self-contained and accessible via the terminal.

---

## 📦 Features

### 1. Console Calculator

A basic CLI calculator supporting:

- Addition, subtraction, multiplication, division
- Input via terminal (left operand, operator, right operand)
- Float64 precision with up to 3 decimal places for division
- Error handling: prompts user again if invalid input

**Example:**
```
Input left operand:
> 10

Input operation:
> +

Input right operand:
> 15

Result:
25.000
```

---

### 2. Most Frequent Words

Reads a space-separated string of words and a number `K`, returning the `K` most frequent words.

- Words sorted by frequency (desc), then lexicographically
- Handles edge cases: fewer than K words, empty input
- Includes unit tests

**Example:**
```
Input:
aa bb cc aa cc cc cc aa ab ac bb
3

Output:
cc aa bb
```

---

### 3. Slice Intersection

Reads two lists of integers and returns their intersection:

- Maintains the order from the first list
- Returns `"Empty intersection"` if none found
- Validates integer-only input

**Example:**
```
Input:
5 3 4 2 1 6
6 4 2 4

Output:
4 2 6
```

---

### 4. Visit Log System

An interactive in-memory patient visit tracking system. Supports:

- `Save`: Stores full name, doctor specialization, and date
- `GetHistory`: Returns all visits of a patient
- `GetLastVisit`: Returns last visit to a particular specialist
- Custom error: `UserNotFoundError`

**Example:**
```
Save
Ivanov Ivan Ivanovich
orthopedist
2024-04-13

GetHistory
Ivanov Ivan Ivanovich

Output:
orthopedist 2024-04-13

GetLastVisit
Ivanov Ivan Ivanovich
orthopedist

Output:
2024-04-13
```

---

## 🧪 Running Tests

All packages use the `testing` standard library.

Run all tests:

```bash
make test
```

Get coverage report:

```bash
make coverage
```

---

## 🌐 View Documentation Locally

You can run a local documentation server with `godoc` or with `pkgsite` and open it in your browser.

```bash
make dvi
```

And then open: [http://localhost:8080](http://localhost:8080)

---

## 📂 Running

```bash
make
```

And then just write `make ...` - all main targets will be shown after comiling executable files!

---

## ⚙️ Requirements

- Go 1.20+
- Only standard library is used

---

## 📄 Author
- [Anton Evgenev](https://t.me/tdutanton)

2025