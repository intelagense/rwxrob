## Tuesday, April 20, 2021, 11:49:02AM EDT <1618933742>

Today after noticing how fucking long the GitHub repo names are getting
I'm seriously wondering what is the best way to organize labs. I'm
feeling like I should be using a monolith repo instead of a bunch of
tiny ones. In fact, why can't this just go into `rwx.gg` with all the
other content. People are more likely to want to follow a single repo
for changes than a bunch of tiny ones.

Recently it became clear (again) that the simplest approach is just a
`README.md` file that describes the activity rather than a complicated
system for validating, in other words, a Knowledge Node. It can contain
`Dockerfile` and other informational files and starting assets. 

One point of confusion is what to do about labs that involve creation of
GitHub repos and such. Again, the solution is to just say, "go create a
repo someplace."

## Tuesday, April 20, 2021, 11:25:53AM EDT <1618932353>

Realizing how completely valuable Kind is. The simple fact that you can
use local docker tags in place of container registry settings in pod
configurations is worth everything about it. I know that `minikube` does
*not* do this and I suspect that `microk8s` also does not because they
have different goals in mind. Because `kind` was and is being used to
test Kubernetes itself having these iterative conveniences is not only
very nice, but mandatory in a modern development environment.

The more I reflect I realize that I've been thrown completely into the
deepest end of the Kubernetes application development pool. I'm
customizing and extending an existing complicated multi-user application
(JupyterHub) so that is uses network attached storage that it identifies
from a query against the Kubernetes API (Python client) and alters the
core behavior for notebook identification. And to top it off modify the
web front-end to accept all the changes and prompt for user directory
entry, securely. I have a feeling getting the CKAD will be a *lot*
easier later after this project (if I decide to get it).

One thing is for sure, *every* assumption I've made over the last eight
years about what tech is required to learn for a beginner has been spot
on --- except one: Docker.

## Tuesday, April 20, 2021, 8:49:50AM EDT <1618922990>

(I really need to start using the `log` as a log and get the
`cmdtab-zet` done for the other Zettelkasten type of notes.)

Today I want to get the entire collection of JupyterHub components
working and installed as a Kubernetes deployment (since I don't think I
can do the entire thing just as a single pod). 

And I really need to document the entire process because there are
several interdependencies that are not entirely clear even to people who
know about this stuff. I'll keep this work under the repo for
`lab-jupyterhub-nas-ns`
