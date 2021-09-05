## How to contribute to climon operator

#### **Did you find a bug?**

* **Ensure the bug was not already reported** by searching on GitHub under [Issues](https://github.com/devopstoday11/climon/issues)

* If you're unable to find an open issue addressing the problem, [open a new one](https://github.com/devopstoday11/climon/issues/new). Be sure to include a **title and clear description**, as much relevant information as possible, and a **code sample** or an **executable test case** demonstrating the expected behavior that is not occurring.

#### **Did you write a patch that fixes a bug?**

* Open a new GitHub pull request with the patch.

* Ensure the PR description clearly describes the problem and solution. Include the relevant issue number if applicable.


## Guidelines

We recommend discussing your plans on our Slack (join our [channel](https://soi-dev.slack.com/archives/C022EFUL0GH) before starting to code - especially for more ambitious contributions. This gives other contributors a chance to point you in the right direction, give feedback on your design, and maybe point out if someone else is working on the same thing.

## Conventions

Fork the repo and make the changes on your fork in a feature branch based on the master branch:

- If it’s a bugfix branch, name it fix/XXX-something where XXX is the number of the issue
- If it’s a feature branch, create an enhancement issue to announce your intentions, and name it feature/XXX-something where XXX is the number of the issue.
- Submit unit tests for your changes. Go has a great test framework built in; use it! Take a look at existing tests for inspiration. Run the full test suite on your branch before submitting a pull request.
- Make sure you include relevant updates or additions to documentation when creating or modifying features.
- Write clean code. Universally formatted code promotes ease of writing, reading, and maintenance. Always run go fmt before committing your changes. Most editors have plugins that do this automatically.

