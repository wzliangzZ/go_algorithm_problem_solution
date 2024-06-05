from collections import defaultdict

#count属性表示以node结尾的串的个数
class Trie(object):
    def __init__(self):
        self.node = defaultdict(Trie)
        self.count = 0

    def insert(self, word: str):
        cur = self
        for c in word: cur = cur.node[c]
        cur.count += 1

    #删除Trie中的word
    def delete(self, word: str):
        cur = self
        for c in word:
            if not cur.node.get(c, None): return
            cur = cur.node[c]
        if cur.count > 0: cur.count -= 1


    #默认统计word在Trie里有多少前缀相同的字串
    def query(self, word: str) -> int:
        res, cur = 0, self
        for c in word:
            if not cur.node.get(c, None):
                #return 0
                return res
            cur = cur.node[c]
            res += cur.count    #前缀相同就累加
        #return cur.count
        return res