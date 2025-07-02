/*
 * @lc app=leetcode.cn id=3186 lang=golang
 * @lcpr version=30204
 *
 * [3186] 施咒的最大总伤害
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// maximumTotalDamage 计算给定力量数组中可以造成的最大总伤害。
// 参数:
//   power - 一个整数数组，表示每个个体的力量值。
// 返回值:
//   这些力量值可以造成的最大总伤害。
func maximumTotalDamage(power []int) int64 {
    // 使用 map 统计每个力量值的出现次数。
    count := map[int]int{}
    for _, p := range power {
        count[p]++
    }

    // 将 map 中的键（力量值）提取到数组 arr 中。
    arr := []int{}
    for k := range count {
        arr = append(arr, k)
    }

    // 对 arr 进行排序，以便后续动态规划处理。
    sort.Ints(arr)
    n := len(arr)

    // 初始化动态规划数组 f，用于存储最大伤害值。
    f := make([]int64, n+1)
    j := 0

    // 动态规划计算最大伤害值。
    for i, v := range arr {
        f[i+1] = f[i]
        // 找到不冲突的最大索引 j。
        for arr[j] < v-2 {
            j++
        }
        // 更新 f[i+1] 为当前最大值与 f[j] + 当前力量值 * 出现次数的较大值。
        f[i+1] = max(f[i+1], f[j]+int64(arr[i]*count[arr[i]]))
    }

    // 返回最大总伤害值。
    return f[n]
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,1,6,6]\n
// @lcpr case=end

*/

