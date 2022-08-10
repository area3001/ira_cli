# IRA go CLI

# Getting Started
```shell
# set pixel length of specific device
./goira config set <dev> pixel_length 150

# enable fx mode for a specific device
./goira fx enable <dev>

# list available fx's
./goira fx list

# set specific fx for a specific device (colors are hex3 (rgb) or hex 6 (rrggbb)
./goira fx set <dev> 2 fg=f0f bg=0f0

# ask devices to send a hearthbeat (best to do before listing)
./goira devices sync

# list devices
./goira devices list

# reset a device
./goira devices reset

# set the name of a device
./goira config name <dev> <name>
```