# TODO List

* Unicode (go) convenience classes in PEGN
* Create a Recording YAML to YouTube converter 
* Change `topic` to use log recording YAML files
* Change `KNPATH` to `KN` everywhere
* Finish the KEG Recording YAML schema
* Universal bot for YouTube, Discord, and Twitch chat.
* Add ability to `select` from `lasturl` 
* Change chat refs to chat.rwx.gg (from rwxrob.live)
* Finish the `gh orgs` alias
* Fix vimrc perl default link replacement to use `urlencode`
* Fix my `sched` command
* Need to create an emoji keyword filter (see log for jan 23) 
* Add a follow section to GitHub profile (RSS, etc.)
* Pull environment info in `gl` out into `.env`
* Simple 00:00:00 converter to seconds for date math
* Shave top margin off of first menu item on rwxrob.live

## AFKWorks Questions for Discussion

* Do we assume a single KEG base per account?
* Or do we allow multiple context/workspaces per account?
* How often would a knowledge worker shift between roles/accounts?
* Would just environment vars cover it for a single account?

## Videos That *Must* Be Made

* Difference between `<<<` and `< <(`
* Everything about `date -d` 
* YAML and JSON from Bash Shell with `yq`/`jq`

## Viewer Video Requests

Yeah, yeah, I'll get to it eventually.

1. Command Line Searches with DuckDuck.go
1. TMUX Screen-centric Configuration 
1. Setting Up Fish / Asciiquarium
1. Review Vim Plugins

## The Rest

* Need to submit change merge request to https://golang.org/pkg/testing/
  to update the docs to say "Output:" or "Unordered output:" instead
  since it isn't covered anywhere else.

## Movie Clips to Add

* "Stay on target."
* "It's working! It's working!"
* "2 years later."
* "Now that's a name I've not heard in a long, long."

## Urgent

* give people their own color
* flesh out the notion of a YAML task list leveraging references
* create a prompt that keeps it under max length chopping at last /
* make `twitch head` (twitch top)
* make `twitch tail`
* fix save on new repos
* add Twitch to `ska`
* `lsof` need to read up on everything
* twitch shell script with raid/vip/etc.
* add *Bash/POSIX Shell Live Coding*
* Running glossary of security terms (CTF)
* make a welcome video for youtube that says "find me on twitch"
* make sat like sunday (send email invites)
* cleanup up gitlab.com/rwxrob/yaml-vim
* Create oscp project in gitlab for collaborative collection of *legal* resources
* CafeSec
* Disable all VODs, use OBS to make vids for posting.
* Make robs.io public
* Convert WeeChat stuff to use `/secure` command so it gets saved as encrypted.
* Carefully vet and test to enable WeeChat config to be public without the secure files. There is so much configuration to do creating a boilerplate with vim-centric bindings and such is worth it.

## Blog Articles to Write

* Coco Chanel Rule of Software Design
* Recursive Functions in Golang
  * ParentWith(".kn",".")
* Just in Time Learning
* How Can I Get Started?
  1. Consider Free Book: Learning the Linux Command Line
  1. Learn to Touch Type
     1. Learn Your Keyboard
  1. Get Linux
     1. Virtual Machine
     1. Install
  1. Learn to Survive the Command Line
     1. Navigation: `ls`, `cd`, `pwd`, `clear`
     1. Files: `mkdir`, `rm`, `rmdir` 
     1. SSH
  1. Learn Vi/m
  1. Learn to Code in Bash
     1. Write Simple Scripts
        1. `touch`, `chown`
     1. Configure Your Bash Shell
  1. Learn TMUX
* Bug Bounty Programs are Economic Future Reality
* Tips for Resumes and LinkedIn
  1. Be specific
  1. Include measurable criteria
  1. Purpose is to Troll Potential Employer
  1. Do *Not* Lie, But Don't Understate Either
  1. Cater to Your Potential Employer
* My Terminal Linux Setup
* End of Salaried Employment?
  * Salaried
  * Hourly (Project)
  * Gaming the System
* How I Caught My Linux Email Server Hackers
* Minimum System for Virtual Machine Hosting
* VPS vs Local Linux for Learning
  1. Cost
  1. Setup
  1. Security Considerations
  1. Interfacing with Devices
  1. What Comes Next?
