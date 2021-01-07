---
Title: TIL What FPGAs are and Why They Matter
Subtitle: 'Logic Gate "Softcore" Chips You can Code'
Contrib:
- '[\@mac__gyver](https://twitch.tv/mac__gyver)'
- '[\@piratescipher](https://twitch.tv/piratescipher)'
- '[\@i_hef_corona](https://twitch.tv/i_hef_corona)'
---

Imagine you could design the logic of an electrical device (gates,
etc.). There have been simulators for a while but relatively recently
*Field Programmable Gate Arrays* now provide the ability to not just
simulate the device, but essentially flash it to a chip. Before this
would have required having the chip manufactured. It sorta reminds me of
the functionality provided by virtualization of operating systems but
for low-level chips, even lower than LLVM. You can actually *code* the
electronics of the device. The term "softcore" gets used a lot.

Programmable logic, or LUTs (logical units, AND, OR, XOR) allow the
chaining of logic primitives that are then made into persistent code
that runs as the core logic of the chip. Imagine being able to code your
own chip like an ARM or AMD64 or MicroChip PIC. Not just coding the
operating system that runs on it, but the electronic logic of the chip
itself. It's like an interface on top of a *real* chip, not a virtual
one.

Apparently the process of programming a core is far more cost effective
than programming --- and producing *actual* cores. FPGA are especially
useful for the interface between the world and computer because they can
parallelize and process input like crazy fast. No wonder they are
trending for AI applications.

## See Also

- [Symbiflow Main Site](https://symbiflow.github.io/)
- [Introduction to Symbiflow](https://symbiflow.readthedocs.io/en/latest/introduction.html)
- [Intel AI and FPGAs](https://www.intel.com/content/www/us/en/artificial-intelligence/programmable/solutions.html)
