# functions

Solutions to the exercises set here: https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/4077556?start=60

1. Write a function that takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that lets us know whether or not the argument was even. For example:

    - `half(1)` returns `(0, false)` 
    - `half(2)` returns `(1, true)`

2. Modify the previous exercise to use a func expression.

3. Write a function with one variadic parameter that finds the greatest number in a list of numbers.

4. What's the value of the expression: `(true && false) || (false && true) || !(false && false)`

5. Write a function, `foo` which can be called in these ways:

```golang
  func main() {
    foo(1, 2)
    foo(1, 2, 3)
    aSlice := []int{1, 2, 3, 4}
    foo(aSlice...)
    foo()
  }
```

6. Find a problem at projecteuler.net then create a solution. Add a comment beneath your solution that includes a description of the problem. Upload your solution to Github. I chose https://projecteuler.net/problem=9