* Why are You Picking That Tech?
  1. Get a job.
  1. It's the best. 
  1. Get chicks.
  1. Other reasons.
* How I Became a Massive GraphQL Fan
* What the Fuck is "Server-Side Rendering"?
* Reasons Rust Will Gain Adoption
* Benefits of Full Screen Context Switching
  1. Your Window Manager Doesn't Matter
  1. Maximize Effective Works Flows
  1. Stop Assuming a Desktop Environment
* Benefit of Search-centric Workflows

## FAQs That I Need TODO

1. "Notation as a Tool of Thought" ... by Iverson.
1. Add sound effects to the stream.
1. Look at remaking `ytop` in Rust as exercise.
1. What OS do you use? 
1. What language should I start with?
    1. What are your favorite programming languages and why?
    1. What order should I learn which language?
1. How did you get your screen like that?
1. Where can I get your config (dotfiles)?
1. Why choose Gitlab over Github?
1. Why certify?
1. Which certifications should I get?
    1. What order should I get them?

## Eventually

* add a random quote terminal screen save to go with `back` cmd
* Add full `date -d` support to a Go date package.

* Submit `FileContains()` to Go testity
* Watch Hidden Figures
* magic vi script to automatically include all *.{jpg|png} as Markdown links
* add autocomplete to my `twitch`
* look at pipes.sh for terminal stuff
* Pirates of Silicon Valley
* fix `tmux.conf` to allow defaults as well as screen-based
* Check out `axel` `wget` replacement.
* Check out <http://zed.0xff.me/>
* Upgrade `ix` to contain http://ix.io/client-ish stuff
* Research Srinivasa Ramanujan
* Go Projects:
    * Static Site Generator with Built-in Search and Pandoc
    * Virtual Agent - Twitch IRC, Twitter, Discord (CAVA)
    * `dtime` Package
    * `cmdtab` Package

* List of ways to describe regular family/neighbor tech support suitable for listing on a resume.

## Soon

