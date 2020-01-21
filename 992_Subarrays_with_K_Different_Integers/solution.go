func subarraysWithKDistinct(A []int, K int) int {
    return atMostK(A, K) - atMostK(A, K-1)
}


func atMostK(A []int, K int) int {
    m := make(map[int]int)
    start, end := 0, 0
    length := len(A)
    result := 0
    for start < length && end < length {
        m[A[end]] += 1
        for len(m) > K {
            m[A[start]] -= 1
            if m[A[start]] == 0 {
                delete(m, A[start])
            }
            start++
        }
        result += (end - start + 1) // 当前位置与前面组成的最多K个元素子数组的个数
        end++
    }    
    return result
}
