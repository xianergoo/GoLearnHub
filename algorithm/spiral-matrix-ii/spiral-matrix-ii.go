package algorithm

func generateMatrix(n int) [][]int {
    left, right := 0, n - 1
    top, bottom := 0, n - 1
    fornum := 0
    target_num := n * n 
    res := make([][]int, n)
    for i := 0; i < n; i ++ {
        res[i] = make([]int, n)
    }
    for fornum < target_num {
        for i := left; i < right; i ++ {
            res[top][i] = fornum
            fornum ++
        }
        for i := top; i < bottom; i ++ {
            res[right][i] = fornum
            fornum ++
        }
        for i := right; i > left; i -- {
            res[bottom][i] = fornum
            fornum ++    
        }
        for i := bottom; i > top; i -- {
            res[left][i] = fornum
            fornum ++    
        }
        left ++ 
        top ++
        right --
        bottom --
    }
    return res
}