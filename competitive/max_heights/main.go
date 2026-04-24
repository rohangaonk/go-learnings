package maxheights
func largestRectangleArea(heights []int) int {
	elems := len(heights)
	left := make([]int, elems)
	right := make([]int, elems)

	stack := make([]int, 0, elems)

	//build left array
	left[0] = -1
	stack = append(stack, 0)

	for i := 1; i < elems; i++ {
        top := heights[stack[len(stack) -1]]

		if len(stack) > 0 && top < heights[i]{
            stack = append(stack, i);
            left[i] = stack[len(stack) -1]
        } else {
            for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
                stack = stack[:len(stack) - 1];
            }
            if len(stack) > 0{
                left[i] = stack[len(stack) -1]
            } else {
                left[i] =  -1
            }
            stack = append(stack, i);
        }
	}

    stack = stack[:0]

    //build right array
	right[elems - 1] = elems
	stack = append(stack, elems - 1)

	for i := elems - 1; i >= 0; i-- {
        top := heights[stack[len(stack) -1]]

		if len(stack) > 0 && top < heights[i]{
            stack = append(stack, i);
            right[i] = stack[len(stack) -1];
        } else {
            for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
                stack = stack[:len(stack) - 1];
            }
            if len(stack) > 0{
                right[i] = stack[len(stack) -1]
            } else {
                right[i] = elems
            }
            stack = append(stack, i);
        }
	}
    

    ans := 0;

    for i, val := range heights {
        total := val * (right[i] - left[i] - 1)
        ans = max(ans, total)
    }

    return ans

}