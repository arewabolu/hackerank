package hackerrank

/*
In this challenge, you are required to calculate and print the sum of the elements in an array,
keeping in mind that some of those integers may be quite large.

Function Description

Complete the aVeryBigSum function in the editor below. It must return the sum of all array elements.

aVeryBigSum has the following parameter(s):

    int ar[n]: an array of integers .

Return

    long: the sum of all array elements

Input Format

The first line of the input consists of an integer
.
The next line contains

space-separated integers contained in the array.

Output Format

Return the integer sum of the elements in the array.
*/

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var bigNum int64
	for _, v := range ar {
		bigNum += v
	}
	return bigNum
}