* temporary email address 10minutemail.com 
* Understand `$CDPATH`
* use `chattr` for changing whole filesystem attributes
* look into https://developer.nvidia.com/video-encode-decode-gpu-support-matrix
* https://pwnable.tw/
* AFL for fuzzing
* ASAN for forensics and reverse engineering
* look into NixOS
* `syncthing` instead of `save`?
* `git clean -xdf`
* `cd $(mktemp -d)`
* Disable [email protected] in lynx
* "suckless terminal" look into it
* https://cyberspace.dev/
* https://systemoverlord.com/2018/02/14/preparing-for-penetration-testing-with-kali-linux.html
* "ASM is too high level" Def Con talk
* https://www.fuzzingbook.org
* study cve.mitre.org and write exploits "all you need"
* learn about Google P0 (https://googleprojectzero.blogspot.com)
*  idea for your twitch bot: https://www.youtube.com/watch?v=QIacthT6c84 
* have weechat pick up on status changes
* https://chmod-calculator.com/
* lookup meute for hacker music
* follow ytcracker
* learn linux.die.net
* get rid of the fucking waisted stuff at bottom of page on Lynx (or switch links) 
* find out if other people post CTF play-through on YT
* suggest twitch start a Cybersecurity CTF category
* tty-share.com
* 5151 rising sun song
* https://en.wikipedia.org/wiki/Maltego
* study yaml code execution vulnerability
* SMBv3 vuln (https://twitter.com/malwrhunterteam/status/1237438376032251904)
* AKG K52 -> Audio Technica M50x
* research SS-1 or CloudLifter or FetHead
* add completion to l
* research odin project
* yet another oscp video to review https://youtu.be/T1AUCXXKzL8
* pwnable.xyz
* pwnagotchi pssh
* TIL pssh
* fix `ix` to /dev/null output
* apt install ncdu && ncdu
* xmonad
* Art of Exploitation
* I learned about Ghidra, the released RE tool from the NSA
* check out FlareOn
* https://github.com/RPISEC/MBE
* https://blog.ret2.io/2018/09/11/scalable-security-education/
* look into Rensselaer Polytechnic Institute (binary exploitation on github)
* read about <http://ctftime.org>
* add a colon after the nicks in the weechat output window
* Get a resources and references list started for pentesting and OSCP (perhaps an "awesome" list, but without the stupid "awesome" in the name).
- video "Safeguard against disconnects (keeping a terminal open, tmux, etc)."
  - upcoming videos list
  - capture requirements for "continuous learning"
  - video "Which certs for security?"
  - research "quanta" (linux v.s. nt kernel scheduling)
  - video "Set Up Automated Google Alerts"
  - video "Creating a free web site with Netlify and GitLab"
  - configure mutt, protonmailbridge and start using again
  - video "Why College?"
  - video "Which College?"
  - video "FaunaDB"
  - video "Netcat (NC)"
  - research Discord voice issue
  - setup gpg signing of commits (again)
  - channel quotes database
  - paths to fruit
  - netlify video
  - jamstack video
  - fix `repo clone` not creating proper path
  - do video about gpl2 vs gpl3
  - do video about GitLab vs GitHub
  - do video about open source v.s. fsf
  - do video summarize all the licenses
  - temporary email addresses
  - research twit.tv
  - research dr. jordon peterson
  - research alan watts
  - research victorian era students teaching students
  - research cgit as alternative to gitea
  - research cview
  - video what's wrong with github PR process and contributions
  - https://fundingcircle.github.io/fc4-framework/
  - change google alerts to once a week
  - find out amazon's position on gplv3 with regard to Alexa (for example)
  - update the fortune quotes to our own
  - research kata and it's benefits v.s docker
  - research modern changes to KVM vs virtualbox vs vmware
  - research SourceHut
  - research zeit.co
  - research MISRA and the Toyota case
  - video creating google alerts
  - code an easy way to rebase/reset a git repo
  - research http://kadgar.net/

News:
  - https://www.linkedin.com/feed/update/urn:li:activity:6635664931192655872/

Ideas:
  - duckduckgo why use it?
  - why i don't like conversational ui
  - how processes work on UNIX/Linux
  - how do i get my new bashrc or variables to work
    - what does export do
  - fg/bg jobs what are they (and why i don't use them)
  - streaming
    - audio setup
  - init states
  - "Which terminal file manager should I use?"
  - Become a carpenter, not a hammer user.
  - make a `repo rename` command function

Later:
- create lists in Twitter
- create a Twitter bot/agent (rwxagent)
- put macOS back on trash can
- figure out how to be notified if a twitch viewer has a question
- finish config for bash
- port all medium blogs to main
- port old blogs from github to robs.io
- create robs.io
- search browser history from the command line
- look into ultrasnips and youcompleteme
- filter all the entries in skilstak.io to robs.io that belong there
- fix functions.bash exports
- create vimscript for my bash extensions
- make `save` use gh or gitlab to commit
- set REPOS_DIR
- tree fully in bash
- mail boys computers
- pilfer oh-my-bash
- move bash config to github
- fix licenses
- rw build
- pego (consider using pidgeon instead)
- dtime

Long Term:
- make a tripwire clone that is multiproc and in golang

* no `function` keyword, at all
* only arrow functions and class methods
* no function hoisting, at all
* no redundant `name:name`
* no use of `lodash` at all
* single quotes for strings
* two space indent
* no parens for single parameter functions
* no `return` for single line functions
* no curlies for single line blocks
* no edge cases that break semicolon omission
* absolutely no use of `var`
* forced declaration with `let` or `const`
* no semicolons, ever
* no `+` joins, only template literals
* go rules for math operators
* no use of `this` unless in a class
* forced `})` when it occurs (no different lines)

## Left Overs

* ezmark linter with emoji conversion
* make a vim-adventures.com killer for free in Phaser3
# create a discline installer in `home-config`

# GET OFF GITHUB

# search from the command line medium blog post

# eliminate any reference to anything that isn't priority for empowerment 

# write a `send` bash function to send files to email addresses or aliases

# soil bash completion restored

# add `config.bash` to home-config sconsify to use the .bash_private token

# add a circle with a tux in it to skilstak.io

# hang the nets

# add autolinker for `edit` SOIL titles

# cancel adobe membership for me

# make `class` into a function with completion

# migrate adobe membership for gordon to new email

# move offlineimap to cron

# update all packages on laptop

# create a `curl`-able search interface to skilstak.io

# setup talkwalker instead of google alerts

### Major Initiatives

* SkilStak (skilstak.{io,sh})
* KnowledgeNet (know.sh)
* Urban Adventure
* RetroPi Arcade




