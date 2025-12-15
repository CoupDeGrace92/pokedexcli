Uses PokeAPI to create a working CLI pokedex.  Currently supported commands:
    -Exit: Exits the program
    -Catch {...pokemon}: Attempts to catch all the listed pokemon in order
    -Map: Displays the name of 20 locations starting at the current location.  Initializes to the 1 location
    -Mapb: Displays the previous 20 locations.  If no previous location, no location is printed
    -help {...commands}: Lists all commands and uses, lists only the listed commands if list is given
    -catch {...pokemon}: Attempts to catch pokemon in the list
    -inspect {...pokemon}: Lists attributes of the listed pokemon if they are caught
    -pokedex: Lists all pokemon that have been caught
    -explore {...location}: Takes either names or id numbers, lists pokemon catchable at that location

pokemon or location variables can have names or ids as the identifier and work since the pokeApi endpoints take either

The cache helps improve searching performance by not requiring multiple calls.  Both the pokemon calls and location calls
are cached.

Catch probability - determined by BaseExperience (experience gained by defeating in battle). I normalized it and gave
it a buffer probability on the low and high end to ensure no perfect catch rates and no miniscule catch rates.

POSSIBLE PLANS FOR EXTENSION:
    -Save pokemon state (needs some sort of savable database layer)
    -Sortable pokedex/searchable pokedex for attributes/types etc.
    -Simplified combat
    -Random Encounters 
    -Map location and movement via graph instead of batch exploring