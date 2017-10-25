package set

type Strings map[string]bool

func NewStrings(initial ...string) Strings {
    result := make(Strings)
    for _, value := range initial {
        result[value] = true
    }

    return result
}

func (s Strings)Add(value string) {
    if s == nil {
        panic("uninitialied value")
    }

    s[value] = true
}

func (s Strings)Equal(s1 Strings) bool {
    if len(s) != len(s1) {
        return false
    }

    for value := range s1 {
        if s[value] != true {
            return false
        }
    }

    return true
}
