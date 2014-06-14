# Git

![Git](../../assets/images/git.png)

- Version Control System
- Source Control Management Tool - Handles the management of changes to documents. It allows developers to manage multiple versions of the same codebase, which comes in handy when bugs get accidently introduced and the application needs to be reverted to a previous working state.

## Commands

### git init - initializes a new git repository in current working directory

### git add
- if done on a file that is not being tracked, this command will add file to tracked files and adds it to staging area
- when done on a file that is already being tracked, this command will just add current state of file to staging area
- you can pass a filename to this command to add one file of you can type '.' to add all unstaged changes

###### Examples:
`git add a.txt`

`git add .`
***

### git commit
- commits staged changes
- this command takes one mandatory flag ('-m'), which stands for the message you want to attach to this commit

###### Examples:
`git commit -m 'Initial commit'`
***

### git pull
- pulls codebase from hosting server and merges it with your codebase

###### Examples:
`git pull origin master`
***

### git push
- pushs your code to hosting server and merges it with its codebase
- this command takes two arguments:
	- the first argument is the remote server to which you wish to push your code
	- the second argument is the branch that you wish to push

###### Examples:
`git push origin master`
***

### git branch
- duplicates your current branch into a new branch
- changes made in one branch do not affect other branches until the two branches are merged together

###### Examples:
`git branch new_branch`
***

### git checkout
- moves you into the branch whose name is passed to this command

###### Examples:
`git checkout new_branch`

- you can create a new branch and checkout into it in one command by doing:
`git checkout -b new_branch`

***

### git merge
- merges one branch with the current branch

###### Examples:
`git merge new_branch`
***

## Best Practices
- When writing a commit message, explicitly state the changes you have made
