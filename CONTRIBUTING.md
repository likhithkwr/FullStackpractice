# Contributing

Use this repository for shared practice. Organize code by language and topic, and use a separate branch for each person's work.

## Git workflow

Start from an up-to-date `main` branch:

```sh
git switch main
git pull --ff-only
git switch -c practice/your-name/two-sum-python
```

Replace `your-name` with your name and `two-sum-python` with the topic and language you are practicing.

1. Add or update code in the appropriate language/topic folder.
2. Add or update the shared explanation in `Notes`.
3. Update the topic links in the root README and relevant folder READMEs.
4. Run the example or relevant checks and include the command in the topic instructions.
5. Review and commit only the files belonging to your change.

For a change to the existing Python example:

```sh
git diff
git add Python/two-sum/solution.py Notes/two-sum.md README.md
git commit -m "Improve Python Two Sum explanation"
git push -u origin HEAD
```

Open a pull request into `main` for the other contributor to review. Once it is merged, switch to `main` and pull before starting another topic.

Direct pushes require repository write access. A repository owner can add a friend's GitHub account as a collaborator. Contributors without write access can work in a fork and open a pull request.

## Adding another language or framework

- Create a folder with a simple conventional name, such as `Java`, `Python`, `React`, or `Vue`.
- Add a README describing prerequisites and how its topic folders are organized.
- Add one folder per topic using lowercase names with hyphens.
- Keep each exercise's dependency manifest and lockfile with that exercise.
- Record any required runtime version, setup command, and run command in the exercise README when applicable.

## Keeping practice easy to review

- Preserve other contributors' work and describe the approach in shared notes.
- Use descriptive filenames for alternative solutions.
- Commit dependency lockfiles for runnable applications.
- Keep generated output, installed dependencies, local environments, credentials, and personal settings out of commits.
- Use the normal branch and pull request workflow when updating shared work.
