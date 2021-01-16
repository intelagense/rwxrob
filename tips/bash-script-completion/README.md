# Drop Dead Simple Bash Script Tab Completion

Most people don't even know about `complete -C` from the [Programmable
Completion in the Bash Man
page](https://duck.com/lite?kd=-1&kp=-1&q=Programmable+Completion+in+the+Bash+Man+page).
Now you do. Here's a template of a Bash script you can use to get
started. Just create a function for each of your subcommands and add
them to the list of commands. Then add `complete -C myscript myscript`
to your `.bashrc` (or whatever). Note, however, that this only words for
Bash. ([Zsh is stupid, don't use it.](https://rwx.gg/advice/dont/zsh))

```bash
#!/usr/bin/bash

foo () {
  local first=$1
  echo "Would foo with $first."
}

############################## Main Command  #############################

subcommand="${1-usage}"
shift 2>/dev/null

if [ -z "$subcommand" ]; then
    usage
    exit 1
fi

commands () {
    echo ${commands[@]}
}

declare -a commands=(rm commands)

for i in ${commands[@]}; do
    if [[ "$i" == "$subcommand" ]]; then
        "$subcommand" "$@"
        exit 0
    fi
done

######################### Tab Completion Context ########################

# remember COMP_LINE is changed by shift
if [ -n "$COMP_LINE" ]; then
    for cmd in ${commands[@]}; do 
        [[ "${cmd:0:${#1}}" == "$1" ]] && echo $cmd
    done
    exit 0
fi
```

