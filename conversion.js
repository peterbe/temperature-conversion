function c2f(c) {
  return (c * 9) / 5 + 32;
}

function isMirror(a, b) {
  function massage(n) {
    if (n < 10) return `0${n}`;
    else if (n >= 100) return massage(n - 100);
    return `${n}`;
  }
  return reverseString(massage(a)) === massage(b);
}

function reverseString(str) {
  return str.split("").reverse().join("");
}

function printConversion(c, f) {
  console.log(`${c}°C ~= ${f}°F`);
}

for (let c = 4; c < 100; c += 12) {
  const f = c2f(c);
  if (isMirror(c, Math.ceil(f))) {
    printConversion(c, Math.ceil(f));
  } else if (isMirror(c, Math.floor(f))) {
    printConversion(c, Math.floor(f));
  } else {
    break;
  }
}
