package piscine

func Index(s string, toFind string) int {
    a := -1
    if len(s) == 0 || len(toFind) == 0 {
        return 0
    }
    for i, c := range []rune(s) {
        if c == rune(toFind[0]) {
            a := i
            for j, v := range []rune(toFind) {
                if v != rune(s[a+j]) {
                    a = -1
                }
            }
            if a != -1 {
                return a
            }
        }
    }
    return a
}