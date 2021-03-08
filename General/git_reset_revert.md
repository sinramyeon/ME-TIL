# Git reset vs Git revert

### git reset

The `git reset` command is a complex and versatile tool for undoing changes. It has three primary forms of invocation. These forms correspond to command line arguments `--soft, --mixed, --hard`.

let's see the example. if we do

`git reset --hard a0fvf8` then we got,

![](https://www.pixelstech.net/article/images/git-reset-1.jpg)

But now the remote origin still has HEAD point to commit D, if we directly use  **git push**  to push the changes, it will not update the remote repo, we need to add a  **-f**  option to force pushing the changes.

```bash
git push -f
```

### git reset options

#### 1. --hard

When passed `--hard` The Commit History ref pointers are updated to the specified commit. Then, the Staging Index and Working Directory are reset to match that of the specified commit. Any previously pending changes to the Staging Index and the Working Directory gets reset to match the state of the Commit Tree. This means any pending work that was hanging out in the Staging Index and Working Directory will be lost. 

#### 2. --mixed

This is the default operating mode. The ref pointers are updated. The Staging Index is reset to the state of the specified commit. Any changes that have been undone from the Staging Index are moved to the Working Directory. `-mixed` is the default mode and the same effect as executing `git reset`.

#### 3. --soft

the ref pointers are updated and the reset stops there. The Staging Index and the Working Directory are left untouched. 

#### usually we use the --hard way.

ex)
```
$ git reset --soft HEAD~1
```

Reset will rewind your current HEAD branch to the specified revision. In our example above, we'd like to return to the  _one before the current_  revision - effectively making our last commit undone.

If you don't want to keep these changes, simply use the --hard flag. Be sure to only do this when you're sure you don't need these changes anymore.

```
$ git reset --hard HEAD~1
```

### Don't Reset Public History

You should never use  `git reset`  when any snapshots after  have been pushed to a public repository. After publishing a commit, you have to assume that other developers are reliant upon it. 

