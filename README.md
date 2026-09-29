# Full Stack Practice

Shared interview preparation and programming practice, organized by language or framework. Each topic has one shared explanation and separate implementations in the languages being practiced.

## Get started

```sh
git clone https://github.com/likhithkwr/FullStackpractice.git
cd FullStackpractice
```

Install the runtime or tools for the language you want to practice. Each language folder describes its prerequisites; runnable projects keep their own dependencies and instructions inside their topic folder.

Start with a Two Sum practice file from the repository root:

```sh
python3 Python/two-sum/practice.py
```

Or use the Java practice file with JDK 11 or newer:

```sh
java Java/two-sum/TwoSumPractice.java
```

Implement the `twoSum()` method marked `TODO`, then run the same file again. Its checker reports `PASS`, `FAIL`, or `TODO` for each case. The initial unfinished file reports `TODO`; see the language guide for the exit codes.

Each problem has a completed **reference** and a separate **practice** file for your own implementation. Use your own Git branch when practicing with a friend.

## Folder structure

```text
Explanation/            Shared guides, diagrams, examples, and terminology
Java/                   Java practice
Python/                 Python practice
React/                  React components and applications
JavaScript/             JavaScript practice
TypeScript/             TypeScript practice
SQL/                    SQL queries and database exercises
Go/                     Go practice
```

Add another language or framework folder when you decide to practice it, using a simple conventional name.

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
Python/two-sum/practice.py
Java/two-sum/TwoSum.java
Java/two-sum/TwoSumPractice.java
```

## Topics

| Topic | Explanation | References | Your practice |
| --- | --- | --- | --- |
| Two Sum | [Explanation](Explanation/two-sum.md) | [Java](Java/two-sum/TwoSum.java), [Python](Python/two-sum/solution.py) | [Java](Java/two-sum/TwoSumPractice.java), [Python](Python/two-sum/practice.py) |

See [the explanation format](Explanation/README.md), the [terminology guide](Explanation/terminology.md), and [CONTRIBUTING.md](CONTRIBUTING.md) for the shared Git workflow.
