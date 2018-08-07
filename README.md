# 10chapters

This is a small helper program, for [Professor Grant Horner's Bible-Reading System][system].
In this system you read ten chapters of the bible per day.
Grant Horner divided the bible into ten lists.
You read one chapter from every list every day.

[system]: http://www.sohmer.net/media/professor_grant_horners_bible_reading_system.pdf

This command line tool prints the ten chapters for a given day.
(Currenty in German only.)

## Usage

`10chapters` allows two ways of retrieving the list for a day.
Info about all options can be found calling `10chapters -help`.

1. Using the start date:

   ```
   10chapters -date-started 2018-02-16 [-days-advanced 4] [-days-skipped 3]
   ```

   Will print the chapters for today, when the first list was read on 2018-02-16.
   The default is todays date.
   This parameter can be useful in an alias:
   `alias 10chapters='10chapters -date-started 2018-02-16'`
   Now `10chapters` will always print the correct chapters for today.

   `-days-advanced` and `-days-skipped` can be used to adjust the calculations.
   Increment `-days-advanced` when you read 20 chapters in a day.
   Increment `-days-skipped` when you missed reading the chapters one day.
   This will shift the printed list for the same start date.
   Default is 0 for both.

2. Using the number of the day:

   ```
   10chapters -day n
   ```

   `n` can be any number starting from 1, which is also the default.

Note that `-day` trumps `-date-started`.
`10chapters -date-started 2010-10-11 -day 113`
will always print the list for day 113.

## Example: the ten chapters using start date

```
$ 10chapters -date-started 2018-02-16 -days-advanced 4 -days-skipped 3
Your 10 Chapters for today (day 167):
List 0: Johannes 10 (78/89)
List 1: Deuteronomium 14 (167/187)
List 2: Römer 11 (11/78)
List 3: 1. Johannes 2 (37/65)
List 4: Prediger 1 (43/62)
List 5: Psalm 17 (17/150)
List 6: Sprüche 12 (12/31)
List 7: 1. Chronik 16 (167/249)
List 8: Hesekiel 44 (167/250)
List 9: Apostelgeschichte 27 (27/28)
```

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
