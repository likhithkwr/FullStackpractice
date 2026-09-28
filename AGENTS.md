# Repository instructions

This is a shared interview preparation and programming practice repository.

## Organization

- Keep shared explanations in `Explanation/<topic>.md` and reusable definitions in `Explanation/terminology.md`.
- Keep implementations in `<Language>/<topic>/`, using existing language/framework folders where appropriate.
- Use simple conventional folder names such as `Java`, `Python`, and `React`.
- The current practice folders are `Java`, `Python`, `React`, `JavaScript`, `TypeScript`, `SQL`, and `Go`.
- Use lowercase topic slugs with hyphens, such as `two-sum`.
- Follow the language's source-file naming conventions.
- Add another language or framework folder only when the user requests practice in it, with a short README.
- Keep each runnable project's dependencies and lockfile inside its topic folder.

## Teaching requirements

Follow `Explanation/README.md`: explain what the topic is, how it works, where it is useful, and its tradeoffs. Include worked examples, diagrams, a clearly labeled hypothetical industry scenario, edge cases, and interview follow-up questions with answers. Define terminology and technical words in plain language when first introduced. Explain relevant language syntax and execution tools, and add reusable definitions to the terminology guide.

Use the language requested by the learner. Add links to available implementations and commands to run them. Update the root topic index and relevant folder indexes when adding a topic or implementation.

## Verification and collaboration

- Run the affected example or relevant checks before reporting it as working.
- Preserve other contributors' work and unrelated local changes.
- Keep generated files, dependencies, local environments, and secrets out of commits.
- Use ordinary Git commits and the shared branch workflow in `CONTRIBUTING.md`.
- Never force push shared history as part of routine practice work.
