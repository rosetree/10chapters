# 10chapters

This is a small helper program, for [Professor Grant Horner's Bible-Reading System][system].
In this system you read ten chapters of the bible per day.
Grant Horner divided the bible into ten lists.
You read one chapter from every list every day.

[system]: http://www.sohmer.net/media/professor_grant_horners_bible_reading_system.pdf

This command line tool prints the ten chapters for a given day.
(Currenty in German only.)

## Usage

```
10chapters [-day n]
```

`n` can be any number starting from 1.
When the option is omitted, the chapters for the first day are printed.

## Example: the ten chapters for day 113

```
$ 10chapters -day 113
List 0: Matthäus 24
List 1: Levitikus 23
List 2: 2. Korinther 3
List 3: Offenbarung 5
List 4: Prediger 9
List 5: Psalm 113
List 6: Sprüche 20
List 7: 1. Könige 9
List 8: Jeremia 47
List 9: Apostelgeschichte 1
```
