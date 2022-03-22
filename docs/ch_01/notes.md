# Chapter 1: Introduction to algorithms

>An *algorithm* is a set of instructions for accomplishing a task.

## Binary search

Binary search compares the target value to the middle element of the array. If they are not equal, the half in which the target cannot lie is eliminated and the search continues on the remaining half, again taking the middle element to compare to the target value, and repeating this until the target value is found. If the search ends with the remaining half being empty, the target is not in the array.

Binary search takes log<sub>2</sub>(*n*) steps. *Simple search* takes *n* steps. So, for a list containing 128 elements, the maximum number of steps would be 7 (log<sub>2</sub>(128) = 7). As an aside, if the number of elements in the list is an exponent with base 2 (i.e. 2^2, 2^3, ... etc.), the maximum number of steps is log<sub>2</sub>(*n*) + 1.

For binary search to work, the list must be sorted.

**Big O Notation**: `O(log(n))`

## Big O Notation

*Big O Notation* is special notation that tells you how fast an algorithm is. It lets you compare the number of operations, not duration. That is, how fast the algorithm grows. Big O Notation is always log base-2.

Common Big O run times:
* `O(log(n))`: log time, ex. binary search
* `O(n)`: linear time, ex. simple search
* `O(n*log(n))`: ex. quicksort
* `O(n^2)`: ex. selection sort
* `O(n!)`: factorial time, ex. traveling salesman

Big O establishes worst-case run time, not best-case run time. Additionally, constants are ignored, as they have no material impact on the growth of the algorithm.
