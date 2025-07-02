/*
 * @lc app=leetcode.cn id=290 lang=golang
 * @lcpr version=30204
 *
 * [290] 单词规律
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// wordPattern 检查字符串s是否能够按照pattern给出的模式进行匹配。
// pattern: 模式字符串，每个字符代表一个单词的模式。
// s: 待匹配的字符串，单词之间以空格分隔。
// 返回值: 如果s能够按照pattern的模式进行匹配，则返回true，否则返回false。
func wordPattern(pattern string, s string) bool {
    // 将字符串s按空格分割成单词列表。
    words := strings.Split(s, " ")
    // 如果pattern的长度与单词列表的长度不一致，则无法完全匹配，返回false。
    if len(pattern) != len(words) {
        return false
    }
    // patterntoWord用于存储pattern字符到单词的映射。
    patterntoWord := make(map[byte]string)
    // wordtoPattern用于存储单词到pattern字符的映射。
    wordtoPattern := make(map[string]byte)
    // 遍历pattern中的每个字符和对应的单词。
    for i := 0; i < len(pattern); i++ {
        p := pattern[i]
        word := words[i]
        // 如果当前pattern字符已经在patterntoWord中有映射，
        // 且映射的单词与当前单词不一致，则匹配失败，返回false。
        if w, ok := patterntoWord[p]; ok {
            if w != word {
                return false
            }
        } else {
            // 如果当前pattern字符在patterntoWord中没有映射，则添加当前字符和单词的映射。
            patterntoWord[p] = word
        }
        // 如果当前单词已经在wordtoPattern中有映射，
        // 且映射的pattern字符与当前字符不一致，则匹配失败，返回false。
        if p2, ok := wordtoPattern[word]; ok {
            if p2 != p {
                return false
            }
        } else {
            // 如果当前单词在wordtoPattern中没有映射，则添加当前单词和pattern字符的映射。
            wordtoPattern[word] = p
        }
    }
    // 所有字符和单词都能一一匹配，返回true。
    return true
}

// @lc code=end

/*
// @lcpr case=start
// "abba"\n"dog cat cat dog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dog cat cat fish"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n"dog cat cat dog"\n
// @lcpr case=end

*/

