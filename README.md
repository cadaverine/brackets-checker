# brackets-checker

## Скобки в коде

Проверить, правильно ли расставлены скобки в данном коде.

**Вход.** Исходный код программы.<br>
**Выход.** Проверить, верно ли расставлены скобки. Если нет, выдать индекс первой ошибки.

**Формат входа.** Строка s[1 . . . n], состоящая из заглавных и прописных
букв латинского алфавита, цифр, знаков препинания и скобок
из множества []{}().<br>
**Формат выхода.** Если скобки в s расставлены правильно, выведите
строку “Success". В противном случае выведите индекс (используя
индексацию с единицы) первой закрывающей скобки, для
которой нет соответствующей открывающей. Если такой нет,
выведите индекс первой открывающей скобки, для которой нет
соответствующей закрывающей.

<hr>

**Tests results:**
```
=== RUN   TestFindWrongBrackets
--- PASS: TestFindWrongBrackets (0.00s)
PASS
ok      _/home/vlad/programs/go/brackets-checker        0.001s
```