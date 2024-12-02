package array

/*
ReverseInPlace checks the start and end indices. If the check and end indices are
valid, then iterate through the array and reverse the integers between start and
end indices.

time = O(n), space = O(1)
*/
func ReverseInPlace(list []int, start, end int) {
	// list = [1, 2, 3, 4, 5]
	// start, end = 1, 2
	// expected(list) = [1, 3, 2, 4, 5]
	// swap ints within i from start to (end-i)+start
	// could further constrain i from start to (start+end)/2
	for i := start; i < end-i+start && i <= (start+end)/2; i++ {
		list[i], list[end-i+start] = list[end-i+start], list[i]
	}
}
