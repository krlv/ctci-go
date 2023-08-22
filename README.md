# Cracking the Coding Interview in Go

Solving programming questions from ["Cracking the Coding Interview 6th Edition"](http://www.crackingthecodinginterview.com/)  by [Gayle Laakmann McDowell](http://www.gayle.com/). In [Go](https://golang.org/).

## Chapter 1: Arrays and Strings

| # | Problem                       | Tested        | Solved    |
|---|-------------------------------|:-------------:|:---------:|
|1.1| [Is Unique][1]                | [tests][2]    |     ✓     |
|1.2| [Check Permutation][3]        | [tests][4]    |     ✓     |
|1.3| [URLify][5]                   | [tests][6]    |     ✓     |
|1.4| [Palindrome Permutation][7]   | [tests][8]    |     ✓     |
|1.5| [One Away][9]                 | [tests][10]   |     ✓     |
|1.6| [String Compression][11]      | [tests][12]   |     ✓     |
|1.7| [Rotate Matrix][13]           | [tests][14]   |     ✓     |
|1.8| [Zero Matrix][15]             | [tests][16]   |     ✓     |
|1.9| [String Rotation][17]         | [tests][18]   |     ✓     |

## Chapter 2: Linked List

| # | Problem                       | Tested        | Solved    |
|---|-------------------------------|:-------------:|:---------:|
|2.1| [Remove Sups][19]             | [tests][20]   |     ✓     |
|2.2| [Return `k`th to Last][21]    | [tests][22]   |     ✓     |
|2.3| [Delete Middle Node][23]      | [tests][24]   |     ✓     |
|2.4| [Partition][25]               | [tests][26]   |     ✓     |
|2.5| [Sum Lists][27]               | [tests][28]   |     ✓     |
|2.6| [Palindrome][29]              | [tests][30]   |     ✓     |
|2.7| [Intersection][31]            | [tests][32]   |     ✓     |
|2.8| [Loop Detection][33]          | [tests][34]   |     ✓     |

## Chapter 3: Stacks and Queues

[Stack implementation][35] (with Linked List as a storage); [tests][36]

[Queue implementation][37] (with Linked List as a storage); [tests][38]

[Fixed (capped) Stack implementation][39] (with Linked List as a storage); [tests][40]

| #   | Problem                |   Tested    | Solved |
|-----|------------------------|:-----------:|:------:|
| 3.1 | [Three in one][41]     | [tests][42] |   ✓    |
| 3.2 | [Stack Min][43]        | [tests][44] |   ✓    |
| 3.3 | [Stack of Plates][45]  | [tests][46] |   ✓    |
| 3.4 | [Queue via Stacks][47] | [tests][48] |   ✓    |
| 3.5 | Sort Stack             |    tests    |        |
| 3.6 | Animal Shelter         |    tests    |        |

[1]:  ch01/01_is_unique.go
[2]:  ch01/01_is_uniqie_test.go
[3]:  ch01/02_check_permutation.go
[4]:  ch01/02_check_permutation_test.go
[5]:  ch01/03_urlify.go
[6]:  ch01/03_urlify_test.go
[7]:  ch01/04_palindrome_permutation.go
[8]:  ch01/04_palindrome_permutation_test.go
[9]:  ch01/05_one_away.go
[10]:  ch01/05_one_away_test.go
[11]:  ch01/06_string_compression.go
[12]:  ch01/06_string_compression_test.go
[13]:  ch01/07_rotate_matrix.go
[14]:  ch01/07_rotate_matrix_test.go
[15]:  ch01/08_zero_matrix.go
[16]:  ch01/08_zero_matrix_test.go
[17]:  ch01/09_string_rotation.go
[18]:  ch01/09_string_rotation_test.go
[19]:  ch02/01_remove_dups.go
[20]:  ch02/01_remove_dups_test.go
[21]:  ch02/02_kth_to_last.go
[22]:  ch02/02_kth_to_last_test.go
[23]:  ch02/03_delete_middle.go
[24]:  ch02/03_delete_middle_test.go
[25]:  ch02/04_partition.go
[26]:  ch02/04_partition_test.go
[27]:  ch02/05_sum_lists.go
[28]:  ch02/05_sum_lists_test.go
[29]:  ch02/06_palindrome.go
[30]:  ch02/06_palindrome_test.go
[31]:  ch02/07_intersection.go
[32]:  ch02/07_intersection_test.go
[33]:  ch02/08_loop_detection.go
[34]:  ch02/08_loop_detection_test.go
[35]:  ch03/stack.go
[36]:  ch03/stack_test.go
[37]:  ch03/queue.go
[38]:  ch03/queue_test.go
[39]:  ch03/fixed_stack.go
[40]:  ch03/fixed_stack_test.go
[41]:  ch03/01_three_in_one.go
[42]:  ch03/01_three_in_one_test.go
[43]:  ch03/02_stack_min.go
[44]:  ch03/02_stack_min_test.go
[45]:  ch03/03_stack_of_plates.go
[46]:  ch03/03_stack_of_plates_test.go
[47]:  ch03/04_my_queue.go
[48]:  ch03/04_my_queue_test.go
