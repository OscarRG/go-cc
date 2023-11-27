
# go-cc - Conventional Commit CLI

`go-cc` is a versatile CLI tool written in Go, designed to facilitate the generation and management of conventional commits in your Git repository. It adheres to the Conventional Commit standard for commit messages, enhancing your project's versioning and making your commit history more organized and easier to manage.


## Features

- **Repository-Level Configuration**: Set default conventional commit configurations specific to your repository with ease.
- **Interactive Command-Line Prompts**: Enjoy a user-friendly experience with interactive command-line prompts for creating conventional commits.
- **Automatic Initialization**: When no .gocc file is found in the repository, go-cc seamlessly defaults to a predefined configuration for conventional commits.
- **Supports Commit Types, Optional Scopes, and Commit Descriptions**: Effortlessly define commit types, optional scopes, and commit descriptions tailored to your project.
- **Review and Confirm Commit**: go-cc provides the option to review and confirm commits before execution, ensuring a final check for changes.
- **Seamless Git Workflow Integration**: Integrate go-cc seamlessly into your Git workflow, establishing a consistent and standardized approach to conventional commits.
- **Easy Versioning**: Manage GitHub conventional commits effortlessly and consistently, promoting a streamlined versioning process.


## Installation

To install `go-cc`, you need to have [Go](https://go.dev) installed. You can install the tool using the following steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/OscarRG/go-cc.git
   ```

2. Change to the project directory:

   ```bash
   cd go-cc
   ```

3. Build the executable:

   ```bash
   make build
   ```

4. Install the executable (optional, for system-wide usage):

   ```bash
   make install
   ```

## Usage

To use `go-cc`, follow these steps:

1. Open your terminal and navigate to your Git repository.

2. Run the `go-cc` command:

   ```bash
   go-cc
   ```

3. You will see a menu with various commit types, each accompanied by an emoji.

4. Use the arrow keys to select the desired commit type and press Enter.

5. You will be prompted to enter an optional scope (leave empty for none) and the commit message description.

6. After providing this information, `go-cc` will generate a conventional commit message and display it for your review.

7. Press Enter to confirm and execute the commit, or press 'C' to cancel.

## Example

Here's an example of using `go-cc`:

```bash
$ go-cc

Select a commit type:
➔ ✨ feat
  🐛 fix
  📄 docs
  💅 style
  🛠️  refactor
  🎯 perf
  🧪 test
  👷 build
  🔃 ci
  🧹 chore
  🔙 revert

Enter the commit type (feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert): feat
Enter an optional scope (leave empty for none): user-auth
Enter the commit message description: implement user login

Generated commit: feat(user-auth): implement user login

Press Enter to confirm or 'C' to cancel: 
```

## Uninstallation

If you installed `go-cc` system-wide and want to remove it, you can run the following command:

```bash
make clean
```
