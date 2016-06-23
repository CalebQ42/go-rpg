# go-rpg
Some RPG variables written in Go (Mainly for the Fantasy Age system) to help keep things organized while GMing

# Contents  
Die (A single die)  
Dice (A collection of die)  
AgeDice (Specially made for Fantasy Age system)  
Item (Basic inventory item)  
Shield  
Armor  
Weapon (A little muddled right now due to the fact it needs to be compatible with ranged and melee weapons)  
Stats (Basic framework for base stats)  
Player  
InitPlayer (Player and initiative number struct)  
rpgxui/PlayListAdapter (An array of Player that works as a list adapter for gxui)
rpgxui/InitiativeListAdapter (An array of InitPlayer that works as a list adapter for gxui. Orders from highest initiative to lowest)
rpgxui/ArmorListAdapter (An array of Armor that works as a list adapter for gxui)
rpgxui/WeaponListAdapter (An array of Weapon that works as a list adapter for gxui)


# TODO
Documentation on rpgxui needs more work, but for the most part is done  
Add simple initiative example program (Have it done, just need to add it to the repo)  
   Actually making a program to keep track of a bunch of characters, armor, weapons, and initiative. Still needs a lot of work.  
Add ability to load and save to file easily(rpgxui package does implement this)  
Add age system focuses (or some other type of more universal implementation)
Make more universal  

# UI
As of right now, I currently use gxui for UI, though I may change to github.com/andlabs/ui if/when listboxes are added
