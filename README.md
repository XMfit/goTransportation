# PLC Project 2
Professor Smith (remember Professor Smith?) is transporting the manatees to green “pastures” where the supply of seagrass is better. The FIT Biology department has a special boat for transporting manatees, basically the boat has two long tubs the width of a manatee. The tubs are located on either side of the boat and have the same length. The manatees are packed "nose-to-tail" in the two tubs. After all the manatees have entered through the aft gates, the boat goes to the destination, and the fore gates are opened releasing the manatees.

The manatees are herded into a single queue, and Professor Smith coaxes each manatee in turn onto the port (left) or starboard (right) tub of the boat. The manatees in the queue have different lengths, which the graduate students in the water have determined. Based on the lengths Professor Smith decides which tub each manatee should board, and boards as many manatees as possible from the queue, subject to the length limit of the tubs. Your job is to write a program that will tell Professor Smith how to load the manatees in order so as to maximize the number of manatees loaded. No manatees in the queue can be skipped, but some manatees may have to be left behind for the next trip. 

# Input / Output
All input comes from the standard input stream and all output is to be written to the standard output stream.

The first line of input contains a single integer between 1 and 100: the length of the tubs (in meters). The queue of manatees follow. For each manatee in the queue there is a line of input specifying the length of the manatee (in centimeters, an integer between 100 and 3000 inclusive). A final line of input contains the integer 0 marking the end of the input. There can be any number of manatees in the queue (including none). The manatees must be loaded in order, subject to the constraint that the total length of manatees on either side does not exceed the length of the tubs. Subject to this constraint as many manatees should be loaded as possible, starting with the first manatee in the queue and loading manatees in order until no more can be loaded.

The first line of output should give the number of manatees that can be transported. For each manatee that can be loaded onto the boat, in the order the manatees appear in the input, output a line containing the word `port` if the manatee is to be directed to the port side and `starboard` if the manatee is to be directed to the starboard side. If several arrangements of the manatees meet the criteria above, any one will do. 

### Input

```
50
2500
3000
1000
1000
1500
700
800
0
```

### Output
```
6
port
starboard
starboard
starboard
port
port
```