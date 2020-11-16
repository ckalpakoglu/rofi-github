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

### i3 keybinding
`bindsym $mod+b  exec  --no-startup-id "/usr/local/bin/rofi-github"` 

