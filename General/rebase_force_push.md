# git rebase

Rebasing is the process of moving or combining a sequence of commits to a new base commit, like below.

![image](https://wac-cdn.atlassian.com/dam/jcr:4e576671-1b7f-43db-afb5-cf8db8df8e4a/01%20What%20is%20git%20rebase.svg?cdnVersion=31)

rebasing is changing the base of your branch from one commit to another making it appear as if you'd created your branch from a different commit. 

**Unlike merging, rebasing flattens history.** It transfers the completed work from one branch to another. In the process, unwanted history is eliminated.

# Git force push


Git prevents you from overwriting the central repository’s history by refusing push requests when they result in a non-fast-forward merge. So, if the remote history has diverged from your history, you need to pull the remote branch and merge it into your local one, then try pushing again. 

The  `--force`  flag overrides this behavior and makes the remote repository’s branch match your local one, deleting any upstream changes that may have occurred since you last pulled. The only time you should ever need to force push is when you realize that the commits you just shared were not quite right and you fixed them with a  `git commit --amend`  or an interactive rebase. 

However, you must be absolutely certain that none of your teammates have pulled those commits before using the  `--force`  option.

## `--force-with-lease`

`--force-with-lease` is alternative -force option, such a safer option. it will not overwrite any work on the remote branch if more commits were added to the remote branch (by another team-member or coworker or what have you). It ensures you do not overwrite someone elses work by force pushing. I just think of `--force-with-lease` as the option to use when I want to make sure I don't overwrite any teammates code. A lot of teams at my company use `--force-with-lease` as the default option for a fail-safe.
