package bytes

import "testing"
import "strings"

var containsSlice = func() []string {
    return []string{
        "12312312",    
        "abcsdsfw",
        "abcdefgh",
        "qereqwre",
        "gwertdsg",
        "hellowod",
        "iamgroot",
        "theiswer",
        "dg323sdf",
        "gadsewwe",
        "g42dg4t3",
        "4hre2323",
        "23eg4325",
        "13234234",
        "32dfgsdg",
        "23fgre34",
        "43rerrer",
        "hh2s2443",
        "hhwesded",
        "1swdf23d",
        "gwcdrwer",
        "bfgwertd",
        "badgwe3g",
        "lhoejyop",
    }
}()

var x = strings.Repeat("x", 8)

func IndexStringA(strs []string, str string) int {
    for i := range strs {
        if strs[i] == str {
            return i
        }
    }
    return -1
}

func IndexStringB(strs []string, str string) int {
    for i := range strs {
        if str == strs[i] {
            return i
        }
    }
    return -1
}

func IndexStringC(strs []string, str string) int {
    for i, s := range strs {
        if s == str {
            return i
        }
    }
    return -1
}

func IndexStringD(strs []string, str string) int {
    for i, s := range strs {
        if str == s {
            return i
        }
    }
    return -1
}

var r int

func BenchmarkIndexStringA_Constant(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringA(containsSlice, "xxxxxxxx")
    }
}

func BenchmarkIndexStringB_Constant(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringB(containsSlice, "xxxxxxxx")
    }
}

func BenchmarkIndexStringC_Constant(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringC(containsSlice, "xxxxxxxx")
    }
}

func BenchmarkIndexStringD_Constant(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringD(containsSlice, "xxxxxxxx")
    }
}



func BenchmarkIndexStringA_Variable(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringA(containsSlice, x)
    }
}

func BenchmarkIndexStringB_Variable(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringB(containsSlice, x)
    }
}

func BenchmarkIndexStringC_Variable(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringC(containsSlice, x)
    }
}

func BenchmarkIndexStringD_Variable(b *testing.B) {
    for i := 0; i < b.N; i++ {
        r = IndexStringD(containsSlice, x)
    }
}
