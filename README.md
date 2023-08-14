# gonotes
**gonotes** is simple cli program to manage notes or in table with class, date and id 

## library used 
1. cobra
2. sqlite driver

## install 
you dont need to install any dependency\
just run **go install**

## use
in gonotes you have four subcommands: 
1. initdb
2. write
3. read
4. remove

### initdb 
    - before you start write some notes you need to initialization database with:\
        **gonotes initdb**
    - you can find the database in ~/.cache/gonotes/ with name gonotes.db
### write
    to write new notes with class, date and id 
        - write command has four flags :
            1. --content "set the content of notes if you leave it empty the programe will exit"
            2. --class, -c  "set a class for note default value is empty"
            2. --id "       "set a id for note default value is empty"
            4. --date -d "set date for note default current date with format YYYY:MM:DD"
### read 
    to read note from database 
    - read has three flags:
        1. --class -c "set class for note to search at datebase default empty
        2. --id "set id for note to search at database default empty
        3. --date -d "set date for note to search at datebase default empty
    - if you run gonotes read without class,id or date gonotes will return all notes in database
### remove
    to remove note from database 
    - to remove the note we use non(number of note) its autoincrement column its set auto when you add new note to database
    - remove command just has one flag :
        1. -n "set non that you want remove it"
## todo
1. edit command 
2. multi remove for remove command

    



