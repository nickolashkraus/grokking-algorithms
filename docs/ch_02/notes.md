# Chapter 2: Selection sort

## How memory works

Each time you want to store an item in memory, you ask the computer for some space, and it gives you an address where you can store your item. If you want to store multiple items, there are two basic ways to do so: arrays and linked lists.

## Arrays and linked lists

Using an array means all your items are stored contiguously (right next to each other) in memory (note, all elements in an array are the same type). With arrays, if you have reached its capacity and want to add a new item, you need to ask for a different chunk of memory that can fit the array plus the new item. This can be really slow if done many times. This issue is solved by using linked lists.

With linked listed, your items can be anywhere in memory. Each item stores the address of the next item in the list. Linked lists are great if you're going to read all the items one at a time, but if you're going to keep jumping around, linked lists are inefficient, since you must traverse the linked list starting with the first item.

With arrays, however, the address of an item can be computed quickly. Arrays are great if you want to read random elements, because you can look up any element instantly.

### Runtimes

|        | Arrays | Linked lists |
| ------ | ------ | ------------ |
| Read   | O(1)   | O(n)         |
| Insert | O(n)   | O(1)         |
| Delete | O(n)   | O(1)         |

**NOTE**: Insertions and deletions are O(1) time only if you can instantly access the element to be deleted. It's common practice to keep track of the first and last element in a linked list, so it would take only O(1) time to delete them.

### Inserting and deleting

It is easy to insert an element into the middle of a linked list, but for arrays, you have to shift all the rest of the elements down. Additionally, an insertion may fail if the array is out of space. Likewise for deletions: with linked lists, you just need to change what the previous element points to. With arrays, everything needs to be moved up when you delete an element.

### Random vs. sequential access

*Random access* means you can jump directly to the desired element. Arrays are good at random access and search. *Sequential access* means reading the elements one by one, starting at the first element. Linked list can only do sequential access.

## Selection sort

The selection sort algorithm divides the input list into two parts: a sorted sublist of items which is built up from left to right at the front (left) of the list and a sublist of the remaining unsorted items that occupy the rest of the list. Initially, the sorted sublist is empty and the unsorted sublist is the entire input list. The algorithm proceeds by finding the smallest (or largest, depending on sorting order) element in the unsorted sublist, exchanging (swapping) it with the leftmost unsorted element (putting it in sorted order), and moving the sublist boundaries one element to the right.

**Big O Notation**: `O(n^2)`
