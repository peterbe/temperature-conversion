# https://crystal-lang.org/

def c2f(c)
    c * 9.0 / 5 + 32;
end

def is_mirror(a, b)
    massage(a).reverse == massage(b)
end

def massage(n)
    if n < 10
        "0#{n}"
    elsif n >= 100
        massage(n - 100)
    else
        n.to_s
    end
end

def print_conv(c, f)
    puts "#{c}°C ~= #{f}°F"
end

(4...100).step(12).each do |c|
    f = c2f(c)
    if is_mirror(c, f.ceil.to_i)
        print_conv(c, f.ceil.to_i)
    elsif is_mirror(c, f.floor.to_i)
        print_conv(c, f.floor.to_i)
    else
        break
    end
end
