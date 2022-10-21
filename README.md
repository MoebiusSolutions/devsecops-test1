# devsecops-test1

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/MoebiusSolutions/devsecops-test1/tree/main.svg?style=shield)](https://app.circleci.com/pipelines/github/MoebiusSolutions/devsecops-test1)
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/t/moesol-interview/devsecops-test1)

__In this DevSecOps test, interviewees will have to repair the `circleci` pipeline.
This test is intended for a 1-on-1 interview where BOTH interviewer and interviewee
are in the same room or screen sharing. After the interview is completed, the branch
that the candidate push to the repo MUST NOT be merged into `main` and NEEDS to be
deleted. This ensures that the next candidate cannot crib off of previous conducted
tests.__

## Prerequisite

- GitHub Account

## Getting Started

1. Login to GitHub
2. Create a [new branch](https://github.com/MoebiusSolutions/devsecops-test1/branches)
from `main` named: `firstname-lastname-interview`
3. Click `Open in Gitpod` badge/button found in the `README.md`. Gitpod takes a short amount
of time to show the branch created. Refresh the page every couple of seconds.
4. In Gitpod, click `New Workspace` on the newly created branch.
5. Candidate will solve this assessment. Git Commits must be made and pushed for CircleCI
to pick up the changes to run the CICD Pipeline.
6. Click the `circleci` badge/button found in the `README.md`. View the pipeline by filtering
on the branch the cadidate created.

- __As the interviewer, after the interview is completed and the assessment has been graded.
Make sure to delete all [branches](https://github.com/MoebiusSolutions/devsecops-test1/branches)
except for `main`. This ensure branches cannot be cribbed off of.__

## Problems

- Candidate must fix the `.circleci/config.yml` so that it satisfies
the following pipeline criterias:
  - test
  - build
    - (optional bonus) store artifacts
- (optional bonus) Candidate must fix the test(s) that are broken
