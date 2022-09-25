# gate

gate is a learning project about logic gates and is inspired by the wonderful book: The Manga Guide to Microprocessors

Built with Aloha in Hawaii 🌊

Thank you to the Hawaii Public Library System for making such an amazing book available for free!

```bash
import "github.com/mbarley333/gate"
```

# testing
This library uses [gotestdox](https://github.com/bitfield/gotestdox) to create sentences in the testing output.

```bash
❯ gotestdox
gate:
 ✔ AND output is false when input a equals false and input b equals false (0.00s)
 ✔ AND output is false when input a equals false and input b equals true (0.00s)
 ✔ AND output is false when input a equals true and input b equals false (0.00s)
 ✔ AND output is true when input a equals true and input b equals true (0.00s)
 ✔ NAND output is false when input a equals true and input b equals true (0.00s)
 ✔ NAND output is true when input a equals false and input b equals false (0.00s)
 ✔ NAND output is true when input a equals false and input b equals true (0.00s)
 ✔ NAND output is true when input a equals true and input b equals false (0.00s)
 ✔ NOR output is false when input a equals false and input b equals true (0.00s)
 ✔ NOR output is false when input a equals true and input b equals false (0.00s)
 ✔ NOR output is false when input a equals true and input b equals true (0.00s)
 ✔ NOR output is true when input a equals false and input b equals false (0.00s)
 ✔ NOT output is false when input a equals true (0.00s)
 ✔ NOT output is true when input a equals false (0.00s)
 ✔ OR output is false when input a and input b equals true (0.00s)
 ✔ OR output is true when input a and input b equals false (0.00s)
 ✔ OR output is true when input a equals false and input b equals true (0.00s)
 ✔ OR output is true when input a equals true and input b equals false (0.00s)
 ✔ XOR output is false when input a equals false and input b equals false (0.00s)
 ✔ XOR output is false when input a equals true and input b equals true (0.00s)
 ✔ XOR output is true when input a equals false and input b equals true (0.00s)
 ✔ XOR output is true when input a equals true and input b equals false (0.00s)
```
