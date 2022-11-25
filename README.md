# Yall

Yall stands for "Yet another Lox language" implementation, referring to the Lox language in ["Crafting Interpreters"](https://github.com/munificent/craftinginterpreters), by Bob Nystrom.

## Changes to the Interpreter

- REPL exits when input line is `exit`.
- Supports UTF-8 source code.
- Adds fancier error printing (like in the book).
- Adds Newline token to partially preserve newlines for code formatter.
