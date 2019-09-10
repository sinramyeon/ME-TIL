1. Install A Gitlab CI on Mac


Download the binary

`sudo curl --output /usr/local/bin/gitlab-runner https://gitlab-runner-downloads.s3.amazonaws.com/latest/binaries/gitlab-runner-darwin-amd64
Give it permissions to execute:

sudo chmod +x /usr/local/bin/gitlab-runner
2.Setup A Gitlab Ci on Gitlab.com


Register Runner

gitlab-runner register
Enter your GitLab instance URL(Can FInd From https://gitlab.com/basking/graphql-api/settings/ci_cd)

Please enter the gitlab-ci coordinator URL (e.g. https://gitlab.com )
https://gitlab.com
Enter the token you obtained to register the Runner:
Please enter the gitlab-ci token for this runner
tokentokenbasking
Enter a description for the Runner(you can change this later in GitLab’s UI)
Please enter the gitlab-ci description for this runner [hostame]
my-runner
Enter the tags associated with the Runner (you can change this later in GitLab’s UI)
Please enter the gitlab-ci tags for this runner (comma separated):
graphql


5. Install the Runner as service and start it:

cd ~ gitlab-runner install gitlab-runner start


6. Reboot a system & check status

gitlab-runner status


results :

Runtime platform arch=amd64 os=darwin pid=87172 revision=8bb608ff version=11.7.0
gitlab-runner: Service is running!


(reference) https://docs.gitlab.com/runner/install/osx.html

https://allroundplaying.tistory.com/21

From https://gitlab.com/basking/graphql-api/settings/ci_cd , now we can find running CI we made.



3.Test Fail / Success / On Merge Request


once the runner started, we can check our test(or build) is failed/ or succeed in CI(https://gitlab.com/basking/graphql-api/-/jobs)





also if the test failed and we made a merge request, it looks like







and also the commiter receives a todo





we can check which erros occurs in script via (https://gitlab.com/basking/graphql-api/-/jobs)





4.gitlab-ci.yml

```
(TODO)
```


5.Automated Test / Build / Deploy


We currently have 3 CI stages, 1. Build, 2. Test, 3. Deploy. 

since we have this line below in our Deploy job,

only:
    refs: #it will run only in master branch
      - dev
      - master


the deploy scripts will only run in the dev or master branch.



