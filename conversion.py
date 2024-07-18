import math


def c2f(c):
    return c * 9 / 5 + 32


def is_mirror(a, b):
    def massage(n):
        if n < 10:
            return f"0{n}"
        elif n >= 100:
            return massage(n - 100)
        else:
            return str(n)

    return massage(a)[::-1] == massage(b)


def print_conv(c, f):
    print(f"{c}°C ~= {f}°F")


for c in range(4, 100, 12):
    f = c2f(c)
    if is_mirror(c, math.ceil(f)):
        print_conv(c, math.ceil(f))
    elif is_mirror(c, math.floor(f)):
        print_conv(c, math.floor(f))
    else:
        break
