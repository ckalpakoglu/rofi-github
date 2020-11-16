# rofi-github
A program that generates a rofi menu to access easily to most used github repositories.

It uses a text file as a history. If the input is not in the cache(history) retrieves projects from github and adds selection to the cache for future use.

Program uses [github-cli's](https://github.com/cli/cli) `hosts.yaml` file for details (oauth token)

### Installation
Add precompiled binary to your `$PATH` or: 
1. `git clone https://github.com/ckalpakoglu/rofi-github`
2. `cd rofi-github`
3. `make` 
4. `cd _release` and `mv rofi-github /usr/local/bin` 


### Arguments
```
NAME:
   rofi-github - A new cli application

USAGE:
   rofi-github [global options] command [command options] [arguments...]

DESCRIPTION:
   A Rofi plugin to access github repositories quickly. It does not requires configuration, uses github cli (gh) oauth token

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --browser value, -b value  open in browser (default: "google-chrome-stable")
   --conf value, -c value     config location (default: "/your/home/.config/gh/hosts.yml")
   --cache value              cache file location (default: "/your/home/.config/gh/rofi-github.cache")
   --org value, -o value      organization
   --version, -v              prints version (default: false)
   --help, -h                 show help (default: false)
```

### i3 keybinding
`bindsym $mod+b  exec  --no-startup-id "/usr/local/bin/rofi-github"` 

<div align="center">
<h3>See In Action</h3>
<img src="https://github.com/ckalpakoglu/rofi-github/blob/master/.meta/rofi-github.gif">
