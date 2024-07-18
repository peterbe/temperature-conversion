fn c2f(c: i8) -> f32 {
    let c = c as f32;
    c * 9.0 / 5.0 + 32.0
}

fn is_mirror(a: i8, b: i8) -> bool {
    let a = massage(a);
    let b = reverse_string(massage(b));
    a == b
}

fn massage(n: i8) -> String {
    if n < 10 {
        return format!("0{}", n);
    } else if n >= 100 {
        return massage(n - 100);
    } else {
        return format!("{}", n);
    }
}

fn reverse_string(s: String) -> String {
    s.chars().rev().collect()
}

fn print_conversion(c: i8, f: i8) {
    println!("{}°C ~= {}°F", c, f);
}

fn main() {
    let mut c = 4;
    while c < 100 {
        let f = c2f(c);
        if is_mirror(c, f.ceil() as i8) {
            print_conversion(c, f.ceil() as i8)
        } else if is_mirror(c, f.floor() as i8) {
            print_conversion(c, f.floor() as i8)
        } else {
            break;
        }
        c += 12;
    }
}
