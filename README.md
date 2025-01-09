<h1 align="center"> ✨ actions-io ✨ </h1>
<p align="center">A web application that displays each input and output of a GitHub Action</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/3be9e1a3-5fcd-42a9-8922-f81c096f4ad5" width=320 /><img src="https://github.com/user-attachments/assets/d6c521e2-75d5-433d-8a20-27b9495f19cd" width=320 />
</p>

## Why?

* **No standard**: There is no standard way to display GitHub Action inputs and outputs.
* **Maintenance effort**: Maintaining them in the README file requires effort and can be forgotten.
* **Going back in time**: Viewing previous inputs and outputs requires checking out a branch on GitHub.

## How?

I try to fix the metnioned problems with this approach.
Just visit https://actions-io.vercel.ap/[OWNER]/[REPO] to see all input and output of the action.
The site can be used as you use your actions in your workflow yml files.
This means, it can also handle specific versions (git tags).
Just add `@[ACTION_VERSION]` to the `[REPO]` URL to see that specific version.

## Badge

The site supports also displaying a badge.
Just start with `/badge` in the URL path like `https://actions-io.vercel.app/badge/ioki-mobility/summaraizer-action@main`.
This will be rendered to: <img src="https://actions-io.vercel.app/badge/ioki-mobility/summaraizer-action@main" />

## Examples

* [StefMa/Upgrade-Go-Action](https://actions-io.vercel.app/StefMa/Upgrade-Go-Action) [![Actions I/O](https://actions-io.vercel.app/badge/StefMa/Upgrade-Go-Action@main)](https://actions-io.vercel.app/StefMa/Upgrade-Go-Action@main)
* [ioki-mobility/summaraizer-action](https://actions-io.vercel.app/ioki-mobility/summaraizer-action) [![Actions I/O](https://actions-io.vercel.app/badge/ioki-mobility/summaraizer-action@main)](https://actions-io.vercel.app/ioki-mobility/summaraizer-action@main)
* [actions/setup-go@v2](https://actions-io.vercel.app/actions/setup-go@v5) [![Actions I/O](https://actions-io.vercel.app/badge/actions/setup-go@v5)](https://actions-io.vercel.app/actions/setup-go@v2)
* [docker/build-push-action@v6.10.0](https://actions-io.vercel.app/docker/build-push-action@v6.10.0) [![Actions I/O](https://actions-io.vercel.app/badge/docker/build-push-action@v6.10.0)](https://actions-io.vercel.app/docker/build-push-action@v6.10.0)

## Background

Have you ever worked with GitHub Actions?
Do you know all of their inputs and outputs?
If so, lucky you. Probably the maintainer did a good job by document them in the README file.
But, this is not always the case.
Also, there is no standard way to do this.
Lets look at some examples:

[actions/setup-go@v5.2.0](https://github.com/actions/setup-go/tree/v5.2.0)</br>
This action doesn't display all possible IO in a unified way.
It explains them, but spread accross the README. 

[docker/build-push-action@v6.10.0](https://github.com/docker/build-push-action/tree/v6.10.0)
This action does a pretty good job.
It has a nice table with the explanation of all their IO.
*But*, do you see which of them are required and which not?
Can you directly spot the defaults for each?

[ioki-mobility/summaraizer-action@0aa210e3dd7513fae6c40a55d8a304592554da0c](https://github.com/ioki-mobility/summaraizer-action/tree/0aa210e3dd7513fae6c40a55d8a304592554da0c)
This action uses a different style to display its IO than the docker action.
Seems also valid.
*But*, did you spot that the input `summaraizer-version` is not documented?

There is no standard way to display action IO in actions.
Its like wild west.

Also, even if a maintainer is trying to do a good job, he/she might simply forget to document it on the README.
We are humans, that happens.

I try to solve this with two things:
A website and a badge.

The website should aim as the defacto standard of checking action input and outputs.
The badge acts as the entry point for the website, as maintainer can add them to their README
to always have an up-to-date information about the IO.

I encurage all readers of this to make my dream happen.
Spread `actions-io` to the world and add it to your GitHub Action **today**.


