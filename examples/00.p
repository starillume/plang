int x -> 1;
float y -> 2.0;

func sum -> x int, y int: int {
    return x + y * y / x;
}

func inline_sun -> x int, y int: int => x - y;

string aurora_name -> "alice";

sum(x, y);
