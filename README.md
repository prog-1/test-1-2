
# Mid-term test 1 (variant 2)

The test consists of three tasks located in directories `app`, `physics` and
`ping-pong`. Fork the repository, add/edit files in the directories in order to
complete the test. The completed test must be submitted as a GitHub pull request
to `prog-1/test-1-1` repository.

1. Explain what the program in the directory `app` does. Add your explanations
   as comments in the source code. Feel free to add any comments that you think
   are necessary.

2. Write a program in the directory `physics`. The user provides three numbers:
   inductance of three  inductors `L1` `L2` and `L3`. The program must
   calculate and print the total inductance of three inductors connected in
   parallel. The total inductance `L` can be found using the formula `1/L =
   1/L1 + 1/L2 + 1/L3`.

   Example:

   ```
   The program finds the capacity of three inductors connected in parallel.
   Enter inductance of three inductors: 1 2 3
   The total inductance of the three inductors connected in parallel is 0.5454545454545454.
   ```

3. Write a program in the directory `ping-pong` that prints the following for
a provided number:

    - `Ping` if the number is NOT divisible by 2;
    - `Pong` if the number is NOT divisible by 5;
    - `PingPong` if the number is NOT divisible by both 2 and 5;
    - The number itself otherwise.

    Example 1:

    ```
    The program prints Ping, Pong or a number.
    Enter a number: 12
    Pong
    ```

    Example 2:

    ```
    The program pints Ping, Pong or a numbers.
    Enter a number: 10
    10
    ```
