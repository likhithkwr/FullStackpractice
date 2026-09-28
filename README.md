# Full Stack Practice

Shared interview preparation and programming practice, organized by language or framework. Each topic has one shared explanation and separate implementations in the languages being practiced.

## Get started

```sh
git clone https://github.com/likhithkwr/FullStackpractice.git
cd FullStackpractice
```

Install the runtime or tools for the language you want to practice. Each language folder describes its prerequisites; runnable projects keep their own dependencies and instructions inside their topic folder.

Run a Two Sum example from the repository root:

```sh
python3 Python/two-sum/solution.py
```

Or run the Java version with JDK 11 or newer:

```sh
java Java/two-sum/TwoSum.java
```

## Folder structure

```text
Explanation/            Shared guides, diagrams, examples, and terminology
Java/                   Java practice
Python/                 Python practice
React/                  React components and applications
JavaScript/             JavaScript practice
TypeScript/             TypeScript practice
SQL/                    SQL queries and database exercises
HTML/                   HTML practice
CSS/                    CSS practice
C/                      C practice
Cpp/                    C++ practice
CSharp/                 C# practice
Go/                     Go practice
Rust/                   Rust practice
Kotlin/                 Kotlin practice
Swift/                  Swift practice
PHP/                    PHP practice
Ruby/                   Ruby practice
```

The structure can accommodate any other language or framework. Add a folder with its conventional name, such as `Vue`, `Django`, or `SpringBoot`, when starting practice in it.

## Naming

- Use simple language and framework names: `Java`, `Python`, `React`.
- Use lowercase topic names with hyphens: `two-sum`, `binary-search`, `todo-app`.
- Put shared explanations in `Explanation/<topic>.md` and common definitions in `Explanation/terminology.md`.
- Put code in `<Language>/<topic>/`.
- Follow source-file conventions for the language: `TwoSum.java`, `solution.py`, `TwoSum.jsx`.
- Give alternative approaches descriptive filenames, such as `brute_force.py` and `hash_map.py`.

Example of the same problem practiced in multiple languages:

```text
Explanation/two-sum.md
Python/two-sum/solution.py
Java/two-sum/TwoSum.java
```

## Topics

| Topic | Explanation | Available implementations |
| --- | --- | --- |
| Two Sum | [Explanation](Explanation/two-sum.md) | [Python](Python/two-sum/solution.py), [Java](Java/two-sum/README.md) |

See [the explanation format](Explanation/README.md), the [terminology guide](Explanation/terminology.md), and [CONTRIBUTING.md](CONTRIBUTING.md) for the shared Git workflow.
