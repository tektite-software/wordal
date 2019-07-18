# Wordal
You've tried octal, you've tried hexadecimal, now try wordal!  It's not a Pokémon, it's a word-based numeral system with a base of 65,536.  In other words, Wordal is a dictionary containing precisely 2<sup>16</sup> words, allowing any arbitrary data to be expressed with a memorable English word for every 16 bits.

## Implementations

#### Official
* [Golang](https://github.com/tektite-software/wordal)

#### Community

If you would like to have your implementation featured here, feel free to send a PR!

#### [Wordal Master Dictionary](https://github.com/tektite-software/wordal-dict)

## Features

Wordal's primary intended use case is to efficiently communicate data using memorable phrases or speech.  As such, it was created with the following key features:

* 2<sup>16</sup> word entries, allowing 16 bits per word or 32 bits for two-word phrases (over 4 billion unique two-word combinations)
* Emphasizes memorable and common words, avoids overly technical, archaic, or scientific words
* Avoids homophones and pseudo-homophones (about 2500 were removed from the source dictionaries)
* No innapropriate words (removes Google's "Bad Words" list)
* Favors American spelling; British, Canadian, and Australian spellings and colloquialisms have been removed as much as possible (e.g., `check` instead of `cheque`)
* No words shorter than 3 letters or longer than 15, prefer shorter words
* Average word length is about 8 letters

## Acknowledgements

#### Source Dictionaries
* [Alan Beale's 12dicts](http://wordlist.aspell.net/12dicts/) lists were used at the primary word sources of Wordal.
* Hunspell dictionaries were used to remove regional variants (en_AU, etc).

#### Homophone Sources
  * [Phillip M. Feldman et. al., "List of English Homophones"](http://phillipmfeldman.org/English/Homophones.html)
  * [Ian Miller, "English Homophones"](http://www.singularis.ltd.uk/bifroest/misc/homophones-list.html)
  * [Evan Antworth, "American Homophones"](http://members.peak.org/~jeremy/dictionaryclassic/chapters/homophones.php)

## License

Copyright (c) Tektite Software, LLC

The Wordal Dictionary and all official implementations are liscenced under the MIT License.