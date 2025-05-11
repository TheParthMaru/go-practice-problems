# Questions on pointers

## Easy

1. Declare an integer variable x with the value 10. Create a pointer ptr that stores the address of x. Print the following:

- The value of x.
- The memory address of x (using the pointer).
- The value stored at the memory address using dereferencing.

2. Write a function named square that accepts a pointer to an integer. Inside the function, modify the original value by squaring it (i.e., multiply it by itself). In main(), declare a variable x = 5, pass its address to the square function, and print the updated value of x.

3. Write a function named swap that accepts two `*int` parameters. The function should swap the values of the two variables pointed to. In main(), define two integer variables, print their values, call swap, and print their values again to confirm they’ve been exchanged.

4. Write a function resetSlice that accepts a pointer to a slice of integers `*[]int` and sets it to an empty slice. In main(), define a slice with some values, pass its pointer to resetSlice, and print it after to confirm it's empty.

5. Write a function maxMin that accepts two `*int` arguments and returns two int values: the larger and the smaller of the two. In main(), define two variables, call maxMin with their pointers, and print the result.