![](https://wac-cdn.atlassian.com/dam/jcr:b616f03d-5257-4ea8-a6eb-db1a0207a78a/07%20(1).svg?cdnVersion=1491)

The point is, make sure that you’re using `git reset` on a local experiment that went wrong—not on published changes. If you need to fix a public commit, the `git revert` command was designed specifically for this purpose.

### Git revert

Git revert is like 'undoing' the commits, but this prevents Git from losing history, which is important for the integrity of your revision history and for reliable collaboration.
Where the `reset` command moves the branch pointer back in the chain (typically) to "undo" changes, the `revert` command adds a new commit at the end of the chain to "cancel" changes.

`$ git revert HEAD`

Because this adds a new commit, Git will prompt for the commit message:
```
Revert "File with three lines"  
  
This reverts commit b764644bad524b804577684bf74e7bca3117f554.  
  
# Please enter the commit message for your changes. Lines starting  
# with '#' will be ignored, and an empty message aborts the commit.  
# On branch master  
# Changes to be committed:  
#       modified:   file1.txt  
#
```
Figure 3 (below) shows the result after the  `revert`  operation is completed.

If we do a  `git log`  now, we'll see a new commit that reflects the contents before the previous commit.

```
$ git log --oneline  
11b7712 Revert "File with three lines"  
b764644 File with three lines  
7c709f0 File with two lines  
9ef9173 File with one line
```

![Git Revert – NUKE Designs | Blog](http://blog.nakulrajput.com/wp-content/uploads/2018/10/Git-Reverting-Resetting.jpg)

### git revert options

The  `git revert`  command can be considered an 'undo' type command, however, it is not a traditional undo operation. Instead of removing the commit from the project history, it figures out how to invert the changes introduced by the commit and appends a new commit with the resulting inverse content. This prevents Git from losing history, which is important for the integrity of your revision history and for reliable collaboration.

Reverting should be used when you want to apply the inverse of a commit from your project history. This can be useful, for example, if you’re tracking down a bug and find that it was introduced by a single commit. Instead of manually going in, fixing it, and committing a new snapshot, you can use  `git revert`  to automatically do all of this for you.

![Git revert - Atlassian git tutorials](https://wac-cdn.atlassian.com/dam/jcr:b6fcf82b-5b15-4569-8f4f-a76454f9ca5b/03%20(7).svg?cdnVersion=1491)

## How it works

The  `git revert`  command is used for undoing changes to a repository's commit history. Other 'undo' commands like,  `[git checkout](https://www.atlassian.com/git/tutorials/using-branches/git-checkout)`  and  `[git reset](https://www.atlassian.com/git/tutorials/undoing-changes/git-reset)`, move the  `HEAD`  and branch ref pointers to a specified commit.  `Git revert`  also takes a specified commit, however,  `git revert`  does not move ref pointers to this commit. A revert operation will take the specified commit, inverse the changes from that commit, and create a new "revert commit". The ref pointers are then updated to point at the new revert commit making it the tip of the branch.  
  
To demonstrate let’s create an example repo using the command line examples below:

```
$ mkdir git_revert_test$ cd git_revert_test/$ git init .Initialized empty Git repository in /git_revert_test/.git/$ touch demo_file$ git add demo_file$ git commit -am"initial commit"[master (root-commit) 299b15f] initial commit 1 file changed, 0 insertions(+), 0 deletions(-) create mode 100644 demo_file$ echo "initial content" >> demo_file$ git commit -am"add new content to demo file"[master 3602d88] add new content to demo filen 1 file changed, 1 insertion(+)$ echo "prepended line content" >> demo_file$ git commit -am"prepend content to demo file"[master 86bb32e] prepend content to demo file 1 file changed, 1 insertion(+)$ git log --oneline86bb32e prepend content to demo file3602d88 add new content to demo file299b15f initial commit
```

Here we have initialized a repo in a newly created directory named  `git_revert_test`. We have made 3 commits to the repo in which we have added a file  `demo_file`  and modified its content twice. At the end of the repo setup procedure, we invoke  `git log`  to display the commit history, showing a total of 3 commits. With the repo in this state, we are ready to initiate a  `git revert.`

```
$ git revert HEAD [master b9cd081] Revert "prepend content to demo file" 1 file changed, 1 deletion(-)
```

`Git revert`  expects a commit ref was passed in and will not execute without one. Here we have passed in the  `HEAD`  ref. This will revert the latest commit. This is the same behavior as if we reverted to commit  `3602d8815dbfa78cd37cd4d189552764b5e96c58`. Similar to a merge, a revert will create a new commit which will open up the configured system editor prompting for a new commit message. Once a commit message has been entered and saved Git will resume operation. We can now examine the state of the repo using  `git log`  and see that there is a new commit added to the previous log:

```
$ git log --oneline 1061e79 Revert "prepend content to demo file" 86bb32e prepend content to demo file 3602d88 add new content to demo file 299b15f initial commit
```

Note that the 3rd commit is still in the project history after the revert. Instead of deleting it,  `git revert`  added a new commit to undo its changes. As a result, the 2nd and 4th commits represent the exact same code base and the 3rd commit is still in our history just in case we want to go back to it down the road.

## Common options

```
-e--edit
```

This is a default option and doesn't need to be specified. This option will open the configured system editor and prompts you to edit the commit message prior to committing the revert

```
--no-edit
```

This is the inverse of the  `-e`  option. The revert will not open the editor.

```
-n--no-commit
```

Passing this option will prevent  `git revert`  from creating a new commit that inverses the target commit. Instead of creating the new commit this option will add the inverse changes to the Staging Index and Working Directory.

### Revert or reset?

Why would you choose to do a  `revert`  over a  `reset`  operation? If you have already pushed your chain of commits to the remote repository (where others may have pulled your code and started working with it), a **revert** is a nicer way to cancel out changes for them. This is because the Git workflow works well for picking up additional commits at the end of a branch, but it can be challenging if a set of commits is no longer seen in the chain when someone resets the branch pointer back.

This brings us to one of the fundamental rules when working with Git in this manner: **Making these kinds of changes in your  _local repository_**  to code you haven't pushed yet is fine. But avoid making changes that rewrite history if the commits have already been pushed to the remote repository and others may be working with them.

In short, if you rollback, undo, or rewrite the history of a commit chain that others are working with, your colleagues may have a lot more work when they try to merge in changes based on the original chain they pulled. If you must make changes against code that has already been pushed and is being used by others, consider communicating before you make the changes and give people the chance to merge their changes first. Then they can pull a fresh copy after the infringing operation without needing to merge.

---

Wait.... i've heard there are `git rm` and `git clean` ....

### 1. git rm

The primary function of  `git rm`  is to remove tracked files from the Git index. Additionally,  `git rm`  can be used to remove files from both the staging index and the working directory. There is no option to remove a file from only the working directory. The files being operated on must be identical to the files in the current  `HEAD`. If there is a discrepancy between the  `HEAD`  version of a file and the staging index or working tree version, Git will block the removal. This block is a safety mechanism to prevent removal of in-progress changes.

Note that  `git rm`  does not remove branches.  to remove branch, do `git branch -d 브랜치명`

### 2. git clean

_This is builtin command to cleanup the untracked files._  **Be careful with this one, it deletes files permanently!**

Always add  `-n`  or  `--dry-run`  options to preview the damage you'll do!

-   If you just clean untracked files, run  `git clean -f`
-   If you want to also remove directories, run  `git clean -f -d`
-   If you just want to remove ignored files, run  `git clean -f -X`
-   If you want to remove ignored as well as non-ignored files, run  `git clean -f -x`

Note, that there is  `-f`  /  `--force`  option in each example, this is because of  [default configuration of git](https://www.kernel.org/pub/software/scm/git/docs/git-clean.html):  _If the git configuration variable clean.requireForce is not set to false, git clean will refuse to run unless given -f or -n._
