type MyHashMap struct {
    Ele []int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    init:=make([]int,1000000)
    for i:=range init{
        init[i]=-1
    }
    return MyHashMap{Ele:init}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    this.Ele[key]=value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    return this.Ele[key]
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    this.Ele[key]=-1
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */