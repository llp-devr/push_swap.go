# push_swap.go

A GO implementation of the push_swap project (version 8.3) from 42 School.

## Introduction

The Push swap project is a simple yet highly structured algorithmic challenge:
you need to sort data.

You have at your disposal a set of integer values, 2 stacks, and a set of
instructions to manipulate both stacks.

Your goal? Write a C program called push_swap that calculates and displays the
shortest sequence of Push_swap instructions needed to sort the given integers.

## Mandatory part

### The rules

- You have **two stacks** named `a` and `b`.

- At the beginning:

  - Stack `a` contains a **random number of unique** negative and/or positive
    integers.
  - Stack `b` is **empty**.

- The goal is to **sort the numbers in stack `a` in ascending order** using the
  available operations.

#### Available Operations

##### Swap Operations

- **`sa` (swap a)**: Swap the first **two elements** at the top of stack `a`.  
  _(Does nothing if `a` has only one or no elements.)_
- **`sb` (swap b)**: Swap the first **two elements** at the top of stack `b`.  
  _(Does nothing if `b` has only one or no elements.)_
- **`ss`**: Perform `sa` and `sb` **simultaneously**.

##### Push Operations

- **`pa` (push a)**: Take the **top element** of `b` and put it at the
  **top of `a`**.  
  _(Does nothing if `b` is empty.)_
- **`pb` (push b)**: Take the **top element** of `a` and put it at the
  **top of `b`**.  
  _(Does nothing if `a` is empty.)_

##### Rotate Operations (Shift Up)

- **`ra` (rotate a)**: Shift **all elements of `a` up** by one.  
  _(The first element becomes the last one.)_
- **`rb` (rotate b)**: Shift **all elements of `b` up** by one.  
  _(The first element becomes the last one.)_
- **`rr`**: Perform `ra` and `rb` **simultaneously**.

##### Reverse Rotate Operations (Shift Down)

- **`rra` (reverse rotate a)**: Shift **all elements of `a` down** by one.  
  _(The last element becomes the first one.)_
- **`rrb` (reverse rotate b)**: Shift **all elements of `b` down** by one.  
  _(The last element becomes the first one.)_
- **`rrr`**: Perform `rra` and `rrb` **simultaneously**.

## Visualizer

Understanding the **push_swap** algorithm can be challenging without visual
feedback.

To help with this, you can use a **GUI visualizer** to see the sorting process
in action:

- [Push Swap Visualizer by o-reo](https://github.com/o-reo/push_swap_visualizer)
- [Push Swap GUI by elijahkash](https://github.com/elijahkash/push_swap_gui)
