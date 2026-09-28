# Full Stack Practice

Shared interview preparation and programming practice, organized by language or framework. Each topic has one shared explanation and separate implementations in the languages being practiced.

## Get started

```sh
git clone https://github.com/likhithkwr/FullStackpractice.git
cd FullStackpractice
```

Install the runtime or tools for the language you want to practice. Each language folder describes its prerequisites; runnable projects keep their own dependencies and instructions inside their topic folder.

Run the existing Python example from the repository root:

```sh
python3 Python/two-sum/solution.py
```

## Folder structure

```text
Notes/                  Shared explanations, examples, diagrams, and scenarios
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
- Put shared notes in `Notes/<topic>.md`.
- Put code in `<Language>/<topic>/`.
- Follow source-file conventions for the language: `TwoSum.java`, `solution.py`, `TwoSum.jsx`.
- Give alternative approaches descriptive filenames, such as `brute_force.py` and `hash_map.py`.

Example of the same problem practiced in multiple languages:

```text
Notes/two-sum.md
Python/two-sum/solution.py
Java/two-sum/TwoSum.java       # Add when practicing the Java version
```

## Topics

| Topic | Explanation | Available implementations |
| --- | --- | --- |
| Two Sum | [Notes](Notes/two-sum.md) | [Python](Python/two-sum/solution.py) |

See [the notes format](Notes/README.md) for what each explanation covers, and [CONTRIBUTING.md](CONTRIBUTING.md) for the shared Git workflow.
