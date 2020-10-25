type WordsFrequency struct {
    words map[string]int
}


func Constructor(book []string) WordsFrequency {
    init:=WordsFrequency{words:make(map[string]int)}
    for _,v:=range book{
        init.words[v]++
    }
    return init
}


func (this *WordsFrequency) Get(word string) int {
    if _,ok:=this.words[word];ok{
        return this.words[word]
    }
    return 0
}