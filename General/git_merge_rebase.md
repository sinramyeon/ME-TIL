# Git merge vs Git rebase

### What is Git merge?

`git merge`.

A `merge` is a way to put a forked history back together.

**The  `git merge`  command lets us take independent branches of development and combine them into a single branch**.

`git merge`, the current branch will be updated to reflect the merge, but the target branch remains untouched.

![](https://reflectoring.io/assets/img/posts/git-rebase-merge/git-merge-working1.png)

![](https://reflectoring.io/assets/img/posts/git-rebase-merge/git-merge-working2.png)

so, `git merge` is jooin two or even more branches to development historrieess together.

```
_git merge_ [-n] [--stat] [--no-commit] [--squash] [--[no-]edit]
	[--no-verify] [-s <strategy>] [-X <strategy-option>] [-S[<keyid>]]
	[--[no-]allow-unrelated-histories]
	[--[no-]rerere-autoupdate] [-m <msg>] [-F <file>] [<commit>…​]
_git merge_ (--continue | --abort | --quit)
```

### Merge Conflict everywhere

sometimes we have `merge confflict` , when you make commits on separate branches that alter the same line in conflicting ways. If this happens, Git will not know which version of the file to keep in an error message similar to the following:

```
shell
CONFLICT (content): Merge conflict in resumé.txt Automatic merge failed; fix conflicts and then commit the result.
```

If you look at the  conflict file in your code editor, you can see where the conflict took place:

```
shell
<<<<<<< HEAD
잉잉이
=======
징징이
>>>>>>> 어쩌구
```

so you can easily find conflilct from searching `<<<<<` or `======` or `>>>>>` in your ide.

1. content of the current brrarnch(HEAD : ref is pointing to)

```
<<<<<<< HEAD
잉잉이
=======
```

2. content in the branch being merged. `>>>>>> 어쩌구` iss the content in the brranch being merged.

```
=======
징징이
>>>>>>> 어쩌구
```

### git merge options

#### Recursive

`git merge -s recursive branch1 branch2`

Reecursive is the default merge strategy when pulling or merging one branch. 

#### Resolve

`git merge -s resolve branch1 branch2`

This can only resolve two heads using a 3-way merge algorithm. 

#### Octopus

`git merge -s octopus branch1 branch2 branch3 branchN`

The default merge strategy for more than two heads. When more than one branch is passed octopus is automatically engaged. If a merge has conflicts that need manual resolution octopus will refuse the merge attempt. It is primarily used for bundling similar feature branch heads together.

#### Ours

`git merge -s ours branch1 branch2 branchN`

The Ours strategy operates on multiple number of branches.(more than 2)

#### Subtree

`git merge -s subtree branchA branchB`

This is an extension of the recursive strategy. When merging A and B, if B is a child subtree of A, B is first updated to reflect the tree structure of A, This update is also done to the common ancestor tree that is shared between A and B.

---

### Git rebase

`git rebase`

`rebase` is the way of migrating or combining a sequence of commits to a new base commit. so summary, Reapply commits on top of another base tip.

![](https://reflectoring.io/assets/img/posts/git-rebase-merge/git-rebasing-basic.png)

**Unlike merging, rebasing flattens history.** It transfers the completed work from one branch to another. In the process, unwanted history is eliminated. 

### Interactive Rebasing

Interactive rebasing gives you the opportunity to alter commits as they are moved to the new branch. This is even more powerful than an automated rebase, since it offers complete control over the branch’s commit history. Typically, this is used to clean up a messy history before merging a feature branch into  `master`.

`git rebase -i master`

This will open a text editor listing all of the commits that are about to be moved:

```
pick 33d5b7a Message for commit #1
pick 9480b3d Message for commit #2
pick 5c67e61 Message for commit #3
```

and you can pick / fixup / sqaush commits you want.

![](https://wac-cdn.atlassian.com/dam/jcr:fe6942b4-7a60-4464-9181-b67e59e50788/04.svg?cdnVersion=1486)

### Tip of Rebase

the most important thing to learn is when  _not_  to do it. The golden rule of  `git rebase`  is to **never use it on  _public_  branches.**

For example, think about what would happen if you rebased `master` onto your `feature` branch:

![](https://wac-cdn.atlassian.com/dam/jcr:1d22f018-b2c7-4096-9db1-c54940cf4f4e/05.svg?cdnVersion=1486)

The rebase moves all of the commits in `master` onto the tip of `feature`. The problem is that this only happened in _your_ repository. All of the other developers are still working with the original `master`. Since rebasing results in brand new commits, Git will think that your `master` branch’s history has diverged from everybody else’s.

So, before you run  `git rebase`, always ask yourself, “Is anyone else looking at this branch?” If the answer is yes, take your hands off the keyboard and start thinking about a non-destructive way to make your changes (e.g., the  `git revert`  command). Otherwise, you’re safe to re-write history as much as you like.

Also, what is the differrence between `git rebase` and `git merge`?

ex1 ) git merge

it keeps gold,red,green commits son the basee branch. the merge commit s

![](https://reflectoring.io/assets/img/posts/git-rebase-merge/git-merge-history.png)

**`git merge`  preserves the ancestry of commits.**

ex2) git rebase

_re-writes_  the changes of one branch onto another branch without the creation of a merge commit.

![](https://reflectoring.io/assets/img/posts/git-rebase-merge/git-rebase-history.png)

**It will be just like all commits have been written on top of the main branch all along.**

---

### Git Rebase vs. Merge: Which to Use

#### Git Rebase

-   Streamlines a potentially complex history.
-   Avoids merge commit “noise” in busy repos with busy branches.
-   Cleans intermediate commits by making them a single commit, which can be helpful for DevOps teams.

#### Git Merge

-   Simple and familiar.
-   Preserves complete history and chronological order.
-   Maintains the context of the branch.
