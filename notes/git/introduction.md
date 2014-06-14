# Git

![Git](../../assets/images/git.png)

- Version Control System
- Source Control Management Tool - Handles the management of changes to documents. It allows developers to manage multiple versions of the same codebase, which comes in handy when bugs get accidently introduced and the application needs to be reverted to a previous working state.

## Commands

### git init

```
~> git init
```

- initialises a new git repository in current working directory

### git add

```
~>  git add a.txt
```

```
~>  git add --all
```

- if done on a file that is not being tracked, this command will add file to tracked files and adds it to staging area
- when done on a file that is already being tracked, this command will just add current state of file to staging area
- you can pass a filename to this command to add one file of you can type '.' to add all unstaged changes

### git add

```
~>  git diff
```

```
~>  git diff a.txt
```

### git rm


```
~>  git rm a.txt
```

This command is used to remove files from Git. Running it as shown above will
also remove the file from your computer. You can run this command with `--cached`
to remove the file from the git staging area.

### git commit

```
~> git commit -m 'Initial commit'
```
- commits staged changes
- this command takes one mandatory flag ('-m'), which stands for the message you want to attach to this commit

### git pull

```
~> git pull origin master
```

- pulls codebase from hosting server and merges it with your codebase

### git push

```
~> git push origin master
```

- pushs your code to hosting server and merges it with its codebase
- This command takes two arguments. The first argument is the remote server to
 which you wish to push your code. The second argument is the branch that you
 wish to push to.



### git branch

```
~> git branch new_branch
```

- duplicates your current branch into a new branch
- changes made in one branch do not affect other branches until the two branches are merged together

### git checkout



### git merge

```
~> git merge new_branch
```

- merges one branch with the current branch


## Branching

![Git Branching](../../assets/images/git-branching.png)

Branching means you diverge from the main line of development and continue to
do work without messing with that main line.  

When working on a new feature or a bug fix, it is best practice to create a
new branch for the new code you are going to write and to give that branch
a descriptive name, such as 'feature/add-navigation-bar'.

To create a branch, you can use the git checkout command like so:

```
~> git checkout -b new_branch
```

The above command will create a new branch with the name new_branch and
immediately move you into that branch.

Once a branch already exists and you would like to move into it, use the
checkout command like so:

```
~> git checkout new_branch
```





## Best Practices
- When writing a commit message, explicitly state the changes you have made
