import "./my_lib.p"

int x -> 1;
float y -> 2.0;

func sum -> x int, y int: int {
    return x + y * y / x;
}

// NOTE: criminal comment
func inline_sum -> x int, y int: int => x - y;

string aurora_name -> "alice";

boolean b -> true;

if b && true || !false != b || 1 >= 2 && -2 <= 1 {
    sum(x, y);
} else {
    inline_sum(x, y);
}
