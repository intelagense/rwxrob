## Thursday, January 28, 2021, 6:36:53PM EST <1611877013>

Learning from Greg that there is a thing called
[LocalStack](https://duck.com/lite?kd=-1&kp=-1&q=LocalStack) technology
that uses containerd tech in order to emulate Amazon cloud services
locally for local development, including the AWS CLI and Cloud Formation
templates.

Obviously, the benefit for developers is amazing but for training and
education is also amazing. You wouldn't run this in production, but for
testing everything *before* you deploy.

## Thursday, January 28, 2021, 10:41:15AM EST <1611848475>

Realized that `kn` Actions can communicate by setting system environment
variables. If they need more than can call one another and connect their
stdin and stdout.

What's even crazier if that with Bash an Action can actually define and
export a function to be used by other subsequent Actions within that
running process. In other words, this allows first-class functions to be
used and passed between Actions.

## Thursday, January 28, 2021, 8:36:10AM EST <1611840970>

It's really becoming clear how valuable the UNIX philosophy is in all
things. Because of a real need for a "`sed` for JSON" we got `jq` which
standardizes a command-line query language for JSON (that frankly blows
away GraphQL). But when combining that simply solution with a solution
for a larger problem related to a query language for all of the
Knowledge Exchange Grid we can how compose URIs that contain implicit
queries drilling down into the data that we need directly.

I'm also sure of something else. 

We don't need no fucking HTML!

There I said it. HTML was a seriously bad design that has damaged *true*
knowledge exchange from the very beginning, but it was the best we had.
These days simplified Markdown and YAML have become the de facto
standards for knowledge capture and exchange. So fuck HTML! 😍

