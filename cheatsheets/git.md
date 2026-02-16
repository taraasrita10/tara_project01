# Git Cheatsheet

## Frequently used commands

| Command | Description | Comments |
| --------| ------------| ---------|
| git add <file> | Adds untracked file | . - for all, -p - for part of file |
| git rm <file> | Deletes file | - |
| git mv <old> <new> | Moves file | - |
| git reset | Unstages everything | For a particular file, mention <file> |
| git status | Check what you added | - |
| git commit -m <message> | Make a commit with the message | -am - all untracked changes |
| git checkout <name> | Change branch | Switch can be used instead of checkout |
| git branch -d <name> | Delete branch| -D for force delete |
| git checkout -b <name> | Create branch | Switch can be used instead of checkout |
| git clone <url> | Clone an existing repo | - |
| git init | Create a new repo | - |
| git stash | Stashes all changes | - |
| git restore | Deletes unstaged changes to a file | - |
| git reset | Resets the file to the previous commit | --hard, HEAD^ |
| git commit --amend | To amend the commit message, add forgotten files | - |
| git push origin main | Push the main to a remote origin | - |
| git push | Push current branch to its remote tracking branch | - |
| git pull | Update your current branch with changes | - |
| git fetch origin main | Fetch changes but does not update the current branch | - |
| git pull --rebase | Fetch changes and rebase your current branch | - |
| git rebase -i | Interactive rebase - opens a text editor with a list of all commits with actions to be taken | - |
| git rebase --onto <to> <from> <branch> | Move a branch from one branch to another | - |



